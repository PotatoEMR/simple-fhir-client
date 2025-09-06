package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
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
	SourceUri         *string           `json:"sourceUri,omitempty"`
	SourceCanonical   *string           `json:"sourceCanonical,omitempty"`
	TargetUri         *string           `json:"targetUri,omitempty"`
	TargetCanonical   *string           `json:"targetCanonical,omitempty"`
	Group             []ConceptMapGroup `json:"group,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
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

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
type ConceptMapGroupElement struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *string                        `json:"code,omitempty"`
	Display           *string                        `json:"display,omitempty"`
	Target            []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
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

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	System            *string     `json:"system,omitempty"`
	Value             string      `json:"value"`
	Display           *string     `json:"display,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ConceptMap
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

func (resource *ConceptMap) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Id", nil)
	}
	return StringInput("ConceptMap.Id", resource.Id)
}
func (resource *ConceptMap) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.ImplicitRules", nil)
	}
	return StringInput("ConceptMap.ImplicitRules", resource.ImplicitRules)
}
func (resource *ConceptMap) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ConceptMap.Language", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Language", resource.Language, optionsValueSet)
}
func (resource *ConceptMap) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Url", nil)
	}
	return StringInput("ConceptMap.Url", resource.Url)
}
func (resource *ConceptMap) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Version", nil)
	}
	return StringInput("ConceptMap.Version", resource.Version)
}
func (resource *ConceptMap) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Name", nil)
	}
	return StringInput("ConceptMap.Name", resource.Name)
}
func (resource *ConceptMap) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Title", nil)
	}
	return StringInput("ConceptMap.Title", resource.Title)
}
func (resource *ConceptMap) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ConceptMap.Status", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Status", &resource.Status, optionsValueSet)
}
func (resource *ConceptMap) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ConceptMap.Experimental", nil)
	}
	return BoolInput("ConceptMap.Experimental", resource.Experimental)
}
func (resource *ConceptMap) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Date", nil)
	}
	return StringInput("ConceptMap.Date", resource.Date)
}
func (resource *ConceptMap) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Publisher", nil)
	}
	return StringInput("ConceptMap.Publisher", resource.Publisher)
}
func (resource *ConceptMap) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Description", nil)
	}
	return StringInput("ConceptMap.Description", resource.Description)
}
func (resource *ConceptMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ConceptMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConceptMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ConceptMap) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Purpose", nil)
	}
	return StringInput("ConceptMap.Purpose", resource.Purpose)
}
func (resource *ConceptMap) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.Copyright", nil)
	}
	return StringInput("ConceptMap.Copyright", resource.Copyright)
}
func (resource *ConceptMap) T_GroupId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Id", resource.Group[numGroup].Id)
}
func (resource *ConceptMap) T_GroupSource(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Source", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Source", resource.Group[numGroup].Source)
}
func (resource *ConceptMap) T_GroupSourceVersion(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].SourceVersion", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].SourceVersion", resource.Group[numGroup].SourceVersion)
}
func (resource *ConceptMap) T_GroupTarget(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", resource.Group[numGroup].Target)
}
func (resource *ConceptMap) T_GroupTargetVersion(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].TargetVersion", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].TargetVersion", resource.Group[numGroup].TargetVersion)
}
func (resource *ConceptMap) T_GroupElementId(numGroup int, numElement int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Id", resource.Group[numGroup].Element[numElement].Id)
}
func (resource *ConceptMap) T_GroupElementCode(numGroup int, numElement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementDisplay(numGroup int, numElement int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Display", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Display", resource.Group[numGroup].Element[numElement].Display)
}
func (resource *ConceptMap) T_GroupElementTargetId(numGroup int, numElement int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Id", resource.Group[numGroup].Element[numElement].Target[numTarget].Id)
}
func (resource *ConceptMap) T_GroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetDisplay(numGroup int, numElement int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", resource.Group[numGroup].Element[numElement].Target[numTarget].Display)
}
func (resource *ConceptMap) T_GroupElementTargetEquivalence(numGroup int, numElement int, numTarget int) templ.Component {
	optionsValueSet := VSConcept_map_equivalence

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Equivalence", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Equivalence", &resource.Group[numGroup].Element[numElement].Target[numTarget].Equivalence, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnId(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Id", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Id)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnProperty(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Property", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Property", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Property)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnSystem(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].System", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].System", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].System)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValue(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Value", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Value", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Value)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnDisplay(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Display", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Display", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Display)
}
func (resource *ConceptMap) T_GroupUnmappedId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Id", resource.Group[numGroup].Unmapped.Id)
}
func (resource *ConceptMap) T_GroupUnmappedMode(numGroup int) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Mode", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet)
}
func (resource *ConceptMap) T_GroupUnmappedCode(numGroup int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Code", resource.Group[numGroup].Unmapped.Code, optionsValueSet)
}
func (resource *ConceptMap) T_GroupUnmappedDisplay(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Display", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Display", resource.Group[numGroup].Unmapped.Display)
}
func (resource *ConceptMap) T_GroupUnmappedUrl(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Url", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Url", resource.Group[numGroup].Unmapped.Url)
}
