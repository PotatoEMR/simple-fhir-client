//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSet struct {
	Id                     *string            `json:"id,omitempty"`
	Meta                   *Meta              `json:"meta,omitempty"`
	ImplicitRules          *string            `json:"implicitRules,omitempty"`
	Language               *string            `json:"language,omitempty"`
	Text                   *Narrative         `json:"text,omitempty"`
	Contained              []Resource         `json:"contained,omitempty"`
	Extension              []Extension        `json:"extension,omitempty"`
	ModifierExtension      []Extension        `json:"modifierExtension,omitempty"`
	Url                    *string            `json:"url,omitempty"`
	Identifier             []Identifier       `json:"identifier,omitempty"`
	Version                *string            `json:"version,omitempty"`
	VersionAlgorithmString *string            `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding            `json:"versionAlgorithmCoding"`
	Name                   *string            `json:"name,omitempty"`
	Title                  *string            `json:"title,omitempty"`
	Status                 string             `json:"status"`
	Experimental           *bool              `json:"experimental,omitempty"`
	Date                   *string            `json:"date,omitempty"`
	Publisher              *string            `json:"publisher,omitempty"`
	Contact                []ContactDetail    `json:"contact,omitempty"`
	Description            *string            `json:"description,omitempty"`
	UseContext             []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept  `json:"jurisdiction,omitempty"`
	Immutable              *bool              `json:"immutable,omitempty"`
	Purpose                *string            `json:"purpose,omitempty"`
	Copyright              *string            `json:"copyright,omitempty"`
	CopyrightLabel         *string            `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string            `json:"approvalDate,omitempty"`
	LastReviewDate         *string            `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period            `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept  `json:"topic,omitempty"`
	Author                 []ContactDetail    `json:"author,omitempty"`
	Editor                 []ContactDetail    `json:"editor,omitempty"`
	Reviewer               []ContactDetail    `json:"reviewer,omitempty"`
	Endorser               []ContactDetail    `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact  `json:"relatedArtifact,omitempty"`
	Compose                *ValueSetCompose   `json:"compose,omitempty"`
	Expansion              *ValueSetExpansion `json:"expansion,omitempty"`
	Scope                  *ValueSetScope     `json:"scope,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
	Property          []string                 `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetComposeInclude struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	System            *string                         `json:"system,omitempty"`
	Version           *string                         `json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet          []string                        `json:"valueSet,omitempty"`
	Copyright         *string                         `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	AdditionalUse     []Coding    `json:"additionalUse,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetComposeIncludeFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	Op                string      `json:"op"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetExpansion struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Next              *string                      `json:"next,omitempty"`
	Timestamp         string                       `json:"timestamp"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Property          []ValueSetExpansionProperty  `json:"property,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
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

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetExpansionProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetExpansionContains struct {
	Id                *string                             `json:"id,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	System            *string                             `json:"system,omitempty"`
	Abstract          *bool                               `json:"abstract,omitempty"`
	Inactive          *bool                               `json:"inactive,omitempty"`
	Version           *string                             `json:"version,omitempty"`
	Code              *string                             `json:"code,omitempty"`
	Display           *string                             `json:"display,omitempty"`
	Property          []ValueSetExpansionContainsProperty `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetExpansionContainsProperty struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Code              string                                         `json:"code"`
	ValueCode         string                                         `json:"valueCode"`
	ValueCoding       Coding                                         `json:"valueCoding"`
	ValueString       string                                         `json:"valueString"`
	ValueInteger      int                                            `json:"valueInteger"`
	ValueBoolean      bool                                           `json:"valueBoolean"`
	ValueDateTime     string                                         `json:"valueDateTime"`
	ValueDecimal      float64                                        `json:"valueDecimal"`
	SubProperty       []ValueSetExpansionContainsPropertySubProperty `json:"subProperty,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetExpansionContainsPropertySubProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCode         string      `json:"valueCode"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueDecimal      float64     `json:"valueDecimal"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ValueSet
type ValueSetScope struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	InclusionCriteria *string     `json:"inclusionCriteria,omitempty"`
	ExclusionCriteria *string     `json:"exclusionCriteria,omitempty"`
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
