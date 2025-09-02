package r4

//generated with command go run ./bultaoreune -nodownload
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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupStatus() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupIntent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Action[numAction].Priority, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionCode(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Action) >= numAction {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Action[numAction].Code[0], optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Action) >= numAction {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Action[numAction].Type, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("groupingBehavior", nil, optionsValueSet)
	}
	return CodeSelect("groupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("selectionBehavior", nil, optionsValueSet)
	}
	return CodeSelect("selectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("requiredBehavior", nil, optionsValueSet)
	}
	return CodeSelect("requiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("precheckBehavior", nil, optionsValueSet)
	}
	return CodeSelect("precheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil && len(resource.Action) >= numAction {
		return CodeSelect("cardinalityBehavior", nil, optionsValueSet)
	}
	return CodeSelect("cardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil && len(resource.Action[numAction].Condition) >= numCondition {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet)
}
func (resource *RequestGroup) RequestGroupActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		return CodeSelect("relationship", nil, optionsValueSet)
	}
	return CodeSelect("relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet)
}
