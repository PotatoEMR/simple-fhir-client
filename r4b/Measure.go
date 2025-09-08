package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
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
	Name                            *string                   `json:"name,omitempty"`
	Title                           *string                   `json:"title,omitempty"`
	Subtitle                        *string                   `json:"subtitle,omitempty"`
	Status                          string                    `json:"status"`
	Experimental                    *bool                     `json:"experimental,omitempty"`
	SubjectCodeableConcept          *CodeableConcept          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference                *Reference                `json:"subjectReference,omitempty"`
	Date                            *time.Time                `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher                       *string                   `json:"publisher,omitempty"`
	Contact                         []ContactDetail           `json:"contact,omitempty"`
	Description                     *string                   `json:"description,omitempty"`
	UseContext                      []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                         *string                   `json:"purpose,omitempty"`
	Usage                           *string                   `json:"usage,omitempty"`
	Copyright                       *string                   `json:"copyright,omitempty"`
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
	CompositeScoring                *CodeableConcept          `json:"compositeScoring,omitempty"`
	Type                            []CodeableConcept         `json:"type,omitempty"`
	RiskAdjustment                  *string                   `json:"riskAdjustment,omitempty"`
	RateAggregation                 *string                   `json:"rateAggregation,omitempty"`
	Rationale                       *string                   `json:"rationale,omitempty"`
	ClinicalRecommendationStatement *string                   `json:"clinicalRecommendationStatement,omitempty"`
	ImprovementNotation             *CodeableConcept          `json:"improvementNotation,omitempty"`
	Definition                      []string                  `json:"definition,omitempty"`
	Guidance                        *string                   `json:"guidance,omitempty"`
	Group                           []MeasureGroup            `json:"group,omitempty"`
	SupplementalData                []MeasureSupplementalData `json:"supplementalData,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
type MeasureGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept         `json:"code,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Population        []MeasureGroupPopulation `json:"population,omitempty"`
	Stratifier        []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
type MeasureGroupPopulation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          Expression       `json:"criteria"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
type MeasureGroupStratifier struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept                  `json:"code,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	Criteria          *Expression                       `json:"criteria,omitempty"`
	Component         []MeasureGroupStratifierComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
type MeasureGroupStratifierComponent struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          Expression       `json:"criteria"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Measure
type MeasureSupplementalData struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
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
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *Measure) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *Measure) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *Measure) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *Measure) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Subtitle", nil, htmlAttrs)
	}
	return StringInput("Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Measure) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Measure) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *Measure) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Measure) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *Measure) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Measure) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Usage", nil, htmlAttrs)
	}
	return StringInput("Usage", resource.Usage, htmlAttrs)
}
func (resource *Measure) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Measure) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Measure) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Measure) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("Library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *Measure) T_Disclaimer(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Disclaimer", nil, htmlAttrs)
	}
	return StringInput("Disclaimer", resource.Disclaimer, htmlAttrs)
}
func (resource *Measure) T_Scoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Scoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Scoring", resource.Scoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_CompositeScoring(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CompositeScoring", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CompositeScoring", resource.CompositeScoring, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_RiskAdjustment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("RiskAdjustment", nil, htmlAttrs)
	}
	return StringInput("RiskAdjustment", resource.RiskAdjustment, htmlAttrs)
}
func (resource *Measure) T_RateAggregation(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("RateAggregation", nil, htmlAttrs)
	}
	return StringInput("RateAggregation", resource.RateAggregation, htmlAttrs)
}
func (resource *Measure) T_Rationale(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Rationale", nil, htmlAttrs)
	}
	return StringInput("Rationale", resource.Rationale, htmlAttrs)
}
func (resource *Measure) T_ClinicalRecommendationStatement(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ClinicalRecommendationStatement", nil, htmlAttrs)
	}
	return StringInput("ClinicalRecommendationStatement", resource.ClinicalRecommendationStatement, htmlAttrs)
}
func (resource *Measure) T_ImprovementNotation(htmlAttrs string) templ.Component {
	optionsValueSet := VSMeasure_improvement_notation

	if resource == nil {
		return CodeableConceptSelect("ImprovementNotation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImprovementNotation", resource.ImprovementNotation, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_Definition(numDefinition int, htmlAttrs string) templ.Component {
	if resource == nil || numDefinition >= len(resource.Definition) {
		return StringInput("Definition["+strconv.Itoa(numDefinition)+"]", nil, htmlAttrs)
	}
	return StringInput("Definition["+strconv.Itoa(numDefinition)+"]", &resource.Definition[numDefinition], htmlAttrs)
}
func (resource *Measure) T_Guidance(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Guidance", nil, htmlAttrs)
	}
	return StringInput("Guidance", resource.Guidance, htmlAttrs)
}
func (resource *Measure) T_GroupCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Code", resource.Group[numGroup].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupDescription(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Description", resource.Group[numGroup].Description, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationCode(numGroup int, numPopulation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Code", resource.Group[numGroup].Population[numPopulation].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupPopulationDescription(numGroup int, numPopulation int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numPopulation >= len(resource.Group[numGroup].Population) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Population["+strconv.Itoa(numPopulation)+"].Description", resource.Group[numGroup].Population[numPopulation].Description, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierCode(numGroup int, numStratifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierDescription(numGroup int, numStratifier int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Description", resource.Group[numGroup].Stratifier[numStratifier].Description, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentCode(numGroup int, numStratifier int, numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Code", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_GroupStratifierComponentDescription(numGroup int, numStratifier int, numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numStratifier >= len(resource.Group[numGroup].Stratifier) || numComponent >= len(resource.Group[numGroup].Stratifier[numStratifier].Component) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Stratifier["+strconv.Itoa(numStratifier)+"].Component["+strconv.Itoa(numComponent)+"].Description", resource.Group[numGroup].Stratifier[numStratifier].Component[numComponent].Description, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataCode(numSupplementalData int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return CodeableConceptSelect("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Code", resource.SupplementalData[numSupplementalData].Code, optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataUsage(numSupplementalData int, numUsage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) || numUsage >= len(resource.SupplementalData[numSupplementalData].Usage) {
		return CodeableConceptSelect("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Usage["+strconv.Itoa(numUsage)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Usage["+strconv.Itoa(numUsage)+"]", &resource.SupplementalData[numSupplementalData].Usage[numUsage], optionsValueSet, htmlAttrs)
}
func (resource *Measure) T_SupplementalDataDescription(numSupplementalData int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplementalData >= len(resource.SupplementalData) {
		return StringInput("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Description", nil, htmlAttrs)
	}
	return StringInput("SupplementalData["+strconv.Itoa(numSupplementalData)+"]Description", resource.SupplementalData[numSupplementalData].Description, htmlAttrs)
}
