package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                   *time.Time           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r TestPlan) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "TestPlan/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "TestPlan"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *TestPlan) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *TestPlan) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *TestPlan) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *TestPlan) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *TestPlan) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *TestPlan) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *TestPlan) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TestPlan) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *TestPlan) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *TestPlan) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *TestPlan) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *TestPlan) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *TestPlan) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *TestPlan) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *TestPlan) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *TestPlan) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *TestPlan) T_TestTools(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TestTools", nil, htmlAttrs)
	}
	return StringInput("TestTools", resource.TestTools, htmlAttrs)
}
func (resource *TestPlan) T_ExitCriteria(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExitCriteria", nil, htmlAttrs)
	}
	return StringInput("ExitCriteria", resource.ExitCriteria, htmlAttrs)
}
func (resource *TestPlan) T_DependencyDescription(numDependency int, htmlAttrs string) templ.Component {
	if resource == nil || numDependency >= len(resource.Dependency) {
		return StringInput("Dependency["+strconv.Itoa(numDependency)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Dependency["+strconv.Itoa(numDependency)+"]Description", resource.Dependency[numDependency].Description, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseSequence(numTestCase int, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) {
		return IntInput("TestCase["+strconv.Itoa(numTestCase)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("TestCase["+strconv.Itoa(numTestCase)+"]Sequence", resource.TestCase[numTestCase].Sequence, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseDependencyDescription(numTestCase int, numDependency int, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numDependency >= len(resource.TestCase[numTestCase].Dependency) {
		return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]Dependency["+strconv.Itoa(numDependency)+"].Description", nil, htmlAttrs)
	}
	return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]Dependency["+strconv.Itoa(numDependency)+"].Description", resource.TestCase[numTestCase].Dependency[numDependency].Description, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseTestRunNarrative(numTestCase int, numTestRun int, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numTestRun >= len(resource.TestCase[numTestCase].TestRun) {
		return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestRun["+strconv.Itoa(numTestRun)+"].Narrative", nil, htmlAttrs)
	}
	return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestRun["+strconv.Itoa(numTestRun)+"].Narrative", resource.TestCase[numTestCase].TestRun[numTestRun].Narrative, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseTestRunScriptSourceString(numTestCase int, numTestRun int, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numTestRun >= len(resource.TestCase[numTestCase].TestRun) {
		return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestRun["+strconv.Itoa(numTestRun)+"].Script.SourceString", nil, htmlAttrs)
	}
	return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestRun["+strconv.Itoa(numTestRun)+"].Script.SourceString", resource.TestCase[numTestCase].TestRun[numTestRun].Script.SourceString, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseTestDataType(numTestCase int, numTestData int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numTestData >= len(resource.TestCase[numTestCase].TestData) {
		return CodingSelect("TestCase["+strconv.Itoa(numTestCase)+"]TestData["+strconv.Itoa(numTestData)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("TestCase["+strconv.Itoa(numTestCase)+"]TestData["+strconv.Itoa(numTestData)+"].Type", &resource.TestCase[numTestCase].TestData[numTestData].Type, optionsValueSet, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseTestDataSourceString(numTestCase int, numTestData int, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numTestData >= len(resource.TestCase[numTestCase].TestData) {
		return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestData["+strconv.Itoa(numTestData)+"].SourceString", nil, htmlAttrs)
	}
	return StringInput("TestCase["+strconv.Itoa(numTestCase)+"]TestData["+strconv.Itoa(numTestData)+"].SourceString", resource.TestCase[numTestCase].TestData[numTestData].SourceString, htmlAttrs)
}
func (resource *TestPlan) T_TestCaseAssertionType(numTestCase int, numAssertion int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTestCase >= len(resource.TestCase) || numAssertion >= len(resource.TestCase[numTestCase].Assertion) || numType >= len(resource.TestCase[numTestCase].Assertion[numAssertion].Type) {
		return CodeableConceptSelect("TestCase["+strconv.Itoa(numTestCase)+"]Assertion["+strconv.Itoa(numAssertion)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("TestCase["+strconv.Itoa(numTestCase)+"]Assertion["+strconv.Itoa(numAssertion)+"].Type["+strconv.Itoa(numType)+"]", &resource.TestCase[numTestCase].Assertion[numAssertion].Type[numType], optionsValueSet, htmlAttrs)
}
