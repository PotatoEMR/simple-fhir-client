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
func (resource *Organization) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Organization) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
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
func (resource *Organization) T_QualificationCode(numQualification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", &resource.Qualification[numQualification].Code, optionsValueSet, htmlAttrs)
}
func (resource *Organization) T_QualificationPeriod(numQualification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return PeriodInput("qualification["+strconv.Itoa(numQualification)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("qualification["+strconv.Itoa(numQualification)+"].period", resource.Qualification[numQualification].Period, htmlAttrs)
}
func (resource *Organization) T_QualificationIssuer(frs []FhirResource, numQualification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return ReferenceInput(frs, "qualification["+strconv.Itoa(numQualification)+"].issuer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "qualification["+strconv.Itoa(numQualification)+"].issuer", resource.Qualification[numQualification].Issuer, htmlAttrs)
}
