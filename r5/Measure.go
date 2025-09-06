package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type Measure struct {
	Id                              *string                   `json:"id,omitempty"`
	Meta                            *Meta                     `json:"meta,omitempty"`
	ImplicitRules                   *string                   `json:"implicitRules,omitempty"`
	Language                        *string                   `json:"language,omitempty"`
	Text                            *Narrative                `json:"text,omitempty"`
	Contained                       []Resource                `json:"contained,omitempty"`
	Extension                       []Extension               `json:"extension,omitempty"`
	ModifierExtension               []Extension               `json:"modifierExtension,omitempty"`
	Url                             *string                   `json:"url,omitempty"`
	Identifier                      []Identifier              `json:"identifier,omitempty"`
	Version                         *string                   `json:"version,omitempty"`
	VersionAlgorithmString          *string                   `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding          *Coding                   `json:"versionAlgorithmCoding,omitempty"`
	Name                            *string                   `json:"name,omitempty"`
	Title                           *string                   `json:"title,omitempty"`
	Subtitle                        *string                   `json:"subtitle,omitempty"`
	Status                          string                    `json:"status"`
	Experimental                    *bool                     `json:"experimental,omitempty"`
	SubjectCodeableConcept          *CodeableConcept          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference                *Reference                `json:"subjectReference,omitempty"`
	Basis                           *string                   `json:"basis,omitempty"`
	Date                            *string                   `json:"date,omitempty"`
	Publisher                       *string                   `json:"publisher,omitempty"`
	Contact                         []ContactDetail           `json:"contact,omitempty"`
	Description                     *string                   `json:"description,omitempty"`
	UseContext                      []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                         *string                   `json:"purpose,omitempty"`
	Usage                           *string                   `json:"usage,omitempty"`
	Copyright                       *string                   `json:"copyright,omitempty"`
	CopyrightLabel                  *string                   `json:"copyrightLabel,omitempty"`
	ApprovalDate                    *string                   `json:"approvalDate,omitempty"`
	LastReviewDate                  *string                   `json:"lastReviewDate,omitempty"`
	EffectivePeriod                 *Period                   `json:"effectivePeriod,omitempty"`
	Topic                           []CodeableConcept         `json:"topic,omitempty"`
	Author                          []ContactDetail           `json:"author,omitempty"`
	Editor                          []ContactDetail           `json:"editor,omitempty"`
	Reviewer                        []ContactDetail           `json:"reviewer,omitempty"`
	Endorser                        []ContactDetail           `json:"endorser,omitempty"`
	RelatedArtifact                 []RelatedArtifact         `json:"relatedArtifact,omitempty"`
	Library                         []string                  `json:"library,omitempty"`
	Disclaimer                      *string                   `json:"disclaimer,omitempty"`
	Scoring                         *CodeableConcept          `json:"scoring,omitempty"`
	ScoringUnit                     *CodeableConcept          `json:"scoringUnit,omitempty"`
	CompositeScoring                *CodeableConcept          `json:"compositeScoring,omitempty"`
	Type                            []CodeableConcept         `json:"type,omitempty"`
	RiskAdjustment                  *string                   `json:"riskAdjustment,omitempty"`
	RateAggregation                 *string                   `json:"rateAggregation,omitempty"`
	Rationale                       *string                   `json:"rationale,omitempty"`
	ClinicalRecommendationStatement *string                   `json:"clinicalRecommendationStatement,omitempty"`
	ImprovementNotation             *CodeableConcept          `json:"improvementNotation,omitempty"`
	Term                            []MeasureTerm             `json:"term,omitempty"`
	Guidance                        *string                   `json:"guidance,omitempty"`
	Group                           []MeasureGroup            `json:"group,omitempty"`
	SupplementalData                []MeasureSupplementalData `json:"supplementalData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureTerm struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Definition        *string          `json:"definition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureGroup struct {
	Id                     *string                  `json:"id,omitempty"`
	Extension              []Extension              `json:"extension,omitempty"`
	ModifierExtension      []Extension              `json:"modifierExtension,omitempty"`
	LinkId                 *string                  `json:"linkId,omitempty"`
	Code                   *CodeableConcept         `json:"code,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	Type                   []CodeableConcept        `json:"type,omitempty"`
	SubjectCodeableConcept *CodeableConcept         `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference               `json:"subjectReference,omitempty"`
	Basis                  *string                  `json:"basis,omitempty"`
	Scoring                *CodeableConcept         `json:"scoring,omitempty"`
	ScoringUnit            *CodeableConcept         `json:"scoringUnit,omitempty"`
	RateAggregation        *string                  `json:"rateAggregation,omitempty"`
	ImprovementNotation    *CodeableConcept         `json:"improvementNotation,omitempty"`
	Library                []string                 `json:"library,omitempty"`
	Population             []MeasureGroupPopulation `json:"population,omitempty"`
	Stratifier             []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureGroupPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	LinkId            *string          `json:"linkId,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          *Expression      `json:"criteria,omitempty"`
	GroupDefinition   *Reference       `json:"groupDefinition,omitempty"`
	InputPopulationId *string          `json:"inputPopulationId,omitempty"`
	AggregateMethod   *CodeableConcept `json:"aggregateMethod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureGroupStratifier struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	LinkId            *string                           `json:"linkId,omitempty"`
	Code              *CodeableConcept                  `json:"code,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	Criteria          *Expression                       `json:"criteria,omitempty"`
	GroupDefinition   *Reference                        `json:"groupDefinition,omitempty"`
	Component         []MeasureGroupStratifierComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureGroupStratifierComponent struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	LinkId            *string          `json:"linkId,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          *Expression      `json:"criteria,omitempty"`
	GroupDefinition   *Reference       `json:"groupDefinition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Measure
type MeasureSupplementalData struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	LinkId            *string           `json:"linkId,omitempty"`
	Code              *CodeableConcept  `json:"code,omitempty"`
	Usage             []CodeableConcept `json:"usage,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Criteria          Expression        `json:"criteria"`
}

type OtherMeasure Measure

// on convert struct to json, automatically add resourceType=Measure
func (r Measure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasure
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasure: OtherMeasure(r),
		ResourceType: "Measure",
	})
}

