package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Binary
type Binary struct {
	Id              *string    `json:"id,omitempty"`
	Meta            *Meta      `json:"meta,omitempty"`
	ImplicitRules   *string    `json:"implicitRules,omitempty"`
	Language        *string    `json:"language,omitempty"`
	ContentType     string     `json:"contentType"`
	SecurityContext *Reference `json:"securityContext,omitempty"`
	Data            *string    `json:"data,omitempty"`
}

func (resource *Binary) T_ContentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("contentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("contentType", &resource.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *Binary) T_Data(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("data", nil, htmlAttrs)
	}
	return StringInput("data", resource.Data, htmlAttrs)
}
