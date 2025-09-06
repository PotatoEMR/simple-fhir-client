package r5

//generated with command go run ./bultaoreune -nodownload
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
	OccurrenceDateTime  *string                       `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn          *string                       `json:"authoredOn,omitempty"`
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

func (resource *CommunicationRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CommunicationRequest.Id", nil)
	}
	return StringInput("CommunicationRequest.Id", resource.Id)
}
func (resource *CommunicationRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CommunicationRequest.ImplicitRules", nil)
	}
	return StringInput("CommunicationRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *CommunicationRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CommunicationRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("CommunicationRequest.Language", resource.Language, optionsValueSet)
}
func (resource *CommunicationRequest) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("CommunicationRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("CommunicationRequest.Status", &resource.Status, optionsValueSet)
}
func (resource *CommunicationRequest) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("CommunicationRequest.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CommunicationRequest.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *CommunicationRequest) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("CommunicationRequest.Intent", nil, optionsValueSet)
	}
	return CodeSelect("CommunicationRequest.Intent", &resource.Intent, optionsValueSet)
}
func (resource *CommunicationRequest) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("CommunicationRequest.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CommunicationRequest.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *CommunicationRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("CommunicationRequest.Priority", nil, optionsValueSet)
	}
	return CodeSelect("CommunicationRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *CommunicationRequest) T_DoNotPerform() templ.Component {

	if resource == nil {
		return BoolInput("CommunicationRequest.DoNotPerform", nil)
	}
	return BoolInput("CommunicationRequest.DoNotPerform", resource.DoNotPerform)
}
func (resource *CommunicationRequest) T_Medium(numMedium int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Medium) >= numMedium {
		return CodeableConceptSelect("CommunicationRequest.Medium["+strconv.Itoa(numMedium)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CommunicationRequest.Medium["+strconv.Itoa(numMedium)+"]", &resource.Medium[numMedium], optionsValueSet)
}
func (resource *CommunicationRequest) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("CommunicationRequest.AuthoredOn", nil)
	}
	return StringInput("CommunicationRequest.AuthoredOn", resource.AuthoredOn)
}
func (resource *CommunicationRequest) T_PayloadId(numPayload int) templ.Component {

	if resource == nil || len(resource.Payload) >= numPayload {
		return StringInput("CommunicationRequest.Payload["+strconv.Itoa(numPayload)+"].Id", nil)
	}
	return StringInput("CommunicationRequest.Payload["+strconv.Itoa(numPayload)+"].Id", resource.Payload[numPayload].Id)
}
