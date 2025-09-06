package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Goal
type Goal struct {
	Id                   *string             `json:"id,omitempty"`
	Meta                 *Meta               `json:"meta,omitempty"`
	ImplicitRules        *string             `json:"implicitRules,omitempty"`
	Language             *string             `json:"language,omitempty"`
	Text                 *Narrative          `json:"text,omitempty"`
	Contained            []Resource          `json:"contained,omitempty"`
	Extension            []Extension         `json:"extension,omitempty"`
	ModifierExtension    []Extension         `json:"modifierExtension,omitempty"`
	Identifier           []Identifier        `json:"identifier,omitempty"`
	LifecycleStatus      string              `json:"lifecycleStatus"`
	AchievementStatus    *CodeableConcept    `json:"achievementStatus,omitempty"`
	Category             []CodeableConcept   `json:"category,omitempty"`
	Continuous           *bool               `json:"continuous,omitempty"`
	Priority             *CodeableConcept    `json:"priority,omitempty"`
	Description          CodeableConcept     `json:"description"`
	Subject              Reference           `json:"subject"`
	StartDate            *string             `json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept    `json:"startCodeableConcept,omitempty"`
	Target               []GoalTarget        `json:"target,omitempty"`
	StatusDate           *string             `json:"statusDate,omitempty"`
	StatusReason         *string             `json:"statusReason,omitempty"`
	Source               *Reference          `json:"source,omitempty"`
	Addresses            []Reference         `json:"addresses,omitempty"`
	Note                 []Annotation        `json:"note,omitempty"`
	Outcome              []CodeableReference `json:"outcome,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Goal
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
	DueDate               *string          `json:"dueDate,omitempty"`
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

func (resource *Goal) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Goal.Id", nil)
	}
	return StringInput("Goal.Id", resource.Id)
}
func (resource *Goal) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Goal.ImplicitRules", nil)
	}
	return StringInput("Goal.ImplicitRules", resource.ImplicitRules)
}
func (resource *Goal) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Goal.Language", nil, optionsValueSet)
	}
	return CodeSelect("Goal.Language", resource.Language, optionsValueSet)
}
func (resource *Goal) T_LifecycleStatus() templ.Component {
	optionsValueSet := VSGoal_status

	if resource == nil {
		return CodeSelect("Goal.LifecycleStatus", nil, optionsValueSet)
	}
	return CodeSelect("Goal.LifecycleStatus", &resource.LifecycleStatus, optionsValueSet)
}
func (resource *Goal) T_AchievementStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Goal.AchievementStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Goal.AchievementStatus", resource.AchievementStatus, optionsValueSet)
}
func (resource *Goal) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Goal.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Goal.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Goal) T_Continuous() templ.Component {

	if resource == nil {
		return BoolInput("Goal.Continuous", nil)
	}
	return BoolInput("Goal.Continuous", resource.Continuous)
}
func (resource *Goal) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Goal.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Goal.Priority", resource.Priority, optionsValueSet)
}
func (resource *Goal) T_Description(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Goal.Description", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Goal.Description", &resource.Description, optionsValueSet)
}
func (resource *Goal) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("Goal.StatusDate", nil)
	}
	return StringInput("Goal.StatusDate", resource.StatusDate)
}
func (resource *Goal) T_StatusReason() templ.Component {

	if resource == nil {
		return StringInput("Goal.StatusReason", nil)
	}
	return StringInput("Goal.StatusReason", resource.StatusReason)
}
func (resource *Goal) T_TargetId(numTarget int) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return StringInput("Goal.Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("Goal.Target["+strconv.Itoa(numTarget)+"].Id", resource.Target[numTarget].Id)
}
func (resource *Goal) T_TargetMeasure(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("Goal.Target["+strconv.Itoa(numTarget)+"].Measure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Goal.Target["+strconv.Itoa(numTarget)+"].Measure", resource.Target[numTarget].Measure, optionsValueSet)
}
