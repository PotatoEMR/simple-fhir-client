//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/Endpoint
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
