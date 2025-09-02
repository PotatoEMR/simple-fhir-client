package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Sent                  *string                `json:"sent,omitempty"`
	Received              *string                `json:"received,omitempty"`
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

func (resource *Communication) CommunicationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Communication) CommunicationStatus() templ.Component {
	optionsValueSet := VSEvent_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Communication) CommunicationStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Communication) CommunicationCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Communication) CommunicationPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource != nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Communication) CommunicationMedium(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("medium", nil, optionsValueSet)
	}
	return CodeableConceptSelect("medium", &resource.Medium[0], optionsValueSet)
}
func (resource *Communication) CommunicationTopic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", resource.Topic, optionsValueSet)
}
func (resource *Communication) CommunicationReasonCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
