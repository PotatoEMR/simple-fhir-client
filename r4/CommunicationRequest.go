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

// http://hl7.org/fhir/r4/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                 *string                       `json:"id,omitempty"`
	Meta               *Meta                         `json:"meta,omitempty"`
	ImplicitRules      *string                       `json:"implicitRules,omitempty"`
	Language           *string                       `json:"language,omitempty"`
	Text               *Narrative                    `json:"text,omitempty"`
	Contained          []Resource                    `json:"contained,omitempty"`
	Extension          []Extension                   `json:"extension,omitempty"`
	ModifierExtension  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `json:"identifier,omitempty"`
	BasedOn            []Reference                   `json:"basedOn,omitempty"`
	Replaces           []Reference                   `json:"replaces,omitempty"`
	GroupIdentifier    *Identifier                   `json:"groupIdentifier,omitempty"`
	Status             string                        `json:"status"`
	StatusReason       *CodeableConcept              `json:"statusReason,omitempty"`
	Category           []CodeableConcept             `json:"category,omitempty"`
	Priority           *string                       `json:"priority,omitempty"`
	DoNotPerform       *bool                         `json:"doNotPerform,omitempty"`
	Medium             []CodeableConcept             `json:"medium,omitempty"`
	Subject            *Reference                    `json:"subject,omitempty"`
	About              []Reference                   `json:"about,omitempty"`
	Encounter          *Reference                    `json:"encounter,omitempty"`
	Payload            []CommunicationRequestPayload `json:"payload,omitempty"`
	OccurrenceDateTime *time.Time                    `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod   *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn         *time.Time                    `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Requester          *Reference                    `json:"requester,omitempty"`
	Recipient          []Reference                   `json:"recipient,omitempty"`
	Sender             *Reference                    `json:"sender,omitempty"`
	ReasonCode         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `json:"reasonReference,omitempty"`
	Note               []Annotation                  `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CommunicationRequest
type CommunicationRequestPayload struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentString     string      `json:"contentString"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
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
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *CommunicationRequest) T_Medium(numMedium int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedium >= len(resource.Medium) {
		return CodeableConceptSelect("Medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *CommunicationRequest) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *CommunicationRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CommunicationRequest) T_PayloadContentString(numPayload int, htmlAttrs string) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return StringInput("Payload["+strconv.Itoa(numPayload)+"]ContentString", nil, htmlAttrs)
	}
	return StringInput("Payload["+strconv.Itoa(numPayload)+"]ContentString", &resource.Payload[numPayload].ContentString, htmlAttrs)
}
