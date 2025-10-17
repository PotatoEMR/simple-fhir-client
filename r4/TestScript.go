package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScript struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Url               string                  `json:"url"`
	Identifier        *Identifier             `json:"identifier,omitempty"`
	Version           *string                 `json:"version,omitempty"`
	Name              string                  `json:"name"`
	Title             *string                 `json:"title,omitempty"`
	Status            string                  `json:"status"`
	Experimental      *bool                   `json:"experimental,omitempty"`
	Date              *FhirDateTime           `json:"date,omitempty"`
	Publisher         *string                 `json:"publisher,omitempty"`
	Contact           []ContactDetail         `json:"contact,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	UseContext        []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose           *string                 `json:"purpose,omitempty"`
	Copyright         *string                 `json:"copyright,omitempty"`
	Origin            []TestScriptOrigin      `json:"origin,omitempty"`
	Destination       []TestScriptDestination `json:"destination,omitempty"`
	Metadata          *TestScriptMetadata     `json:"metadata,omitempty"`
	Fixture           []TestScriptFixture     `json:"fixture,omitempty"`
	Profile           []Reference             `json:"profile,omitempty"`
	Variable          []TestScriptVariable    `json:"variable,omitempty"`
	Setup             *TestScriptSetup        `json:"setup,omitempty"`
	Test              []TestScriptTest        `json:"test,omitempty"`
	Teardown          *TestScriptTeardown     `json:"teardown,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptOrigin struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptDestination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptMetadata struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `json:"capability"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptMetadataLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptFixture struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Autocreate        bool        `json:"autocreate"`
	Autodelete        bool        `json:"autodelete"`
	Resource          *Reference  `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Field             string      `json:"field"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptSetupActionAssert struct {
	Id                        *string     `json:"id,omitempty"`
	Extension                 []Extension `json:"extension,omitempty"`
	ModifierExtension         []Extension `json:"modifierExtension,omitempty"`
	Label                     *string     `json:"label,omitempty"`
	Description               *string     `json:"description,omitempty"`
	Direction                 *string     `json:"direction,omitempty"`
	CompareToSourceId         *string     `json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string     `json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string     `json:"compareToSourcePath,omitempty"`
	ContentType               *string     `json:"contentType,omitempty"`
	Expression                *string     `json:"expression,omitempty"`
	HeaderField               *string     `json:"headerField,omitempty"`
	MinimumId                 *string     `json:"minimumId,omitempty"`
	NavigationLinks           *bool       `json:"navigationLinks,omitempty"`
	Operator                  *string     `json:"operator,omitempty"`
	Path                      *string     `json:"path,omitempty"`
	RequestMethod             *string     `json:"requestMethod,omitempty"`
	RequestURL                *string     `json:"requestURL,omitempty"`
	Resource                  *string     `json:"resource,omitempty"`
	Response                  *string     `json:"response,omitempty"`
	ResponseCode              *string     `json:"responseCode,omitempty"`
	SourceId                  *string     `json:"sourceId,omitempty"`
	ValidateProfileId         *string     `json:"validateProfileId,omitempty"`
	Value                     *string     `json:"value,omitempty"`
	WarningOnly               bool        `json:"warningOnly"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestScriptTestAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TestScript
type TestScriptTeardownAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

type OtherTestScript TestScript

// struct -> json, automatically add resourceType=Patient
func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

func (r TestScript) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "TestScript/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "TestScript"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r TestScript) ResourceType() string {
	return "TestScript"
}

func (resource *TestScript) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *TestScript) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *TestScript) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *TestScript) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *TestScript) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *TestScript) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *TestScript) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *TestScript) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *TestScript) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *TestScript) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *TestScript) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *TestScript) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *TestScript) T_Profile(frs []FhirResource, numProfile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProfile >= len(resource.Profile) {
		return ReferenceInput(frs, "profile["+strconv.Itoa(numProfile)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "profile["+strconv.Itoa(numProfile)+"]", &resource.Profile[numProfile], htmlAttrs)
}
func (resource *TestScript) T_OriginIndex(numOrigin int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrigin >= len(resource.Origin) {
		return IntInput("origin["+strconv.Itoa(numOrigin)+"].index", nil, htmlAttrs)
	}
	return IntInput("origin["+strconv.Itoa(numOrigin)+"].index", &resource.Origin[numOrigin].Index, htmlAttrs)
}
func (resource *TestScript) T_OriginProfile(numOrigin int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrigin >= len(resource.Origin) {
		return CodingSelect("origin["+strconv.Itoa(numOrigin)+"].profile", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("origin["+strconv.Itoa(numOrigin)+"].profile", &resource.Origin[numOrigin].Profile, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_DestinationIndex(numDestination int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return IntInput("destination["+strconv.Itoa(numDestination)+"].index", nil, htmlAttrs)
	}
	return IntInput("destination["+strconv.Itoa(numDestination)+"].index", &resource.Destination[numDestination].Index, htmlAttrs)
}
func (resource *TestScript) T_DestinationProfile(numDestination int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return CodingSelect("destination["+strconv.Itoa(numDestination)+"].profile", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("destination["+strconv.Itoa(numDestination)+"].profile", &resource.Destination[numDestination].Profile, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_MetadataLinkUrl(numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Metadata.Link) {
		return StringInput("metadata.link["+strconv.Itoa(numLink)+"].url", nil, htmlAttrs)
	}
	return StringInput("metadata.link["+strconv.Itoa(numLink)+"].url", &resource.Metadata.Link[numLink].Url, htmlAttrs)
}
func (resource *TestScript) T_MetadataLinkDescription(numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Metadata.Link) {
		return StringInput("metadata.link["+strconv.Itoa(numLink)+"].description", nil, htmlAttrs)
	}
	return StringInput("metadata.link["+strconv.Itoa(numLink)+"].description", resource.Metadata.Link[numLink].Description, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityRequired(numCapability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return BoolInput("metadata.capability["+strconv.Itoa(numCapability)+"].required", nil, htmlAttrs)
	}
	return BoolInput("metadata.capability["+strconv.Itoa(numCapability)+"].required", &resource.Metadata.Capability[numCapability].Required, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityValidated(numCapability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return BoolInput("metadata.capability["+strconv.Itoa(numCapability)+"].validated", nil, htmlAttrs)
	}
	return BoolInput("metadata.capability["+strconv.Itoa(numCapability)+"].validated", &resource.Metadata.Capability[numCapability].Validated, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityDescription(numCapability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].description", nil, htmlAttrs)
	}
	return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].description", resource.Metadata.Capability[numCapability].Description, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityOrigin(numCapability int, numOrigin int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) || numOrigin >= len(resource.Metadata.Capability[numCapability].Origin) {
		return IntInput("metadata.capability["+strconv.Itoa(numCapability)+"].origin["+strconv.Itoa(numOrigin)+"]", nil, htmlAttrs)
	}
	return IntInput("metadata.capability["+strconv.Itoa(numCapability)+"].origin["+strconv.Itoa(numOrigin)+"]", &resource.Metadata.Capability[numCapability].Origin[numOrigin], htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityDestination(numCapability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return IntInput("metadata.capability["+strconv.Itoa(numCapability)+"].destination", nil, htmlAttrs)
	}
	return IntInput("metadata.capability["+strconv.Itoa(numCapability)+"].destination", resource.Metadata.Capability[numCapability].Destination, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityLink(numCapability int, numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) || numLink >= len(resource.Metadata.Capability[numCapability].Link) {
		return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].link["+strconv.Itoa(numLink)+"]", nil, htmlAttrs)
	}
	return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].link["+strconv.Itoa(numLink)+"]", &resource.Metadata.Capability[numCapability].Link[numLink], htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityCapabilities(numCapability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].capabilities", nil, htmlAttrs)
	}
	return StringInput("metadata.capability["+strconv.Itoa(numCapability)+"].capabilities", &resource.Metadata.Capability[numCapability].Capabilities, htmlAttrs)
}
func (resource *TestScript) T_FixtureAutocreate(numFixture int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFixture >= len(resource.Fixture) {
		return BoolInput("fixture["+strconv.Itoa(numFixture)+"].autocreate", nil, htmlAttrs)
	}
	return BoolInput("fixture["+strconv.Itoa(numFixture)+"].autocreate", &resource.Fixture[numFixture].Autocreate, htmlAttrs)
}
func (resource *TestScript) T_FixtureAutodelete(numFixture int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFixture >= len(resource.Fixture) {
		return BoolInput("fixture["+strconv.Itoa(numFixture)+"].autodelete", nil, htmlAttrs)
	}
	return BoolInput("fixture["+strconv.Itoa(numFixture)+"].autodelete", &resource.Fixture[numFixture].Autodelete, htmlAttrs)
}
func (resource *TestScript) T_FixtureResource(frs []FhirResource, numFixture int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFixture >= len(resource.Fixture) {
		return ReferenceInput(frs, "fixture["+strconv.Itoa(numFixture)+"].resource", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "fixture["+strconv.Itoa(numFixture)+"].resource", resource.Fixture[numFixture].Resource, htmlAttrs)
}
func (resource *TestScript) T_VariableName(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].name", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].name", &resource.Variable[numVariable].Name, htmlAttrs)
}
func (resource *TestScript) T_VariableDefaultValue(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].defaultValue", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].defaultValue", resource.Variable[numVariable].DefaultValue, htmlAttrs)
}
func (resource *TestScript) T_VariableDescription(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].description", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].description", resource.Variable[numVariable].Description, htmlAttrs)
}
func (resource *TestScript) T_VariableExpression(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].expression", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].expression", resource.Variable[numVariable].Expression, htmlAttrs)
}
func (resource *TestScript) T_VariableHeaderField(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].headerField", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].headerField", resource.Variable[numVariable].HeaderField, htmlAttrs)
}
func (resource *TestScript) T_VariableHint(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].hint", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].hint", resource.Variable[numVariable].Hint, htmlAttrs)
}
func (resource *TestScript) T_VariablePath(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].path", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].path", resource.Variable[numVariable].Path, htmlAttrs)
}
func (resource *TestScript) T_VariableSourceId(numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("variable["+strconv.Itoa(numVariable)+"].sourceId", nil, htmlAttrs)
	}
	return StringInput("variable["+strconv.Itoa(numVariable)+"].sourceId", resource.Variable[numVariable].SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodingSelect("setup.action["+strconv.Itoa(numAction)+"].operation.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("setup.action["+strconv.Itoa(numAction)+"].operation.type", resource.Setup.Action[numAction].Operation.Type, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationResource(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.resource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.resource", resource.Setup.Action[numAction].Operation.Resource, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationLabel(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.label", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.label", resource.Setup.Action[numAction].Operation.Label, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationDescription(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.description", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.description", resource.Setup.Action[numAction].Operation.Description, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationAccept(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.accept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.accept", resource.Setup.Action[numAction].Operation.Accept, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationContentType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.contentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.contentType", resource.Setup.Action[numAction].Operation.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationDestination(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return IntInput("setup.action["+strconv.Itoa(numAction)+"].operation.destination", nil, htmlAttrs)
	}
	return IntInput("setup.action["+strconv.Itoa(numAction)+"].operation.destination", resource.Setup.Action[numAction].Operation.Destination, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationEncodeRequestUrl(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("setup.action["+strconv.Itoa(numAction)+"].operation.encodeRequestUrl", nil, htmlAttrs)
	}
	return BoolInput("setup.action["+strconv.Itoa(numAction)+"].operation.encodeRequestUrl", &resource.Setup.Action[numAction].Operation.EncodeRequestUrl, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationMethod(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].operation.method", resource.Setup.Action[numAction].Operation.Method, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationOrigin(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return IntInput("setup.action["+strconv.Itoa(numAction)+"].operation.origin", nil, htmlAttrs)
	}
	return IntInput("setup.action["+strconv.Itoa(numAction)+"].operation.origin", resource.Setup.Action[numAction].Operation.Origin, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationParams(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.params", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.params", resource.Setup.Action[numAction].Operation.Params, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestId", resource.Setup.Action[numAction].Operation.RequestId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationResponseId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.responseId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.responseId", resource.Setup.Action[numAction].Operation.ResponseId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationSourceId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.sourceId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.sourceId", resource.Setup.Action[numAction].Operation.SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationTargetId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.targetId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.targetId", resource.Setup.Action[numAction].Operation.TargetId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationUrl(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.url", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.url", resource.Setup.Action[numAction].Operation.Url, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderField(numAction int, numRequestHeader int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) || numRequestHeader >= len(resource.Setup.Action[numAction].Operation.RequestHeader) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestHeader["+strconv.Itoa(numRequestHeader)+"].field", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestHeader["+strconv.Itoa(numRequestHeader)+"].field", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Field, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderValue(numAction int, numRequestHeader int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) || numRequestHeader >= len(resource.Setup.Action[numAction].Operation.RequestHeader) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestHeader["+strconv.Itoa(numRequestHeader)+"].value", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].operation.requestHeader["+strconv.Itoa(numRequestHeader)+"].value", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Value, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertLabel(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.label", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.label", resource.Setup.Action[numAction].Assert.Label, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertDescription(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.description", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.description", resource.Setup.Action[numAction].Assert.Description, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertDirection(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAssert_direction_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.direction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.direction", resource.Setup.Action[numAction].Assert.Direction, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourceId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourceId", resource.Setup.Action[numAction].Assert.CompareToSourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceExpression(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourceExpression", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourceExpression", resource.Setup.Action[numAction].Assert.CompareToSourceExpression, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourcePath(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourcePath", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.compareToSourcePath", resource.Setup.Action[numAction].Assert.CompareToSourcePath, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertContentType(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.contentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.contentType", resource.Setup.Action[numAction].Assert.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertExpression(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.expression", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.expression", resource.Setup.Action[numAction].Assert.Expression, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertHeaderField(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.headerField", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.headerField", resource.Setup.Action[numAction].Assert.HeaderField, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertMinimumId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.minimumId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.minimumId", resource.Setup.Action[numAction].Assert.MinimumId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertNavigationLinks(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("setup.action["+strconv.Itoa(numAction)+"].assert.navigationLinks", nil, htmlAttrs)
	}
	return BoolInput("setup.action["+strconv.Itoa(numAction)+"].assert.navigationLinks", resource.Setup.Action[numAction].Assert.NavigationLinks, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertOperator(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAssert_operator_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.operator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.operator", resource.Setup.Action[numAction].Assert.Operator, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertPath(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.path", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.path", resource.Setup.Action[numAction].Assert.Path, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertRequestMethod(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.requestMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.requestMethod", resource.Setup.Action[numAction].Assert.RequestMethod, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertRequestURL(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.requestURL", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.requestURL", resource.Setup.Action[numAction].Assert.RequestURL, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResource(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.resource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.resource", resource.Setup.Action[numAction].Assert.Resource, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResponse(numAction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAssert_response_code_types

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.response", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("setup.action["+strconv.Itoa(numAction)+"].assert.response", resource.Setup.Action[numAction].Assert.Response, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResponseCode(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.responseCode", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.responseCode", resource.Setup.Action[numAction].Assert.ResponseCode, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertSourceId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.sourceId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.sourceId", resource.Setup.Action[numAction].Assert.SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertValidateProfileId(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.validateProfileId", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.validateProfileId", resource.Setup.Action[numAction].Assert.ValidateProfileId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertValue(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.value", nil, htmlAttrs)
	}
	return StringInput("setup.action["+strconv.Itoa(numAction)+"].assert.value", resource.Setup.Action[numAction].Assert.Value, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertWarningOnly(numAction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("setup.action["+strconv.Itoa(numAction)+"].assert.warningOnly", nil, htmlAttrs)
	}
	return BoolInput("setup.action["+strconv.Itoa(numAction)+"].assert.warningOnly", &resource.Setup.Action[numAction].Assert.WarningOnly, htmlAttrs)
}
func (resource *TestScript) T_TestName(numTest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("test["+strconv.Itoa(numTest)+"].name", nil, htmlAttrs)
	}
	return StringInput("test["+strconv.Itoa(numTest)+"].name", resource.Test[numTest].Name, htmlAttrs)
}
func (resource *TestScript) T_TestDescription(numTest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("test["+strconv.Itoa(numTest)+"].description", nil, htmlAttrs)
	}
	return StringInput("test["+strconv.Itoa(numTest)+"].description", resource.Test[numTest].Description, htmlAttrs)
}
