/*
generates structs, client, forms. usage: go run ./bulteaoreune
...honestly codegen became a big ball of mud, very messy and hard to read or modify easily
kind of based on these but their approach to golang codegen might cleaner:
https://github.com/DAMEDIC/fhir-toolbox-go/ (DAMEDIC)
https://github.com/samply/golang-fhir-models (Samply)
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

var (
	fhirVersions = []string{"r5", "r4", "r4b"}
	fhirURL      = "https://www.hl7.org/fhir" //link to download fhir spec
	extractDir   = "download"                 //dir we download fhir spec to
	generateDir  = ""                         //dir we generate into, currently base dir
	zipFileNames = []string{
		"profiles-resources.json",
		"profiles-types.json",
		"valuesets.json",
	}
)

func main() {
	for _, v := range fhirVersions {
		os.RemoveAll(v)
		os.RemoveAll(v + "Client")
	}

	//get spec json file from fhir site, if run without -nodownload
	// region Download
	nodownload := flag.Bool("nodownload", false, "download FHIR definitions")
	r4 := flag.Bool("r4", false, "r4 only")
	r4b := flag.Bool("r4b", false, "r4b only")
	r5 := flag.Bool("r5", false, "r5 only")
	flag.Parse()
	if *r4 {
		fhirVersions = []string{"r4"}
	} else if *r4b {
		fhirVersions = []string{"r4b"}
	} else if *r5 {
		fhirVersions = []string{"r5"}
	}

	fmt.Println("🔥🔥🔥 bultaoreune")
	fmt.Println("fhir version", strings.Join(fhirVersions, ","))

	if !*nodownload {
		os.RemoveAll(extractDir)
	}

	for _, fhirVersion := range fhirVersions {
		var header strings.Builder
		header.WriteString("//generated with command go run ./bultaoreune\n")
		header.WriteString(fmt.Sprintf("//inputs %s/[%s]\n", fhirURL+"/"+fhirVersion, strings.Join(zipFileNames, " ")))
		header.WriteString("//for details see https://github.com/PotatoEMR/simple-fhir-client\n\n")

		extractDirVer := path.Join(extractDir, fhirVersion)
		if *nodownload {
			fmt.Println("-nodownload flag on (turn off to download FHIR definitions)")
		} else {
			os.MkdirAll(extractDirVer, 0755)
			fmt.Printf("downloading files from %s into %s\n", fhirURL+fhirVersion, extractDirVer)
			fmt.Println("use -nodownload flag to skip FHIR definition download")
			for _, f := range zipFileNames {
				if r, e := http.Get(fhirURL + "/" + fhirVersion + "/" + f); e != nil {
					panic(e)
				} else {
					o, _ := os.Create(filepath.Join(extractDirVer, f))
					io.Copy(o, r.Body)
					r.Body.Close()
					o.Close()
				}
			}
		}
		generateDirVer := path.Join(generateDir, fhirVersion)
		generateClientDirVer := path.Join(generateDir, fhirVersion+"Client")
		os.MkdirAll(generateDirVer, 0755)
		os.MkdirAll(generateClientDirVer, 0755)
		//endregion

		//these are all valuesets, regardless of if field binding is required or just example
		//maybe bad for compile time to include in resources package?
		//so the full list goes into their own package, required stay in resources for use in form
		valueset_list := GenValuesets(path.Join(extractDirVer, "valuesets.json"), generateDirVer, fhirVersion)

		//write structs
		fileToStructs(path.Join(extractDirVer, "profiles-types.json"), generateDirVer, fhirVersion, valueset_list, header.String(), false)
		//write resources
		fileToStructs(path.Join(extractDirVer, "profiles-resources.json"), generateDirVer, fhirVersion, valueset_list, header.String(), true)
		fmt.Println("created resources in", generateDirVer)
		fileToClient(path.Join(extractDirVer, "profiles-resources.json"), generateDirVer, fhirVersion, generateClientDirVer, header.String())
		fmt.Println("created client in", generateClientDirVer)

		filename := filepath.Join(generateDirVer, "zzzFieldComponents.templ")
		file, _ := os.Create(filename)
		fmt.Fprintln(file, FormText(fhirVersion))

		timeFormat := filepath.Join(generateDirVer, "zzzTimeFormat.go")
		tf, _ := os.Create(timeFormat)
		fmt.Fprintln(tf, "package "+fhirVersion+`	
const DateFormat string = "2006-01-02"
const DateTimeFormat string = "2006-01-02T15:04:05-07:00"`)
	}
	fmt.Println("wrote files, running templ generate for form components")
	cmd := exec.Command("go", "tool", "templ", "generate")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(string(out))

	fmt.Println("finished writing files, formatting with goimports")
	exec.Command("goimports", "-w", ".").Run()
}

// region Profiles Parsing
// follows resource structure, used to parse .json
type Bundle struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Resource json.RawMessage `json:"resource"`
}

type Resource struct {
	Rest           []Rest   `json:"rest"`
	Snapshot       Snapshot `json:"snapshot"`
	ResourceType   string   `json:"resourceType"`
	Name           string   `json:"name"`
	Url            string   `json:"url"`
	Kind           string   `json:"kind"`
	BaseDefinition string   `json:"baseDefinition"`
}

type Rest struct {
	Resource []RestResource `json:"resource"`
}

type RestResource struct {
	Type             string        `json:"type"`
	SearchInclude    []string      `json:"searchInclude"`
	SearchRevInclude []string      `json:"searchRevInclude"`
	SearchParam      []SearchParam `json:"searchParam"`
}

type SearchParam struct {
	Name string `json:"name"`
}

type Snapshot struct {
	Element []Element `json:"element"`
}
type Element struct {
	Path    string  `json:"path"`
	Min     int     `json:"min,omitempty"`
	Max     string  `json:"max,omitempty"`
	Type    []Type  `json:"type,omitempty"`
	Binding Binding `json:"binding,omitempty"`
}

type Type struct {
	Code          string   `json:"code"`
	TargetProfile []string `json:"targetProfile"`
}

type Binding struct {
	Strength string `json:"strength,omitempty"`
	ValueSet string `json:"valueSet,omitempty"`
}

// endregion
func fileToStructs(spec_file, generateStructsDir, fhirVersion string, valueset_list []string, header string, isDR bool) {
	fmt.Printf("parse %s into structs\n", spec_file)
	spec, err := os.Open(spec_file)
	if err != nil {
		fmt.Printf("did not find %s, probably messed up download\n", spec_file)
	}
	defer spec.Close()
	resStrs, _ := ResourceStrings()
	var bundle Bundle
	if err := json.NewDecoder(spec).Decode(&bundle); err != nil {
		panic(err)
	}

	//region Structs

	for _, entry := range bundle.Entry {
		var res Resource
		if err := json.Unmarshal(entry.Resource, &res); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		if !(res.Kind == "complex-type" || res.Kind == "resource") {
			//primitive are just 1 field types so don't make struct
			//also don't make structs for random junk resources like Base FHIR Capability Statement
			continue
		}
		if res.Name == "Base" || res.Name == "BackboneElement" {
			continue //we use other struct instead of backbone element, we don't use base
		}

		var formFuncs strings.Builder

		isDomainResource := false
		if res.BaseDefinition == "http://hl7.org/fhir/StructureDefinition/DomainResource" {
			isDomainResource = true
			//this resource has json serializer to add resourceType on json.marshal
			//also adding forms input field funcs as we go through
		}

		identifierType := "none" //for reference function

		pathCard := make(map[string]string)
		//just for forms, to know if backbone elements in field path are [] or not
		for _, elt := range res.Snapshot.Element {
			pathCard[elt.Path] = elt.Max
		}

		//map of structs -> fields needed to list out all the BackboneElements
		//because they're subcomponents eg Allergy -> Reaction
		//and rather than nested, we want to write structs one after the other
		structs := make(map[string][]Element)
		structs[res.Name] = []Element{}
		//we want to write structs in order of parsing backbone elements, but map order random
		structInsertionOrder := []string{res.Name}
		//put elts into name of current struct, eg AllergyIntolerance, AllergyIntoleranceReaction...
		for _, elt := range res.Snapshot.Element {
			pathParts := strings.Split(elt.Path, ".")
			fieldPathMinusLast := ""
			for _, pp := range pathParts[:len(pathParts)-1] {
				fieldPathMinusLast = fieldPathMinusLast + strings.Title(pp)
			}
			fieldPath := fieldPathMinusLast + strings.Title(pathParts[len(pathParts)-1])
			if len(elt.Type) > 0 && elt.Type[0].Code == "BackboneElement" {
				structs[fieldPath] = []Element{}
				structInsertionOrder = append(structInsertionOrder, fieldPath)
				structs[fieldPathMinusLast] = append(structs[fieldPathMinusLast], elt)
			} else {
				structs[fieldPathMinusLast] = append(structs[fieldPathMinusLast], elt)
			}
		}

		var sb strings.Builder
		// Generate go struct using parsed element fields

		if isDR {
			sb.WriteString("import \"github.com/a-h/templ\"\n")
		} //there are 2 domain resource bools that do slightly different things or the same thing idk probably terrible either way will fix it later
		sb.WriteString("\n")

		for _, s_name := range structInsertionOrder {
			sb.WriteString(fmt.Sprintf("// %s\n", strings.Replace(res.Url, "hl7.org/fhir", "hl7.org/fhir/"+fhirVersion, 1)))
			sb.WriteString(fmt.Sprintf("type %s struct {\n", s_name))
			s_elts := structs[s_name]
			for _, elt := range s_elts {
				// getting last part of path eg allergy.reaction.severity -> severity
				parts := strings.Split(elt.Path, ".")
				fieldName_lower := parts[len(parts)-1]
				fieldName := strings.Title(fieldName_lower)

				if elt.Path == res.Name+".identifier" {
					identifierType = elt.Max //1 or *
				}

				//region Forms
				//take a break from structs to write form methods...kind of mixed up with structs
				//dont think you want a form for one of these...make it yourself if you do
				if !slices.Contains([]string{"id", "meta", "implicitRules", "language", "contained", "extension", "modifierExtension", "identifier"}, fieldName_lower) {
					for _, eltType := range elt.Type {
						eltGoType, ok := fhirPrimitive_to_GolangType[eltType.Code]
						//currently only forms for primitive types and coding/codeableconcept/annotation
						if ok || eltType.Code == "Coding" || eltType.Code == "CodeableConcept" || eltType.Code == "Annotation" {
							optionsValueSet := ""
							vsParam := ""
							vsReq := ""
							if eltType.Code == "code" || eltType.Code == "Coding" || eltType.Code == "CodeableConcept" {
								optionsValueSet = ", optionsValueSet"
								varname := urlToVarname(elt.Binding.ValueSet)
								if elt.Binding.Strength == "required" && slices.Contains(valueset_list, varname) {
									vsReq = "optionsValueSet := " + "VS" + varname + "\n"
								} else {
									vsParam = "optionsValueSet []Coding, "
								}
							}
							//forms
							caps := []string{}
							for _, s := range parts[1:] {
								caps = append(caps, strings.Title(s))
							}
							backbonePath := "resource" //this one actually used in field
							intParams := ""
							bbCheck := ""
							formName := "" //this one used in submitted json

							//need to check if each backbone element along the way is cardinality * or not to make [n] or not
							pathString := res.Name
							for i, p := range parts[1:] {
								pUpper := strings.Title(p)
								pUpperNum := "num" + pUpper
								pathString = pathString + "." + p
								if i != 0 {
									formName = formName + "."
								}
								if pathCard[pathString] == "*" {
									intParams = intParams + pUpperNum + " int, "
									bbCheck = bbCheck + " || " + pUpperNum + ">= len(" + backbonePath + "." + pUpper + ")"
									backbonePath = backbonePath + "." + pUpper + "[" + pUpperNum + "]"
									formName = formName + p + "[\"+strconv.Itoa(" + pUpperNum + ")+\"]"
								} else {
									backbonePath = backbonePath + "." + pUpper
									formName = formName + p
								}
							}

							funcName := strings.Join(caps, "")
							if len(elt.Type) > 1 {
								//choice properties end in [x], and have multiple types
								//they can be one of multiple types, we don't bother enforcing that though
								//just provide fields Address_url Address_string Address_ContactPoint etc
								funcName = funcName[:len(funcName)-3] + strings.Title(eltType.Code)
								backbonePath = backbonePath[:len(backbonePath)-3] + strings.Title(eltType.Code)
								formName = formName[:len(formName)-3] + strings.Title(eltType.Code)
							}
							formFuncs.WriteString(`func ` + `(resource *` + res.Name + ") T_" + funcName + "(" + intParams + vsParam + `htmlAttrs templ.Attributes) templ.Component {` + vsReq + `
								        if resource == nil ` + bbCheck + `{`)
							if elt.Max == "*" || elt.Min == 1 {
								backbonePath = "&" + backbonePath
							}

							inputType := "StringInput"
							if eltGoType == "bool" {
								inputType = "BoolInput"
							} else if eltGoType == "float64" || eltGoType == "int" || eltGoType == "int64" {
								inputType = strings.Title(eltGoType) + "Input"
							} else if eltType.Code == "date" || eltType.Code == "dateTime" {
								inputType = strings.Title(eltType.Code) + "Input"
							} else if eltType.Code == "code" || eltType.Code == "Coding" || eltType.Code == "CodeableConcept" {
								inputType = strings.Title(eltType.Code + "Select")
							} else if eltType.Code == "Annotation" {
								inputType = "AnnotationTextArea"
							}

							formFuncs.WriteString(`
					        return ` + inputType + `("` + formName + `", nil` + optionsValueSet + `, htmlAttrs)
					        }
					    return ` + inputType + `("` + formName + `", ` + backbonePath + optionsValueSet + `, htmlAttrs)
					    }
						`)
						}
					}
				}
				//endregion
				//done with form methods, back to writing struct fields
				//region Structs again

				if len(elt.Type) < 1 {
					//some fields eg AuditEvent.entity.agent have no type
					//think they link back to other field in resource, outside backboneelement?
					//not sure what to do with these, just skipping
					continue
				}

				//depending on cardinality make field one of slice, value, or pointer
				c := "[]"
				if elt.Max != "*" {
					if elt.Min == 0 {
						c = "*"
					} else {
						c = ""
					}
				}
				omitempty := ",omitempty"
				if elt.Min == 1 {
					omitempty = ""
				}
				for _, t := range elt.Type {
					ft := t.Code
					gt, ok := fhirPrimitive_to_GolangType[ft]
					if ok {
						ft = gt
					}
					if ft == "BackboneElement" {
						ft = ""
						for _, p := range parts {
							ft = ft + strings.Title(p)
						}
					}
					f := fieldName
					fl := fieldName_lower
					if len(elt.Type) > 1 {
						//choice properties end in [x], and have multiple types
						//they can be one of multiple types, we don't bother enforcing that though
						//just provide fields Address_url Address_string Address_ContactPoint etc
						f = f[:len(f)-3] + strings.Title(t.Code)
						fl = fieldName_lower[:len(fieldName_lower)-3] + strings.Title(t.Code)
					}
					jsonEnc := omitempty
					/*if t.Code == "date" {
						jsonEnc = jsonEnc + ",format:'2006-01-02'"
					} else if t.Code == "dateTime" {
						jsonEnc = jsonEnc + ",format:'2006-01-02T15:04:05Z07:00'"
					}*/

					sb.WriteString(fmt.Sprintf("%s %s%s `json:\"%s%s\"`\n", f, c, ft, fl, jsonEnc))
				}
			}
			//add string representation if struct has one in resourceStrings.go
			sb.WriteString("}\n")
			if s, ok := resStrs[s_name]; ok {
				sb.WriteString(fmt.Sprintf("%s\n", s))
			}
			sb.WriteString("\n")
		}

		// domain resources convert to json and get search params
		if isDomainResource {
			sb.WriteString(fmt.Sprintf(`
				type Other%s %s
				// on convert struct to json, automatically add resourceType=%s
				func (r %s) MarshalJSON() ([]byte, error) {
					return json.Marshal(struct {
						Other%s`, res.Name, res.Name, res.Name, res.Name, res.Name))
			sb.WriteString("\nResourceType string `json:\"resourceType\"`")
			sb.WriteString(fmt.Sprintf(`
							}{
						Other%s: Other%s(r),
						ResourceType:            "%s",
					})
				}
				`, res.Name, res.Name, res.Name))

			refFromIdentifier := ""
			if identifierType == "1" {
				refFromIdentifier = `ref.Identifier = r.Identifier`
			} else if identifierType == "*" {
				refFromIdentifier = `if len(r.Identifier) != 0 {
								ref.Identifier = &r.Identifier[0]
							}`
			}
			sb.WriteString(`func (r ` + res.Name + `) ToRef() Reference {
							var ref Reference
							if r.Id != nil {
								refStr := "` + res.Name + `/" + *r.Id
								ref.Reference = &refStr
							}
							` + refFromIdentifier + `
							rtype := "` + res.Name + `"
							ref.Type = &rtype
							//rDisplay := r.String()
							//ref.Display = &rDisplay
							return ref
						}`)
		}
		filename := filepath.Join(generateStructsDir, res.Name+".go")
		file, err := os.Create(filename)
		if err != nil {
			panic("Error creating file:" + err.Error())
		}
		fmt.Fprintf(file, "package %s\n\n", fhirVersion)
		fmt.Fprint(file, header)
		fmt.Fprintln(file, sb.String())
		if isDR {
			fmt.Fprintln(file, formFuncs.String())
		}
		file.Close()
	}
	//endregion
}

