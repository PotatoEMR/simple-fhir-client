package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *ResearchSubject) ResearchSubjectLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectProgressType(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Progress[numProgress].Type, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectProgressSubjectState(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("subjectState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subjectState", resource.Progress[numProgress].SubjectState, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectProgressMilestone(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("milestone", nil, optionsValueSet)
	}
	return CodeableConceptSelect("milestone", resource.Progress[numProgress].Milestone, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectProgressReason(numProgress int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Progress) >= numProgress {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", resource.Progress[numProgress].Reason, optionsValueSet)
}
