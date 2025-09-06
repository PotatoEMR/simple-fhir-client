package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestration struct {
	Id                    *string                      `json:"id,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	Contained             []Resource                   `json:"contained,omitempty"`
	Extension             []Extension                  `json:"extension,omitempty"`
	ModifierExtension     []Extension                  `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                 `json:"identifier,omitempty"`
	InstantiatesCanonical []string                     `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                     `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                  `json:"basedOn,omitempty"`
	Replaces              []Reference                  `json:"replaces,omitempty"`
	GroupIdentifier       *Identifier                  `json:"groupIdentifier,omitempty"`
	Status                string                       `json:"status"`
	Intent                string                       `json:"intent"`
	Priority              *string                      `json:"priority,omitempty"`
	Code                  *CodeableConcept             `json:"code,omitempty"`
	Subject               *Reference                   `json:"subject,omitempty"`
	Encounter             *Reference                   `json:"encounter,omitempty"`
	AuthoredOn            *string                      `json:"authoredOn,omitempty"`
	Author                *Reference                   `json:"author,omitempty"`
	Reason                []CodeableReference          `json:"reason,omitempty"`
	Goal                  []Reference                  `json:"goal,omitempty"`
	Note                  []Annotation                 `json:"note,omitempty"`
	Action                []RequestOrchestrationAction `json:"action,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationAction struct {
	Id                  *string                                   `json:"id,omitempty"`
	Extension           []Extension                               `json:"extension,omitempty"`
	ModifierExtension   []Extension                               `json:"modifierExtension,omitempty"`
	LinkId              *string                                   `json:"linkId,omitempty"`
	Prefix              *string                                   `json:"prefix,omitempty"`
	Title               *string                                   `json:"title,omitempty"`
	Description         *string                                   `json:"description,omitempty"`
	TextEquivalent      *string                                   `json:"textEquivalent,omitempty"`
	Priority            *string                                   `json:"priority,omitempty"`
	Code                []CodeableConcept                         `json:"code,omitempty"`
	Documentation       []RelatedArtifact                         `json:"documentation,omitempty"`
	Goal                []Reference                               `json:"goal,omitempty"`
	Condition           []RequestOrchestrationActionCondition     `json:"condition,omitempty"`
	Input               []RequestOrchestrationActionInput         `json:"input,omitempty"`
	Output              []RequestOrchestrationActionOutput        `json:"output,omitempty"`
	RelatedAction       []RequestOrchestrationActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime      *string                                   `json:"timingDateTime,omitempty"`
	TimingAge           *Age                                      `json:"timingAge,omitempty"`
	TimingPeriod        *Period                                   `json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                                 `json:"timingDuration,omitempty"`
	TimingRange         *Range                                    `json:"timingRange,omitempty"`
	TimingTiming        *Timing                                   `json:"timingTiming,omitempty"`
	Location            *CodeableReference                        `json:"location,omitempty"`
	Participant         []RequestOrchestrationActionParticipant   `json:"participant,omitempty"`
	Type                *CodeableConcept                          `json:"type,omitempty"`
	GroupingBehavior    *string                                   `json:"groupingBehavior,omitempty"`
	SelectionBehavior   *string                                   `json:"selectionBehavior,omitempty"`
	RequiredBehavior    *string                                   `json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *string                                   `json:"precheckBehavior,omitempty"`
	CardinalityBehavior *string                                   `json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                                `json:"resource,omitempty"`
	DefinitionCanonical *string                                   `json:"definitionCanonical,omitempty"`
	DefinitionUri       *string                                   `json:"definitionUri,omitempty"`
	Transform           *string                                   `json:"transform,omitempty"`
	DynamicValue        []RequestOrchestrationActionDynamicValue  `json:"dynamicValue,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionInput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionOutput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionRelatedAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	TargetId          string      `json:"targetId"`
	Relationship      string      `json:"relationship"`
	EndRelationship   *string     `json:"endRelationship,omitempty"`
	OffsetDuration    *Duration   `json:"offsetDuration,omitempty"`
	OffsetRange       *Range      `json:"offsetRange,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *string          `json:"type,omitempty"`
	TypeCanonical     *string          `json:"typeCanonical,omitempty"`
	TypeReference     *Reference       `json:"typeReference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	ActorCanonical    *string          `json:"actorCanonical,omitempty"`
	ActorReference    *Reference       `json:"actorReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionDynamicValue struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Path              *string     `json:"path,omitempty"`
	Expression        *Expression `json:"expression,omitempty"`
}

type OtherRequestOrchestration RequestOrchestration

// on convert struct to json, automatically add resourceType=RequestOrchestration
func (r RequestOrchestration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestOrchestration
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestOrchestration: OtherRequestOrchestration(r),
		ResourceType:              "RequestOrchestration",
	})
}

