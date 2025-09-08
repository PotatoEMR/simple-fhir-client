package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Goal
type Goal struct {
	Id                   *string           `json:"id,omitempty"`
	Meta                 *Meta             `json:"meta,omitempty"`
	ImplicitRules        *string           `json:"implicitRules,omitempty"`
	Language             *string           `json:"language,omitempty"`
	Text                 *Narrative        `json:"text,omitempty"`
	Contained            []Resource        `json:"contained,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `json:"identifier,omitempty"`
	LifecycleStatus      string            `json:"lifecycleStatus"`
	AchievementStatus    *CodeableConcept  `json:"achievementStatus,omitempty"`
	Category             []CodeableConcept `json:"category,omitempty"`
	Priority             *CodeableConcept  `json:"priority,omitempty"`
	Description          CodeableConcept   `json:"description"`
	Subject              Reference         `json:"subject"`
	StartDate            *time.Time        `json:"startDate,omitempty,format:'2006-01-02'"`
	StartCodeableConcept *CodeableConcept  `json:"startCodeableConcept,omitempty"`
	Target               []GoalTarget      `json:"target,omitempty"`
	StatusDate           *time.Time        `json:"statusDate,omitempty,format:'2006-01-02'"`
	StatusReason         *string           `json:"statusReason,omitempty"`
	ExpressedBy          *Reference        `json:"expressedBy,omitempty"`
	Addresses            []Reference       `json:"addresses,omitempty"`
	Note                 []Annotation      `json:"note,omitempty"`
	OutcomeCode          []CodeableConcept `json:"outcomeCode,omitempty"`
	OutcomeReference     []Reference       `json:"outcomeReference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Goal
type GoalTarget struct {
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
	DueDate               *time.Time       `json:"dueDate,omitempty,format:'2006-01-02'"`
	DueDuration           *Duration        `json:"dueDuration,omitempty"`
}

type OtherGoal Goal

// on convert struct to json, automatically add resourceType=Goal
func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}
func (r Goal) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Goal/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Goal"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Goal) T_LifecycleStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSGoal_status

	if resource == nil {
		return CodeSelect("LifecycleStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("LifecycleStatus", &resource.LifecycleStatus, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_AchievementStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("AchievementStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AchievementStatus", resource.AchievementStatus, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Description(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Description", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Description", &resource.Description, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_StartDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("StartDate", nil, htmlAttrs)
	}
	return DateInput("StartDate", resource.StartDate, htmlAttrs)
}
func (resource *Goal) T_StartCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StartCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StartCodeableConcept", resource.StartCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_StatusDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("StatusDate", nil, htmlAttrs)
	}
	return DateInput("StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *Goal) T_StatusReason(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("StatusReason", nil, htmlAttrs)
	}
	return StringInput("StatusReason", resource.StatusReason, htmlAttrs)
}
func (resource *Goal) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Goal) T_OutcomeCode(numOutcomeCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutcomeCode >= len(resource.OutcomeCode) {
		return CodeableConceptSelect("OutcomeCode["+strconv.Itoa(numOutcomeCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OutcomeCode["+strconv.Itoa(numOutcomeCode)+"]", &resource.OutcomeCode[numOutcomeCode], optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetMeasure(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("Target["+strconv.Itoa(numTarget)+"]Measure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Target["+strconv.Itoa(numTarget)+"]Measure", resource.Target[numTarget].Measure, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetDetailCodeableConcept(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("Target["+strconv.Itoa(numTarget)+"]DetailCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Target["+strconv.Itoa(numTarget)+"]DetailCodeableConcept", resource.Target[numTarget].DetailCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetDetailString(numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return StringInput("Target["+strconv.Itoa(numTarget)+"]DetailString", nil, htmlAttrs)
	}
	return StringInput("Target["+strconv.Itoa(numTarget)+"]DetailString", resource.Target[numTarget].DetailString, htmlAttrs)
}
func (resource *Goal) T_TargetDetailBoolean(numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return BoolInput("Target["+strconv.Itoa(numTarget)+"]DetailBoolean", nil, htmlAttrs)
	}
	return BoolInput("Target["+strconv.Itoa(numTarget)+"]DetailBoolean", resource.Target[numTarget].DetailBoolean, htmlAttrs)
}
func (resource *Goal) T_TargetDetailInteger(numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return IntInput("Target["+strconv.Itoa(numTarget)+"]DetailInteger", nil, htmlAttrs)
	}
	return IntInput("Target["+strconv.Itoa(numTarget)+"]DetailInteger", resource.Target[numTarget].DetailInteger, htmlAttrs)
}
func (resource *Goal) T_TargetDueDate(numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return DateInput("Target["+strconv.Itoa(numTarget)+"]DueDate", nil, htmlAttrs)
	}
	return DateInput("Target["+strconv.Itoa(numTarget)+"]DueDate", resource.Target[numTarget].DueDate, htmlAttrs)
}
