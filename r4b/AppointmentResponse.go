package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/AppointmentResponse
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
func (resource *AppointmentResponse) T_Appointment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("appointment", nil, htmlAttrs)
	}
	return ReferenceInput("appointment", &resource.Appointment, htmlAttrs)
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
func (resource *AppointmentResponse) T_Actor(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("actor", nil, htmlAttrs)
	}
	return ReferenceInput("actor", resource.Actor, htmlAttrs)
}
func (resource *AppointmentResponse) T_ParticipantStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSParticipationstatus

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
