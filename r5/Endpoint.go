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
func (resource *Endpoint) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEndpoint_status

	if resource == nil {
		return CodeSelect("Endpoint.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Endpoint.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_ConnectionType(numConnectionType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numConnectionType >= len(resource.ConnectionType) {
		return CodeableConceptSelect("Endpoint.ConnectionType."+strconv.Itoa(numConnectionType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Endpoint.ConnectionType."+strconv.Itoa(numConnectionType)+".", &resource.ConnectionType[numConnectionType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Name", nil, htmlAttrs)
	}
	return StringInput("Endpoint.Name", resource.Name, htmlAttrs)
}
func (resource *Endpoint) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Description", nil, htmlAttrs)
	}
	return StringInput("Endpoint.Description", resource.Description, htmlAttrs)
}
func (resource *Endpoint) T_EnvironmentType(numEnvironmentType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numEnvironmentType >= len(resource.EnvironmentType) {
		return CodeableConceptSelect("Endpoint.EnvironmentType."+strconv.Itoa(numEnvironmentType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Endpoint.EnvironmentType."+strconv.Itoa(numEnvironmentType)+".", &resource.EnvironmentType[numEnvironmentType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_Address(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Address", nil, htmlAttrs)
	}
	return StringInput("Endpoint.Address", &resource.Address, htmlAttrs)
}
func (resource *Endpoint) T_Header(numHeader int, htmlAttrs string) templ.Component {

	if resource == nil || numHeader >= len(resource.Header) {
		return StringInput("Endpoint.Header."+strconv.Itoa(numHeader)+".", nil, htmlAttrs)
	}
	return StringInput("Endpoint.Header."+strconv.Itoa(numHeader)+".", &resource.Header[numHeader], htmlAttrs)
}
func (resource *Endpoint) T_PayloadType(numPayload int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPayload >= len(resource.Payload) || numType >= len(resource.Payload[numPayload].Type) {
		return CodeableConceptSelect("Endpoint.Payload."+strconv.Itoa(numPayload)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Endpoint.Payload."+strconv.Itoa(numPayload)+"..Type."+strconv.Itoa(numType)+".", &resource.Payload[numPayload].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_PayloadMimeType(numPayload int, numMimeType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPayload >= len(resource.Payload) || numMimeType >= len(resource.Payload[numPayload].MimeType) {
		return CodeSelect("Endpoint.Payload."+strconv.Itoa(numPayload)+"..MimeType."+strconv.Itoa(numMimeType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Endpoint.Payload."+strconv.Itoa(numPayload)+"..MimeType."+strconv.Itoa(numMimeType)+".", &resource.Payload[numPayload].MimeType[numMimeType], optionsValueSet, htmlAttrs)
}
