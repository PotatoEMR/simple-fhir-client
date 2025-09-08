package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Active            *bool                   `json:"active,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	Practitioner      *Reference              `json:"practitioner,omitempty"`
	Organization      *Reference              `json:"organization,omitempty"`
	Code              []CodeableConcept       `json:"code,omitempty"`
	Specialty         []CodeableConcept       `json:"specialty,omitempty"`
	Location          []Reference             `json:"location,omitempty"`
	HealthcareService []Reference             `json:"healthcareService,omitempty"`
	Contact           []ExtendedContactDetail `json:"contact,omitempty"`
	Characteristic    []CodeableConcept       `json:"characteristic,omitempty"`
	Communication     []CodeableConcept       `json:"communication,omitempty"`
	Availability      []Availability          `json:"availability,omitempty"`
	Endpoint          []Reference             `json:"endpoint,omitempty"`
}

type OtherPractitionerRole PractitionerRole

// on convert struct to json, automatically add resourceType=PractitionerRole
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}
func (r PractitionerRole) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PractitionerRole/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PractitionerRole"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PractitionerRole) T_Active(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("PractitionerRole.Active", nil, htmlAttrs)
	}
	return BoolInput("PractitionerRole.Active", resource.Active, htmlAttrs)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("PractitionerRole.Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Code."+strconv.Itoa(numCode)+".", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("PractitionerRole.Specialty."+strconv.Itoa(numSpecialty)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Specialty."+strconv.Itoa(numSpecialty)+".", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("PractitionerRole.Characteristic."+strconv.Itoa(numCharacteristic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Characteristic."+strconv.Itoa(numCharacteristic)+".", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Communication(numCommunication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCommunication >= len(resource.Communication) {
		return CodeableConceptSelect("PractitionerRole.Communication."+strconv.Itoa(numCommunication)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Communication."+strconv.Itoa(numCommunication)+".", &resource.Communication[numCommunication], optionsValueSet, htmlAttrs)
}
