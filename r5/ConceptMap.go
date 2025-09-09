package r5

//generated with command go run ./bultaoreune
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
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ConceptMap) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ConceptMap) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ConceptMap) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ConceptMap) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ConceptMap) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConceptMap) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ConceptMap) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConceptMap) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ConceptMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ConceptMap) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ConceptMap) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ConceptMap) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ConceptMap) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ConceptMap) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_SourceScopeUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("sourceScopeUri", nil, htmlAttrs)
	}
	return StringInput("sourceScopeUri", resource.SourceScopeUri, htmlAttrs)
}
func (resource *ConceptMap) T_SourceScopeCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("sourceScopeCanonical", nil, htmlAttrs)
	}
	return StringInput("sourceScopeCanonical", resource.SourceScopeCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_TargetScopeUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("targetScopeUri", nil, htmlAttrs)
	}
	return StringInput("targetScopeUri", resource.TargetScopeUri, htmlAttrs)
}
func (resource *ConceptMap) T_TargetScopeCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("targetScopeCanonical", nil, htmlAttrs)
	}
	return StringInput("targetScopeCanonical", resource.TargetScopeCanonical, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("property["+strconv.Itoa(numProperty)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("property["+strconv.Itoa(numProperty)+"].code", &resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyUri(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].uri", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].uri", resource.Property[numProperty].Uri, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyDescription(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].description", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].description", resource.Property[numProperty].Description, htmlAttrs)
}
func (resource *ConceptMap) T_PropertyType(numProperty int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_property_type

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_PropertySystem(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].system", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].system", resource.Property[numProperty].System, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeCode(numAdditionalAttribute int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return CodeSelect("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].code", &resource.AdditionalAttribute[numAdditionalAttribute].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeUri(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return StringInput("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].uri", nil, htmlAttrs)
	}
	return StringInput("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].uri", resource.AdditionalAttribute[numAdditionalAttribute].Uri, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeDescription(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return StringInput("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].description", nil, htmlAttrs)
	}
	return StringInput("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].description", resource.AdditionalAttribute[numAdditionalAttribute].Description, htmlAttrs)
}
func (resource *ConceptMap) T_AdditionalAttributeType(numAdditionalAttribute int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_attribute_type

	if resource == nil || numAdditionalAttribute >= len(resource.AdditionalAttribute) {
		return CodeSelect("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("additionalAttribute["+strconv.Itoa(numAdditionalAttribute)+"].type", &resource.AdditionalAttribute[numAdditionalAttribute].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupSource(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].source", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].source", resource.Group[numGroup].Source, htmlAttrs)
}
func (resource *ConceptMap) T_GroupTarget(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].target", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].target", resource.Group[numGroup].Target, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementCode(numGroup int, numElement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].code", resource.Group[numGroup].Element[numElement].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementDisplay(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].display", resource.Group[numGroup].Element[numElement].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementValueSet(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].valueSet", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].valueSet", resource.Group[numGroup].Element[numElement].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementNoMap(numGroup int, numElement int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].noMap", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].noMap", resource.Group[numGroup].Element[numElement].NoMap, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetCode(numGroup int, numElement int, numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].code", resource.Group[numGroup].Element[numElement].Target[numTarget].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDisplay(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].display", resource.Group[numGroup].Element[numElement].Target[numTarget].Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetValueSet(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].valueSet", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].valueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetRelationship(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].relationship", &resource.Group[numGroup].Element[numElement].Target[numTarget].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetComment(numGroup int, numElement int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].comment", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].comment", resource.Group[numGroup].Element[numElement].Target[numTarget].Comment, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyCode(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].code", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueCoding(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodingSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueCoding", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueString(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueString", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueString, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueInteger(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueBoolean(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueDateTime(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return DateTimeInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueDateTime", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueDateTime, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueDecimal(numGroup int, numElement int, numTarget int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return Float64Input("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueDecimal", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueDecimal, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetPropertyValueCode(numGroup int, numElement int, numTarget int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numProperty >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].Property) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].property["+strconv.Itoa(numProperty)+"].valueCode", &resource.Group[numGroup].Element[numElement].Target[numTarget].Property[numProperty].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnAttribute(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].attribute", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].attribute", &resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].Attribute, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueCode(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueCode", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueCoding(numGroup int, numElement int, numTarget int, numDependsOn int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return CodingSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueCoding", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueString(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueString", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueString, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueBoolean(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueBoolean", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueBoolean, htmlAttrs)
}
func (resource *ConceptMap) T_GroupElementTargetDependsOnValueSet(numGroup int, numElement int, numTarget int, numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numElement >= len(resource.Group[numGroup].Element) || numTarget >= len(resource.Group[numGroup].Element[numElement].Target) || numDependsOn >= len(resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueSet", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].element["+strconv.Itoa(numElement)+"].target["+strconv.Itoa(numTarget)+"].dependsOn["+strconv.Itoa(numDependsOn)+"].valueSet", resource.Group[numGroup].Element[numElement].Target[numTarget].DependsOn[numDependsOn].ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedMode(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConceptmap_unmapped_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.mode", &resource.Group[numGroup].Unmapped.Mode, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedCode(numGroup int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.code", resource.Group[numGroup].Unmapped.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedDisplay(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.display", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.display", resource.Group[numGroup].Unmapped.Display, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedValueSet(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.valueSet", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.valueSet", resource.Group[numGroup].Unmapped.ValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedRelationship(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_map_relationship

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].unmapped.relationship", resource.Group[numGroup].Unmapped.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ConceptMap) T_GroupUnmappedOtherMap(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.otherMap", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].unmapped.otherMap", resource.Group[numGroup].Unmapped.OtherMap, htmlAttrs)
}
