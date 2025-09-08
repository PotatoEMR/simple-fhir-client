package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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
	Date              *time.Time              `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptOrigin struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptDestination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptMetadata struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `json:"capability"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptMetadataLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptFixture struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Autocreate        bool        `json:"autocreate"`
	Autodelete        bool        `json:"autodelete"`
	Resource          *Reference  `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Field             string      `json:"field"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptTest struct {
	Id                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestScriptTestAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptTestAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
type TestScriptTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `json:"action"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/TestScript
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
func (resource *TestScript) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", &resource.Url, htmlAttrs)
}
func (resource *TestScript) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *TestScript) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", &resource.Name, htmlAttrs)
}
func (resource *TestScript) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *TestScript) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *TestScript) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *TestScript) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *TestScript) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *TestScript) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *TestScript) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *TestScript) T_OriginIndex(numOrigin int, htmlAttrs string) templ.Component {
	if resource == nil || numOrigin >= len(resource.Origin) {
		return IntInput("Origin["+strconv.Itoa(numOrigin)+"]Index", nil, htmlAttrs)
	}
	return IntInput("Origin["+strconv.Itoa(numOrigin)+"]Index", &resource.Origin[numOrigin].Index, htmlAttrs)
}
func (resource *TestScript) T_OriginProfile(numOrigin int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOrigin >= len(resource.Origin) {
		return CodingSelect("Origin["+strconv.Itoa(numOrigin)+"]Profile", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Origin["+strconv.Itoa(numOrigin)+"]Profile", &resource.Origin[numOrigin].Profile, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_DestinationIndex(numDestination int, htmlAttrs string) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return IntInput("Destination["+strconv.Itoa(numDestination)+"]Index", nil, htmlAttrs)
	}
	return IntInput("Destination["+strconv.Itoa(numDestination)+"]Index", &resource.Destination[numDestination].Index, htmlAttrs)
}
func (resource *TestScript) T_DestinationProfile(numDestination int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDestination >= len(resource.Destination) {
		return CodingSelect("Destination["+strconv.Itoa(numDestination)+"]Profile", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Destination["+strconv.Itoa(numDestination)+"]Profile", &resource.Destination[numDestination].Profile, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_MetadataLinkUrl(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Metadata.Link) {
		return StringInput("MetadataLink["+strconv.Itoa(numLink)+"].Url", nil, htmlAttrs)
	}
	return StringInput("MetadataLink["+strconv.Itoa(numLink)+"].Url", &resource.Metadata.Link[numLink].Url, htmlAttrs)
}
func (resource *TestScript) T_MetadataLinkDescription(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Metadata.Link) {
		return StringInput("MetadataLink["+strconv.Itoa(numLink)+"].Description", nil, htmlAttrs)
	}
	return StringInput("MetadataLink["+strconv.Itoa(numLink)+"].Description", resource.Metadata.Link[numLink].Description, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityRequired(numCapability int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return BoolInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Required", nil, htmlAttrs)
	}
	return BoolInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Required", &resource.Metadata.Capability[numCapability].Required, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityValidated(numCapability int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return BoolInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Validated", nil, htmlAttrs)
	}
	return BoolInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Validated", &resource.Metadata.Capability[numCapability].Validated, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityDescription(numCapability int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Description", nil, htmlAttrs)
	}
	return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Description", resource.Metadata.Capability[numCapability].Description, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityOrigin(numCapability int, numOrigin int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) || numOrigin >= len(resource.Metadata.Capability[numCapability].Origin) {
		return IntInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Origin["+strconv.Itoa(numOrigin)+"]", nil, htmlAttrs)
	}
	return IntInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Origin["+strconv.Itoa(numOrigin)+"]", &resource.Metadata.Capability[numCapability].Origin[numOrigin], htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityDestination(numCapability int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return IntInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Destination", nil, htmlAttrs)
	}
	return IntInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Destination", resource.Metadata.Capability[numCapability].Destination, htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityLink(numCapability int, numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) || numLink >= len(resource.Metadata.Capability[numCapability].Link) {
		return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Link["+strconv.Itoa(numLink)+"]", nil, htmlAttrs)
	}
	return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Link["+strconv.Itoa(numLink)+"]", &resource.Metadata.Capability[numCapability].Link[numLink], htmlAttrs)
}
func (resource *TestScript) T_MetadataCapabilityCapabilities(numCapability int, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Metadata.Capability) {
		return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Capabilities", nil, htmlAttrs)
	}
	return StringInput("MetadataCapability["+strconv.Itoa(numCapability)+"].Capabilities", &resource.Metadata.Capability[numCapability].Capabilities, htmlAttrs)
}
func (resource *TestScript) T_FixtureAutocreate(numFixture int, htmlAttrs string) templ.Component {
	if resource == nil || numFixture >= len(resource.Fixture) {
		return BoolInput("Fixture["+strconv.Itoa(numFixture)+"]Autocreate", nil, htmlAttrs)
	}
	return BoolInput("Fixture["+strconv.Itoa(numFixture)+"]Autocreate", &resource.Fixture[numFixture].Autocreate, htmlAttrs)
}
func (resource *TestScript) T_FixtureAutodelete(numFixture int, htmlAttrs string) templ.Component {
	if resource == nil || numFixture >= len(resource.Fixture) {
		return BoolInput("Fixture["+strconv.Itoa(numFixture)+"]Autodelete", nil, htmlAttrs)
	}
	return BoolInput("Fixture["+strconv.Itoa(numFixture)+"]Autodelete", &resource.Fixture[numFixture].Autodelete, htmlAttrs)
}
func (resource *TestScript) T_VariableName(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]Name", &resource.Variable[numVariable].Name, htmlAttrs)
}
func (resource *TestScript) T_VariableDefaultValue(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]DefaultValue", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]DefaultValue", resource.Variable[numVariable].DefaultValue, htmlAttrs)
}
func (resource *TestScript) T_VariableDescription(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]Description", resource.Variable[numVariable].Description, htmlAttrs)
}
func (resource *TestScript) T_VariableExpression(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]Expression", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]Expression", resource.Variable[numVariable].Expression, htmlAttrs)
}
func (resource *TestScript) T_VariableHeaderField(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]HeaderField", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]HeaderField", resource.Variable[numVariable].HeaderField, htmlAttrs)
}
func (resource *TestScript) T_VariableHint(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]Hint", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]Hint", resource.Variable[numVariable].Hint, htmlAttrs)
}
func (resource *TestScript) T_VariablePath(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]Path", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]Path", resource.Variable[numVariable].Path, htmlAttrs)
}
func (resource *TestScript) T_VariableSourceId(numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numVariable >= len(resource.Variable) {
		return StringInput("Variable["+strconv.Itoa(numVariable)+"]SourceId", nil, htmlAttrs)
	}
	return StringInput("Variable["+strconv.Itoa(numVariable)+"]SourceId", resource.Variable[numVariable].SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationType(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodingSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Type", resource.Setup.Action[numAction].Operation.Type, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationResource(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Resource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Resource", resource.Setup.Action[numAction].Operation.Resource, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationLabel(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Label", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Label", resource.Setup.Action[numAction].Operation.Label, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationDescription(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Description", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Description", resource.Setup.Action[numAction].Operation.Description, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationAccept(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Accept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Accept", resource.Setup.Action[numAction].Operation.Accept, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationContentType(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.ContentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.ContentType", resource.Setup.Action[numAction].Operation.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationDestination(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return IntInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Destination", nil, htmlAttrs)
	}
	return IntInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Destination", resource.Setup.Action[numAction].Operation.Destination, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationEncodeRequestUrl(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.EncodeRequestUrl", nil, htmlAttrs)
	}
	return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.EncodeRequestUrl", &resource.Setup.Action[numAction].Operation.EncodeRequestUrl, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationMethod(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Operation.Method", resource.Setup.Action[numAction].Operation.Method, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationOrigin(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return IntInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Origin", nil, htmlAttrs)
	}
	return IntInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Origin", resource.Setup.Action[numAction].Operation.Origin, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationParams(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Params", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Params", resource.Setup.Action[numAction].Operation.Params, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestId", resource.Setup.Action[numAction].Operation.RequestId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationResponseId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.ResponseId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.ResponseId", resource.Setup.Action[numAction].Operation.ResponseId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationSourceId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.SourceId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.SourceId", resource.Setup.Action[numAction].Operation.SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationTargetId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.TargetId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.TargetId", resource.Setup.Action[numAction].Operation.TargetId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationUrl(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Url", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.Url", resource.Setup.Action[numAction].Operation.Url, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderField(numAction int, numRequestHeader int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) || numRequestHeader >= len(resource.Setup.Action[numAction].Operation.RequestHeader) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Field", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Field", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Field, htmlAttrs)
}
func (resource *TestScript) T_SetupActionOperationRequestHeaderValue(numAction int, numRequestHeader int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) || numRequestHeader >= len(resource.Setup.Action[numAction].Operation.RequestHeader) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Value", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Operation.RequestHeader["+strconv.Itoa(numRequestHeader)+"].Value", &resource.Setup.Action[numAction].Operation.RequestHeader[numRequestHeader].Value, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertLabel(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Label", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Label", resource.Setup.Action[numAction].Assert.Label, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertDescription(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Description", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Description", resource.Setup.Action[numAction].Assert.Description, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertDirection(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAssert_direction_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Direction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Direction", resource.Setup.Action[numAction].Assert.Direction, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourceId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourceId", resource.Setup.Action[numAction].Assert.CompareToSourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourceExpression(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourceExpression", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourceExpression", resource.Setup.Action[numAction].Assert.CompareToSourceExpression, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertCompareToSourcePath(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourcePath", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.CompareToSourcePath", resource.Setup.Action[numAction].Assert.CompareToSourcePath, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertContentType(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.ContentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.ContentType", resource.Setup.Action[numAction].Assert.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertExpression(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Expression", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Expression", resource.Setup.Action[numAction].Assert.Expression, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertHeaderField(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.HeaderField", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.HeaderField", resource.Setup.Action[numAction].Assert.HeaderField, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertMinimumId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.MinimumId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.MinimumId", resource.Setup.Action[numAction].Assert.MinimumId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertNavigationLinks(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.NavigationLinks", nil, htmlAttrs)
	}
	return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.NavigationLinks", resource.Setup.Action[numAction].Assert.NavigationLinks, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertOperator(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAssert_operator_codes

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Operator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Operator", resource.Setup.Action[numAction].Assert.Operator, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertPath(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Path", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Path", resource.Setup.Action[numAction].Assert.Path, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertRequestMethod(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.RequestMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.RequestMethod", resource.Setup.Action[numAction].Assert.RequestMethod, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertRequestURL(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.RequestURL", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.RequestURL", resource.Setup.Action[numAction].Assert.RequestURL, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResource(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Resource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Resource", resource.Setup.Action[numAction].Assert.Resource, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResponse(numAction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAssert_response_code_types

	if resource == nil || numAction >= len(resource.Setup.Action) {
		return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Response", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SetupAction["+strconv.Itoa(numAction)+"].Assert.Response", resource.Setup.Action[numAction].Assert.Response, optionsValueSet, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertResponseCode(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.ResponseCode", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.ResponseCode", resource.Setup.Action[numAction].Assert.ResponseCode, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertSourceId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.SourceId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.SourceId", resource.Setup.Action[numAction].Assert.SourceId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertValidateProfileId(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.ValidateProfileId", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.ValidateProfileId", resource.Setup.Action[numAction].Assert.ValidateProfileId, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertValue(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Value", nil, htmlAttrs)
	}
	return StringInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.Value", resource.Setup.Action[numAction].Assert.Value, htmlAttrs)
}
func (resource *TestScript) T_SetupActionAssertWarningOnly(numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Setup.Action) {
		return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.WarningOnly", nil, htmlAttrs)
	}
	return BoolInput("SetupAction["+strconv.Itoa(numAction)+"].Assert.WarningOnly", &resource.Setup.Action[numAction].Assert.WarningOnly, htmlAttrs)
}
func (resource *TestScript) T_TestName(numTest int, htmlAttrs string) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("Test["+strconv.Itoa(numTest)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Test["+strconv.Itoa(numTest)+"]Name", resource.Test[numTest].Name, htmlAttrs)
}
func (resource *TestScript) T_TestDescription(numTest int, htmlAttrs string) templ.Component {
	if resource == nil || numTest >= len(resource.Test) {
		return StringInput("Test["+strconv.Itoa(numTest)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Test["+strconv.Itoa(numTest)+"]Description", resource.Test[numTest].Description, htmlAttrs)
}
