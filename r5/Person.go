package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	DeceasedBoolean      *bool                 `json:"deceasedBoolean"`
	DeceasedDateTime     *string               `json:"deceasedDateTime"`
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

func (resource *Person) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Person) T_Gender() templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("gender", nil, optionsValueSet)
	}
	return CodeSelect("gender", resource.Gender, optionsValueSet)
}
func (resource *Person) T_MaritalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("maritalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("maritalStatus", resource.MaritalStatus, optionsValueSet)
}
func (resource *Person) T_CommunicationLanguage(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.Communication[numCommunication].Language, optionsValueSet)
}
func (resource *Person) T_LinkAssurance(numLink int) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel

	if resource == nil && len(resource.Link) >= numLink {
		return CodeSelect("assurance", nil, optionsValueSet)
	}
	return CodeSelect("assurance", resource.Link[numLink].Assurance, optionsValueSet)
}
