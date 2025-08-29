package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                  *string              `json:"id,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	ImplicitRules       *string              `json:"implicitRules,omitempty"`
	Language            *string              `json:"language,omitempty"`
	Text                *Narrative           `json:"text,omitempty"`
	Contained           []Resource           `json:"contained,omitempty"`
	Extension           []Extension          `json:"extension,omitempty"`
	ModifierExtension   []Extension          `json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `json:"identifier,omitempty"`
	Status              string               `json:"status"`
	Type                string               `json:"type"`
	Measure             string               `json:"measure"`
	Subject             *Reference           `json:"subject,omitempty"`
	Date                *string              `json:"date,omitempty"`
	Reporter            *Reference           `json:"reporter,omitempty"`
	Period              Period               `json:"period"`
	ImprovementNotation *CodeableConcept     `json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup `json:"group,omitempty"`
	EvaluatedResource   []Reference          `json:"evaluatedResource,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroup struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept               `json:"code,omitempty"`
	Population        []MeasureReportGroupPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                      `json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroupPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroupStratifier struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept                     `json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratum struct {
	Id                *string                                         `json:"id,omitempty"`
	Extension         []Extension                                     `json:"extension,omitempty"`
	ModifierExtension []Extension                                     `json:"modifierExtension,omitempty"`
	Value             *CodeableConcept                                `json:"value,omitempty"`
	Component         []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                                       `json:"measureScore,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratumComponent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Value             CodeableConcept `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
}

type OtherMeasureReport MeasureReport

// on convert struct to json, automatically add resourceType=MeasureReport
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

func (resource *MeasureReport) MeasureReportLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportStatus() templ.Component {
	optionsValueSet := VSMeasure_report_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportType() templ.Component {
	optionsValueSet := VSMeasure_report_type
	currentVal := ""
	if resource != nil {
		currentVal = resource.Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
