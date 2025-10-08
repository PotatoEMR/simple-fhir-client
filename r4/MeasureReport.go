package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                *FhirDateTime        `json:"date,omitempty"`
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
func (resource *MeasureReport) T_Measure(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("measure", nil, htmlAttrs)
	}
	return StringInput("measure", &resource.Measure, htmlAttrs)
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
func (resource *MeasureReport) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", &resource.Period, htmlAttrs)
}
func (resource *MeasureReport) T_ImprovementNotation(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMeasure_improvement_notation

	if resource == nil {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("improvementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_EvaluatedResource(frs []FhirResource, numEvaluatedResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvaluatedResource >= len(resource.EvaluatedResource) {
		return ReferenceInput(frs, "evaluatedResource["+strconv.Itoa(numEvaluatedResource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "evaluatedResource["+strconv.Itoa(numEvaluatedResource)+"]", &resource.EvaluatedResource[numEvaluatedResource], htmlAttrs)
}
func (resource *MeasureReport) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupMeasureScore(numGroup int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].measureScore", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].measureScore", resource.Group[numGroup].MeasureScore, optionsValueSet, htmlAttrs)
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
func (resource *MeasureReport) T_GroupStratifierCode(numGroup int, numStratifier int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numCode >= len(resource.Group[numGroup].Stratifier[numStratifier].Code) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code["+strconv.Itoa(numCode)+"]", &resource.Group[numGroup].Stratifier[numStratifier].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumValue(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].value", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].value", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Value, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumMeasureScore(numGroup int, numStratifier int, numStratum int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScore", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].measureScore", resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].MeasureScore, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentCode(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].code", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *MeasureReport) T_GroupStratifierStratumComponentValue(numGroup int, numStratifier int, numStratum int, numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numStratum >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].value", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].stratum["+strconv.Itoa(numStratum)+"].component["+strconv.Itoa(numComponent)+"].value", &resource.Group[numGroup].Stratifier[numStratifier].Stratum[numStratum].Component[numComponent].Value, optionsValueSet, htmlAttrs)
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
