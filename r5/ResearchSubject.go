package r5

//generated with command go run ./bultaoreune
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
func (resource *ResearchSubject) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_AssignedComparisonGroup(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("assignedComparisonGroup", nil, htmlAttrs)
	}
	return StringInput("assignedComparisonGroup", resource.AssignedComparisonGroup, htmlAttrs)
}
func (resource *ResearchSubject) T_ActualComparisonGroup(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("actualComparisonGroup", nil, htmlAttrs)
	}
	return StringInput("actualComparisonGroup", resource.ActualComparisonGroup, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressType(numProgress int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].type", resource.Progress[numProgress].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressSubjectState(numProgress int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].subjectState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].subjectState", resource.Progress[numProgress].SubjectState, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressMilestone(numProgress int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].milestone", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].milestone", resource.Progress[numProgress].Milestone, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressReason(numProgress int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("progress["+strconv.Itoa(numProgress)+"].reason", resource.Progress[numProgress].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressStartDate(numProgress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return DateTimeInput("progress["+strconv.Itoa(numProgress)+"].startDate", nil, htmlAttrs)
	}
	return DateTimeInput("progress["+strconv.Itoa(numProgress)+"].startDate", resource.Progress[numProgress].StartDate, htmlAttrs)
}
func (resource *ResearchSubject) T_ProgressEndDate(numProgress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgress >= len(resource.Progress) {
		return DateTimeInput("progress["+strconv.Itoa(numProgress)+"].endDate", nil, htmlAttrs)
	}
	return DateTimeInput("progress["+strconv.Itoa(numProgress)+"].endDate", resource.Progress[numProgress].EndDate, htmlAttrs)
}
