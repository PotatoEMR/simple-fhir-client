//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
