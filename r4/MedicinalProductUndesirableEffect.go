package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductUndesirableEffect) MedicinalProductUndesirableEffectLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) MedicinalProductUndesirableEffectSymptomConditionEffect(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("symptomConditionEffect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("symptomConditionEffect", resource.SymptomConditionEffect, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) MedicinalProductUndesirableEffectClassification(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classification", resource.Classification, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) MedicinalProductUndesirableEffectFrequencyOfOccurrence(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("frequencyOfOccurrence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("frequencyOfOccurrence", resource.FrequencyOfOccurrence, optionsValueSet)
}
