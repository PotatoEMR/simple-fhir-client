package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	TimingDateTime      *string                                   `json:"timingDateTime"`
	TimingAge           *Age                                      `json:"timingAge"`
	TimingPeriod        *Period                                   `json:"timingPeriod"`
	TimingDuration      *Duration                                 `json:"timingDuration"`
	TimingRange         *Range                                    `json:"timingRange"`
	TimingTiming        *Timing                                   `json:"timingTiming"`
	Location            *CodeableReference                        `json:"location,omitempty"`
	Participant         []RequestOrchestrationActionParticipant   `json:"participant,omitempty"`
	Type                *CodeableConcept                          `json:"type,omitempty"`
	GroupingBehavior    *string                                   `json:"groupingBehavior,omitempty"`
	SelectionBehavior   *string                                   `json:"selectionBehavior,omitempty"`
	RequiredBehavior    *string                                   `json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *string                                   `json:"precheckBehavior,omitempty"`
	CardinalityBehavior *string                                   `json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                                `json:"resource,omitempty"`
	DefinitionCanonical *string                                   `json:"definitionCanonical"`
	DefinitionUri       *string                                   `json:"definitionUri"`
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
	OffsetDuration    *Duration   `json:"offsetDuration"`
	OffsetRange       *Range      `json:"offsetRange"`
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
	ActorCanonical    *string          `json:"actorCanonical"`
	ActorReference    *Reference       `json:"actorReference"`
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

func (resource *RequestOrchestration) RequestOrchestrationLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationStatus() templ.Component {
	optionsValueSet := VSRequest_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationIntent() templ.Component {
	optionsValueSet := VSRequest_intent
	currentVal := ""
	if resource != nil {
		currentVal = resource.Intent
	}
	return CodeSelect("intent", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationPriority() templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].GroupingBehavior
	}
	return CodeSelect("groupingBehavior", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].SelectionBehavior
	}
	return CodeSelect("selectionBehavior", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].RequiredBehavior
	}
	return CodeSelect("requiredBehavior", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].PrecheckBehavior
	}
	return CodeSelect("precheckBehavior", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].CardinalityBehavior
	}
	return CodeSelect("cardinalityBehavior", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].Condition) >= numCondition {
		currentVal = resource.Action[numAction].Condition[numCondition].Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		currentVal = resource.Action[numAction].RelatedAction[numRelatedAction].Relationship
	}
	return CodeSelect("relationship", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionRelatedActionEndRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		currentVal = *resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship
	}
	return CodeSelect("endRelationship", currentVal, optionsValueSet)
}
func (resource *RequestOrchestration) RequestOrchestrationActionParticipantType(numAction int, numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].Participant) >= numParticipant {
		currentVal = *resource.Action[numAction].Participant[numParticipant].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
