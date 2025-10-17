package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	AuthoredOn            *FhirDateTime                `json:"authoredOn,omitempty"`
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
	TimingDateTime      *FhirDateTime                             `json:"timingDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r RequestOrchestration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestOrchestration
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestOrchestration: OtherRequestOrchestration(r),
		ResourceType:              "RequestOrchestration",
	})
}

// json -> struct, first reject if resourceType != RequestOrchestration
func (r *RequestOrchestration) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "RequestOrchestration" {
		return errors.New("resourceType not RequestOrchestration")
	}
	return json.Unmarshal(data, (*OtherRequestOrchestration)(r))
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
func (resource *RequestOrchestration) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *RequestOrchestration) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *RequestOrchestration) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *RequestOrchestration) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *RequestOrchestration) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
}
func (resource *RequestOrchestration) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *RequestOrchestration) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *RequestOrchestration) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *RequestOrchestration) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
func (resource *RequestOrchestration) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *RequestOrchestration) T_Goal(frs []FhirResource, numGoal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return ReferenceInput(frs, "goal["+strconv.Itoa(numGoal)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "goal["+strconv.Itoa(numGoal)+"]", &resource.Goal[numGoal], htmlAttrs)
}
func (resource *RequestOrchestration) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionLinkId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].linkId", resource.Action[numAction].LinkId, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPrefix(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].prefix", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].prefix", resource.Action[numAction].Prefix, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTitle(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].title", resource.Action[numAction].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDescription(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].description", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].description", resource.Action[numAction].Description, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTextEquivalent(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", resource.Action[numAction].TextEquivalent, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPriority(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", resource.Action[numAction].Priority, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCode >= len(resource.Action[numAction].Code) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDocumentation(numAction int, numDocumentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDocumentation >= len(resource.Action[numAction].Documentation) {
		return RelatedArtifactInput("action["+strconv.Itoa(numAction)+"].documentation["+strconv.Itoa(numDocumentation)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("action["+strconv.Itoa(numAction)+"].documentation["+strconv.Itoa(numDocumentation)+"]", &resource.Action[numAction].Documentation[numDocumentation], htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionGoal(frs []FhirResource, numAction int, numGoal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numGoal >= len(resource.Action[numAction].Goal) {
		return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].goal["+strconv.Itoa(numGoal)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].goal["+strconv.Itoa(numGoal)+"]", &resource.Action[numAction].Goal[numGoal], htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingDateTime(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return FhirDateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", resource.Action[numAction].TimingDateTime, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingAge(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return AgeInput("action["+strconv.Itoa(numAction)+"].timingAge", nil, htmlAttrs)
	}
	return AgeInput("action["+strconv.Itoa(numAction)+"].timingAge", resource.Action[numAction].TimingAge, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingPeriod(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return PeriodInput("action["+strconv.Itoa(numAction)+"].timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("action["+strconv.Itoa(numAction)+"].timingPeriod", resource.Action[numAction].TimingPeriod, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingDuration(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return DurationInput("action["+strconv.Itoa(numAction)+"].timingDuration", nil, htmlAttrs)
	}
	return DurationInput("action["+strconv.Itoa(numAction)+"].timingDuration", resource.Action[numAction].TimingDuration, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingRange(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return RangeInput("action["+strconv.Itoa(numAction)+"].timingRange", nil, htmlAttrs)
	}
	return RangeInput("action["+strconv.Itoa(numAction)+"].timingRange", resource.Action[numAction].TimingRange, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTimingTiming(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return TimingInput("action["+strconv.Itoa(numAction)+"].timingTiming", nil, htmlAttrs)
	}
	return TimingInput("action["+strconv.Itoa(numAction)+"].timingTiming", resource.Action[numAction].TimingTiming, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionLocation(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableReferenceInput("action["+strconv.Itoa(numAction)+"].location", nil, htmlAttrs)
	}
	return CodeableReferenceInput("action["+strconv.Itoa(numAction)+"].location", resource.Action[numAction].Location, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", resource.Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionGroupingBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionSelectionBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRequiredBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionPrecheckBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionCardinalityBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionResource(frs []FhirResource, numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].resource", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].resource", resource.Action[numAction].Resource, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDefinitionCanonical(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].definitionCanonical", resource.Action[numAction].DefinitionCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDefinitionUri(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].definitionUri", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].definitionUri", resource.Action[numAction].DefinitionUri, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionTransform(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].transform", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].transform", resource.Action[numAction].Transform, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionConditionKind(numAction int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionConditionExpression(numAction int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return ExpressionInput("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].expression", nil, htmlAttrs)
	}
	return ExpressionInput("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].expression", resource.Action[numAction].Condition[numCondition].Expression, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionInputTitle(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].title", resource.Action[numAction].Input[numInput].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionInputRequirement(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return DataRequirementInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].requirement", nil, htmlAttrs)
	}
	return DataRequirementInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].requirement", resource.Action[numAction].Input[numInput].Requirement, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionInputRelatedData(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].relatedData", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].relatedData", resource.Action[numAction].Input[numInput].RelatedData, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionOutputTitle(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].title", resource.Action[numAction].Output[numOutput].Title, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionOutputRequirement(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return DataRequirementInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].requirement", nil, htmlAttrs)
	}
	return DataRequirementInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].requirement", resource.Action[numAction].Output[numOutput].Requirement, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionOutputRelatedData(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].relatedData", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].relatedData", resource.Action[numAction].Output[numOutput].RelatedData, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionTargetId(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].targetId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].targetId", &resource.Action[numAction].RelatedAction[numRelatedAction].TargetId, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionEndRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].endRelationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].endRelationship", resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionOffsetDuration(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return DurationInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetDuration", nil, htmlAttrs)
	}
	return DurationInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetDuration", resource.Action[numAction].RelatedAction[numRelatedAction].OffsetDuration, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionRelatedActionOffsetRange(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return RangeInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetRange", nil, htmlAttrs)
	}
	return RangeInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetRange", resource.Action[numAction].RelatedAction[numRelatedAction].OffsetRange, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantType(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantTypeCanonical(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeCanonical", resource.Action[numAction].Participant[numParticipant].TypeCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantTypeReference(frs []FhirResource, numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeReference", resource.Action[numAction].Participant[numParticipant].TypeReference, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantFunction(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].function", resource.Action[numAction].Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantActorCanonical(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorCanonical", resource.Action[numAction].Participant[numParticipant].ActorCanonical, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionParticipantActorReference(frs []FhirResource, numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorReference", resource.Action[numAction].Participant[numParticipant].ActorReference, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDynamicValuePath(numAction int, numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", resource.Action[numAction].DynamicValue[numDynamicValue].Path, htmlAttrs)
}
func (resource *RequestOrchestration) T_ActionDynamicValueExpression(numAction int, numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return ExpressionInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", nil, htmlAttrs)
	}
	return ExpressionInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", resource.Action[numAction].DynamicValue[numDynamicValue].Expression, htmlAttrs)
}
