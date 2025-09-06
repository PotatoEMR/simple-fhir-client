package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                 *string                        `json:"id,omitempty"`
	Meta               *Meta                          `json:"meta,omitempty"`
	ImplicitRules      *string                        `json:"implicitRules,omitempty"`
	Language           *string                        `json:"language,omitempty"`
	Text               *Narrative                     `json:"text,omitempty"`
	Contained          []Resource                     `json:"contained,omitempty"`
	Extension          []Extension                    `json:"extension,omitempty"`
	ModifierExtension  []Extension                    `json:"modifierExtension,omitempty"`
	Url                *string                        `json:"url,omitempty"`
	Identifier         []Identifier                   `json:"identifier,omitempty"`
	Version            *string                        `json:"version,omitempty"`
	Name               *string                        `json:"name,omitempty"`
	Title              *string                        `json:"title,omitempty"`
	Label              []ResearchStudyLabel           `json:"label,omitempty"`
	Protocol           []Reference                    `json:"protocol,omitempty"`
	PartOf             []Reference                    `json:"partOf,omitempty"`
	RelatedArtifact    []RelatedArtifact              `json:"relatedArtifact,omitempty"`
	Date               *string                        `json:"date,omitempty"`
	Status             string                         `json:"status"`
	PrimaryPurposeType *CodeableConcept               `json:"primaryPurposeType,omitempty"`
	Phase              *CodeableConcept               `json:"phase,omitempty"`
	StudyDesign        []CodeableConcept              `json:"studyDesign,omitempty"`
	Focus              []CodeableReference            `json:"focus,omitempty"`
	Condition          []CodeableConcept              `json:"condition,omitempty"`
	Keyword            []CodeableConcept              `json:"keyword,omitempty"`
	Region             []CodeableConcept              `json:"region,omitempty"`
	DescriptionSummary *string                        `json:"descriptionSummary,omitempty"`
	Description        *string                        `json:"description,omitempty"`
	Period             *Period                        `json:"period,omitempty"`
	Site               []Reference                    `json:"site,omitempty"`
	Note               []Annotation                   `json:"note,omitempty"`
	Classifier         []CodeableConcept              `json:"classifier,omitempty"`
	AssociatedParty    []ResearchStudyAssociatedParty `json:"associatedParty,omitempty"`
	ProgressStatus     []ResearchStudyProgressStatus  `json:"progressStatus,omitempty"`
	WhyStopped         *CodeableConcept               `json:"whyStopped,omitempty"`
	Recruitment        *ResearchStudyRecruitment      `json:"recruitment,omitempty"`
	ComparisonGroup    []ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	Objective          []ResearchStudyObjective       `json:"objective,omitempty"`
	OutcomeMeasure     []ResearchStudyOutcomeMeasure  `json:"outcomeMeasure,omitempty"`
	Result             []Reference                    `json:"result,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyLabel struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Value             *string          `json:"value,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyAssociatedParty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Name              *string           `json:"name,omitempty"`
	Role              CodeableConcept   `json:"role"`
	Period            []Period          `json:"period,omitempty"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
	Party             *Reference        `json:"party,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyProgressStatus struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	State             CodeableConcept `json:"state"`
	Actual            *bool           `json:"actual,omitempty"`
	Period            *Period         `json:"period,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyRecruitment struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	TargetNumber      *int        `json:"targetNumber,omitempty"`
	ActualNumber      *int        `json:"actualNumber,omitempty"`
	Eligibility       *Reference  `json:"eligibility,omitempty"`
	ActualGroup       *Reference  `json:"actualGroup,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyComparisonGroup struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	LinkId            *string          `json:"linkId,omitempty"`
	Name              string           `json:"name"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
	IntendedExposure  []Reference      `json:"intendedExposure,omitempty"`
	ObservedGroup     *Reference       `json:"observedGroup,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyObjective struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchStudy
