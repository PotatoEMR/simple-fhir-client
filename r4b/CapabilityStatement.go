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

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
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
	Date                time.Time                          `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementSoftware struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
	ReleaseDate       *time.Time  `json:"releaseDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
	Custodian         *Reference  `json:"custodian,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
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

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementRestSecurity struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Cors              *bool             `json:"cors,omitempty"`
	Service           []CodeableConcept `json:"service,omitempty"`
	Description       *string           `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
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

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        *string     `json:"definition,omitempty"`
	Type              string      `json:"type"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        string      `json:"definition"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementRestInteraction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementMessaging struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache     *int                                           `json:"reliableCache,omitempty"`
	Documentation     *string                                        `json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Protocol          Coding      `json:"protocol"`
	Address           string      `json:"address"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Definition        string      `json:"definition"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CapabilityStatement
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

	rtype := "CapabilityStatement"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CapabilityStatement) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Url", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Url", resource.Url, htmlAttrs)
}
func (resource *CapabilityStatement) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Version", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Version", resource.Version, htmlAttrs)
}
func (resource *CapabilityStatement) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Name", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Name", resource.Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Title", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Title", resource.Title, htmlAttrs)
}
func (resource *CapabilityStatement) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CapabilityStatement.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("CapabilityStatement.Experimental", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *CapabilityStatement) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CapabilityStatement.Date", nil, htmlAttrs)
	}
	return DateTimeInput("CapabilityStatement.Date", &resource.Date, htmlAttrs)
}
func (resource *CapabilityStatement) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Publisher", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *CapabilityStatement) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Description", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Description", resource.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("CapabilityStatement.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CapabilityStatement.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Purpose", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *CapabilityStatement) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Copyright", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *CapabilityStatement) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("CapabilityStatement.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Instantiates(numInstantiates int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("CapabilityStatement.Instantiates."+strconv.Itoa(numInstantiates)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Instantiates."+strconv.Itoa(numInstantiates)+".", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *CapabilityStatement) T_Imports(numImports int, htmlAttrs string) templ.Component {

	if resource == nil || numImports >= len(resource.Imports) {
		return StringInput("CapabilityStatement.Imports."+strconv.Itoa(numImports)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Imports."+strconv.Itoa(numImports)+".", &resource.Imports[numImports], htmlAttrs)
}
func (resource *CapabilityStatement) T_FhirVersion(htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("CapabilityStatement.FhirVersion", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.FhirVersion", &resource.FhirVersion, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_Format(numFormat int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numFormat >= len(resource.Format) {
		return CodeSelect("CapabilityStatement.Format."+strconv.Itoa(numFormat)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Format."+strconv.Itoa(numFormat)+".", &resource.Format[numFormat], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_PatchFormat(numPatchFormat int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPatchFormat >= len(resource.PatchFormat) {
		return CodeSelect("CapabilityStatement.PatchFormat."+strconv.Itoa(numPatchFormat)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.PatchFormat."+strconv.Itoa(numPatchFormat)+".", &resource.PatchFormat[numPatchFormat], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationGuide(numImplementationGuide int, htmlAttrs string) templ.Component {

	if resource == nil || numImplementationGuide >= len(resource.ImplementationGuide) {
		return StringInput("CapabilityStatement.ImplementationGuide."+strconv.Itoa(numImplementationGuide)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.ImplementationGuide."+strconv.Itoa(numImplementationGuide)+".", &resource.ImplementationGuide[numImplementationGuide], htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareName(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.Name", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Software.Name", &resource.Software.Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareVersion(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Software.Version", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Software.Version", resource.Software.Version, htmlAttrs)
}
func (resource *CapabilityStatement) T_SoftwareReleaseDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CapabilityStatement.Software.ReleaseDate", nil, htmlAttrs)
	}
	return DateTimeInput("CapabilityStatement.Software.ReleaseDate", resource.Software.ReleaseDate, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationDescription(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Implementation.Description", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Implementation.Description", &resource.Implementation.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_ImplementationUrl(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CapabilityStatement.Implementation.Url", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Implementation.Url", resource.Implementation.Url, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestMode(numRest int, htmlAttrs string) templ.Component {
	optionsValueSet := VSRestful_capability_mode

	if resource == nil || numRest >= len(resource.Rest) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Mode", &resource.Rest[numRest].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestDocumentation(numRest int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Documentation", resource.Rest[numRest].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestCompartment(numRest int, numCompartment int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numCompartment >= len(resource.Rest[numRest].Compartment) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Compartment."+strconv.Itoa(numCompartment)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Compartment."+strconv.Itoa(numCompartment)+".", &resource.Rest[numRest].Compartment[numCompartment], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityCors(numRest int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) {
		return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Cors", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Cors", resource.Rest[numRest].Security.Cors, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityService(numRest int, numService int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numService >= len(resource.Rest[numRest].Security.Service) {
		return CodeableConceptSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Service."+strconv.Itoa(numService)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Service."+strconv.Itoa(numService)+".", &resource.Rest[numRest].Security.Service[numService], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestSecurityDescription(numRest int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Description", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Security.Description", resource.Rest[numRest].Security.Description, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceType(numRest int, numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Type", &resource.Rest[numRest].Resource[numResource].Type, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceProfile(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Profile", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Profile", resource.Rest[numRest].Resource[numResource].Profile, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSupportedProfile(numRest int, numResource int, numSupportedProfile int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSupportedProfile >= len(resource.Rest[numRest].Resource[numResource].SupportedProfile) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SupportedProfile."+strconv.Itoa(numSupportedProfile)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SupportedProfile."+strconv.Itoa(numSupportedProfile)+".", &resource.Rest[numRest].Resource[numResource].SupportedProfile[numSupportedProfile], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceDocumentation(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Documentation", resource.Rest[numRest].Resource[numResource].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceVersioning(numRest int, numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVersioning_policy

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Versioning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Versioning", resource.Rest[numRest].Resource[numResource].Versioning, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceReadHistory(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ReadHistory", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ReadHistory", resource.Rest[numRest].Resource[numResource].ReadHistory, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceUpdateCreate(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..UpdateCreate", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..UpdateCreate", resource.Rest[numRest].Resource[numResource].UpdateCreate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalCreate(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalCreate", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalCreate", resource.Rest[numRest].Resource[numResource].ConditionalCreate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalRead(numRest int, numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConditional_read_status

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalRead", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalRead", resource.Rest[numRest].Resource[numResource].ConditionalRead, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalUpdate(numRest int, numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalUpdate", nil, htmlAttrs)
	}
	return BoolInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalUpdate", resource.Rest[numRest].Resource[numResource].ConditionalUpdate, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceConditionalDelete(numRest int, numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConditional_delete_status

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalDelete", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ConditionalDelete", resource.Rest[numRest].Resource[numResource].ConditionalDelete, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceReferencePolicy(numRest int, numResource int, numReferencePolicy int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReference_handling_policy

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numReferencePolicy >= len(resource.Rest[numRest].Resource[numResource].ReferencePolicy) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ReferencePolicy."+strconv.Itoa(numReferencePolicy)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..ReferencePolicy."+strconv.Itoa(numReferencePolicy)+".", &resource.Rest[numRest].Resource[numResource].ReferencePolicy[numReferencePolicy], optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchInclude(numRest int, numResource int, numSearchInclude int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchInclude >= len(resource.Rest[numRest].Resource[numResource].SearchInclude) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchInclude."+strconv.Itoa(numSearchInclude)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchInclude."+strconv.Itoa(numSearchInclude)+".", &resource.Rest[numRest].Resource[numResource].SearchInclude[numSearchInclude], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchRevInclude(numRest int, numResource int, numSearchRevInclude int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchRevInclude >= len(resource.Rest[numRest].Resource[numResource].SearchRevInclude) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchRevInclude."+strconv.Itoa(numSearchRevInclude)+".", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchRevInclude."+strconv.Itoa(numSearchRevInclude)+".", &resource.Rest[numRest].Resource[numResource].SearchRevInclude[numSearchRevInclude], htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceInteractionCode(numRest int, numResource int, numInteraction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSType_restful_interaction

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numInteraction >= len(resource.Rest[numRest].Resource[numResource].Interaction) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Interaction."+strconv.Itoa(numInteraction)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Interaction."+strconv.Itoa(numInteraction)+"..Code", &resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Code, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceInteractionDocumentation(numRest int, numResource int, numInteraction int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numInteraction >= len(resource.Rest[numRest].Resource[numResource].Interaction) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Interaction."+strconv.Itoa(numInteraction)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Interaction."+strconv.Itoa(numInteraction)+"..Documentation", resource.Rest[numRest].Resource[numResource].Interaction[numInteraction].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamName(numRest int, numResource int, numSearchParam int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Name", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Name", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDefinition(numRest int, numResource int, numSearchParam int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Definition", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamType(numRest int, numResource int, numSearchParam int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Type", &resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Type, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceSearchParamDocumentation(numRest int, numResource int, numSearchParam int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numSearchParam >= len(resource.Rest[numRest].Resource[numResource].SearchParam) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..SearchParam."+strconv.Itoa(numSearchParam)+"..Documentation", resource.Rest[numRest].Resource[numResource].SearchParam[numSearchParam].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationName(numRest int, numResource int, numOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Name", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Name", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Name, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationDefinition(numRest int, numResource int, numOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Definition", &resource.Rest[numRest].Resource[numResource].Operation[numOperation].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestResourceOperationDocumentation(numRest int, numResource int, numOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numResource >= len(resource.Rest[numRest].Resource) || numOperation >= len(resource.Rest[numRest].Resource[numResource].Operation) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Resource."+strconv.Itoa(numResource)+"..Operation."+strconv.Itoa(numOperation)+"..Documentation", resource.Rest[numRest].Resource[numResource].Operation[numOperation].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestInteractionCode(numRest int, numInteraction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSystem_restful_interaction

	if resource == nil || numRest >= len(resource.Rest) || numInteraction >= len(resource.Rest[numRest].Interaction) {
		return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Interaction."+strconv.Itoa(numInteraction)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Interaction."+strconv.Itoa(numInteraction)+"..Code", &resource.Rest[numRest].Interaction[numInteraction].Code, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_RestInteractionDocumentation(numRest int, numInteraction int, htmlAttrs string) templ.Component {

	if resource == nil || numRest >= len(resource.Rest) || numInteraction >= len(resource.Rest[numRest].Interaction) {
		return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Interaction."+strconv.Itoa(numInteraction)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Rest."+strconv.Itoa(numRest)+"..Interaction."+strconv.Itoa(numInteraction)+"..Documentation", resource.Rest[numRest].Interaction[numInteraction].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingReliableCache(numMessaging int, htmlAttrs string) templ.Component {

	if resource == nil || numMessaging >= len(resource.Messaging) {
		return IntInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..ReliableCache", nil, htmlAttrs)
	}
	return IntInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..ReliableCache", resource.Messaging[numMessaging].ReliableCache, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingDocumentation(numMessaging int, htmlAttrs string) templ.Component {

	if resource == nil || numMessaging >= len(resource.Messaging) {
		return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Documentation", resource.Messaging[numMessaging].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingEndpointProtocol(numMessaging int, numEndpoint int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMessaging >= len(resource.Messaging) || numEndpoint >= len(resource.Messaging[numMessaging].Endpoint) {
		return CodingSelect("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Endpoint."+strconv.Itoa(numEndpoint)+"..Protocol", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Endpoint."+strconv.Itoa(numEndpoint)+"..Protocol", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Protocol, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingEndpointAddress(numMessaging int, numEndpoint int, htmlAttrs string) templ.Component {

	if resource == nil || numMessaging >= len(resource.Messaging) || numEndpoint >= len(resource.Messaging[numMessaging].Endpoint) {
		return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Endpoint."+strconv.Itoa(numEndpoint)+"..Address", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..Endpoint."+strconv.Itoa(numEndpoint)+"..Address", &resource.Messaging[numMessaging].Endpoint[numEndpoint].Address, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageMode(numMessaging int, numSupportedMessage int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_capability_mode

	if resource == nil || numMessaging >= len(resource.Messaging) || numSupportedMessage >= len(resource.Messaging[numMessaging].SupportedMessage) {
		return CodeSelect("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..SupportedMessage."+strconv.Itoa(numSupportedMessage)+"..Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..SupportedMessage."+strconv.Itoa(numSupportedMessage)+"..Mode", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_MessagingSupportedMessageDefinition(numMessaging int, numSupportedMessage int, htmlAttrs string) templ.Component {

	if resource == nil || numMessaging >= len(resource.Messaging) || numSupportedMessage >= len(resource.Messaging[numMessaging].SupportedMessage) {
		return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..SupportedMessage."+strconv.Itoa(numSupportedMessage)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Messaging."+strconv.Itoa(numMessaging)+"..SupportedMessage."+strconv.Itoa(numSupportedMessage)+"..Definition", &resource.Messaging[numMessaging].SupportedMessage[numSupportedMessage].Definition, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentMode(numDocument int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_mode

	if resource == nil || numDocument >= len(resource.Document) {
		return CodeSelect("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Mode", &resource.Document[numDocument].Mode, optionsValueSet, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentDocumentation(numDocument int, htmlAttrs string) templ.Component {

	if resource == nil || numDocument >= len(resource.Document) {
		return StringInput("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Documentation", resource.Document[numDocument].Documentation, htmlAttrs)
}
func (resource *CapabilityStatement) T_DocumentProfile(numDocument int, htmlAttrs string) templ.Component {

	if resource == nil || numDocument >= len(resource.Document) {
		return StringInput("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Profile", nil, htmlAttrs)
	}
	return StringInput("CapabilityStatement.Document."+strconv.Itoa(numDocument)+"..Profile", &resource.Document[numDocument].Profile, htmlAttrs)
}
