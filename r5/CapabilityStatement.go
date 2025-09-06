package r5

//generated with command go run ./bultaoreune -nodownload
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
	Date                   string                             `json:"date"`
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
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
	ReleaseDate       *string     `json:"releaseDate,omitempty"`
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

func (resource *CapabilityStatement) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Id", nil)
	}
	return StringInput("CapabilityStatement.Id", resource.Id)
}
func (resource *CapabilityStatement) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.ImplicitRules", nil)
	}
	return StringInput("CapabilityStatement.ImplicitRules", resource.ImplicitRules)
}
func (resource *CapabilityStatement) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CapabilityStatement.Language", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Language", resource.Language, optionsValueSet)
}
func (resource *CapabilityStatement) T_Url() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Url", nil)
	}
	return StringInput("CapabilityStatement.Url", resource.Url)
}
func (resource *CapabilityStatement) T_Version() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Version", nil)
	}
	return StringInput("CapabilityStatement.Version", resource.Version)
}
func (resource *CapabilityStatement) T_Name() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Name", nil)
	}
	return StringInput("CapabilityStatement.Name", resource.Name)
}
func (resource *CapabilityStatement) T_Title() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Title", nil)
	}
	return StringInput("CapabilityStatement.Title", resource.Title)
}
func (resource *CapabilityStatement) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CapabilityStatement.Status", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Status", &resource.Status, optionsValueSet)
}
func (resource *CapabilityStatement) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("CapabilityStatement.Experimental", nil)
	}
	return BoolInput("CapabilityStatement.Experimental", resource.Experimental)
}
func (resource *CapabilityStatement) T_Date() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Date", nil)
	}
	return StringInput("CapabilityStatement.Date", &resource.Date)
}
func (resource *CapabilityStatement) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Publisher", nil)
	}
	return StringInput("CapabilityStatement.Publisher", resource.Publisher)
}
func (resource *CapabilityStatement) T_Description() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Description", nil)
	}
	return StringInput("CapabilityStatement.Description", resource.Description)
}
func (resource *CapabilityStatement) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("CapabilityStatement.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CapabilityStatement.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *CapabilityStatement) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Purpose", nil)
	}
	return StringInput("CapabilityStatement.Purpose", resource.Purpose)
}
func (resource *CapabilityStatement) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Copyright", nil)
	}
	return StringInput("CapabilityStatement.Copyright", resource.Copyright)
}
func (resource *CapabilityStatement) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.CopyrightLabel", nil)
	}
	return StringInput("CapabilityStatement.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *CapabilityStatement) T_Kind() templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("CapabilityStatement.Kind", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Kind", &resource.Kind, optionsValueSet)
}
func (resource *CapabilityStatement) T_Instantiates(numInstantiates int) templ.Component {

	if resource == nil || len(resource.Instantiates) >= numInstantiates {
		return StringInput("CapabilityStatement.Instantiates["+strconv.Itoa(numInstantiates)+"]", nil)
	}
	return StringInput("CapabilityStatement.Instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates])
}
func (resource *CapabilityStatement) T_Imports(numImports int) templ.Component {

	if resource == nil || len(resource.Imports) >= numImports {
		return StringInput("CapabilityStatement.Imports["+strconv.Itoa(numImports)+"]", nil)
	}
	return StringInput("CapabilityStatement.Imports["+strconv.Itoa(numImports)+"]", &resource.Imports[numImports])
}
func (resource *CapabilityStatement) T_FhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("CapabilityStatement.FhirVersion", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.FhirVersion", &resource.FhirVersion, optionsValueSet)
}
func (resource *CapabilityStatement) T_Format(numFormat int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Format) >= numFormat {
		return CodeSelect("CapabilityStatement.Format["+strconv.Itoa(numFormat)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Format["+strconv.Itoa(numFormat)+"]", &resource.Format[numFormat], optionsValueSet)
}
func (resource *CapabilityStatement) T_PatchFormat(numPatchFormat int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PatchFormat) >= numPatchFormat {
		return CodeSelect("CapabilityStatement.PatchFormat["+strconv.Itoa(numPatchFormat)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.PatchFormat["+strconv.Itoa(numPatchFormat)+"]", &resource.PatchFormat[numPatchFormat], optionsValueSet)
}
func (resource *CapabilityStatement) T_AcceptLanguage(numAcceptLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AcceptLanguage) >= numAcceptLanguage {
		return CodeSelect("CapabilityStatement.AcceptLanguage["+strconv.Itoa(numAcceptLanguage)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.AcceptLanguage["+strconv.Itoa(numAcceptLanguage)+"]", &resource.AcceptLanguage[numAcceptLanguage], optionsValueSet)
}
func (resource *CapabilityStatement) T_ImplementationGuide(numImplementationGuide int) templ.Component {

	if resource == nil || len(resource.ImplementationGuide) >= numImplementationGuide {
		return StringInput("CapabilityStatement.ImplementationGuide["+strconv.Itoa(numImplementationGuide)+"]", nil)
	}
	return StringInput("CapabilityStatement.ImplementationGuide["+strconv.Itoa(numImplementationGuide)+"]", &resource.ImplementationGuide[numImplementationGuide])
}
func (resource *CapabilityStatement) T_SoftwareId() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.Id", nil)
	}
	return StringInput("CapabilityStatement.Software.Id", resource.Software.Id)
}
func (resource *CapabilityStatement) T_SoftwareName() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.Name", nil)
	}
	return StringInput("CapabilityStatement.Software.Name", &resource.Software.Name)
}
func (resource *CapabilityStatement) T_SoftwareVersion() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.Version", nil)
	}
	return StringInput("CapabilityStatement.Software.Version", resource.Software.Version)
}
func (resource *CapabilityStatement) T_SoftwareReleaseDate() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.ReleaseDate", nil)
	}
	return StringInput("CapabilityStatement.Software.ReleaseDate", resource.Software.ReleaseDate)
}
func (resource *CapabilityStatement) T_ImplementationId() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Implementation.Id", nil)
	}
	return StringInput("CapabilityStatement.Implementation.Id", resource.Implementation.Id)
}
func (resource *CapabilityStatement) T_ImplementationDescription() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Implementation.Description", nil)
	}
	return StringInput("CapabilityStatement.Implementation.Description", &resource.Implementation.Description)
}
func (resource *CapabilityStatement) T_ImplementationUrl() templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Implementation.Url", nil)
	}
	return StringInput("CapabilityStatement.Implementation.Url", resource.Implementation.Url)
}
func (resource *CapabilityStatement) T_RestId(numRest int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Id", resource.Rest[numRest].Id)
}
func (resource *CapabilityStatement) T_RestMode(numRest int) templ.Component {
	optionsValueSet := VSRestful_capability_mode

	if resource == nil || len(resource.Rest) >= numRest {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Mode", &resource.Rest[numRest].Mode, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestDocumentation(numRest int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Documentation", resource.Rest[numRest].Documentation)
}
func (resource *CapabilityStatement) T_RestCompartment(numRest int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Compartment) >= numCompartment {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Compartment["+strconv.Itoa(numCompartment)+"]", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Compartment["+strconv.Itoa(numCompartment)+"]", &resource.Rest[numRest].Compartment[numCompartment])
}
func (resource *CapabilityStatement) T_RestSecurityId(numRest int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Id", resource.Rest[numRest].Security.Id)
}
func (resource *CapabilityStatement) T_RestSecurityCors(numRest int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Cors", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Cors", resource.Rest[numRest].Security.Cors)
}
func (resource *CapabilityStatement) T_RestSecurityService(numRest int, numService int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Security.Service) >= numService {
		return CodeableConceptSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Service["+strconv.Itoa(numService)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Service["+strconv.Itoa(numService)+"]", &resource.Rest[numRest].Security.Service[numService], optionsValueSet)
}
func (resource *CapabilityStatement) T_RestSecurityDescription(numRest int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Description", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Security.Description", resource.Rest[numRest].Security.Description)
}
func (resource *CapabilityStatement) T_RestResourceId(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Id", resource.Rest[numRest].Resource[numResource].Id)
}
func (resource *CapabilityStatement) T_RestResourceType(numRest int, numResource int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Type", &resource.Rest[numRest].Resource[numResource].Type, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceProfile(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Profile", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Profile", resource.Rest[numRest].Resource[numResource].Profile)
}
func (resource *CapabilityStatement) T_RestResourceSupportedProfile(numRest int, numResource int, numSupportedProfile int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SupportedProfile) >= numSupportedProfile {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SupportedProfile["+strconv.Itoa(numSupportedProfile)+"]", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SupportedProfile["+strconv.Itoa(numSupportedProfile)+"]", &resource.Rest[numRest].Resource[numResource].SupportedProfile[numSupportedProfile])
}
func (resource *CapabilityStatement) T_RestResourceDocumentation(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Documentation", resource.Rest[numRest].Resource[numResource].Documentation)
}
func (resource *CapabilityStatement) T_RestResourceVersioning(numRest int, numResource int) templ.Component {
	optionsValueSet := VSVersioning_policy

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Versioning", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Versioning", resource.Rest[numRest].Resource[numResource].Versioning, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceReadHistory(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ReadHistory", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ReadHistory", resource.Rest[numRest].Resource[numResource].ReadHistory)
}
func (resource *CapabilityStatement) T_RestResourceUpdateCreate(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].UpdateCreate", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].UpdateCreate", resource.Rest[numRest].Resource[numResource].UpdateCreate)
}
func (resource *CapabilityStatement) T_RestResourceConditionalCreate(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalCreate", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalCreate", resource.Rest[numRest].Resource[numResource].ConditionalCreate)
}
func (resource *CapabilityStatement) T_RestResourceConditionalRead(numRest int, numResource int) templ.Component {
	optionsValueSet := VSConditional_read_status

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalRead", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalRead", resource.Rest[numRest].Resource[numResource].ConditionalRead, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceConditionalUpdate(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalUpdate", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalUpdate", resource.Rest[numRest].Resource[numResource].ConditionalUpdate)
}
func (resource *CapabilityStatement) T_RestResourceConditionalPatch(numRest int, numResource int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalPatch", nil)
	}
	return BoolInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalPatch", resource.Rest[numRest].Resource[numResource].ConditionalPatch)
}
func (resource *CapabilityStatement) T_RestResourceConditionalDelete(numRest int, numResource int) templ.Component {
	optionsValueSet := VSConditional_delete_status

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalDelete", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ConditionalDelete", resource.Rest[numRest].Resource[numResource].ConditionalDelete, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceReferencePolicy(numRest int, numResource int, numReferencePolicy int) templ.Component {
	optionsValueSet := VSReference_handling_policy

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].ReferencePolicy) >= numReferencePolicy {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ReferencePolicy["+strconv.Itoa(numReferencePolicy)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].ReferencePolicy["+strconv.Itoa(numReferencePolicy)+"]", &resource.Rest[numRest].Resource[numResource].ReferencePolicy[numReferencePolicy], optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceSearchInclude(numRest int, numResource int, numSearchInclude int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchInclude) >= numSearchInclude {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchInclude["+strconv.Itoa(numSearchInclude)+"]", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchInclude["+strconv.Itoa(numSearchInclude)+"]", &resource.Rest[numRest].Resource[numResource].SearchInclude[numSearchInclude])
}
func (resource *CapabilityStatement) T_RestResourceSearchRevInclude(numRest int, numResource int, numSearchRevInclude int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchRevInclude) >= numSearchRevInclude {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchRevInclude["+strconv.Itoa(numSearchRevInclude)+"]", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchRevInclude["+strconv.Itoa(numSearchRevInclude)+"]", &resource.Rest[numRest].Resource[numResource].SearchRevInclude[numSearchRevInclude])
}
func (resource *CapabilityStatement) T_RestResourceInteractionId(numRest int, numResource int, numInteraction int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Interaction) >= numInteraction {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Id", resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Id)
}
func (resource *CapabilityStatement) T_RestResourceInteractionCode(numRest int, numResource int, numInteraction int) templ.Component {
	optionsValueSet := VSType_restful_interaction

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Interaction) >= numInteraction {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Code", &resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Code, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceInteractionDocumentation(numRest int, numResource int, numInteraction int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Interaction) >= numInteraction {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Interaction["+strconv.Itoa(numInteraction)+"].Documentation", resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Documentation)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamId(numRest int, numResource int, numSearchParam int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Id", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Id)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamName(numRest int, numResource int, numSearchParam int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Name", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Name", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Name)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDefinition(numRest int, numResource int, numSearchParam int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Definition", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Definition", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Definition)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamType(numRest int, numResource int, numSearchParam int) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Type", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Type, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDocumentation(numRest int, numResource int, numSearchParam int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].SearchParam) >= numSearchParam {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].SearchParam["+strconv.Itoa(numSearchParam)+"].Documentation", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Documentation)
}
func (resource *CapabilityStatement) T_RestResourceOperationId(numRest int, numResource int, numOperation int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Operation) >= numOperation {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Id", resource.Rest[numRest].Resource[numResource].Operation[numOperation].Id)
}
func (resource *CapabilityStatement) T_RestResourceOperationName(numRest int, numResource int, numOperation int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Operation) >= numOperation {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Name", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Name", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Name)
}
func (resource *CapabilityStatement) T_RestResourceOperationDefinition(numRest int, numResource int, numOperation int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Operation) >= numOperation {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Definition", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Definition", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Definition)
}
func (resource *CapabilityStatement) T_RestResourceOperationDocumentation(numRest int, numResource int, numOperation int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Resource) >= numResource || len(resource.Rest[numRest].Resource[numResource].Operation) >= numOperation {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Resource["+strconv.Itoa(numResource)+"].Operation["+strconv.Itoa(numOperation)+"].Documentation", resource.Rest[numRest].Resource[numResource].Operation[numOperation].Documentation)
}
func (resource *CapabilityStatement) T_RestInteractionId(numRest int, numInteraction int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Interaction) >= numInteraction {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Id", resource.Rest[numRest].Interaction[numInteraction].Id)
}
func (resource *CapabilityStatement) T_RestInteractionCode(numRest int, numInteraction int) templ.Component {
	optionsValueSet := VSSystem_restful_interaction

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Interaction) >= numInteraction {
		return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Code", &resource.Rest[numRest].Interaction[numInteraction].Code, optionsValueSet)
}
func (resource *CapabilityStatement) T_RestInteractionDocumentation(numRest int, numInteraction int) templ.Component {

	if resource == nil || len(resource.Rest) >= numRest || len(resource.Rest[numRest].Interaction) >= numInteraction {
		return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Rest["+strconv.Itoa(numRest)+"].Interaction["+strconv.Itoa(numInteraction)+"].Documentation", resource.Rest[numRest].Interaction[numInteraction].Documentation)
}
func (resource *CapabilityStatement) T_MessagingId(numMessaging int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Id", resource.Messaging[numMessaging].Id)
}
func (resource *CapabilityStatement) T_MessagingReliableCache(numMessaging int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging {
		return IntInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].ReliableCache", nil)
	}
	return IntInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].ReliableCache", resource.Messaging[numMessaging].ReliableCache)
}
func (resource *CapabilityStatement) T_MessagingDocumentation(numMessaging int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Documentation", resource.Messaging[numMessaging].Documentation)
}
func (resource *CapabilityStatement) T_MessagingEndpointId(numMessaging int, numEndpoint int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].Endpoint) >= numEndpoint {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Id", resource.Messaging[numMessaging].Endpoint[numEndpoint].Id)
}
func (resource *CapabilityStatement) T_MessagingEndpointProtocol(numMessaging int, numEndpoint int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].Endpoint) >= numEndpoint {
		return CodingSelect("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Protocol", nil, optionsValueSet)
	}
	return CodingSelect("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Protocol", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Protocol, optionsValueSet)
}
func (resource *CapabilityStatement) T_MessagingEndpointAddress(numMessaging int, numEndpoint int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].Endpoint) >= numEndpoint {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Address", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].Endpoint["+strconv.Itoa(numEndpoint)+"].Address", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Address)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageId(numMessaging int, numSupportedMessage int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].SupportedMessage) >= numSupportedMessage {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Id", resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Id)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageMode(numMessaging int, numSupportedMessage int) templ.Component {
	optionsValueSet := VSEvent_capability_mode

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].SupportedMessage) >= numSupportedMessage {
		return CodeSelect("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Mode", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Mode, optionsValueSet)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageDefinition(numMessaging int, numSupportedMessage int) templ.Component {

	if resource == nil || len(resource.Messaging) >= numMessaging || len(resource.Messaging[numMessaging].SupportedMessage) >= numSupportedMessage {
		return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Definition", nil)
	}
	return StringInput("CapabilityStatement.Messaging["+strconv.Itoa(numMessaging)+"].SupportedMessage["+strconv.Itoa(numSupportedMessage)+"].Definition", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Definition)
}
func (resource *CapabilityStatement) T_DocumentId(numDocument int) templ.Component {

	if resource == nil || len(resource.Document) >= numDocument {
		return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Id", nil)
	}
	return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Id", resource.Document[numDocument].Id)
}
func (resource *CapabilityStatement) T_DocumentMode(numDocument int) templ.Component {
	optionsValueSet := VSDocument_mode

	if resource == nil || len(resource.Document) >= numDocument {
		return CodeSelect("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Mode", &resource.Document[numDocument].Mode, optionsValueSet)
}
func (resource *CapabilityStatement) T_DocumentDocumentation(numDocument int) templ.Component {

	if resource == nil || len(resource.Document) >= numDocument {
		return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Documentation", nil)
	}
	return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Documentation", resource.Document[numDocument].Documentation)
}
func (resource *CapabilityStatement) T_DocumentProfile(numDocument int) templ.Component {

	if resource == nil || len(resource.Document) >= numDocument {
		return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Profile", nil)
	}
	return StringInput("CapabilityStatement.Document["+strconv.Itoa(numDocument)+"].Profile", &resource.Document[numDocument].Profile)
}
