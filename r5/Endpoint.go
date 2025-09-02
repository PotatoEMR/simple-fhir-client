package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Endpoint) EndpointLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Endpoint) EndpointStatus() templ.Component {
	optionsValueSet := VSEndpoint_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Endpoint) EndpointConnectionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("connectionType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("connectionType", &resource.ConnectionType[0], optionsValueSet)
}
func (resource *Endpoint) EndpointEnvironmentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("environmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("environmentType", &resource.EnvironmentType[0], optionsValueSet)
}
func (resource *Endpoint) EndpointPayloadType(numPayload int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Payload) >= numPayload {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Payload[numPayload].Type[0], optionsValueSet)
}
func (resource *Endpoint) EndpointPayloadMimeType(numPayload int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Payload) >= numPayload {
		return CodeSelect("mimeType", nil, optionsValueSet)
	}
	return CodeSelect("mimeType", &resource.Payload[numPayload].MimeType[0], optionsValueSet)
}
