package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Date              *string                 `json:"date,omitempty"`
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

func (resource *TestScript) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *TestScript) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *TestScript) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *TestScript) T_OriginProfile(numOrigin int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Origin) >= numOrigin {
		return CodingSelect("profile", nil, optionsValueSet)
	}
	return CodingSelect("profile", &resource.Origin[numOrigin].Profile, optionsValueSet)
}
func (resource *TestScript) T_DestinationProfile(numDestination int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Destination) >= numDestination {
		return CodingSelect("profile", nil, optionsValueSet)
	}
	return CodingSelect("profile", &resource.Destination[numDestination].Profile, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", resource.Setup.Action[numAction].Operation.Type, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationResource(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("resource", nil, optionsValueSet)
	}
	return CodeSelect("resource", resource.Setup.Action[numAction].Operation.Resource, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationAccept(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("accept", nil, optionsValueSet)
	}
	return CodeSelect("accept", resource.Setup.Action[numAction].Operation.Accept, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationContentType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("contentType", nil, optionsValueSet)
	}
	return CodeSelect("contentType", resource.Setup.Action[numAction].Operation.ContentType, optionsValueSet)
}
func (resource *TestScript) T_SetupActionOperationMethod(numAction int) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("method", nil, optionsValueSet)
	}
	return CodeSelect("method", resource.Setup.Action[numAction].Operation.Method, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertDirection(numAction int) templ.Component {
	optionsValueSet := VSAssert_direction_codes

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("direction", nil, optionsValueSet)
	}
	return CodeSelect("direction", resource.Setup.Action[numAction].Assert.Direction, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertContentType(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("contentType", nil, optionsValueSet)
	}
	return CodeSelect("contentType", resource.Setup.Action[numAction].Assert.ContentType, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertOperator(numAction int) templ.Component {
	optionsValueSet := VSAssert_operator_codes

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("operator", nil, optionsValueSet)
	}
	return CodeSelect("operator", resource.Setup.Action[numAction].Assert.Operator, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertRequestMethod(numAction int) templ.Component {
	optionsValueSet := VSHttp_operations

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("requestMethod", nil, optionsValueSet)
	}
	return CodeSelect("requestMethod", resource.Setup.Action[numAction].Assert.RequestMethod, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertResource(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("resource", nil, optionsValueSet)
	}
	return CodeSelect("resource", resource.Setup.Action[numAction].Assert.Resource, optionsValueSet)
}
func (resource *TestScript) T_SetupActionAssertResponse(numAction int) templ.Component {
	optionsValueSet := VSAssert_response_code_types

	if resource == nil && len(resource.Setup.Action) >= numAction {
		return CodeSelect("response", nil, optionsValueSet)
	}
	return CodeSelect("response", resource.Setup.Action[numAction].Assert.Response, optionsValueSet)
}
