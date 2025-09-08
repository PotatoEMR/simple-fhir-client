package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Appointment
type Appointment struct {
	Id                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []Resource               `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	Status                string                   `json:"status"`
	CancelationReason     *CodeableConcept         `json:"cancelationReason,omitempty"`
	ServiceCategory       []CodeableConcept        `json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept        `json:"serviceType,omitempty"`
	Specialty             []CodeableConcept        `json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept         `json:"appointmentType,omitempty"`
	ReasonCode            []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference       []Reference              `json:"reasonReference,omitempty"`
	Priority              *int                     `json:"priority,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	SupportingInformation []Reference              `json:"supportingInformation,omitempty"`
	Start                 *string                  `json:"start,omitempty"`
	End                   *string                  `json:"end,omitempty"`
	MinutesDuration       *int                     `json:"minutesDuration,omitempty"`
	Slot                  []Reference              `json:"slot,omitempty"`
	Created               *time.Time               `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Comment               *string                  `json:"comment,omitempty"`
	PatientInstruction    *string                  `json:"patientInstruction,omitempty"`
	BasedOn               []Reference              `json:"basedOn,omitempty"`
	Participant           []AppointmentParticipant `json:"participant"`
	RequestedPeriod       []Period                 `json:"requestedPeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Appointment
type AppointmentParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Actor             *Reference        `json:"actor,omitempty"`
	Required          *string           `json:"required,omitempty"`
	Status            string            `json:"status"`
	Period            *Period           `json:"period,omitempty"`
}

type OtherAppointment Appointment

// on convert struct to json, automatically add resourceType=Appointment
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}
func (r Appointment) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Appointment/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Appointment"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Appointment) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("Appointment.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Appointment.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_CancelationReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.CancelationReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.CancelationReason", resource.CancelationReason, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("Appointment.ServiceCategory."+strconv.Itoa(numServiceCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.ServiceCategory."+strconv.Itoa(numServiceCategory)+".", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceType(numServiceType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableConceptSelect("Appointment.ServiceType."+strconv.Itoa(numServiceType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.ServiceType."+strconv.Itoa(numServiceType)+".", &resource.ServiceType[numServiceType], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("Appointment.Specialty."+strconv.Itoa(numSpecialty)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Specialty."+strconv.Itoa(numSpecialty)+".", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.AppointmentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.AppointmentType", resource.AppointmentType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Appointment.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Priority(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("Appointment.Priority", nil, htmlAttrs)
	}
	return IntInput("Appointment.Priority", resource.Priority, htmlAttrs)
}
func (resource *Appointment) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.Description", nil, htmlAttrs)
	}
	return StringInput("Appointment.Description", resource.Description, htmlAttrs)
}
func (resource *Appointment) T_Start(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.Start", nil, htmlAttrs)
	}
	return StringInput("Appointment.Start", resource.Start, htmlAttrs)
}
func (resource *Appointment) T_End(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.End", nil, htmlAttrs)
	}
	return StringInput("Appointment.End", resource.End, htmlAttrs)
}
func (resource *Appointment) T_MinutesDuration(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("Appointment.MinutesDuration", nil, htmlAttrs)
	}
	return IntInput("Appointment.MinutesDuration", resource.MinutesDuration, htmlAttrs)
}
func (resource *Appointment) T_Created(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Appointment.Created", nil, htmlAttrs)
	}
	return DateTimeInput("Appointment.Created", resource.Created, htmlAttrs)
}
func (resource *Appointment) T_Comment(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.Comment", nil, htmlAttrs)
	}
	return StringInput("Appointment.Comment", resource.Comment, htmlAttrs)
}
func (resource *Appointment) T_PatientInstruction(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.PatientInstruction", nil, htmlAttrs)
	}
	return StringInput("Appointment.PatientInstruction", resource.PatientInstruction, htmlAttrs)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSParticipantrequired

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Required", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Required", resource.Participant[numParticipant].Required, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Status", &resource.Participant[numParticipant].Status, optionsValueSet, htmlAttrs)
}
