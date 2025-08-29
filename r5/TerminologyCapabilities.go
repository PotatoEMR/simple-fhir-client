package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString *string                                `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                                `json:"versionAlgorithmCoding"`
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
func (resource *TerminologyCapabilities) TerminologyCapabilitiesLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesKind() templ.Component {
	optionsValueSet := VSCapability_statement_kind
	currentVal := ""
	if resource != nil {
		currentVal = resource.Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSearch() templ.Component {
	optionsValueSet := VSCode_search_support
	currentVal := ""
	if resource != nil {
		currentVal = *resource.CodeSearch
	}
	return CodeSelect("codeSearch", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemContent(numCodeSystem int) templ.Component {
	optionsValueSet := VSCodesystem_content_mode
	currentVal := ""
	if resource != nil && len(resource.CodeSystem) >= numCodeSystem {
		currentVal = resource.CodeSystem[numCodeSystem].Content
	}
	return CodeSelect("content", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionLanguage(numCodeSystem int, numVersion int) templ.Component {
	optionsValueSet := VSLanguages
	currentVal := ""
	if resource != nil && len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		currentVal = resource.CodeSystem[numCodeSystem].Version[numVersion].Language[0]
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionProperty(numCodeSystem int, numVersion int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		currentVal = resource.CodeSystem[numCodeSystem].Version[numVersion].Property[0]
	}
	return CodeSelect("property", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionFilterCode(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		currentVal = resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionFilterOp(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		currentVal = resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op[0]
	}
	return CodeSelect("op", currentVal, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesExpansionParameterName(numParameter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Expansion.Parameter) >= numParameter {
		currentVal = resource.Expansion.Parameter[numParameter].Name
	}
	return CodeSelect("name", currentVal, optionsValueSet)
}
