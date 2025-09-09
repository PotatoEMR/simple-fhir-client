package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Organization
type Organization struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Active            *bool                       `json:"active,omitempty"`
	Type              []CodeableConcept           `json:"type,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	Alias             []string                    `json:"alias,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Contact           []ExtendedContactDetail     `json:"contact,omitempty"`
	PartOf            *Reference                  `json:"partOf,omitempty"`
	Endpoint          []Reference                 `json:"endpoint,omitempty"`
	Qualification     []OrganizationQualification `json:"qualification,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Organization
type OrganizationQualification struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Period            *Period         `json:"period,omitempty"`
	Issuer            *Reference      `json:"issuer,omitempty"`
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
		return BoolInput("Organization.Active", nil, htmlAttrs)
	}
	return BoolInput("Organization.Active", resource.Active, htmlAttrs)
}
func (resource *Organization) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Organization.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Organization.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Organization) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Organization.Name", nil, htmlAttrs)
	}
	return StringInput("Organization.Name", resource.Name, htmlAttrs)
}
func (resource *Organization) T_Alias(numAlias int, htmlAttrs string) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("Organization.Alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("Organization.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Organization) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Organization.Description", nil, htmlAttrs)
	}
	return StringInput("Organization.Description", resource.Description, htmlAttrs)
}
func (resource *Organization) T_QualificationCode(numQualification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return CodeableConceptSelect("Organization.Qualification["+strconv.Itoa(numQualification)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Organization.Qualification["+strconv.Itoa(numQualification)+"].Code", &resource.Qualification[numQualification].Code, optionsValueSet, htmlAttrs)
}
