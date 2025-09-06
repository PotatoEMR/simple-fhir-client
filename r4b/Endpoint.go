package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Endpoint
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
	ConnectionType       Coding            `json:"connectionType"`
	Name                 *string           `json:"name,omitempty"`
	ManagingOrganization *Reference        `json:"managingOrganization,omitempty"`
	Contact              []ContactPoint    `json:"contact,omitempty"`
	Period               *Period           `json:"period,omitempty"`
	PayloadType          []CodeableConcept `json:"payloadType"`
	PayloadMimeType      []string          `json:"payloadMimeType,omitempty"`
	Address              string            `json:"address"`
	Header               []string          `json:"header,omitempty"`
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
func (resource *Endpoint) T_ConnectionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("Endpoint.ConnectionType", nil, optionsValueSet)
	}
	return CodingSelect("Endpoint.ConnectionType", &resource.ConnectionType, optionsValueSet)
}
func (resource *Endpoint) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Endpoint.Name", nil)
	}
	return StringInput("Endpoint.Name", resource.Name)
}
func (resource *Endpoint) T_PayloadType(numPayloadType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PayloadType) >= numPayloadType {
		return CodeableConceptSelect("Endpoint.PayloadType["+strconv.Itoa(numPayloadType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Endpoint.PayloadType["+strconv.Itoa(numPayloadType)+"]", &resource.PayloadType[numPayloadType], optionsValueSet)
}
func (resource *Endpoint) T_PayloadMimeType(numPayloadMimeType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PayloadMimeType) >= numPayloadMimeType {
		return CodeSelect("Endpoint.PayloadMimeType["+strconv.Itoa(numPayloadMimeType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("Endpoint.PayloadMimeType["+strconv.Itoa(numPayloadMimeType)+"]", &resource.PayloadMimeType[numPayloadMimeType], optionsValueSet)
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
