//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"
import "strings"

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

func (oo OperationOutcome) String() string {
	var b strings.Builder
	b.WriteString("OperationOutcome ")
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