type ResearchStudyOutcomeMeasure struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Name              *string           `json:"name,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Reference         *Reference        `json:"reference,omitempty"`
}

type OtherResearchStudy ResearchStudy

// on convert struct to json, automatically add resourceType=ResearchStudy
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

func (resource *ResearchStudy) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Id", nil)
	}
	return StringInput("ResearchStudy.Id", resource.Id)
}
func (resource *ResearchStudy) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.ImplicitRules", nil)
	}
	return StringInput("ResearchStudy.ImplicitRules", resource.ImplicitRules)
}
func (resource *ResearchStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ResearchStudy.Language", nil, optionsValueSet)
	}
	return CodeSelect("ResearchStudy.Language", resource.Language, optionsValueSet)
}
func (resource *ResearchStudy) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Url", nil)
	}
	return StringInput("ResearchStudy.Url", resource.Url)
}
func (resource *ResearchStudy) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Version", nil)
	}
	return StringInput("ResearchStudy.Version", resource.Version)
}
func (resource *ResearchStudy) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Name", nil)
	}
	return StringInput("ResearchStudy.Name", resource.Name)
}
func (resource *ResearchStudy) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Title", nil)
	}
	return StringInput("ResearchStudy.Title", resource.Title)
}
func (resource *ResearchStudy) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Date", nil)
	}
	return StringInput("ResearchStudy.Date", resource.Date)
}
func (resource *ResearchStudy) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ResearchStudy.Status", nil, optionsValueSet)
	}
	return CodeSelect("ResearchStudy.Status", &resource.Status, optionsValueSet)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.PrimaryPurposeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.PrimaryPurposeType", resource.PrimaryPurposeType, optionsValueSet)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.Phase", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Phase", resource.Phase, optionsValueSet)
}
func (resource *ResearchStudy) T_StudyDesign(numStudyDesign int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StudyDesign) >= numStudyDesign {
		return CodeableConceptSelect("ResearchStudy.StudyDesign["+strconv.Itoa(numStudyDesign)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.StudyDesign["+strconv.Itoa(numStudyDesign)+"]", &resource.StudyDesign[numStudyDesign], optionsValueSet)
}
func (resource *ResearchStudy) T_Condition(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("ResearchStudy.Condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet)
}
func (resource *ResearchStudy) T_Keyword(numKeyword int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Keyword) >= numKeyword {
		return CodeableConceptSelect("ResearchStudy.Keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet)
}
func (resource *ResearchStudy) T_Region(numRegion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Region) >= numRegion {
		return CodeableConceptSelect("ResearchStudy.Region["+strconv.Itoa(numRegion)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Region["+strconv.Itoa(numRegion)+"]", &resource.Region[numRegion], optionsValueSet)
}
func (resource *ResearchStudy) T_DescriptionSummary() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.DescriptionSummary", nil)
	}
	return StringInput("ResearchStudy.DescriptionSummary", resource.DescriptionSummary)
}
func (resource *ResearchStudy) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Description", nil)
	}
	return StringInput("ResearchStudy.Description", resource.Description)
}
func (resource *ResearchStudy) T_Classifier(numClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classifier) >= numClassifier {
		return CodeableConceptSelect("ResearchStudy.Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Classifier[numClassifier], optionsValueSet)
}
func (resource *ResearchStudy) T_WhyStopped(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.WhyStopped", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.WhyStopped", resource.WhyStopped, optionsValueSet)
}
func (resource *ResearchStudy) T_LabelId(numLabel int) templ.Component {

	if resource == nil || len(resource.Label) >= numLabel {
		return StringInput("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Id", nil)
	}
	return StringInput("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Id", resource.Label[numLabel].Id)
}
func (resource *ResearchStudy) T_LabelType(numLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Label) >= numLabel {
		return CodeableConceptSelect("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Type", resource.Label[numLabel].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_LabelValue(numLabel int) templ.Component {

	if resource == nil || len(resource.Label) >= numLabel {
		return StringInput("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Value", nil)
	}
	return StringInput("ResearchStudy.Label["+strconv.Itoa(numLabel)+"].Value", resource.Label[numLabel].Value)
}
func (resource *ResearchStudy) T_AssociatedPartyId(numAssociatedParty int) templ.Component {

	if resource == nil || len(resource.AssociatedParty) >= numAssociatedParty {
		return StringInput("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Id", nil)
	}
	return StringInput("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Id", resource.AssociatedParty[numAssociatedParty].Id)
}
func (resource *ResearchStudy) T_AssociatedPartyName(numAssociatedParty int) templ.Component {

	if resource == nil || len(resource.AssociatedParty) >= numAssociatedParty {
		return StringInput("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Name", nil)
	}
	return StringInput("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Name", resource.AssociatedParty[numAssociatedParty].Name)
}
func (resource *ResearchStudy) T_AssociatedPartyRole(numAssociatedParty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AssociatedParty) >= numAssociatedParty {
		return CodeableConceptSelect("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Role", &resource.AssociatedParty[numAssociatedParty].Role, optionsValueSet)
}
func (resource *ResearchStudy) T_AssociatedPartyClassifier(numAssociatedParty int, numClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AssociatedParty) >= numAssociatedParty || len(resource.AssociatedParty[numAssociatedParty].Classifier) >= numClassifier {
		return CodeableConceptSelect("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.AssociatedParty["+strconv.Itoa(numAssociatedParty)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.AssociatedParty[numAssociatedParty].Classifier[numClassifier], optionsValueSet)
}
func (resource *ResearchStudy) T_ProgressStatusId(numProgressStatus int) templ.Component {

	if resource == nil || len(resource.ProgressStatus) >= numProgressStatus {
		return StringInput("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].Id", nil)
	}
	return StringInput("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].Id", resource.ProgressStatus[numProgressStatus].Id)
}
func (resource *ResearchStudy) T_ProgressStatusState(numProgressStatus int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProgressStatus) >= numProgressStatus {
		return CodeableConceptSelect("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].State", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].State", &resource.ProgressStatus[numProgressStatus].State, optionsValueSet)
}
func (resource *ResearchStudy) T_ProgressStatusActual(numProgressStatus int) templ.Component {

	if resource == nil || len(resource.ProgressStatus) >= numProgressStatus {
		return BoolInput("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].Actual", nil)
	}
	return BoolInput("ResearchStudy.ProgressStatus["+strconv.Itoa(numProgressStatus)+"].Actual", resource.ProgressStatus[numProgressStatus].Actual)
}
func (resource *ResearchStudy) T_RecruitmentId() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Recruitment.Id", nil)
	}
	return StringInput("ResearchStudy.Recruitment.Id", resource.Recruitment.Id)
}
func (resource *ResearchStudy) T_RecruitmentTargetNumber() templ.Component {

	if resource == nil {
		return IntInput("ResearchStudy.Recruitment.TargetNumber", nil)
	}
	return IntInput("ResearchStudy.Recruitment.TargetNumber", resource.Recruitment.TargetNumber)
}
func (resource *ResearchStudy) T_RecruitmentActualNumber() templ.Component {

	if resource == nil {
		return IntInput("ResearchStudy.Recruitment.ActualNumber", nil)
	}
	return IntInput("ResearchStudy.Recruitment.ActualNumber", resource.Recruitment.ActualNumber)
}
func (resource *ResearchStudy) T_ComparisonGroupId(numComparisonGroup int) templ.Component {

	if resource == nil || len(resource.ComparisonGroup) >= numComparisonGroup {
		return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Id", nil)
	}
	return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Id", resource.ComparisonGroup[numComparisonGroup].Id)
}
func (resource *ResearchStudy) T_ComparisonGroupLinkId(numComparisonGroup int) templ.Component {

	if resource == nil || len(resource.ComparisonGroup) >= numComparisonGroup {
		return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].LinkId", nil)
	}
	return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].LinkId", resource.ComparisonGroup[numComparisonGroup].LinkId)
}
func (resource *ResearchStudy) T_ComparisonGroupName(numComparisonGroup int) templ.Component {

	if resource == nil || len(resource.ComparisonGroup) >= numComparisonGroup {
		return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Name", nil)
	}
	return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Name", &resource.ComparisonGroup[numComparisonGroup].Name)
}
func (resource *ResearchStudy) T_ComparisonGroupType(numComparisonGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ComparisonGroup) >= numComparisonGroup {
		return CodeableConceptSelect("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Type", resource.ComparisonGroup[numComparisonGroup].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_ComparisonGroupDescription(numComparisonGroup int) templ.Component {

	if resource == nil || len(resource.ComparisonGroup) >= numComparisonGroup {
		return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Description", nil)
	}
	return StringInput("ResearchStudy.ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"].Description", resource.ComparisonGroup[numComparisonGroup].Description)
}
func (resource *ResearchStudy) T_ObjectiveId(numObjective int) templ.Component {

	if resource == nil || len(resource.Objective) >= numObjective {
		return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Id", nil)
	}
	return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Id", resource.Objective[numObjective].Id)
}
func (resource *ResearchStudy) T_ObjectiveName(numObjective int) templ.Component {

	if resource == nil || len(resource.Objective) >= numObjective {
		return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Name", nil)
	}
	return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Name", resource.Objective[numObjective].Name)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Objective) >= numObjective {
		return CodeableConceptSelect("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Type", resource.Objective[numObjective].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_ObjectiveDescription(numObjective int) templ.Component {

	if resource == nil || len(resource.Objective) >= numObjective {
		return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Description", nil)
	}
	return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Description", resource.Objective[numObjective].Description)
}
func (resource *ResearchStudy) T_OutcomeMeasureId(numOutcomeMeasure int) templ.Component {

	if resource == nil || len(resource.OutcomeMeasure) >= numOutcomeMeasure {
		return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Id", nil)
	}
	return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Id", resource.OutcomeMeasure[numOutcomeMeasure].Id)
}
func (resource *ResearchStudy) T_OutcomeMeasureName(numOutcomeMeasure int) templ.Component {

	if resource == nil || len(resource.OutcomeMeasure) >= numOutcomeMeasure {
		return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Name", nil)
	}
	return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Name", resource.OutcomeMeasure[numOutcomeMeasure].Name)
}
func (resource *ResearchStudy) T_OutcomeMeasureType(numOutcomeMeasure int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OutcomeMeasure) >= numOutcomeMeasure || len(resource.OutcomeMeasure[numOutcomeMeasure].Type) >= numType {
		return CodeableConceptSelect("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Type["+strconv.Itoa(numType)+"]", &resource.OutcomeMeasure[numOutcomeMeasure].Type[numType], optionsValueSet)
}
func (resource *ResearchStudy) T_OutcomeMeasureDescription(numOutcomeMeasure int) templ.Component {

	if resource == nil || len(resource.OutcomeMeasure) >= numOutcomeMeasure {
		return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Description", nil)
	}
	return StringInput("ResearchStudy.OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"].Description", resource.OutcomeMeasure[numOutcomeMeasure].Description)
}
