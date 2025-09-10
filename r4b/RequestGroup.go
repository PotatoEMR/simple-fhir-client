package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/RequestGroup
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

// http://hl7.org/fhir/r4b/StructureDefinition/RequestGroup
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

// http://hl7.org/fhir/r4b/StructureDefinition/RequestGroup
type RequestGroupActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/RequestGroup
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
func (r RequestGroup) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RequestGroup/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RequestGroup"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *RequestGroup) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *RequestGroup) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *RequestGroup) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *RequestGroup) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RequestGroup) T_ActionPrefix(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].prefix", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].prefix", resource.Action[numAction].Prefix, htmlAttrs)
}
func (resource *RequestGroup) T_ActionTitle(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].title", resource.Action[numAction].Title, htmlAttrs)
}
func (resource *RequestGroup) T_ActionDescription(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].description", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].description", resource.Action[numAction].Description, htmlAttrs)
}
func (resource *RequestGroup) T_ActionTextEquivalent(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", resource.Action[numAction].TextEquivalent, htmlAttrs)
}
func (resource *RequestGroup) T_ActionPriority(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", resource.Action[numAction].Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCode >= len(resource.Action[numAction].Code) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionTimingDateTime(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return DateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", resource.Action[numAction].TimingDateTime, htmlAttrs)
}
func (resource *RequestGroup) T_ActionType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", resource.Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionGroupingBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionSelectionBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionRequiredBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionPrecheckBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionCardinalityBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionConditionKind(numAction int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet, htmlAttrs)
}
func (resource *RequestGroup) T_ActionRelatedActionActionId(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].actionId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].actionId", &resource.Action[numAction].RelatedAction[numRelatedAction].ActionId, htmlAttrs)
}
func (resource *RequestGroup) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
