package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Appointment       Reference         `json:"appointment"`
	Start             *string           `json:"start,omitempty"`
	End               *string           `json:"end,omitempty"`
	ParticipantType   []CodeableConcept `json:"participantType,omitempty"`
	Actor             *Reference        `json:"actor,omitempty"`
	ParticipantStatus string            `json:"participantStatus"`
	Comment           *string           `json:"comment,omitempty"`
}

type OtherAppointmentResponse AppointmentResponse

// on convert struct to json, automatically add resourceType=AppointmentResponse
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointmentResponse: OtherAppointmentResponse(r),
		ResourceType:             "AppointmentResponse",
	})
}

func (resource *AppointmentResponse) T_Id() templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.Id", nil)
	}
	return StringInput("AppointmentResponse.Id", resource.Id)
}
func (resource *AppointmentResponse) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.ImplicitRules", nil)
	}
	return StringInput("AppointmentResponse.ImplicitRules", resource.ImplicitRules)
}
func (resource *AppointmentResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("AppointmentResponse.Language", nil, optionsValueSet)
	}
	return CodeSelect("AppointmentResponse.Language", resource.Language, optionsValueSet)
}
func (resource *AppointmentResponse) T_Start() templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.Start", nil)
	}
	return StringInput("AppointmentResponse.Start", resource.Start)
}
func (resource *AppointmentResponse) T_End() templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.End", nil)
	}
	return StringInput("AppointmentResponse.End", resource.End)
}
func (resource *AppointmentResponse) T_ParticipantType(numParticipantType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ParticipantType) >= numParticipantType {
		return CodeableConceptSelect("AppointmentResponse.ParticipantType["+strconv.Itoa(numParticipantType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AppointmentResponse.ParticipantType["+strconv.Itoa(numParticipantType)+"]", &resource.ParticipantType[numParticipantType], optionsValueSet)
}
func (resource *AppointmentResponse) T_ParticipantStatus() templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil {
		return CodeSelect("AppointmentResponse.ParticipantStatus", nil, optionsValueSet)
	}
	return CodeSelect("AppointmentResponse.ParticipantStatus", &resource.ParticipantStatus, optionsValueSet)
}
func (resource *AppointmentResponse) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.Comment", nil)
	}
	return StringInput("AppointmentResponse.Comment", resource.Comment)
}
