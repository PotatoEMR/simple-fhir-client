package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *ResearchStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ResearchStudy) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("primaryPurposeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("primaryPurposeType", resource.PrimaryPurposeType, optionsValueSet)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("phase", nil, optionsValueSet)
	}
	return CodeableConceptSelect("phase", resource.Phase, optionsValueSet)
}
func (resource *ResearchStudy) T_StudyDesign(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("studyDesign", nil, optionsValueSet)
	}
	return CodeableConceptSelect("studyDesign", &resource.StudyDesign[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Condition(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("condition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("condition", &resource.Condition[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Keyword(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("keyword", nil, optionsValueSet)
	}
	return CodeableConceptSelect("keyword", &resource.Keyword[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Region(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("region", nil, optionsValueSet)
	}
	return CodeableConceptSelect("region", &resource.Region[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Classifier(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.Classifier[0], optionsValueSet)
}
func (resource *ResearchStudy) T_WhyStopped(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("whyStopped", nil, optionsValueSet)
	}
	return CodeableConceptSelect("whyStopped", resource.WhyStopped, optionsValueSet)
}
func (resource *ResearchStudy) T_LabelType(numLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Label) >= numLabel {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Label[numLabel].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_AssociatedPartyRole(numAssociatedParty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AssociatedParty) >= numAssociatedParty {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.AssociatedParty[numAssociatedParty].Role, optionsValueSet)
}
func (resource *ResearchStudy) T_AssociatedPartyClassifier(numAssociatedParty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AssociatedParty) >= numAssociatedParty {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.AssociatedParty[numAssociatedParty].Classifier[0], optionsValueSet)
}
func (resource *ResearchStudy) T_ProgressStatusState(numProgressStatus int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProgressStatus) >= numProgressStatus {
		return CodeableConceptSelect("state", nil, optionsValueSet)
	}
	return CodeableConceptSelect("state", &resource.ProgressStatus[numProgressStatus].State, optionsValueSet)
}
func (resource *ResearchStudy) T_ComparisonGroupType(numComparisonGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ComparisonGroup) >= numComparisonGroup {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.ComparisonGroup[numComparisonGroup].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Objective) >= numObjective {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Objective[numObjective].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_OutcomeMeasureType(numOutcomeMeasure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.OutcomeMeasure) >= numOutcomeMeasure {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.OutcomeMeasure[numOutcomeMeasure].Type[0], optionsValueSet)
}
