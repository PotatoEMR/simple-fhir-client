package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Sent                  *string                `json:"sent,omitempty"`
	Received              *string                `json:"received,omitempty"`
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

func (resource *Communication) CommunicationLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Communication) CommunicationStatus() templ.Component {
	optionsValueSet := VSEvent_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Communication) CommunicationPriority() templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
