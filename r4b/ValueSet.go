package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSet struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Url               *string            `json:"url,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Version           *string            `json:"version,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Status            string             `json:"status"`
	Experimental      *bool              `json:"experimental,omitempty"`
	Date              *string            `json:"date,omitempty"`
	Publisher         *string            `json:"publisher,omitempty"`
	Contact           []ContactDetail    `json:"contact,omitempty"`
	Description       *string            `json:"description,omitempty"`
	UseContext        []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept  `json:"jurisdiction,omitempty"`
	Immutable         *bool              `json:"immutable,omitempty"`
	Purpose           *string            `json:"purpose,omitempty"`
	Copyright         *string            `json:"copyright,omitempty"`
	Compose           *ValueSetCompose   `json:"compose,omitempty"`
	Expansion         *ValueSetExpansion `json:"expansion,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeInclude struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	System            *string                         `json:"system,omitempty"`
	Version           *string                         `json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet          []string                        `json:"valueSet,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	Op                string      `json:"op"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetExpansion struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Timestamp         string                       `json:"timestamp"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetExpansionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	ValueString       *string     `json:"valueString"`
	ValueBoolean      *bool       `json:"valueBoolean"`
	ValueInteger      *int        `json:"valueInteger"`
	ValueDecimal      *float64    `json:"valueDecimal"`
	ValueUri          *string     `json:"valueUri"`
	ValueCode         *string     `json:"valueCode"`
	ValueDateTime     *string     `json:"valueDateTime"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetExpansionContains struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	System            *string     `json:"system,omitempty"`
	Abstract          *bool       `json:"abstract,omitempty"`
	Inactive          *bool       `json:"inactive,omitempty"`
	Version           *string     `json:"version,omitempty"`
	Code              *string     `json:"code,omitempty"`
	Display           *string     `json:"display,omitempty"`
}

type OtherValueSet ValueSet

// on convert struct to json, automatically add resourceType=ValueSet
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}
func (resource *ValueSet) ValueSetLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetComposeIncludeConceptCode(numInclude int, numConcept int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Compose.Include[numInclude].Concept) >= numConcept {
		currentVal = resource.Compose.Include[numInclude].Concept[numConcept].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetComposeIncludeConceptDesignationLanguage(numInclude int, numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) >= numDesignation {
		currentVal = *resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetComposeIncludeFilterProperty(numInclude int, numFilter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		currentVal = resource.Compose.Include[numInclude].Filter[numFilter].Property
	}
	return CodeSelect("property", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetComposeIncludeFilterOp(numInclude int, numFilter int) templ.Component {
	optionsValueSet := VSFilter_operator
	currentVal := ""
	if resource != nil && len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		currentVal = resource.Compose.Include[numInclude].Filter[numFilter].Op
	}
	return CodeSelect("op", currentVal, optionsValueSet)
}
func (resource *ValueSet) ValueSetExpansionContainsCode(numContains int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Expansion.Contains) >= numContains {
		currentVal = *resource.Expansion.Contains[numContains].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
