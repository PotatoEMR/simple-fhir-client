package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	StartDate         *time.Time       `json:"startDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	EndDate           *time.Time       `json:"endDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ResearchSubject) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ResearchSubject/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ResearchSubject"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ResearchSubject) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ResearchSubject.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ResearchSubject.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_AssignedComparisonGroup(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.AssignedComparisonGroup", nil, htmlAttrs)
	}
	return StringInput("ResearchSubject.AssignedComparisonGroup", resource.AssignedComparisonGroup, htmlAttrs)
}
func (resource *ResearchSubject) T_ActualComparisonGroup(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.ActualComparisonGroup", nil, htmlAttrs)
	}
	return StringInput("ResearchSubject.ActualComparisonGroup", resource.ActualComparisonGroup, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressType(numProgress int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Type", resource.Progress[numProgress].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressSubjectState(numProgress int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..SubjectState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..SubjectState", resource.Progress[numProgress].SubjectState, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressMilestone(numProgress int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Milestone", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Milestone", resource.Progress[numProgress].Milestone, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressReason(numProgress int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..Reason", resource.Progress[numProgress].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressStartDate(numProgress int, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return DateTimeInput("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..StartDate", nil, htmlAttrs)
	}
	return DateTimeInput("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..StartDate", resource.Progress[numProgress].StartDate, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressEndDate(numProgress int, htmlAttrs string) templ.Component {

	if resource == nil || numProgress >= len(resource.Progress) {
		return DateTimeInput("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..EndDate", nil, htmlAttrs)
	}
	return DateTimeInput("ResearchSubject.Progress."+strconv.Itoa(numProgress)+"..EndDate", resource.Progress[numProgress].EndDate, htmlAttrs)
}
