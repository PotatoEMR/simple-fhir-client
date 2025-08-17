//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
