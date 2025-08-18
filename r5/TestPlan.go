//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	VersionAlgorithmString *string              `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding              `json:"versionAlgorithmCoding"`
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
	SourceString      *string          `json:"sourceString"`
	SourceReference   *Reference       `json:"sourceReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestPlan
type TestPlanTestCaseTestData struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type"`
	Content           *Reference  `json:"content,omitempty"`
	SourceString      *string     `json:"sourceString"`
	SourceReference   *Reference  `json:"sourceReference"`
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
