package r5

//generated with command go run ./bultaoreune
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
	Date                   FhirDateTime                           `json:"date"`
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

// struct -> json, automatically add resourceType=Patient
func (r TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTerminologyCapabilities
		ResourceType string `json:"resourceType"`
	}{
		OtherTerminologyCapabilities: OtherTerminologyCapabilities(r),
		ResourceType:                 "TerminologyCapabilities",
	})
}

func (r TerminologyCapabilities) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "TerminologyCapabilities/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "TerminologyCapabilities"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r TerminologyCapabilities) ResourceType() string {
	return "TerminologyCapabilities"
}

func (resource *TerminologyCapabilities) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_LockedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("lockedDate", nil, htmlAttrs)
	}
	return BoolInput("lockedDate", resource.LockedDate, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSearch(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCode_search_support

	if resource == nil {
		return CodeSelect("codeSearch", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("codeSearch", resource.CodeSearch, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_SoftwareName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("software.name", nil, htmlAttrs)
	}
	return StringInput("software.name", &resource.Software.Name, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_SoftwareVersion(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("software.version", nil, htmlAttrs)
	}
	return StringInput("software.version", resource.Software.Version, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ImplementationDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("implementation.description", nil, htmlAttrs)
	}
	return StringInput("implementation.description", &resource.Implementation.Description, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ImplementationUrl(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("implementation.url", nil, htmlAttrs)
	}
	return StringInput("implementation.url", resource.Implementation.Url, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemUri(numCodeSystem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return StringInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].uri", nil, htmlAttrs)
	}
	return StringInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].uri", resource.CodeSystem[numCodeSystem].Uri, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemContent(numCodeSystem int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].content", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].content", &resource.CodeSystem[numCodeSystem].Content, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemSubsumption(numCodeSystem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].subsumption", nil, htmlAttrs)
	}
	return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].subsumption", resource.CodeSystem[numCodeSystem].Subsumption, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCode(numCodeSystem int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return StringInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].code", nil, htmlAttrs)
	}
	return StringInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].code", resource.CodeSystem[numCodeSystem].Version[numVersion].Code, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionIsDefault(numCodeSystem int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].isDefault", nil, htmlAttrs)
	}
	return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].isDefault", resource.CodeSystem[numCodeSystem].Version[numVersion].IsDefault, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCompositional(numCodeSystem int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].compositional", nil, htmlAttrs)
	}
	return BoolInput("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].compositional", resource.CodeSystem[numCodeSystem].Version[numVersion].Compositional, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionProperty(numCodeSystem int, numVersion int, numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numProperty >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Property) {
		return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].property["+strconv.Itoa(numProperty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].property["+strconv.Itoa(numProperty)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Property[numProperty], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterCode(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numFilter >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) {
		return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].filter["+strconv.Itoa(numFilter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].filter["+strconv.Itoa(numFilter)+"].code", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Code, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterOp(numCodeSystem int, numVersion int, numFilter int, numOp int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numFilter >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) || numOp >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op) {
		return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].filter["+strconv.Itoa(numFilter)+"].op["+strconv.Itoa(numOp)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("codeSystem["+strconv.Itoa(numCodeSystem)+"].version["+strconv.Itoa(numVersion)+"].filter["+strconv.Itoa(numFilter)+"].op["+strconv.Itoa(numOp)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op[numOp], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionHierarchical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("expansion.hierarchical", nil, htmlAttrs)
	}
	return BoolInput("expansion.hierarchical", resource.Expansion.Hierarchical, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionPaging(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("expansion.paging", nil, htmlAttrs)
	}
	return BoolInput("expansion.paging", resource.Expansion.Paging, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionIncomplete(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("expansion.incomplete", nil, htmlAttrs)
	}
	return BoolInput("expansion.incomplete", resource.Expansion.Incomplete, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionTextFilter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("expansion.textFilter", nil, htmlAttrs)
	}
	return StringInput("expansion.textFilter", resource.Expansion.TextFilter, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterName(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return CodeSelect("expansion.parameter["+strconv.Itoa(numParameter)+"].name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("expansion.parameter["+strconv.Itoa(numParameter)+"].name", &resource.Expansion.Parameter[numParameter].Name, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterDocumentation(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].documentation", resource.Expansion.Parameter[numParameter].Documentation, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ValidateCodeTranslations(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("validateCode.translations", nil, htmlAttrs)
	}
	return BoolInput("validateCode.translations", &resource.ValidateCode.Translations, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_TranslationNeedsMap(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("translation.needsMap", nil, htmlAttrs)
	}
	return BoolInput("translation.needsMap", &resource.Translation.NeedsMap, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ClosureTranslation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("closure.translation", nil, htmlAttrs)
	}
	return BoolInput("closure.translation", resource.Closure.Translation, htmlAttrs)
}
