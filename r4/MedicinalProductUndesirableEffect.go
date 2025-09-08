package r4

//generated with command go run ./bultaoreune
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
func (r MedicinalProductUndesirableEffect) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductUndesirableEffect/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MedicinalProductUndesirableEffect"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductUndesirableEffect) T_SymptomConditionEffect(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.SymptomConditionEffect", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.SymptomConditionEffect", resource.SymptomConditionEffect, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_Classification(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.Classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.Classification", resource.Classification, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_FrequencyOfOccurrence(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductUndesirableEffect.FrequencyOfOccurrence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductUndesirableEffect.FrequencyOfOccurrence", resource.FrequencyOfOccurrence, optionsValueSet, htmlAttrs)
}
