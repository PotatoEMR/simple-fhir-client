package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/RiskAssessment
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
	OccurrenceDateTime *string                    `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                    `json:"occurrencePeriod,omitempty"`
	Condition          *Reference                 `json:"condition,omitempty"`
	Performer          *Reference                 `json:"performer,omitempty"`
	Reason             []CodeableReference        `json:"reason,omitempty"`
	Basis              []Reference                `json:"basis,omitempty"`
	Prediction         []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation         *string                    `json:"mitigation,omitempty"`
	Note               []Annotation               `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RiskAssessment
type RiskAssessmentPrediction struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ProbabilityDecimal *float64         `json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `json:"probabilityRange,omitempty"`
	QualitativeRisk    *CodeableConcept `json:"qualitativeRisk,omitempty"`
	RelativeRisk       *float64         `json:"relativeRisk,omitempty"`
	WhenPeriod         *Period          `json:"whenPeriod,omitempty"`
	WhenRange          *Range           `json:"whenRange,omitempty"`
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

func (resource *RiskAssessment) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RiskAssessment.Id", nil)
	}
	return StringInput("RiskAssessment.Id", resource.Id)
}
func (resource *RiskAssessment) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RiskAssessment.ImplicitRules", nil)
	}
	return StringInput("RiskAssessment.ImplicitRules", resource.ImplicitRules)
}
func (resource *RiskAssessment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RiskAssessment.Language", nil, optionsValueSet)
	}
	return CodeSelect("RiskAssessment.Language", resource.Language, optionsValueSet)
}
func (resource *RiskAssessment) T_Status() templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("RiskAssessment.Status", nil, optionsValueSet)
	}
	return CodeSelect("RiskAssessment.Status", &resource.Status, optionsValueSet)
}
func (resource *RiskAssessment) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskAssessment.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskAssessment.Method", resource.Method, optionsValueSet)
}
func (resource *RiskAssessment) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskAssessment.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskAssessment.Code", resource.Code, optionsValueSet)
}
func (resource *RiskAssessment) T_Mitigation() templ.Component {

	if resource == nil {
		return StringInput("RiskAssessment.Mitigation", nil)
	}
	return StringInput("RiskAssessment.Mitigation", resource.Mitigation)
}
func (resource *RiskAssessment) T_PredictionId(numPrediction int) templ.Component {

	if resource == nil || len(resource.Prediction) >= numPrediction {
		return StringInput("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Id", nil)
	}
	return StringInput("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Id", resource.Prediction[numPrediction].Id)
}
func (resource *RiskAssessment) T_PredictionOutcome(numPrediction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Prediction) >= numPrediction {
		return CodeableConceptSelect("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Outcome", resource.Prediction[numPrediction].Outcome, optionsValueSet)
}
func (resource *RiskAssessment) T_PredictionQualitativeRisk(numPrediction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Prediction) >= numPrediction {
		return CodeableConceptSelect("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].QualitativeRisk", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].QualitativeRisk", resource.Prediction[numPrediction].QualitativeRisk, optionsValueSet)
}
func (resource *RiskAssessment) T_PredictionRelativeRisk(numPrediction int) templ.Component {

	if resource == nil || len(resource.Prediction) >= numPrediction {
		return Float64Input("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].RelativeRisk", nil)
	}
	return Float64Input("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].RelativeRisk", resource.Prediction[numPrediction].RelativeRisk)
}
func (resource *RiskAssessment) T_PredictionRationale(numPrediction int) templ.Component {

	if resource == nil || len(resource.Prediction) >= numPrediction {
		return StringInput("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Rationale", nil)
	}
	return StringInput("RiskAssessment.Prediction["+strconv.Itoa(numPrediction)+"].Rationale", resource.Prediction[numPrediction].Rationale)
}
