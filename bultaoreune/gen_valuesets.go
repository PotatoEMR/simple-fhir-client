package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
)

//each coding in a valueset has the same system
//so it'd save space to not add system field to each array item but whatever

type Peek struct {
	ResourceType string `json:"resourceType,omitempty"`
}

// some of the valuesets in valueset.json are codesystems, others are valuesets
// peek to check which one, but then just ignoring difference and making both into valueset
type CodeSystem struct {
	Name    *string             `json:"name,omitempty"`
	Url     *string             `json:"url,omitempty"`
	Concept []CodeSystemConcept `json:"concept,omitempty"`
}

type CodeSystemConcept struct {
	Code     string              `json:"code"`
	Display  *string             `json:"display,omitempty"`
	Contains []CodeSystemConcept `json:"contains,omitempty"`
}

// some of the valuesets in valueset.json are codesystems, others are valuesets
type ValueSet struct {
	Contained []Resource         `json:"contained,omitempty"`
	Url       *string            `json:"url,omitempty"`
	Compose   *ValueSetCompose   `json:"compose,omitempty"`
	Expansion *ValueSetExpansion `json:"expansion,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Include []ValueSetComposeInclude `json:"include"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeInclude struct {
	System   *string                         `json:"system,omitempty"`
	Concept  []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter   []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet []string                        `json:"valueSet,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeConcept struct {
	Code        string                                     `json:"code"`
	Display     *string                                    `json:"display,omitempty"`
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeConceptDesignation struct {
	Value string `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeFilter struct {
	Value string `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansion struct {
	Id         *string                      `json:"id,omitempty"`
	Identifier *string                      `json:"identifier,omitempty"`
	Timestamp  string                       `json:"timestamp"`
	Total      *int                         `json:"total,omitempty"`
	Offset     *int                         `json:"offset,omitempty"`
	Parameter  []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains   []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansionParameter struct {
	Id            *string  `json:"id,omitempty"`
	Name          string   `json:"name"`
	ValueString   *string  `json:"valueString"`
	ValueBoolean  *bool    `json:"valueBoolean"`
	ValueInteger  *int     `json:"valueInteger"`
	ValueDecimal  *float64 `json:"valueDecimal"`
	ValueUri      *string  `json:"valueUri"`
	ValueCode     *string  `json:"valueCode"`
	ValueDateTime *string  `json:"valueDateTime"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansionContains struct {
	Id       *string `json:"id,omitempty"`
	System   *string `json:"system,omitempty"`
	Abstract *bool   `json:"abstract,omitempty"`
	Inactive *bool   `json:"inactive,omitempty"`
	Version  *string `json:"version,omitempty"`
	Code     *string `json:"code,omitempty"`
	Display  *string `json:"display,omitempty"`
}

func GenValuesets(spec_vs, generateVSDirVer, fhirVersion string) []string {
	created := []string{}
	fmt.Printf("make valuesets from %s \n", spec_vs)
	spec, err := os.Open(spec_vs)
	if err != nil {
		fmt.Printf("GenValuesets did not find %s, probably messed up download\n", spec_vs)
	}
	defer spec.Close()
	var bundle Bundle
	if err := json.NewDecoder(spec).Decode(&bundle); err != nil {
		panic(err)
	}

	var sb strings.Builder
	sb.WriteString("package " + fhirVersion + "\n\n")

	// Maps to track string variables: Go variable name -> original FHIR value
	stringVars := make(map[string]string)

	var haveVsAlready []string

	// First pass: collect all unique strings and create variables
	for _, entry := range bundle.Entry {
		var peek Peek
		json.Unmarshal(entry.Resource, &peek)
		if peek.ResourceType == "CodeSystem" {
			var cs CodeSystem
			json.Unmarshal(entry.Resource, &cs)
			if cs.Url != nil && len(cs.Concept) > 0 {
				var varname = "s" + urlToVarname(*cs.Url)
				if _, exists := stringVars[varname]; !exists {
					stringVars[varname] = *cs.Url
				}
				for _, concept := range cs.Concept {
					cname := sanitizeVarName(concept.Code)
					if _, exists := stringVars[cname]; !exists {
						stringVars[cname] = concept.Code
					}
					if concept.Display != nil {
						dname := sanitizeVarName(*concept.Display)
						if _, exists := stringVars[dname]; !exists {
							stringVars[dname] = *concept.Display
						}
					}
				}
			}
		} else if peek.ResourceType == "ValueSet" {
			var vs ValueSet
			json.Unmarshal(entry.Resource, &vs)
			if vs.Compose != nil && len(vs.Compose.Include) > 0 && len(vs.Compose.Include[0].Concept) > 0 {
				if vs.Url != nil {
					var varname = "s" + urlToVarname(*vs.Url)
					if _, exists := stringVars[varname]; !exists {
						stringVars[varname] = *vs.Url
					}
				}
				for _, inc := range vs.Compose.Include {
					for _, concept := range inc.Concept {
						cname := sanitizeVarName(concept.Code)
						if _, exists := stringVars[cname]; !exists {
							stringVars[cname] = concept.Code
						}
						if concept.Display != nil {
							dname := sanitizeVarName(*concept.Display)
							if _, exists := stringVars[dname]; !exists {
								stringVars[dname] = *concept.Display
							}
						}
					}
				}
			}
		}
	}

	// Write string variables
	for varName, originalValue := range stringVars {
		escaped := strings.ReplaceAll(originalValue, `"`, `\"`)
		escaped = strings.ReplaceAll(escaped, "\n", " ")
		sb.WriteString("var " + varName + " = \"" + escaped + "\"\n")
	}
	sb.WriteString("\n")

	// Second pass: generate valuesets using references
	for _, entry := range bundle.Entry {
		var peek Peek
		json.Unmarshal(entry.Resource, &peek)
		if peek.ResourceType == "CodeSystem" {
			var cs CodeSystem
			json.Unmarshal(entry.Resource, &cs)
			varname := "s" + urlToVarname(*cs.Url)
			numCodings := len(cs.Concept)
			if numCodings != 0 {
				haveVsAlready = append(haveVsAlready, varname)
				created = append(created, varname[1:])
				sb.WriteString("var VS" + varname[1:] + " = []Coding{\n")
				for _, concept := range cs.Concept {
					if cs.Name != nil && cs.Url != nil {
						systemVar := varname
						codeVar := sanitizeVarName(concept.Code)
						displayRef := "nil"
						if concept.Display != nil {
							displayVar := sanitizeVarName(*concept.Display)
							displayRef = "&" + displayVar
						}
						sb.WriteString(`	{
		System:  &` + systemVar + `,
		Code:    &` + codeVar + `,
		Display: ` + displayRef + `,
	},
`)
					}
				}
				sb.WriteString("}\n\n")
			}
		} else if peek.ResourceType == "ValueSet" {
			var vs ValueSet
			json.Unmarshal(entry.Resource, &vs)
			if vs.Compose != nil && len(vs.Compose.Include) > 0 && len(vs.Compose.Include[0].Concept) > 0 {
				varname := "s" + urlToVarname(*vs.Url)
				if !slices.Contains(haveVsAlready, varname) {
					created = append(created, varname[1:])
					sb.WriteString("var VS" + varname[1:] + " = []Coding{\n")
					for _, inc := range vs.Compose.Include {
						for _, concept := range inc.Concept {
							systemVar := varname
							codeVar := sanitizeVarName(concept.Code)
							displayRef := "nil"
							if concept.Display != nil {
								displayVar := sanitizeVarName(*concept.Display)
								displayRef = "&" + displayVar
							}
							sb.WriteString(`	{
		System:  &` + systemVar + `,
		Code:    &` + codeVar + `,
		Display: ` + displayRef + `,
	},
`)
						}
					}
					sb.WriteString("}\n\n")
				}
			}
		}
	}

	searchParamFile, _ := os.Create(path.Join(generateVSDirVer, "zzzValueSets.go"))
	fmt.Fprint(searchParamFile, sb.String())
	searchParamFile.Close()
	return created
}

func urlToVarname(url string) string {
	if i := strings.Index(url, "|"); i >= 0 {
		url = url[:i]
	}
	parts := strings.SplitN(url, "://", 2)
	if len(parts) == 1 {
		parts = []string{"", url}
	}
	last := parts[1][strings.LastIndex(parts[1], "/")+1:]
	ret := strings.Title(strings.NewReplacer("-", "_", ":", "_", ".", "_").Replace(last))
	return ret
}

// ai idk whatever
func sanitizeVarName(s string) string {
	// Remove/replace invalid characters for Go variable names
	reg := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	sanitized := reg.ReplaceAllString(s, "_")

	// Ensure it starts with a letter or underscore
	if len(sanitized) > 0 && sanitized[0] >= '0' && sanitized[0] <= '9' {
		sanitized = "_" + sanitized
	}

	// Limit length to reasonable size
	if len(sanitized) > 50 {
		sanitized = sanitized[:50]
	}

	return "s" + sanitized
}
