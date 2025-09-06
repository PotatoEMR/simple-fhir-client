package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *Slot) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Slot.Id", nil)
	}
	return StringInput("Slot.Id", resource.Id)
}
func (resource *Slot) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Slot.ImplicitRules", nil)
	}
	return StringInput("Slot.ImplicitRules", resource.ImplicitRules)
}
func (resource *Slot) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Slot.Language", nil, optionsValueSet)
	}
	return CodeSelect("Slot.Language", resource.Language, optionsValueSet)
}
func (resource *Slot) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceCategory) >= numServiceCategory {
		return CodeableConceptSelect("Slot.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Slot.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet)
}
func (resource *Slot) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("Slot.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Slot.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *Slot) T_AppointmentType(numAppointmentType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AppointmentType) >= numAppointmentType {
		return CodeableConceptSelect("Slot.AppointmentType["+strconv.Itoa(numAppointmentType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Slot.AppointmentType["+strconv.Itoa(numAppointmentType)+"]", &resource.AppointmentType[numAppointmentType], optionsValueSet)
}
func (resource *Slot) T_Status() templ.Component {
	optionsValueSet := VSSlotstatus

	if resource == nil {
		return CodeSelect("Slot.Status", nil, optionsValueSet)
	}
	return CodeSelect("Slot.Status", &resource.Status, optionsValueSet)
}
func (resource *Slot) T_Start() templ.Component {

	if resource == nil {
		return StringInput("Slot.Start", nil)
	}
	return StringInput("Slot.Start", &resource.Start)
}
func (resource *Slot) T_End() templ.Component {

	if resource == nil {
		return StringInput("Slot.End", nil)
	}
	return StringInput("Slot.End", &resource.End)
}
func (resource *Slot) T_Overbooked() templ.Component {

	if resource == nil {
		return BoolInput("Slot.Overbooked", nil)
	}
	return BoolInput("Slot.Overbooked", resource.Overbooked)
}
func (resource *Slot) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("Slot.Comment", nil)
	}
	return StringInput("Slot.Comment", resource.Comment)
}
