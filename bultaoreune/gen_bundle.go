package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// custom bundle unmarshal to parse into typed resources
// parts before and after generated switch statement
func writeUnmarshalBundle(domainResources []string, generateDirVer string) {
	fmt.Println("adding custom marshal to bundle, to use resource entry structs")
	//replace resource from spec with 2:
	//one to take raw json and one to hold any typed Account, AllergyIntolerance, etc
	bp := path.Join(generateDirVer, "Bundle.go")
	file, _ := os.Open(bp)
	var outputLines []string
	outputLines = append(outputLines, "//Note: Bundle is special case in gen_bundle.go, to handle entry resources with types")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "Resource") &&
			strings.Contains(trimmed, "*Resource") &&
			strings.Contains(trimmed, "`json:\"resource,omitempty\"`") {
			outputLines = append(outputLines, "UntypedResource   *json.RawMessage     `json:\"resource,omitempty\"`")
			outputLines = append(outputLines, "Resource          any                  `json:\"-\"`")
		} else {
			outputLines = append(outputLines, line)
		}
	}
	file.Close()
	os.WriteFile(bp, []byte(strings.Join(outputLines, "\n")), 0644)

	before := `
	type OtherBundle Bundle

	// per Lukas Schmierer, check resourceType string in raw resource json,
	// then use string in switch to parse into correct resource,
	// and user can use resource with type assertion
	func (b *Bundle) UnmarshalJSON(data []byte) error {
		var tmp struct {
			OtherBundle
		}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		*b = Bundle(tmp.OtherBundle)

		for i := range b.Entry {
			e := &b.Entry[i] //weird looking because need to modify entry at i, not copy of i
			if e.UntypedResource == nil {
				return errors.New("Bundle UnmarshalJSON got nil resource")
			}

			var peek struct {` + "ResourceType string `json:\"resourceType\"`" + `
			}
			if err := json.Unmarshal(*e.UntypedResource, &peek); err != nil {
				return errors.New("Bundle UnmarshalJSON failed to get resource type from bundle entry" + string(*e.UntypedResource))
			}

			switch peek.ResourceType {`
	after := `default:
				return errors.New("Bundle UnmarshalJSON switch default, no resource type" + peek.ResourceType)
			}
		}
		return nil
	}
		
	// MarshalJSON copies Resource fields to UntypedResource before marshaling
	func (b *Bundle) MarshalJSON() ([]byte, error) {
		tmp := struct {
			*OtherBundle
		}{
			OtherBundle: (*OtherBundle)(b),
		}

		// Convert Resource to UntypedResource for each entry
		for i := range tmp.Entry {
			e := &tmp.Entry[i]
			if e.Resource != nil {
				resourceBytes, err := json.Marshal(e.Resource)
				if err != nil {
					return nil, err
				}
				rawMsg := json.RawMessage(resourceBytes)
				e.UntypedResource = &rawMsg
			}
		}
		return json.Marshal(tmp.OtherBundle)
	}`
	var sb strings.Builder
	sb.WriteString(before)
	for _, dr := range domainResources {
		sb.WriteString(`case "`)
		sb.WriteString(dr)
		sb.WriteString(`":
			var res `)
		sb.WriteString(dr)
		sb.WriteString(`
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		`)
	}
	sb.WriteString(after)
	//add unmarshal method with cases for each resource in middle
	bf, _ := os.OpenFile(bp, os.O_APPEND|os.O_WRONLY, 0644)
	bf.WriteString(sb.String())
	bf.Close()
}

func WriteResourceGroup(domainResouces []string, generateClientDirVer, fhirVersion string) {
	var sb strings.Builder
	sb.WriteString(`
	package ` + fhirVersion + `Client

import ` + fhirVersion + ` "github.com/PotatoEMR/simple-fhir-client/` + fhirVersion + `"

type ResourceGroup struct {`)
	for _, dr := range domainResouces {
		sb.WriteString(dr + "_list []*" + fhirVersion + "." + dr + "\n")
	}
	sb.WriteString("}")
	f, _ := os.Create(path.Join(generateClientDirVer, "resourceGroup.go"))
	f.WriteString(sb.String())
	f.Close()
}
