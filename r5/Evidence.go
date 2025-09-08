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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type Evidence struct {
	Id                     *string                      `json:"id,omitempty"`
	Meta                   *Meta                        `json:"meta,omitempty"`
	ImplicitRules          *string                      `json:"implicitRules,omitempty"`
	Language               *string                      `json:"language,omitempty"`
	Text                   *Narrative                   `json:"text,omitempty"`
	Contained              []Resource                   `json:"contained,omitempty"`
	Extension              []Extension                  `json:"extension,omitempty"`
	ModifierExtension      []Extension                  `json:"modifierExtension,omitempty"`
	Url                    *string                      `json:"url,omitempty"`
	Identifier             []Identifier                 `json:"identifier,omitempty"`
	Version                *string                      `json:"version,omitempty"`
	VersionAlgorithmString *string                      `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                      `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                      `json:"name,omitempty"`
	Title                  *string                      `json:"title,omitempty"`
	CiteAsReference        *Reference                   `json:"citeAsReference,omitempty"`
	CiteAsMarkdown         *string                      `json:"citeAsMarkdown,omitempty"`
	Status                 string                       `json:"status"`
	Experimental           *bool                        `json:"experimental,omitempty"`
	Date                   *time.Time                   `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ApprovalDate           *time.Time                   `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                   `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
	Publisher              *string                      `json:"publisher,omitempty"`
	Contact                []ContactDetail              `json:"contact,omitempty"`
	Author                 []ContactDetail              `json:"author,omitempty"`
	Editor                 []ContactDetail              `json:"editor,omitempty"`
	Reviewer               []ContactDetail              `json:"reviewer,omitempty"`
	Endorser               []ContactDetail              `json:"endorser,omitempty"`
	UseContext             []UsageContext               `json:"useContext,omitempty"`
	Purpose                *string                      `json:"purpose,omitempty"`
	Copyright              *string                      `json:"copyright,omitempty"`
	CopyrightLabel         *string                      `json:"copyrightLabel,omitempty"`
	RelatedArtifact        []RelatedArtifact            `json:"relatedArtifact,omitempty"`
	Description            *string                      `json:"description,omitempty"`
	Assertion              *string                      `json:"assertion,omitempty"`
	Note                   []Annotation                 `json:"note,omitempty"`
	VariableDefinition     []EvidenceVariableDefinition `json:"variableDefinition"`
	SynthesisType          *CodeableConcept             `json:"synthesisType,omitempty"`
	StudyDesign            []CodeableConcept            `json:"studyDesign,omitempty"`
	Statistic              []EvidenceStatistic          `json:"statistic,omitempty"`
	Certainty              []EvidenceCertainty          `json:"certainty,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceVariableDefinition struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	VariableRole      CodeableConcept  `json:"variableRole"`
	Observed          *Reference       `json:"observed,omitempty"`
	Intended          *Reference       `json:"intended,omitempty"`
	DirectnessMatch   *CodeableConcept `json:"directnessMatch,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatistic struct {
	Id                  *string                                `json:"id,omitempty"`
	Extension           []Extension                            `json:"extension,omitempty"`
	ModifierExtension   []Extension                            `json:"modifierExtension,omitempty"`
	Description         *string                                `json:"description,omitempty"`
	Note                []Annotation                           `json:"note,omitempty"`
	StatisticType       *CodeableConcept                       `json:"statisticType,omitempty"`
	Category            *CodeableConcept                       `json:"category,omitempty"`
	Quantity            *Quantity                              `json:"quantity,omitempty"`
	NumberOfEvents      *int                                   `json:"numberOfEvents,omitempty"`
	NumberAffected      *int                                   `json:"numberAffected,omitempty"`
	SampleSize          *EvidenceStatisticSampleSize           `json:"sampleSize,omitempty"`
	AttributeEstimate   []EvidenceStatisticAttributeEstimate   `json:"attributeEstimate,omitempty"`
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatisticSampleSize struct {
	Id                   *string      `json:"id,omitempty"`
	Extension            []Extension  `json:"extension,omitempty"`
	ModifierExtension    []Extension  `json:"modifierExtension,omitempty"`
	Description          *string      `json:"description,omitempty"`
	Note                 []Annotation `json:"note,omitempty"`
	NumberOfStudies      *int         `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int         `json:"numberOfParticipants,omitempty"`
	KnownDataCount       *int         `json:"knownDataCount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatisticAttributeEstimate struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
	Level             *float64         `json:"level,omitempty"`
	Range             *Range           `json:"range,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatisticModelCharacteristic struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                `json:"code"`
	Value             *Quantity                                      `json:"value,omitempty"`
	Variable          []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatisticModelCharacteristicVariable struct {
	Id                 *string           `json:"id,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	VariableDefinition Reference         `json:"variableDefinition"`
	Handling           *string           `json:"handling,omitempty"`
	ValueCategory      []CodeableConcept `json:"valueCategory,omitempty"`
	ValueQuantity      []Quantity        `json:"valueQuantity,omitempty"`
	ValueRange         []Range           `json:"valueRange,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceCertainty struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Rating            *CodeableConcept `json:"rating,omitempty"`
	Rater             *string          `json:"rater,omitempty"`
}

type OtherEvidence Evidence

// on convert struct to json, automatically add resourceType=Evidence
func (r Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidence: OtherEvidence(r),
		ResourceType:  "Evidence",
	})
}
func (r Evidence) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Evidence/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Evidence"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Evidence) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Url", nil, htmlAttrs)
	}
	return StringInput("Evidence.Url", resource.Url, htmlAttrs)
}
func (resource *Evidence) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Version", nil, htmlAttrs)
	}
	return StringInput("Evidence.Version", resource.Version, htmlAttrs)
}
func (resource *Evidence) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("Evidence.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Evidence) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("Evidence.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Evidence.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Name", nil, htmlAttrs)
	}
	return StringInput("Evidence.Name", resource.Name, htmlAttrs)
}
func (resource *Evidence) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Title", nil, htmlAttrs)
	}
	return StringInput("Evidence.Title", resource.Title, htmlAttrs)
}
func (resource *Evidence) T_CiteAsMarkdown(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.CiteAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("Evidence.CiteAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *Evidence) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Evidence.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Evidence.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Evidence.Experimental", nil, htmlAttrs)
	}
	return BoolInput("Evidence.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Evidence) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Evidence.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Evidence.Date", resource.Date, htmlAttrs)
}
func (resource *Evidence) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Evidence.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("Evidence.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Evidence) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Evidence.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("Evidence.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Evidence) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Publisher", nil, htmlAttrs)
	}
	return StringInput("Evidence.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Evidence) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Purpose", nil, htmlAttrs)
	}
	return StringInput("Evidence.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Evidence) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Copyright", nil, htmlAttrs)
	}
	return StringInput("Evidence.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Evidence) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("Evidence.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Evidence) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.Description", resource.Description, htmlAttrs)
}
func (resource *Evidence) T_Assertion(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Evidence.Assertion", nil, htmlAttrs)
	}
	return StringInput("Evidence.Assertion", resource.Assertion, htmlAttrs)
}
func (resource *Evidence) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Evidence.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_SynthesisType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Evidence.SynthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.SynthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StudyDesign(numStudyDesign int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStudyDesign >= len(resource.StudyDesign) {
		return CodeableConceptSelect("Evidence.StudyDesign."+strconv.Itoa(numStudyDesign)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.StudyDesign."+strconv.Itoa(numStudyDesign)+".", &resource.StudyDesign[numStudyDesign], optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionDescription(numVariableDefinition int, htmlAttrs string) templ.Component {

	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return StringInput("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..Description", resource.VariableDefinition[numVariableDefinition].Description, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionNote(numVariableDefinition int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) || numNote >= len(resource.VariableDefinition[numVariableDefinition].Note) {
		return AnnotationTextArea("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..Note."+strconv.Itoa(numNote)+".", &resource.VariableDefinition[numVariableDefinition].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionVariableRole(numVariableDefinition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return CodeableConceptSelect("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..VariableRole", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..VariableRole", &resource.VariableDefinition[numVariableDefinition].VariableRole, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionDirectnessMatch(numVariableDefinition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return CodeableConceptSelect("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..DirectnessMatch", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.VariableDefinition."+strconv.Itoa(numVariableDefinition)+"..DirectnessMatch", resource.VariableDefinition[numVariableDefinition].DirectnessMatch, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticDescription(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Description", resource.Statistic[numStatistic].Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticNote(numStatistic int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numNote >= len(resource.Statistic[numStatistic].Note) {
		return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Note."+strconv.Itoa(numNote)+".", &resource.Statistic[numStatistic].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticStatisticType(numStatistic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..StatisticType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..StatisticType", resource.Statistic[numStatistic].StatisticType, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticCategory(numStatistic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..Category", resource.Statistic[numStatistic].Category, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticNumberOfEvents(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..NumberOfEvents", nil, htmlAttrs)
	}
	return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..NumberOfEvents", resource.Statistic[numStatistic].NumberOfEvents, htmlAttrs)
}
func (resource *Evidence) T_StatisticNumberAffected(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..NumberAffected", nil, htmlAttrs)
	}
	return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..NumberAffected", resource.Statistic[numStatistic].NumberAffected, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeDescription(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.Description", resource.Statistic[numStatistic].SampleSize.Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNote(numStatistic int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numNote >= len(resource.Statistic[numStatistic].SampleSize.Note) {
		return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.Note."+strconv.Itoa(numNote)+".", &resource.Statistic[numStatistic].SampleSize.Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfStudies(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.NumberOfStudies", nil, htmlAttrs)
	}
	return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.NumberOfStudies", resource.Statistic[numStatistic].SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfParticipants(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.NumberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.NumberOfParticipants", resource.Statistic[numStatistic].SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeKnownDataCount(numStatistic int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.KnownDataCount", nil, htmlAttrs)
	}
	return IntInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..SampleSize.KnownDataCount", resource.Statistic[numStatistic].SampleSize.KnownDataCount, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateDescription(numStatistic int, numAttributeEstimate int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Description", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateNote(numStatistic int, numAttributeEstimate int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) || numNote >= len(resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Note) {
		return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Note."+strconv.Itoa(numNote)+".", &resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateType(numStatistic int, numAttributeEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Type", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateLevel(numStatistic int, numAttributeEstimate int, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return Float64Input("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Level", nil, htmlAttrs)
	}
	return Float64Input("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..AttributeEstimate."+strconv.Itoa(numAttributeEstimate)+"..Level", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Level, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicCode(numStatistic int, numModelCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) {
		return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Code", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableHandling(numStatistic int, numModelCharacteristic int, numVariable int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) {
		return CodeSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Variable."+strconv.Itoa(numVariable)+"..Handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Variable."+strconv.Itoa(numVariable)+"..Handling", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Handling, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueCategory(numStatistic int, numModelCharacteristic int, numVariable int, numValueCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) || numValueCategory >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory) {
		return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Variable."+strconv.Itoa(numVariable)+"..ValueCategory."+strconv.Itoa(numValueCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Statistic."+strconv.Itoa(numStatistic)+"..ModelCharacteristic."+strconv.Itoa(numModelCharacteristic)+"..Variable."+strconv.Itoa(numVariable)+"..ValueCategory."+strconv.Itoa(numValueCategory)+".", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory[numValueCategory], optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_CertaintyDescription(numCertainty int, htmlAttrs string) templ.Component {

	if resource == nil || numCertainty >= len(resource.Certainty) {
		return StringInput("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Description", nil, htmlAttrs)
	}
	return StringInput("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Description", resource.Certainty[numCertainty].Description, htmlAttrs)
}
func (resource *Evidence) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Note."+strconv.Itoa(numNote)+".", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_CertaintyType(numCertainty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCertainty >= len(resource.Certainty) {
		return CodeableConceptSelect("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Type", resource.Certainty[numCertainty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_CertaintyRating(numCertainty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCertainty >= len(resource.Certainty) {
		return CodeableConceptSelect("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Rating", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Rating", resource.Certainty[numCertainty].Rating, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_CertaintyRater(numCertainty int, htmlAttrs string) templ.Component {

	if resource == nil || numCertainty >= len(resource.Certainty) {
		return StringInput("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Rater", nil, htmlAttrs)
	}
	return StringInput("Evidence.Certainty."+strconv.Itoa(numCertainty)+"..Rater", resource.Certainty[numCertainty].Rater, htmlAttrs)
}
