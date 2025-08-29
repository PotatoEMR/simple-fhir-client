package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Slot) SlotLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Slot) SlotStatus() templ.Component {
	optionsValueSet := VSSlotstatus
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
