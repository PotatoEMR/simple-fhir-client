package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Communication
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

// http://hl7.org/fhir/r4b/StructureDefinition/Communication
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

func (resource *Communication) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Communication.Id", nil)
	}
	return StringInput("Communication.Id", resource.Id)
}
func (resource *Communication) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Communication.ImplicitRules", nil)
	}
	return StringInput("Communication.ImplicitRules", resource.ImplicitRules)
}
func (resource *Communication) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Communication.Language", nil, optionsValueSet)
	}
	return CodeSelect("Communication.Language", resource.Language, optionsValueSet)
}
func (resource *Communication) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("Communication.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("Communication.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *Communication) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("Communication.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("Communication.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *Communication) T_Status() templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Communication.Status", nil, optionsValueSet)
	}
	return CodeSelect("Communication.Status", &resource.Status, optionsValueSet)
}
func (resource *Communication) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Communication.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Communication.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Communication) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Communication.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Communication.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Communication) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Communication.Priority", nil, optionsValueSet)
	}
	return CodeSelect("Communication.Priority", resource.Priority, optionsValueSet)
}
func (resource *Communication) T_Medium(numMedium int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Medium) >= numMedium {
		return CodeableConceptSelect("Communication.Medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Communication.Medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet)
}
func (resource *Communication) T_Topic(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Communication.Topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Communication.Topic", resource.Topic, optionsValueSet)
}
func (resource *Communication) T_Sent() templ.Component {

	if resource == nil {
		return StringInput("Communication.Sent", nil)
	}
	return StringInput("Communication.Sent", resource.Sent)
}
func (resource *Communication) T_Received() templ.Component {

	if resource == nil {
		return StringInput("Communication.Received", nil)
	}
	return StringInput("Communication.Received", resource.Received)
}
func (resource *Communication) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Communication.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Communication.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Communication) T_PayloadId(numPayload int) templ.Component {

	if resource == nil || len(resource.Payload) >= numPayload {
		return StringInput("Communication.Payload["+strconv.Itoa(numPayload)+"].Id", nil)
	}
	return StringInput("Communication.Payload["+strconv.Itoa(numPayload)+"].Id", resource.Payload[numPayload].Id)
}
