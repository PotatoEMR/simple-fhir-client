//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/RelatedPerson
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

// http://hl7.org/fhir/r5/StructureDefinition/RelatedPerson
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
