package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	BirthDate         *FhirDate                   `json:"birthDate,omitempty"`
	DeceasedBoolean   *bool                       `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime  *FhirDateTime               `json:"deceasedDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// json -> struct, first reject if resourceType != Practitioner
func (r *Practitioner) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Practitioner" {
		return errors.New("resourceType not Practitioner")
	}
	return json.Unmarshal(data, (*OtherPractitioner)(r))
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
func (resource *Practitioner) T_Name(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return HumanNameInput("name["+strconv.Itoa(numName)+"]", nil, htmlAttrs)
	}
	return HumanNameInput("name["+strconv.Itoa(numName)+"]", &resource.Name[numName], htmlAttrs)
}
func (resource *Practitioner) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
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
		return FhirDateInput("birthDate", nil, htmlAttrs)
	}
	return FhirDateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Practitioner) T_DeceasedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("deceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("deceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *Practitioner) T_DeceasedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("deceasedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("deceasedDateTime", resource.DeceasedDateTime, htmlAttrs)
}
func (resource *Practitioner) T_Address(numAddress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddress >= len(resource.Address) {
		return AddressInput("address["+strconv.Itoa(numAddress)+"]", nil, htmlAttrs)
	}
	return AddressInput("address["+strconv.Itoa(numAddress)+"]", &resource.Address[numAddress], htmlAttrs)
}
func (resource *Practitioner) T_Photo(numPhoto int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPhoto >= len(resource.Photo) {
		return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", nil, htmlAttrs)
	}
	return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", &resource.Photo[numPhoto], htmlAttrs)
}
func (resource *Practitioner) T_QualificationCode(numQualification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("qualification["+strconv.Itoa(numQualification)+"].code", &resource.Qualification[numQualification].Code, optionsValueSet, htmlAttrs)
}
func (resource *Practitioner) T_QualificationPeriod(numQualification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return PeriodInput("qualification["+strconv.Itoa(numQualification)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("qualification["+strconv.Itoa(numQualification)+"].period", resource.Qualification[numQualification].Period, htmlAttrs)
}
func (resource *Practitioner) T_QualificationIssuer(frs []FhirResource, numQualification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualification >= len(resource.Qualification) {
		return ReferenceInput(frs, "qualification["+strconv.Itoa(numQualification)+"].issuer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "qualification["+strconv.Itoa(numQualification)+"].issuer", resource.Qualification[numQualification].Issuer, htmlAttrs)
}
func (resource *Practitioner) T_CommunicationPreferred(numCommunication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", resource.Communication[numCommunication].Preferred, htmlAttrs)
}
