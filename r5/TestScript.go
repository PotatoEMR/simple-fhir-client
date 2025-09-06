package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScript struct {
	Id                     *string                 `json:"id,omitempty"`
	Meta                   *Meta                   `json:"meta,omitempty"`
	ImplicitRules          *string                 `json:"implicitRules,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Text                   *Narrative              `json:"text,omitempty"`
	Contained              []Resource              `json:"contained,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	Url                    *string                 `json:"url,omitempty"`
	Identifier             []Identifier            `json:"identifier,omitempty"`
	Version                *string                 `json:"version,omitempty"`
	VersionAlgorithmString *string                 `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                  `json:"name"`
	Title                  *string                 `json:"title,omitempty"`
	Status                 string                  `json:"status"`
	Experimental           *bool                   `json:"experimental,omitempty"`
	Date                   *string                 `json:"date,omitempty"`
	Publisher              *string                 `json:"publisher,omitempty"`
	Contact                []ContactDetail         `json:"contact,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	UseContext             []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                *string                 `json:"purpose,omitempty"`
	Copyright              *string                 `json:"copyright,omitempty"`
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty"`
	Origin                 []TestScriptOrigin      `json:"origin,omitempty"`
	Destination            []TestScriptDestination `json:"destination,omitempty"`
	Metadata               *TestScriptMetadata     `json:"metadata,omitempty"`
	Scope                  []TestScriptScope       `json:"scope,omitempty"`
	Fixture                []TestScriptFixture     `json:"fixture,omitempty"`
	Profile                []string                `json:"profile,omitempty"`
	Variable               []TestScriptVariable    `json:"variable,omitempty"`
	Setup                  *TestScriptSetup        `json:"setup,omitempty"`
	Test                   []TestScriptTest        `json:"test,omitempty"`
	Teardown               *TestScriptTeardown     `json:"teardown,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptOrigin struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
	Url               *string     `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptDestination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
	Url               *string     `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptMetadata struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `json:"capability"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptMetadataLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptMetadataCapability struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Required          bool        `json:"required"`
	Validated         bool        `json:"validated"`
	Description       *string     `json:"description,omitempty"`
	Origin            []int       `json:"origin,omitempty"`
	Destination       *int        `json:"destination,omitempty"`
	Link              []string    `json:"link,omitempty"`
	Capabilities      string      `json:"capabilities"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptScope struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Artifact          string           `json:"artifact"`
	Conformance       *CodeableConcept `json:"conformance,omitempty"`
	Phase             *CodeableConcept `json:"phase,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptFixture struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Autocreate        bool        `json:"autocreate"`
	Autodelete        bool        `json:"autodelete"`
	Resource          *Reference  `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptVariable struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	DefaultValue      *string     `json:"defaultValue,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Expression        *string     `json:"expression,omitempty"`
	HeaderField       *string     `json:"headerField,omitempty"`
	Hint              *string     `json:"hint,omitempty"`
	Path              *string     `json:"path,omitempty"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetupActionOperation struct {
	Id                *string                                       `json:"id,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Type              *Coding                                       `json:"type,omitempty"`
	Resource          *string                                       `json:"resource,omitempty"`
	Label             *string                                       `json:"label,omitempty"`
	Description       *string                                       `json:"description,omitempty"`
	Accept            *string                                       `json:"accept,omitempty"`
	ContentType       *string                                       `json:"contentType,omitempty"`
	Destination       *int                                          `json:"destination,omitempty"`
	EncodeRequestUrl  bool                                          `json:"encodeRequestUrl"`
	Method            *string                                       `json:"method,omitempty"`
	Origin            *int                                          `json:"origin,omitempty"`
	Params            *string                                       `json:"params,omitempty"`
	RequestHeader     []TestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`
	RequestId         *string                                       `json:"requestId,omitempty"`
	ResponseId        *string                                       `json:"responseId,omitempty"`
	SourceId          *string                                       `json:"sourceId,omitempty"`
	TargetId          *string                                       `json:"targetId,omitempty"`
	Url               *string                                       `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Field             string      `json:"field"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetupActionAssert struct {
	Id                        *string                                  `json:"id,omitempty"`
	Extension                 []Extension                              `json:"extension,omitempty"`
	ModifierExtension         []Extension                              `json:"modifierExtension,omitempty"`
	Label                     *string                                  `json:"label,omitempty"`
	Description               *string                                  `json:"description,omitempty"`
	Direction                 *string                                  `json:"direction,omitempty"`
	CompareToSourceId         *string                                  `json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string                                  `json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string                                  `json:"compareToSourcePath,omitempty"`
	ContentType               *string                                  `json:"contentType,omitempty"`
	DefaultManualCompletion   *string                                  `json:"defaultManualCompletion,omitempty"`
	Expression                *string                                  `json:"expression,omitempty"`
	HeaderField               *string                                  `json:"headerField,omitempty"`
	MinimumId                 *string                                  `json:"minimumId,omitempty"`
	NavigationLinks           *bool                                    `json:"navigationLinks,omitempty"`
	Operator                  *string                                  `json:"operator,omitempty"`
	Path                      *string                                  `json:"path,omitempty"`
	RequestMethod             *string                                  `json:"requestMethod,omitempty"`
	RequestURL                *string                                  `json:"requestURL,omitempty"`
	Resource                  *string                                  `json:"resource,omitempty"`
	Response                  *string                                  `json:"response,omitempty"`
	ResponseCode              *string                                  `json:"responseCode,omitempty"`
	SourceId                  *string                                  `json:"sourceId,omitempty"`
	StopTestOnFail            bool                                     `json:"stopTestOnFail"`
	ValidateProfileId         *string                                  `json:"validateProfileId,omitempty"`
	Value                     *string                                  `json:"value,omitempty"`
	WarningOnly               bool                                     `json:"warningOnly"`
	Requirement               []TestScriptSetupActionAssertRequirement `json:"requirement,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptSetupActionAssertRequirement struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LinkUri           *string     `json:"linkUri,omitempty"`
	LinkCanonical     *string     `json:"linkCanonical,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestScriptTestAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TestScript
type TestScriptTeardownAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

type OtherTestScript TestScript

// on convert struct to json, automatically add resourceType=TestScript
func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

func (resource *TestScript) T_Id() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Id", nil)
	}
	return StringInput("TestScript.Id", resource.Id)
}
func (resource *TestScript) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("TestScript.ImplicitRules", nil)
	}
	return StringInput("TestScript.ImplicitRules", resource.ImplicitRules)
}
func (resource *TestScript) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("TestScript.Language", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Language", resource.Language, optionsValueSet)
}
func (resource *TestScript) T_Url() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Url", nil)
	}
	return StringInput("TestScript.Url", resource.Url)
}
func (resource *TestScript) T_Version() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Version", nil)
	}
	return StringInput("TestScript.Version", resource.Version)
}
func (resource *TestScript) T_Name() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Name", nil)
	}
	return StringInput("TestScript.Name", &resource.Name)
}
func (resource *TestScript) T_Title() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Title", nil)
	}
	return StringInput("TestScript.Title", resource.Title)
}
func (resource *TestScript) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("TestScript.Status", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Status", &resource.Status, optionsValueSet)
}
func (resource *TestScript) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("TestScript.Experimental", nil)
	}
	return BoolInput("TestScript.Experimental", resource.Experimental)
}
func (resource *TestScript) T_Date() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Date", nil)
	}
	return StringInput("TestScript.Date", resource.Date)
}
func (resource *TestScript) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Publisher", nil)
	}
	return StringInput("TestScript.Publisher", resource.Publisher)
}
func (resource *TestScript) T_Description() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Description", nil)
	}
	return StringInput("TestScript.Description", resource.Description)
}
func (resource *TestScript) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("TestScript.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestScript.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *TestScript) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Purpose", nil)
	}
	return StringInput("TestScript.Purpose", resource.Purpose)
}
func (resource *TestScript) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Copyright", nil)
	}
	return StringInput("TestScript.Copyright", resource.Copyright)
}
func (resource *TestScript) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("TestScript.CopyrightLabel", nil)
	}
	return StringInput("TestScript.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *TestScript) T_Profile(numProfile int) templ.Component {

	if resource == nil || len(resource.Profile) >= numProfile {
		return StringInput("TestScript.Profile["+strconv.Itoa(numProfile)+"]", nil)
	}
	return StringInput("TestScript.Profile["+strconv.Itoa(numProfile)+"]", &resource.Profile[numProfile])
}
func (resource *TestScript) T_OriginId(numOrigin int) templ.Component {

	if resource == nil || len(resource.Origin) >= numOrigin {
		return StringInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Id", nil)
	}
	return StringInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Id", resource.Origin[numOrigin].Id)
}
func (resource *TestScript) T_OriginIndex(numOrigin int) templ.Component {

	if resource == nil || len(resource.Origin) >= numOrigin {
		return IntInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Index", nil)
	}
	return IntInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Index", &resource.Origin[numOrigin].Index)
}
func (resource *TestScript) T_OriginProfile(numOrigin int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Origin) >= numOrigin {
		return CodingSelect("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Profile", nil, optionsValueSet)
	}
	return CodingSelect("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Profile", &resource.Origin[numOrigin].Profile, optionsValueSet)
}
func (resource *TestScript) T_OriginUrl(numOrigin int) templ.Component {

	if resource == nil || len(resource.Origin) >= numOrigin {
		return StringInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Url", nil)
	}
	return StringInput("TestScript.Origin["+strconv.Itoa(numOrigin)+"].Url", resource.Origin[numOrigin].Url)
}
func (resource *TestScript) T_DestinationId(numDestination int) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return StringInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Id", nil)
	}
	return StringInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Id", resource.Destination[numDestination].Id)
}
func (resource *TestScript) T_DestinationIndex(numDestination int) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return IntInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Index", nil)
	}
	return IntInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Index", &resource.Destination[numDestination].Index)
}
func (resource *TestScript) T_DestinationProfile(numDestination int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return CodingSelect("TestScript.Destination["+strconv.Itoa(numDestination)+"].Profile", nil, optionsValueSet)
	}
	return CodingSelect("TestScript.Destination["+strconv.Itoa(numDestination)+"].Profile", &resource.Destination[numDestination].Profile, optionsValueSet)
}
func (resource *TestScript) T_DestinationUrl(numDestination int) templ.Component {

	if resource == nil || len(resource.Destination) >= numDestination {
		return StringInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Url", nil)
	}
	return StringInput("TestScript.Destination["+strconv.Itoa(numDestination)+"].Url", resource.Destination[numDestination].Url)
}
func (resource *TestScript) T_MetadataId() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Metadata.Id", nil)
	}
	return StringInput("TestScript.Metadata.Id", resource.Metadata.Id)
}
func (resource *TestScript) T_MetadataLinkId(numLink int) templ.Component {

	if resource == nil || len(resource.Metadata.Link) >= numLink {
		return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Id", nil)
	}
	return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Id", resource.Metadata.Link[numLink].Id)
}
func (resource *TestScript) T_MetadataLinkUrl(numLink int) templ.Component {

	if resource == nil || len(resource.Metadata.Link) >= numLink {
		return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Url", nil)
	}
	return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Url", &resource.Metadata.Link[numLink].Url)
}
func (resource *TestScript) T_MetadataLinkDescription(numLink int) templ.Component {

	if resource == nil || len(resource.Metadata.Link) >= numLink {
		return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Description", nil)
	}
	return StringInput("TestScript.Metadata.Link["+strconv.Itoa(numLink)+"].Description", resource.Metadata.Link[numLink].Description)
}
func (resource *TestScript) T_MetadataCapabilityId(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Id", nil)
	}
	return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Id", resource.Metadata.Capability[numCapability].Id)
}
func (resource *TestScript) T_MetadataCapabilityRequired(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return BoolInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Required", nil)
	}
	return BoolInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Required", &resource.Metadata.Capability[numCapability].Required)
}
func (resource *TestScript) T_MetadataCapabilityValidated(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return BoolInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Validated", nil)
	}
	return BoolInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Validated", &resource.Metadata.Capability[numCapability].Validated)
}
func (resource *TestScript) T_MetadataCapabilityDescription(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Description", nil)
	}
	return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Description", resource.Metadata.Capability[numCapability].Description)
}
func (resource *TestScript) T_MetadataCapabilityOrigin(numCapability int, numOrigin int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability || len(resource.Metadata.Capability[numCapability].Origin) >= numOrigin {
		return IntInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Origin["+strconv.Itoa(numOrigin)+"]", nil)
	}
	return IntInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Origin["+strconv.Itoa(numOrigin)+"]", &resource.Metadata.Capability[numCapability].Origin[numOrigin])
}
func (resource *TestScript) T_MetadataCapabilityDestination(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return IntInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Destination", nil)
	}
	return IntInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Destination", resource.Metadata.Capability[numCapability].Destination)
}
func (resource *TestScript) T_MetadataCapabilityLink(numCapability int, numLink int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability || len(resource.Metadata.Capability[numCapability].Link) >= numLink {
		return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Link["+strconv.Itoa(numLink)+"]", nil)
	}
	return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Link["+strconv.Itoa(numLink)+"]", &resource.Metadata.Capability[numCapability].Link[numLink])
}
func (resource *TestScript) T_MetadataCapabilityCapabilities(numCapability int) templ.Component {

	if resource == nil || len(resource.Metadata.Capability) >= numCapability {
		return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Capabilities", nil)
	}
	return StringInput("TestScript.Metadata.Capability["+strconv.Itoa(numCapability)+"].Capabilities", &resource.Metadata.Capability[numCapability].Capabilities)
}
func (resource *TestScript) T_ScopeId(numScope int) templ.Component {

	if resource == nil || len(resource.Scope) >= numScope {
		return StringInput("TestScript.Scope["+strconv.Itoa(numScope)+"].Id", nil)
	}
	return StringInput("TestScript.Scope["+strconv.Itoa(numScope)+"].Id", resource.Scope[numScope].Id)
}
func (resource *TestScript) T_ScopeArtifact(numScope int) templ.Component {

	if resource == nil || len(resource.Scope) >= numScope {
		return StringInput("TestScript.Scope["+strconv.Itoa(numScope)+"].Artifact", nil)
	}
	return StringInput("TestScript.Scope["+strconv.Itoa(numScope)+"].Artifact", &resource.Scope[numScope].Artifact)
}
func (resource *TestScript) T_ScopeConformance(numScope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Scope) >= numScope {
		return CodeableConceptSelect("TestScript.Scope["+strconv.Itoa(numScope)+"].Conformance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestScript.Scope["+strconv.Itoa(numScope)+"].Conformance", resource.Scope[numScope].Conformance, optionsValueSet)
}
func (resource *TestScript) T_ScopePhase(numScope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Scope) >= numScope {
		return CodeableConceptSelect("TestScript.Scope["+strconv.Itoa(numScope)+"].Phase", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TestScript.Scope["+strconv.Itoa(numScope)+"].Phase", resource.Scope[numScope].Phase, optionsValueSet)
}
func (resource *TestScript) T_FixtureId(numFixture int) templ.Component {

	if resource == nil || len(resource.Fixture) >= numFixture {
		return StringInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Id", nil)
	}
	return StringInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Id", resource.Fixture[numFixture].Id)
}
func (resource *TestScript) T_FixtureAutocreate(numFixture int) templ.Component {

	if resource == nil || len(resource.Fixture) >= numFixture {
		return BoolInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Autocreate", nil)
	}
	return BoolInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Autocreate", &resource.Fixture[numFixture].Autocreate)
}
func (resource *TestScript) T_FixtureAutodelete(numFixture int) templ.Component {

	if resource == nil || len(resource.Fixture) >= numFixture {
		return BoolInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Autodelete", nil)
	}
	return BoolInput("TestScript.Fixture["+strconv.Itoa(numFixture)+"].Autodelete", &resource.Fixture[numFixture].Autodelete)
}
func (resource *TestScript) T_VariableId(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Id", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Id", resource.Variable[numVariable].Id)
}
func (resource *TestScript) T_VariableName(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Name", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Name", &resource.Variable[numVariable].Name)
}
func (resource *TestScript) T_VariableDefaultValue(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].DefaultValue", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].DefaultValue", resource.Variable[numVariable].DefaultValue)
}
func (resource *TestScript) T_VariableDescription(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Description", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Description", resource.Variable[numVariable].Description)
}
func (resource *TestScript) T_VariableExpression(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Expression", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Expression", resource.Variable[numVariable].Expression)
}
func (resource *TestScript) T_VariableHeaderField(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].HeaderField", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].HeaderField", resource.Variable[numVariable].HeaderField)
}
func (resource *TestScript) T_VariableHint(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Hint", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Hint", resource.Variable[numVariable].Hint)
}
func (resource *TestScript) T_VariablePath(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Path", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].Path", resource.Variable[numVariable].Path)
}
func (resource *TestScript) T_VariableSourceId(numVariable int) templ.Component {

	if resource == nil || len(resource.Variable) >= numVariable {
		return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].SourceId", nil)
	}
	return StringInput("TestScript.Variable["+strconv.Itoa(numVariable)+"].SourceId", resource.Variable[numVariable].SourceId)
}
func (resource *TestScript) T_SetupId() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Setup.Id", nil)
	}
	return StringInput("TestScript.Setup.Id", resource.Setup.Id)
}
func (resource *TestScript) T_SetupActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Id", resource.Setup.Action[numAction].Id)
}
func (resource *TestScript) T_SetupActionOperationId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Id", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Id", resource.Setup.Action[numAction].Operation.Id)
}
func (resource *TestScript) T_SetupActionOperationType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodingSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Type", nil, optionsValueSet)
	}
	return CodingSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Type", resource.Setup.Action[numAction].Operation.Type, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationResource(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Resource", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Resource", resource.Setup.Action[numAction].Operation.Resource)
}
func (resource *TestScript) T_SetupActionOperationLabel(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Label", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Label", resource.Setup.Action[numAction].Operation.Label)
}
func (resource *TestScript) T_SetupActionOperationDescription(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Description", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Description", resource.Setup.Action[numAction].Operation.Description)
}
func (resource *TestScript) T_SetupActionOperationAccept(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Accept", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Accept", resource.Setup.Action[numAction].Operation.Accept, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationContentType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.ContentType", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.ContentType", resource.Setup.Action[numAction].Operation.ContentType, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationDestination(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return IntInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Destination", nil)
	}
	return IntInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Destination", resource.Setup.Action[numAction].Operation.Destination)
}
func (resource *TestScript) T_SetupActionOperationEncodeRequestUrl(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.EncodeRequestUrl", nil)
	}
	return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.EncodeRequestUrl", &resource.Setup.Action[numAction].Operation.EncodeRequestUrl)
}
func (resource *TestScript) T_SetupActionOperationMethod(numAction int) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Method", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Method", resource.Setup.Action[numAction].Operation.Method, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationOrigin(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return IntInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Origin", nil)
	}
	return IntInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Origin", resource.Setup.Action[numAction].Operation.Origin)
}
func (resource *TestScript) T_SetupActionOperationParams(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Params", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Params", resource.Setup.Action[numAction].Operation.Params)
}
func (resource *TestScript) T_SetupActionOperationRequestId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestId", resource.Setup.Action[numAction].Operation.RequestId)
}
func (resource *TestScript) T_SetupActionOperationResponseId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.ResponseId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.ResponseId", resource.Setup.Action[numAction].Operation.ResponseId)
}
func (resource *TestScript) T_SetupActionOperationSourceId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.SourceId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.SourceId", resource.Setup.Action[numAction].Operation.SourceId)
}
func (resource *TestScript) T_SetupActionOperationTargetId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.TargetId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.TargetId", resource.Setup.Action[numAction].Operation.TargetId)
}
func (resource *TestScript) T_SetupActionOperationUrl(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Url", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.Url", resource.Setup.Action[numAction].Operation.Url)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderId(numAction int, numRequestHeader int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction || len(resource.Setup.Action[numAction].Operation.RequestHeader) >= numRequestHeader {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Id", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Id", resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Id)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderField(numAction int, numRequestHeader int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction || len(resource.Setup.Action[numAction].Operation.RequestHeader) >= numRequestHeader {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Field", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Field", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Field)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderValue(numAction int, numRequestHeader int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction || len(resource.Setup.Action[numAction].Operation.RequestHeader) >= numRequestHeader {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Value", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Value", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Value)
}
func (resource *TestScript) T_SetupActionAssertId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Id", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Id", resource.Setup.Action[numAction].Assert.Id)
}
func (resource *TestScript) T_SetupActionAssertLabel(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Label", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Label", resource.Setup.Action[numAction].Assert.Label)
}
func (resource *TestScript) T_SetupActionAssertDescription(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Description", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Description", resource.Setup.Action[numAction].Assert.Description)
}
func (resource *TestScript) T_SetupActionAssertDirection(numAction int) templ.Component {
	optionsValueSet := VSAssert_direction_codes

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Direction", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Direction", resource.Setup.Action[numAction].Assert.Direction, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourceId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourceId", resource.Setup.Action[numAction].Assert.CompareToSourceId)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceExpression(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourceExpression", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourceExpression", resource.Setup.Action[numAction].Assert.CompareToSourceExpression)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourcePath(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourcePath", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.CompareToSourcePath", resource.Setup.Action[numAction].Assert.CompareToSourcePath)
}
func (resource *TestScript) T_SetupActionAssertContentType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ContentType", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ContentType", resource.Setup.Action[numAction].Assert.ContentType, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertDefaultManualCompletion(numAction int) templ.Component {
	optionsValueSet := VSAssert_manual_completion_codes

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.DefaultManualCompletion", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.DefaultManualCompletion", resource.Setup.Action[numAction].Assert.DefaultManualCompletion, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertExpression(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Expression", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Expression", resource.Setup.Action[numAction].Assert.Expression)
}
func (resource *TestScript) T_SetupActionAssertHeaderField(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.HeaderField", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.HeaderField", resource.Setup.Action[numAction].Assert.HeaderField)
}
func (resource *TestScript) T_SetupActionAssertMinimumId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.MinimumId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.MinimumId", resource.Setup.Action[numAction].Assert.MinimumId)
}
func (resource *TestScript) T_SetupActionAssertNavigationLinks(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.NavigationLinks", nil)
	}
	return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.NavigationLinks", resource.Setup.Action[numAction].Assert.NavigationLinks)
}
func (resource *TestScript) T_SetupActionAssertOperator(numAction int) templ.Component {
	optionsValueSet := VSAssert_operator_codes

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Operator", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Operator", resource.Setup.Action[numAction].Assert.Operator, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertPath(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Path", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Path", resource.Setup.Action[numAction].Assert.Path)
}
func (resource *TestScript) T_SetupActionAssertRequestMethod(numAction int) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.RequestMethod", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.RequestMethod", resource.Setup.Action[numAction].Assert.RequestMethod, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertRequestURL(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.RequestURL", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.RequestURL", resource.Setup.Action[numAction].Assert.RequestURL)
}
func (resource *TestScript) T_SetupActionAssertResource(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Resource", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Resource", resource.Setup.Action[numAction].Assert.Resource)
}
func (resource *TestScript) T_SetupActionAssertResponse(numAction int) templ.Component {
	optionsValueSet := VSAssert_response_code_types

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Response", nil, optionsValueSet)
	}
	return CodeSelect("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Response", resource.Setup.Action[numAction].Assert.Response, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertResponseCode(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ResponseCode", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ResponseCode", resource.Setup.Action[numAction].Assert.ResponseCode)
}
func (resource *TestScript) T_SetupActionAssertSourceId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.SourceId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.SourceId", resource.Setup.Action[numAction].Assert.SourceId)
}
func (resource *TestScript) T_SetupActionAssertStopTestOnFail(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.StopTestOnFail", nil)
	}
	return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.StopTestOnFail", &resource.Setup.Action[numAction].Assert.StopTestOnFail)
}
func (resource *TestScript) T_SetupActionAssertValidateProfileId(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ValidateProfileId", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.ValidateProfileId", resource.Setup.Action[numAction].Assert.ValidateProfileId)
}
func (resource *TestScript) T_SetupActionAssertValue(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Value", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Value", resource.Setup.Action[numAction].Assert.Value)
}
func (resource *TestScript) T_SetupActionAssertWarningOnly(numAction int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction {
		return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.WarningOnly", nil)
	}
	return BoolInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.WarningOnly", &resource.Setup.Action[numAction].Assert.WarningOnly)
}
func (resource *TestScript) T_SetupActionAssertRequirementId(numAction int, numRequirement int) templ.Component {

	if resource == nil || len(resource.Setup.Action) >= numAction || len(resource.Setup.Action[numAction].Assert.Requirement) >= numRequirement {
		return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Requirement["+strconv.Itoa(numRequirement)+"].Id", nil)
	}
	return StringInput("TestScript.Setup.Action["+strconv.Itoa(numAction)+"].Assert.Requirement["+strconv.Itoa(numRequirement)+"].Id", resource.Setup.Action[numAction].Assert.Requirement[numRequirement].Id)
}
func (resource *TestScript) T_TestId(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Id", nil)
	}
	return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Id", resource.Test[numTest].Id)
}
func (resource *TestScript) T_TestName(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Name", nil)
	}
	return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Name", resource.Test[numTest].Name)
}
func (resource *TestScript) T_TestDescription(numTest int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest {
		return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Description", nil)
	}
	return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Description", resource.Test[numTest].Description)
}
func (resource *TestScript) T_TestActionId(numTest int, numAction int) templ.Component {

	if resource == nil || len(resource.Test) >= numTest || len(resource.Test[numTest].Action) >= numAction {
		return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestScript.Test["+strconv.Itoa(numTest)+"].Action["+strconv.Itoa(numAction)+"].Id", resource.Test[numTest].Action[numAction].Id)
}
func (resource *TestScript) T_TeardownId() templ.Component {

	if resource == nil {
		return StringInput("TestScript.Teardown.Id", nil)
	}
	return StringInput("TestScript.Teardown.Id", resource.Teardown.Id)
}
func (resource *TestScript) T_TeardownActionId(numAction int) templ.Component {

	if resource == nil || len(resource.Teardown.Action) >= numAction {
		return StringInput("TestScript.Teardown.Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("TestScript.Teardown.Action["+strconv.Itoa(numAction)+"].Id", resource.Teardown.Action[numAction].Id)
}
