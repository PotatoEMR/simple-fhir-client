//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/Slot
type Slot struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []Resource          `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `json:"identifier,omitempty"`
	ServiceCategory   []CodeableConcept   `json:"serviceCategory,omitempty"`
	ServiceType       []CodeableReference `json:"serviceType,omitempty"`
	Specialty         []CodeableConcept   `json:"specialty,omitempty"`
	AppointmentType   []CodeableConcept   `json:"appointmentType,omitempty"`
	Schedule          Reference           `json:"schedule"`
	Status            string              `json:"status"`
	Start             string              `json:"start"`
	End               string              `json:"end"`
	Overbooked        *bool               `json:"overbooked,omitempty"`
	Comment           *string             `json:"comment,omitempty"`
}

type OtherSlot Slot

// on convert struct to json, automatically add resourceType=Slot
func (r Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSlot
		ResourceType string `json:"resourceType"`
	}{
		OtherSlot:    OtherSlot(r),
		ResourceType: "Slot",
	})
}
