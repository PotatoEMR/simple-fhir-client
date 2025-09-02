package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilities struct {
	Id                *string                                `json:"id,omitempty"`
	Meta              *Meta                                  `json:"meta,omitempty"`
	ImplicitRules     *string                                `json:"implicitRules,omitempty"`
	Language          *string                                `json:"language,omitempty"`
	Text              *Narrative                             `json:"text,omitempty"`
	Contained         []Resource                             `json:"contained,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Url               *string                                `json:"url,omitempty"`
	Version           *string                                `json:"version,omitempty"`
	Name              *string                                `json:"name,omitempty"`
	Title             *string                                `json:"title,omitempty"`
	Status            string                                 `json:"status"`
	Experimental      *bool                                  `json:"experimental,omitempty"`
	Date              string                                 `json:"date"`
	Publisher         *string                                `json:"publisher,omitempty"`
	Contact           []ContactDetail                        `json:"contact,omitempty"`
	Description       *string                                `json:"description,omitempty"`
	UseContext        []UsageContext                         `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                      `json:"jurisdiction,omitempty"`
	Purpose           *string                                `json:"purpose,omitempty"`
	Copyright         *string                                `json:"copyright,omitempty"`
	Kind              string                                 `json:"kind"`
	Software          *TerminologyCapabilitiesSoftware       `json:"software,omitempty"`
	Implementation    *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`
	LockedDate        *bool                                  `json:"lockedDate,omitempty"`
	CodeSystem        []TerminologyCapabilitiesCodeSystem    `json:"codeSystem,omitempty"`
	Expansion         *TerminologyCapabilitiesExpansion      `json:"expansion,omitempty"`
	CodeSearch        *string                                `json:"codeSearch,omitempty"`
	ValidateCode      *TerminologyCapabilitiesValidateCode   `json:"validateCode,omitempty"`
	Translation       *TerminologyCapabilitiesTranslation    `json:"translation,omitempty"`
	Closure           *TerminologyCapabilitiesClosure        `json:"closure,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesSoftware struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesImplementation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystem struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Uri               *string                                    `json:"uri,omitempty"`
	Version           []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`
	Subsumption       *bool                                      `json:"subsumption,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
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

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Op                []string    `json:"op"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
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

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesExpansionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesValidateCode struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Translations      bool        `json:"translations"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesTranslation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NeedsMap          bool        `json:"needsMap"`
}

// http://hl7.org/fhir/r4/StructureDefinition/TerminologyCapabilities
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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesKind() templ.Component {
	optionsValueSet := VSCapability_statement_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSearch() templ.Component {
	optionsValueSet := VSCode_search_support

	if resource == nil {
		return CodeSelect("codeSearch", nil, optionsValueSet)
	}
	return CodeSelect("codeSearch", resource.CodeSearch, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionLanguage(numCodeSystem int, numVersion int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", &resource.CodeSystem[numCodeSystem].Version[numVersion].Language[0], optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionProperty(numCodeSystem int, numVersion int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CodeSystem[numCodeSystem].Version) >= numVersion {
		return CodeSelect("property", nil, optionsValueSet)
	}
	return CodeSelect("property", &resource.CodeSystem[numCodeSystem].Version[numVersion].Property[0], optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionFilterCode(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Code, optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesCodeSystemVersionFilterOp(numCodeSystem int, numVersion int, numFilter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CodeSystem[numCodeSystem].Version[numVersion].Filter) >= numFilter {
		return CodeSelect("op", nil, optionsValueSet)
	}
	return CodeSelect("op", &resource.CodeSystem[numCodeSystem].Version[numVersion].Filter[numFilter].Op[0], optionsValueSet)
}
func (resource *TerminologyCapabilities) TerminologyCapabilitiesExpansionParameterName(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Expansion.Parameter) >= numParameter {
		return CodeSelect("name", nil, optionsValueSet)
	}
	return CodeSelect("name", &resource.Expansion.Parameter[numParameter].Name, optionsValueSet)
}
