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

// http://hl7.org/fhir/r4/StructureDefinition/Communication
type Communication struct {
	Id                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []Resource             `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	PartOf                []Reference            `json:"partOf,omitempty"`
	InResponseTo          []Reference            `json:"inResponseTo,omitempty"`
	Status                string                 `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	Category              []CodeableConcept      `json:"category,omitempty"`
	Priority              *string                `json:"priority,omitempty"`
	Medium                []CodeableConcept      `json:"medium,omitempty"`
	Subject               *Reference             `json:"subject,omitempty"`
	Topic                 *CodeableConcept       `json:"topic,omitempty"`
	About                 []Reference            `json:"about,omitempty"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	Sent                  *time.Time             `json:"sent,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Received              *time.Time             `json:"received,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Recipient             []Reference            `json:"recipient,omitempty"`
	Sender                *Reference             `json:"sender,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	Payload               []CommunicationPayload `json:"payload,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Communication
type CommunicationPayload struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentString     string      `json:"contentString"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

type OtherCommunication Communication

// on convert struct to json, automatically add resourceType=Communication
func (r Communication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
}
func (r Communication) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Communication/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Communication"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Communication) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("Communication.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("Communication.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *Communication) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("Communication.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", nil, htmlAttrs)
	}
	return StringInput("Communication.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *Communication) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Communication.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Communication.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Communication.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Communication.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Communication.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Communication.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Medium(numMedium int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMedium >= len(resource.Medium) {
		return CodeableConceptSelect("Communication.Medium."+strconv.Itoa(numMedium)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication.Medium."+strconv.Itoa(numMedium)+".", &resource.Medium[numMedium], optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Topic(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Communication.Topic", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication.Topic", resource.Topic, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Sent(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Communication.Sent", nil, htmlAttrs)
	}
	return DateTimeInput("Communication.Sent", resource.Sent, htmlAttrs)
}
func (resource *Communication) T_Received(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Communication.Received", nil, htmlAttrs)
	}
	return DateTimeInput("Communication.Received", resource.Received, htmlAttrs)
}
func (resource *Communication) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Communication.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Communication.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Communication.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *Communication) T_PayloadContentString(numPayload int, htmlAttrs string) templ.Component {

	if resource == nil || numPayload >= len(resource.Payload) {
		return StringInput("Communication.Payload."+strconv.Itoa(numPayload)+"..ContentString", nil, htmlAttrs)
	}
	return StringInput("Communication.Payload."+strconv.Itoa(numPayload)+"..ContentString", &resource.Payload[numPayload].ContentString, htmlAttrs)
}
