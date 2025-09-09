package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	AuthoredOn            *time.Time                   `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	TimingDateTime      *time.Time                                `json:"timingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r RequestOrchestration) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RequestOrchestration/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RequestOrchestration"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *RequestOrchestration) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("RequestOrchestration.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *RequestOrchestration) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("RequestOrchestration.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *RequestOrchestration) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("RequestOrchestration.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("RequestOrchestration.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("RequestOrchestration.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("RequestOrchestration.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RequestOrchestration.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("RequestOrchestration.AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("RequestOrchestration.AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *RequestOrchestration) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("RequestOrchestration.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("RequestOrchestration.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionLinkId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].LinkId", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].LinkId", resource.Action[numAction].LinkId, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPrefix(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Prefix", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Prefix", resource.Action[numAction].Prefix, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTitle(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Title", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Title", resource.Action[numAction].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDescription(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Description", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Description", resource.Action[numAction].Description, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTextEquivalent(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TextEquivalent", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TextEquivalent", resource.Action[numAction].TextEquivalent, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPriority(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Priority", resource.Action[numAction].Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCode >= len(resource.Action[numAction].Code) {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingDateTime(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return DateTimeInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].TimingDateTime", resource.Action[numAction].TimingDateTime, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionType(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Type", resource.Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionGroupingBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionSelectionBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRequiredBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPrecheckBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionCardinalityBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDefinitionCanonical(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DefinitionCanonical", resource.Action[numAction].DefinitionCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDefinitionUri(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DefinitionUri", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DefinitionUri", resource.Action[numAction].DefinitionUri, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTransform(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Transform", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Transform", resource.Action[numAction].Transform, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionConditionKind(numAction int, numCondition int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionInputTitle(numAction int, numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Title", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].Title", resource.Action[numAction].Input[numInput].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionInputRelatedData(numAction int, numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].RelatedData", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Input["+strconv.Itoa(numInput)+"].RelatedData", resource.Action[numAction].Input[numInput].RelatedData, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionOutputTitle(numAction int, numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Title", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].Title", resource.Action[numAction].Output[numOutput].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionOutputRelatedData(numAction int, numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].RelatedData", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Output["+strconv.Itoa(numOutput)+"].RelatedData", resource.Action[numAction].Output[numOutput].RelatedData, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionTargetId(numAction int, numRelatedAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].TargetId", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].TargetId", &resource.Action[numAction].RelatedAction[numRelatedAction].TargetId, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionEndRelationship(numAction int, numRelatedAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].EndRelationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].EndRelationship", resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantType(numAction int, numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantTypeCanonical(numAction int, numParticipant int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].TypeCanonical", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].TypeCanonical", resource.Action[numAction].Participant[numParticipant].TypeCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantFunction(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Action[numAction].Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantActorCanonical(numAction int, numParticipant int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].ActorCanonical", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].ActorCanonical", resource.Action[numAction].Participant[numParticipant].ActorCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDynamicValuePath(numAction int, numDynamicValue int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil, htmlAttrs)
	}
	return StringInput("RequestOrchestration.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", resource.Action[numAction].DynamicValue[numDynamicValue].Path, htmlAttrs)
}
