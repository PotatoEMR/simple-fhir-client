package r4b

//generated with command go run ./bultaoreune
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
	optionsValueSet := VSResearch_subject_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_AssignedArm(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AssignedArm", nil, htmlAttrs)
	}
	return StringInput("AssignedArm", resource.AssignedArm, htmlAttrs)
}
func (resource *ResearchSubject) T_ActualArm(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActualArm", nil, htmlAttrs)
	}
	return StringInput("ActualArm", resource.ActualArm, htmlAttrs)
}
