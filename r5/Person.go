package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Person
type Person struct {
	Id                   *string               `json:"id,omitempty"`
	Meta                 *Meta                 `json:"meta,omitempty"`
	ImplicitRules        *string               `json:"implicitRules,omitempty"`
	Language             *string               `json:"language,omitempty"`
	Text                 *Narrative            `json:"text,omitempty"`
	Contained            []Resource            `json:"contained,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `json:"identifier,omitempty"`
	Active               *bool                 `json:"active,omitempty"`
	Name                 []HumanName           `json:"name,omitempty"`
	Telecom              []ContactPoint        `json:"telecom,omitempty"`
	Gender               *string               `json:"gender,omitempty"`
	BirthDate            *string               `json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                 `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string               `json:"deceasedDateTime,omitempty"`
	Address              []Address             `json:"address,omitempty"`
	MaritalStatus        *CodeableConcept      `json:"maritalStatus,omitempty"`
	Photo                []Attachment          `json:"photo,omitempty"`
	Communication        []PersonCommunication `json:"communication,omitempty"`
	ManagingOrganization *Reference            `json:"managingOrganization,omitempty"`
	Link                 []PersonLink          `json:"link,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Person
type PersonCommunication struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Person
type PersonLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Target            Reference   `json:"target"`
	Assurance         *string     `json:"assurance,omitempty"`
}

type OtherPerson Person

// on convert struct to json, automatically add resourceType=Person
func (r Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherPerson:  OtherPerson(r),
		ResourceType: "Person",
	})
}

func (resource *Person) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Person.Id", nil)
	}
	return StringInput("Person.Id", resource.Id)
}
func (resource *Person) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Person.ImplicitRules", nil)
	}
	return StringInput("Person.ImplicitRules", resource.ImplicitRules)
}
func (resource *Person) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Person.Language", nil, optionsValueSet)
	}
	return CodeSelect("Person.Language", resource.Language, optionsValueSet)
}
func (resource *Person) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Person.Active", nil)
	}
	return BoolInput("Person.Active", resource.Active)
}
func (resource *Person) T_Gender() templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("Person.Gender", nil, optionsValueSet)
	}
	return CodeSelect("Person.Gender", resource.Gender, optionsValueSet)
}
func (resource *Person) T_BirthDate() templ.Component {

	if resource == nil {
		return StringInput("Person.BirthDate", nil)
	}
	return StringInput("Person.BirthDate", resource.BirthDate)
}
func (resource *Person) T_MaritalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Person.MaritalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Person.MaritalStatus", resource.MaritalStatus, optionsValueSet)
}
func (resource *Person) T_CommunicationId(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return StringInput("Person.Communication["+strconv.Itoa(numCommunication)+"].Id", nil)
	}
	return StringInput("Person.Communication["+strconv.Itoa(numCommunication)+"].Id", resource.Communication[numCommunication].Id)
}
func (resource *Person) T_CommunicationLanguage(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("Person.Communication["+strconv.Itoa(numCommunication)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Person.Communication["+strconv.Itoa(numCommunication)+"].Language", &resource.Communication[numCommunication].Language, optionsValueSet)
}
func (resource *Person) T_CommunicationPreferred(numCommunication int) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return BoolInput("Person.Communication["+strconv.Itoa(numCommunication)+"].Preferred", nil)
	}
	return BoolInput("Person.Communication["+strconv.Itoa(numCommunication)+"].Preferred", resource.Communication[numCommunication].Preferred)
}
func (resource *Person) T_LinkId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("Person.Link["+strconv.Itoa(numLink)+"].Id", nil)
	}
	return StringInput("Person.Link["+strconv.Itoa(numLink)+"].Id", resource.Link[numLink].Id)
}
func (resource *Person) T_LinkAssurance(numLink int) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel

	if resource == nil || len(resource.Link) >= numLink {
		return CodeSelect("Person.Link["+strconv.Itoa(numLink)+"].Assurance", nil, optionsValueSet)
	}
	return CodeSelect("Person.Link["+strconv.Itoa(numLink)+"].Assurance", resource.Link[numLink].Assurance, optionsValueSet)
}
