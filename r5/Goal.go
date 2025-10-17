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
	StartDate            *FhirDate           `json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept    `json:"startCodeableConcept,omitempty"`
	Target               []GoalTarget        `json:"target,omitempty"`
	StatusDate           *FhirDate           `json:"statusDate,omitempty"`
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
	DueDate               *FhirDate        `json:"dueDate,omitempty"`
	DueDuration           *Duration        `json:"dueDuration,omitempty"`
}

type OtherGoal Goal

// struct -> json, automatically add resourceType=Patient
func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// json -> struct, first reject if resourceType != Goal
func (r *Goal) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Goal" {
		return errors.New("resourceType not Goal")
	}
	return json.Unmarshal(data, (*OtherGoal)(r))
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
func (resource *Goal) T_LifecycleStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGoal_status

	if resource == nil {
		return CodeSelect("lifecycleStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("lifecycleStatus", &resource.LifecycleStatus, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_AchievementStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("achievementStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("achievementStatus", resource.AchievementStatus, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Continuous(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("continuous", nil, htmlAttrs)
	}
	return BoolInput("continuous", resource.Continuous, htmlAttrs)
}
func (resource *Goal) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Description(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("description", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("description", &resource.Description, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *Goal) T_StartDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("startDate", nil, htmlAttrs)
	}
	return FhirDateInput("startDate", resource.StartDate, htmlAttrs)
}
func (resource *Goal) T_StartCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("startCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("startCodeableConcept", resource.StartCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_StatusDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("statusDate", nil, htmlAttrs)
	}
	return FhirDateInput("statusDate", resource.StatusDate, htmlAttrs)
}
func (resource *Goal) T_StatusReason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("statusReason", nil, htmlAttrs)
	}
	return StringInput("statusReason", resource.StatusReason, htmlAttrs)
}
func (resource *Goal) T_Source(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source", resource.Source, htmlAttrs)
}
func (resource *Goal) T_Addresses(frs []FhirResource, numAddresses int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddresses >= len(resource.Addresses) {
		return ReferenceInput(frs, "addresses["+strconv.Itoa(numAddresses)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "addresses["+strconv.Itoa(numAddresses)+"]", &resource.Addresses[numAddresses], htmlAttrs)
}
func (resource *Goal) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Goal) T_Outcome(numOutcome int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutcome >= len(resource.Outcome) {
		return CodeableReferenceInput("outcome["+strconv.Itoa(numOutcome)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("outcome["+strconv.Itoa(numOutcome)+"]", &resource.Outcome[numOutcome], htmlAttrs)
}
func (resource *Goal) T_TargetMeasure(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].measure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].measure", resource.Target[numTarget].Measure, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetDetailQuantity(numTarget int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return QuantityInput("target["+strconv.Itoa(numTarget)+"].detailQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("target["+strconv.Itoa(numTarget)+"].detailQuantity", resource.Target[numTarget].DetailQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetDetailRange(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return RangeInput("target["+strconv.Itoa(numTarget)+"].detailRange", nil, htmlAttrs)
	}
	return RangeInput("target["+strconv.Itoa(numTarget)+"].detailRange", resource.Target[numTarget].DetailRange, htmlAttrs)
}
func (resource *Goal) T_TargetDetailCodeableConcept(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].detailCodeableConcept", resource.Target[numTarget].DetailCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Goal) T_TargetDetailString(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return StringInput("target["+strconv.Itoa(numTarget)+"].detailString", nil, htmlAttrs)
	}
	return StringInput("target["+strconv.Itoa(numTarget)+"].detailString", resource.Target[numTarget].DetailString, htmlAttrs)
}
func (resource *Goal) T_TargetDetailBoolean(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return BoolInput("target["+strconv.Itoa(numTarget)+"].detailBoolean", nil, htmlAttrs)
	}
	return BoolInput("target["+strconv.Itoa(numTarget)+"].detailBoolean", resource.Target[numTarget].DetailBoolean, htmlAttrs)
}
func (resource *Goal) T_TargetDetailInteger(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return IntInput("target["+strconv.Itoa(numTarget)+"].detailInteger", nil, htmlAttrs)
	}
	return IntInput("target["+strconv.Itoa(numTarget)+"].detailInteger", resource.Target[numTarget].DetailInteger, htmlAttrs)
}
func (resource *Goal) T_TargetDetailRatio(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return RatioInput("target["+strconv.Itoa(numTarget)+"].detailRatio", nil, htmlAttrs)
	}
	return RatioInput("target["+strconv.Itoa(numTarget)+"].detailRatio", resource.Target[numTarget].DetailRatio, htmlAttrs)
}
func (resource *Goal) T_TargetDueDate(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return FhirDateInput("target["+strconv.Itoa(numTarget)+"].dueDate", nil, htmlAttrs)
	}
	return FhirDateInput("target["+strconv.Itoa(numTarget)+"].dueDate", resource.Target[numTarget].DueDate, htmlAttrs)
}
func (resource *Goal) T_TargetDueDuration(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return DurationInput("target["+strconv.Itoa(numTarget)+"].dueDuration", nil, htmlAttrs)
	}
	return DurationInput("target["+strconv.Itoa(numTarget)+"].dueDuration", resource.Target[numTarget].DueDuration, htmlAttrs)
}
