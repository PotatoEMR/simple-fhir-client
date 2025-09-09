package r5

//generated with command go run ./bultaoreune
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
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Measure) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Measure) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Measure) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Measure) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Measure) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Measure) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *Measure) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Basis(htmlAttrs string) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil {
		return CodeSelect("basis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("basis", resource.Basis, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Measure) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *Measure) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Measure) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *Measure) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *Measure) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *Measure) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Measure) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Measure) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Measure) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *Measure) T_Disclaimer(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("disclaimer", nil, htmlAttrs)
	}
	return StringInput("disclaimer", resource.Disclaimer, htmlAttrs)
}
func (resource *Measure) T_Scoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("scoring", resource.Scoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_ScoringUnit(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("scoringUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("scoringUnit", resource.ScoringUnit, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_CompositeScoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("compositeScoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("compositeScoring", resource.CompositeScoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_RiskAdjustment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("riskAdjustment", nil, htmlAttrs)
	}
	return StringInput("riskAdjustment", resource.RiskAdjustment, htmlAttrs)
}
func (resource *Measure) T_RateAggregation(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("rateAggregation", nil, htmlAttrs)
	}
	return StringInput("rateAggregation", resource.RateAggregation, htmlAttrs)
}
func (resource *Measure) T_Rationale(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("rationale", nil, htmlAttrs)
	}
	return StringInput("rationale", resource.Rationale, htmlAttrs)
}
func (resource *Measure) T_ClinicalRecommendationStatement(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("clinicalRecommendationStatement", nil, htmlAttrs)
	}
	return StringInput("clinicalRecommendationStatement", resource.ClinicalRecommendationStatement, htmlAttrs)
}
func (resource *Measure) T_ImprovementNotation(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("improvementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("improvementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Guidance(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("guidance", nil, htmlAttrs)
	}
	return StringInput("guidance", resource.Guidance, htmlAttrs)
}
func (resource *Measure) T_TermCode(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].code", resource.Term[numTerm].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_TermDefinition(numTerm int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].definition", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].definition", resource.Term[numTerm].Definition, htmlAttrs)
}
func (resource *Measure) T_GroupLinkId(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].linkId", resource.Group[numGroup].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupDescription(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].description", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].description", resource.Group[numGroup].Description, htmlAttrs)
}
func (resource *Measure) T_GroupType(numGroup int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numType >= len(resource.Group[numGroup].Type) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].type["+strconv.Itoa(numType)+"]", &resource.Group[numGroup].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupSubjectCodeableConcept(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].subjectCodeableConcept", resource.Group[numGroup].SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupBasis(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].basis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].basis", resource.Group[numGroup].Basis, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupScoring(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].scoring", resource.Group[numGroup].Scoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupScoringUnit(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].scoringUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].scoringUnit", resource.Group[numGroup].ScoringUnit, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupRateAggregation(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rateAggregation", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rateAggregation", resource.Group[numGroup].RateAggregation, htmlAttrs)
}
func (resource *Measure) T_GroupImprovementNotation(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].improvementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].improvementNotation", resource.Group[numGroup].ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupLibrary(numGroup int, numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numLibrary >= len(resource.Group[numGroup].Library) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].library["+strconv.Itoa(numLibrary)+"]", &resource.Group[numGroup].Library[numLibrary], htmlAttrs)
}
func (resource *Measure) T_GroupPopulationLinkId(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].linkId", resource.Group[numGroup].Population[numPopulation].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationDescription(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].description", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].description", resource.Group[numGroup].Population[numPopulation].Description, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationInputPopulationId(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].inputPopulationId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].inputPopulationId", resource.Group[numGroup].Population[numPopulation].InputPopulationId, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationAggregateMethod(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].aggregateMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].population["+strconv.Itoa(numPopulation)+"].aggregateMethod", resource.Group[numGroup].Population[numPopulation].AggregateMethod, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierLinkId(numGroup int, numStratifier int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].linkId", resource.Group[numGroup].Stratifier[numStratifier].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierDescription(numGroup int, numStratifier int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].description", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].description", resource.Group[numGroup].Stratifier[numStratifier].Description, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentLinkId(numGroup int, numStratifier int, numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].linkId", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].LinkId, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentCode(numGroup int, numStratifier int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].code", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentDescription(numGroup int, numStratifier int, numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].description", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].stratifier["+strconv.Itoa(numStratifier)+"].component["+strconv.Itoa(numComponent)+"].description", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Description, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataLinkId(numSupplementalData int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return StringInput("supplementalData["+strconv.Itoa(numSupplementalData)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("supplementalData["+strconv.Itoa(numSupplementalData)+"].linkId", resource.SupplementalData[numSupplementalData].LinkId, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataCode(numSupplementalData int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return CodeableConceptSelect("supplementalData["+strconv.Itoa(numSupplementalData)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supplementalData["+strconv.Itoa(numSupplementalData)+"].code", resource.SupplementalData[numSupplementalData].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataUsage(numSupplementalData int, numUsage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) || numUsage >= len(resource.SupplementalData[numSupplementalData].Usage) {
		return CodeableConceptSelect("supplementalData["+strconv.Itoa(numSupplementalData)+"].usage["+strconv.Itoa(numUsage)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supplementalData["+strconv.Itoa(numSupplementalData)+"].usage["+strconv.Itoa(numUsage)+"]", &resource.SupplementalData[numSupplementalData].Usage[numUsage], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataDescription(numSupplementalData int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return StringInput("supplementalData["+strconv.Itoa(numSupplementalData)+"].description", nil, htmlAttrs)
	}
	return StringInput("supplementalData["+strconv.Itoa(numSupplementalData)+"].description", resource.SupplementalData[numSupplementalData].Description, htmlAttrs)
}
