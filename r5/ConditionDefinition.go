package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinition struct {
	Id                     *string                            `json:"id,omitempty"`
	Meta                   *Meta                              `json:"meta,omitempty"`
	ImplicitRules          *string                            `json:"implicitRules,omitempty"`
	Language               *string                            `json:"language,omitempty"`
	Text                   *Narrative                         `json:"text,omitempty"`
	Contained              []Resource                         `json:"contained,omitempty"`
	Extension              []Extension                        `json:"extension,omitempty"`
	ModifierExtension      []Extension                        `json:"modifierExtension,omitempty"`
	Url                    *string                            `json:"url,omitempty"`
	Identifier             []Identifier                       `json:"identifier,omitempty"`
	Version                *string                            `json:"version,omitempty"`
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                            `json:"name,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	Subtitle               *string                            `json:"subtitle,omitempty"`
	Status                 string                             `json:"status"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	Date                   *string                            `json:"date,omitempty"`
	Publisher              *string                            `json:"publisher,omitempty"`
	Contact                []ContactDetail                    `json:"contact,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	UseContext             []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Code                   CodeableConcept                    `json:"code"`
	Severity               *CodeableConcept                   `json:"severity,omitempty"`
	BodySite               *CodeableConcept                   `json:"bodySite,omitempty"`
	Stage                  *CodeableConcept                   `json:"stage,omitempty"`
	HasSeverity            *bool                              `json:"hasSeverity,omitempty"`
	HasBodySite            *bool                              `json:"hasBodySite,omitempty"`
	HasStage               *bool                              `json:"hasStage,omitempty"`
	Definition             []string                           `json:"definition,omitempty"`
	Observation            []ConditionDefinitionObservation   `json:"observation,omitempty"`
	Medication             []ConditionDefinitionMedication    `json:"medication,omitempty"`
	Precondition           []ConditionDefinitionPrecondition  `json:"precondition,omitempty"`
	Team                   []Reference                        `json:"team,omitempty"`
	Questionnaire          []ConditionDefinitionQuestionnaire `json:"questionnaire,omitempty"`
	Plan                   []ConditionDefinitionPlan          `json:"plan,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionObservation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionMedication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionPrecondition struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 string           `json:"type"`
	Code                 CodeableConcept  `json:"code"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionQuestionnaire struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Purpose           string      `json:"purpose"`
	Reference         Reference   `json:"reference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionPlan struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Reference         Reference        `json:"reference"`
}

type OtherConditionDefinition ConditionDefinition

// on convert struct to json, automatically add resourceType=ConditionDefinition
func (r ConditionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConditionDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherConditionDefinition: OtherConditionDefinition(r),
		ResourceType:             "ConditionDefinition",
	})
}

