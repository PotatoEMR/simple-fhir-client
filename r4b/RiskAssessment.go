//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                 *string                    `json:"id,omitempty"`
	Meta               *Meta                      `json:"meta,omitempty"`
	ImplicitRules      *string                    `json:"implicitRules,omitempty"`
	Language           *string                    `json:"language,omitempty"`
	Text               *Narrative                 `json:"text,omitempty"`
	Contained          []Resource                 `json:"contained,omitempty"`
	Extension          []Extension                `json:"extension,omitempty"`
	ModifierExtension  []Extension                `json:"modifierExtension,omitempty"`
	Identifier         []Identifier               `json:"identifier,omitempty"`
	BasedOn            *Reference                 `json:"basedOn,omitempty"`
	Parent             *Reference                 `json:"parent,omitempty"`
	Status             string                     `json:"status"`
	Method             *CodeableConcept           `json:"method,omitempty"`
	Code               *CodeableConcept           `json:"code,omitempty"`
	Subject            Reference                  `json:"subject"`
	Encounter          *Reference                 `json:"encounter,omitempty"`
	OccurrenceDateTime *string                    `json:"occurrenceDateTime"`
	OccurrencePeriod   *Period                    `json:"occurrencePeriod"`
	Condition          *Reference                 `json:"condition,omitempty"`
	Performer          *Reference                 `json:"performer,omitempty"`
	ReasonCode         []CodeableConcept          `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                `json:"reasonReference,omitempty"`
	Basis              []Reference                `json:"basis,omitempty"`
	Prediction         []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation         *string                    `json:"mitigation,omitempty"`
	Note               []Annotation               `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/RiskAssessment
type RiskAssessmentPrediction struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ProbabilityDecimal *float64         `json:"probabilityDecimal"`
	ProbabilityRange   *Range           `json:"probabilityRange"`
	QualitativeRisk    *CodeableConcept `json:"qualitativeRisk,omitempty"`
	RelativeRisk       *float64         `json:"relativeRisk,omitempty"`
	WhenPeriod         *Period          `json:"whenPeriod"`
	WhenRange          *Range           `json:"whenRange"`
	Rationale          *string          `json:"rationale,omitempty"`
}

type OtherRiskAssessment RiskAssessment

// on convert struct to json, automatically add resourceType=RiskAssessment
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
}
