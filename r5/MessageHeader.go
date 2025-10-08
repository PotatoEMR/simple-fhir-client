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
func (resource *MessageHeader) T_EventCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("eventCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("eventCoding", &resource.EventCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_EventCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("eventCanonical", nil, htmlAttrs)
	}
	return StringInput("eventCanonical", &resource.EventCanonical, htmlAttrs)
}
func (resource *MessageHeader) T_Sender(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "sender", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "sender", resource.Sender, htmlAttrs)
}
func (resource *MessageHeader) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
func (resource *MessageHeader) T_Responsible(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "responsible", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "responsible", resource.Responsible, htmlAttrs)
}
func (resource *MessageHeader) T_Reason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason", resource.Reason, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_Focus(frs []FhirResource, numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], htmlAttrs)
}
func (resource *MessageHeader) T_Definition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("definition", nil, htmlAttrs)
	}
	return StringInput("definition", resource.Definition, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationEndpointUrl(numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("destination["+strconv.Itoa(numDestination)+"].endpointUrl", nil, htmlAttrs)
	}
	return StringInput("destination["+strconv.Itoa(numDestination)+"].endpointUrl", resource.Destination[numDestination].EndpointUrl, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationEndpointReference(frs []FhirResource, numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].endpointReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].endpointReference", resource.Destination[numDestination].EndpointReference, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationName(numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return StringInput("destination["+strconv.Itoa(numDestination)+"].name", nil, htmlAttrs)
	}
	return StringInput("destination["+strconv.Itoa(numDestination)+"].name", resource.Destination[numDestination].Name, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationTarget(frs []FhirResource, numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].target", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].target", resource.Destination[numDestination].Target, htmlAttrs)
}
func (resource *MessageHeader) T_DestinationReceiver(frs []FhirResource, numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].receiver", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination["+strconv.Itoa(numDestination)+"].receiver", resource.Destination[numDestination].Receiver, htmlAttrs)
}
func (resource *MessageHeader) T_SourceEndpointUrl(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source.endpointUrl", nil, htmlAttrs)
	}
	return StringInput("source.endpointUrl", resource.Source.EndpointUrl, htmlAttrs)
}
func (resource *MessageHeader) T_SourceEndpointReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source.endpointReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source.endpointReference", resource.Source.EndpointReference, htmlAttrs)
}
func (resource *MessageHeader) T_SourceName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source.name", nil, htmlAttrs)
	}
	return StringInput("source.name", resource.Source.Name, htmlAttrs)
}
func (resource *MessageHeader) T_SourceSoftware(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source.software", nil, htmlAttrs)
	}
	return StringInput("source.software", resource.Source.Software, htmlAttrs)
}
func (resource *MessageHeader) T_SourceVersion(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source.version", nil, htmlAttrs)
	}
	return StringInput("source.version", resource.Source.Version, htmlAttrs)
}
func (resource *MessageHeader) T_SourceContact(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ContactPointInput("source.contact", nil, htmlAttrs)
	}
	return ContactPointInput("source.contact", resource.Source.Contact, htmlAttrs)
}
func (resource *MessageHeader) T_ResponseCode(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResponse_code

	if resource == nil {
		return CodeSelect("response.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("response.code", &resource.Response.Code, optionsValueSet, htmlAttrs)
}
func (resource *MessageHeader) T_ResponseDetails(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "response.details", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "response.details", resource.Response.Details, htmlAttrs)
}
