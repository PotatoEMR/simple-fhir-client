package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	MeasureScoreQuantity        *Quantity                      `json:"measureScoreQuantity"`
	MeasureScoreDateTime        *string                        `json:"measureScoreDateTime"`
	MeasureScoreCodeableConcept *CodeableConcept               `json:"measureScoreCodeableConcept"`
	MeasureScorePeriod          *Period                        `json:"measureScorePeriod"`
	MeasureScoreRange           *Range                         `json:"measureScoreRange"`
	MeasureScoreDuration        *Duration                      `json:"measureScoreDuration"`
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
	ValueCodeableConcept        *CodeableConcept                                `json:"valueCodeableConcept"`
	ValueBoolean                *bool                                           `json:"valueBoolean"`
	ValueQuantity               *Quantity                                       `json:"valueQuantity"`
	ValueRange                  *Range                                          `json:"valueRange"`
	ValueReference              *Reference                                      `json:"valueReference"`
	Component                   []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population                  []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScoreQuantity        *Quantity                                       `json:"measureScoreQuantity"`
	MeasureScoreDateTime        *string                                         `json:"measureScoreDateTime"`
	MeasureScoreCodeableConcept *CodeableConcept                                `json:"measureScoreCodeableConcept"`
	MeasureScorePeriod          *Period                                         `json:"measureScorePeriod"`
	MeasureScoreRange           *Range                                          `json:"measureScoreRange"`
	MeasureScoreDuration        *Duration                                       `json:"measureScoreDuration"`
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

func (resource *MeasureReport) MeasureReportLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportStatus() templ.Component {
	optionsValueSet := VSMeasure_report_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportType() templ.Component {
	optionsValueSet := VSMeasure_report_type

	if resource != nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportDataUpdateType() templ.Component {
	optionsValueSet := VSSubmit_data_update_type

	if resource != nil {
		return CodeSelect("dataUpdateType", nil, optionsValueSet)
	}
	return CodeSelect("dataUpdateType", resource.DataUpdateType, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportScoring(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scoring", resource.Scoring, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportImprovementNotation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("improvementNotation", resource.ImprovementNotation, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportGroupCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Code, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportGroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportGroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportGroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) >= numComponent {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet)
}
func (resource *MeasureReport) MeasureReportGroupStratifierStratumPopulationCode(numGroup int, numStratifier int, numStratum int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) >= numPopulation {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Code, optionsValueSet)
}
