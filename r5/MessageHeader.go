package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MessageHeader
type MessageHeader struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	EventCoding       Coding                     `json:"eventCoding"`
	EventCanonical    string                     `json:"eventCanonical"`
	Destination       []MessageHeaderDestination `json:"destination,omitempty"`
	Sender            *Reference                 `json:"sender,omitempty"`
	Author            *Reference                 `json:"author,omitempty"`
	Source            MessageHeaderSource        `json:"source"`
	Responsible       *Reference                 `json:"responsible,omitempty"`
	Reason            *CodeableConcept           `json:"reason,omitempty"`
	Response          *MessageHeaderResponse     `json:"response,omitempty"`
	Focus             []Reference                `json:"focus,omitempty"`
	Definition        *string                    `json:"definition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MessageHeader
type MessageHeaderDestination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	EndpointUrl       *string     `json:"endpointUrl,omitempty"`
	EndpointReference *Reference  `json:"endpointReference,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Target            *Reference  `json:"target,omitempty"`
	Receiver          *Reference  `json:"receiver,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MessageHeader
type MessageHeaderSource struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	EndpointUrl       *string       `json:"endpointUrl,omitempty"`
	EndpointReference *Reference    `json:"endpointReference,omitempty"`
	Name              *string       `json:"name,omitempty"`
	Software          *string       `json:"software,omitempty"`
	Version           *string       `json:"version,omitempty"`
	Contact           *ContactPoint `json:"contact,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MessageHeader
type MessageHeaderResponse struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        Identifier  `json:"identifier"`
	Code              string      `json:"code"`
	Details           *Reference  `json:"details,omitempty"`
}

type OtherMessageHeader MessageHeader

// on convert struct to json, automatically add resourceType=MessageHeader
func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageHeader
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageHeader: OtherMessageHeader(r),
		ResourceType:       "MessageHeader",
	})
}

func (resource *MessageHeader) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Id", nil)
	}
	return StringInput("MessageHeader.Id", resource.Id)
}
func (resource *MessageHeader) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.ImplicitRules", nil)
	}
	return StringInput("MessageHeader.ImplicitRules", resource.ImplicitRules)
}
func (resource *MessageHeader) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MessageHeader.Language", nil, optionsValueSet)
	}
	return CodeSelect("MessageHeader.Language", resource.Language, optionsValueSet)
}
func (resource *MessageHeader) T_Reason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MessageHeader.Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MessageHeader.Reason", resource.Reason, optionsValueSet)
}
func (resource *MessageHeader) T_Definition() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Definition", nil)
	}
	return StringInput("MessageHeader.Definition", resource.Definition)
}
func (resource *MessageHeader) T_DestinationId(numDestination int) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return StringInput("MessageHeader.Destination["+strconv.Itoa(numDestination)+"].Id", nil)
	}
	return StringInput("MessageHeader.Destination["+strconv.Itoa(numDestination)+"].Id", resource.Destination[numDestination].Id)
}
func (resource *MessageHeader) T_DestinationName(numDestination int) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return StringInput("MessageHeader.Destination["+strconv.Itoa(numDestination)+"].Name", nil)
	}
	return StringInput("MessageHeader.Destination["+strconv.Itoa(numDestination)+"].Name", resource.Destination[numDestination].Name)
}
func (resource *MessageHeader) T_SourceId() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Id", nil)
	}
	return StringInput("MessageHeader.Source.Id", resource.Source.Id)
}
func (resource *MessageHeader) T_SourceName() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Name", nil)
	}
	return StringInput("MessageHeader.Source.Name", resource.Source.Name)
}
func (resource *MessageHeader) T_SourceSoftware() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Software", nil)
	}
	return StringInput("MessageHeader.Source.Software", resource.Source.Software)
}
func (resource *MessageHeader) T_SourceVersion() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Version", nil)
	}
	return StringInput("MessageHeader.Source.Version", resource.Source.Version)
}
func (resource *MessageHeader) T_ResponseId() templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Response.Id", nil)
	}
	return StringInput("MessageHeader.Response.Id", resource.Response.Id)
}
func (resource *MessageHeader) T_ResponseCode() templ.Component {
	optionsValueSet := VSResponse_code

	if resource == nil {
		return CodeSelect("MessageHeader.Response.Code", nil, optionsValueSet)
	}
	return CodeSelect("MessageHeader.Response.Code", &resource.Response.Code, optionsValueSet)
}
