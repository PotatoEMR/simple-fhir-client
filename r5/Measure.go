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
	Date                            *time.Time                `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher                       *string                   `json:"publisher,omitempty"`
	Contact                         []ContactDetail           `json:"contact,omitempty"`
	Description                     *string                   `json:"description,omitempty"`
	UseContext                      []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                         *string                   `json:"purpose,omitempty"`
	Usage                           *string                   `json:"usage,omitempty"`
	Copyright                       *string                   `json:"copyright,omitempty"`
	CopyrightLabel                  *string                   `json:"copyrightLabel,omitempty"`
	ApprovalDate                    *time.Time                `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate                  *time.Time                `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (r Measure) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Measure/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Measure"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Measure) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Url", nil, htmlAttrs)
	}
	return StringInput("Measure.Url", resource.Url, htmlAttrs)
}
func (resource *Measure) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Version", nil, htmlAttrs)
	}
	return StringInput("Measure.Version", resource.Version, htmlAttrs)
}
func (resource *Measure) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("Measure.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Measure) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("Measure.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Measure.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Name", nil, htmlAttrs)
	}
	return StringInput("Measure.Name", resource.Name, htmlAttrs)
}
func (resource *Measure) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Title", nil, htmlAttrs)
	}
	return StringInput("Measure.Title", resource.Title, htmlAttrs)
}
func (resource *Measure) T_Subtitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Subtitle", nil, htmlAttrs)
	}
	return StringInput("Measure.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Measure) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Measure.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Measure.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Measure.Experimental", nil, htmlAttrs)
	}
	return BoolInput("Measure.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Measure) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Basis(htmlAttrs string) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil {
		return CodeSelect("Measure.Basis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Measure.Basis", resource.Basis, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Measure.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Measure.Date", resource.Date, htmlAttrs)
}
func (resource *Measure) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Publisher", nil, htmlAttrs)
	}
	return StringInput("Measure.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Measure) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Description", nil, htmlAttrs)
	}
	return StringInput("Measure.Description", resource.Description, htmlAttrs)
}
func (resource *Measure) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Measure.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Purpose", nil, htmlAttrs)
	}
	return StringInput("Measure.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Measure) T_Usage(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Usage", nil, htmlAttrs)
	}
	return StringInput("Measure.Usage", resource.Usage, htmlAttrs)
}
func (resource *Measure) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Copyright", nil, htmlAttrs)
	}
	return StringInput("Measure.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Measure) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("Measure.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Measure) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Measure.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("Measure.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Measure) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Measure.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("Measure.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Measure) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Measure.Topic."+strconv.Itoa(numTopic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Topic."+strconv.Itoa(numTopic)+".", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Library(numLibrary int, htmlAttrs string) templ.Component {

	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("Measure.Library."+strconv.Itoa(numLibrary)+".", nil, htmlAttrs)
	}
	return StringInput("Measure.Library."+strconv.Itoa(numLibrary)+".", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *Measure) T_Disclaimer(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Disclaimer", nil, htmlAttrs)
	}
	return StringInput("Measure.Disclaimer", resource.Disclaimer, htmlAttrs)
}
func (resource *Measure) T_Scoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.Scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Scoring", resource.Scoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_ScoringUnit(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.ScoringUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.ScoringUnit", resource.ScoringUnit, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_CompositeScoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.CompositeScoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.CompositeScoring", resource.CompositeScoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Measure.Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Type."+strconv.Itoa(numType)+".", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_RiskAdjustment(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.RiskAdjustment", nil, htmlAttrs)
	}
	return StringInput("Measure.RiskAdjustment", resource.RiskAdjustment, htmlAttrs)
}
func (resource *Measure) T_RateAggregation(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.RateAggregation", nil, htmlAttrs)
	}
	return StringInput("Measure.RateAggregation", resource.RateAggregation, htmlAttrs)
}
func (resource *Measure) T_Rationale(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Rationale", nil, htmlAttrs)
	}
	return StringInput("Measure.Rationale", resource.Rationale, htmlAttrs)
}
func (resource *Measure) T_ClinicalRecommendationStatement(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.ClinicalRecommendationStatement", nil, htmlAttrs)
	}
	return StringInput("Measure.ClinicalRecommendationStatement", resource.ClinicalRecommendationStatement, htmlAttrs)
}
func (resource *Measure) T_ImprovementNotation(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Measure.ImprovementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.ImprovementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Guidance(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Measure.Guidance", nil, htmlAttrs)
	}
	return StringInput("Measure.Guidance", resource.Guidance, htmlAttrs)
}
func (resource *Measure) T_TermCode(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Measure.Term."+strconv.Itoa(numTerm)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Term."+strconv.Itoa(numTerm)+"..Code", resource.Term[numTerm].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_TermDefinition(numTerm int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("Measure.Term."+strconv.Itoa(numTerm)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("Measure.Term."+strconv.Itoa(numTerm)+"..Definition", resource.Term[numTerm].Definition, htmlAttrs)
}
func (resource *Measure) T_GroupLinkId(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..LinkId", resource.Group[numGroup].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupDescription(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Description", resource.Group[numGroup].Description, htmlAttrs)
}
func (resource *Measure) T_GroupType(numGroup int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numType >= len(resource.Group[numGroup].Type) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Type."+strconv.Itoa(numType)+".", &resource.Group[numGroup].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupSubjectCodeableConcept(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..SubjectCodeableConcept", resource.Group[numGroup].SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupBasis(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Basis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Basis", resource.Group[numGroup].Basis, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupScoring(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Scoring", resource.Group[numGroup].Scoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupScoringUnit(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..ScoringUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..ScoringUnit", resource.Group[numGroup].ScoringUnit, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupRateAggregation(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..RateAggregation", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..RateAggregation", resource.Group[numGroup].RateAggregation, htmlAttrs)
}
func (resource *Measure) T_GroupImprovementNotation(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..ImprovementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..ImprovementNotation", resource.Group[numGroup].ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupLibrary(numGroup int, numLibrary int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numLibrary >= len(resource.Group[numGroup].Library) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Library."+strconv.Itoa(numLibrary)+".", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Library."+strconv.Itoa(numLibrary)+".", &resource.Group[numGroup].Library[numLibrary], htmlAttrs)
}
func (resource *Measure) T_GroupPopulationLinkId(numGroup int, numPopulation int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..LinkId", resource.Group[numGroup].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationDescription(numGroup int, numPopulation int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..Description", resource.Group[numGroup].Population[numPopulation].Description, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationInputPopulationId(numGroup int, numPopulation int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..InputPopulationId", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..InputPopulationId", resource.Group[numGroup].Population[numPopulation].InputPopulationId, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationAggregateMethod(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..AggregateMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Population."+strconv.Itoa(numPopulation)+"..AggregateMethod", resource.Group[numGroup].Population[numPopulation].AggregateMethod, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierLinkId(numGroup int, numStratifier int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..LinkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierDescription(numGroup int, numStratifier int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Description", resource.Group[numGroup].Stratifier[numStratifier].Description, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentLinkId(numGroup int, numStratifier int, numComponent int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..LinkId", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentCode(numGroup int, numStratifier int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..Code", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentDescription(numGroup int, numStratifier int, numComponent int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Measure.Group."+strconv.Itoa(numGroup)+"..Stratifier."+strconv.Itoa(numStratifier)+"..Component."+strconv.Itoa(numComponent)+"..Description", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Description, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataLinkId(numSupplementalData int, htmlAttrs string) templ.Component {

	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return StringInput("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..LinkId", resource.SupplementalData[numSupplementalData].LinkId, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataCode(numSupplementalData int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return CodeableConceptSelect("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Code", resource.SupplementalData[numSupplementalData].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataUsage(numSupplementalData int, numUsage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSupplementalData >= len(resource.SupplementalData) || numUsage >= len(resource.SupplementalData[numSupplementalData].Usage) {
		return CodeableConceptSelect("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Usage."+strconv.Itoa(numUsage)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Usage."+strconv.Itoa(numUsage)+".", &resource.SupplementalData[numSupplementalData].Usage[numUsage], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataDescription(numSupplementalData int, htmlAttrs string) templ.Component {

	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return StringInput("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Measure.SupplementalData."+strconv.Itoa(numSupplementalData)+"..Description", resource.SupplementalData[numSupplementalData].Description, htmlAttrs)
}
