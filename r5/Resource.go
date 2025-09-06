package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Resource
type Resource struct {
	Id            *string `json:"id,omitempty"`
	Meta          *Meta   `json:"meta,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Language      *string `json:"language,omitempty"`
}

func (resource *Resource) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Resource.Id", nil)
	}
	return StringInput("Resource.Id", resource.Id)
}
func (resource *Resource) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Resource.ImplicitRules", nil)
	}
	return StringInput("Resource.ImplicitRules", resource.ImplicitRules)
}
func (resource *Resource) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Resource.Language", nil, optionsValueSet)
	}
	return CodeSelect("Resource.Language", resource.Language, optionsValueSet)
}
