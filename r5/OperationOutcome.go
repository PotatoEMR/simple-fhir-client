package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"
import "strings"

// http://hl7.org/fhir/r5/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `json:"issue"`
}

func (oo OperationOutcome) String() string {
	var b strings.Builder
	b.WriteString("OperationOutcome ")
	for _, v := range oo.Issue {
		b.WriteString(v.String())
		b.WriteString(" ")
	}
	return b.String()
}

// http://hl7.org/fhir/r5/StructureDefinition/OperationOutcome
type OperationOutcomeIssue struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Severity          string           `json:"severity"`
	Code              string           `json:"code"`
	Details           *CodeableConcept `json:"details,omitempty"`
	Diagnostics       *string          `json:"diagnostics,omitempty"`
	Location          []string         `json:"location,omitempty"`
	Expression        []string         `json:"expression,omitempty"`
}

func (ooi OperationOutcomeIssue) String() string {
	ret := "Issue: "
	if ooi.Diagnostics != nil {
		ret += *ooi.Diagnostics
	}
	ret += " "
	errorPath := append(ooi.Location, ooi.Expression...) //maybe error fields? idk
	ret += strings.Join(errorPath, ", ")
	return ret
}

type OtherOperationOutcome OperationOutcome

// on convert struct to json, automatically add resourceType=OperationOutcome
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
}

func (resource *OperationOutcome) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueSeverity(numIssue int) templ.Component {
	optionsValueSet := VSIssue_severity

	if resource == nil && len(resource.Issue) >= numIssue {
		return CodeSelect("severity", nil, optionsValueSet)
	}
	return CodeSelect("severity", &resource.Issue[numIssue].Severity, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueCode(numIssue int) templ.Component {
	optionsValueSet := VSIssue_type

	if resource == nil && len(resource.Issue) >= numIssue {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.Issue[numIssue].Code, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueDetails(numIssue int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Issue) >= numIssue {
		return CodeableConceptSelect("details", nil, optionsValueSet)
	}
	return CodeableConceptSelect("details", resource.Issue[numIssue].Details, optionsValueSet)
}
