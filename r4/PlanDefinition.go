package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	SubjectCodeableConcept *CodeableConcept       `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference             `json:"subjectReference,omitempty"`
	Date                   *time.Time             `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                `json:"publisher,omitempty"`
	Contact                []ContactDetail        `json:"contact,omitempty"`
	Description            *string                `json:"description,omitempty"`
	UseContext             []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                *string                `json:"purpose,omitempty"`
	Usage                  *string                `json:"usage,omitempty"`
	Copyright              *string                `json:"copyright,omitempty"`
	ApprovalDate           *time.Time             `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time             `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range           `json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
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
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                          `json:"subjectReference,omitempty"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []DataRequirement                   `json:"input,omitempty"`
	Output                 []DataRequirement                   `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime         *time.Time                          `json:"timingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	TimingAge              *Age                                `json:"timingAge,omitempty"`
	TimingPeriod           *Period                             `json:"timingPeriod,omitempty"`
	TimingDuration         *Duration                           `json:"timingDuration,omitempty"`
	TimingRange            *Range                              `json:"timingRange,omitempty"`
	TimingTiming           *Timing                             `json:"timingTiming,omitempty"`
	Participant            []PlanDefinitionActionParticipant   `json:"participant,omitempty"`
	Type                   *CodeableConcept                    `json:"type,omitempty"`
	GroupingBehavior       *string                             `json:"groupingBehavior,omitempty"`
	SelectionBehavior      *string                             `json:"selectionBehavior,omitempty"`
	RequiredBehavior       *string                             `json:"requiredBehavior,omitempty"`
	PrecheckBehavior       *string                             `json:"precheckBehavior,omitempty"`
	CardinalityBehavior    *string                             `json:"cardinalityBehavior,omitempty"`
	DefinitionCanonical    *string                             `json:"definitionCanonical,omitempty"`
	DefinitionUri          *string                             `json:"definitionUri,omitempty"`
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
	OffsetDuration    *Duration   `json:"offsetDuration,omitempty"`
	OffsetRange       *Range      `json:"offsetRange,omitempty"`
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
func (r PlanDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PlanDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PlanDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PlanDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *PlanDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *PlanDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *PlanDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *PlanDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Subtitle", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *PlanDefinition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PlanDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("PlanDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("PlanDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("PlanDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *PlanDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PlanDefinition.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PlanDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("PlanDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *PlanDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *PlanDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *PlanDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("PlanDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *PlanDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Usage", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Usage", resource.Usage, htmlAttrs)
}
func (resource *PlanDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PlanDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *PlanDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PlanDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("PlanDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *PlanDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PlanDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("PlanDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *PlanDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("PlanDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("PlanDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *PlanDefinition) T_GoalCategory(numGoal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Category", resource.Goal[numGoal].Category, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalDescription(numGoal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Description", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Description", &resource.Goal[numGoal].Description, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalPriority(numGoal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Priority", resource.Goal[numGoal].Priority, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalStart(numGoal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Start", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Start", resource.Goal[numGoal].Start, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalAddresses(numGoal int, numAddresses int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numAddresses >= len(resource.Goal[numGoal].Addresses) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Addresses["+strconv.Itoa(numAddresses)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Addresses["+strconv.Itoa(numAddresses)+"]", &resource.Goal[numGoal].Addresses[numAddresses], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetMeasure(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Measure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].Measure", resource.Goal[numGoal].Target[numTarget].Measure, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailCodeableConcept(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].DetailCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Goal["+strconv.Itoa(numGoal)+"].Target["+strconv.Itoa(numTarget)+"].DetailCodeableConcept", resource.Goal[numGoal].Target[numTarget].DetailCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPrefix(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Prefix", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Prefix", resource.Action[numAction].Prefix, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTitle(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Title", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Title", resource.Action[numAction].Title, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDescription(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Description", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Description", resource.Action[numAction].Description, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTextEquivalent(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TextEquivalent", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TextEquivalent", resource.Action[numAction].TextEquivalent, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPriority(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Priority", resource.Action[numAction].Priority, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCode >= len(resource.Action[numAction].Code) {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionReason(numAction int, numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numReason >= len(resource.Action[numAction].Reason) {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", &resource.Action[numAction].Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionGoalId(numAction int, numGoalId int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numGoalId >= len(resource.Action[numAction].GoalId) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GoalId["+strconv.Itoa(numGoalId)+"]", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GoalId["+strconv.Itoa(numGoalId)+"]", &resource.Action[numAction].GoalId[numGoalId], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionSubjectCodeableConcept(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SubjectCodeableConcept", resource.Action[numAction].SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingDateTime(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return DateTimeInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].TimingDateTime", resource.Action[numAction].TimingDateTime, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionType(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Type", resource.Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionGroupingBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].GroupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionSelectionBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].SelectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRequiredBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RequiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPrecheckBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].PrecheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionCardinalityBehavior(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].CardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDefinitionCanonical(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DefinitionCanonical", resource.Action[numAction].DefinitionCanonical, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDefinitionUri(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DefinitionUri", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DefinitionUri", resource.Action[numAction].DefinitionUri, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTransform(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Transform", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Transform", resource.Action[numAction].Transform, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionConditionKind(numAction int, numCondition int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Condition["+strconv.Itoa(numCondition)+"].Kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionActionId(numAction int, numRelatedAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].ActionId", &resource.Action[numAction].RelatedAction[numRelatedAction].ActionId, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].RelatedAction["+strconv.Itoa(numRelatedAction)+"].Relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantType(numAction int, numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Type", &resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PlanDefinition.Action["+strconv.Itoa(numAction)+"].Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDynamicValuePath(numAction int, numDynamicValue int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil, htmlAttrs)
	}
	return StringInput("PlanDefinition.Action["+strconv.Itoa(numAction)+"].DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", resource.Action[numAction].DynamicValue[numDynamicValue].Path, htmlAttrs)
}
