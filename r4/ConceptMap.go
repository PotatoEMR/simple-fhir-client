package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *ConceptMap) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *ConceptMap) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ConceptMap) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *ConceptMap) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ConceptMap) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConceptMap) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ConceptMap) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConceptMap) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ConceptMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ConceptMap) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ConceptMap) T_SourceUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SourceUri", nil, htmlAttrs)
	}
	return StringInput("SourceUri", resource.SourceUri, htmlAttrs)
}
func (resource *ConceptMap) T_SourceCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SourceCanonical", nil, htmlAttrs)
	}
	return StringInput("SourceCanonical", resource.SourceCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_TargetUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TargetUri", nil, htmlAttrs)
	}
	return StringInput("TargetUri", resource.TargetUri, htmlAttrs)
}
func (resource *ConceptMap) T_TargetCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("TargetCanonical", nil, htmlAttrs)
	}
	return StringInput("TargetCanonical", resource.TargetCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSource(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Source", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Source", resource.Group[numGroup].Source, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSourceVersion(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]SourceVersion", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]SourceVersion", resource.Group[numGroup].SourceVersion, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTarget(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Target", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Target", resource.Group[numGroup].Target, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTargetVersion(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]TargetVersion", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]TargetVersion", resource.Group[numGroup].TargetVersion, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementCode(numGroup int, numElement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementDisplay(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Display", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Display", resource.Group[numGroup].Element[numElement].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDisplay(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", resource.Group[numGroup].Element[numElement].Target[numTarget].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetEquivalence(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_map_equivalence

	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Equivalence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Equivalence", &resource.Group[numGroup].Element[numElement].Target[numTarget].Equivalence, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnProperty(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Property", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Property", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Property, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnSystem(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].System", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].System", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].System, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValue(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Value", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Value", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Value, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnDisplay(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Display", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Display", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedMode(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Unmapped.Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Unmapped.Mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Unmapped.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Group["+strconv.Itoa(numGroup)+"]Unmapped.Code", resource.Group[numGroup].Unmapped.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedDisplay(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Unmapped.Display", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Unmapped.Display", resource.Group[numGroup].Unmapped.Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedUrl(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("Group["+strconv.Itoa(numGroup)+"]Unmapped.Url", nil, htmlAttrs)
	}
	return StringInput("Group["+strconv.Itoa(numGroup)+"]Unmapped.Url", resource.Group[numGroup].Unmapped.Url, htmlAttrs)
}
