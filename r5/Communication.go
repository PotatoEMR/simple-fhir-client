package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Communication
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
	Sent                  *FhirDateTime          `json:"sent,omitempty"`
	Received              *FhirDateTime          `json:"received,omitempty"`
	Recipient             []Reference            `json:"recipient,omitempty"`
	Sender                *Reference             `json:"sender,omitempty"`
	Reason                []CodeableReference    `json:"reason,omitempty"`
	Payload               []CommunicationPayload `json:"payload,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Communication
type CommunicationPayload struct {
	Id                     *string         `json:"id,omitempty"`
	Extension              []Extension     `json:"extension,omitempty"`
	ModifierExtension      []Extension     `json:"modifierExtension,omitempty"`
	ContentAttachment      Attachment      `json:"contentAttachment"`
	ContentReference       Reference       `json:"contentReference"`
	ContentCodeableConcept CodeableConcept `json:"contentCodeableConcept"`
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
func (resource *Communication) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *Communication) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *Communication) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Communication) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *Communication) T_InResponseTo(frs []FhirResource, numInResponseTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInResponseTo >= len(resource.InResponseTo) {
		return ReferenceInput(frs, "inResponseTo["+strconv.Itoa(numInResponseTo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "inResponseTo["+strconv.Itoa(numInResponseTo)+"]", &resource.InResponseTo[numInResponseTo], htmlAttrs)
}
func (resource *Communication) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Medium(numMedium int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedium >= len(resource.Medium) {
		return CodeableConceptSelect("medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Communication) T_Topic(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic", resource.Topic, optionsValueSet, htmlAttrs)
}
func (resource *Communication) T_About(frs []FhirResource, numAbout int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAbout >= len(resource.About) {
		return ReferenceInput(frs, "about["+strconv.Itoa(numAbout)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "about["+strconv.Itoa(numAbout)+"]", &resource.About[numAbout], htmlAttrs)
}
func (resource *Communication) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Communication) T_Sent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("sent", nil, htmlAttrs)
	}
	return FhirDateTimeInput("sent", resource.Sent, htmlAttrs)
}
func (resource *Communication) T_Received(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("received", nil, htmlAttrs)
	}
	return FhirDateTimeInput("received", resource.Received, htmlAttrs)
}
func (resource *Communication) T_Recipient(frs []FhirResource, numRecipient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecipient >= len(resource.Recipient) {
		return ReferenceInput(frs, "recipient["+strconv.Itoa(numRecipient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recipient["+strconv.Itoa(numRecipient)+"]", &resource.Recipient[numRecipient], htmlAttrs)
}
func (resource *Communication) T_Sender(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "sender", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "sender", resource.Sender, htmlAttrs)
}
func (resource *Communication) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *Communication) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Communication) T_PayloadContentAttachment(numPayload int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return AttachmentInput("payload["+strconv.Itoa(numPayload)+"].contentAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("payload["+strconv.Itoa(numPayload)+"].contentAttachment", &resource.Payload[numPayload].ContentAttachment, htmlAttrs)
}
func (resource *Communication) T_PayloadContentReference(frs []FhirResource, numPayload int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return ReferenceInput(frs, "payload["+strconv.Itoa(numPayload)+"].contentReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "payload["+strconv.Itoa(numPayload)+"].contentReference", &resource.Payload[numPayload].ContentReference, htmlAttrs)
}
func (resource *Communication) T_PayloadContentCodeableConcept(numPayload int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].contentCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].contentCodeableConcept", &resource.Payload[numPayload].ContentCodeableConcept, optionsValueSet, htmlAttrs)
}
