package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *FhirDateTime          `json:"date,omitempty"`
	Publisher              *string                `json:"publisher,omitempty"`
	Contact                []ContactDetail        `json:"contact,omitempty"`
	Description            *string                `json:"description,omitempty"`
	UseContext             []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                *string                `json:"purpose,omitempty"`
	Usage                  *string                `json:"usage,omitempty"`
	Copyright              *string                `json:"copyright,omitempty"`
	ApprovalDate           *FhirDate              `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate              `json:"lastReviewDate,omitempty"`
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
	TimingDateTime         *FhirDateTime                       `json:"timingDateTime,omitempty"`
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
func (resource *PlanDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *PlanDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *PlanDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *PlanDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *PlanDefinition) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *PlanDefinition) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *PlanDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_SubjectReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subjectReference", resource.SubjectReference, htmlAttrs)
}
func (resource *PlanDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *PlanDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *PlanDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *PlanDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *PlanDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *PlanDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *PlanDefinition) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *PlanDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *PlanDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *PlanDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *PlanDefinition) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *PlanDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *PlanDefinition) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *PlanDefinition) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *PlanDefinition) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *PlanDefinition) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *PlanDefinition) T_Library(numLibrary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *PlanDefinition) T_GoalCategory(numGoal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].category", resource.Goal[numGoal].Category, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalDescription(numGoal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].description", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].description", &resource.Goal[numGoal].Description, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalPriority(numGoal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].priority", resource.Goal[numGoal].Priority, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalStart(numGoal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].start", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].start", resource.Goal[numGoal].Start, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalAddresses(numGoal int, numAddresses int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numAddresses >= len(resource.Goal[numGoal].Addresses) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].addresses["+strconv.Itoa(numAddresses)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].addresses["+strconv.Itoa(numAddresses)+"]", &resource.Goal[numGoal].Addresses[numAddresses], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalDocumentation(numGoal int, numDocumentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numDocumentation >= len(resource.Goal[numGoal].Documentation) {
		return RelatedArtifactInput("goal["+strconv.Itoa(numGoal)+"].documentation["+strconv.Itoa(numDocumentation)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("goal["+strconv.Itoa(numGoal)+"].documentation["+strconv.Itoa(numDocumentation)+"]", &resource.Goal[numGoal].Documentation[numDocumentation], htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetMeasure(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].measure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].measure", resource.Goal[numGoal].Target[numTarget].Measure, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailQuantity(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return QuantityInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailQuantity", resource.Goal[numGoal].Target[numTarget].DetailQuantity, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailRange(numGoal int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return RangeInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailRange", nil, htmlAttrs)
	}
	return RangeInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailRange", resource.Goal[numGoal].Target[numTarget].DetailRange, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailCodeableConcept(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", resource.Goal[numGoal].Target[numTarget].DetailCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDue(numGoal int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return DurationInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].due", nil, htmlAttrs)
	}
	return DurationInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].due", resource.Goal[numGoal].Target[numTarget].Due, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPrefix(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].prefix", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].prefix", resource.Action[numAction].Prefix, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTitle(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].title", resource.Action[numAction].Title, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDescription(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].description", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].description", resource.Action[numAction].Description, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTextEquivalent(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].textEquivalent", resource.Action[numAction].TextEquivalent, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPriority(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].priority", resource.Action[numAction].Priority, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionCode(numAction int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCode >= len(resource.Action[numAction].Code) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code["+strconv.Itoa(numCode)+"]", &resource.Action[numAction].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionReason(numAction int, numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numReason >= len(resource.Action[numAction].Reason) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", &resource.Action[numAction].Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDocumentation(numAction int, numDocumentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDocumentation >= len(resource.Action[numAction].Documentation) {
		return RelatedArtifactInput("action["+strconv.Itoa(numAction)+"].documentation["+strconv.Itoa(numDocumentation)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("action["+strconv.Itoa(numAction)+"].documentation["+strconv.Itoa(numDocumentation)+"]", &resource.Action[numAction].Documentation[numDocumentation], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionGoalId(numAction int, numGoalId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numGoalId >= len(resource.Action[numAction].GoalId) {
		return StringInput("action["+strconv.Itoa(numAction)+"].goalId["+strconv.Itoa(numGoalId)+"]", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].goalId["+strconv.Itoa(numGoalId)+"]", &resource.Action[numAction].GoalId[numGoalId], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionSubjectCodeableConcept(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].subjectCodeableConcept", resource.Action[numAction].SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionSubjectReference(frs []FhirResource, numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "action["+strconv.Itoa(numAction)+"].subjectReference", resource.Action[numAction].SubjectReference, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTrigger(numAction int, numTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numTrigger >= len(resource.Action[numAction].Trigger) {
		return TriggerDefinitionInput("action["+strconv.Itoa(numAction)+"].trigger["+strconv.Itoa(numTrigger)+"]", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("action["+strconv.Itoa(numAction)+"].trigger["+strconv.Itoa(numTrigger)+"]", &resource.Action[numAction].Trigger[numTrigger], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionInput(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return DataRequirementInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"]", nil, htmlAttrs)
	}
	return DataRequirementInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"]", &resource.Action[numAction].Input[numInput], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionOutput(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return DataRequirementInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"]", nil, htmlAttrs)
	}
	return DataRequirementInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"]", &resource.Action[numAction].Output[numOutput], htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingDateTime(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return FhirDateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("action["+strconv.Itoa(numAction)+"].timingDateTime", resource.Action[numAction].TimingDateTime, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingAge(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return AgeInput("action["+strconv.Itoa(numAction)+"].timingAge", nil, htmlAttrs)
	}
	return AgeInput("action["+strconv.Itoa(numAction)+"].timingAge", resource.Action[numAction].TimingAge, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingPeriod(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return PeriodInput("action["+strconv.Itoa(numAction)+"].timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("action["+strconv.Itoa(numAction)+"].timingPeriod", resource.Action[numAction].TimingPeriod, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingDuration(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return DurationInput("action["+strconv.Itoa(numAction)+"].timingDuration", nil, htmlAttrs)
	}
	return DurationInput("action["+strconv.Itoa(numAction)+"].timingDuration", resource.Action[numAction].TimingDuration, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingRange(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return RangeInput("action["+strconv.Itoa(numAction)+"].timingRange", nil, htmlAttrs)
	}
	return RangeInput("action["+strconv.Itoa(numAction)+"].timingRange", resource.Action[numAction].TimingRange, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTimingTiming(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return TimingInput("action["+strconv.Itoa(numAction)+"].timingTiming", nil, htmlAttrs)
	}
	return TimingInput("action["+strconv.Itoa(numAction)+"].timingTiming", resource.Action[numAction].TimingTiming, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].type", resource.Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionGroupingBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_grouping_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].groupingBehavior", resource.Action[numAction].GroupingBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionSelectionBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_selection_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].selectionBehavior", resource.Action[numAction].SelectionBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRequiredBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_required_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].requiredBehavior", resource.Action[numAction].RequiredBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionPrecheckBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_precheck_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].precheckBehavior", resource.Action[numAction].PrecheckBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionCardinalityBehavior(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_cardinality_behavior

	if resource == nil || numAction >= len(resource.Action) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].cardinalityBehavior", resource.Action[numAction].CardinalityBehavior, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDefinitionCanonical(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].definitionCanonical", resource.Action[numAction].DefinitionCanonical, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDefinitionUri(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].definitionUri", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].definitionUri", resource.Action[numAction].DefinitionUri, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionTransform(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].transform", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].transform", resource.Action[numAction].Transform, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionConditionKind(numAction int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_condition_kind

	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].kind", &resource.Action[numAction].Condition[numCondition].Kind, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionConditionExpression(numAction int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numCondition >= len(resource.Action[numAction].Condition) {
		return ExpressionInput("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].expression", nil, htmlAttrs)
	}
	return ExpressionInput("action["+strconv.Itoa(numAction)+"].condition["+strconv.Itoa(numCondition)+"].expression", resource.Action[numAction].Condition[numCondition].Expression, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionActionId(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].actionId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].actionId", &resource.Action[numAction].RelatedAction[numRelatedAction].ActionId, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionOffsetDuration(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return DurationInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetDuration", nil, htmlAttrs)
	}
	return DurationInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetDuration", resource.Action[numAction].RelatedAction[numRelatedAction].OffsetDuration, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionOffsetRange(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return RangeInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetRange", nil, htmlAttrs)
	}
	return RangeInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].offsetRange", resource.Action[numAction].RelatedAction[numRelatedAction].OffsetRange, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantType(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", &resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDynamicValuePath(numAction int, numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", resource.Action[numAction].DynamicValue[numDynamicValue].Path, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDynamicValueExpression(numAction int, numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return ExpressionInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", nil, htmlAttrs)
	}
	return ExpressionInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", resource.Action[numAction].DynamicValue[numDynamicValue].Expression, htmlAttrs)
}
