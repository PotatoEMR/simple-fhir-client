package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Patient
type Patient struct {
	Id                   *string                `json:"id,omitempty"`
	Meta                 *Meta                  `json:"meta,omitempty"`
	ImplicitRules        *string                `json:"implicitRules,omitempty"`
	Language             *string                `json:"language,omitempty"`
	Text                 *Narrative             `json:"text,omitempty"`
	Contained            []Resource             `json:"contained,omitempty"`
	Extension            []Extension            `json:"extension,omitempty"`
	ModifierExtension    []Extension            `json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `json:"identifier,omitempty"`
	Active               *bool                  `json:"active,omitempty"`
	Name                 []HumanName            `json:"name,omitempty"`
	Telecom              []ContactPoint         `json:"telecom,omitempty"`
	Gender               *string                `json:"gender,omitempty"`
	BirthDate            *FhirDate              `json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                  `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *FhirDateTime          `json:"deceasedDateTime,omitempty"`
	Address              []Address              `json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                  `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int                   `json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment           `json:"photo,omitempty"`
	Contact              []PatientContact       `json:"contact,omitempty"`
	Communication        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []Reference            `json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference             `json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `json:"link,omitempty"`
}

func (p Patient) String() string {
	var b strings.Builder
	lastNewline := len(p.Name) - 1
	for i, v := range p.Name {
		b.WriteString(v.String())
		if i != lastNewline {
			b.WriteString("\n")
		}
	}
	if b.String() == "" {
		return "Unnamed Patient"
	}
	return b.String()
}

// http://hl7.org/fhir/r4/StructureDefinition/Patient
type PatientContact struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept `json:"relationship,omitempty"`
	Name              *HumanName        `json:"name,omitempty"`
	Telecom           []ContactPoint    `json:"telecom,omitempty"`
	Address           *Address          `json:"address,omitempty"`
	Gender            *string           `json:"gender,omitempty"`
	Organization      *Reference        `json:"organization,omitempty"`
	Period            *Period           `json:"period,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Patient
type PatientCommunication struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Patient
type PatientLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Other             Reference   `json:"other"`
	Type              string      `json:"type"`
}

type OtherPatient Patient

// struct -> json, automatically add resourceType=Patient
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

func (r Patient) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Patient/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Patient"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r Patient) ResourceType() string {
	return "Patient"
}

func (resource *Patient) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Patient) T_Name(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return HumanNameInput("name["+strconv.Itoa(numName)+"]", nil, htmlAttrs)
	}
	return HumanNameInput("name["+strconv.Itoa(numName)+"]", &resource.Name[numName], htmlAttrs)
}
func (resource *Patient) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *Patient) T_Gender(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil {
		return CodeSelect("gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("gender", resource.Gender, optionsValueSet, htmlAttrs)
}
func (resource *Patient) T_BirthDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("birthDate", nil, htmlAttrs)
	}
	return FhirDateInput("birthDate", resource.BirthDate, htmlAttrs)
}
func (resource *Patient) T_DeceasedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("deceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("deceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *Patient) T_DeceasedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("deceasedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("deceasedDateTime", resource.DeceasedDateTime, htmlAttrs)
}
func (resource *Patient) T_Address(numAddress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddress >= len(resource.Address) {
		return AddressInput("address["+strconv.Itoa(numAddress)+"]", nil, htmlAttrs)
	}
	return AddressInput("address["+strconv.Itoa(numAddress)+"]", &resource.Address[numAddress], htmlAttrs)
}
func (resource *Patient) T_MaritalStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("maritalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("maritalStatus", resource.MaritalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Patient) T_MultipleBirthBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("multipleBirthBoolean", nil, htmlAttrs)
	}
	return BoolInput("multipleBirthBoolean", resource.MultipleBirthBoolean, htmlAttrs)
}
func (resource *Patient) T_MultipleBirthInteger(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("multipleBirthInteger", nil, htmlAttrs)
	}
	return IntInput("multipleBirthInteger", resource.MultipleBirthInteger, htmlAttrs)
}
func (resource *Patient) T_Photo(numPhoto int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPhoto >= len(resource.Photo) {
		return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", nil, htmlAttrs)
	}
	return AttachmentInput("photo["+strconv.Itoa(numPhoto)+"]", &resource.Photo[numPhoto], htmlAttrs)
}
func (resource *Patient) T_GeneralPractitioner(frs []FhirResource, numGeneralPractitioner int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGeneralPractitioner >= len(resource.GeneralPractitioner) {
		return ReferenceInput(frs, "generalPractitioner["+strconv.Itoa(numGeneralPractitioner)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "generalPractitioner["+strconv.Itoa(numGeneralPractitioner)+"]", &resource.GeneralPractitioner[numGeneralPractitioner], htmlAttrs)
}
func (resource *Patient) T_ManagingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization", resource.ManagingOrganization, htmlAttrs)
}
func (resource *Patient) T_ContactRelationship(numContact int, numRelationship int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) || numRelationship >= len(resource.Contact[numContact].Relationship) {
		return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].relationship["+strconv.Itoa(numRelationship)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].relationship["+strconv.Itoa(numRelationship)+"]", &resource.Contact[numContact].Relationship[numRelationship], optionsValueSet, htmlAttrs)
}
func (resource *Patient) T_ContactName(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", nil, htmlAttrs)
	}
	return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", resource.Contact[numContact].Name, htmlAttrs)
}
func (resource *Patient) T_ContactTelecom(numContact int, numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) || numTelecom >= len(resource.Contact[numContact].Telecom) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", &resource.Contact[numContact].Telecom[numTelecom], htmlAttrs)
}
func (resource *Patient) T_ContactAddress(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return AddressInput("contact["+strconv.Itoa(numContact)+"].address", nil, htmlAttrs)
	}
	return AddressInput("contact["+strconv.Itoa(numContact)+"].address", resource.Contact[numContact].Address, htmlAttrs)
}
func (resource *Patient) T_ContactGender(numContact int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil || numContact >= len(resource.Contact) {
		return CodeSelect("contact["+strconv.Itoa(numContact)+"].gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("contact["+strconv.Itoa(numContact)+"].gender", resource.Contact[numContact].Gender, optionsValueSet, htmlAttrs)
}
func (resource *Patient) T_ContactOrganization(frs []FhirResource, numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ReferenceInput(frs, "contact["+strconv.Itoa(numContact)+"].organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contact["+strconv.Itoa(numContact)+"].organization", resource.Contact[numContact].Organization, htmlAttrs)
}
func (resource *Patient) T_ContactPeriod(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return PeriodInput("contact["+strconv.Itoa(numContact)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("contact["+strconv.Itoa(numContact)+"].period", resource.Contact[numContact].Period, htmlAttrs)
}
func (resource *Patient) T_CommunicationPreferred(numCommunication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("communication["+strconv.Itoa(numCommunication)+"].preferred", resource.Communication[numCommunication].Preferred, htmlAttrs)
}
func (resource *Patient) T_LinkOther(frs []FhirResource, numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return ReferenceInput(frs, "link["+strconv.Itoa(numLink)+"].other", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "link["+strconv.Itoa(numLink)+"].other", &resource.Link[numLink].Other, htmlAttrs)
}
func (resource *Patient) T_LinkType(numLink int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSLink_type

	if resource == nil || numLink >= len(resource.Link) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].type", &resource.Link[numLink].Type, optionsValueSet, htmlAttrs)
}
