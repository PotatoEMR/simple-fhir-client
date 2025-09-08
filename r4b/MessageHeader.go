package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/MessageHeader
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
	EventUri          string                     `json:"eventUri"`
	Destination       []MessageHeaderDestination `json:"destination,omitempty"`
	Sender            *Reference                 `json:"sender,omitempty"`
	Enterer           *Reference                 `json:"enterer,omitempty"`
	Author            *Reference                 `json:"author,omitempty"`
	Source            MessageHeaderSource        `json:"source"`
	Responsible       *Reference                 `json:"responsible,omitempty"`
	Reason            *CodeableConcept           `json:"reason,omitempty"`
	Response          *MessageHeaderResponse     `json:"response,omitempty"`
	Focus             []Reference                `json:"focus,omitempty"`
	Definition        *string                    `json:"definition,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MessageHeader
type MessageHeaderDestination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Target            *Reference  `json:"target,omitempty"`
	Endpoint          string      `json:"endpoint"`
	Receiver          *Reference  `json:"receiver,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MessageHeader
type MessageHeaderSource struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Name              *string       `json:"name,omitempty"`
	Software          *string       `json:"software,omitempty"`
	Version           *string       `json:"version,omitempty"`
	Contact           *ContactPoint `json:"contact,omitempty"`
	Endpoint          string        `json:"endpoint"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MessageHeader
type MessageHeaderResponse struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        string      `json:"identifier"`
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
		return CodingSelect("MessageHeader.EventCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("MessageHeader.EventCoding", &resource.EventCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_EventUri(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.EventUri", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.EventUri", &resource.EventUri, htmlAttrs)
}
func (resource *MessageHeader) T_Reason(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MessageHeader.Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MessageHeader.Reason", resource.Reason, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_Definition(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Definition", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Definition", resource.Definition, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationName(numDestination int, htmlAttrs string) templ.Component {

	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("MessageHeader.Destination."+strconv.Itoa(numDestination)+"..Name", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Destination."+strconv.Itoa(numDestination)+"..Name", resource.Destination[numDestination].Name, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationEndpoint(numDestination int, htmlAttrs string) templ.Component {

	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("MessageHeader.Destination."+strconv.Itoa(numDestination)+"..Endpoint", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Destination."+strconv.Itoa(numDestination)+"..Endpoint", &resource.Destination[numDestination].Endpoint, htmlAttrs)
}
func (resource *MessageHeader) T_SourceName(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Name", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Source.Name", resource.Source.Name, htmlAttrs)
}
func (resource *MessageHeader) T_SourceSoftware(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Software", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Source.Software", resource.Source.Software, htmlAttrs)
}
func (resource *MessageHeader) T_SourceVersion(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Version", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Source.Version", resource.Source.Version, htmlAttrs)
}
func (resource *MessageHeader) T_SourceEndpoint(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageHeader.Source.Endpoint", nil, htmlAttrs)
	}
	return StringInput("MessageHeader.Source.Endpoint", &resource.Source.Endpoint, htmlAttrs)
}
func (resource *MessageHeader) T_ResponseCode(htmlAttrs string) templ.Component {
	optionsValueSet := VSResponse_code

	if resource == nil {
		return CodeSelect("MessageHeader.Response.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MessageHeader.Response.Code", &resource.Response.Code, optionsValueSet, htmlAttrs)
}
