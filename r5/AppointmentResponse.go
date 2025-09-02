package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/AppointmentResponse
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
	ProposedNewTime   *bool             `json:"proposedNewTime,omitempty"`
	Start             *string           `json:"start,omitempty"`
	End               *string           `json:"end,omitempty"`
	ParticipantType   []CodeableConcept `json:"participantType,omitempty"`
	Actor             *Reference        `json:"actor,omitempty"`
	ParticipantStatus string            `json:"participantStatus"`
	Comment           *string           `json:"comment,omitempty"`
	Recurring         *bool             `json:"recurring,omitempty"`
	OccurrenceDate    *string           `json:"occurrenceDate,omitempty"`
	RecurrenceId      *int              `json:"recurrenceId,omitempty"`
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

func (resource *AppointmentResponse) AppointmentResponseLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AppointmentResponse) AppointmentResponseParticipantType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("participantType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("participantType", &resource.ParticipantType[0], optionsValueSet)
}
func (resource *AppointmentResponse) AppointmentResponseParticipantStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("participantStatus", nil, optionsValueSet)
	}
	return CodeSelect("participantStatus", &resource.ParticipantStatus, optionsValueSet)
}
