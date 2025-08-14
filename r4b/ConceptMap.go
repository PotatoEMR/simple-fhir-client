//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMap struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Url               *string           `json:"url,omitempty"`
	Identifier        *Identifier       `json:"identifier,omitempty"`
	Version           *string           `json:"version,omitempty"`
	Name              *string           `json:"name,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Status            string            `json:"status"`
	Experimental      *bool             `json:"experimental,omitempty"`
	Date              *string           `json:"date,omitempty"`
	Publisher         *string           `json:"publisher,omitempty"`
	Contact           []ContactDetail   `json:"contact,omitempty"`
	Description       *string           `json:"description,omitempty"`
	UseContext        []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose           *string           `json:"purpose,omitempty"`
	Copyright         *string           `json:"copyright,omitempty"`
	SourceUri         *string           `json:"sourceUri"`
	SourceCanonical   *string           `json:"sourceCanonical"`
	TargetUri         *string           `json:"targetUri"`
	TargetCanonical   *string           `json:"targetCanonical"`
	Group             []ConceptMapGroup `json:"group,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMapGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Source            *string                  `json:"source,omitempty"`
	SourceVersion     *string                  `json:"sourceVersion,omitempty"`
	Target            *string                  `json:"target,omitempty"`
	TargetVersion     *string                  `json:"targetVersion,omitempty"`
	Element           []ConceptMapGroupElement `json:"element"`
	Unmapped          *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMapGroupElement struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *string                        `json:"code,omitempty"`
	Display           *string                        `json:"display,omitempty"`
	Target            []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMapGroupElementTarget struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Code              *string                                 `json:"code,omitempty"`
	Display           *string                                 `json:"display,omitempty"`
	Equivalence       string                                  `json:"equivalence"`
	Comment           *string                                 `json:"comment,omitempty"`
	DependsOn         []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	System            *string     `json:"system,omitempty"`
	Value             string      `json:"value"`
	Display           *string     `json:"display,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ConceptMap
type ConceptMapGroupUnmapped struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Code              *string     `json:"code,omitempty"`
	Display           *string     `json:"display,omitempty"`
	Url               *string     `json:"url,omitempty"`
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
