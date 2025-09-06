package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Endpoint) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Id", nil)
	}
	return StringInput("Endpoint.Id", resource.Id)
}
func (resource *Endpoint) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.ImplicitRules", nil)
	}
	return StringInput("Endpoint.ImplicitRules", resource.ImplicitRules)
}
func (resource *Endpoint) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Endpoint.Language", nil, optionsValueSet)
	}
	return CodeSelect("Endpoint.Language", resource.Language, optionsValueSet)
}
func (resource *Endpoint) T_Status() templ.Component {
	optionsValueSet := VSEndpoint_status

	if resource == nil {
		return CodeSelect("Endpoint.Status", nil, optionsValueSet)
	}
	return CodeSelect("Endpoint.Status", &resource.Status, optionsValueSet)
}
func (resource *Endpoint) T_ConnectionType(numConnectionType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConnectionType) >= numConnectionType {
		return CodeableConceptSelect("Endpoint.ConnectionType["+strconv.Itoa(numConnectionType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Endpoint.ConnectionType["+strconv.Itoa(numConnectionType)+"]", &resource.ConnectionType[numConnectionType], optionsValueSet)
}
func (resource *Endpoint) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Name", nil)
	}
	return StringInput("Endpoint.Name", resource.Name)
}
func (resource *Endpoint) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Description", nil)
	}
	return StringInput("Endpoint.Description", resource.Description)
}
func (resource *Endpoint) T_EnvironmentType(numEnvironmentType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EnvironmentType) >= numEnvironmentType {
		return CodeableConceptSelect("Endpoint.EnvironmentType["+strconv.Itoa(numEnvironmentType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Endpoint.EnvironmentType["+strconv.Itoa(numEnvironmentType)+"]", &resource.EnvironmentType[numEnvironmentType], optionsValueSet)
}
func (resource *Endpoint) T_Address() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Address", nil)
	}
	return StringInput("Endpoint.Address", &resource.Address)
}
func (resource *Endpoint) T_Header(numHeader int) templ.Component {

	if resource == nil || len(resource.Header) >= numHeader {
		return StringInput("Endpoint.Header["+strconv.Itoa(numHeader)+"]", nil)
	}
	return StringInput("Endpoint.Header["+strconv.Itoa(numHeader)+"]", &resource.Header[numHeader])
}
func (resource *Endpoint) T_PayloadId(numPayload int) templ.Component {

	if resource == nil || len(resource.Payload) >= numPayload {
		return StringInput("Endpoint.Payload["+strconv.Itoa(numPayload)+"].Id", nil)
	}
	return StringInput("Endpoint.Payload["+strconv.Itoa(numPayload)+"].Id", resource.Payload[numPayload].Id)
}
func (resource *Endpoint) T_PayloadType(numPayload int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Payload) >= numPayload || len(resource.Payload[numPayload].Type) >= numType {
		return CodeableConceptSelect("Endpoint.Payload["+strconv.Itoa(numPayload)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Endpoint.Payload["+strconv.Itoa(numPayload)+"].Type["+strconv.Itoa(numType)+"]", &resource.Payload[numPayload].Type[numType], optionsValueSet)
}
func (resource *Endpoint) T_PayloadMimeType(numPayload int, numMimeType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Payload) >= numPayload || len(resource.Payload[numPayload].MimeType) >= numMimeType {
		return CodeSelect("Endpoint.Payload["+strconv.Itoa(numPayload)+"].MimeType["+strconv.Itoa(numMimeType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("Endpoint.Payload["+strconv.Itoa(numPayload)+"].MimeType["+strconv.Itoa(numMimeType)+"]", &resource.Payload[numPayload].MimeType[numMimeType], optionsValueSet)
}
