package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	SubjectCodeableConcept *CodeableConcept       `json:"subjectCodeableConcept"`
	SubjectReference       *Reference             `json:"subjectReference"`
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
	DetailQuantity        *Quantity        `json:"detailQuantity"`
	DetailRange           *Range           `json:"detailRange"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept"`
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
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept"`
	SubjectReference       *Reference                          `json:"subjectReference"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []DataRequirement                   `json:"input,omitempty"`
	Output                 []DataRequirement                   `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime         *string                             `json:"timingDateTime"`
	TimingAge              *Age                                `json:"timingAge"`
	TimingPeriod           *Period                             `json:"timingPeriod"`
	TimingDuration         *Duration                           `json:"timingDuration"`
	TimingRange            *Range                              `json:"timingRange"`
	TimingTiming           *Timing                             `json:"timingTiming"`
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
	OffsetDuration    *Duration   `json:"offsetDuration"`
	OffsetRange       *Range      `json:"offsetRange"`
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

func (resource *PlanDefinition) PlanDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionPriority(numAction int) templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionGroupingBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_grouping_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].GroupingBehavior
	}
	return CodeSelect("groupingBehavior", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionSelectionBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_selection_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].SelectionBehavior
	}
	return CodeSelect("selectionBehavior", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionRequiredBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_required_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].RequiredBehavior
	}
	return CodeSelect("requiredBehavior", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionPrecheckBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_precheck_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].PrecheckBehavior
	}
	return CodeSelect("precheckBehavior", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionCardinalityBehavior(numAction int) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior
	currentVal := ""
	if resource != nil && len(resource.Action) >= numAction {
		currentVal = *resource.Action[numAction].CardinalityBehavior
	}
	return CodeSelect("cardinalityBehavior", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionConditionKind(numAction int, numCondition int) templ.Component {
	optionsValueSet := VSAction_condition_kind
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].Condition) >= numCondition {
		currentVal = resource.Action[numAction].Condition[numCondition].Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionRelatedActionRelationship(numAction int, numRelatedAction int) templ.Component {
	optionsValueSet := VSAction_relationship_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].RelatedAction) >= numRelatedAction {
		currentVal = resource.Action[numAction].RelatedAction[numRelatedAction].Relationship
	}
	return CodeSelect("relationship", currentVal, optionsValueSet)
}
func (resource *PlanDefinition) PlanDefinitionActionParticipantType(numAction int, numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type
	currentVal := ""
	if resource != nil && len(resource.Action[numAction].Participant) >= numParticipant {
		currentVal = resource.Action[numAction].Participant[numParticipant].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
