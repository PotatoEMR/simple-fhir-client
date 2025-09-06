package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                *string      `json:"id,omitempty"`
	Meta              *Meta        `json:"meta,omitempty"`
	ImplicitRules     *string      `json:"implicitRules,omitempty"`
	Language          *string      `json:"language,omitempty"`
	Text              *Narrative   `json:"text,omitempty"`
	Contained         []Resource   `json:"contained,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier `json:"identifier,omitempty"`
	Status            string       `json:"status"`
	Period            *Period      `json:"period,omitempty"`
	Study             Reference    `json:"study"`
	Individual        Reference    `json:"individual"`
	AssignedArm       *string      `json:"assignedArm,omitempty"`
	ActualArm         *string      `json:"actualArm,omitempty"`
	Consent           *Reference   `json:"consent,omitempty"`
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
	optionsValueSet := VSResearch_subject_status

	if resource == nil {
		return CodeSelect("ResearchSubject.Status", nil, optionsValueSet)
	}
	return CodeSelect("ResearchSubject.Status", &resource.Status, optionsValueSet)
}
func (resource *ResearchSubject) T_AssignedArm() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.AssignedArm", nil)
	}
	return StringInput("ResearchSubject.AssignedArm", resource.AssignedArm)
}
func (resource *ResearchSubject) T_ActualArm() templ.Component {

	if resource == nil {
		return StringInput("ResearchSubject.ActualArm", nil)
	}
	return StringInput("ResearchSubject.ActualArm", resource.ActualArm)
}
