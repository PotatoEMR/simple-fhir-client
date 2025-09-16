package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	OccurrenceDate    *FhirDate         `json:"occurrenceDate,omitempty"`
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
func (resource *AppointmentResponse) T_ProposedNewTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("proposedNewTime", nil, htmlAttrs)
	}
	return BoolInput("proposedNewTime", resource.ProposedNewTime, htmlAttrs)
}
func (resource *AppointmentResponse) T_Start(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("start", nil, htmlAttrs)
	}
	return StringInput("start", resource.Start, htmlAttrs)
}
func (resource *AppointmentResponse) T_End(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("end", nil, htmlAttrs)
	}
	return StringInput("end", resource.End, htmlAttrs)
}
func (resource *AppointmentResponse) T_ParticipantType(numParticipantType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipantType >= len(resource.ParticipantType) {
		return CodeableConceptSelect("participantType["+strconv.Itoa(numParticipantType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participantType["+strconv.Itoa(numParticipantType)+"]", &resource.ParticipantType[numParticipantType], optionsValueSet, htmlAttrs)
}
func (resource *AppointmentResponse) T_ParticipantStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("participantStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("participantStatus", &resource.ParticipantStatus, optionsValueSet, htmlAttrs)
}
func (resource *AppointmentResponse) T_Comment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *AppointmentResponse) T_Recurring(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("recurring", nil, htmlAttrs)
	}
	return BoolInput("recurring", resource.Recurring, htmlAttrs)
}
func (resource *AppointmentResponse) T_OccurrenceDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("occurrenceDate", nil, htmlAttrs)
	}
	return DateInput("occurrenceDate", resource.OccurrenceDate, htmlAttrs)
}
func (resource *AppointmentResponse) T_RecurrenceId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("recurrenceId", nil, htmlAttrs)
	}
	return IntInput("recurrenceId", resource.RecurrenceId, htmlAttrs)
}
