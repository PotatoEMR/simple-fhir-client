package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string     `json:"id,omitempty"`
	Meta              *Meta       `json:"meta,omitempty"`
	ImplicitRules     *string     `json:"implicitRules,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Text              *Narrative  `json:"text,omitempty"`
	Contained         []Resource  `json:"contained,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

func (resource *DomainResource) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DomainResource.Id", nil)
	}
	return StringInput("DomainResource.Id", resource.Id)
}
func (resource *DomainResource) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DomainResource.ImplicitRules", nil)
	}
	return StringInput("DomainResource.ImplicitRules", resource.ImplicitRules)
}
func (resource *DomainResource) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DomainResource.Language", nil, optionsValueSet)
	}
	return CodeSelect("DomainResource.Language", resource.Language, optionsValueSet)
}