func (resource *RequestOrchestration) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RequestOrchestration.Id", nil)
	}
	return StringInput("RequestOrchestration.Id", resource.Id)
}
func (resource *RequestOrchestration) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RequestOrchestration.ImplicitRules", nil)
	}
	return StringInput("RequestOrchestration.ImplicitRules", resource.ImplicitRules)
}
func (resource *RequestOrchestration) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RequestOrchestration.Language", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Language", resource.Language, optionsValueSet)
}
func (resource *RequestOrchestration) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("RequestOrchestration.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("RequestOrchestration.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *RequestOrchestration) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("RequestOrchestration.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("RequestOrchestration.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *RequestOrchestration) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("RequestOrchestration.Status", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Status", &resource.Status, optionsValueSet)
}
func (resource *RequestOrchestration) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("RequestOrchestration.Intent", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Intent", &resource.Intent, optionsValueSet)
}
func (resource *RequestOrchestration) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("RequestOrchestration.Priority", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Priority", resource.Priority, optionsValueSet)
}
func (resource *RequestOrchestration) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RequestOrchestration.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestOrchestration.Code", resource.Code, optionsValueSet)
}
func (resource *RequestOrchestration) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("RequestOrchestration.AuthoredOn", nil)
	}
	return StringInput("RequestOrchestration.AuthoredOn", resource.AuthoredOn)
}
func (resource *RequestOrchestration) T_ActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Id", resource.Action[numAction].Id)
}
func (resource *RequestOrchestration) T_ActionLinkId(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].LinkId", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].LinkId", resource.Action[numAction].LinkId)
}
func (resource *RequestOrchestration) T_ActionPrefix(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Prefix", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Prefix", resource.Action[numAction].Prefix)
}
func (resource *RequestOrchestration) T_ActionTitle(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Title", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Title", resource.Action[numAction].Title)
}
func (resource *RequestOrchestration) T_ActionDescription(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Description", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Description", resource.Action[numAction].Description)
}
func (resource *RequestOrchestration) T_ActionTextEquivalent(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TextEquivalent", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TextEquivalent", resource.Action[numAction].TextEquivalent)
}
func (resource *RequestOrchestration) T_ActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Priority", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Priority", resource.Action[numAction].Priority, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Code) >= numCode {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Type", resource.Action[numAction].Type, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionTransform(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Transform", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Transform", resource.Action[numAction].Transform)
}
func (resource *RequestOrchestration) T_ActionConditionId(numAction int, numCondition int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", resource.Action[numAction].Condition[numCondition].Id)
}
func (resource *RequestOrchestration) T_ActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionInputId(numAction int, numInput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Input) >= numInput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Id", resource.Action[numAction].Input[numInput].Id)
}
func (resource *RequestOrchestration) T_ActionInputTitle(numAction int, numInput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Input) >= numInput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Title", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Title", resource.Action[numAction].Input[numInput].Title)
}
func (resource *RequestOrchestration) T_ActionInputRelatedData(numAction int, numInput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Input) >= numInput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].RelatedData", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].RelatedData", resource.Action[numAction].Input[numInput].RelatedData)
}
func (resource *RequestOrchestration) T_ActionOutputId(numAction int, numOutput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Output) >= numOutput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Id", resource.Action[numAction].Output[numOutput].Id)
}
func (resource *RequestOrchestration) T_ActionOutputTitle(numAction int, numOutput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Output) >= numOutput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Title", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Title", resource.Action[numAction].Output[numOutput].Title)
}
func (resource *RequestOrchestration) T_ActionOutputRelatedData(numAction int, numOutput int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Output) >= numOutput {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].RelatedData", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].RelatedData", resource.Action[numAction].Output[numOutput].RelatedData)
}
func (resource *RequestOrchestration) T_ActionRelatedActionId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", resource.Action[numAction].RelatedAction[numRelatedAction].Id)
}
func (resource *RequestOrchestration) T_ActionRelatedActionTargetId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].TargetId", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].TargetId", &resource.Action[numAction].RelatedAction[numRelatedAction].TargetId)
}
func (resource *RequestOrchestration) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionRelatedActionEndRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].EndRelationship", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].EndRelationship", resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionParticipantId(numAction int, numParticipant int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Action[numAction].Participant[numParticipant].Id)
}
func (resource *RequestOrchestration) T_ActionParticipantType(numAction int, numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionParticipantTypeCanonical(numAction int, numParticipant int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].TypeCanonical", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].TypeCanonical", resource.Action[numAction].Participant[numParticipant].TypeCanonical)
}
func (resource *RequestOrchestration) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionParticipantFunction(numAction int, numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Participant) >= numParticipant {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Action[numAction].Participant[numParticipant].Function, optionsValueSet)
}
func (resource *RequestOrchestration) T_ActionDynamicValueId(numAction int, numDynamicValue int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].DynamicValue) >= numDynamicValue {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", resource.Action[numAction].DynamicValue[numDynamicValue].Id)
}
func (resource *RequestOrchestration) T_ActionDynamicValuePath(numAction int, numDynamicValue int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].DynamicValue) >= numDynamicValue {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", resource.Action[numAction].DynamicValue[numDynamicValue].Path)
}
