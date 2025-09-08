package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

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
func (resource *OperationOutcome) T_IssueSeverity(numIssue int, htmlAttrs string) templ.Component {
	optionsValueSet := VSIssue_severity

	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Severity", &resource.Issue[numIssue].Severity, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueCode(numIssue int, htmlAttrs string) templ.Component {
	optionsValueSet := VSIssue_type

	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Code", &resource.Issue[numIssue].Code, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueDetails(numIssue int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIssue >= len(resource.Issue) {
		return CodeableConceptSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Details", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Details", resource.Issue[numIssue].Details, optionsValueSet, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueDiagnostics(numIssue int, htmlAttrs string) templ.Component {

	if resource == nil || numIssue >= len(resource.Issue) {
		return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Diagnostics", nil, htmlAttrs)
	}
	return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Diagnostics", resource.Issue[numIssue].Diagnostics, htmlAttrs)
}
func (resource *OperationOutcome) T_IssueLocation(numIssue int, numLocation int, htmlAttrs string) templ.Component {

	if resource == nil || numIssue >= len(resource.Issue) || numLocation >= len(resource.Issue[numIssue].Location) {
		return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Location."+strconv.Itoa(numLocation)+".", nil, htmlAttrs)
	}
	return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Location."+strconv.Itoa(numLocation)+".", &resource.Issue[numIssue].Location[numLocation], htmlAttrs)
}
func (resource *OperationOutcome) T_IssueExpression(numIssue int, numExpression int, htmlAttrs string) templ.Component {

	if resource == nil || numIssue >= len(resource.Issue) || numExpression >= len(resource.Issue[numIssue].Expression) {
		return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Expression."+strconv.Itoa(numExpression)+".", nil, htmlAttrs)
	}
	return StringInput("OperationOutcome.Issue."+strconv.Itoa(numIssue)+"..Expression."+strconv.Itoa(numExpression)+".", &resource.Issue[numIssue].Expression[numExpression], htmlAttrs)
}
