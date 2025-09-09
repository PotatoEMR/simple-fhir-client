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

// http://hl7.org/fhir/r5/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                  *string                       `json:"id,omitempty"`
	Meta                *Meta                         `json:"meta,omitempty"`
	ImplicitRules       *string                       `json:"implicitRules,omitempty"`
	Language            *string                       `json:"language,omitempty"`
	Text                *Narrative                    `json:"text,omitempty"`
	Contained           []Resource                    `json:"contained,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	Identifier          []Identifier                  `json:"identifier,omitempty"`
	BasedOn             []Reference                   `json:"basedOn,omitempty"`
	Replaces            []Reference                   `json:"replaces,omitempty"`
	GroupIdentifier     *Identifier                   `json:"groupIdentifier,omitempty"`
	Status              string                        `json:"status"`
	StatusReason        *CodeableConcept              `json:"statusReason,omitempty"`
	Intent              string                        `json:"intent"`
	Category            []CodeableConcept             `json:"category,omitempty"`
	Priority            *string                       `json:"priority,omitempty"`
	DoNotPerform        *bool                         `json:"doNotPerform,omitempty"`
	Medium              []CodeableConcept             `json:"medium,omitempty"`
	Subject             *Reference                    `json:"subject,omitempty"`
	About               []Reference                   `json:"about,omitempty"`
	Encounter           *Reference                    `json:"encounter,omitempty"`
	Payload             []CommunicationRequestPayload `json:"payload,omitempty"`
	OccurrenceDateTime  *time.Time                    `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod    *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn          *time.Time                    `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Requester           *Reference                    `json:"requester,omitempty"`
	Recipient           []Reference                   `json:"recipient,omitempty"`
	InformationProvider []Reference                   `json:"informationProvider,omitempty"`
	Reason              []CodeableReference           `json:"reason,omitempty"`
	Note                []Annotation                  `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CommunicationRequest
type CommunicationRequestPayload struct {
	Id                     *string         `json:"id,omitempty"`
	Extension              []Extension     `json:"extension,omitempty"`
	ModifierExtension      []Extension     `json:"modifierExtension,omitempty"`
	ContentAttachment      Attachment      `json:"contentAttachment"`
	ContentReference       Reference       `json:"contentReference"`
	ContentCodeableConcept CodeableConcept `json:"contentCodeableConcept"`
}

type OtherCommunicationRequest CommunicationRequest

// on convert struct to json, automatically add resourceType=CommunicationRequest
func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunicationRequest: OtherCommunicationRequest(r),
		ResourceType:              "CommunicationRequest",
	})
}
func (r CommunicationRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CommunicationRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CommunicationRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CommunicationRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("CommunicationRequest.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CommunicationRequest.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CommunicationRequest.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CommunicationRequest.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("CommunicationRequest.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CommunicationRequest.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("CommunicationRequest.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CommunicationRequest.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("CommunicationRequest.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CommunicationRequest.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("CommunicationRequest.DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("CommunicationRequest.DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *CommunicationRequest) T_Medium(numMedium int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedium >= len(resource.Medium) {
		return CodeableConceptSelect("CommunicationRequest.Medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CommunicationRequest.Medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CommunicationRequest.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("CommunicationRequest.OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *CommunicationRequest) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CommunicationRequest.AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("CommunicationRequest.AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *CommunicationRequest) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("CommunicationRequest.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("CommunicationRequest.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CommunicationRequest) T_PayloadContentCodeableConcept(numPayload int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return CodeableConceptSelect("CommunicationRequest.Payload["+strconv.Itoa(numPayload)+"].ContentCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CommunicationRequest.Payload["+strconv.Itoa(numPayload)+"].ContentCodeableConcept", &resource.Payload[numPayload].ContentCodeableConcept, optionsValueSet, htmlAttrs)
}
