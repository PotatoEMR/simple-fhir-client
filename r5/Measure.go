package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString          *string                   `json:"versionAlgorithmString"`
	VersionAlgorithmCoding          *Coding                   `json:"versionAlgorithmCoding"`
	Name                            *string                   `json:"name,omitempty"`
	Title                           *string                   `json:"title,omitempty"`
	Subtitle                        *string                   `json:"subtitle,omitempty"`
	Status                          string                    `json:"status"`
	Experimental                    *bool                     `json:"experimental,omitempty"`
	SubjectCodeableConcept          *CodeableConcept          `json:"subjectCodeableConcept"`
	SubjectReference                *Reference                `json:"subjectReference"`
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
	SubjectCodeableConcept *CodeableConcept         `json:"subjectCodeableConcept"`
	SubjectReference       *Reference               `json:"subjectReference"`
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

func (resource *Measure) MeasureLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Measure) MeasureStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Measure) MeasureBasis() templ.Component {
	optionsValueSet := VSFhir_types

	if resource != nil {
		return CodeSelect("basis", nil, optionsValueSet)
	}
	return CodeSelect("basis", resource.Basis, optionsValueSet)
}
func (resource *Measure) MeasureJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *Measure) MeasureTopic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", &resource.Topic[0], optionsValueSet)
}
func (resource *Measure) MeasureScoring(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scoring", resource.Scoring, optionsValueSet)
}
func (resource *Measure) MeasureScoringUnit(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("scoringUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scoringUnit", resource.ScoringUnit, optionsValueSet)
}
func (resource *Measure) MeasureCompositeScoring(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("compositeScoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("compositeScoring", resource.CompositeScoring, optionsValueSet)
}
func (resource *Measure) MeasureType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *Measure) MeasureImprovementNotation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("improvementNotation", resource.ImprovementNotation, optionsValueSet)
}
func (resource *Measure) MeasureTermCode(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Term[numTerm].Code, optionsValueSet)
}
func (resource *Measure) MeasureGroupCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Code, optionsValueSet)
}
func (resource *Measure) MeasureGroupType(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Group[numGroup].Type[0], optionsValueSet)
}
func (resource *Measure) MeasureGroupBasis(numGroup int) templ.Component {
	optionsValueSet := VSFhir_types

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeSelect("basis", nil, optionsValueSet)
	}
	return CodeSelect("basis", resource.Group[numGroup].Basis, optionsValueSet)
}
func (resource *Measure) MeasureGroupScoring(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("scoring", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scoring", resource.Group[numGroup].Scoring, optionsValueSet)
}
func (resource *Measure) MeasureGroupScoringUnit(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("scoringUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scoringUnit", resource.Group[numGroup].ScoringUnit, optionsValueSet)
}
func (resource *Measure) MeasureGroupImprovementNotation(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("improvementNotation", resource.Group[numGroup].ImprovementNotation, optionsValueSet)
}
func (resource *Measure) MeasureGroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet)
}
func (resource *Measure) MeasureGroupPopulationAggregateMethod(numGroup int, numPopulation int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Population) >= numPopulation {
		return CodeableConceptSelect("aggregateMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("aggregateMethod", resource.Group[numGroup].Population[numPopulation].AggregateMethod, optionsValueSet)
}
func (resource *Measure) MeasureGroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Stratifier) >= numStratifier {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet)
}
func (resource *Measure) MeasureGroupStratifierComponentCode(numGroup int, numStratifier int, numComponent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Stratifier[numStratifier].Component) >= numComponent {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Code, optionsValueSet)
}
func (resource *Measure) MeasureSupplementalDataCode(numSupplementalData int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SupplementalData) >= numSupplementalData {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.SupplementalData[numSupplementalData].Code, optionsValueSet)
}
func (resource *Measure) MeasureSupplementalDataUsage(numSupplementalData int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SupplementalData) >= numSupplementalData {
		return CodeableConceptSelect("usage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("usage", &resource.SupplementalData[numSupplementalData].Usage[0], optionsValueSet)
}
