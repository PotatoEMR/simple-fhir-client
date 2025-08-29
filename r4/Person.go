package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/Person
type Person struct {
	Id                   *string        `json:"id,omitempty"`
	Meta                 *Meta          `json:"meta,omitempty"`
	ImplicitRules        *string        `json:"implicitRules,omitempty"`
	Language             *string        `json:"language,omitempty"`
	Text                 *Narrative     `json:"text,omitempty"`
	Contained            []Resource     `json:"contained,omitempty"`
	Extension            []Extension    `json:"extension,omitempty"`
	ModifierExtension    []Extension    `json:"modifierExtension,omitempty"`
	Identifier           []Identifier   `json:"identifier,omitempty"`
	Name                 []HumanName    `json:"name,omitempty"`
	Telecom              []ContactPoint `json:"telecom,omitempty"`
	Gender               *string        `json:"gender,omitempty"`
	BirthDate            *string        `json:"birthDate,omitempty"`
	Address              []Address      `json:"address,omitempty"`
	Photo                *Attachment    `json:"photo,omitempty"`
	ManagingOrganization *Reference     `json:"managingOrganization,omitempty"`
	Active               *bool          `json:"active,omitempty"`
	Link                 []PersonLink   `json:"link,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Person
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

func (resource *Person) PersonLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Person) PersonGender() templ.Component {
	optionsValueSet := VSAdministrative_gender
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Gender
	}
	return CodeSelect("gender", currentVal, optionsValueSet)
}
func (resource *Person) PersonLinkAssurance(numLink int) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel
	currentVal := ""
	if resource != nil && len(resource.Link) >= numLink {
		currentVal = *resource.Link[numLink].Assurance
	}
	return CodeSelect("assurance", currentVal, optionsValueSet)
}
