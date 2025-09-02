package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/DomainResource
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

func (resource *DomainResource) DomainResourceLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
