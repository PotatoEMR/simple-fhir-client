package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	VersionAlgorithmString *string                         `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding,omitempty"`
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
	SourceScopeUri         *string                         `json:"sourceScopeUri,omitempty"`
	SourceScopeCanonical   *string                         `json:"sourceScopeCanonical,omitempty"`
	TargetScopeUri         *string                         `json:"targetScopeUri,omitempty"`
	TargetScopeCanonical   *string                         `json:"targetScopeCanonical,omitempty"`
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
	ValueCode         *string     `json:"valueCode,omitempty"`
	ValueCoding       *Coding     `json:"valueCoding,omitempty"`
	ValueString       *string     `json:"valueString,omitempty"`
	ValueBoolean      *bool       `json:"valueBoolean,omitempty"`
	ValueQuantity     *Quantity   `json:"valueQuantity,omitempty"`
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
func (resource *ConceptMap) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.CopyrightLabel", nil)
	}
	return StringInput("ConceptMap.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *ConceptMap) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.ApprovalDate", nil)
	}
	return StringInput("ConceptMap.ApprovalDate", resource.ApprovalDate)
}
func (resource *ConceptMap) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ConceptMap.LastReviewDate", nil)
	}
	return StringInput("ConceptMap.LastReviewDate", resource.LastReviewDate)
}
func (resource *ConceptMap) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("ConceptMap.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ConceptMap.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *ConceptMap) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *ConceptMap) T_PropertyCode(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Code", &resource.Property[numProperty].Code, optionsValueSet)
}
func (resource *ConceptMap) T_PropertyUri(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Uri", nil)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Uri", resource.Property[numProperty].Uri)
}
func (resource *ConceptMap) T_PropertyDescription(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Description", nil)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Description", resource.Property[numProperty].Description)
}
func (resource *ConceptMap) T_PropertyType(numProperty int) templ.Component {
	optionsValueSet := VSConceptmap_property_type

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *ConceptMap) T_PropertySystem(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].System", nil)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].System", resource.Property[numProperty].System)
}
func (resource *ConceptMap) T_AdditionalAttributeId(numAdditionalAttribute int) templ.Component {

	if resource == nil || len(resource.AdditionalAttribute) >= numAdditionalAttribute {
		return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Id", nil)
	}
	return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Id", resource.AdditionalAttribute[numAdditionalAttribute].Id)
}
func (resource *ConceptMap) T_AdditionalAttributeCode(numAdditionalAttribute int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AdditionalAttribute) >= numAdditionalAttribute {
		return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Code", &resource.AdditionalAttribute[numAdditionalAttribute].Code, optionsValueSet)
}
func (resource *ConceptMap) T_AdditionalAttributeUri(numAdditionalAttribute int) templ.Component {

	if resource == nil || len(resource.AdditionalAttribute) >= numAdditionalAttribute {
		return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Uri", nil)
	}
	return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Uri", resource.AdditionalAttribute[numAdditionalAttribute].Uri)
}
func (resource *ConceptMap) T_AdditionalAttributeDescription(numAdditionalAttribute int) templ.Component {

	if resource == nil || len(resource.AdditionalAttribute) >= numAdditionalAttribute {
		return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Description", nil)
	}
	return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Description", resource.AdditionalAttribute[numAdditionalAttribute].Description)
}
func (resource *ConceptMap) T_AdditionalAttributeType(numAdditionalAttribute int) templ.Component {
	optionsValueSet := VSConceptmap_attribute_type

	if resource == nil || len(resource.AdditionalAttribute) >= numAdditionalAttribute {
		return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Type", &resource.AdditionalAttribute[numAdditionalAttribute].Type, optionsValueSet)
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
func (resource *ConceptMap) T_GroupTarget(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", resource.Group[numGroup].Target)
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
func (resource *ConceptMap) T_GroupElementValueSet(numGroup int, numElement int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].ValueSet", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].ValueSet", resource.Group[numGroup].Element[numElement].ValueSet)
}
func (resource *ConceptMap) T_GroupElementNoMap(numGroup int, numElement int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement {
		return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].NoMap", nil)
	}
	return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].NoMap", resource.Group[numGroup].Element[numElement].NoMap)
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
func (resource *ConceptMap) T_GroupElementTargetValueSet(numGroup int, numElement int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].ValueSet", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].ValueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].ValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetRelationship(numGroup int, numElement int, numTarget int) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Relationship", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Relationship", &resource.Group[numGroup].Element[numElement].Target[numTarget].Relationship, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyId(numGroup int, numElement int, numTarget int, numProperty int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) >= numProperty {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Id", resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].Id)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyCode(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) >= numProperty {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Code", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].Code, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnId(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Id", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Id", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Id)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnAttribute(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Attribute", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Attribute", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Attribute, optionsValueSet)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueSet(numGroup int, numElement int, numTarget int, numDependsOn int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Element) >= numElement || len(resource.Group[numGroup].Element[numElement].Target) >= numTarget || len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) >= numDependsOn {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueSet", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueSet)
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
func (resource *ConceptMap) T_GroupUnmappedValueSet(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.ValueSet", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.ValueSet", resource.Group[numGroup].Unmapped.ValueSet)
}
func (resource *ConceptMap) T_GroupUnmappedRelationship(numGroup int) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Relationship", nil, optionsValueSet)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Relationship", resource.Group[numGroup].Unmapped.Relationship, optionsValueSet)
}
func (resource *ConceptMap) T_GroupUnmappedOtherMap(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.OtherMap", nil)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.OtherMap", resource.Group[numGroup].Unmapped.OtherMap)
}
