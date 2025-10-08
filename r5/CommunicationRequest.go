package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	OccurrenceDateTime  *FhirDateTime                 `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn          *FhirDateTime                 `json:"authoredOn,omitempty"`
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
func (resource *CommunicationRequest) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *CommunicationRequest) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *CommunicationRequest) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
}
func (resource *CommunicationRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *CommunicationRequest) T_Medium(numMedium int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedium >= len(resource.Medium) {
		return CodeableConceptSelect("medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet, htmlAttrs)
}
func (resource *CommunicationRequest) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *CommunicationRequest) T_About(frs []FhirResource, numAbout int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAbout >= len(resource.About) {
		return ReferenceInput(frs, "about["+strconv.Itoa(numAbout)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "about["+strconv.Itoa(numAbout)+"]", &resource.About[numAbout], htmlAttrs)
}
func (resource *CommunicationRequest) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *CommunicationRequest) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *CommunicationRequest) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *CommunicationRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *CommunicationRequest) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *CommunicationRequest) T_Recipient(frs []FhirResource, numRecipient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecipient >= len(resource.Recipient) {
		return ReferenceInput(frs, "recipient["+strconv.Itoa(numRecipient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recipient["+strconv.Itoa(numRecipient)+"]", &resource.Recipient[numRecipient], htmlAttrs)
}
func (resource *CommunicationRequest) T_InformationProvider(frs []FhirResource, numInformationProvider int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInformationProvider >= len(resource.InformationProvider) {
		return ReferenceInput(frs, "informationProvider["+strconv.Itoa(numInformationProvider)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "informationProvider["+strconv.Itoa(numInformationProvider)+"]", &resource.InformationProvider[numInformationProvider], htmlAttrs)
}
func (resource *CommunicationRequest) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *CommunicationRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CommunicationRequest) T_PayloadContentAttachment(numPayload int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return AttachmentInput("payload["+strconv.Itoa(numPayload)+"].contentAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("payload["+strconv.Itoa(numPayload)+"].contentAttachment", &resource.Payload[numPayload].ContentAttachment, htmlAttrs)
}
func (resource *CommunicationRequest) T_PayloadContentReference(frs []FhirResource, numPayload int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return ReferenceInput(frs, "payload["+strconv.Itoa(numPayload)+"].contentReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "payload["+strconv.Itoa(numPayload)+"].contentReference", &resource.Payload[numPayload].ContentReference, htmlAttrs)
}
func (resource *CommunicationRequest) T_PayloadContentCodeableConcept(numPayload int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) {
		return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].contentCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].contentCodeableConcept", &resource.Payload[numPayload].ContentCodeableConcept, optionsValueSet, htmlAttrs)
}
