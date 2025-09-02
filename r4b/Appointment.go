package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Appointment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Appointment) T_Status() templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Appointment) T_CancelationReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("cancelationReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("cancelationReason", resource.CancelationReason, optionsValueSet)
}
func (resource *Appointment) T_ServiceCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("serviceCategory", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceCategory", &resource.ServiceCategory[0], optionsValueSet)
}
func (resource *Appointment) T_ServiceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceType", &resource.ServiceType[0], optionsValueSet)
}
func (resource *Appointment) T_Specialty(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("appointmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appointmentType", resource.AppointmentType, optionsValueSet)
}
func (resource *Appointment) T_ReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Participant[numParticipant].Type[0], optionsValueSet)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int) templ.Component {
	optionsValueSet := VSParticipantrequired

	if resource == nil && len(resource.Participant) >= numParticipant {
		return CodeSelect("required", nil, optionsValueSet)
	}
	return CodeSelect("required", resource.Participant[numParticipant].Required, optionsValueSet)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil && len(resource.Participant) >= numParticipant {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Participant[numParticipant].Status, optionsValueSet)
}
