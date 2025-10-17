package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
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
	DataUpdateType      *string              `json:"dataUpdateType,omitempty"`
	Measure             *string              `json:"measure,omitempty"`
	Subject             *Reference           `json:"subject,omitempty"`
	Date                *FhirDateTime        `json:"date,omitempty"`
	Reporter            *Reference           `json:"reporter,omitempty"`
	ReportingVendor     *Reference           `json:"reportingVendor,omitempty"`
	Location            *Reference           `json:"location,omitempty"`
	Period              Period               `json:"period"`
	InputParameters     *Reference           `json:"inputParameters,omitempty"`
	Scoring             *CodeableConcept     `json:"scoring,omitempty"`
	ImprovementNotation *CodeableConcept     `json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup `json:"group,omitempty"`
	SupplementalData    []Reference          `json:"supplementalData,omitempty"`
	EvaluatedResource   []Reference          `json:"evaluatedResource,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroup struct {
	Id                          *string                        `json:"id,omitempty"`
	Extension                   []Extension                    `json:"extension,omitempty"`
	ModifierExtension           []Extension                    `json:"modifierExtension,omitempty"`
	LinkId                      *string                        `json:"linkId,omitempty"`
	Code                        *CodeableConcept               `json:"code,omitempty"`
	Subject                     *Reference                     `json:"subject,omitempty"`
	Population                  []MeasureReportGroupPopulation `json:"population,omitempty"`
	MeasureScoreQuantity        *Quantity                      `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *FhirDateTime                  `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept               `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period                        `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                         `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration                      `json:"measureScoreDuration,omitempty"`
	Stratifier                  []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroupPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	LinkId            *string          `json:"linkId,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
	SubjectReport     []Reference      `json:"subjectReport,omitempty"`
	Subjects          *Reference       `json:"subjects,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroupStratifier struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	LinkId            *string                               `json:"linkId,omitempty"`
	Code              *CodeableConcept                      `json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratum struct {
	Id                          *string                                         `json:"id,omitempty"`
	Extension                   []Extension                                     `json:"extension,omitempty"`
	ModifierExtension           []Extension                                     `json:"modifierExtension,omitempty"`
	ValueCodeableConcept        *CodeableConcept                                `json:"valueCodeableConcept,omitempty"`
	ValueBoolean                *bool                                           `json:"valueBoolean,omitempty"`
	ValueQuantity               *Quantity                                       `json:"valueQuantity,omitempty"`
	ValueRange                  *Range                                          `json:"valueRange,omitempty"`
	ValueReference              *Reference                                      `json:"valueReference,omitempty"`
	Component                   []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population                  []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScoreQuantity        *Quantity                                       `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *FhirDateTime                                   `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept                                `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period                                         `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                                          `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration                                       `json:"measureScoreDuration,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratumComponent struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	LinkId               *string         `json:"linkId,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	LinkId            *string          `json:"linkId,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
	SubjectReport     []Reference      `json:"subjectReport,omitempty"`
	Subjects          *Reference       `json:"subjects,omitempty"`
}

type OtherMeasureReport MeasureReport

// struct -> json, automatically add resourceType=Patient
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

func (r MeasureReport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MeasureReport/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MeasureReport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r MeasureReport) ResourceType() string {
	return "MeasureReport"
}

func (resource *MeasureReport) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMeasure_report_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMeasure_report_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_DataUpdateType(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubmit_data_update_type

	if resource == nil {
		return CodeSelect("dataUpdateType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("dataUpdateType", resource.DataUpdateType, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Measure(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("measure", nil, htmlAttrs)
	}
	return StringInput("measure", resource.Measure, htmlAttrs)
}
func (resource *MeasureReport) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *MeasureReport) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *MeasureReport) T_Reporter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "reporter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reporter", resource.Reporter, htmlAttrs)
}
func (resource *MeasureReport) T_ReportingVendor(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "reportingVendor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reportingVendor", resource.ReportingVendor, htmlAttrs)
}
func (resource *MeasureReport) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *MeasureReport) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", &resource.Period, htmlAttrs)
}
func (resource *MeasureReport) T_InputParameters(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "inputParameters", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "inputParameters", resource.InputParameters, htmlAttrs)
}
func (resource *MeasureReport) T_Scoring(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("scoring", resource.Scoring, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_ImprovementNotation(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("improvementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_SupplementalData(frs []FhirResource, numSupplementalData int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return ReferenceInput(frs, "supplementalData["+strconv.Itoa(numSupplementalData)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supplementalData["+strconv.Itoa(numSupplementalData)+"]", &resource.SupplementalData[numSupplementalData], htmlAttrs)
}
func (resource *MeasureReport) T_EvaluatedResource(frs []FhirResource, numEvaluatedResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvaluatedResource >= len(resource.EvaluatedResource) {
		return ReferenceInput(frs, "evaluatedResource["+strconv.Itoa(numEvaluatedResource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "evaluatedResource["+strconv.Itoa(numEvaluatedResource)+"]", &resource.EvaluatedResource[numEvaluatedResource], htmlAttrs)
}
func (resource *MeasureReport) T_GroupLinkId(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].linkId", resource.Group[numGroup].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupSubject(frs []FhirResource, numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].subject", resource.Group[numGroup].Subject, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreQuantity(numGroup int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].measureScoreQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].measureScoreQuantity", resource.Group[numGroup].MeasureScoreQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreDateTime(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].measureScoreDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].measureScoreDateTime", resource.Group[numGroup].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreCodeableConcept(numGroup int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].measureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].measureScoreCodeableConcept", resource.Group[numGroup].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScorePeriod(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return PeriodInput("group["+strconv.Itoa(numGroup)+"].measureScorePeriod", nil, htmlAttrs)
	}
	return PeriodInput("group["+strconv.Itoa(numGroup)+"].measureScorePeriod", resource.Group[numGroup].MeasureScorePeriod, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreRange(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return RangeInput("group["+strconv.Itoa(numGroup)+"].measureScoreRange", nil, htmlAttrs)
	}
	return RangeInput("group["+strconv.Itoa(numGroup)+"].measureScoreRange", resource.Group[numGroup].MeasureScoreRange, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreDuration(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return DurationInput("group["+strconv.Itoa(numGroup)+"].measureScoreDuration", nil, htmlAttrs)
	}
	return DurationInput("group["+strconv.Itoa(numGroup)+"].measureScoreDuration", resource.Group[numGroup].MeasureScoreDuration, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationLinkId(numGroup int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].linkId", resource.Group[numGroup].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCount(numGroup int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].count", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].count", resource.Group[numGroup].Population[numPopulation].Count, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationSubjectResults(frs []FhirResource, numGroup int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjectResults", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjectResults", resource.Group[numGroup].Population[numPopulation].SubjectResults, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationSubjectReport(frs []FhirResource, numGroup int, numPopulation int, numSubjectReport int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) || numSubjectReport >= len(resource.Group[numGroup].Population[numPopulation].SubjectReport) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjectReport["+strconv.Itoa(numSubjectReport)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjectReport["+strconv.Itoa(numSubjectReport)+"]", &resource.Group[numGroup].Population[numPopulation].SubjectReport[numSubjectReport], htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationSubjects(frs []FhirResource, numGroup int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjects", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].subjects", resource.Group[numGroup].Population[numPopulation].Subjects, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierLinkId(numGroup int, numStratifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].linkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueCodeableConcept(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueCodeableConcept", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueBoolean(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueBoolean", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueBoolean, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueQuantity(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueQuantity", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueRange(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueRange", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueRange, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueReference(frs []FhirResource, numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].valueReference", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueReference, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreQuantity(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreQuantity", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreDateTime(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDateTime", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreCodeableConcept(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreCodeableConcept", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScorePeriod(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return PeriodInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScorePeriod", nil, htmlAttrs)
	}
	return PeriodInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScorePeriod", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScorePeriod, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreRange(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreRange", nil, htmlAttrs)
	}
	return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreRange", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreRange, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreDuration(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return DurationInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDuration", nil, htmlAttrs)
	}
	return DurationInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDuration", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreDuration, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentLinkId(numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].linkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueCodeableConcept(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueCodeableConcept", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueBoolean(numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueBoolean", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueBoolean, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueQuantity(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueQuantity", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueRange(numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueRange", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueRange, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueReference(frs []FhirResource, numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].valueReference", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueReference, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationLinkId(numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].linkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCode(numGroup int, numStratifier int, numStratum int, numPopulation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].code", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCount(numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].count", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].count", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Count, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationSubjectResults(frs []FhirResource, numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjectResults", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjectResults", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].SubjectResults, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationSubjectReport(frs []FhirResource, numGroup int, numStratifier int, numStratum int, numPopulation int, numSubjectReport int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) || numSubjectReport >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].SubjectReport) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjectReport["+strconv.Itoa(numSubjectReport)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjectReport["+strconv.Itoa(numSubjectReport)+"]", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].SubjectReport[numSubjectReport], htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationSubjects(frs []FhirResource, numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjects", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].population["+strconv.Itoa(numPopulation)+"].subjects", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Subjects, htmlAttrs)
}
