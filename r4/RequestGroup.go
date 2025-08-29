package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	TimingDateTime      *string                           `json:"timingDateTime"`
	TimingAge           *Age                              `json:"timingAge"`
	TimingPeriod        *Period                           `json:"timingPeriod"`
	TimingDuration      *Duration                         `json:"timingDuration"`
	TimingRange         *Range                            `json:"timingRange"`
	TimingTiming        *Timing                           `json:"timingTiming"`
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
	OffsetDuration    *Duration   `json:"offsetDuration"`
	OffsetRange       *Range      `json:"offsetRange"`
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
func (resource *RequestGroup) RequestGroupLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupStatus() templ.Component {
	optionsValueSet := VSRequest_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupIntent() templ.Component {
	optionsValueSet := VSRequest_intent
	currentVal := ""
	if resource != nil {
		currentVal = resource.Intent
	}
	return CodeSelect("intent", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupPriority() templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].GroupingBehavior
	}
	return CodeSelect("groupingBehavior", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].SelectionBehavior
	}
	return CodeSelect("selectionBehavior", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].RequiredBehavior
	}
	return CodeSelect("requiredBehavior", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].PrecheckBehavior
	}
	return CodeSelect("precheckBehavior", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].CardinalityBehavior
	}
	return CodeSelect("cardinalityBehavior", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].Condition) >= numCondition {
		currentVal = resource.Action[numAction].Condition[numCondition].Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		currentVal = resource.Action[numAction].RelatedAction[numRelatedAction].Relationship
	}
	return CodeSelect("relationship", currentVal, optionsValueSet)
}
