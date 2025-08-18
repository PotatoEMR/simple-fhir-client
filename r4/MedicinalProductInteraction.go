//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductInteraction
type MedicinalProductInteraction struct {
	Id                *string                                  `json:"id,omitempty"`
	Meta              *Meta                                    `json:"meta,omitempty"`
	ImplicitRules     *string                                  `json:"implicitRules,omitempty"`
	Language          *string                                  `json:"language,omitempty"`
	Text              *Narrative                               `json:"text,omitempty"`
	Contained         []Resource                               `json:"contained,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Subject           []Reference                              `json:"subject,omitempty"`
	Description       *string                                  `json:"description,omitempty"`
	Interactant       []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`
	Type              *CodeableConcept                         `json:"type,omitempty"`
	Effect            *CodeableConcept                         `json:"effect,omitempty"`
	Incidence         *CodeableConcept                         `json:"incidence,omitempty"`
	Management        *CodeableConcept                         `json:"management,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductInteraction
type MedicinalProductInteractionInteractant struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

type OtherMedicinalProductInteraction MedicinalProductInteraction

// on convert struct to json, automatically add resourceType=MedicinalProductInteraction
func (r MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductInteraction
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductInteraction: OtherMedicinalProductInteraction(r),
		ResourceType:                     "MedicinalProductInteraction",
	})
}
