//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMap struct {
	Id                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []Resource                      `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Url                    *string                         `json:"url,omitempty"`
	Identifier             []Identifier                    `json:"identifier,omitempty"`
	Version                *string                         `json:"version,omitempty"`
	VersionAlgorithmString *string                         `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding"`
	Name                   *string                         `json:"name,omitempty"`
	Title                  *string                         `json:"title,omitempty"`
	Status                 string                          `json:"status"`
	Experimental           *bool                           `json:"experimental,omitempty"`
	Date                   *string                         `json:"date,omitempty"`
	Publisher              *string                         `json:"publisher,omitempty"`
	Contact                []ContactDetail                 `json:"contact,omitempty"`
	Description            *string                         `json:"description,omitempty"`
	UseContext             []UsageContext                  `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept               `json:"jurisdiction,omitempty"`
	Purpose                *string                         `json:"purpose,omitempty"`
	Copyright              *string                         `json:"copyright,omitempty"`
	CopyrightLabel         *string                         `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                         `json:"approvalDate,omitempty"`
	LastReviewDate         *string                         `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                         `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept               `json:"topic,omitempty"`
	Author                 []ContactDetail                 `json:"author,omitempty"`
	Editor                 []ContactDetail                 `json:"editor,omitempty"`
	Reviewer               []ContactDetail                 `json:"reviewer,omitempty"`
	Endorser               []ContactDetail                 `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact               `json:"relatedArtifact,omitempty"`
	Property               []ConceptMapProperty            `json:"property,omitempty"`
	AdditionalAttribute    []ConceptMapAdditionalAttribute `json:"additionalAttribute,omitempty"`
	SourceScopeUri         *string                         `json:"sourceScopeUri"`
	SourceScopeCanonical   *string                         `json:"sourceScopeCanonical"`
	TargetScopeUri         *string                         `json:"targetScopeUri"`
	TargetScopeCanonical   *string                         `json:"targetScopeCanonical"`
	Group                  []ConceptMapGroup               `json:"group,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
	System            *string     `json:"system,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapAdditionalAttribute struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Source            *string                  `json:"source,omitempty"`
	Target            *string                  `json:"target,omitempty"`
	Element           []ConceptMapGroupElement `json:"element"`
	Unmapped          *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroupElement struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *string                        `json:"code,omitempty"`
	Display           *string                        `json:"display,omitempty"`
	ValueSet          *string                        `json:"valueSet,omitempty"`
	NoMap             *bool                          `json:"noMap,omitempty"`
	Target            []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroupElementTarget struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Code              *string                                 `json:"code,omitempty"`
	Display           *string                                 `json:"display,omitempty"`
	ValueSet          *string                                 `json:"valueSet,omitempty"`
	Relationship      string                                  `json:"relationship"`
	Comment           *string                                 `json:"comment,omitempty"`
	Property          []ConceptMapGroupElementTargetProperty  `json:"property,omitempty"`
	DependsOn         []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroupElementTargetProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueCode         string      `json:"valueCode"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Attribute         string      `json:"attribute"`
	ValueCode         *string     `json:"valueCode"`
	ValueCoding       *Coding     `json:"valueCoding"`
	ValueString       *string     `json:"valueString"`
	ValueBoolean      *bool       `json:"valueBoolean"`
	ValueQuantity     *Quantity   `json:"valueQuantity"`
	ValueSet          *string     `json:"valueSet,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConceptMap
type ConceptMapGroupUnmapped struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Code              *string     `json:"code,omitempty"`
	Display           *string     `json:"display,omitempty"`
	ValueSet          *string     `json:"valueSet,omitempty"`
	Relationship      *string     `json:"relationship,omitempty"`
	OtherMap          *string     `json:"otherMap,omitempty"`
}

type OtherConceptMap ConceptMap

// on convert struct to json, automatically add resourceType=ConceptMap
func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType"`
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
}
