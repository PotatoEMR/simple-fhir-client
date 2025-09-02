package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *ConceptMap) ConceptMapLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ConceptMap) ConceptMapGroupElementCode(numGroup int, numElement int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Element) >= numElement {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapGroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapGroupElementTargetEquivalence(numGroup int, numElement int, numTarget int) templ.Component {
	optionsValueSet := VSConcept_map_equivalence

	if resource != nil && len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return CodeSelect("equivalence", nil, optionsValueSet)
	}
	return CodeSelect("equivalence", &resource.Group[numGroup].Element[numElement].Target[numTarget].Equivalence, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapGroupUnmappedMode(numGroup int) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet)
}
func (resource *ConceptMap) ConceptMapGroupUnmappedCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Group) >= numGroup {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", resource.Group[numGroup].Unmapped.Code, optionsValueSet)
}
