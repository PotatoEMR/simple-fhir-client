package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Organization
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

// http://hl7.org/fhir/r4/StructureDefinition/Organization
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

// on convert struct to json, automatically add resourceType=Organization
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
func (resource *Organization) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Active", nil, htmlAttrs)
	}
	return BoolInput("Active", resource.Active, htmlAttrs)
}
func (resource *Organization) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Organization) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *Organization) T_Alias(numAlias int, htmlAttrs string) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("Alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Organization) T_ContactPurpose(numContact int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("Contact["+strconv.Itoa(numContact)+"]Purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contact["+strconv.Itoa(numContact)+"]Purpose", resource.Contact[numContact].Purpose, optionsValueSet, htmlAttrs)
}
