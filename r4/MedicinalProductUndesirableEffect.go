package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

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

// struct -> json, automatically add resourceType=Patient
func (r MedicinalProductUndesirableEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductUndesirableEffect
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductUndesirableEffect: OtherMedicinalProductUndesirableEffect(r),
		ResourceType:                           "MedicinalProductUndesirableEffect",
	})
}

// json -> struct, first reject if resourceType != MedicinalProductUndesirableEffect
func (r *MedicinalProductUndesirableEffect) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MedicinalProductUndesirableEffect" {
		return errors.New("resourceType not MedicinalProductUndesirableEffect")
	}
	return json.Unmarshal(data, (*OtherMedicinalProductUndesirableEffect)(r))
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
func (resource *MedicinalProductUndesirableEffect) T_Subject(frs []FhirResource, numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_SymptomConditionEffect(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("symptomConditionEffect", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("symptomConditionEffect", resource.SymptomConditionEffect, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_Classification(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification", resource.Classification, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_FrequencyOfOccurrence(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("frequencyOfOccurrence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("frequencyOfOccurrence", resource.FrequencyOfOccurrence, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductUndesirableEffect) T_Population(numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPopulation >= len(resource.Population) {
		return PopulationInput("population["+strconv.Itoa(numPopulation)+"]", nil, htmlAttrs)
	}
	return PopulationInput("population["+strconv.Itoa(numPopulation)+"]", &resource.Population[numPopulation], htmlAttrs)
}
