package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
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
	Date                *time.Time           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Reporter            *Reference           `json:"reporter,omitempty"`
	Period              Period               `json:"period"`
	ImprovementNotation *CodeableConcept     `json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup `json:"group,omitempty"`
	EvaluatedResource   []Reference          `json:"evaluatedResource,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
type MeasureReportGroup struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept               `json:"code,omitempty"`
	Population        []MeasureReportGroupPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                      `json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
type MeasureReportGroupPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
type MeasureReportGroupStratifier struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept                     `json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratum struct {
	Id                *string                                         `json:"id,omitempty"`
	Extension         []Extension                                     `json:"extension,omitempty"`
	ModifierExtension []Extension                                     `json:"modifierExtension,omitempty"`
	Value             *CodeableConcept                                `json:"value,omitempty"`
	Component         []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                                       `json:"measureScore,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
type MeasureReportGroupStratifierStratumComponent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Value             CodeableConcept `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MeasureReport
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
		return CodeSelect("MeasureReport.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MeasureReport.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSMeasure_report_type

	if resource == nil {
		return CodeSelect("MeasureReport.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MeasureReport.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_Measure(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("MeasureReport.Measure", nil, htmlAttrs)
	}
	return StringInput("MeasureReport.Measure", &resource.Measure, htmlAttrs)
}
func (resource *MeasureReport) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MeasureReport.Date", nil, htmlAttrs)
	}
	return DateTimeInput("MeasureReport.Date", resource.Date, htmlAttrs)
}
func (resource *MeasureReport) T_ImprovementNotation(htmlAttrs string) templ.Component {
	optionsValueSet := VSMeasure_improvement_notation

	if resource == nil {
		return CodeableConceptSelect("MeasureReport.ImprovementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.ImprovementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupPopulationCount(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Count", nil, htmlAttrs)
	}
	return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Population[numPopulation].Count, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierCode(numGroup int, numStratifier int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numCode >= len(resource.Group[numGroup].Stratifier[numStratifier].Code) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Group[numGroup].Stratifier[numStratifier].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValue(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Value", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Value", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Value, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValue(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Value", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Component["+strconv.Itoa(numComponent)+"].Value", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Value, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCode(numGroup int, numStratifier int, numStratum int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumPopulationCount(numGroup int, numStratifier int, numStratum int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numPopulation >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population) {
		return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", nil, htmlAttrs)
	}
	return IntInput("MeasureReport.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Stratum["+strconv.Itoa(numStratum)+"].Population["+strconv.Itoa(numPopulation)+"].Count", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Population[numPopulation].Count, htmlAttrs)
}
