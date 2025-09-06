package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                     *string                `json:"id,omitempty"`
	Meta                   *Meta                  `json:"meta,omitempty"`
	ImplicitRules          *string                `json:"implicitRules,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Text                   *Narrative             `json:"text,omitempty"`
	Contained              []Resource             `json:"contained,omitempty"`
	Extension              []Extension            `json:"extension,omitempty"`
	ModifierExtension      []Extension            `json:"modifierExtension,omitempty"`
	Url                    *string                `json:"url,omitempty"`
	Identifier             []Identifier           `json:"identifier,omitempty"`
	Version                *string                `json:"version,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Title                  *string                `json:"title,omitempty"`
	Subtitle               *string                `json:"subtitle,omitempty"`
	Type                   *CodeableConcept       `json:"type,omitempty"`
	Status                 string                 `json:"status"`
	Experimental           *bool                  `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept       `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference             `json:"subjectReference,omitempty"`
	Date                   *string                `json:"date,omitempty"`
	Publisher              *string                `json:"publisher,omitempty"`
	Contact                []ContactDetail        `json:"contact,omitempty"`
	Description            *string                `json:"description,omitempty"`
	UseContext             []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                *string                `json:"purpose,omitempty"`
	Usage                  *string                `json:"usage,omitempty"`
	Copyright              *string                `json:"copyright,omitempty"`
	ApprovalDate           *string                `json:"approvalDate,omitempty"`
	LastReviewDate         *string                `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept      `json:"topic,omitempty"`
	Author                 []ContactDetail        `json:"author,omitempty"`
	Editor                 []ContactDetail        `json:"editor,omitempty"`
	Reviewer               []ContactDetail        `json:"reviewer,omitempty"`
	Endorser               []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Library                []string               `json:"library,omitempty"`
	Goal                   []PlanDefinitionGoal   `json:"goal,omitempty"`
	Action                 []PlanDefinitionAction `json:"action,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionGoal struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept           `json:"category,omitempty"`
	Description       CodeableConcept            `json:"description"`
	Priority          *CodeableConcept           `json:"priority,omitempty"`
	Start             *CodeableConcept           `json:"start,omitempty"`
	Addresses         []CodeableConcept          `json:"addresses,omitempty"`
	Documentation     []RelatedArtifact          `json:"documentation,omitempty"`
	Target            []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionGoalTarget struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range           `json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
	Due                   *Duration        `json:"due,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionAction struct {
	Id                     *string                             `json:"id,omitempty"`
	Extension              []Extension                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                         `json:"modifierExtension,omitempty"`
	Prefix                 *string                             `json:"prefix,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	TextEquivalent         *string                             `json:"textEquivalent,omitempty"`
	Priority               *string                             `json:"priority,omitempty"`
	Code                   []CodeableConcept                   `json:"code,omitempty"`
	Reason                 []CodeableConcept                   `json:"reason,omitempty"`
	Documentation          []RelatedArtifact                   `json:"documentation,omitempty"`
	GoalId                 []string                            `json:"goalId,omitempty"`
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                          `json:"subjectReference,omitempty"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []DataRequirement                   `json:"input,omitempty"`
	Output                 []DataRequirement                   `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime         *string                             `json:"timingDateTime,omitempty"`
	TimingAge              *Age                                `json:"timingAge,omitempty"`
	TimingPeriod           *Period                             `json:"timingPeriod,omitempty"`
	TimingDuration         *Duration                           `json:"timingDuration,omitempty"`
	TimingRange            *Range                              `json:"timingRange,omitempty"`
	TimingTiming           *Timing                             `json:"timingTiming,omitempty"`
	Participant            []PlanDefinitionActionParticipant   `json:"participant,omitempty"`
	Type                   *CodeableConcept                    `json:"type,omitempty"`
	GroupingBehavior       *string                             `json:"groupingBehavior,omitempty"`
	SelectionBehavior      *string                             `json:"selectionBehavior,omitempty"`
	RequiredBehavior       *string                             `json:"requiredBehavior,omitempty"`
	PrecheckBehavior       *string                             `json:"precheckBehavior,omitempty"`
	CardinalityBehavior    *string                             `json:"cardinalityBehavior,omitempty"`
	DefinitionCanonical    *string                             `json:"definitionCanonical,omitempty"`
	DefinitionUri          *string                             `json:"definitionUri,omitempty"`
	Transform              *string                             `json:"transform,omitempty"`
	DynamicValue           []PlanDefinitionActionDynamicValue  `json:"dynamicValue,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionActionRelatedAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ActionId          string      `json:"actionId"`
	Relationship      string      `json:"relationship"`
	OffsetDuration    *Duration   `json:"offsetDuration,omitempty"`
	OffsetRange       *Range      `json:"offsetRange,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionActionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              string           `json:"type"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PlanDefinition
