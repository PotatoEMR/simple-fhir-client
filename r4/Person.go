package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	BirthDate            *time.Time     `json:"birthDate,omitempty,format:'2006-01-02'"`
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
func (r Person) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Person/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Person"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Person) T_Gender(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("Gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *Person) T_BirthDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("BirthDate", nil, htmlAttrs)
	}
	return DateInput("BirthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Person) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Active", nil, htmlAttrs)
	}
	return BoolInput("Active", resource.Active, htmlAttrs)
}
func (resource *Person) T_LinkAssurance(numLink int, htmlAttrs string) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel

	if resource == nil || numLink >= len(resource.Link) {
		return CodeSelect("Link["+strconv.Itoa(numLink)+"]Assurance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Link["+strconv.Itoa(numLink)+"]Assurance", resource.Link[numLink].Assurance, optionsValueSet, htmlAttrs)
}
