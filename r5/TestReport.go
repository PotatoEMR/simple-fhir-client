package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
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
	TestScript        string                  `json:"testScript"`
	Result            string                  `json:"result"`
	Score             *float64                `json:"score,omitempty"`
	Tester            *string                 `json:"tester,omitempty"`
	Issued            *string                 `json:"issued,omitempty"`
	Participant       []TestReportParticipant `json:"participant,omitempty"`
	Setup             *TestReportSetup        `json:"setup,omitempty"`
	Test              []TestReportTest        `json:"test,omitempty"`
	Teardown          *TestReportTeardown     `json:"teardown,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportParticipant struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Uri               string      `json:"uri"`
	Display           *string     `json:"display,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportSetupActionOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Result            string      `json:"result"`
	Message           *string     `json:"message,omitempty"`
	Detail            *string     `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportSetupActionAssert struct {
	Id                *string                                  `json:"id,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Result            string                                   `json:"result"`
	Message           *string                                  `json:"message,omitempty"`
	Detail            *string                                  `json:"detail,omitempty"`
	Requirement       []TestReportSetupActionAssertRequirement `json:"requirement,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportSetupActionAssertRequirement struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LinkUri           *string     `json:"linkUri,omitempty"`
	LinkCanonical     *string     `json:"linkCanonical,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestReportTestAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
type TestReportTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestReport
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

func (resource *TestReport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Id", nil)
	}
	return StringInput("TestReport.Id", resource.Id)
}
func (resource *TestReport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("TestReport.ImplicitRules", nil)
	}
	return StringInput("TestReport.ImplicitRules", resource.ImplicitRules)
}
func (resource *TestReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("TestReport.Language", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Language", resource.Language, optionsValueSet)
}
func (resource *TestReport) T_Name() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Name", nil)
	}
	return StringInput("TestReport.Name", resource.Name)
}
func (resource *TestReport) T_Status() templ.Component {
	optionsValueSet := VSReport_status_codes

	if resource == nil {
		return CodeSelect("TestReport.Status", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Status", &resource.Status, optionsValueSet)
}
func (resource *TestReport) T_TestScript() templ.Component {

	if resource == nil {
		return StringInput("TestReport.TestScript", nil)
	}
	return StringInput("TestReport.TestScript", &resource.TestScript)
}
func (resource *TestReport) T_Result() templ.Component {
	optionsValueSet := VSReport_result_codes

	if resource == nil {
		return CodeSelect("TestReport.Result", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Result", &resource.Result, optionsValueSet)
}
func (resource *TestReport) T_Score() templ.Component {

	if resource == nil {
		return Float64Input("TestReport.Score", nil)
	}
	return Float64Input("TestReport.Score", resource.Score)
}
func (resource *TestReport) T_Tester() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Tester", nil)
	}
	return StringInput("TestReport.Tester", resource.Tester)
}
func (resource *TestReport) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Issued", nil)
	}
	return StringInput("TestReport.Issued", resource.Issued)
}
func (resource *TestReport) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *TestReport) T_ParticipantType(numParticipant int) templ.Component {
	optionsValueSet := VSReport_participant_type

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeSelect("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Type", &resource.Participant[numParticipant].Type, optionsValueSet)
}
func (resource *TestReport) T_ParticipantUri(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Uri", nil)
	}
	return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Uri", &resource.Participant[numParticipant].Uri)
}
func (resource *TestReport) T_ParticipantDisplay(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Display", nil)
	}
	return StringInput("TestReport.Participant["+strconv.Itoa(numParticipant)+"].Display", resource.Participant[numParticipant].Display)
}
func (resource *TestReport) T_SetupId() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Setup.Id", nil)
	}
	return StringInput("TestReport.Setup.Id", resource.Setup.Id)
}
func (resource *TestReport) T_SetupActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Id", resource.Setup.Action[numAction].Id)
}
func (resource *TestReport) T_SetupActionOperationId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Id", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Id", resource.Setup.Action[numAction].Operation.Id)
}
func (resource *TestReport) T_SetupActionOperationResult(numAction int) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Result", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Result", &resource.Setup.Action[numAction].Operation.Result, optionsValueSet)
}
func (resource *TestReport) T_SetupActionOperationMessage(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Message", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Message", resource.Setup.Action[numAction].Operation.Message)
}
func (resource *TestReport) T_SetupActionOperationDetail(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Detail", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Detail", resource.Setup.Action[numAction].Operation.Detail)
}
func (resource *TestReport) T_SetupActionAssertId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Id", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Id", resource.Setup.Action[numAction].Assert.Id)
}
func (resource *TestReport) T_SetupActionAssertResult(numAction int) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Result", nil, optionsValueSet)
	}
	return CodeSelect("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Result", &resource.Setup.Action[numAction].Assert.Result, optionsValueSet)
}
func (resource *TestReport) T_SetupActionAssertMessage(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Message", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Message", resource.Setup.Action[numAction].Assert.Message)
}
func (resource *TestReport) T_SetupActionAssertDetail(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Detail", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Detail", resource.Setup.Action[numAction].Assert.Detail)
}
func (resource *TestReport) T_SetupActionAssertRequirementId(numAction int, numRequirement int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction || len(resource.Setup.Action[numAction].Assert.Requirement) >= numRequirement {
		return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Requirement["+strconv.Itoa(numRequirement)+"].Id", nil)
	}
	return StringInput("TestReport.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Requirement["+strconv.Itoa(numRequirement)+"].Id", resource.Setup.Action[numAction].Assert.Requirement[numRequirement].Id)
}
func (resource *TestReport) T_TestId(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Id", nil)
	}
	return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Id", resource.Test[numTest].Id)
}
func (resource *TestReport) T_TestName(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Name", nil)
	}
	return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Name", resource.Test[numTest].Name)
}
func (resource *TestReport) T_TestDescription(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Description", nil)
	}
	return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Description", resource.Test[numTest].Description)
}
func (resource *TestReport) T_TestActionId(numTest int, numAction int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest || len(resource.Test[numTest].Action) >= numAction {
		return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestReport.Test["+strconv.Itoa(numTest)+"].Action["+strconv.Itoa(numAction)+"].Id", resource.Test[numTest].Action[numAction].Id)
}
func (resource *TestReport) T_TeardownId() templ.Component {

	if resource == nil {
		return StringInput("TestReport.Teardown.Id", nil)
	}
	return StringInput("TestReport.Teardown.Id", resource.Teardown.Id)
}
func (resource *TestReport) T_TeardownActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Teardown.Action) >= numAction {
		return StringInput("TestReport.Teardown.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestReport.Teardown.Action["+strconv.Itoa(numAction)+"].Id", resource.Teardown.Action[numAction].Id)
}
