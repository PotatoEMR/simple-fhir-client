package r4

//generated with command go run ./bultaoreune -nodownload
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
	BirthDate         *string                      `json:"birthDate,omitempty"`
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

// on convert struct to json, automatically add resourceType=RelatedPerson
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
}

func (resource *RelatedPerson) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RelatedPerson.Id", nil)
	}
	return StringInput("RelatedPerson.Id", resource.Id)
}
func (resource *RelatedPerson) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RelatedPerson.ImplicitRules", nil)
	}
	return StringInput("RelatedPerson.ImplicitRules", resource.ImplicitRules)
}
func (resource *RelatedPerson) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RelatedPerson.Language", nil, optionsValueSet)
	}
	return CodeSelect("RelatedPerson.Language", resource.Language, optionsValueSet)
}
func (resource *RelatedPerson) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("RelatedPerson.Active", nil)
	}
	return BoolInput("RelatedPerson.Active", resource.Active)
}
func (resource *RelatedPerson) T_Relationship(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("RelatedPerson.Relationship["+strconv.Itoa(numRelationship)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RelatedPerson.Relationship["+strconv.Itoa(numRelationship)+"]", &resource.Relationship[numRelationship], optionsValueSet)
}
func (resource *RelatedPerson) T_Gender() templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("RelatedPerson.Gender", nil, optionsValueSet)
	}
	return CodeSelect("RelatedPerson.Gender", resource.Gender, optionsValueSet)
}
func (resource *RelatedPerson) T_BirthDate() templ.Component {

	if resource == nil {
		return StringInput("RelatedPerson.BirthDate", nil)
	}
	return StringInput("RelatedPerson.BirthDate", resource.BirthDate)
}
func (resource *RelatedPerson) T_CommunicationId(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return StringInput("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Id", nil)
	}
	return StringInput("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Id", resource.Communication[numCommunication].Id)
}
func (resource *RelatedPerson) T_CommunicationLanguage(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Language", &resource.Communication[numCommunication].Language, optionsValueSet)
}
func (resource *RelatedPerson) T_CommunicationPreferred(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return BoolInput("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Preferred", nil)
	}
	return BoolInput("RelatedPerson.Communication["+strconv.Itoa(numCommunication)+"].Preferred", resource.Communication[numCommunication].Preferred)
}