type PlanDefinitionActionDynamicValue struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Path              *string     `json:"path,omitempty"`
	Expression        *Expression `json:"expression,omitempty"`
}

type OtherPlanDefinition PlanDefinition

// on convert struct to json, automatically add resourceType=PlanDefinition
func (r PlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPlanDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPlanDefinition: OtherPlanDefinition(r),
		ResourceType:        "PlanDefinition",
	})
}

func (resource *PlanDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Id", nil)
	}
	return StringInput("PlanDefinition.Id", resource.Id)
}
func (resource *PlanDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.ImplicitRules", nil)
	}
	return StringInput("PlanDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *PlanDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PlanDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *PlanDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Url", nil)
	}
	return StringInput("PlanDefinition.Url", resource.Url)
}
func (resource *PlanDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Version", nil)
	}
	return StringInput("PlanDefinition.Version", resource.Version)
}
func (resource *PlanDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Name", nil)
	}
	return StringInput("PlanDefinition.Name", resource.Name)
}
func (resource *PlanDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Title", nil)
	}
	return StringInput("PlanDefinition.Title", resource.Title)
}
func (resource *PlanDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Subtitle", nil)
	}
	return StringInput("PlanDefinition.Subtitle", resource.Subtitle)
}
func (resource *PlanDefinition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PlanDefinition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Type", resource.Type, optionsValueSet)
}
func (resource *PlanDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("PlanDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *PlanDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("PlanDefinition.Experimental", nil)
	}
	return BoolInput("PlanDefinition.Experimental", resource.Experimental)
}
func (resource *PlanDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Date", nil)
	}
	return StringInput("PlanDefinition.Date", resource.Date)
}
func (resource *PlanDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Publisher", nil)
	}
	return StringInput("PlanDefinition.Publisher", resource.Publisher)
}
func (resource *PlanDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Description", nil)
	}
	return StringInput("PlanDefinition.Description", resource.Description)
}
func (resource *PlanDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("PlanDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *PlanDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Purpose", nil)
	}
	return StringInput("PlanDefinition.Purpose", resource.Purpose)
}
func (resource *PlanDefinition) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Usage", nil)
	}
	return StringInput("PlanDefinition.Usage", resource.Usage)
}
func (resource *PlanDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.Copyright", nil)
	}
	return StringInput("PlanDefinition.Copyright", resource.Copyright)
}
func (resource *PlanDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.ApprovalDate", nil)
	}
	return StringInput("PlanDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *PlanDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("PlanDefinition.LastReviewDate", nil)
	}
	return StringInput("PlanDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *PlanDefinition) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("PlanDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *PlanDefinition) T_Library(numLibrary int) templ.Component {

	if resource == nil || len(resource.Library) >= numLibrary {
		return StringInput("PlanDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("PlanDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary])
}
func (resource *PlanDefinition) T_GoalId(numGoal int) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal {
		return StringInput("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Id", resource.Goal[numGoal].Id)
}
func (resource *PlanDefinition) T_GoalCategory(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Category", resource.Goal[numGoal].Category, optionsValueSet)
}
func (resource *PlanDefinition) T_GoalDescription(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Description", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Description", &resource.Goal[numGoal].Description, optionsValueSet)
}
func (resource *PlanDefinition) T_GoalPriority(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Priority", resource.Goal[numGoal].Priority, optionsValueSet)
}
func (resource *PlanDefinition) T_GoalStart(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Start", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Start", resource.Goal[numGoal].Start, optionsValueSet)
}
func (resource *PlanDefinition) T_GoalAddresses(numGoal int, numAddresses int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal || len(resource.Goal[numGoal].Addresses) >= numAddresses {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Addresses["+strconv.Itoa(numAddresses)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Addresses["+strconv.Itoa(numAddresses)+"]", &resource.Goal[numGoal].Addresses[numAddresses], optionsValueSet)
}
func (resource *PlanDefinition) T_GoalTargetId(numGoal int, numTarget int) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal || len(resource.Goal[numGoal].Target) >= numTarget {
		return StringInput("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Id", resource.Goal[numGoal].Target[numTarget].Id)
}
func (resource *PlanDefinition) T_GoalTargetMeasure(numGoal int, numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Goal) >= numGoal || len(resource.Goal[numGoal].Target) >= numTarget {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Measure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Measure", resource.Goal[numGoal].Target[numTarget].Measure, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Id", resource.Action[numAction].Id)
}
func (resource *PlanDefinition) T_ActionPrefix(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Prefix", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Prefix", resource.Action[numAction].Prefix)
}
func (resource *PlanDefinition) T_ActionTitle(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Title", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Title", resource.Action[numAction].Title)
}
func (resource *PlanDefinition) T_ActionDescription(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Description", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Description", resource.Action[numAction].Description)
}
func (resource *PlanDefinition) T_ActionTextEquivalent(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TextEquivalent", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TextEquivalent", resource.Action[numAction].TextEquivalent)
}
func (resource *PlanDefinition) T_ActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Priority", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Priority", resource.Action[numAction].Priority, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Code) >= numCode {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet)
}
func (resource *PlanDefinition) T_ActionReason(numAction int, numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Reason) >= numReason {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", &resource.Action[numAction].Reason[numReason], optionsValueSet)
}
func (resource *PlanDefinition) T_ActionGoalId(numAction int, numGoalId int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].GoalId) >= numGoalId {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GoalId["+strconv.Itoa(numGoalId)+"]", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GoalId["+strconv.Itoa(numGoalId)+"]", &resource.Action[numAction].GoalId[numGoalId])
}
func (resource *PlanDefinition) T_ActionType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Type", resource.Action[numAction].Type, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionTransform(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Transform", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Transform", resource.Action[numAction].Transform)
}
func (resource *PlanDefinition) T_ActionConditionId(numAction int, numCondition int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", resource.Action[numAction].Condition[numCondition].Id)
}
func (resource *PlanDefinition) T_ActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionRelatedActionId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", resource.Action[numAction].RelatedAction[numRelatedAction].Id)
}
func (resource *PlanDefinition) T_ActionRelatedActionActionId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", &resource.Action[numAction].RelatedAction[numRelatedAction].ActionId)
}
func (resource *PlanDefinition) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionParticipantId(numAction int, numParticipant int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Action[numAction].Participant[numParticipant].Id)
}
func (resource *PlanDefinition) T_ActionParticipantType(numAction int, numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", &resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet)
}
func (resource *PlanDefinition) T_ActionDynamicValueId(numAction int, numDynamicValue int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].DynamicValue) >= numDynamicValue {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", resource.Action[numAction].DynamicValue[numDynamicValue].Id)
}
func (resource *PlanDefinition) T_ActionDynamicValuePath(numAction int, numDynamicValue int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].DynamicValue) >= numDynamicValue {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", resource.Action[numAction].DynamicValue[numDynamicValue].Path)
}
