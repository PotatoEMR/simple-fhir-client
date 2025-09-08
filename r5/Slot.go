package r5

//generated with command go run ./bultaoreune
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
func (r Slot) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Slot/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Slot"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Slot) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_AppointmentType(numAppointmentType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAppointmentType >= len(resource.AppointmentType) {
		return CodeableConceptSelect("AppointmentType["+strconv.Itoa(numAppointmentType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AppointmentType["+strconv.Itoa(numAppointmentType)+"]", &resource.AppointmentType[numAppointmentType], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSlotstatus

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Start(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Start", nil, htmlAttrs)
	}
	return StringInput("Start", &resource.Start, htmlAttrs)
}
func (resource *Slot) T_End(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("End", nil, htmlAttrs)
	}
	return StringInput("End", &resource.End, htmlAttrs)
}
func (resource *Slot) T_Overbooked(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Overbooked", nil, htmlAttrs)
	}
	return BoolInput("Overbooked", resource.Overbooked, htmlAttrs)
}
func (resource *Slot) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Comment", nil, htmlAttrs)
	}
	return StringInput("Comment", resource.Comment, htmlAttrs)
}
