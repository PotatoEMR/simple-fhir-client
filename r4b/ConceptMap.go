package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	Date              *FhirDateTime     `json:"date,omitempty"`
	Publisher         *string           `json:"publisher,omitempty"`
	Contact           []ContactDetail   `json:"contact,omitempty"`
	Description       *string           `json:"description,omitempty"`
	UseContext        []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose           *string           `json:"purpose,omitempty"`
	Copyright         *string           `json:"copyright,omitempty"`
	SourceUri         *string           `json:"sourceUri,omitempty"`
	SourceCanonical   *string           `json:"sourceCanonical,omitempty"`
	TargetUri         *string           `json:"targetUri,omitempty"`
	TargetCanonical   *string           `json:"targetCanonical,omitempty"`
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
func (r ConceptMap) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ConceptMap/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "ConceptMap"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ConceptMap) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ConceptMap) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ConceptMap) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ConceptMap) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ConceptMap) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConceptMap) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ConceptMap) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConceptMap) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ConceptMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ConceptMap) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ConceptMap) T_SourceUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("sourceUri", nil, htmlAttrs)
	}
	return StringInput("sourceUri", resource.SourceUri, htmlAttrs)
}
func (resource *ConceptMap) T_SourceCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("sourceCanonical", nil, htmlAttrs)
	}
	return StringInput("sourceCanonical", resource.SourceCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_TargetUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("targetUri", nil, htmlAttrs)
	}
	return StringInput("targetUri", resource.TargetUri, htmlAttrs)
}
func (resource *ConceptMap) T_TargetCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("targetCanonical", nil, htmlAttrs)
	}
	return StringInput("targetCanonical", resource.TargetCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSource(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].source", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].source", resource.Group[numGroup].Source, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSourceVersion(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].sourceVersion", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].sourceVersion", resource.Group[numGroup].SourceVersion, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTarget(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].target", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].target", resource.Group[numGroup].Target, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTargetVersion(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].targetVersion", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].targetVersion", resource.Group[numGroup].TargetVersion, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementCode(numGroup int, numElement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementDisplay(numGroup int, numElement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].display", resource.Group[numGroup].Element[numElement].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDisplay(numGroup int, numElement int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].display", resource.Group[numGroup].Element[numElement].Target[numTarget].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetEquivalence(numGroup int, numElement int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConcept_map_equivalence

	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].equivalence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].equivalence", &resource.Group[numGroup].Element[numElement].Target[numTarget].Equivalence, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].comment", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnProperty(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].property", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].property", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Property, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnSystem(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].system", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].system", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].System, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValue(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].value", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].value", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Value, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnDisplay(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].display", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedMode(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedCode(numGroup int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.code", resource.Group[numGroup].Unmapped.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedDisplay(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.display", resource.Group[numGroup].Unmapped.Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedUrl(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.url", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.url", resource.Group[numGroup].Unmapped.Url, htmlAttrs)
}
