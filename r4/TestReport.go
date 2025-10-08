package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Issued            *FhirDateTime           `json:"issued,omitempty"`
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
func (resource *TestReport) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *TestReport) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_status_codes

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_TestScript(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "testScript", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "testScript", &resource.TestScript, htmlAttrs)
}
func (resource *TestReport) T_Result(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_result_codes

	if resource == nil {
		return CodeSelect("result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("result", &resource.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_Score(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("score", nil, htmlAttrs)
	}
	return Float64Input("score", resource.Score, htmlAttrs)
}
func (resource *TestReport) T_Tester(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("tester", nil, htmlAttrs)
	}
	return StringInput("tester", resource.Tester, htmlAttrs)
}
func (resource *TestReport) T_Issued(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("issued", nil, htmlAttrs)
	}
	return FhirDateTimeInput("issued", resource.Issued, htmlAttrs)
}
func (resource *TestReport) T_ParticipantType(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_participant_type

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].type", &resource.Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_ParticipantUri(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return StringInput("participant["+strconv.Itoa(numParticipant)+"].uri", nil, htmlAttrs)
	}
	return StringInput("participant["+strconv.Itoa(numParticipant)+"].uri", &resource.Participant[numParticipant].Uri, htmlAttrs)
}
func (resource *TestReport) T_ParticipantDisplay(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return StringInput("participant["+strconv.Itoa(numParticipant)+"].display", nil, htmlAttrs)
	}
	return StringInput("participant["+strconv.Itoa(numParticipant)+"].display", resource.Participant[numParticipant].Display, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationResult(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.result", &resource.Setup.Action[numAction].Operation.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationMessage(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.message", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.message", resource.Setup.Action[numAction].Operation.Message, htmlAttrs)
}
func (resource *TestReport) T_SetupActionOperationDetail(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.detail", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.detail", resource.Setup.Action[numAction].Operation.Detail, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertResult(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_action_result_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.result", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.result", &resource.Setup.Action[numAction].Assert.Result, optionsValueSet, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertMessage(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.message", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.message", resource.Setup.Action[numAction].Assert.Message, htmlAttrs)
}
func (resource *TestReport) T_SetupActionAssertDetail(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.detail", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.detail", resource.Setup.Action[numAction].Assert.Detail, htmlAttrs)
}
func (resource *TestReport) T_TestName(numTest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("test["+strconv.Itoa(numTest)+"].name", nil, htmlAttrs)
	}
	return StringInput("test["+strconv.Itoa(numTest)+"].name", resource.Test[numTest].Name, htmlAttrs)
}
func (resource *TestReport) T_TestDescription(numTest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("test["+strconv.Itoa(numTest)+"].description", nil, htmlAttrs)
	}
	return StringInput("test["+strconv.Itoa(numTest)+"].description", resource.Test[numTest].Description, htmlAttrs)
}