// region Primitive -> Go
// if struct field is a fhir primitive, map it to appropriate go type
var fhirPrimitive_to_GolangType = map[string]string{
	"base64Binary":                          "string",
	"boolean":                               "bool",
	"canonical":                             "string",
	"code":                                  "string",
	"date":                                  "string",
	"dateTime":                              "string",
	"decimal":                               "float64",
	"id":                                    "string",
	"instant":                               "string",
	"integer":                               "int",
	"integer64":                             "int64",
	"markdown":                              "string",
	"oid":                                   "string",
	"positiveInt":                           "int",
	"string":                                "string",
	"time":                                  "string",
	"unsignedInt":                           "int",
	"uri":                                   "string",
	"url":                                   "string",
	"uuid":                                  "string",
	"xhtml":                                 "string",
	"http://hl7.org/fhirpath/System.String": "string",
}

//endregion

// region Client
func fileToClient(spec_file, generateDirVer, fhirVersion, generateClientDirVer, header string) {
	fmt.Printf("parse %s into client funcs and searchparams\n", spec_file)
	spec, err := os.Open(spec_file)
	if err != nil {
		fmt.Printf("did not find %s, probably messed up download\n", spec_file)
	}
	defer spec.Close()
	var bundle Bundle
	if err := json.NewDecoder(spec).Decode(&bundle); err != nil {
		panic(err)
	}

	domainResources := []string{}
	// each domain resource has search params, maybe belongs in client more than resource itself
	writtenSP := false //writing search params from capabilityStatment but it appears twice, second empty
	for _, entry := range bundle.Entry {
		var res Resource
		if err := json.Unmarshal(entry.Resource, &res); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		if res.ResourceType == "CapabilityStatement" && !writtenSP {
			for _, rest := range res.Rest {
				var sps strings.Builder
				sps.WriteString("package " + fhirVersion + "Client \n\n")
				sps.WriteString(header)
				sps.WriteString("import \"net/url\"\n")
				sps.WriteString(`type SearchParam interface{
									SpEncode(baseURL* string)(*string, error)
								}`)
				for _, restRes := range rest.Resource {
					sps.WriteString("//search params for " + restRes.Type + "\n")
					sps.WriteString("type Sp" + restRes.Type + " struct{\n")
					for _, sp := range restRes.SearchParam {
						paramName := strings.ReplaceAll(strings.Title(sp.Name), "-", "")
						sps.WriteString(paramName + " string\n")
					}
					sps.WriteString("}\n")

					sps.WriteString("func (p Sp" + restRes.Type + ")SpEncode (baseURL *string) (*string, error) {")
					sps.WriteString(`v := url.Values{}

						params := []struct {
							key   string
							value string
						}{`)
					sps.WriteString("\n")
					for _, sp := range restRes.SearchParam {
						paramName := strings.ReplaceAll(strings.Title(sp.Name), "-", "")
						nameLower := strings.ReplaceAll(sp.Name, "-", "")
						sps.WriteString("{\"" + nameLower + "\", p." + paramName + "},\n")
					}
					sps.WriteString(`
					}
						for _, kv := range params {
							if kv.value != "" {
								v.Set(kv.key, kv.value)
							}
						}

						resUrl, err := url.JoinPath(*baseURL, "` + restRes.Type + `")
						if err != nil{
							return nil, err
						}
						u, err := url.Parse(resUrl)
						if err != nil {
							return nil, err
						}
						u.RawQuery = v.Encode()
						s := u.String()
						return &s, nil
					}

					`)
				}
				searchParamFile, _ := os.Create(path.Join(generateClientDirVer, "searchParam.go"))
				fmt.Fprint(searchParamFile, sps.String())
				searchParamFile.Close()
				writtenSP = true //there's another CapabilityStatement, which is empty, and we dont want to overwrite searchParamFile
			}
		}
		if res.BaseDefinition == "http://hl7.org/fhir/StructureDefinition/DomainResource" {
			if res.Name != "MetadataResource" {
				domainResources = append(domainResources, res.Name)
			}
		}
	}

	cf, _ := os.Create(path.Join(generateClientDirVer, "client.go"))
	fmt.Fprintf(cf, "package %sClient", fhirVersion)
	fmt.Fprintf(cf, `
	import (
		"net/http"
		"strings"
		"time"
	)

	// http client for fhir server, supports Create, Read, Update, Patch, Delete, Search.
	// create with base url eg client.New("https://r4.smarthealthit.org")
	type FhirClient struct {
		BaseUrl string
		*http.Client
	}

	// http client for fhir server, supports Create, Read, Update, Patch, Delete, Search.
	// create with base url eg client.New("https://r4.smarthealthit.org")
	func New(FhirServer_BaseUrl string) *FhirClient {
		if !strings.HasPrefix(FhirServer_BaseUrl, "https://") {
			FhirServer_BaseUrl = "https://" + FhirServer_BaseUrl
		}
		return &FhirClient{
			BaseUrl: FhirServer_BaseUrl,
			Client: &http.Client{
				Timeout: 15 * time.Second,
			},
		}
	}`)

	//domainResources has list of resources we want client REST operations for
	//Binary is not a domainResource, but still gets REST operations
	domainResources = append(domainResources, "Binary")
	slices.Sort(domainResources)

	//custom bundle unmarshal for handling resourc
	// ...kind of belongs in structs (writes to struct) but related to client
	writeUnmarshalBundle(domainResources, generateDirVer)
	//bundle to group of resources, for easier use
	WriteResourceGroup(domainResources, generateClientDirVer, fhirVersion)
	writeRest(domainResources, generateClientDirVer, fhirVersion)
}

//endregion

func generateIndexes(parts []string) string {
	vars := []string{}
	pathStr := ""
	for _, p := range parts {
		pathStr += strings.Title(p)
		vars = append(vars, "num"+strings.Title(p))
	}
	return strings.Join(vars, ", ")
}
