package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Binary) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Binary.Id", nil)
	}
	return StringInput("Binary.Id", resource.Id)
}
func (resource *Binary) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Binary.ImplicitRules", nil)
	}
	return StringInput("Binary.ImplicitRules", resource.ImplicitRules)
}
func (resource *Binary) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Binary.Language", nil, optionsValueSet)
	}
	return CodeSelect("Binary.Language", resource.Language, optionsValueSet)
}
func (resource *Binary) T_ContentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Binary.ContentType", nil, optionsValueSet)
	}
	return CodeSelect("Binary.ContentType", &resource.ContentType, optionsValueSet)
}
func (resource *Binary) T_Data() templ.Component {

	if resource == nil {
		return StringInput("Binary.Data", nil)
	}
	return StringInput("Binary.Data", resource.Data)
}
