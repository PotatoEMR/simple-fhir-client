package r5

//generated with command go run ./bultaoreune
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
	BirthDate            *FhirDate             `json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                 `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *FhirDateTime         `json:"deceasedDateTime,omitempty"`
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
func (resource *Person) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Person) T_Gender(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *Person) T_BirthDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("birthDate", nil, htmlAttrs)
	}
	return DateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Person) T_DeceasedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("deceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("deceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *Person) T_DeceasedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("deceasedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("deceasedDateTime", resource.DeceasedDateTime, htmlAttrs)
}
func (resource *Person) T_MaritalStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("maritalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("maritalStatus", resource.MaritalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Person) T_CommunicationPreferred(numCommunication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", resource.Communication[numCommunication].Preferred, htmlAttrs)
}
func (resource *Person) T_LinkAssurance(numLink int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel

	if resource == nil || numLink >= len(resource.Link) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].assurance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].assurance", resource.Link[numLink].Assurance, optionsValueSet, htmlAttrs)
}
