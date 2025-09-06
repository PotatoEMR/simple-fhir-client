package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                    *string              `json:"id,omitempty"`
	Meta                  *Meta                `json:"meta,omitempty"`
	ImplicitRules         *string              `json:"implicitRules,omitempty"`
	Language              *string              `json:"language,omitempty"`
	Text                  *Narrative           `json:"text,omitempty"`
	Contained             []Resource           `json:"contained,omitempty"`
	Extension             []Extension          `json:"extension,omitempty"`
	ModifierExtension     []Extension          `json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `json:"identifier,omitempty"`
	InstantiatesCanonical []string             `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string             `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference          `json:"basedOn,omitempty"`
	Replaces              []Reference          `json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `json:"groupIdentifier,omitempty"`
	Status                string               `json:"status"`
	Intent                string               `json:"intent"`
	Priority              *string              `json:"priority,omitempty"`
	Code                  *CodeableConcept     `json:"code,omitempty"`
	Subject               *Reference           `json:"subject,omitempty"`
	Encounter             *Reference           `json:"encounter,omitempty"`
	AuthoredOn            *string              `json:"authoredOn,omitempty"`
	Author                *Reference           `json:"author,omitempty"`
	ReasonCode            []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference       []Reference          `json:"reasonReference,omitempty"`
	Note                  []Annotation         `json:"note,omitempty"`
	Action                []RequestGroupAction `json:"action,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RequestGroup
type RequestGroupAction struct {
	Id                  *string                           `json:"id,omitempty"`
	Extension           []Extension                       `json:"extension,omitempty"`
	ModifierExtension   []Extension                       `json:"modifierExtension,omitempty"`
	Prefix              *string                           `json:"prefix,omitempty"`
	Title               *string                           `json:"title,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	TextEquivalent      *string                           `json:"textEquivalent,omitempty"`
	Priority            *string                           `json:"priority,omitempty"`
	Code                []CodeableConcept                 `json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime      *string                           `json:"timingDateTime,omitempty"`
	TimingAge           *Age                              `json:"timingAge,omitempty"`
	TimingPeriod        *Period                           `json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                         `json:"timingDuration,omitempty"`
	TimingRange         *Range                            `json:"timingRange,omitempty"`
	TimingTiming        *Timing                           `json:"timingTiming,omitempty"`
	Participant         []Reference                       `json:"participant,omitempty"`
	Type                *CodeableConcept                  `json:"type,omitempty"`
	GroupingBehavior    *string                           `json:"groupingBehavior,omitempty"`
	SelectionBehavior   *string                           `json:"selectionBehavior,omitempty"`
	RequiredBehavior    *string                           `json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *string                           `json:"precheckBehavior,omitempty"`
	CardinalityBehavior *string                           `json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RequestGroup
type RequestGroupActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RequestGroup
type RequestGroupActionRelatedAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ActionId          string      `json:"actionId"`
	Relationship      string      `json:"relationship"`
	OffsetDuration    *Duration   `json:"offsetDuration,omitempty"`
	OffsetRange       *Range      `json:"offsetRange,omitempty"`
}

type OtherRequestGroup RequestGroup

// on convert struct to json, automatically add resourceType=RequestGroup
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

func (resource *RequestGroup) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RequestGroup.Id", nil)
	}
	return StringInput("RequestGroup.Id", resource.Id)
}
func (resource *RequestGroup) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RequestGroup.ImplicitRules", nil)
	}
	return StringInput("RequestGroup.ImplicitRules", resource.ImplicitRules)
}
func (resource *RequestGroup) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RequestGroup.Language", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Language", resource.Language, optionsValueSet)
}
func (resource *RequestGroup) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("RequestGroup.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("RequestGroup.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *RequestGroup) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("RequestGroup.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("RequestGroup.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *RequestGroup) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("RequestGroup.Status", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Status", &resource.Status, optionsValueSet)
}
func (resource *RequestGroup) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("RequestGroup.Intent", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Intent", &resource.Intent, optionsValueSet)
}
func (resource *RequestGroup) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("RequestGroup.Priority", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Priority", resource.Priority, optionsValueSet)
}
func (resource *RequestGroup) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RequestGroup.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestGroup.Code", resource.Code, optionsValueSet)
}
func (resource *RequestGroup) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("RequestGroup.AuthoredOn", nil)
	}
	return StringInput("RequestGroup.AuthoredOn", resource.AuthoredOn)
}
func (resource *RequestGroup) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("RequestGroup.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestGroup.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *RequestGroup) T_ActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Id", resource.Action[numAction].Id)
}
func (resource *RequestGroup) T_ActionPrefix(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Prefix", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Prefix", resource.Action[numAction].Prefix)
}
func (resource *RequestGroup) T_ActionTitle(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Title", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Title", resource.Action[numAction].Title)
}
func (resource *RequestGroup) T_ActionDescription(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Description", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Description", resource.Action[numAction].Description)
}
func (resource *RequestGroup) T_ActionTextEquivalent(numAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].TextEquivalent", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].TextEquivalent", resource.Action[numAction].TextEquivalent)
}
func (resource *RequestGroup) T_ActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Priority", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Priority", resource.Action[numAction].Priority, optionsValueSet)
}
func (resource *RequestGroup) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Code) >= numCode {
		return CodeableConceptSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet)
}
func (resource *RequestGroup) T_ActionType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Action) >= numAction {
		return CodeableConceptSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Type", resource.Action[numAction].Type, optionsValueSet)
}
func (resource *RequestGroup) T_ActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet)
}
func (resource *RequestGroup) T_ActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet)
}
func (resource *RequestGroup) T_ActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet)
}
func (resource *RequestGroup) T_ActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet)
}
func (resource *RequestGroup) T_ActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || len(resource.Action) >= numAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet)
}
func (resource *RequestGroup) T_ActionConditionId(numAction int, numCondition int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Id", resource.Action[numAction].Condition[numCondition].Id)
}
func (resource *RequestGroup) T_ActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].Condition) >= numCondition {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet)
}
func (resource *RequestGroup) T_ActionRelatedActionId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Id", resource.Action[numAction].RelatedAction[numRelatedAction].Id)
}
func (resource *RequestGroup) T_ActionRelatedActionActionId(numAction int, numRelatedAction int) templ.Component {

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", nil)
	}
	return StringInput("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", &resource.Action[numAction].RelatedAction[numRelatedAction].ActionId)
}
func (resource *RequestGroup) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || len(resource.Action) >= numAction || len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", nil, optionsValueSet)
	}
	return CodeSelect("RequestGroup.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet)
}
