package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *MeasureReport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MeasureReport.Id", nil)
	}
	return StringInput("MeasureReport.Id", resource.Id)
}
func (resource *MeasureReport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MeasureReport.ImplicitRules", nil)
	}
	return StringInput("MeasureReport.ImplicitRules", resource.ImplicitRules)
}
func (resource *MeasureReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MeasureReport.Language", nil, optionsValueSet)
	}
	return CodeSelect("MeasureReport.Language", resource.Language, optionsValueSet)
}
func (resource *MeasureReport) T_Status() templ.Component {
	optionsValueSet := VSMeasure_report_status

	if resource == nil {
		return CodeSelect("MeasureReport.Status", nil, optionsValueSet)
	}
	return CodeSelect("MeasureReport.Status", &resource.Status, optionsValueSet)
}
func (resource *MeasureReport) T_Type() templ.Component {
	optionsValueSet := VSMeasure_report_type

	if resource == nil {
		return CodeSelect("MeasureReport.Type", nil, optionsValueSet)
	}
	return CodeSelect("MeasureReport.Type", &resource.Type, optionsValueSet)
}
func (resource *MeasureReport) T_DataUpdateType() templ.Component {
	optionsValueSet := VSSubmit_data_update_type

	if resource == nil {
		return CodeSelect("MeasureReport.DataUpdateType", nil, optionsValueSet)
	}
	return CodeSelect("MeasureReport.DataUpdateType", resource.DataUpdateType, optionsValueSet)
}
func (resource *MeasureReport) T_Measure() templ.Component {

	if resource == nil {
		return StringInput("MeasureReport.Measure", nil)
	}
	return StringInput("MeasureReport.Measure", resource.Measure)
}
func (resource *MeasureReport) T_Date() templ.Component {

	if resource == nil {
		return StringInput("MeasureReport.Date", nil)
	}
	return StringInput("MeasureReport.Date", resource.Date)
}
func (resource *MeasureReport) T_Scoring(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MeasureReport.Scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Scoring", resource.Scoring, optionsValueSet)
}
func (resource *MeasureReport) T_ImprovementNotation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MeasureReport.ImprovementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.ImprovementNotation", resource.ImprovementNotation, optionsValueSet)
}
func (resource *MeasureReport) T_GroupId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Id", resource.Group[numGroup].Id)
}
func (resource *MeasureReport) T_GroupLinkId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].LinkId", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].LinkId", resource.Group[numGroup].LinkId)
}
func (resource *MeasureReport) T_GroupCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Code", resource.Group[numGroup].Code, optionsValueSet)
}
func (resource *MeasureReport) T_GroupPopulationId(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Id", resource.Group[numGroup].Population[numPopulation].Id)
}
func (resource *MeasureReport) T_GroupPopulationLinkId(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", resource.Group[numGroup].Population[numPopulation].LinkId)
}
func (resource *MeasureReport) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet)
}
func (resource *MeasureReport) T_GroupPopulationCount(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Count", nil)
	}
	return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Population[numPopulation].Count)
}
func (resource *MeasureReport) T_GroupStratifierId(numGroup int, numStratifier int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Id)
}
func (resource *MeasureReport) T_GroupStratifierLinkId(numGroup int, numStratifier int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId)
}
func (resource *MeasureReport) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet)
}
func (resource *MeasureReport) T_GroupStratifierStratumId(numGroup int, numStratifier int, numStratum int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Id)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentId(numGroup int, numStratifier int, numStratum int, numComponent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) >= numComponent {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Id)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentLinkId(numGroup int, numStratifier int, numStratum int, numComponent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) >= numComponent {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].LinkId)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) >= numComponent {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationId(numGroup int, numStratifier int, numStratum int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) >= numPopulation {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Id", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Id)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationLinkId(numGroup int, numStratifier int, numStratum int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) >= numPopulation {
		return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", nil)
	}
	return StringInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].LinkId)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCode(numGroup int, numStratifier int, numStratum int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) >= numPopulation {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Code, optionsValueSet)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCount(numGroup int, numStratifier int, numStratum int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) >= numStratum || len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) >= numPopulation {
		return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", nil)
	}
	return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Count)
}
