package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/OperationOutcome
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

// http://hl7.org/fhir/r4b/StructureDefinition/OperationOutcome
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

func (resource *OperationOutcome) T_Id() templ.Component {

	if resource == nil {
		return StringInput("OperationOutcome.Id", nil)
	}
	return StringInput("OperationOutcome.Id", resource.Id)
}
func (resource *OperationOutcome) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("OperationOutcome.ImplicitRules", nil)
	}
	return StringInput("OperationOutcome.ImplicitRules", resource.ImplicitRules)
}
func (resource *OperationOutcome) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("OperationOutcome.Language", nil, optionsValueSet)
	}
	return CodeSelect("OperationOutcome.Language", resource.Language, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueId(numIssue int) templ.Component {

	if resource == nil || len(resource.Issue) >= numIssue {
		return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Id", nil)
	}
	return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Id", resource.Issue[numIssue].Id)
}
func (resource *OperationOutcome) T_IssueSeverity(numIssue int) templ.Component {
	optionsValueSet := VSIssue_severity

	if resource == nil || len(resource.Issue) >= numIssue {
		return CodeSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Severity", nil, optionsValueSet)
	}
	return CodeSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Severity", &resource.Issue[numIssue].Severity, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueCode(numIssue int) templ.Component {
	optionsValueSet := VSIssue_type

	if resource == nil || len(resource.Issue) >= numIssue {
		return CodeSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Code", &resource.Issue[numIssue].Code, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueDetails(numIssue int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Issue) >= numIssue {
		return CodeableConceptSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Details", nil, optionsValueSet)
	}
	return CodeableConceptSelect("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Details", resource.Issue[numIssue].Details, optionsValueSet)
}
func (resource *OperationOutcome) T_IssueDiagnostics(numIssue int) templ.Component {

	if resource == nil || len(resource.Issue) >= numIssue {
		return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Diagnostics", nil)
	}
	return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Diagnostics", resource.Issue[numIssue].Diagnostics)
}
func (resource *OperationOutcome) T_IssueLocation(numIssue int, numLocation int) templ.Component {

	if resource == nil || len(resource.Issue) >= numIssue || len(resource.Issue[numIssue].Location) >= numLocation {
		return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Location["+strconv.Itoa(numLocation)+"]", nil)
	}
	return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Location["+strconv.Itoa(numLocation)+"]", &resource.Issue[numIssue].Location[numLocation])
}
func (resource *OperationOutcome) T_IssueExpression(numIssue int, numExpression int) templ.Component {

	if resource == nil || len(resource.Issue) >= numIssue || len(resource.Issue[numIssue].Expression) >= numExpression {
		return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Expression["+strconv.Itoa(numExpression)+"]", nil)
	}
	return StringInput("OperationOutcome.Issue["+strconv.Itoa(numIssue)+"].Expression["+strconv.Itoa(numExpression)+"]", &resource.Issue[numIssue].Expression[numExpression])
}
