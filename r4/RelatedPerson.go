package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []Resource                   `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Active            *bool                        `json:"active,omitempty"`
	Patient           Reference                    `json:"patient"`
	Relationship      []CodeableConcept            `json:"relationship,omitempty"`
	Name              []HumanName                  `json:"name,omitempty"`
	Telecom           []ContactPoint               `json:"telecom,omitempty"`
	Gender            *string                      `json:"gender,omitempty"`
	BirthDate         *FhirDate                    `json:"birthDate,omitempty"`
	Address           []Address                    `json:"address,omitempty"`
	Photo             []Attachment                 `json:"photo,omitempty"`
	Period            *Period                      `json:"period,omitempty"`
	Communication     []RelatedPersonCommunication `json:"communication,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RelatedPerson
type RelatedPersonCommunication struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
}

type OtherRelatedPerson RelatedPerson

// struct -> json, automatically add resourceType=Patient
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
}

func (r RelatedPerson) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RelatedPerson/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RelatedPerson"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r RelatedPerson) ResourceType() string {
	return "RelatedPerson"
}

func (resource *RelatedPerson) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *RelatedPerson) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *RelatedPerson) T_Relationship(numRelationship int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"]", &resource.Relationship[numRelationship], optionsValueSet, htmlAttrs)
}
func (resource *RelatedPerson) T_Name(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return HumanNameInput("name["+strconv.Itoa(numName)+"]", nil, htmlAttrs)
	}
	return HumanNameInput("name["+strconv.Itoa(numName)+"]", &resource.Name[numName], htmlAttrs)
}
func (resource *RelatedPerson) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *RelatedPerson) T_Gender(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *RelatedPerson) T_BirthDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("birthDate", nil, htmlAttrs)
	}
	return FhirDateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *RelatedPerson) T_Address(numAddress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddress >= len(resource.Address) {
		return AddressInput("address["+strconv.Itoa(numAddress)+"]", nil, htmlAttrs)
	}
	return AddressInput("address["+strconv.Itoa(numAddress)+"]", &resource.Address[numAddress], htmlAttrs)
}
func (resource *RelatedPerson) T_Photo(numPhoto int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPhoto >= len(resource.Photo) {
		return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", nil, htmlAttrs)
	}
	return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", &resource.Photo[numPhoto], htmlAttrs)
}
func (resource *RelatedPerson) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *RelatedPerson) T_CommunicationPreferred(numCommunication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", resource.Communication[numCommunication].Preferred, htmlAttrs)
}