func (resource *ConditionDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Id", nil)
	}
	return StringInput("ConditionDefinition.Id", resource.Id)
}
func (resource *ConditionDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.ImplicitRules", nil)
	}
	return StringInput("ConditionDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ConditionDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ConditionDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ConditionDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ConditionDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Url", nil)
	}
	return StringInput("ConditionDefinition.Url", resource.Url)
}
func (resource *ConditionDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Version", nil)
	}
	return StringInput("ConditionDefinition.Version", resource.Version)
}
func (resource *ConditionDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Name", nil)
	}
	return StringInput("ConditionDefinition.Name", resource.Name)
}
func (resource *ConditionDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Title", nil)
	}
	return StringInput("ConditionDefinition.Title", resource.Title)
}
func (resource *ConditionDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Subtitle", nil)
	}
	return StringInput("ConditionDefinition.Subtitle", resource.Subtitle)
}
func (resource *ConditionDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ConditionDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ConditionDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ConditionDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ConditionDefinition.Experimental", nil)
	}
	return BoolInput("ConditionDefinition.Experimental", resource.Experimental)
}
func (resource *ConditionDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Date", nil)
	}
	return StringInput("ConditionDefinition.Date", resource.Date)
}
func (resource *ConditionDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Publisher", nil)
	}
	return StringInput("ConditionDefinition.Publisher", resource.Publisher)
}
func (resource *ConditionDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ConditionDefinition.Description", nil)
	}
	return StringInput("ConditionDefinition.Description", resource.Description)
}
func (resource *ConditionDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ConditionDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ConditionDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Code", &resource.Code, optionsValueSet)
}
func (resource *ConditionDefinition) T_Severity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Severity", resource.Severity, optionsValueSet)
}
func (resource *ConditionDefinition) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.BodySite", resource.BodySite, optionsValueSet)
}
func (resource *ConditionDefinition) T_Stage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Stage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Stage", resource.Stage, optionsValueSet)
}
func (resource *ConditionDefinition) T_HasSeverity() templ.Component {

	if resource == nil {
		return BoolInput("ConditionDefinition.HasSeverity", nil)
	}
	return BoolInput("ConditionDefinition.HasSeverity", resource.HasSeverity)
}
func (resource *ConditionDefinition) T_HasBodySite() templ.Component {

	if resource == nil {
		return BoolInput("ConditionDefinition.HasBodySite", nil)
	}
	return BoolInput("ConditionDefinition.HasBodySite", resource.HasBodySite)
}
func (resource *ConditionDefinition) T_HasStage() templ.Component {

	if resource == nil {
		return BoolInput("ConditionDefinition.HasStage", nil)
	}
	return BoolInput("ConditionDefinition.HasStage", resource.HasStage)
}
func (resource *ConditionDefinition) T_Definition(numDefinition int) templ.Component {

	if resource == nil || len(resource.Definition) >= numDefinition {
		return StringInput("ConditionDefinition.Definition["+strconv.Itoa(numDefinition)+"]", nil)
	}
	return StringInput("ConditionDefinition.Definition["+strconv.Itoa(numDefinition)+"]", &resource.Definition[numDefinition])
}
func (resource *ConditionDefinition) T_ObservationId(numObservation int) templ.Component {

	if resource == nil || len(resource.Observation) >= numObservation {
		return StringInput("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Id", nil)
	}
	return StringInput("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Id", resource.Observation[numObservation].Id)
}
func (resource *ConditionDefinition) T_ObservationCategory(numObservation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Observation) >= numObservation {
		return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Category", resource.Observation[numObservation].Category, optionsValueSet)
}
func (resource *ConditionDefinition) T_ObservationCode(numObservation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Observation) >= numObservation {
		return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Code", resource.Observation[numObservation].Code, optionsValueSet)
}
func (resource *ConditionDefinition) T_MedicationId(numMedication int) templ.Component {

	if resource == nil || len(resource.Medication) >= numMedication {
		return StringInput("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Id", nil)
	}
	return StringInput("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Id", resource.Medication[numMedication].Id)
}
func (resource *ConditionDefinition) T_MedicationCategory(numMedication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Medication) >= numMedication {
		return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Category", resource.Medication[numMedication].Category, optionsValueSet)
}
func (resource *ConditionDefinition) T_MedicationCode(numMedication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Medication) >= numMedication {
		return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Code", resource.Medication[numMedication].Code, optionsValueSet)
}
func (resource *ConditionDefinition) T_PreconditionId(numPrecondition int) templ.Component {

	if resource == nil || len(resource.Precondition) >= numPrecondition {
		return StringInput("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Id", nil)
	}
	return StringInput("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Id", resource.Precondition[numPrecondition].Id)
}
func (resource *ConditionDefinition) T_PreconditionType(numPrecondition int) templ.Component {
	optionsValueSet := VSCondition_precondition_type

	if resource == nil || len(resource.Precondition) >= numPrecondition {
		return CodeSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Type", &resource.Precondition[numPrecondition].Type, optionsValueSet)
}
func (resource *ConditionDefinition) T_PreconditionCode(numPrecondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Precondition) >= numPrecondition {
		return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Code", &resource.Precondition[numPrecondition].Code, optionsValueSet)
}
func (resource *ConditionDefinition) T_QuestionnaireId(numQuestionnaire int) templ.Component {

	if resource == nil || len(resource.Questionnaire) >= numQuestionnaire {
		return StringInput("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Id", nil)
	}
	return StringInput("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Id", resource.Questionnaire[numQuestionnaire].Id)
}
func (resource *ConditionDefinition) T_QuestionnairePurpose(numQuestionnaire int) templ.Component {
	optionsValueSet := VSCondition_questionnaire_purpose

	if resource == nil || len(resource.Questionnaire) >= numQuestionnaire {
		return CodeSelect("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Purpose", nil, optionsValueSet)
	}
	return CodeSelect("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Purpose", &resource.Questionnaire[numQuestionnaire].Purpose, optionsValueSet)
}
func (resource *ConditionDefinition) T_PlanId(numPlan int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan {
		return StringInput("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Id", nil)
	}
	return StringInput("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Id", resource.Plan[numPlan].Id)
}
func (resource *ConditionDefinition) T_PlanRole(numPlan int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan {
		return CodeableConceptSelect("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Role", resource.Plan[numPlan].Role, optionsValueSet)
}
