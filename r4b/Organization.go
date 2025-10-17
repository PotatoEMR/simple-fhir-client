package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Organization
type Organization struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Active            *bool                 `json:"active,omitempty"`
	Type              []CodeableConcept     `json:"type,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Alias             []string              `json:"alias,omitempty"`
	Telecom           []ContactPoint        `json:"telecom,omitempty"`
	Address           []Address             `json:"address,omitempty"`
	PartOf            *Reference            `json:"partOf,omitempty"`
	Contact           []OrganizationContact `json:"contact,omitempty"`
	Endpoint          []Reference           `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Organization
type OrganizationContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}

type OtherOrganization Organization

// struct -> json, automatically add resourceType=Patient
func (r Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganization
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganization: OtherOrganization(r),
		ResourceType:      "Organization",
	})
}

func (r Organization) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Organization/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Organization"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r Organization) ResourceType() string {
	return "Organization"
}

func (resource *Organization) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Organization) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Organization) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Organization) T_Alias(numAlias int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Organization) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *Organization) T_Address(numAddress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddress >= len(resource.Address) {
		return AddressInput("address["+strconv.Itoa(numAddress)+"]", nil, htmlAttrs)
	}
	return AddressInput("address["+strconv.Itoa(numAddress)+"]", &resource.Address[numAddress], htmlAttrs)
}
func (resource *Organization) T_PartOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "partOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf", resource.PartOf, htmlAttrs)
}
func (resource *Organization) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
func (resource *Organization) T_ContactPurpose(numContact int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].purpose", resource.Contact[numContact].Purpose, optionsValueSet, htmlAttrs)
}
func (resource *Organization) T_ContactName(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", nil, htmlAttrs)
	}
	return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", resource.Contact[numContact].Name, htmlAttrs)
}
func (resource *Organization) T_ContactTelecom(numContact int, numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) || numTelecom >= len(resource.Contact[numContact].Telecom) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", &resource.Contact[numContact].Telecom[numTelecom], htmlAttrs)
}
func (resource *Organization) T_ContactAddress(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return AddressInput("contact["+strconv.Itoa(numContact)+"].address", nil, htmlAttrs)
	}
	return AddressInput("contact["+strconv.Itoa(numContact)+"].address", resource.Contact[numContact].Address, htmlAttrs)
}
