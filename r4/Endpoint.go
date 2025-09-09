package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_ConnectionType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("connectionType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("connectionType", &resource.ConnectionType, optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Endpoint) T_PayloadType(numPayloadType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPayloadType >= len(resource.PayloadType) {
		return CodeableConceptSelect("payloadType["+strconv.Itoa(numPayloadType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payloadType["+strconv.Itoa(numPayloadType)+"]", &resource.PayloadType[numPayloadType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_PayloadMimeType(numPayloadMimeType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPayloadMimeType >= len(resource.PayloadMimeType) {
		return CodeSelect("payloadMimeType["+strconv.Itoa(numPayloadMimeType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("payloadMimeType["+strconv.Itoa(numPayloadMimeType)+"]", &resource.PayloadMimeType[numPayloadMimeType], optionsValueSet, htmlAttrs)
}
func (resource *Endpoint) T_Address(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("address", nil, htmlAttrs)
	}
	return StringInput("address", &resource.Address, htmlAttrs)
}
func (resource *Endpoint) T_Header(numHeader int, htmlAttrs string) templ.Component {
	if resource == nil || numHeader >= len(resource.Header) {
		return StringInput("header["+strconv.Itoa(numHeader)+"]", nil, htmlAttrs)
	}
	return StringInput("header["+strconv.Itoa(numHeader)+"]", &resource.Header[numHeader], htmlAttrs)
}
