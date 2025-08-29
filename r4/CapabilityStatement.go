package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                  *string                            `json:"id,omitempty"`
	Meta                *Meta                              `json:"meta,omitempty"`
	ImplicitRules       *string                            `json:"implicitRules,omitempty"`
	Language            *string                            `json:"language,omitempty"`
	Text                *Narrative                         `json:"text,omitempty"`
	Contained           []Resource                         `json:"contained,omitempty"`
	Extension           []Extension                        `json:"extension,omitempty"`
	ModifierExtension   []Extension                        `json:"modifierExtension,omitempty"`
	Url                 *string                            `json:"url,omitempty"`
	Version             *string                            `json:"version,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	Title               *string                            `json:"title,omitempty"`
	Status              string                             `json:"status"`
	Experimental        *bool                              `json:"experimental,omitempty"`
	Date                string                             `json:"date"`
	Publisher           *string                            `json:"publisher,omitempty"`
	Contact             []ContactDetail                    `json:"contact,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	UseContext          []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose             *string                            `json:"purpose,omitempty"`
	Copyright           *string                            `json:"copyright,omitempty"`
	Kind                string                             `json:"kind"`
	Instantiates        []string                           `json:"instantiates,omitempty"`
	Imports             []string                           `json:"imports,omitempty"`
	Software            *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FhirVersion         string                             `json:"fhirVersion"`
	Format              []string                           `json:"format"`
	PatchFormat         []string                           `json:"patchFormat,omitempty"`
	ImplementationGuide []string                           `json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging           []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document            []CapabilityStatementDocument      `json:"document,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementSoftware struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
	ReleaseDate       *string     `json:"releaseDate,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
	Custodian         *Reference  `json:"custodian,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRest struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Mode              string                               `json:"mode"`
	Documentation     *string                              `json:"documentation,omitempty"`
	Security          *CapabilityStatementRestSecurity     `json:"security,omitempty"`
	Resource          []CapabilityStatementRestResource    `json:"resource,omitempty"`
	Interaction       []CapabilityStatementRestInteraction `json:"interaction,omitempty"`
	Compartment       []string                             `json:"compartment,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestSecurity struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Cors              *bool             `json:"cors,omitempty"`
	Service           []CodeableConcept `json:"service,omitempty"`
	Description       *string           `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResource struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Type              string                                       `json:"type"`
	Profile           *string                                      `json:"profile,omitempty"`
	SupportedProfile  []string                                     `json:"supportedProfile,omitempty"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	Versioning        *string                                      `json:"versioning,omitempty"`
	ReadHistory       *bool                                        `json:"readHistory,omitempty"`
	UpdateCreate      *bool                                        `json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                        `json:"conditionalCreate,omitempty"`
	ConditionalRead   *string                                      `json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                        `json:"conditionalUpdate,omitempty"`
	ConditionalDelete *string                                      `json:"conditionalDelete,omitempty"`
	ReferencePolicy   []string                                     `json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        *string     `json:"definition,omitempty"`
	Type              string      `json:"type"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        string      `json:"definition"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementRestInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementMessaging struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache     *int                                           `json:"reliableCache,omitempty"`
	Documentation     *string                                        `json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Protocol          Coding      `json:"protocol"`
	Address           string      `json:"address"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Definition        string      `json:"definition"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CapabilityStatement
type CapabilityStatementDocument struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Documentation     *string     `json:"documentation,omitempty"`
	Profile           string      `json:"profile"`
}

type OtherCapabilityStatement CapabilityStatement

// on convert struct to json, automatically add resourceType=CapabilityStatement
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
}
func (resource *CapabilityStatement) CapabilityStatementLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementKind() templ.Component {
	optionsValueSet := VSCapability_statement_kind
	currentVal := ""
	if resource != nil {
		currentVal = resource.Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementFhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version
	currentVal := ""
	if resource != nil {
		currentVal = resource.FhirVersion
	}
	return CodeSelect("fhirVersion", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementFormat(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = resource.Format[0]
	}
	return CodeSelect("format", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementPatchFormat(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = resource.PatchFormat[0]
	}
	return CodeSelect("patchFormat", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestMode(numRest int) templ.Component {
	optionsValueSet := VSRestful_capability_mode
	currentVal := ""
	if resource != nil && len(resource.Rest) >= numRest {
		currentVal = resource.Rest[numRest].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceType(numRest int, numResource int) templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource) >= numResource {
		currentVal = resource.Rest[numRest].Resource[numResource].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceVersioning(numRest int, numResource int) templ.Component {
	optionsValueSet := VSVersioning_policy
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource) >= numResource {
		currentVal = *resource.Rest[numRest].Resource[numResource].Versioning
	}
	return CodeSelect("versioning", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceConditionalRead(numRest int, numResource int) templ.Component {
	optionsValueSet := VSConditional_read_status
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource) >= numResource {
		currentVal = *resource.Rest[numRest].Resource[numResource].ConditionalRead
	}
	return CodeSelect("conditionalRead", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceConditionalDelete(numRest int, numResource int) templ.Component {
	optionsValueSet := VSConditional_delete_status
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource) >= numResource {
		currentVal = *resource.Rest[numRest].Resource[numResource].ConditionalDelete
	}
	return CodeSelect("conditionalDelete", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceReferencePolicy(numRest int, numResource int) templ.Component {
	optionsValueSet := VSReference_handling_policy
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource) >= numResource {
		currentVal = resource.Rest[numRest].Resource[numResource].ReferencePolicy[0]
	}
	return CodeSelect("referencePolicy", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceInteractionCode(numRest int, numResource int, numInteraction int) templ.Component {
	optionsValueSet := VSType_restful_interaction
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource[numResource].Interaction) >= numInteraction {
		currentVal = resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestResourceSearchParamType(numRest int, numResource int, numSearchParam int) templ.Component {
	optionsValueSet := VSSearch_param_type
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		currentVal = resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementRestInteractionCode(numRest int, numInteraction int) templ.Component {
	optionsValueSet := VSSystem_restful_interaction
	currentVal := ""
	if resource != nil && len(resource.Rest[numRest].Interaction) >= numInteraction {
		currentVal = resource.Rest[numRest].Interaction[numInteraction].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementMessagingSupportedMessageMode(numMessaging int, numSupportedMessage int) templ.Component {
	optionsValueSet := VSEvent_capability_mode
	currentVal := ""
	if resource != nil && len(resource.Messaging[numMessaging].SupportedMessage) >= numSupportedMessage {
		currentVal = resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
func (resource *CapabilityStatement) CapabilityStatementDocumentMode(numDocument int) templ.Component {
	optionsValueSet := VSDocument_mode
	currentVal := ""
	if resource != nil && len(resource.Document) >= numDocument {
		currentVal = resource.Document[numDocument].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
