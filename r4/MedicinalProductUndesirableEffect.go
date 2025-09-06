package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

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

func (resource *MedicinalProductUndesirableEffect) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductUndesirableEffect.Id", nil)
	}
	return StringInput("MedicinalProductUndesirableEffect.Id", resource.Id)
}
func (resource *MedicinalProductUndesirableEffect) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductUndesirableEffect.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductUndesirableEffect.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductUndesirableEffect) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductUndesirableEffect.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductUndesirableEffect.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) T_SymptomConditionEffect(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.SymptomConditionEffect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.SymptomConditionEffect", resource.SymptomConditionEffect, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) T_Classification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.Classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.Classification", resource.Classification, optionsValueSet)
}
func (resource *MedicinalProductUndesirableEffect) T_FrequencyOfOccurrence(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.FrequencyOfOccurrence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.FrequencyOfOccurrence", resource.FrequencyOfOccurrence, optionsValueSet)
}
