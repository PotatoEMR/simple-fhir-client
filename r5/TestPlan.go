package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlan struct {
	Id                     *string              `json:"id,omitempty"`
	Meta                   *Meta                `json:"meta,omitempty"`
	ImplicitRules          *string              `json:"implicitRules,omitempty"`
	Language               *string              `json:"language,omitempty"`
	Text                   *Narrative           `json:"text,omitempty"`
	Contained              []Resource           `json:"contained,omitempty"`
	Extension              []Extension          `json:"extension,omitempty"`
	ModifierExtension      []Extension          `json:"modifierExtension,omitempty"`
	Url                    *string              `json:"url,omitempty"`
	Identifier             []Identifier         `json:"identifier,omitempty"`
	Version                *string              `json:"version,omitempty"`
	VersionAlgorithmString *string              `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding              `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string              `json:"name,omitempty"`
	Title                  *string              `json:"title,omitempty"`
	Status                 string               `json:"status"`
	Experimental           *bool                `json:"experimental,omitempty"`
	Date                   *string              `json:"date,omitempty"`
	Publisher              *string              `json:"publisher,omitempty"`
	Contact                []ContactDetail      `json:"contact,omitempty"`
	Description            *string              `json:"description,omitempty"`
	UseContext             []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose                *string              `json:"purpose,omitempty"`
	Copyright              *string              `json:"copyright,omitempty"`
	CopyrightLabel         *string              `json:"copyrightLabel,omitempty"`
	Category               []CodeableConcept    `json:"category,omitempty"`
	Scope                  []Reference          `json:"scope,omitempty"`
	TestTools              *string              `json:"testTools,omitempty"`
	Dependency             []TestPlanDependency `json:"dependency,omitempty"`
	ExitCriteria           *string              `json:"exitCriteria,omitempty"`
	TestCase               []TestPlanTestCase   `json:"testCase,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanDependency struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Predecessor       *Reference  `json:"predecessor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCase struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Sequence          *int                         `json:"sequence,omitempty"`
	Scope             []Reference                  `json:"scope,omitempty"`
	Dependency        []TestPlanTestCaseDependency `json:"dependency,omitempty"`
	TestRun           []TestPlanTestCaseTestRun    `json:"testRun,omitempty"`
	TestData          []TestPlanTestCaseTestData   `json:"testData,omitempty"`
	Assertion         []TestPlanTestCaseAssertion  `json:"assertion,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseDependency struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Predecessor       *Reference  `json:"predecessor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseTestRun struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Narrative         *string                        `json:"narrative,omitempty"`
	Script            *TestPlanTestCaseTestRunScript `json:"script,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseTestRunScript struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
	SourceString      *string          `json:"sourceString,omitempty"`
	SourceReference   *Reference       `json:"sourceReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseTestData struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type"`
	Content           *Reference  `json:"content,omitempty"`
	SourceString      *string     `json:"sourceString,omitempty"`
	SourceReference   *Reference  `json:"sourceReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseAssertion struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept   `json:"type,omitempty"`
	Object            []CodeableReference `json:"object,omitempty"`
	Result            []CodeableReference `json:"result,omitempty"`
}

type OtherTestPlan TestPlan

// on convert struct to json, automatically add resourceType=TestPlan
func (r TestPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestPlan
		ResourceType string `json:"resourceType"`
	}{
		OtherTestPlan: OtherTestPlan(r),
		ResourceType:  "TestPlan",
	})
}

