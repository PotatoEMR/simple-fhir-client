package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                      *string                   `json:"id,omitempty"`
	Meta                    *Meta                     `json:"meta,omitempty"`
	ImplicitRules           *string                   `json:"implicitRules,omitempty"`
	Language                *string                   `json:"language,omitempty"`
	Text                    *Narrative                `json:"text,omitempty"`
	Contained               []Resource                `json:"contained,omitempty"`
	Extension               []Extension               `json:"extension,omitempty"`
	ModifierExtension       []Extension               `json:"modifierExtension,omitempty"`
	Identifier              []Identifier              `json:"identifier,omitempty"`
	Status                  string                    `json:"status"`
	Progress                []ResearchSubjectProgress `json:"progress,omitempty"`
	Period                  *Period                   `json:"period,omitempty"`
	Study                   Reference                 `json:"study"`
	Subject                 Reference                 `json:"subject"`
	AssignedComparisonGroup *string                   `json:"assignedComparisonGroup,omitempty"`
	ActualComparisonGroup   *string                   `json:"actualComparisonGroup,omitempty"`
	Consent                 []Reference               `json:"consent,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ResearchSubject
type ResearchSubjectProgress struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	SubjectState      *CodeableConcept `json:"subjectState,omitempty"`
	Milestone         *CodeableConcept `json:"milestone,omitempty"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	StartDate         *string          `json:"startDate,omitempty"`
	EndDate           *string          `json:"endDate,omitempty"`
}

type OtherResearchSubject ResearchSubject

// on convert struct to json, automatically add resourceType=ResearchSubject
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchSubject: OtherResearchSubject(r),
		ResourceType:         "ResearchSubject",
	})
}

func (resource *ResearchSubject) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.Id", nil)
	}
	return StringInput("ResearchSubject.Id", resource.Id)
}
func (resource *ResearchSubject) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.ImplicitRules", nil)
	}
	return StringInput("ResearchSubject.ImplicitRules", resource.ImplicitRules)
}
func (resource *ResearchSubject) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ResearchSubject.Language", nil, optionsValueSet)
	}
	return CodeSelect("ResearchSubject.Language", resource.Language, optionsValueSet)
}
func (resource *ResearchSubject) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ResearchSubject.Status", nil, optionsValueSet)
	}
	return CodeSelect("ResearchSubject.Status", &resource.Status, optionsValueSet)
}
func (resource *ResearchSubject) T_AssignedComparisonGroup() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.AssignedComparisonGroup", nil)
	}
	return StringInput("ResearchSubject.AssignedComparisonGroup", resource.AssignedComparisonGroup)
}
func (resource *ResearchSubject) T_ActualComparisonGroup() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.ActualComparisonGroup", nil)
	}
	return StringInput("ResearchSubject.ActualComparisonGroup", resource.ActualComparisonGroup)
}
func (resource *ResearchSubject) T_ProgressId(numProgress int) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Id", nil)
	}
	return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Id", resource.Progress[numProgress].Id)
}
func (resource *ResearchSubject) T_ProgressType(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Type", resource.Progress[numProgress].Type, optionsValueSet)
}
func (resource *ResearchSubject) T_ProgressSubjectState(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].SubjectState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].SubjectState", resource.Progress[numProgress].SubjectState, optionsValueSet)
}
func (resource *ResearchSubject) T_ProgressMilestone(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Milestone", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Milestone", resource.Progress[numProgress].Milestone, optionsValueSet)
}
func (resource *ResearchSubject) T_ProgressReason(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].Reason", resource.Progress[numProgress].Reason, optionsValueSet)
}
func (resource *ResearchSubject) T_ProgressStartDate(numProgress int) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].StartDate", nil)
	}
	return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].StartDate", resource.Progress[numProgress].StartDate)
}
func (resource *ResearchSubject) T_ProgressEndDate(numProgress int) templ.Component {

	if resource == nil || len(resource.Progress) >= numProgress {
		return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].EndDate", nil)
	}
	return StringInput("ResearchSubject.Progress["+strconv.Itoa(numProgress)+"].EndDate", resource.Progress[numProgress].EndDate)
}
