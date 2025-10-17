package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

// struct -> json, automatically add resourceType=Patient
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
func (r Slot) ResourceType() string {
	return "Slot"
}

func (resource *Slot) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_ServiceType(numServiceType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableConceptSelect("serviceType["+strconv.Itoa(numServiceType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_AppointmentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("appointmentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("appointmentType", resource.AppointmentType, optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Schedule(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "schedule", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "schedule", &resource.Schedule, htmlAttrs)
}
func (resource *Slot) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSlotstatus

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Slot) T_Start(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("start", nil, htmlAttrs)
	}
	return StringInput("start", &resource.Start, htmlAttrs)
}
func (resource *Slot) T_End(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("end", nil, htmlAttrs)
	}
	return StringInput("end", &resource.End, htmlAttrs)
}
func (resource *Slot) T_Overbooked(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("overbooked", nil, htmlAttrs)
	}
	return BoolInput("overbooked", resource.Overbooked, htmlAttrs)
}
func (resource *Slot) T_Comment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
