package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	StartDate            *string           `json:"startDate"`
	StartCodeableConcept *CodeableConcept  `json:"startCodeableConcept"`
	Target               []GoalTarget      `json:"target,omitempty"`
	StatusDate           *string           `json:"statusDate,omitempty"`
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
	DetailQuantity        *Quantity        `json:"detailQuantity"`
	DetailRange           *Range           `json:"detailRange"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept"`
	DetailString          *string          `json:"detailString"`
	DetailBoolean         *bool            `json:"detailBoolean"`
	DetailInteger         *int             `json:"detailInteger"`
	DetailRatio           *Ratio           `json:"detailRatio"`
	DueDate               *string          `json:"dueDate"`
	DueDuration           *Duration        `json:"dueDuration"`
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

func (resource *Goal) GoalLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Goal) GoalLifecycleStatus() templ.Component {
	optionsValueSet := VSGoal_status

	if resource != nil {
		return CodeSelect("lifecycleStatus", nil, optionsValueSet)
	}
	return CodeSelect("lifecycleStatus", &resource.LifecycleStatus, optionsValueSet)
}
func (resource *Goal) GoalAchievementStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("achievementStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("achievementStatus", resource.AchievementStatus, optionsValueSet)
}
func (resource *Goal) GoalCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Goal) GoalPriority(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Goal) GoalDescription(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("description", nil, optionsValueSet)
	}
	return CodeableConceptSelect("description", &resource.Description, optionsValueSet)
}
func (resource *Goal) GoalOutcomeCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("outcomeCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcomeCode", &resource.OutcomeCode[0], optionsValueSet)
}
func (resource *Goal) GoalTargetMeasure(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("measure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("measure", resource.Target[numTarget].Measure, optionsValueSet)
}
