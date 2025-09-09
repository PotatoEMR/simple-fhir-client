package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                   *time.Time                      `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                         `json:"publisher,omitempty"`
	Contact                []ContactDetail                 `json:"contact,omitempty"`
	Description            *string                         `json:"description,omitempty"`
	UseContext             []UsageContext                  `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept               `json:"jurisdiction,omitempty"`
	Purpose                *string                         `json:"purpose,omitempty"`
	Copyright              *string                         `json:"copyright,omitempty"`
	CopyrightLabel         *string                         `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time                      `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                      `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	ValueDateTime     time.Time   `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ConceptMap) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ConceptMap/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ConceptMap"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ConceptMap) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Url", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Url", resource.Url, htmlAttrs)
}
func (resource *ConceptMap) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Version", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Version", resource.Version, htmlAttrs)
}
func (resource *ConceptMap) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ConceptMap) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("ConceptMap.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ConceptMap.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Name", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Name", resource.Name, htmlAttrs)
}
func (resource *ConceptMap) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Title", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Title", resource.Title, htmlAttrs)
}
func (resource *ConceptMap) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ConceptMap.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ConceptMap.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ConceptMap.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConceptMap) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ConceptMap.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ConceptMap.Date", resource.Date, htmlAttrs)
}
func (resource *ConceptMap) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Publisher", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConceptMap) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Description", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Description", resource.Description, htmlAttrs)
}
func (resource *ConceptMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ConceptMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConceptMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Purpose", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ConceptMap) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.Copyright", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ConceptMap) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ConceptMap) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ConceptMap.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ConceptMap.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ConceptMap) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ConceptMap.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("ConceptMap.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ConceptMap) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("ConceptMap.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConceptMap.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_SourceScopeUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.SourceScopeUri", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.SourceScopeUri", resource.SourceScopeUri, htmlAttrs)
}
func (resource *ConceptMap) T_SourceScopeCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.SourceScopeCanonical", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.SourceScopeCanonical", resource.SourceScopeCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_TargetScopeUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.TargetScopeUri", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.TargetScopeUri", resource.TargetScopeUri, htmlAttrs)
}
func (resource *ConceptMap) T_TargetScopeCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConceptMap.TargetScopeCanonical", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.TargetScopeCanonical", resource.TargetScopeCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Code", &resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyUri(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Uri", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Uri", resource.Property[numProperty].Uri, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyDescription(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Description", resource.Property[numProperty].Description, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyType(numProperty int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_property_type

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_PropertySystem(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].System", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Property["+strconv.Itoa(numProperty)+"].System", resource.Property[numProperty].System, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeCode(numAdditionalAttribute int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Code", &resource.AdditionalAttribute[numAdditionalAttribute].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeUri(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Uri", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Uri", resource.AdditionalAttribute[numAdditionalAttribute].Uri, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeDescription(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Description", resource.AdditionalAttribute[numAdditionalAttribute].Description, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeType(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_attribute_type

	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.AdditionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].Type", &resource.AdditionalAttribute[numAdditionalAttribute].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSource(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Source", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Source", resource.Group[numGroup].Source, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTarget(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Target", resource.Group[numGroup].Target, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementCode(numGroup int, numElement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementDisplay(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Display", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Display", resource.Group[numGroup].Element[numElement].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementValueSet(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].ValueSet", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].ValueSet", resource.Group[numGroup].Element[numElement].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementNoMap(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].NoMap", nil, htmlAttrs)
	}
	return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].NoMap", resource.Group[numGroup].Element[numElement].NoMap, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDisplay(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Display", resource.Group[numGroup].Element[numElement].Target[numTarget].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetValueSet(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].ValueSet", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].ValueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetRelationship(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Relationship", &resource.Group[numGroup].Element[numElement].Target[numTarget].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyCode(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].Code", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueCoding(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodingSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueCoding", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueString(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueString", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueString, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueInteger(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return IntInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueInteger", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueBoolean(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueBoolean", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueDateTime(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return DateTimeInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueDateTime", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueDateTime, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueDecimal(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return Float64Input("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueDecimal", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueDecimal, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueCode(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].Property["+strconv.Itoa(numProperty)+"].ValueCode", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnAttribute(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Attribute", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].Attribute", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Attribute, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueCode(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueCode", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueCoding(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodingSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueCoding", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueString(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueString", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueString, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueBoolean(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueBoolean", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueBoolean, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueSet(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueSet", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Element["+strconv.Itoa(numElement)+"].Target["+strconv.Itoa(numTarget)+"].DependsOn["+strconv.Itoa(numDependsOn)+"].ValueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedMode(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Code", resource.Group[numGroup].Unmapped.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedDisplay(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Display", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Display", resource.Group[numGroup].Unmapped.Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedValueSet(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.ValueSet", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.ValueSet", resource.Group[numGroup].Unmapped.ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedRelationship(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.Relationship", resource.Group[numGroup].Unmapped.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedOtherMap(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.OtherMap", nil, htmlAttrs)
	}
	return StringInput("ConceptMap.Group["+strconv.Itoa(numGroup)+"].Unmapped.OtherMap", resource.Group[numGroup].Unmapped.OtherMap, htmlAttrs)
}
