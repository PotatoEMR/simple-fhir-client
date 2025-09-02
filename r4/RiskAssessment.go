package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/RiskAssessment
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

// http://hl7.org/fhir/r4/StructureDefinition/RiskAssessment
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

func (resource *RiskAssessment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *RiskAssessment) T_Status() templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *RiskAssessment) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet)
}
func (resource *RiskAssessment) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *RiskAssessment) T_ReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *RiskAssessment) T_PredictionOutcome(numPrediction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Prediction) >= numPrediction {
		return CodeableConceptSelect("outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcome", resource.Prediction[numPrediction].Outcome, optionsValueSet)
}
func (resource *RiskAssessment) T_PredictionQualitativeRisk(numPrediction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Prediction) >= numPrediction {
		return CodeableConceptSelect("qualitativeRisk", nil, optionsValueSet)
	}
	return CodeableConceptSelect("qualitativeRisk", resource.Prediction[numPrediction].QualitativeRisk, optionsValueSet)
}
