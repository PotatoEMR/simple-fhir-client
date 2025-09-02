package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                      *string                `json:"id,omitempty"`
	Meta                    *Meta                  `json:"meta,omitempty"`
	ImplicitRules           *string                `json:"implicitRules,omitempty"`
	Language                *string                `json:"language,omitempty"`
	Text                    *Narrative             `json:"text,omitempty"`
	Contained               []Resource             `json:"contained,omitempty"`
	Extension               []Extension            `json:"extension,omitempty"`
	ModifierExtension       []Extension            `json:"modifierExtension,omitempty"`
	Url                     *string                `json:"url,omitempty"`
	Identifier              []Identifier           `json:"identifier,omitempty"`
	Version                 *string                `json:"version,omitempty"`
	VersionAlgorithmString  *string                `json:"versionAlgorithmString"`
	VersionAlgorithmCoding  *Coding                `json:"versionAlgorithmCoding"`
	Name                    *string                `json:"name,omitempty"`
	Title                   *string                `json:"title,omitempty"`
	Subtitle                *string                `json:"subtitle,omitempty"`
	Type                    *CodeableConcept       `json:"type,omitempty"`
	Status                  string                 `json:"status"`
	Experimental            *bool                  `json:"experimental,omitempty"`
	SubjectCodeableConcept  *CodeableConcept       `json:"subjectCodeableConcept"`
	SubjectReference        *Reference             `json:"subjectReference"`
	SubjectCanonical        *string                `json:"subjectCanonical"`
	Date                    *string                `json:"date,omitempty"`
	Publisher               *string                `json:"publisher,omitempty"`
	Contact                 []ContactDetail        `json:"contact,omitempty"`
	Description             *string                `json:"description,omitempty"`
	UseContext              []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction            []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                 *string                `json:"purpose,omitempty"`
	Usage                   *string                `json:"usage,omitempty"`
	Copyright               *string                `json:"copyright,omitempty"`
	CopyrightLabel          *string                `json:"copyrightLabel,omitempty"`
	ApprovalDate            *string                `json:"approvalDate,omitempty"`
	LastReviewDate          *string                `json:"lastReviewDate,omitempty"`
	EffectivePeriod         *Period                `json:"effectivePeriod,omitempty"`
	Topic                   []CodeableConcept      `json:"topic,omitempty"`
	Author                  []ContactDetail        `json:"author,omitempty"`
	Editor                  []ContactDetail        `json:"editor,omitempty"`
	Reviewer                []ContactDetail        `json:"reviewer,omitempty"`
	Endorser                []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact         []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Library                 []string               `json:"library,omitempty"`
	Goal                    []PlanDefinitionGoal   `json:"goal,omitempty"`
	Actor                   []PlanDefinitionActor  `json:"actor,omitempty"`
	Action                  []PlanDefinitionAction `json:"action,omitempty"`
	AsNeededBoolean         *bool                  `json:"asNeededBoolean"`
	AsNeededCodeableConcept *CodeableConcept       `json:"asNeededCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionGoalTarget struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *Quantity        `json:"detailQuantity"`
	DetailRange           *Range           `json:"detailRange"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept"`
	DetailString          *string          `json:"detailString"`
	DetailBoolean         *bool            `json:"detailBoolean"`
	DetailInteger         *int             `json:"detailInteger"`
	DetailRatio           *Ratio           `json:"detailRatio"`
	Due                   *Duration        `json:"due,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActor struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Title             *string                     `json:"title,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Option            []PlanDefinitionActorOption `json:"option"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActorOption struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *string          `json:"type,omitempty"`
	TypeCanonical     *string          `json:"typeCanonical,omitempty"`
	TypeReference     *Reference       `json:"typeReference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionAction struct {
	Id                     *string                             `json:"id,omitempty"`
	Extension              []Extension                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                         `json:"modifierExtension,omitempty"`
	LinkId                 *string                             `json:"linkId,omitempty"`
	Prefix                 *string                             `json:"prefix,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	TextEquivalent         *string                             `json:"textEquivalent,omitempty"`
	Priority               *string                             `json:"priority,omitempty"`
	Code                   *CodeableConcept                    `json:"code,omitempty"`
	Reason                 []CodeableConcept                   `json:"reason,omitempty"`
	Documentation          []RelatedArtifact                   `json:"documentation,omitempty"`
	GoalId                 []string                            `json:"goalId,omitempty"`
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept"`
	SubjectReference       *Reference                          `json:"subjectReference"`
	SubjectCanonical       *string                             `json:"subjectCanonical"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []PlanDefinitionActionInput         `json:"input,omitempty"`
	Output                 []PlanDefinitionActionOutput        `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingAge              *Age                                `json:"timingAge"`
	TimingDuration         *Duration                           `json:"timingDuration"`
	TimingRange            *Range                              `json:"timingRange"`
	TimingTiming           *Timing                             `json:"timingTiming"`
	Location               *CodeableReference                  `json:"location,omitempty"`
	Participant            []PlanDefinitionActionParticipant   `json:"participant,omitempty"`
	Type                   *CodeableConcept                    `json:"type,omitempty"`
	GroupingBehavior       *string                             `json:"groupingBehavior,omitempty"`
	SelectionBehavior      *string                             `json:"selectionBehavior,omitempty"`
	RequiredBehavior       *string                             `json:"requiredBehavior,omitempty"`
	PrecheckBehavior       *string                             `json:"precheckBehavior,omitempty"`
	CardinalityBehavior    *string                             `json:"cardinalityBehavior,omitempty"`
	DefinitionCanonical    *string                             `json:"definitionCanonical"`
	DefinitionUri          *string                             `json:"definitionUri"`
	Transform              *string                             `json:"transform,omitempty"`
	DynamicValue           []PlanDefinitionActionDynamicValue  `json:"dynamicValue,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionInput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionOutput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionRelatedAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	TargetId          string      `json:"targetId"`
	Relationship      string      `json:"relationship"`
	EndRelationship   *string     `json:"endRelationship,omitempty"`
	OffsetDuration    *Duration   `json:"offsetDuration"`
	OffsetRange       *Range      `json:"offsetRange"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	ActorId           *string          `json:"actorId,omitempty"`
	Type              *string          `json:"type,omitempty"`
	TypeCanonical     *string          `json:"typeCanonical,omitempty"`
	TypeReference     *Reference       `json:"typeReference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
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

func (resource *PlanDefinition) PlanDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionTopic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", &resource.Topic[0], optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalCategory(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Goal[numGoal].Category, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalDescription(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("description", nil, optionsValueSet)
	}
	return CodeableConceptSelect("description", &resource.Goal[numGoal].Description, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalPriority(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Goal[numGoal].Priority, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalStart(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("start", nil, optionsValueSet)
	}
	return CodeableConceptSelect("start", resource.Goal[numGoal].Start, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalAddresses(numGoal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal) >= numGoal {
		return CodeableConceptSelect("addresses", nil, optionsValueSet)
	}
	return CodeableConceptSelect("addresses", &resource.Goal[numGoal].Addresses[0], optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionGoalTargetMeasure(numGoal int, numTarget int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Goal[numGoal].Target) >= numTarget {
		return CodeableConceptSelect("measure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("measure", resource.Goal[numGoal].Target[numTarget].Measure, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActorOptionType(numActor int, numOption int) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource != nil && len(resource.Actor[numActor].Option) >= numOption {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Actor[numActor].Option[numOption].Type, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActorOptionRole(numActor int, numOption int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Actor[numActor].Option) >= numOption {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Actor[numActor].Option[numOption].Role, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Action[numAction].Priority, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionCode(numAction int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Action) >= numAction {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Action[numAction].Code, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionReason(numAction int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Action) >= numAction {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Action[numAction].Reason[0], optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Action) >= numAction {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Action[numAction].Type, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("groupingBehavior", nil, optionsValueSet)
	}
	return CodeSelect("groupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("selectionBehavior", nil, optionsValueSet)
	}
	return CodeSelect("selectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("requiredBehavior", nil, optionsValueSet)
	}
	return CodeSelect("requiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("precheckBehavior", nil, optionsValueSet)
	}
	return CodeSelect("precheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource != nil && len(resource.Action) >= numAction {
		return CodeSelect("cardinalityBehavior", nil, optionsValueSet)
	}
	return CodeSelect("cardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource != nil && len(resource.Action[numAction].Condition) >= numCondition {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("relationship", nil, optionsValueSet)
	}
	return CodeSelect("relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionRelatedActionEndRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("endRelationship", nil, optionsValueSet)
	}
	return CodeSelect("endRelationship", resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionParticipantType(numAction int, numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource != nil && len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionParticipantFunction(numAction int, numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Action[numAction].Participant[numParticipant].Function, optionsValueSet)
}
