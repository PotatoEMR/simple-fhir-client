package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReport struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        *Identifier             `json:"identifier,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Status            string                  `json:"status"`
	TestScript        Reference               `json:"testScript"`
	Result            string                  `json:"result"`
	Score             *float64                `json:"score,omitempty"`
	Tester            *string                 `json:"tester,omitempty"`
	Issued            *string                 `json:"issued,omitempty"`
	Participant       []TestReportParticipant `json:"participant,omitempty"`
	Setup             *TestReportSetup        `json:"setup,omitempty"`
	Test              []TestReportTest        `json:"test,omitempty"`
	Teardown          *TestReportTeardown     `json:"teardown,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportParticipant struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Uri               string      `json:"uri"`
	Display           *string     `json:"display,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportSetupActionOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Result            string      `json:"result"`
	Message           *string     `json:"message,omitempty"`
	Detail            *string     `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportSetupActionAssert struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Result            string      `json:"result"`
	Message           *string     `json:"message,omitempty"`
	Detail            *string     `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestReportTestAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestReport
type TestReportTeardownAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

type OtherTestReport TestReport

// on convert struct to json, automatically add resourceType=TestReport
func (r TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestReport
		ResourceType string `json:"resourceType"`
	}{
		OtherTestReport: OtherTestReport(r),
		ResourceType:    "TestReport",
	})
}
func (resource *TestReport) TestReportLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *TestReport) TestReportStatus() templ.Component {
	optionsValueSet := VSReport_status_codes
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *TestReport) TestReportResult() templ.Component {
	optionsValueSet := VSReport_result_codes
	currentVal := ""
	if resource != nil {
		currentVal = resource.Result
	}
	return CodeSelect("result", currentVal, optionsValueSet)
}
func (resource *TestReport) TestReportParticipantType(numParticipant int) templ.Component {
	optionsValueSet := VSReport_participant_type
	currentVal := ""
	if resource != nil && len(resource.Participant) >= numParticipant {
		currentVal = resource.Participant[numParticipant].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *TestReport) TestReportSetupActionOperationResult(numAction int) templ.Component {
	optionsValueSet := VSReport_action_result_codes
	currentVal := ""
	if resource != nil && len(resource.Setup.Action) >= numAction {
		currentVal = resource.Setup.Action[numAction].Operation.Result
	}
	return CodeSelect("result", currentVal, optionsValueSet)
}
func (resource *TestReport) TestReportSetupActionAssertResult(numAction int) templ.Component {
	optionsValueSet := VSReport_action_result_codes
	currentVal := ""
	if resource != nil && len(resource.Setup.Action) >= numAction {
		currentVal = resource.Setup.Action[numAction].Assert.Result
	}
	return CodeSelect("result", currentVal, optionsValueSet)
}
