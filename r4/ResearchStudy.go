package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []Resource               `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	Title                 *string                  `json:"title,omitempty"`
	Protocol              []Reference              `json:"protocol,omitempty"`
	PartOf                []Reference              `json:"partOf,omitempty"`
	Status                string                   `json:"status"`
	PrimaryPurposeType    *CodeableConcept         `json:"primaryPurposeType,omitempty"`
	Phase                 *CodeableConcept         `json:"phase,omitempty"`
	Category              []CodeableConcept        `json:"category,omitempty"`
	Focus                 []CodeableConcept        `json:"focus,omitempty"`
	Condition             []CodeableConcept        `json:"condition,omitempty"`
	Contact               []ContactDetail          `json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept        `json:"keyword,omitempty"`
	Location              []CodeableConcept        `json:"location,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Enrollment            []Reference              `json:"enrollment,omitempty"`
	Period                *Period                  `json:"period,omitempty"`
	Sponsor               *Reference               `json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference               `json:"principalInvestigator,omitempty"`
	Site                  []Reference              `json:"site,omitempty"`
	ReasonStopped         *CodeableConcept         `json:"reasonStopped,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	Arm                   []ResearchStudyArm       `json:"arm,omitempty"`
	Objective             []ResearchStudyObjective `json:"objective,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudyArm struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              string           `json:"name"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudyObjective struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
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
func (resource *ResearchStudy) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Title", nil)
	}
	return StringInput("ResearchStudy.Title", resource.Title)
}
func (resource *ResearchStudy) T_Status() templ.Component {
	optionsValueSet := VSResearch_study_status

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
func (resource *ResearchStudy) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ResearchStudy.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ResearchStudy) T_Focus(numFocus int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Focus) >= numFocus {
		return CodeableConceptSelect("ResearchStudy.Focus["+strconv.Itoa(numFocus)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], optionsValueSet)
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
func (resource *ResearchStudy) T_Location(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return CodeableConceptSelect("ResearchStudy.Location["+strconv.Itoa(numLocation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], optionsValueSet)
}
func (resource *ResearchStudy) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ResearchStudy.Description", nil)
	}
	return StringInput("ResearchStudy.Description", resource.Description)
}
func (resource *ResearchStudy) T_ReasonStopped(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.ReasonStopped", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.ReasonStopped", resource.ReasonStopped, optionsValueSet)
}
func (resource *ResearchStudy) T_ArmId(numArm int) templ.Component {

	if resource == nil || len(resource.Arm) >= numArm {
		return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Id", nil)
	}
	return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Id", resource.Arm[numArm].Id)
}
func (resource *ResearchStudy) T_ArmName(numArm int) templ.Component {

	if resource == nil || len(resource.Arm) >= numArm {
		return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Name", nil)
	}
	return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Name", &resource.Arm[numArm].Name)
}
func (resource *ResearchStudy) T_ArmType(numArm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Arm) >= numArm {
		return CodeableConceptSelect("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Type", resource.Arm[numArm].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_ArmDescription(numArm int) templ.Component {

	if resource == nil || len(resource.Arm) >= numArm {
		return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Description", nil)
	}
	return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Description", resource.Arm[numArm].Description)
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
