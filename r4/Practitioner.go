//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/Practitioner
type Practitioner struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Active            *bool                       `json:"active,omitempty"`
	Name              []HumanName                 `json:"name,omitempty"`
	Telecom           []ContactPoint              `json:"telecom,omitempty"`
	Address           []Address                   `json:"address,omitempty"`
	Gender            *string                     `json:"gender,omitempty"`
	BirthDate         *string                     `json:"birthDate,omitempty"`
	Photo             []Attachment                `json:"photo,omitempty"`
	Qualification     []PractitionerQualification `json:"qualification,omitempty"`
	Communication     []CodeableConcept           `json:"communication,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Practitioner
type PractitionerQualification struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Period            *Period         `json:"period,omitempty"`
	Issuer            *Reference      `json:"issuer,omitempty"`
}

type OtherPractitioner Practitioner

// on convert struct to json, automatically add resourceType=Practitioner
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}