func (resource *TestPlan) T_Id() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Id", nil)
	}
	return StringInput("TestPlan.Id", resource.Id)
}
func (resource *TestPlan) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.ImplicitRules", nil)
	}
	return StringInput("TestPlan.ImplicitRules", resource.ImplicitRules)
}
func (resource *TestPlan) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("TestPlan.Language", nil, optionsValueSet)
	}
	return CodeSelect("TestPlan.Language", resource.Language, optionsValueSet)
}
func (resource *TestPlan) T_Url() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Url", nil)
	}
	return StringInput("TestPlan.Url", resource.Url)
}
func (resource *TestPlan) T_Version() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Version", nil)
	}
	return StringInput("TestPlan.Version", resource.Version)
}
func (resource *TestPlan) T_Name() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Name", nil)
	}
	return StringInput("TestPlan.Name", resource.Name)
}
func (resource *TestPlan) T_Title() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Title", nil)
	}
	return StringInput("TestPlan.Title", resource.Title)
}
func (resource *TestPlan) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("TestPlan.Status", nil, optionsValueSet)
	}
	return CodeSelect("TestPlan.Status", &resource.Status, optionsValueSet)
}
func (resource *TestPlan) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("TestPlan.Experimental", nil)
	}
	return BoolInput("TestPlan.Experimental", resource.Experimental)
}
func (resource *TestPlan) T_Date() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Date", nil)
	}
	return StringInput("TestPlan.Date", resource.Date)
}
func (resource *TestPlan) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Publisher", nil)
	}
	return StringInput("TestPlan.Publisher", resource.Publisher)
}
func (resource *TestPlan) T_Description() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Description", nil)
	}
	return StringInput("TestPlan.Description", resource.Description)
}
func (resource *TestPlan) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("TestPlan.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestPlan.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *TestPlan) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Purpose", nil)
	}
	return StringInput("TestPlan.Purpose", resource.Purpose)
}
func (resource *TestPlan) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.Copyright", nil)
	}
	return StringInput("TestPlan.Copyright", resource.Copyright)
}
func (resource *TestPlan) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.CopyrightLabel", nil)
	}
	return StringInput("TestPlan.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *TestPlan) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("TestPlan.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestPlan.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *TestPlan) T_TestTools() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.TestTools", nil)
	}
	return StringInput("TestPlan.TestTools", resource.TestTools)
}
func (resource *TestPlan) T_ExitCriteria() templ.Component {

	if resource == nil {
		return StringInput("TestPlan.ExitCriteria", nil)
	}
	return StringInput("TestPlan.ExitCriteria", resource.ExitCriteria)
}
func (resource *TestPlan) T_DependencyId(numDependency int) templ.Component {

	if resource == nil || len(resource.Dependency) >= numDependency {
		return StringInput("TestPlan.Dependency["+strconv.Itoa(numDependency)+"].Id", nil)
	}
	return StringInput("TestPlan.Dependency["+strconv.Itoa(numDependency)+"].Id", resource.Dependency[numDependency].Id)
}
func (resource *TestPlan) T_DependencyDescription(numDependency int) templ.Component {

	if resource == nil || len(resource.Dependency) >= numDependency {
		return StringInput("TestPlan.Dependency["+strconv.Itoa(numDependency)+"].Description", nil)
	}
	return StringInput("TestPlan.Dependency["+strconv.Itoa(numDependency)+"].Description", resource.Dependency[numDependency].Description)
}
func (resource *TestPlan) T_TestCaseId(numTestCase int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Id", resource.TestCase[numTestCase].Id)
}
func (resource *TestPlan) T_TestCaseSequence(numTestCase int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase {
		return IntInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Sequence", nil)
	}
	return IntInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Sequence", resource.TestCase[numTestCase].Sequence)
}
func (resource *TestPlan) T_TestCaseDependencyId(numTestCase int, numDependency int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].Dependency) >= numDependency {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Dependency["+strconv.Itoa(numDependency)+"].Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Dependency["+strconv.Itoa(numDependency)+"].Id", resource.TestCase[numTestCase].Dependency[numDependency].Id)
}
func (resource *TestPlan) T_TestCaseDependencyDescription(numTestCase int, numDependency int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].Dependency) >= numDependency {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Dependency["+strconv.Itoa(numDependency)+"].Description", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Dependency["+strconv.Itoa(numDependency)+"].Description", resource.TestCase[numTestCase].Dependency[numDependency].Description)
}
func (resource *TestPlan) T_TestCaseTestRunId(numTestCase int, numTestRun int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestRun) >= numTestRun {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Id", resource.TestCase[numTestCase].TestRun[numTestRun].Id)
}
func (resource *TestPlan) T_TestCaseTestRunNarrative(numTestCase int, numTestRun int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestRun) >= numTestRun {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Narrative", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Narrative", resource.TestCase[numTestCase].TestRun[numTestRun].Narrative)
}
func (resource *TestPlan) T_TestCaseTestRunScriptId(numTestCase int, numTestRun int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestRun) >= numTestRun {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Script.Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Script.Id", resource.TestCase[numTestCase].TestRun[numTestRun].Script.Id)
}
func (resource *TestPlan) T_TestCaseTestRunScriptLanguage(numTestCase int, numTestRun int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestRun) >= numTestRun {
		return CodeableConceptSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Script.Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestRun["+strconv.Itoa(numTestRun)+"].Script.Language", resource.TestCase[numTestCase].TestRun[numTestRun].Script.Language, optionsValueSet)
}
func (resource *TestPlan) T_TestCaseTestDataId(numTestCase int, numTestData int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestData) >= numTestData {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestData["+strconv.Itoa(numTestData)+"].Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestData["+strconv.Itoa(numTestData)+"].Id", resource.TestCase[numTestCase].TestData[numTestData].Id)
}
func (resource *TestPlan) T_TestCaseTestDataType(numTestCase int, numTestData int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].TestData) >= numTestData {
		return CodingSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestData["+strconv.Itoa(numTestData)+"].Type", nil, optionsValueSet)
	}
	return CodingSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].TestData["+strconv.Itoa(numTestData)+"].Type", &resource.TestCase[numTestCase].TestData[numTestData].Type, optionsValueSet)
}
func (resource *TestPlan) T_TestCaseAssertionId(numTestCase int, numAssertion int) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].Assertion) >= numAssertion {
		return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Assertion["+strconv.Itoa(numAssertion)+"].Id", nil)
	}
	return StringInput("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Assertion["+strconv.Itoa(numAssertion)+"].Id", resource.TestCase[numTestCase].Assertion[numAssertion].Id)
}
func (resource *TestPlan) T_TestCaseAssertionType(numTestCase int, numAssertion int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TestCase) >= numTestCase || len(resource.TestCase[numTestCase].Assertion) >= numAssertion || len(resource.TestCase[numTestCase].Assertion[numAssertion].Type) >= numType {
		return CodeableConceptSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Assertion["+strconv.Itoa(numAssertion)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestPlan.TestCase["+strconv.Itoa(numTestCase)+"].Assertion["+strconv.Itoa(numAssertion)+"].Type["+strconv.Itoa(numType)+"]", &resource.TestCase[numTestCase].Assertion[numAssertion].Type[numType], optionsValueSet)
}
