//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                     *string              `json:"id,omitempty"`
	Meta                   *Meta                `json:"meta,omitempty"`
	ImplicitRules          *string              `json:"implicitRules,omitempty"`
	Language               *string              `json:"language,omitempty"`
	Text                   *Narrative           `json:"text,omitempty"`
	Contained              []Resource           `json:"contained,omitempty"`
	Extension              []Extension          `json:"extension,omitempty"`
	ModifierExtension      []Extension          `json:"modifierExtension,omitempty"`
	Url                    *string              `json:"url,omitempty"`
	Identifier             []Identifier         `json:"identifier,omitempty"`
	Version                *string              `json:"version,omitempty"`
	VersionAlgorithmString *string              `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding              `json:"versionAlgorithmCoding"`
	Name                   *string              `json:"name,omitempty"`
	Title                  *string              `json:"title,omitempty"`
	Status                 string               `json:"status"`
	Experimental           *bool                `json:"experimental,omitempty"`
	Date                   *string              `json:"date,omitempty"`
	Publisher              *string              `json:"publisher,omitempty"`
	Contact                []ContactDetail      `json:"contact,omitempty"`
	Description            *string              `json:"description,omitempty"`
	UseContext             []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose                *string              `json:"purpose,omitempty"`
	Copyright              *string              `json:"copyright,omitempty"`
	CopyrightLabel         *string              `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string              `json:"approvalDate,omitempty"`
	LastReviewDate         *string              `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period              `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept    `json:"topic,omitempty"`
	Author                 []ContactDetail      `json:"author,omitempty"`
	Editor                 []ContactDetail      `json:"editor,omitempty"`
	Reviewer               []ContactDetail      `json:"reviewer,omitempty"`
	Endorser               []ContactDetail      `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact    `json:"relatedArtifact,omitempty"`
	CaseSensitive          *bool                `json:"caseSensitive,omitempty"`
	ValueSet               *string              `json:"valueSet,omitempty"`
	HierarchyMeaning       *string              `json:"hierarchyMeaning,omitempty"`
	Compositional          *bool                `json:"compositional,omitempty"`
	VersionNeeded          *bool                `json:"versionNeeded,omitempty"`
	Content                string               `json:"content"`
	Supplements            *string              `json:"supplements,omitempty"`
	Count                  *int                 `json:"count,omitempty"`
	Filter                 []CodeSystemFilter   `json:"filter,omitempty"`
	Property               []CodeSystemProperty `json:"property,omitempty"`
	Concept                []CodeSystemConcept  `json:"concept,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Description       *string     `json:"description,omitempty"`
	Operator          []string    `json:"operator"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConcept struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              string                         `json:"code"`
	Display           *string                        `json:"display,omitempty"`
	Definition        *string                        `json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	AdditionalUse     []Coding    `json:"additionalUse,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConceptProperty struct {
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

type OtherCodeSystem CodeSystem

// on convert struct to json, automatically add resourceType=CodeSystem
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}
