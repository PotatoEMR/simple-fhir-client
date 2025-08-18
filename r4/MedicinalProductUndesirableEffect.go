//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductUndesirableEffect
type MedicinalProductUndesirableEffect struct {
	Id                     *string          `json:"id,omitempty"`
	Meta                   *Meta            `json:"meta,omitempty"`
	ImplicitRules          *string          `json:"implicitRules,omitempty"`
	Language               *string          `json:"language,omitempty"`
	Text                   *Narrative       `json:"text,omitempty"`
	Contained              []Resource       `json:"contained,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	Subject                []Reference      `json:"subject,omitempty"`
	SymptomConditionEffect *CodeableConcept `json:"symptomConditionEffect,omitempty"`
	Classification         *CodeableConcept `json:"classification,omitempty"`
	FrequencyOfOccurrence  *CodeableConcept `json:"frequencyOfOccurrence,omitempty"`
	Population             []Population     `json:"population,omitempty"`
}

type OtherMedicinalProductUndesirableEffect MedicinalProductUndesirableEffect

// on convert struct to json, automatically add resourceType=MedicinalProductUndesirableEffect
func (r MedicinalProductUndesirableEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductUndesirableEffect
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductUndesirableEffect: OtherMedicinalProductUndesirableEffect(r),
		ResourceType:                           "MedicinalProductUndesirableEffect",
	})
}
