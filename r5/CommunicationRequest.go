package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	OccurrenceDateTime  *string                       `json:"occurrenceDateTime"`
	OccurrencePeriod    *Period                       `json:"occurrencePeriod"`
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

func (resource *CommunicationRequest) CommunicationRequestLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestStatus() templ.Component {
	optionsValueSet := VSRequest_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestIntent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource != nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource != nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *CommunicationRequest) CommunicationRequestMedium(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("medium", nil, optionsValueSet)
	}
	return CodeableConceptSelect("medium", &resource.Medium[0], optionsValueSet)
}
