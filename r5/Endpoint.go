package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Endpoint
type Endpoint struct {
	Id                   *string           `json:"id,omitempty"`
	Meta                 *Meta             `json:"meta,omitempty"`
	ImplicitRules        *string           `json:"implicitRules,omitempty"`
	Language             *string           `json:"language,omitempty"`
	Text                 *Narrative        `json:"text,omitempty"`
	Contained            []Resource        `json:"contained,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `json:"identifier,omitempty"`
	Status               string            `json:"status"`
	ConnectionType       []CodeableConcept `json:"connectionType"`
	Name                 *string           `json:"name,omitempty"`
	Description          *string           `json:"description,omitempty"`
	EnvironmentType      []CodeableConcept `json:"environmentType,omitempty"`
	ManagingOrganization *Reference        `json:"managingOrganization,omitempty"`
	Contact              []ContactPoint    `json:"contact,omitempty"`
	Period               *Period           `json:"period,omitempty"`
	Payload              []EndpointPayload `json:"payload,omitempty"`
	Address              string            `json:"address"`
	Header               []string          `json:"header,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Endpoint
type EndpointPayload struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	MimeType          []string          `json:"mimeType,omitempty"`
}

type OtherEndpoint Endpoint

// on convert struct to json, automatically add resourceType=Endpoint
func (r Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEndpoint
		ResourceType string `json:"resourceType"`
	}{
		OtherEndpoint: OtherEndpoint(r),
		ResourceType:  "Endpoint",
	})
}
func (r Endpoint) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Endpoint/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Endpoint"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Endpoint) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEndpoint_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_ConnectionType(numConnectionType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConnectionType >= len(resource.ConnectionType) {
		return CodeableConceptSelect("connectionType["+strconv.Itoa(numConnectionType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("connectionType["+strconv.Itoa(numConnectionType)+"]", &resource.ConnectionType[numConnectionType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Endpoint) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Endpoint) T_EnvironmentType(numEnvironmentType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEnvironmentType >= len(resource.EnvironmentType) {
		return CodeableConceptSelect("environmentType["+strconv.Itoa(numEnvironmentType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("environmentType["+strconv.Itoa(numEnvironmentType)+"]", &resource.EnvironmentType[numEnvironmentType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_ManagingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization", resource.ManagingOrganization, htmlAttrs)
}
func (resource *Endpoint) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Endpoint) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *Endpoint) T_Address(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("address", nil, htmlAttrs)
	}
	return StringInput("address", &resource.Address, htmlAttrs)
}
func (resource *Endpoint) T_Header(numHeader int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHeader >= len(resource.Header) {
		return StringInput("header["+strconv.Itoa(numHeader)+"]", nil, htmlAttrs)
	}
	return StringInput("header["+strconv.Itoa(numHeader)+"]", &resource.Header[numHeader], htmlAttrs)
}
func (resource *Endpoint) T_PayloadType(numPayload int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) || numType >= len(resource.Payload[numPayload].Type) {
		return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payload["+strconv.Itoa(numPayload)+"].type["+strconv.Itoa(numType)+"]", &resource.Payload[numPayload].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_PayloadMimeType(numPayload int, numMimeType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPayload >= len(resource.Payload) || numMimeType >= len(resource.Payload[numPayload].MimeType) {
		return CodeSelect("payload["+strconv.Itoa(numPayload)+"].mimeType["+strconv.Itoa(numMimeType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("payload["+strconv.Itoa(numPayload)+"].mimeType["+strconv.Itoa(numMimeType)+"]", &resource.Payload[numPayload].MimeType[numMimeType], optionsValueSet, htmlAttrs)
}
