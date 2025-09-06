package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Practitioner
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
	Gender            *string                     `json:"gender,omitempty"`
	BirthDate         *string                     `json:"birthDate,omitempty"`
	DeceasedBoolean   *bool                       `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime  *string                     `json:"deceasedDateTime,omitempty"`
	Address           []Address                   `json:"address,omitempty"`
	Photo             []Attachment                `json:"photo,omitempty"`
	Qualification     []PractitionerQualification `json:"qualification,omitempty"`
	Communication     []PractitionerCommunication `json:"communication,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Practitioner
type PractitionerQualification struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Period            *Period         `json:"period,omitempty"`
	Issuer            *Reference      `json:"issuer,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Practitioner
type PractitionerCommunication struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
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

func (resource *Practitioner) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Practitioner.Id", nil)
	}
	return StringInput("Practitioner.Id", resource.Id)
}
func (resource *Practitioner) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Practitioner.ImplicitRules", nil)
	}
	return StringInput("Practitioner.ImplicitRules", resource.ImplicitRules)
}
func (resource *Practitioner) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Practitioner.Language", nil, optionsValueSet)
	}
	return CodeSelect("Practitioner.Language", resource.Language, optionsValueSet)
}
func (resource *Practitioner) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Practitioner.Active", nil)
	}
	return BoolInput("Practitioner.Active", resource.Active)
}
func (resource *Practitioner) T_Gender() templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("Practitioner.Gender", nil, optionsValueSet)
	}
	return CodeSelect("Practitioner.Gender", resource.Gender, optionsValueSet)
}
func (resource *Practitioner) T_BirthDate() templ.Component {

	if resource == nil {
		return StringInput("Practitioner.BirthDate", nil)
	}
	return StringInput("Practitioner.BirthDate", resource.BirthDate)
}
func (resource *Practitioner) T_QualificationId(numQualification int) templ.Component {

	if resource == nil || len(resource.Qualification) >= numQualification {
		return StringInput("Practitioner.Qualification["+strconv.Itoa(numQualification)+"].Id", nil)
	}
	return StringInput("Practitioner.Qualification["+strconv.Itoa(numQualification)+"].Id", resource.Qualification[numQualification].Id)
}
func (resource *Practitioner) T_QualificationCode(numQualification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Qualification) >= numQualification {
		return CodeableConceptSelect("Practitioner.Qualification["+strconv.Itoa(numQualification)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Practitioner.Qualification["+strconv.Itoa(numQualification)+"].Code", &resource.Qualification[numQualification].Code, optionsValueSet)
}
func (resource *Practitioner) T_CommunicationId(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return StringInput("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Id", nil)
	}
	return StringInput("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Id", resource.Communication[numCommunication].Id)
}
func (resource *Practitioner) T_CommunicationLanguage(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Language", &resource.Communication[numCommunication].Language, optionsValueSet)
}
func (resource *Practitioner) T_CommunicationPreferred(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return BoolInput("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Preferred", nil)
	}
	return BoolInput("Practitioner.Communication["+strconv.Itoa(numCommunication)+"].Preferred", resource.Communication[numCommunication].Preferred)
}
