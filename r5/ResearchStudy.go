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
	Date               *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ResearchStudy) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ResearchStudy/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ResearchStudy"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ResearchStudy) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *ResearchStudy) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ResearchStudy) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *ResearchStudy) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ResearchStudy) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ResearchStudy) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PrimaryPurposeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PrimaryPurposeType", resource.PrimaryPurposeType, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Phase", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Phase", resource.Phase, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_StudyDesign(numStudyDesign int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStudyDesign >= len(resource.StudyDesign) {
		return CodeableConceptSelect("StudyDesign["+strconv.Itoa(numStudyDesign)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StudyDesign["+strconv.Itoa(numStudyDesign)+"]", &resource.StudyDesign[numStudyDesign], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Keyword(numKeyword int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numKeyword >= len(resource.Keyword) {
		return CodeableConceptSelect("Keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Region(numRegion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRegion >= len(resource.Region) {
		return CodeableConceptSelect("Region["+strconv.Itoa(numRegion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Region["+strconv.Itoa(numRegion)+"]", &resource.Region[numRegion], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_DescriptionSummary(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DescriptionSummary", nil, htmlAttrs)
	}
	return StringInput("DescriptionSummary", resource.DescriptionSummary, htmlAttrs)
}
func (resource *ResearchStudy) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ResearchStudy) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ResearchStudy) T_Classifier(numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassifier >= len(resource.Classifier) {
		return CodeableConceptSelect("Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_WhyStopped(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("WhyStopped", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("WhyStopped", resource.WhyStopped, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_LabelType(numLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLabel >= len(resource.Label) {
		return CodeableConceptSelect("Label["+strconv.Itoa(numLabel)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Label["+strconv.Itoa(numLabel)+"]Type", resource.Label[numLabel].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_LabelValue(numLabel int, htmlAttrs string) templ.Component {
	if resource == nil || numLabel >= len(resource.Label) {
		return StringInput("Label["+strconv.Itoa(numLabel)+"]Value", nil, htmlAttrs)
	}
	return StringInput("Label["+strconv.Itoa(numLabel)+"]Value", resource.Label[numLabel].Value, htmlAttrs)
}
func (resource *ResearchStudy) T_AssociatedPartyName(numAssociatedParty int, htmlAttrs string) templ.Component {
	if resource == nil || numAssociatedParty >= len(resource.AssociatedParty) {
		return StringInput("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Name", nil, htmlAttrs)
	}
	return StringInput("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Name", resource.AssociatedParty[numAssociatedParty].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_AssociatedPartyRole(numAssociatedParty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAssociatedParty >= len(resource.AssociatedParty) {
		return CodeableConceptSelect("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Role", &resource.AssociatedParty[numAssociatedParty].Role, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_AssociatedPartyClassifier(numAssociatedParty int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAssociatedParty >= len(resource.AssociatedParty) || numClassifier >= len(resource.AssociatedParty[numAssociatedParty].Classifier) {
		return CodeableConceptSelect("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AssociatedParty["+strconv.Itoa(numAssociatedParty)+"]Classifier["+strconv.Itoa(numClassifier)+"]", &resource.AssociatedParty[numAssociatedParty].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ProgressStatusState(numProgressStatus int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProgressStatus >= len(resource.ProgressStatus) {
		return CodeableConceptSelect("ProgressStatus["+strconv.Itoa(numProgressStatus)+"]State", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProgressStatus["+strconv.Itoa(numProgressStatus)+"]State", &resource.ProgressStatus[numProgressStatus].State, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ProgressStatusActual(numProgressStatus int, htmlAttrs string) templ.Component {
	if resource == nil || numProgressStatus >= len(resource.ProgressStatus) {
		return BoolInput("ProgressStatus["+strconv.Itoa(numProgressStatus)+"]Actual", nil, htmlAttrs)
	}
	return BoolInput("ProgressStatus["+strconv.Itoa(numProgressStatus)+"]Actual", resource.ProgressStatus[numProgressStatus].Actual, htmlAttrs)
}
func (resource *ResearchStudy) T_RecruitmentTargetNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RecruitmentTargetNumber", nil, htmlAttrs)
	}
	return IntInput("RecruitmentTargetNumber", resource.Recruitment.TargetNumber, htmlAttrs)
}
func (resource *ResearchStudy) T_RecruitmentActualNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RecruitmentActualNumber", nil, htmlAttrs)
	}
	return IntInput("RecruitmentActualNumber", resource.Recruitment.ActualNumber, htmlAttrs)
}
func (resource *ResearchStudy) T_ComparisonGroupLinkId(numComparisonGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numComparisonGroup >= len(resource.ComparisonGroup) {
		return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]LinkId", nil, htmlAttrs)
	}
	return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]LinkId", resource.ComparisonGroup[numComparisonGroup].LinkId, htmlAttrs)
}
func (resource *ResearchStudy) T_ComparisonGroupName(numComparisonGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numComparisonGroup >= len(resource.ComparisonGroup) {
		return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Name", nil, htmlAttrs)
	}
	return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Name", &resource.ComparisonGroup[numComparisonGroup].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ComparisonGroupType(numComparisonGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComparisonGroup >= len(resource.ComparisonGroup) {
		return CodeableConceptSelect("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Type", resource.ComparisonGroup[numComparisonGroup].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ComparisonGroupDescription(numComparisonGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numComparisonGroup >= len(resource.ComparisonGroup) {
		return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Description", nil, htmlAttrs)
	}
	return StringInput("ComparisonGroup["+strconv.Itoa(numComparisonGroup)+"]Description", resource.ComparisonGroup[numComparisonGroup].Description, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveName(numObjective int, htmlAttrs string) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return StringInput("Objective["+strconv.Itoa(numObjective)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Objective["+strconv.Itoa(numObjective)+"]Name", resource.Objective[numObjective].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return CodeableConceptSelect("Objective["+strconv.Itoa(numObjective)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Objective["+strconv.Itoa(numObjective)+"]Type", resource.Objective[numObjective].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveDescription(numObjective int, htmlAttrs string) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return StringInput("Objective["+strconv.Itoa(numObjective)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Objective["+strconv.Itoa(numObjective)+"]Description", resource.Objective[numObjective].Description, htmlAttrs)
}
func (resource *ResearchStudy) T_OutcomeMeasureName(numOutcomeMeasure int, htmlAttrs string) templ.Component {
	if resource == nil || numOutcomeMeasure >= len(resource.OutcomeMeasure) {
		return StringInput("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Name", nil, htmlAttrs)
	}
	return StringInput("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Name", resource.OutcomeMeasure[numOutcomeMeasure].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_OutcomeMeasureType(numOutcomeMeasure int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutcomeMeasure >= len(resource.OutcomeMeasure) || numType >= len(resource.OutcomeMeasure[numOutcomeMeasure].Type) {
		return CodeableConceptSelect("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Type["+strconv.Itoa(numType)+"]", &resource.OutcomeMeasure[numOutcomeMeasure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_OutcomeMeasureDescription(numOutcomeMeasure int, htmlAttrs string) templ.Component {
	if resource == nil || numOutcomeMeasure >= len(resource.OutcomeMeasure) {
		return StringInput("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Description", nil, htmlAttrs)
	}
	return StringInput("OutcomeMeasure["+strconv.Itoa(numOutcomeMeasure)+"]Description", resource.OutcomeMeasure[numOutcomeMeasure].Description, htmlAttrs)
}
