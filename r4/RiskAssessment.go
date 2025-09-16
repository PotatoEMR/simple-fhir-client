package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	OccurrenceDateTime *FhirDateTime              `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                    `json:"occurrencePeriod,omitempty"`
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
func (r RiskAssessment) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RiskAssessment/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RiskAssessment"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *RiskAssessment) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_Method(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *RiskAssessment) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_Mitigation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("mitigation", nil, htmlAttrs)
	}
	return StringInput("mitigation", resource.Mitigation, htmlAttrs)
}
func (resource *RiskAssessment) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionOutcome(numPrediction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrediction >= len(resource.Prediction) {
		return CodeableConceptSelect("prediction["+strconv.Itoa(numPrediction)+"].outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("prediction["+strconv.Itoa(numPrediction)+"].outcome", resource.Prediction[numPrediction].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionProbabilityDecimal(numPrediction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrediction >= len(resource.Prediction) {
		return Float64Input("prediction["+strconv.Itoa(numPrediction)+"].probabilityDecimal", nil, htmlAttrs)
	}
	return Float64Input("prediction["+strconv.Itoa(numPrediction)+"].probabilityDecimal", resource.Prediction[numPrediction].ProbabilityDecimal, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionQualitativeRisk(numPrediction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrediction >= len(resource.Prediction) {
		return CodeableConceptSelect("prediction["+strconv.Itoa(numPrediction)+"].qualitativeRisk", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("prediction["+strconv.Itoa(numPrediction)+"].qualitativeRisk", resource.Prediction[numPrediction].QualitativeRisk, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionRelativeRisk(numPrediction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrediction >= len(resource.Prediction) {
		return Float64Input("prediction["+strconv.Itoa(numPrediction)+"].relativeRisk", nil, htmlAttrs)
	}
	return Float64Input("prediction["+strconv.Itoa(numPrediction)+"].relativeRisk", resource.Prediction[numPrediction].RelativeRisk, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionRationale(numPrediction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrediction >= len(resource.Prediction) {
		return StringInput("prediction["+strconv.Itoa(numPrediction)+"].rationale", nil, htmlAttrs)
	}
	return StringInput("prediction["+strconv.Itoa(numPrediction)+"].rationale", resource.Prediction[numPrediction].Rationale, htmlAttrs)
}
