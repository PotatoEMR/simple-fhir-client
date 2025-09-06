package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *PractitionerRole) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PractitionerRole.Id", nil)
	}
	return StringInput("PractitionerRole.Id", resource.Id)
}
func (resource *PractitionerRole) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PractitionerRole.ImplicitRules", nil)
	}
	return StringInput("PractitionerRole.ImplicitRules", resource.ImplicitRules)
}
func (resource *PractitionerRole) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PractitionerRole.Language", nil, optionsValueSet)
	}
	return CodeSelect("PractitionerRole.Language", resource.Language, optionsValueSet)
}
func (resource *PractitionerRole) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("PractitionerRole.Active", nil)
	}
	return BoolInput("PractitionerRole.Active", resource.Active)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *PractitionerRole) T_Characteristic(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("PractitionerRole.Characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet)
}
func (resource *PractitionerRole) T_Communication(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("PractitionerRole.Communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet)
}
