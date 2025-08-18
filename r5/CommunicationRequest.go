//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
