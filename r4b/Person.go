package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Person
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
	BirthDate            *FhirDate      `json:"birthDate,omitempty"`
	Address              []Address      `json:"address,omitempty"`
	Photo                *Attachment    `json:"photo,omitempty"`
	ManagingOrganization *Reference     `json:"managingOrganization,omitempty"`
	Active               *bool          `json:"active,omitempty"`
	Link                 []PersonLink   `json:"link,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Person
type PersonLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Target            Reference   `json:"target"`
	Assurance         *string     `json:"assurance,omitempty"`
}

type OtherPerson Person

// struct -> json, automatically add resourceType=Patient
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
func (r Person) ResourceType() string {
	return "Person"
}

func (resource *Person) T_Name(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return HumanNameInput("name["+strconv.Itoa(numName)+"]", nil, htmlAttrs)
	}
	return HumanNameInput("name["+strconv.Itoa(numName)+"]", &resource.Name[numName], htmlAttrs)
}
func (resource *Person) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
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
		return FhirDateInput("birthDate", nil, htmlAttrs)
	}
	return FhirDateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Person) T_Address(numAddress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddress >= len(resource.Address) {
		return AddressInput("address["+strconv.Itoa(numAddress)+"]", nil, htmlAttrs)
	}
	return AddressInput("address["+strconv.Itoa(numAddress)+"]", &resource.Address[numAddress], htmlAttrs)
}
func (resource *Person) T_Photo(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AttachmentInput("photo", nil, htmlAttrs)
	}
	return AttachmentInput("photo", resource.Photo, htmlAttrs)
}
func (resource *Person) T_ManagingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization", resource.ManagingOrganization, htmlAttrs)
}
func (resource *Person) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Person) T_LinkTarget(frs []FhirResource, numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return ReferenceInput(frs, "link["+strconv.Itoa(numLink)+"].target", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "link["+strconv.Itoa(numLink)+"].target", &resource.Link[numLink].Target, htmlAttrs)
}
func (resource *Person) T_LinkAssurance(numLink int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSIdentity_assuranceLevel

	if resource == nil || numLink >= len(resource.Link) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].assurance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].assurance", resource.Link[numLink].Assurance, optionsValueSet, htmlAttrs)
}
