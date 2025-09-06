package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/FormularyItem
type FormularyItem struct {
	Id                *string          `json:"id,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	ImplicitRules     *string          `json:"implicitRules,omitempty"`
	Language          *string          `json:"language,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	Contained         []Resource       `json:"contained,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `json:"identifier,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *string          `json:"status,omitempty"`
}

type OtherFormularyItem FormularyItem

// on convert struct to json, automatically add resourceType=FormularyItem
func (r FormularyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFormularyItem
		ResourceType string `json:"resourceType"`
	}{
		OtherFormularyItem: OtherFormularyItem(r),
		ResourceType:       "FormularyItem",
	})
}

func (resource *FormularyItem) T_Id() templ.Component {

	if resource == nil {
		return StringInput("FormularyItem.Id", nil)
	}
	return StringInput("FormularyItem.Id", resource.Id)
}
func (resource *FormularyItem) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("FormularyItem.ImplicitRules", nil)
	}
	return StringInput("FormularyItem.ImplicitRules", resource.ImplicitRules)
}
func (resource *FormularyItem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("FormularyItem.Language", nil, optionsValueSet)
	}
	return CodeSelect("FormularyItem.Language", resource.Language, optionsValueSet)
}
func (resource *FormularyItem) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("FormularyItem.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FormularyItem.Code", resource.Code, optionsValueSet)
}
func (resource *FormularyItem) T_Status() templ.Component {
	optionsValueSet := VSFormularyitem_status

	if resource == nil {
		return CodeSelect("FormularyItem.Status", nil, optionsValueSet)
	}
	return CodeSelect("FormularyItem.Status", resource.Status, optionsValueSet)
}
