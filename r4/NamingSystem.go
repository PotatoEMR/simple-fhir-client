//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `json:"id,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []Resource             `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              string                 `json:"name"`
	Status            string                 `json:"status"`
	Kind              string                 `json:"kind"`
	Date              string                 `json:"date"`
	Publisher         *string                `json:"publisher,omitempty"`
	Contact           []ContactDetail        `json:"contact,omitempty"`
	Responsible       *string                `json:"responsible,omitempty"`
	Type              *CodeableConcept       `json:"type,omitempty"`
	Description       *string                `json:"description,omitempty"`
	UseContext        []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `json:"jurisdiction,omitempty"`
	Usage             *string                `json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `json:"uniqueId"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NamingSystem
type NamingSystemUniqueId struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Value             string      `json:"value"`
	Preferred         *bool       `json:"preferred,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Period            *Period     `json:"period,omitempty"`
}

type OtherNamingSystem NamingSystem

// on convert struct to json, automatically add resourceType=NamingSystem
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}
