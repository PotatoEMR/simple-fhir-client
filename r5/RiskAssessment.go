package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	OccurrenceDateTime *time.Time                 `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *RiskAssessment) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("RiskAssessment.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RiskAssessment.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_Method(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskAssessment.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskAssessment.Method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskAssessment.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskAssessment.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_OccurrenceDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("RiskAssessment.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("RiskAssessment.OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *RiskAssessment) T_Mitigation(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("RiskAssessment.Mitigation", nil, htmlAttrs)
	}
	return StringInput("RiskAssessment.Mitigation", resource.Mitigation, htmlAttrs)
}
func (resource *RiskAssessment) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("RiskAssessment.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("RiskAssessment.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionOutcome(numPrediction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrediction >= len(resource.Prediction) {
		return CodeableConceptSelect("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..Outcome", resource.Prediction[numPrediction].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionProbabilityDecimal(numPrediction int, htmlAttrs string) templ.Component {

	if resource == nil || numPrediction >= len(resource.Prediction) {
		return Float64Input("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..ProbabilityDecimal", nil, htmlAttrs)
	}
	return Float64Input("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..ProbabilityDecimal", resource.Prediction[numPrediction].ProbabilityDecimal, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionQualitativeRisk(numPrediction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrediction >= len(resource.Prediction) {
		return CodeableConceptSelect("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..QualitativeRisk", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..QualitativeRisk", resource.Prediction[numPrediction].QualitativeRisk, optionsValueSet, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionRelativeRisk(numPrediction int, htmlAttrs string) templ.Component {

	if resource == nil || numPrediction >= len(resource.Prediction) {
		return Float64Input("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..RelativeRisk", nil, htmlAttrs)
	}
	return Float64Input("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..RelativeRisk", resource.Prediction[numPrediction].RelativeRisk, htmlAttrs)
}
func (resource *RiskAssessment) T_PredictionRationale(numPrediction int, htmlAttrs string) templ.Component {

	if resource == nil || numPrediction >= len(resource.Prediction) {
		return StringInput("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..Rationale", nil, htmlAttrs)
	}
	return StringInput("RiskAssessment.Prediction."+strconv.Itoa(numPrediction)+"..Rationale", resource.Prediction[numPrediction].Rationale, htmlAttrs)
}
