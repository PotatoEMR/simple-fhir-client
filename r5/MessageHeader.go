package r5

//generated with command go run ./bultaoreune
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
func (r MessageHeader) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MessageHeader/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MessageHeader"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MessageHeader) T_EventCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("eventCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("eventCoding", &resource.EventCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_EventCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("eventCanonical", nil, htmlAttrs)
	}
	return StringInput("eventCanonical", &resource.EventCanonical, htmlAttrs)
}
func (resource *MessageHeader) T_Reason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason", resource.Reason, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_Definition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("definition", nil, htmlAttrs)
	}
	return StringInput("definition", resource.Definition, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationEndpointUrl(numDestination int, htmlAttrs string) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("destination["+strconv.Itoa(numDestination)+"].endpointUrl", nil, htmlAttrs)
	}
	return StringInput("destination["+strconv.Itoa(numDestination)+"].endpointUrl", resource.Destination[numDestination].EndpointUrl, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationName(numDestination int, htmlAttrs string) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("destination["+strconv.Itoa(numDestination)+"].name", nil, htmlAttrs)
	}
	return StringInput("destination["+strconv.Itoa(numDestination)+"].name", resource.Destination[numDestination].Name, htmlAttrs)
}
func (resource *MessageHeader) T_SourceEndpointUrl(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("source.endpointUrl", nil, htmlAttrs)
	}
	return StringInput("source.endpointUrl", resource.Source.EndpointUrl, htmlAttrs)
}
func (resource *MessageHeader) T_SourceName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("source.name", nil, htmlAttrs)
	}
	return StringInput("source.name", resource.Source.Name, htmlAttrs)
}
func (resource *MessageHeader) T_SourceSoftware(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("source.software", nil, htmlAttrs)
	}
	return StringInput("source.software", resource.Source.Software, htmlAttrs)
}
func (resource *MessageHeader) T_SourceVersion(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("source.version", nil, htmlAttrs)
	}
	return StringInput("source.version", resource.Source.Version, htmlAttrs)
}
func (resource *MessageHeader) T_ResponseCode(htmlAttrs string) templ.Component {
	optionsValueSet := VSResponse_code

	if resource == nil {
		return CodeSelect("response.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("response.code", &resource.Response.Code, optionsValueSet, htmlAttrs)
}
