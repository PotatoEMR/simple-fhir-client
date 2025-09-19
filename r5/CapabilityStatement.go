package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                     *string                            `json:"id,omitempty"`
	Meta                   *Meta                              `json:"meta,omitempty"`
	ImplicitRules          *string                            `json:"implicitRules,omitempty"`
	Language               *string                            `json:"language,omitempty"`
	Text                   *Narrative                         `json:"text,omitempty"`
	Contained              []Resource                         `json:"contained,omitempty"`
	Extension              []Extension                        `json:"extension,omitempty"`
	ModifierExtension      []Extension                        `json:"modifierExtension,omitempty"`
	Url                    *string                            `json:"url,omitempty"`
	Identifier             []Identifier                       `json:"identifier,omitempty"`
	Version                *string                            `json:"version,omitempty"`
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                            `json:"name,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	Status                 string                             `json:"status"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	Date                   FhirDateTime                       `json:"date"`
	Publisher              *string                            `json:"publisher,omitempty"`
	Contact                []ContactDetail                    `json:"contact,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	UseContext             []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose                *string                            `json:"purpose,omitempty"`
	Copyright              *string                            `json:"copyright,omitempty"`
	CopyrightLabel         *string                            `json:"copyrightLabel,omitempty"`
	Kind                   string                             `json:"kind"`
	Instantiates           []string                           `json:"instantiates,omitempty"`
	Imports                []string                           `json:"imports,omitempty"`
	Software               *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation         *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FhirVersion            string                             `json:"fhirVersion"`
	Format                 []string                           `json:"format"`
	PatchFormat            []string                           `json:"patchFormat,omitempty"`
	AcceptLanguage         []string                           `json:"acceptLanguage,omitempty"`
	ImplementationGuide    []string                           `json:"implementationGuide,omitempty"`
	Rest                   []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging              []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document               []CapabilityStatementDocument      `json:"document,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementSoftware struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Name              string        `json:"name"`
	Version           *string       `json:"version,omitempty"`
	ReleaseDate       *FhirDateTime `json:"releaseDate,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
	Custodian         *Reference  `json:"custodian,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
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

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementRestSecurity struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Cors              *bool             `json:"cors,omitempty"`
	Service           []CodeableConcept `json:"service,omitempty"`
	Description       *string           `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
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
	ConditionalPatch  *bool                                        `json:"conditionalPatch,omitempty"`
	ConditionalDelete *string                                      `json:"conditionalDelete,omitempty"`
	ReferencePolicy   []string                                     `json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        *string     `json:"definition,omitempty"`
	Type              string      `json:"type"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        string      `json:"definition"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementRestInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementMessaging struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache     *int                                           `json:"reliableCache,omitempty"`
	Documentation     *string                                        `json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Protocol          Coding      `json:"protocol"`
	Address           string      `json:"address"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Definition        string      `json:"definition"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CapabilityStatement
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
func (r CapabilityStatement) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CapabilityStatement/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CapabilityStatement"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CapabilityStatement) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *CapabilityStatement) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *CapabilityStatement) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *CapabilityStatement) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *CapabilityStatement) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *CapabilityStatement) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *CapabilityStatement) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *CapabilityStatement) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *CapabilityStatement) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *CapabilityStatement) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *CapabilityStatement) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *CapabilityStatement) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *CapabilityStatement) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Instantiates(numInstantiates int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *CapabilityStatement) T_Imports(numImports int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImports >= len(resource.Imports) {
		return StringInput("imports["+strconv.Itoa(numImports)+"]", nil, htmlAttrs)
	}
	return StringInput("imports["+strconv.Itoa(numImports)+"]", &resource.Imports[numImports], htmlAttrs)
}
func (resource *CapabilityStatement) T_FhirVersion(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("fhirVersion", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("fhirVersion", &resource.FhirVersion, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Format(numFormat int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFormat >= len(resource.Format) {
		return CodeSelect("format["+strconv.Itoa(numFormat)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("format["+strconv.Itoa(numFormat)+"]", &resource.Format[numFormat], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_PatchFormat(numPatchFormat int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatchFormat >= len(resource.PatchFormat) {
		return CodeSelect("patchFormat["+strconv.Itoa(numPatchFormat)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("patchFormat["+strconv.Itoa(numPatchFormat)+"]", &resource.PatchFormat[numPatchFormat], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_AcceptLanguage(numAcceptLanguage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAcceptLanguage >= len(resource.AcceptLanguage) {
		return CodeSelect("acceptLanguage["+strconv.Itoa(numAcceptLanguage)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("acceptLanguage["+strconv.Itoa(numAcceptLanguage)+"]", &resource.AcceptLanguage[numAcceptLanguage], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationGuide(numImplementationGuide int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImplementationGuide >= len(resource.ImplementationGuide) {
		return StringInput("implementationGuide["+strconv.Itoa(numImplementationGuide)+"]", nil, htmlAttrs)
	}
	return StringInput("implementationGuide["+strconv.Itoa(numImplementationGuide)+"]", &resource.ImplementationGuide[numImplementationGuide], htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("software.name", nil, htmlAttrs)
	}
	return StringInput("software.name", &resource.Software.Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareVersion(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("software.version", nil, htmlAttrs)
	}
	return StringInput("software.version", resource.Software.Version, htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareReleaseDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("software.releaseDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("software.releaseDate", resource.Software.ReleaseDate, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("implementation.description", nil, htmlAttrs)
	}
	return StringInput("implementation.description", &resource.Implementation.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationUrl(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("implementation.url", nil, htmlAttrs)
	}
	return StringInput("implementation.url", resource.Implementation.Url, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationCustodian(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("implementation.custodian", nil, htmlAttrs)
	}
	return ReferenceInput("implementation.custodian", resource.Implementation.Custodian, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestMode(numRest int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRestful_capability_mode

	if resource == nil || numRest >= len(resource.Rest) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].mode", &resource.Rest[numRest].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestDocumentation(numRest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].documentation", resource.Rest[numRest].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestCompartment(numRest int, numCompartment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numCompartment >= len(resource.Rest[numRest].Compartment) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].compartment["+strconv.Itoa(numCompartment)+"]", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].compartment["+strconv.Itoa(numCompartment)+"]", &resource.Rest[numRest].Compartment[numCompartment], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityCors(numRest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].security.cors", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].security.cors", resource.Rest[numRest].Security.Cors, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityService(numRest int, numService int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numService >= len(resource.Rest[numRest].Security.Service) {
		return CodeableConceptSelect("rest["+strconv.Itoa(numRest)+"].security.service["+strconv.Itoa(numService)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("rest["+strconv.Itoa(numRest)+"].security.service["+strconv.Itoa(numService)+"]", &resource.Rest[numRest].Security.Service[numService], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityDescription(numRest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].security.description", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].security.description", resource.Rest[numRest].Security.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceType(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].type", &resource.Rest[numRest].Resource[numResource].Type, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceProfile(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].profile", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].profile", resource.Rest[numRest].Resource[numResource].Profile, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSupportedProfile(numRest int, numResource int, numSupportedProfile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSupportedProfile >= len(resource.Rest[numRest].Resource[numResource].SupportedProfile) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].supportedProfile["+strconv.Itoa(numSupportedProfile)+"]", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].supportedProfile["+strconv.Itoa(numSupportedProfile)+"]", &resource.Rest[numRest].Resource[numResource].SupportedProfile[numSupportedProfile], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceDocumentation(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].documentation", resource.Rest[numRest].Resource[numResource].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceVersioning(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVersioning_policy

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].versioning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].versioning", resource.Rest[numRest].Resource[numResource].Versioning, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceReadHistory(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].readHistory", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].readHistory", resource.Rest[numRest].Resource[numResource].ReadHistory, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceUpdateCreate(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].updateCreate", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].updateCreate", resource.Rest[numRest].Resource[numResource].UpdateCreate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalCreate(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalCreate", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalCreate", resource.Rest[numRest].Resource[numResource].ConditionalCreate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalRead(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConditional_read_status

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalRead", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalRead", resource.Rest[numRest].Resource[numResource].ConditionalRead, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalUpdate(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalUpdate", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalUpdate", resource.Rest[numRest].Resource[numResource].ConditionalUpdate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalPatch(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalPatch", nil, htmlAttrs)
	}
	return BoolInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalPatch", resource.Rest[numRest].Resource[numResource].ConditionalPatch, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalDelete(numRest int, numResource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConditional_delete_status

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalDelete", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].conditionalDelete", resource.Rest[numRest].Resource[numResource].ConditionalDelete, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceReferencePolicy(numRest int, numResource int, numReferencePolicy int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReference_handling_policy

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numReferencePolicy >= len(resource.Rest[numRest].Resource[numResource].ReferencePolicy) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].referencePolicy["+strconv.Itoa(numReferencePolicy)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].referencePolicy["+strconv.Itoa(numReferencePolicy)+"]", &resource.Rest[numRest].Resource[numResource].ReferencePolicy[numReferencePolicy], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchInclude(numRest int, numResource int, numSearchInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchInclude >= len(resource.Rest[numRest].Resource[numResource].SearchInclude) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchInclude["+strconv.Itoa(numSearchInclude)+"]", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchInclude["+strconv.Itoa(numSearchInclude)+"]", &resource.Rest[numRest].Resource[numResource].SearchInclude[numSearchInclude], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchRevInclude(numRest int, numResource int, numSearchRevInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchRevInclude >= len(resource.Rest[numRest].Resource[numResource].SearchRevInclude) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchRevInclude["+strconv.Itoa(numSearchRevInclude)+"]", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchRevInclude["+strconv.Itoa(numSearchRevInclude)+"]", &resource.Rest[numRest].Resource[numResource].SearchRevInclude[numSearchRevInclude], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceInteractionCode(numRest int, numResource int, numInteraction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSType_restful_interaction

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numInteraction >= len(resource.Rest[numRest].Resource[numResource].Interaction) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].interaction["+strconv.Itoa(numInteraction)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].interaction["+strconv.Itoa(numInteraction)+"].code", &resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Code, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceInteractionDocumentation(numRest int, numResource int, numInteraction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numInteraction >= len(resource.Rest[numRest].Resource[numResource].Interaction) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].interaction["+strconv.Itoa(numInteraction)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].interaction["+strconv.Itoa(numInteraction)+"].documentation", resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamName(numRest int, numResource int, numSearchParam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].name", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].name", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDefinition(numRest int, numResource int, numSearchParam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].definition", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].definition", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamType(numRest int, numResource int, numSearchParam int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].type", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Type, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDocumentation(numRest int, numResource int, numSearchParam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].searchParam["+strconv.Itoa(numSearchParam)+"].documentation", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationName(numRest int, numResource int, numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].name", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].name", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationDefinition(numRest int, numResource int, numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].definition", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].definition", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationDocumentation(numRest int, numResource int, numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].resource["+strconv.Itoa(numResource)+"].operation["+strconv.Itoa(numOperation)+"].documentation", resource.Rest[numRest].Resource[numResource].Operation[numOperation].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestInteractionCode(numRest int, numInteraction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSystem_restful_interaction

	if resource == nil || numRest >= len(resource.Rest) || numInteraction >= len(resource.Rest[numRest].Interaction) {
		return CodeSelect("rest["+strconv.Itoa(numRest)+"].interaction["+strconv.Itoa(numInteraction)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("rest["+strconv.Itoa(numRest)+"].interaction["+strconv.Itoa(numInteraction)+"].code", &resource.Rest[numRest].Interaction[numInteraction].Code, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestInteractionDocumentation(numRest int, numInteraction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRest >= len(resource.Rest) || numInteraction >= len(resource.Rest[numRest].Interaction) {
		return StringInput("rest["+strconv.Itoa(numRest)+"].interaction["+strconv.Itoa(numInteraction)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("rest["+strconv.Itoa(numRest)+"].interaction["+strconv.Itoa(numInteraction)+"].documentation", resource.Rest[numRest].Interaction[numInteraction].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingReliableCache(numMessaging int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMessaging >= len(resource.Messaging) {
		return IntInput("messaging["+strconv.Itoa(numMessaging)+"].reliableCache", nil, htmlAttrs)
	}
	return IntInput("messaging["+strconv.Itoa(numMessaging)+"].reliableCache", resource.Messaging[numMessaging].ReliableCache, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingDocumentation(numMessaging int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMessaging >= len(resource.Messaging) {
		return StringInput("messaging["+strconv.Itoa(numMessaging)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("messaging["+strconv.Itoa(numMessaging)+"].documentation", resource.Messaging[numMessaging].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingEndpointProtocol(numMessaging int, numEndpoint int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMessaging >= len(resource.Messaging) || numEndpoint >= len(resource.Messaging[numMessaging].Endpoint) {
		return CodingSelect("messaging["+strconv.Itoa(numMessaging)+"].endpoint["+strconv.Itoa(numEndpoint)+"].protocol", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("messaging["+strconv.Itoa(numMessaging)+"].endpoint["+strconv.Itoa(numEndpoint)+"].protocol", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Protocol, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingEndpointAddress(numMessaging int, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMessaging >= len(resource.Messaging) || numEndpoint >= len(resource.Messaging[numMessaging].Endpoint) {
		return StringInput("messaging["+strconv.Itoa(numMessaging)+"].endpoint["+strconv.Itoa(numEndpoint)+"].address", nil, htmlAttrs)
	}
	return StringInput("messaging["+strconv.Itoa(numMessaging)+"].endpoint["+strconv.Itoa(numEndpoint)+"].address", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Address, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageMode(numMessaging int, numSupportedMessage int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEvent_capability_mode

	if resource == nil || numMessaging >= len(resource.Messaging) || numSupportedMessage >= len(resource.Messaging[numMessaging].SupportedMessage) {
		return CodeSelect("messaging["+strconv.Itoa(numMessaging)+"].supportedMessage["+strconv.Itoa(numSupportedMessage)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("messaging["+strconv.Itoa(numMessaging)+"].supportedMessage["+strconv.Itoa(numSupportedMessage)+"].mode", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageDefinition(numMessaging int, numSupportedMessage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMessaging >= len(resource.Messaging) || numSupportedMessage >= len(resource.Messaging[numMessaging].SupportedMessage) {
		return StringInput("messaging["+strconv.Itoa(numMessaging)+"].supportedMessage["+strconv.Itoa(numSupportedMessage)+"].definition", nil, htmlAttrs)
	}
	return StringInput("messaging["+strconv.Itoa(numMessaging)+"].supportedMessage["+strconv.Itoa(numSupportedMessage)+"].definition", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentMode(numDocument int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDocument_mode

	if resource == nil || numDocument >= len(resource.Document) {
		return CodeSelect("document["+strconv.Itoa(numDocument)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("document["+strconv.Itoa(numDocument)+"].mode", &resource.Document[numDocument].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentDocumentation(numDocument int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDocument >= len(resource.Document) {
		return StringInput("document["+strconv.Itoa(numDocument)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("document["+strconv.Itoa(numDocument)+"].documentation", resource.Document[numDocument].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentProfile(numDocument int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDocument >= len(resource.Document) {
		return StringInput("document["+strconv.Itoa(numDocument)+"].profile", nil, htmlAttrs)
	}
	return StringInput("document["+strconv.Itoa(numDocument)+"].profile", &resource.Document[numDocument].Profile, htmlAttrs)
}
