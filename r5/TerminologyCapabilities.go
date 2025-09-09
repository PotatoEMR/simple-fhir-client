package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                   time.Time                              `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *TerminologyCapabilities) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Url", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Url", resource.Url, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Version", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Version", resource.Version, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("TerminologyCapabilities.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("TerminologyCapabilities.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Name", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Name", resource.Name, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Title", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Title", resource.Title, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Experimental", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("TerminologyCapabilities.Date", nil, htmlAttrs)
	}
	return DateTimeInput("TerminologyCapabilities.Date", &resource.Date, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Publisher", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Description", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Description", resource.Description, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("TerminologyCapabilities.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("TerminologyCapabilities.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Purpose", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Copyright", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_LockedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.LockedDate", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.LockedDate", resource.LockedDate, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSearch(htmlAttrs string) templ.Component {
	optionsValueSet := VSCode_search_support

	if resource == nil {
		return CodeSelect("TerminologyCapabilities.CodeSearch", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.CodeSearch", resource.CodeSearch, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_SoftwareName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Software.Name", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Software.Name", &resource.Software.Name, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_SoftwareVersion(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Software.Version", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Software.Version", resource.Software.Version, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ImplementationDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Implementation.Description", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Implementation.Description", &resource.Implementation.Description, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ImplementationUrl(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Implementation.Url", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Implementation.Url", resource.Implementation.Url, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemUri(numCodeSystem int, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Uri", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Uri", resource.CodeSystem[numCodeSystem].Uri, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemContent(numCodeSystem int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Content", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Content", &resource.CodeSystem[numCodeSystem].Content, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemSubsumption(numCodeSystem int, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Subsumption", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Subsumption", resource.CodeSystem[numCodeSystem].Subsumption, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCode(numCodeSystem int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Code", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Code", resource.CodeSystem[numCodeSystem].Version[numVersion].Code, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionIsDefault(numCodeSystem int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].IsDefault", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].IsDefault", resource.CodeSystem[numCodeSystem].Version[numVersion].IsDefault, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionCompositional(numCodeSystem int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) {
		return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Compositional", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Compositional", resource.CodeSystem[numCodeSystem].Version[numVersion].Compositional, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionProperty(numCodeSystem int, numVersion int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numProperty >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Property) {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Property["+strconv.Itoa(numProperty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Property["+strconv.Itoa(numProperty)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Property[numProperty], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterCode(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numFilter >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Code", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Code, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_CodeSystemVersionFilterOp(numCodeSystem int, numVersion int, numFilter int, numOp int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCodeSystem >= len(resource.CodeSystem) || numVersion >= len(resource.CodeSystem[numCodeSystem].Version) || numFilter >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) || numOp >= len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op) {
		return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Op["+strconv.Itoa(numOp)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.CodeSystem["+strconv.Itoa(numCodeSystem)+"].Version["+strconv.Itoa(numVersion)+"].Filter["+strconv.Itoa(numFilter)+"].Op["+strconv.Itoa(numOp)+"]", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op[numOp], optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionHierarchical(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Hierarchical", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Hierarchical", resource.Expansion.Hierarchical, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionPaging(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Paging", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Paging", resource.Expansion.Paging, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionIncomplete(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Expansion.Incomplete", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Expansion.Incomplete", resource.Expansion.Incomplete, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionTextFilter(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TerminologyCapabilities.Expansion.TextFilter", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Expansion.TextFilter", resource.Expansion.TextFilter, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterName(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return CodeSelect("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Expansion.Parameter[numParameter].Name, optionsValueSet, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ExpansionParameterDocumentation(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Documentation", nil, htmlAttrs)
	}
	return StringInput("TerminologyCapabilities.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Documentation", resource.Expansion.Parameter[numParameter].Documentation, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ValidateCodeTranslations(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.ValidateCode.Translations", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.ValidateCode.Translations", &resource.ValidateCode.Translations, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_TranslationNeedsMap(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Translation.NeedsMap", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Translation.NeedsMap", &resource.Translation.NeedsMap, htmlAttrs)
}
func (resource *TerminologyCapabilities) T_ClosureTranslation(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("TerminologyCapabilities.Closure.Translation", nil, htmlAttrs)
	}
	return BoolInput("TerminologyCapabilities.Closure.Translation", resource.Closure.Translation, htmlAttrs)
}