func (resource *Measure) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Measure.Id", nil)
	}
	return StringInput("Measure.Id", resource.Id)
}
func (resource *Measure) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Measure.ImplicitRules", nil)
	}
	return StringInput("Measure.ImplicitRules", resource.ImplicitRules)
}
func (resource *Measure) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Measure.Language", nil, optionsValueSet)
	}
	return CodeSelect("Measure.Language", resource.Language, optionsValueSet)
}
func (resource *Measure) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Measure.Url", nil)
	}
	return StringInput("Measure.Url", resource.Url)
}
func (resource *Measure) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Measure.Version", nil)
	}
	return StringInput("Measure.Version", resource.Version)
}
func (resource *Measure) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Measure.Name", nil)
	}
	return StringInput("Measure.Name", resource.Name)
}
func (resource *Measure) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Measure.Title", nil)
	}
	return StringInput("Measure.Title", resource.Title)
}
func (resource *Measure) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("Measure.Subtitle", nil)
	}
	return StringInput("Measure.Subtitle", resource.Subtitle)
}
func (resource *Measure) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Measure.Status", nil, optionsValueSet)
	}
	return CodeSelect("Measure.Status", &resource.Status, optionsValueSet)
}
func (resource *Measure) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Measure.Experimental", nil)
	}
	return BoolInput("Measure.Experimental", resource.Experimental)
}
func (resource *Measure) T_Basis() templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil {
		return CodeSelect("Measure.Basis", nil, optionsValueSet)
	}
	return CodeSelect("Measure.Basis", resource.Basis, optionsValueSet)
}
func (resource *Measure) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Measure.Date", nil)
	}
	return StringInput("Measure.Date", resource.Date)
}
func (resource *Measure) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Measure.Publisher", nil)
	}
	return StringInput("Measure.Publisher", resource.Publisher)
}
func (resource *Measure) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Measure.Description", nil)
	}
	return StringInput("Measure.Description", resource.Description)
}
func (resource *Measure) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Measure.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Measure) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Measure.Purpose", nil)
	}
	return StringInput("Measure.Purpose", resource.Purpose)
}
func (resource *Measure) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("Measure.Usage", nil)
	}
	return StringInput("Measure.Usage", resource.Usage)
}
func (resource *Measure) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Measure.Copyright", nil)
	}
	return StringInput("Measure.Copyright", resource.Copyright)
}
func (resource *Measure) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("Measure.CopyrightLabel", nil)
	}
	return StringInput("Measure.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *Measure) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Measure.ApprovalDate", nil)
	}
	return StringInput("Measure.ApprovalDate", resource.ApprovalDate)
}
func (resource *Measure) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Measure.LastReviewDate", nil)
	}
	return StringInput("Measure.LastReviewDate", resource.LastReviewDate)
}
func (resource *Measure) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("Measure.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *Measure) T_Library(numLibrary int) templ.Component {

	if resource == nil || len(resource.Library) >= numLibrary {
		return StringInput("Measure.Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("Measure.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary])
}
func (resource *Measure) T_Disclaimer() templ.Component {

	if resource == nil {
		return StringInput("Measure.Disclaimer", nil)
	}
	return StringInput("Measure.Disclaimer", resource.Disclaimer)
}
func (resource *Measure) T_Scoring(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.Scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Scoring", resource.Scoring, optionsValueSet)
}
func (resource *Measure) T_ScoringUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.ScoringUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.ScoringUnit", resource.ScoringUnit, optionsValueSet)
}
func (resource *Measure) T_CompositeScoring(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.CompositeScoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.CompositeScoring", resource.CompositeScoring, optionsValueSet)
}
func (resource *Measure) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Measure.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Measure) T_RiskAdjustment() templ.Component {

	if resource == nil {
		return StringInput("Measure.RiskAdjustment", nil)
	}
	return StringInput("Measure.RiskAdjustment", resource.RiskAdjustment)
}
func (resource *Measure) T_RateAggregation() templ.Component {

	if resource == nil {
		return StringInput("Measure.RateAggregation", nil)
	}
	return StringInput("Measure.RateAggregation", resource.RateAggregation)
}
func (resource *Measure) T_Rationale() templ.Component {

	if resource == nil {
		return StringInput("Measure.Rationale", nil)
	}
	return StringInput("Measure.Rationale", resource.Rationale)
}
func (resource *Measure) T_ClinicalRecommendationStatement() templ.Component {

	if resource == nil {
		return StringInput("Measure.ClinicalRecommendationStatement", nil)
	}
	return StringInput("Measure.ClinicalRecommendationStatement", resource.ClinicalRecommendationStatement)
}
func (resource *Measure) T_ImprovementNotation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.ImprovementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.ImprovementNotation", resource.ImprovementNotation, optionsValueSet)
}
func (resource *Measure) T_Guidance() templ.Component {

	if resource == nil {
		return StringInput("Measure.Guidance", nil)
	}
	return StringInput("Measure.Guidance", resource.Guidance)
}
func (resource *Measure) T_TermId(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Measure.Term["+strconv.Itoa(numTerm)+"].Id", nil)
	}
	return StringInput("Measure.Term["+strconv.Itoa(numTerm)+"].Id", resource.Term[numTerm].Id)
}
func (resource *Measure) T_TermCode(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return CodeableConceptSelect("Measure.Term["+strconv.Itoa(numTerm)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Term["+strconv.Itoa(numTerm)+"].Code", resource.Term[numTerm].Code, optionsValueSet)
}
func (resource *Measure) T_TermDefinition(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Measure.Term["+strconv.Itoa(numTerm)+"].Definition", nil)
	}
	return StringInput("Measure.Term["+strconv.Itoa(numTerm)+"].Definition", resource.Term[numTerm].Definition)
}
func (resource *Measure) T_GroupId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Id", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Id", resource.Group[numGroup].Id)
}
func (resource *Measure) T_GroupLinkId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].LinkId", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].LinkId", resource.Group[numGroup].LinkId)
}
func (resource *Measure) T_GroupCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Code", resource.Group[numGroup].Code, optionsValueSet)
}
func (resource *Measure) T_GroupDescription(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Description", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Description", resource.Group[numGroup].Description)
}
func (resource *Measure) T_GroupType(numGroup int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Type) >= numType {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Type["+strconv.Itoa(numType)+"]", &resource.Group[numGroup].Type[numType], optionsValueSet)
}
func (resource *Measure) T_GroupBasis(numGroup int) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Basis", nil, optionsValueSet)
	}
	return CodeSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Basis", resource.Group[numGroup].Basis, optionsValueSet)
}
func (resource *Measure) T_GroupScoring(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Scoring", resource.Group[numGroup].Scoring, optionsValueSet)
}
func (resource *Measure) T_GroupScoringUnit(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].ScoringUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].ScoringUnit", resource.Group[numGroup].ScoringUnit, optionsValueSet)
}
func (resource *Measure) T_GroupRateAggregation(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].RateAggregation", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].RateAggregation", resource.Group[numGroup].RateAggregation)
}
func (resource *Measure) T_GroupImprovementNotation(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].ImprovementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].ImprovementNotation", resource.Group[numGroup].ImprovementNotation, optionsValueSet)
}
func (resource *Measure) T_GroupLibrary(numGroup int, numLibrary int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Library) >= numLibrary {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Library["+strconv.Itoa(numLibrary)+"]", &resource.Group[numGroup].Library[numLibrary])
}
func (resource *Measure) T_GroupPopulationId(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Id", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Id", resource.Group[numGroup].Population[numPopulation].Id)
}
func (resource *Measure) T_GroupPopulationLinkId(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].LinkId", resource.Group[numGroup].Population[numPopulation].LinkId)
}
func (resource *Measure) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet)
}
func (resource *Measure) T_GroupPopulationDescription(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Description", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].Description", resource.Group[numGroup].Population[numPopulation].Description)
}
func (resource *Measure) T_GroupPopulationInputPopulationId(numGroup int, numPopulation int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].InputPopulationId", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].InputPopulationId", resource.Group[numGroup].Population[numPopulation].InputPopulationId)
}
func (resource *Measure) T_GroupPopulationAggregateMethod(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].AggregateMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Population["+strconv.Itoa(numPopulation)+"].AggregateMethod", resource.Group[numGroup].Population[numPopulation].AggregateMethod, optionsValueSet)
}
func (resource *Measure) T_GroupStratifierId(numGroup int, numStratifier int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Id", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Id)
}
func (resource *Measure) T_GroupStratifierLinkId(numGroup int, numStratifier int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId)
}
func (resource *Measure) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet)
}
func (resource *Measure) T_GroupStratifierDescription(numGroup int, numStratifier int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Description", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Description", resource.Group[numGroup].Stratifier[numStratifier].Description)
}
func (resource *Measure) T_GroupStratifierComponentId(numGroup int, numStratifier int, numComponent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Component) >= numComponent {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Id", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Id)
}
func (resource *Measure) T_GroupStratifierComponentLinkId(numGroup int, numStratifier int, numComponent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Component) >= numComponent {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].LinkId", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].LinkId)
}
func (resource *Measure) T_GroupStratifierComponentCode(numGroup int, numStratifier int, numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Component) >= numComponent {
		return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Code, optionsValueSet)
}
func (resource *Measure) T_GroupStratifierComponentDescription(numGroup int, numStratifier int, numComponent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Stratifier) >= numStratifier || len(resource.Group[numGroup].Stratifier[numStratifier].Component) >= numComponent {
		return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Description", nil)
	}
	return StringInput("Measure.Group["+strconv.Itoa(numGroup)+"].Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Description", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Description)
}
func (resource *Measure) T_SupplementalDataId(numSupplementalData int) templ.Component {

	if resource == nil || len(resource.SupplementalData) >= numSupplementalData {
		return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Id", nil)
	}
	return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Id", resource.SupplementalData[numSupplementalData].Id)
}
func (resource *Measure) T_SupplementalDataLinkId(numSupplementalData int) templ.Component {

	if resource == nil || len(resource.SupplementalData) >= numSupplementalData {
		return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].LinkId", nil)
	}
	return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].LinkId", resource.SupplementalData[numSupplementalData].LinkId)
}
func (resource *Measure) T_SupplementalDataCode(numSupplementalData int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupplementalData) >= numSupplementalData {
		return CodeableConceptSelect("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Code", resource.SupplementalData[numSupplementalData].Code, optionsValueSet)
}
func (resource *Measure) T_SupplementalDataUsage(numSupplementalData int, numUsage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupplementalData) >= numSupplementalData || len(resource.SupplementalData[numSupplementalData].Usage) >= numUsage {
		return CodeableConceptSelect("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Usage["+strconv.Itoa(numUsage)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Usage["+strconv.Itoa(numUsage)+"]", &resource.SupplementalData[numSupplementalData].Usage[numUsage], optionsValueSet)
}
func (resource *Measure) T_SupplementalDataDescription(numSupplementalData int) templ.Component {

	if resource == nil || len(resource.SupplementalData) >= numSupplementalData {
		return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Description", nil)
	}
	return StringInput("Measure.SupplementalData["+strconv.Itoa(numSupplementalData)+"].Description", resource.SupplementalData[numSupplementalData].Description)
}
