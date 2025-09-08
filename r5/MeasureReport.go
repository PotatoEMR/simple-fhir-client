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
	Date                *time.Time           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	MeasureScoreDateTime        *time.Time                     `json:"measureScoreDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	MeasureScoreDateTime        *time.Time                                      `json:"measureScoreDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *MeasureReport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMeasure_report_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSMeasure_report_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_DataUpdateType(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubmit_data_update_type

	if resource == nil {
		return CodeSelect("DataUpdateType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DataUpdateType", resource.DataUpdateType, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Measure(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Measure", nil, htmlAttrs)
	}
	return StringInput("Measure", resource.Measure, htmlAttrs)
}
func (resource *MeasureReport) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *MeasureReport) T_Scoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Scoring", resource.Scoring, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_ImprovementNotation(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ImprovementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImprovementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupLinkId(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]LinkId", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]LinkId", resource.Group[numGroup].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreDateTime(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return DateTimeInput("Group["+strconv.Itoa(numGroup)+"]MeasureScoreDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Group["+strconv.Itoa(numGroup)+"]MeasureScoreDateTime", resource.Group[numGroup].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScoreCodeableConcept(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]MeasureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]MeasureScoreCodeableConcept", resource.Group[numGroup].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationLinkId(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].LinkId", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].LinkId", resource.Group[numGroup].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCount(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return IntInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Count", nil, htmlAttrs)
	}
	return IntInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Population[numPopulation].Count, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierLinkId(numGroup int, numStratifier int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueCodeableConcept(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].ValueCodeableConcept", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValueBoolean(numGroup int, numStratifier int, numStratum int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return BoolInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].ValueBoolean", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].ValueBoolean, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreDateTime(numGroup int, numStratifier int, numStratum int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return DateTimeInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].MeasureScoreDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].MeasureScoreDateTime", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreDateTime, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScoreCodeableConcept(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].MeasureScoreCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].MeasureScoreCodeableConcept", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScoreCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentLinkId(numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueCodeableConcept(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].ValueCodeableConcept", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValueBoolean(numGroup int, numStratifier int, numStratum int, numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return BoolInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].ValueBoolean", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].ValueBoolean, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationLinkId(numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCode(numGroup int, numStratifier int, numStratum int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCount(numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return IntInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", nil, htmlAttrs)
	}
	return IntInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Count, htmlAttrs)
}
