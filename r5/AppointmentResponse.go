package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	OccurrenceDate    *time.Time        `json:"occurrenceDate,omitempty,format:'2006-01-02'"`
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
func (r AppointmentResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AppointmentResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "AppointmentResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AppointmentResponse) T_ProposedNewTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("AppointmentResponse.ProposedNewTime", nil, htmlAttrs)
	}
	return BoolInput("AppointmentResponse.ProposedNewTime", resource.ProposedNewTime, htmlAttrs)
}
func (resource *AppointmentResponse) T_Start(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.Start", nil, htmlAttrs)
	}
	return StringInput("AppointmentResponse.Start", resource.Start, htmlAttrs)
}
func (resource *AppointmentResponse) T_End(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.End", nil, htmlAttrs)
	}
	return StringInput("AppointmentResponse.End", resource.End, htmlAttrs)
}
func (resource *AppointmentResponse) T_ParticipantType(numParticipantType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipantType >= len(resource.ParticipantType) {
		return CodeableConceptSelect("AppointmentResponse.ParticipantType."+strconv.Itoa(numParticipantType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AppointmentResponse.ParticipantType."+strconv.Itoa(numParticipantType)+".", &resource.ParticipantType[numParticipantType], optionsValueSet, htmlAttrs)
}
func (resource *AppointmentResponse) T_ParticipantStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeSelect("AppointmentResponse.ParticipantStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AppointmentResponse.ParticipantStatus", &resource.ParticipantStatus, optionsValueSet, htmlAttrs)
}
func (resource *AppointmentResponse) T_Comment(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("AppointmentResponse.Comment", nil, htmlAttrs)
	}
	return StringInput("AppointmentResponse.Comment", resource.Comment, htmlAttrs)
}
func (resource *AppointmentResponse) T_Recurring(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("AppointmentResponse.Recurring", nil, htmlAttrs)
	}
	return BoolInput("AppointmentResponse.Recurring", resource.Recurring, htmlAttrs)
}
func (resource *AppointmentResponse) T_OccurrenceDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("AppointmentResponse.OccurrenceDate", nil, htmlAttrs)
	}
	return DateInput("AppointmentResponse.OccurrenceDate", resource.OccurrenceDate, htmlAttrs)
}
func (resource *AppointmentResponse) T_RecurrenceId(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("AppointmentResponse.RecurrenceId", nil, htmlAttrs)
	}
	return IntInput("AppointmentResponse.RecurrenceId", resource.RecurrenceId, htmlAttrs)
}
