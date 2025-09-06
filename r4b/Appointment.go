package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Appointment
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
	Created               *string                  `json:"created,omitempty"`
	Comment               *string                  `json:"comment,omitempty"`
	PatientInstruction    *string                  `json:"patientInstruction,omitempty"`
	BasedOn               []Reference              `json:"basedOn,omitempty"`
	Participant           []AppointmentParticipant `json:"participant"`
	RequestedPeriod       []Period                 `json:"requestedPeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Appointment
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

func (resource *Appointment) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Id", nil)
	}
	return StringInput("Appointment.Id", resource.Id)
}
func (resource *Appointment) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Appointment.ImplicitRules", nil)
	}
	return StringInput("Appointment.ImplicitRules", resource.ImplicitRules)
}
func (resource *Appointment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Appointment.Language", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Language", resource.Language, optionsValueSet)
}
func (resource *Appointment) T_Status() templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("Appointment.Status", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Status", &resource.Status, optionsValueSet)
}
func (resource *Appointment) T_CancelationReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.CancelationReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.CancelationReason", resource.CancelationReason, optionsValueSet)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceCategory) >= numServiceCategory {
		return CodeableConceptSelect("Appointment.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet)
}
func (resource *Appointment) T_ServiceType(numServiceType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceType) >= numServiceType {
		return CodeableConceptSelect("Appointment.ServiceType["+strconv.Itoa(numServiceType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.ServiceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], optionsValueSet)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("Appointment.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.AppointmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.AppointmentType", resource.AppointmentType, optionsValueSet)
}
func (resource *Appointment) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Appointment.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Appointment) T_Priority() templ.Component {

	if resource == nil {
		return IntInput("Appointment.Priority", nil)
	}
	return IntInput("Appointment.Priority", resource.Priority)
}
func (resource *Appointment) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Description", nil)
	}
	return StringInput("Appointment.Description", resource.Description)
}
func (resource *Appointment) T_Start() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Start", nil)
	}
	return StringInput("Appointment.Start", resource.Start)
}
func (resource *Appointment) T_End() templ.Component {

	if resource == nil {
		return StringInput("Appointment.End", nil)
	}
	return StringInput("Appointment.End", resource.End)
}
func (resource *Appointment) T_MinutesDuration() templ.Component {

	if resource == nil {
		return IntInput("Appointment.MinutesDuration", nil)
	}
	return IntInput("Appointment.MinutesDuration", resource.MinutesDuration)
}
func (resource *Appointment) T_Created() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Created", nil)
	}
	return StringInput("Appointment.Created", resource.Created)
}
func (resource *Appointment) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Comment", nil)
	}
	return StringInput("Appointment.Comment", resource.Comment)
}
func (resource *Appointment) T_PatientInstruction() templ.Component {

	if resource == nil {
		return StringInput("Appointment.PatientInstruction", nil)
	}
	return StringInput("Appointment.PatientInstruction", resource.PatientInstruction)
}
func (resource *Appointment) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant || len(resource.Participant[numParticipant].Type) >= numType {
		return CodeableConceptSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int) templ.Component {
	optionsValueSet := VSParticipantrequired

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Required", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Required", resource.Participant[numParticipant].Required, optionsValueSet)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Status", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Status", &resource.Participant[numParticipant].Status, optionsValueSet)
}
