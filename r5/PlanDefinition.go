package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                      *string                `json:"id,omitempty"`
	Meta                    *Meta                  `json:"meta,omitempty"`
	ImplicitRules           *string                `json:"implicitRules,omitempty"`
	Language                *string                `json:"language,omitempty"`
	Text                    *Narrative             `json:"text,omitempty"`
	Contained               []Resource             `json:"contained,omitempty"`
	Extension               []Extension            `json:"extension,omitempty"`
	ModifierExtension       []Extension            `json:"modifierExtension,omitempty"`
	Url                     *string                `json:"url,omitempty"`
	Identifier              []Identifier           `json:"identifier,omitempty"`
	Version                 *string                `json:"version,omitempty"`
	VersionAlgorithmString  *string                `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding  *Coding                `json:"versionAlgorithmCoding,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Title                   *string                `json:"title,omitempty"`
	Subtitle                *string                `json:"subtitle,omitempty"`
	Type                    *CodeableConcept       `json:"type,omitempty"`
	Status                  string                 `json:"status"`
	Experimental            *bool                  `json:"experimental,omitempty"`
	SubjectCodeableConcept  *CodeableConcept       `json:"subjectCodeableConcept,omitempty"`
	SubjectReference        *Reference             `json:"subjectReference,omitempty"`
	SubjectCanonical        *string                `json:"subjectCanonical,omitempty"`
	Date                    *string                `json:"date,omitempty"`
	Publisher               *string                `json:"publisher,omitempty"`
	Contact                 []ContactDetail        `json:"contact,omitempty"`
	Description             *string                `json:"description,omitempty"`
	UseContext              []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction            []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                 *string                `json:"purpose,omitempty"`
	Usage                   *string                `json:"usage,omitempty"`
	Copyright               *string                `json:"copyright,omitempty"`
	CopyrightLabel          *string                `json:"copyrightLabel,omitempty"`
	ApprovalDate            *string                `json:"approvalDate,omitempty"`
	LastReviewDate          *string                `json:"lastReviewDate,omitempty"`
	EffectivePeriod         *Period                `json:"effectivePeriod,omitempty"`
	Topic                   []CodeableConcept      `json:"topic,omitempty"`
	Author                  []ContactDetail        `json:"author,omitempty"`
	Editor                  []ContactDetail        `json:"editor,omitempty"`
	Reviewer                []ContactDetail        `json:"reviewer,omitempty"`
	Endorser                []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact         []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Library                 []string               `json:"library,omitempty"`
	Goal                    []PlanDefinitionGoal   `json:"goal,omitempty"`
	Actor                   []PlanDefinitionActor  `json:"actor,omitempty"`
	Action                  []PlanDefinitionAction `json:"action,omitempty"`
	AsNeededBoolean         *bool                  `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept       `json:"asNeededCodeableConcept,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionGoalTarget struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range           `json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
	DetailString          *string          `json:"detailString,omitempty"`
	DetailBoolean         *bool            `json:"detailBoolean,omitempty"`
	DetailInteger         *int             `json:"detailInteger,omitempty"`
	DetailRatio           *Ratio           `json:"detailRatio,omitempty"`
	Due                   *Duration        `json:"due,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActor struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Title             *string                     `json:"title,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Option            []PlanDefinitionActorOption `json:"option"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActorOption struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *string          `json:"type,omitempty"`
	TypeCanonical     *string          `json:"typeCanonical,omitempty"`
	TypeReference     *Reference       `json:"typeReference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionAction struct {
	Id                     *string                             `json:"id,omitempty"`
	Extension              []Extension                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                         `json:"modifierExtension,omitempty"`
	LinkId                 *string                             `json:"linkId,omitempty"`
	Prefix                 *string                             `json:"prefix,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	TextEquivalent         *string                             `json:"textEquivalent,omitempty"`
	Priority               *string                             `json:"priority,omitempty"`
	Code                   *CodeableConcept                    `json:"code,omitempty"`
	Reason                 []CodeableConcept                   `json:"reason,omitempty"`
	Documentation          []RelatedArtifact                   `json:"documentation,omitempty"`
	GoalId                 []string                            `json:"goalId,omitempty"`
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                          `json:"subjectReference,omitempty"`
	SubjectCanonical       *string                             `json:"subjectCanonical,omitempty"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []PlanDefinitionActionInput         `json:"input,omitempty"`
	Output                 []PlanDefinitionActionOutput        `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingAge              *Age                                `json:"timingAge,omitempty"`
	TimingDuration         *Duration                           `json:"timingDuration,omitempty"`
	TimingRange            *Range                              `json:"timingRange,omitempty"`
	TimingTiming           *Timing                             `json:"timingTiming,omitempty"`
	Location               *CodeableReference                  `json:"location,omitempty"`
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

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionCondition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Kind              string      `json:"kind"`
	Expression        *Expression `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionInput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionOutput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Requirement       *DataRequirement `json:"requirement,omitempty"`
	RelatedData       *string          `json:"relatedData,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionRelatedAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	TargetId          string      `json:"targetId"`
	Relationship      string      `json:"relationship"`
	EndRelationship   *string     `json:"endRelationship,omitempty"`
	OffsetDuration    *Duration   `json:"offsetDuration,omitempty"`
	OffsetRange       *Range      `json:"offsetRange,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
type PlanDefinitionActionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	ActorId           *string          `json:"actorId,omitempty"`
	Type              *string          `json:"type,omitempty"`
	TypeCanonical     *string          `json:"typeCanonical,omitempty"`
	TypeReference     *Reference       `json:"typeReference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PlanDefinition
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
func (resource *PlanDefinition) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *PlanDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
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
func (resource *PlanDefinition) T_SubjectCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subjectCanonical", nil, htmlAttrs)
	}
	return StringInput("subjectCanonical", resource.SubjectCanonical, htmlAttrs)
}
func (resource *PlanDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *PlanDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *PlanDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
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
func (resource *PlanDefinition) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *PlanDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *PlanDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *PlanDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_Library(numLibrary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *PlanDefinition) T_AsNeededBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("asNeededBoolean", nil, htmlAttrs)
	}
	return BoolInput("asNeededBoolean", resource.AsNeededBoolean, htmlAttrs)
}
func (resource *PlanDefinition) T_AsNeededCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("asNeededCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("asNeededCodeableConcept", resource.AsNeededCodeableConcept, optionsValueSet, htmlAttrs)
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
func (resource *PlanDefinition) T_GoalTargetMeasure(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].measure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].measure", resource.Goal[numGoal].Target[numTarget].Measure, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailCodeableConcept(numGoal int, numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", resource.Goal[numGoal].Target[numTarget].DetailCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailString(numGoal int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return StringInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailString", nil, htmlAttrs)
	}
	return StringInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailString", resource.Goal[numGoal].Target[numTarget].DetailString, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailBoolean(numGoal int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return BoolInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailBoolean", nil, htmlAttrs)
	}
	return BoolInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailBoolean", resource.Goal[numGoal].Target[numTarget].DetailBoolean, htmlAttrs)
}
func (resource *PlanDefinition) T_GoalTargetDetailInteger(numGoal int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) || numTarget >= len(resource.Goal[numGoal].Target) {
		return IntInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailInteger", nil, htmlAttrs)
	}
	return IntInput("goal["+strconv.Itoa(numGoal)+"].target["+strconv.Itoa(numTarget)+"].detailInteger", resource.Goal[numGoal].Target[numTarget].DetailInteger, htmlAttrs)
}
func (resource *PlanDefinition) T_ActorTitle(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].title", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].title", resource.Actor[numActor].Title, htmlAttrs)
}
func (resource *PlanDefinition) T_ActorDescription(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].description", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].description", resource.Actor[numActor].Description, htmlAttrs)
}
func (resource *PlanDefinition) T_ActorOptionType(numActor int, numOption int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numActor >= len(resource.Actor) || numOption >= len(resource.Actor[numActor].Option) {
		return CodeSelect("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].type", resource.Actor[numActor].Option[numOption].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActorOptionTypeCanonical(numActor int, numOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) || numOption >= len(resource.Actor[numActor].Option) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].typeCanonical", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].typeCanonical", resource.Actor[numActor].Option[numOption].TypeCanonical, htmlAttrs)
}
func (resource *PlanDefinition) T_ActorOptionRole(numActor int, numOption int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) || numOption >= len(resource.Actor[numActor].Option) {
		return CodeableConceptSelect("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("actor["+strconv.Itoa(numActor)+"].option["+strconv.Itoa(numOption)+"].role", resource.Actor[numActor].Option[numOption].Role, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionLinkId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].linkId", resource.Action[numAction].LinkId, htmlAttrs)
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
func (resource *PlanDefinition) T_ActionCode(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].code", resource.Action[numAction].Code, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionReason(numAction int, numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numReason >= len(resource.Action[numAction].Reason) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", &resource.Action[numAction].Reason[numReason], optionsValueSet, htmlAttrs)
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
func (resource *PlanDefinition) T_ActionSubjectCanonical(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) {
		return StringInput("action["+strconv.Itoa(numAction)+"].subjectCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].subjectCanonical", resource.Action[numAction].SubjectCanonical, htmlAttrs)
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
func (resource *PlanDefinition) T_ActionInputTitle(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].title", resource.Action[numAction].Input[numInput].Title, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionInputRelatedData(numAction int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numInput >= len(resource.Action[numAction].Input) {
		return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].relatedData", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].input["+strconv.Itoa(numInput)+"].relatedData", resource.Action[numAction].Input[numInput].RelatedData, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionOutputTitle(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].title", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].title", resource.Action[numAction].Output[numOutput].Title, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionOutputRelatedData(numAction int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numOutput >= len(resource.Action[numAction].Output) {
		return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].relatedData", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].output["+strconv.Itoa(numOutput)+"].relatedData", resource.Action[numAction].Output[numOutput].RelatedData, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionTargetId(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].targetId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].targetId", &resource.Action[numAction].RelatedAction[numRelatedAction].TargetId, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].relationship", &resource.Action[numAction].RelatedAction[numRelatedAction].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionRelatedActionEndRelationship(numAction int, numRelatedAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_relationship_type

	if resource == nil || numAction >= len(resource.Action) || numRelatedAction >= len(resource.Action[numAction].RelatedAction) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].endRelationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].relatedAction["+strconv.Itoa(numRelatedAction)+"].endRelationship", resource.Action[numAction].RelatedAction[numRelatedAction].EndRelationship, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantActorId(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorId", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].actorId", resource.Action[numAction].Participant[numParticipant].ActorId, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantType(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].type", resource.Action[numAction].Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantTypeCanonical(numAction int, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeCanonical", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].typeCanonical", resource.Action[numAction].Participant[numParticipant].TypeCanonical, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantRole(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].role", resource.Action[numAction].Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionParticipantFunction(numAction int, numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numParticipant >= len(resource.Action[numAction].Participant) {
		return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("action["+strconv.Itoa(numAction)+"].participant["+strconv.Itoa(numParticipant)+"].function", resource.Action[numAction].Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *PlanDefinition) T_ActionDynamicValuePath(numAction int, numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Action) || numDynamicValue >= len(resource.Action[numAction].DynamicValue) {
		return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", nil, htmlAttrs)
	}
	return StringInput("action["+strconv.Itoa(numAction)+"].dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", resource.Action[numAction].DynamicValue[numDynamicValue].Path, htmlAttrs)
}
