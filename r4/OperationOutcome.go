package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/OperationOutcome
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

// client create, read, etc return OperationOutcome err
// Error() includes function trace, String() does not
func (oo OperationOutcome) Error() string {
	return oo.StringWithErrorTrace()
}

// print each issue except internal fhir client error trace issues
func (oo OperationOutcome) String() string {
	var b strings.Builder
	for _, v := range oo.Issue {
		isDebug := false
		for _, coding := range v.Details.Coding {
			if coding.System != nil && *coding.System == "https://potatoemr.github.io/bultaoreune#debug" {
				isDebug = true
				break
			}
		}
		if !isDebug {
			b.WriteString(v.String())
			b.WriteString(" ")
		}
	}
	return b.String()
}

// includes fhir client error trace issues
func (oo OperationOutcome) StringWithErrorTrace() string {
	var b strings.Builder
	b.WriteString("OperationOutcome: ")
	for _, v := range oo.Issue {
		b.WriteString(v.String())
		b.WriteString(" ")
	}
	return b.String()
}

// http://hl7.org/fhir/r4/StructureDefinition/OperationOutcome
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

// struct -> json, automatically add resourceType=Patient
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
}

func (r OperationOutcome) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "OperationOutcome/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "OperationOutcome"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r OperationOutcome) ResourceType() string {
	return "OperationOutcome"
}

func (resource *OperationOutcome) T_IssueSeverity(numIssue int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSIssue_severity

	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeSelect("issue["+strconv.Itoa(numIssue)+"].severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("issue["+strconv.Itoa(numIssue)+"].severity", &resource.Issue[numIssue].Severity, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueCode(numIssue int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSIssue_type

	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeSelect("issue["+strconv.Itoa(numIssue)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("issue["+strconv.Itoa(numIssue)+"].code", &resource.Issue[numIssue].Code, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueDetails(numIssue int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeableConceptSelect("issue["+strconv.Itoa(numIssue)+"].details", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("issue["+strconv.Itoa(numIssue)+"].details", resource.Issue[numIssue].Details, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueDiagnostics(numIssue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIssue >= len(resource.Issue) {
		return StringInput("issue["+strconv.Itoa(numIssue)+"].diagnostics", nil, htmlAttrs)
	}
	return StringInput("issue["+strconv.Itoa(numIssue)+"].diagnostics", resource.Issue[numIssue].Diagnostics, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueLocation(numIssue int, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIssue >= len(resource.Issue) || numLocation >= len(resource.Issue[numIssue].Location) {
		return StringInput("issue["+strconv.Itoa(numIssue)+"].location["+strconv.Itoa(numLocation)+"]", nil, htmlAttrs)
	}
	return StringInput("issue["+strconv.Itoa(numIssue)+"].location["+strconv.Itoa(numLocation)+"]", &resource.Issue[numIssue].Location[numLocation], htmlAttrs)
}
func (resource *OperationOutcome) T_IssueExpression(numIssue int, numExpression int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIssue >= len(resource.Issue) || numExpression >= len(resource.Issue[numIssue].Expression) {
		return StringInput("issue["+strconv.Itoa(numIssue)+"].expression["+strconv.Itoa(numExpression)+"]", nil, htmlAttrs)
	}
	return StringInput("issue["+strconv.Itoa(numIssue)+"].expression["+strconv.Itoa(numExpression)+"]", &resource.Issue[numIssue].Expression[numExpression], htmlAttrs)
}
