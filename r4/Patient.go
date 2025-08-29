package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"
import "strings"

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
	BirthDate            *string                `json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                  `json:"deceasedBoolean"`
	DeceasedDateTime     *string                `json:"deceasedDateTime"`
	Address              []Address              `json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                  `json:"multipleBirthBoolean"`
	MultipleBirthInteger *int                   `json:"multipleBirthInteger"`
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

// on convert struct to json, automatically add resourceType=Patient
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

func (resource *Patient) PatientLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Patient) PatientGender() templ.Component {
	optionsValueSet := VSAdministrative_gender
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Gender
	}
	return CodeSelect("gender", currentVal, optionsValueSet)
}
func (resource *Patient) PatientContactGender(numContact int) templ.Component {
	optionsValueSet := VSAdministrative_gender
	currentVal := ""
	if resource != nil && len(resource.Contact) >= numContact {
		currentVal = *resource.Contact[numContact].Gender
	}
	return CodeSelect("gender", currentVal, optionsValueSet)
}
func (resource *Patient) PatientLinkType(numLink int) templ.Component {
	optionsValueSet := VSLink_type
	currentVal := ""
	if resource != nil && len(resource.Link) >= numLink {
		currentVal = resource.Link[numLink].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
