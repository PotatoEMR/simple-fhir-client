package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ResearchSubject
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
func (resource *ResearchSubject) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResearch_subject_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchSubject) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *ResearchSubject) T_Study(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "study", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "study", &resource.Study, htmlAttrs)
}
func (resource *ResearchSubject) T_Individual(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "individual", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "individual", &resource.Individual, htmlAttrs)
}
func (resource *ResearchSubject) T_AssignedArm(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("assignedArm", nil, htmlAttrs)
	}
	return StringInput("assignedArm", resource.AssignedArm, htmlAttrs)
}
func (resource *ResearchSubject) T_ActualArm(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("actualArm", nil, htmlAttrs)
	}
	return StringInput("actualArm", resource.ActualArm, htmlAttrs)
}
func (resource *ResearchSubject) T_Consent(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "consent", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "consent", resource.Consent, htmlAttrs)
}
