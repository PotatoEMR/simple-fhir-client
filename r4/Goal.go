//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/Goal
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

// http://hl7.org/fhir/r4/StructureDefinition/Goal
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
