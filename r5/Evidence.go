package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *string                      `json:"date,omitempty"`
	ApprovalDate           *string                      `json:"approvalDate,omitempty"`
	LastReviewDate         *string                      `json:"lastReviewDate,omitempty"`
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

func (resource *Evidence) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Id", nil)
	}
	return StringInput("Evidence.Id", resource.Id)
}
func (resource *Evidence) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Evidence.ImplicitRules", nil)
	}
	return StringInput("Evidence.ImplicitRules", resource.ImplicitRules)
}
func (resource *Evidence) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Evidence.Language", nil, optionsValueSet)
	}
	return CodeSelect("Evidence.Language", resource.Language, optionsValueSet)
}
func (resource *Evidence) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Url", nil)
	}
	return StringInput("Evidence.Url", resource.Url)
}
func (resource *Evidence) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Version", nil)
	}
	return StringInput("Evidence.Version", resource.Version)
}
func (resource *Evidence) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Name", nil)
	}
	return StringInput("Evidence.Name", resource.Name)
}
func (resource *Evidence) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Title", nil)
	}
	return StringInput("Evidence.Title", resource.Title)
}
func (resource *Evidence) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Evidence.Status", nil, optionsValueSet)
	}
	return CodeSelect("Evidence.Status", &resource.Status, optionsValueSet)
}
func (resource *Evidence) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Evidence.Experimental", nil)
	}
	return BoolInput("Evidence.Experimental", resource.Experimental)
}
func (resource *Evidence) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Date", nil)
	}
	return StringInput("Evidence.Date", resource.Date)
}
func (resource *Evidence) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Evidence.ApprovalDate", nil)
	}
	return StringInput("Evidence.ApprovalDate", resource.ApprovalDate)
}
func (resource *Evidence) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Evidence.LastReviewDate", nil)
	}
	return StringInput("Evidence.LastReviewDate", resource.LastReviewDate)
}
func (resource *Evidence) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Publisher", nil)
	}
	return StringInput("Evidence.Publisher", resource.Publisher)
}
func (resource *Evidence) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Purpose", nil)
	}
	return StringInput("Evidence.Purpose", resource.Purpose)
}
func (resource *Evidence) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Copyright", nil)
	}
	return StringInput("Evidence.Copyright", resource.Copyright)
}
func (resource *Evidence) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("Evidence.CopyrightLabel", nil)
	}
	return StringInput("Evidence.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *Evidence) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Description", nil)
	}
	return StringInput("Evidence.Description", resource.Description)
}
func (resource *Evidence) T_Assertion() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Assertion", nil)
	}
	return StringInput("Evidence.Assertion", resource.Assertion)
}
func (resource *Evidence) T_SynthesisType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Evidence.SynthesisType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.SynthesisType", resource.SynthesisType, optionsValueSet)
}
func (resource *Evidence) T_StudyDesign(numStudyDesign int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StudyDesign) >= numStudyDesign {
		return CodeableConceptSelect("Evidence.StudyDesign["+strconv.Itoa(numStudyDesign)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.StudyDesign["+strconv.Itoa(numStudyDesign)+"]", &resource.StudyDesign[numStudyDesign], optionsValueSet)
}
func (resource *Evidence) T_VariableDefinitionId(numVariableDefinition int) templ.Component {

	if resource == nil || len(resource.VariableDefinition) >= numVariableDefinition {
		return StringInput("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].Id", nil)
	}
	return StringInput("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].Id", resource.VariableDefinition[numVariableDefinition].Id)
}
func (resource *Evidence) T_VariableDefinitionDescription(numVariableDefinition int) templ.Component {

	if resource == nil || len(resource.VariableDefinition) >= numVariableDefinition {
		return StringInput("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].Description", nil)
	}
	return StringInput("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].Description", resource.VariableDefinition[numVariableDefinition].Description)
}
func (resource *Evidence) T_VariableDefinitionVariableRole(numVariableDefinition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.VariableDefinition) >= numVariableDefinition {
		return CodeableConceptSelect("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].VariableRole", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].VariableRole", &resource.VariableDefinition[numVariableDefinition].VariableRole, optionsValueSet)
}
func (resource *Evidence) T_VariableDefinitionDirectnessMatch(numVariableDefinition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.VariableDefinition) >= numVariableDefinition {
		return CodeableConceptSelect("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].DirectnessMatch", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.VariableDefinition["+strconv.Itoa(numVariableDefinition)+"].DirectnessMatch", resource.VariableDefinition[numVariableDefinition].DirectnessMatch, optionsValueSet)
}
func (resource *Evidence) T_StatisticId(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Id", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Id", resource.Statistic[numStatistic].Id)
}
func (resource *Evidence) T_StatisticDescription(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Description", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Description", resource.Statistic[numStatistic].Description)
}
func (resource *Evidence) T_StatisticStatisticType(numStatistic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].StatisticType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].StatisticType", resource.Statistic[numStatistic].StatisticType, optionsValueSet)
}
func (resource *Evidence) T_StatisticCategory(numStatistic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].Category", resource.Statistic[numStatistic].Category, optionsValueSet)
}
func (resource *Evidence) T_StatisticNumberOfEvents(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].NumberOfEvents", nil)
	}
	return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].NumberOfEvents", resource.Statistic[numStatistic].NumberOfEvents)
}
func (resource *Evidence) T_StatisticNumberAffected(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].NumberAffected", nil)
	}
	return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].NumberAffected", resource.Statistic[numStatistic].NumberAffected)
}
func (resource *Evidence) T_StatisticSampleSizeId(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.Id", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.Id", resource.Statistic[numStatistic].SampleSize.Id)
}
func (resource *Evidence) T_StatisticSampleSizeDescription(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.Description", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.Description", resource.Statistic[numStatistic].SampleSize.Description)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfStudies(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.NumberOfStudies", nil)
	}
	return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.NumberOfStudies", resource.Statistic[numStatistic].SampleSize.NumberOfStudies)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfParticipants(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.NumberOfParticipants", nil)
	}
	return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.NumberOfParticipants", resource.Statistic[numStatistic].SampleSize.NumberOfParticipants)
}
func (resource *Evidence) T_StatisticSampleSizeKnownDataCount(numStatistic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic {
		return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.KnownDataCount", nil)
	}
	return IntInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].SampleSize.KnownDataCount", resource.Statistic[numStatistic].SampleSize.KnownDataCount)
}
func (resource *Evidence) T_StatisticAttributeEstimateId(numStatistic int, numAttributeEstimate int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].AttributeEstimate) >= numAttributeEstimate {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Id", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Id", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Id)
}
func (resource *Evidence) T_StatisticAttributeEstimateDescription(numStatistic int, numAttributeEstimate int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].AttributeEstimate) >= numAttributeEstimate {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Description", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Description", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Description)
}
func (resource *Evidence) T_StatisticAttributeEstimateType(numStatistic int, numAttributeEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].AttributeEstimate) >= numAttributeEstimate {
		return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Type", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Type, optionsValueSet)
}
func (resource *Evidence) T_StatisticAttributeEstimateLevel(numStatistic int, numAttributeEstimate int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].AttributeEstimate) >= numAttributeEstimate {
		return Float64Input("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Level", nil)
	}
	return Float64Input("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].AttributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].Level", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Level)
}
func (resource *Evidence) T_StatisticModelCharacteristicId(numStatistic int, numModelCharacteristic int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Id", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Id", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Id)
}
func (resource *Evidence) T_StatisticModelCharacteristicCode(numStatistic int, numModelCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic {
		return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Code", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Code, optionsValueSet)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableId(numStatistic int, numModelCharacteristic int, numVariable int) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic || len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable {
		return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].Id", nil)
	}
	return StringInput("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].Id", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Id)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableHandling(numStatistic int, numModelCharacteristic int, numVariable int) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic || len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable {
		return CodeSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].Handling", nil, optionsValueSet)
	}
	return CodeSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].Handling", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Handling, optionsValueSet)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueCategory(numStatistic int, numModelCharacteristic int, numVariable int, numValueCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Statistic) >= numStatistic || len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic || len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable || len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory) >= numValueCategory {
		return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].ValueCategory["+strconv.Itoa(numValueCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Statistic["+strconv.Itoa(numStatistic)+"].ModelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].Variable["+strconv.Itoa(numVariable)+"].ValueCategory["+strconv.Itoa(numValueCategory)+"]", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory[numValueCategory], optionsValueSet)
}
func (resource *Evidence) T_CertaintyId(numCertainty int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Id", nil)
	}
	return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Id", resource.Certainty[numCertainty].Id)
}
func (resource *Evidence) T_CertaintyDescription(numCertainty int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Description", nil)
	}
	return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Description", resource.Certainty[numCertainty].Description)
}
func (resource *Evidence) T_CertaintyType(numCertainty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return CodeableConceptSelect("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Type", resource.Certainty[numCertainty].Type, optionsValueSet)
}
func (resource *Evidence) T_CertaintyRating(numCertainty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return CodeableConceptSelect("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Rating", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Rating", resource.Certainty[numCertainty].Rating, optionsValueSet)
}
func (resource *Evidence) T_CertaintyRater(numCertainty int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Rater", nil)
	}
	return StringInput("Evidence.Certainty["+strconv.Itoa(numCertainty)+"].Rater", resource.Certainty[numCertainty].Rater)
}
