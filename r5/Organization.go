package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Organization) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Organization.Id", nil)
	}
	return StringInput("Organization.Id", resource.Id)
}
func (resource *Organization) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Organization.ImplicitRules", nil)
	}
	return StringInput("Organization.ImplicitRules", resource.ImplicitRules)
}
func (resource *Organization) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Organization.Language", nil, optionsValueSet)
	}
	return CodeSelect("Organization.Language", resource.Language, optionsValueSet)
}
func (resource *Organization) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Organization.Active", nil)
	}
	return BoolInput("Organization.Active", resource.Active)
}
func (resource *Organization) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Organization.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Organization.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Organization) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Organization.Name", nil)
	}
	return StringInput("Organization.Name", resource.Name)
}
func (resource *Organization) T_Alias(numAlias int) templ.Component {

	if resource == nil || len(resource.Alias) >= numAlias {
		return StringInput("Organization.Alias["+strconv.Itoa(numAlias)+"]", nil)
	}
	return StringInput("Organization.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias])
}
func (resource *Organization) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Organization.Description", nil)
	}
	return StringInput("Organization.Description", resource.Description)
}
func (resource *Organization) T_QualificationId(numQualification int) templ.Component {

	if resource == nil || len(resource.Qualification) >= numQualification {
		return StringInput("Organization.Qualification["+strconv.Itoa(numQualification)+"].Id", nil)
	}
	return StringInput("Organization.Qualification["+strconv.Itoa(numQualification)+"].Id", resource.Qualification[numQualification].Id)
}
func (resource *Organization) T_QualificationCode(numQualification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Qualification) >= numQualification {
		return CodeableConceptSelect("Organization.Qualification["+strconv.Itoa(numQualification)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Organization.Qualification["+strconv.Itoa(numQualification)+"].Code", &resource.Qualification[numQualification].Code, optionsValueSet)
}
