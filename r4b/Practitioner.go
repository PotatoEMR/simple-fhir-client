package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Practitioner
type Practitioner struct {
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
	Name              []HumanName                 `json:"name,omitempty"`
	Telecom           []ContactPoint              `json:"telecom,omitempty"`
	Address           []Address                   `json:"address,omitempty"`
	Gender            *string                     `json:"gender,omitempty"`
	BirthDate         *FhirDate                   `json:"birthDate,omitempty"`
	Photo             []Attachment                `json:"photo,omitempty"`
	Qualification     []PractitionerQualification `json:"qualification,omitempty"`
	Communication     []CodeableConcept           `json:"communication,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Practitioner
type PractitionerQualification struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Period            *Period         `json:"period,omitempty"`
	Issuer            *Reference      `json:"issuer,omitempty"`
}

type OtherPractitioner Practitioner

// on convert struct to json, automatically add resourceType=Practitioner
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}
func (r Practitioner) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Practitioner/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Practitioner"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Practitioner) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Practitioner) T_Gender(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *Practitioner) T_BirthDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("birthDate", nil, htmlAttrs)
	}
	return DateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Practitioner) T_Communication(numCommunication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet, htmlAttrs)
}
func (resource *Practitioner) T_QualificationCode(numQualification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", &resource.Qualification[numQualification].Code, optionsValueSet, htmlAttrs)
}
