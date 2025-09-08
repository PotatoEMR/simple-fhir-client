package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
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
	Issued            *time.Time              `json:"issued,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Participant       []TestReportParticipant `json:"participant,omitempty"`
	Setup             *TestReportSetup        `json:"setup,omitempty"`
	Test              []TestReportTest        `json:"test,omitempty"`
	Teardown          *TestReportTeardown     `json:"teardown,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportParticipant struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Uri               string      `json:"uri"`
	Display           *string     `json:"display,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportSetupActionOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Result            string      `json:"result"`
	Message           *string     `json:"message,omitempty"`
	Detail            *string     `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportSetupActionAssert struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Result            string      `json:"result"`
	Message           *string     `json:"message,omitempty"`
	Detail            *string     `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestReportTestAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
type TestReportTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestReport
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
func (r TestReport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "TestReport/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "TestReport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *TestReport) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *TestReport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_status_codes

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_Result(htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_result_codes

	if resource == nil {
		return CodeSelect("Result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Result", &resource.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_Score(htmlAttrs string) templ.Component {
	if resource == nil {
		return Float64Input("Score", nil, htmlAttrs)
	}
	return Float64Input("Score", resource.Score, htmlAttrs)
}
func (resource *TestReport) T_Tester(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Tester", nil, htmlAttrs)
	}
	return StringInput("Tester", resource.Tester, htmlAttrs)
}
func (resource *TestReport) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Issued", nil, htmlAttrs)
	}
	return DateTimeInput("Issued", resource.Issued, htmlAttrs)
}
func (resource *TestReport) T_ParticipantType(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_participant_type

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Type", &resource.Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_ParticipantUri(numParticipant int, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return StringInput("Participant["+strconv.Itoa(numParticipant)+"]Uri", nil, htmlAttrs)
	}
	return StringInput("Participant["+strconv.Itoa(numParticipant)+"]Uri", &resource.Participant[numParticipant].Uri, htmlAttrs)
}
func (resource *TestReport) T_ParticipantDisplay(numParticipant int, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return StringInput("Participant["+strconv.Itoa(numParticipant)+"]Display", nil, htmlAttrs)
	}
	return StringInput("Participant["+strconv.Itoa(numParticipant)+"]Display", resource.Participant[numParticipant].Display, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationResult(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Result", &resource.Setup.Action[numAction].Operation.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationMessage(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Message", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Message", resource.Setup.Action[numAction].Operation.Message, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationDetail(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Detail", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Detail", resource.Setup.Action[numAction].Operation.Detail, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertResult(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Result", &resource.Setup.Action[numAction].Assert.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertMessage(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Message", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Message", resource.Setup.Action[numAction].Assert.Message, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertDetail(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Detail", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Detail", resource.Setup.Action[numAction].Assert.Detail, htmlAttrs)
}
func (resource *TestReport) T_TestName(numTest int, htmlAttrs string) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("Test["+strconv.Itoa(numTest)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Test["+strconv.Itoa(numTest)+"]Name", resource.Test[numTest].Name, htmlAttrs)
}
func (resource *TestReport) T_TestDescription(numTest int, htmlAttrs string) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("Test["+strconv.Itoa(numTest)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Test["+strconv.Itoa(numTest)+"]Description", resource.Test[numTest].Description, htmlAttrs)
}
