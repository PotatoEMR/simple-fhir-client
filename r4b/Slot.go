package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Slot
type Slot struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	ServiceCategory   []CodeableConcept `json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `json:"specialty,omitempty"`
	AppointmentType   *CodeableConcept  `json:"appointmentType,omitempty"`
	Schedule          Reference         `json:"schedule"`
	Status            string            `json:"status"`
	Start             string            `json:"start"`
	End               string            `json:"end"`
	Overbooked        *bool             `json:"overbooked,omitempty"`
	Comment           *string           `json:"comment,omitempty"`
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

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Slot) SlotServiceCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("serviceCategory", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceCategory", &resource.ServiceCategory[0], optionsValueSet)
}
func (resource *Slot) SlotServiceType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceType", &resource.ServiceType[0], optionsValueSet)
}
func (resource *Slot) SlotSpecialty(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
func (resource *Slot) SlotAppointmentType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("appointmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appointmentType", resource.AppointmentType, optionsValueSet)
}
func (resource *Slot) SlotStatus() templ.Component {
	optionsValueSet := VSSlotstatus

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
