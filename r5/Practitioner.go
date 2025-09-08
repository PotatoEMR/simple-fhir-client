package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	BirthDate         *time.Time                  `json:"birthDate,omitempty,format:'2006-01-02'"`
	DeceasedBoolean   *bool                       `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime  *time.Time                  `json:"deceasedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Practitioner) T_Active(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Practitioner.Active", nil, htmlAttrs)
	}
	return BoolInput("Practitioner.Active", resource.Active, htmlAttrs)
}
func (resource *Practitioner) T_Gender(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("Practitioner.Gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Practitioner.Gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *Practitioner) T_BirthDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Practitioner.BirthDate", nil, htmlAttrs)
	}
	return DateInput("Practitioner.BirthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Practitioner) T_DeceasedBoolean(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Practitioner.DeceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("Practitioner.DeceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *Practitioner) T_DeceasedDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Practitioner.DeceasedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Practitioner.DeceasedDateTime", resource.DeceasedDateTime, htmlAttrs)
}
func (resource *Practitioner) T_QualificationCode(numQualification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numQualification >= len(resource.Qualification) {
		return CodeableConceptSelect("Practitioner.Qualification."+strconv.Itoa(numQualification)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Practitioner.Qualification."+strconv.Itoa(numQualification)+"..Code", &resource.Qualification[numQualification].Code, optionsValueSet, htmlAttrs)
}
func (resource *Practitioner) T_CommunicationPreferred(numCommunication int, htmlAttrs string) templ.Component {

	if resource == nil || numCommunication >= len(resource.Communication) {
		return BoolInput("Practitioner.Communication."+strconv.Itoa(numCommunication)+"..Preferred", nil, htmlAttrs)
	}
	return BoolInput("Practitioner.Communication."+strconv.Itoa(numCommunication)+"..Preferred", resource.Communication[numCommunication].Preferred, htmlAttrs)
}
