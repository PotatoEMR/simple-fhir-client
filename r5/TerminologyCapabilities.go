package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilities struct {
	Id                     *string                                `json:"id,omitempty"`
	Meta                   *Meta                                  `json:"meta,omitempty"`
	ImplicitRules          *string                                `json:"implicitRules,omitempty"`
	Language               *string                                `json:"language,omitempty"`
	Text                   *Narrative                             `json:"text,omitempty"`
	Contained              []Resource                             `json:"contained,omitempty"`
	Extension              []Extension                            `json:"extension,omitempty"`
	ModifierExtension      []Extension                            `json:"modifierExtension,omitempty"`
	Url                    *string                                `json:"url,omitempty"`
	Identifier             []Identifier                           `json:"identifier,omitempty"`
	Version                *string                                `json:"version,omitempty"`
	VersionAlgorithmString *string                                `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                                `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                                `json:"name,omitempty"`
	Title                  *string                                `json:"title,omitempty"`
	Status                 string                                 `json:"status"`
	Experimental           *bool                                  `json:"experimental,omitempty"`
	Date                   string                                 `json:"date"`
	Publisher              *string                                `json:"publisher,omitempty"`
	Contact                []ContactDetail                        `json:"contact,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	UseContext             []UsageContext                         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                      `json:"jurisdiction,omitempty"`
	Purpose                *string                                `json:"purpose,omitempty"`
	Copyright              *string                                `json:"copyright,omitempty"`
	CopyrightLabel         *string                                `json:"copyrightLabel,omitempty"`
	Kind                   string                                 `json:"kind"`
	Software               *TerminologyCapabilitiesSoftware       `json:"software,omitempty"`
	Implementation         *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`
	LockedDate             *bool                                  `json:"lockedDate,omitempty"`
	CodeSystem             []TerminologyCapabilitiesCodeSystem    `json:"codeSystem,omitempty"`
	Expansion              *TerminologyCapabilitiesExpansion      `json:"expansion,omitempty"`
	CodeSearch             *string                                `json:"codeSearch,omitempty"`
	ValidateCode           *TerminologyCapabilitiesValidateCode   `json:"validateCode,omitempty"`
	Translation            *TerminologyCapabilitiesTranslation    `json:"translation,omitempty"`
	Closure                *TerminologyCapabilitiesClosure        `json:"closure,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesSoftware struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystem struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Uri               *string                                    `json:"uri,omitempty"`
	Version           []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`
	Content           string                                     `json:"content"`
	Subsumption       *bool                                      `json:"subsumption,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystemVersion struct {
	Id                *string                                          `json:"id,omitempty"`
	Extension         []Extension                                      `json:"extension,omitempty"`
	ModifierExtension []Extension                                      `json:"modifierExtension,omitempty"`
	Code              *string                                          `json:"code,omitempty"`
	IsDefault         *bool                                            `json:"isDefault,omitempty"`
	Compositional     *bool                                            `json:"compositional,omitempty"`
	Language          []string                                         `json:"language,omitempty"`
	Filter            []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`
	Property          []string                                         `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Op                []string    `json:"op"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesExpansion struct {
	Id                *string                                     `json:"id,omitempty"`
	Extension         []Extension                                 `json:"extension,omitempty"`
	ModifierExtension []Extension                                 `json:"modifierExtension,omitempty"`
	Hierarchical      *bool                                       `json:"hierarchical,omitempty"`
	Paging            *bool                                       `json:"paging,omitempty"`
	Incomplete        *bool                                       `json:"incomplete,omitempty"`
	Parameter         []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`
	TextFilter        *string                                     `json:"textFilter,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesExpansionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesValidateCode struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Translations      bool        `json:"translations"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesTranslation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NeedsMap          bool        `json:"needsMap"`
}

// http://hl7.org/fhir/r5/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesClosure struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Translation       *bool       `json:"translation,omitempty"`
}

type OtherTerminologyCapabilities TerminologyCapabilities

// on convert struct to json, automatically add resourceType=TerminologyCapabilities
func (r TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTerminologyCapabilities
		ResourceType string `json:"resourceType"`
	}{
		OtherTerminologyCapabilities: OtherTerminologyCapabilities(r),
		ResourceType:                 "TerminologyCapabilities",
	})
}

