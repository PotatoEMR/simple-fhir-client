package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ResearchSubject) ResearchSubjectStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
