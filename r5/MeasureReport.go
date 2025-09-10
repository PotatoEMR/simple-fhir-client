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
	Date                *string              `json:"date,omitempty"`
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
	MeasureScoreDateTime        *string                        `json:"measureScoreDateTime,omitempty"`
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
	MeasureScoreDateTime        *string                                         `json:"measureScoreDateTime,omitempty"`
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
func (resource *MeasureReport) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
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
func (resource *MeasureReport) T_GroupMeasureScoreDateTime(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return DateTimeInput("group["+strconv.Itoa(numGroup)+"].measureScoreDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("group["+strconv.Itoa(numGroup)+"].measureScoreDateTime", resource.Group[numGroup].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreCodeableConcept(numGroup int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].measureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].measureScoreCodeableConcept", resource.Group[numGroup].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
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
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreDateTime(numGroup int, numStratifier int, numStratum int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return DateTimeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreDateTime", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreCodeableConcept(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScoreCodeableConcept", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
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