func (resource *TerminologyCapabilities) T_Id() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Id", resource.Id)
}
func (resource *TerminologyCapabilities) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.ImplicitRules", nil)
	}
	return StringInput("TerminologyCapabilities.ImplicitRules", resource.ImplicitRules)
}
func (resource *TerminologyCapabilities) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.Language", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.Language", resource.Language, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_Url() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Url", nil)
	}
	return StringInput("TerminologyCapabilities.Url", resource.Url)
}
func (resource *TerminologyCapabilities) T_Version() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Version", nil)
	}
	return StringInput("TerminologyCapabilities.Version", resource.Version)
}
func (resource *TerminologyCapabilities) T_Name() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Name", nil)
	}
	return StringInput("TerminologyCapabilities.Name", resource.Name)
}
func (resource *TerminologyCapabilities) T_Title() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Title", nil)
	}
	return StringInput("TerminologyCapabilities.Title", resource.Title)
}
func (resource *TerminologyCapabilities) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.Status", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.Status", &resource.Status, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Experimental", nil)
	}
	return BoolInput("TerminologyCapabilities.Experimental", resource.Experimental)
}
func (resource *TerminologyCapabilities) T_Date() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Date", nil)
	}
	return StringInput("TerminologyCapabilities.Date", &resource.Date)
}
func (resource *TerminologyCapabilities) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Publisher", nil)
	}
	return StringInput("TerminologyCapabilities.Publisher", resource.Publisher)
}
func (resource *TerminologyCapabilities) T_Description() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Description", nil)
	}
	return StringInput("TerminologyCapabilities.Description", resource.Description)
}
func (resource *TerminologyCapabilities) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("TerminologyCapabilities.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("TerminologyCapabilities.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *TerminologyCapabilities) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Purpose", nil)
	}
	return StringInput("TerminologyCapabilities.Purpose", resource.Purpose)
}
func (resource *TerminologyCapabilities) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Copyright", nil)
	}
	return StringInput("TerminologyCapabilities.Copyright", resource.Copyright)
}
func (resource *TerminologyCapabilities) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.CopyrightLabel", nil)
	}
	return StringInput("TerminologyCapabilities.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *TerminologyCapabilities) T_Kind() templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.Kind", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.Kind", &resource.Kind, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_LockedDate() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.LockedDate", nil)
	}
	return BoolInput("TerminologyCapabilities.LockedDate", resource.LockedDate)
}
func (resource *TerminologyCapabilities) T_CodeSearch() templ.Component {
	optionsValueSet := VSCode_search_support

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.CodeSearch", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSearch", resource.CodeSearch, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_SoftwareId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Software.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Software.Id", resource.Software.Id)
}
func (resource *TerminologyCapabilities) T_SoftwareName() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Software.Name", nil)
	}
	return StringInput("TerminologyCapabilities.Software.Name", &resource.Software.Name)
}
func (resource *TerminologyCapabilities) T_SoftwareVersion() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Software.Version", nil)
	}
	return StringInput("TerminologyCapabilities.Software.Version", resource.Software.Version)
}
func (resource *TerminologyCapabilities) T_ImplementationId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Implementation.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Implementation.Id", resource.Implementation.Id)
}
func (resource *TerminologyCapabilities) T_ImplementationDescription() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Implementation.Description", nil)
	}
	return StringInput("TerminologyCapabilities.Implementation.Description", &resource.Implementation.Description)
}
func (resource *TerminologyCapabilities) T_ImplementationUrl() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Implementation.Url", nil)
	}
	return StringInput("TerminologyCapabilities.Implementation.Url", resource.Implementation.Url)
}
func (resource *TerminologyCapabilities) T_CodeSystemId(numCodeSystem int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Id", nil)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Id", resource.CodeSystem[numCodeSystem].Id)
}
func (resource *TerminologyCapabilities) T_CodeSystemUri(numCodeSystem int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Uri", nil)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Uri", resource.CodeSystem[numCodeSystem].Uri)
}
func (resource *TerminologyCapabilities) T_CodeSystemContent(numCodeSystem int) templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Content", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Content", &resource.CodeSystem[numCodeSystem].Content, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_CodeSystemSubsumption(numCodeSystem int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Subsumption", nil)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Subsumption", resource.CodeSystem[numCodeSystem].Subsumption)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionId(numCodeSystem int, numVersion int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Id", nil)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Id", resource.CodeSystem[numCodeSystem].Version[numVersion].Id)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCode(numCodeSystem int, numVersion int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Code", nil)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Code", resource.CodeSystem[numCodeSystem].Version[numVersion].Code)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionIsDefault(numCodeSystem int, numVersion int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].IsDefault", nil)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].IsDefault", resource.CodeSystem[numCodeSystem].Version[numVersion].IsDefault)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCompositional(numCodeSystem int, numVersion int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Compositional", nil)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Compositional", resource.CodeSystem[numCodeSystem].Version[numVersion].Compositional)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionLanguage(numCodeSystem int, numVersion int, numLanguage int) templ.Component {
	optionsValueSet := VSLanguages

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Language) >= numLanguage {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Language["+strconv.Itoa(numLanguage)+"]", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Language["+strconv.Itoa(numLanguage)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Language[numLanguage], optionsValueSet)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionProperty(numCodeSystem int, numVersion int, numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Property) >= numProperty {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Property["+strconv.Itoa(numProperty)+"]", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Property["+strconv.Itoa(numProperty)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Property[numProperty], optionsValueSet)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterId(numCodeSystem int, numVersion int, numFilter int) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Id", nil)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Id", resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Id)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterCode(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Code", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Code, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterOp(numCodeSystem int, numVersion int, numFilter int, numOp int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CodeSystem) >= numCodeSystem || len(resource.CodeSystem[numCodeSystem].Version) >= numVersion || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter || len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op) >= numOp {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Op["+strconv.Itoa(numOp)+"]", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Op["+strconv.Itoa(numOp)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op[numOp], optionsValueSet)
}
func (resource *TerminologyCapabilities) T_ExpansionId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Expansion.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Expansion.Id", resource.Expansion.Id)
}
func (resource *TerminologyCapabilities) T_ExpansionHierarchical() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Hierarchical", nil)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Hierarchical", resource.Expansion.Hierarchical)
}
func (resource *TerminologyCapabilities) T_ExpansionPaging() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Paging", nil)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Paging", resource.Expansion.Paging)
}
func (resource *TerminologyCapabilities) T_ExpansionIncomplete() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Incomplete", nil)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Incomplete", resource.Expansion.Incomplete)
}
func (resource *TerminologyCapabilities) T_ExpansionTextFilter() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Expansion.TextFilter", nil)
	}
	return StringInput("TerminologyCapabilities.Expansion.TextFilter", resource.Expansion.TextFilter)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Expansion.Parameter) >= numParameter {
		return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Expansion.Parameter[numParameter].Id)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterName(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Expansion.Parameter) >= numParameter {
		return CodeSelect("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", nil, optionsValueSet)
	}
	return CodeSelect("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Expansion.Parameter[numParameter].Name, optionsValueSet)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterDocumentation(numParameter int) templ.Component {

	if resource == nil || len(resource.Expansion.Parameter) >= numParameter {
		return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Documentation", nil)
	}
	return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Documentation", resource.Expansion.Parameter[numParameter].Documentation)
}
func (resource *TerminologyCapabilities) T_ValidateCodeId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.ValidateCode.Id", nil)
	}
	return StringInput("TerminologyCapabilities.ValidateCode.Id", resource.ValidateCode.Id)
}
func (resource *TerminologyCapabilities) T_ValidateCodeTranslations() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.ValidateCode.Translations", nil)
	}
	return BoolInput("TerminologyCapabilities.ValidateCode.Translations", &resource.ValidateCode.Translations)
}
func (resource *TerminologyCapabilities) T_TranslationId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Translation.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Translation.Id", resource.Translation.Id)
}
func (resource *TerminologyCapabilities) T_TranslationNeedsMap() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Translation.NeedsMap", nil)
	}
	return BoolInput("TerminologyCapabilities.Translation.NeedsMap", &resource.Translation.NeedsMap)
}
func (resource *TerminologyCapabilities) T_ClosureId() templ.Component {

	if resource == nil {
		return StringInput("TerminologyCapabilities.Closure.Id", nil)
	}
	return StringInput("TerminologyCapabilities.Closure.Id", resource.Closure.Id)
}
func (resource *TerminologyCapabilities) T_ClosureTranslation() templ.Component {

	if resource == nil {
		return BoolInput("TerminologyCapabilities.Closure.Translation", nil)
	}
	return BoolInput("TerminologyCapabilities.Closure.Translation", resource.Closure.Translation)
}
