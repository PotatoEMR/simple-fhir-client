package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSet struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Url               *string            `json:"url,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Version           *string            `json:"version,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Status            string             `json:"status"`
	Experimental      *bool              `json:"experimental,omitempty"`
	Date              *string            `json:"date,omitempty"`
	Publisher         *string            `json:"publisher,omitempty"`
	Contact           []ContactDetail    `json:"contact,omitempty"`
	Description       *string            `json:"description,omitempty"`
	UseContext        []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept  `json:"jurisdiction,omitempty"`
	Immutable         *bool              `json:"immutable,omitempty"`
	Purpose           *string            `json:"purpose,omitempty"`
	Copyright         *string            `json:"copyright,omitempty"`
	Compose           *ValueSetCompose   `json:"compose,omitempty"`
	Expansion         *ValueSetExpansion `json:"expansion,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeInclude struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	System            *string                         `json:"system,omitempty"`
	Version           *string                         `json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet          []string                        `json:"valueSet,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetComposeIncludeFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	Op                string      `json:"op"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansion struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Timestamp         string                       `json:"timestamp"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	ValueString       *string     `json:"valueString,omitempty"`
	ValueBoolean      *bool       `json:"valueBoolean,omitempty"`
	ValueInteger      *int        `json:"valueInteger,omitempty"`
	ValueDecimal      *float64    `json:"valueDecimal,omitempty"`
	ValueUri          *string     `json:"valueUri,omitempty"`
	ValueCode         *string     `json:"valueCode,omitempty"`
	ValueDateTime     *string     `json:"valueDateTime,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ValueSet
type ValueSetExpansionContains struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	System            *string     `json:"system,omitempty"`
	Abstract          *bool       `json:"abstract,omitempty"`
	Inactive          *bool       `json:"inactive,omitempty"`
	Version           *string     `json:"version,omitempty"`
	Code              *string     `json:"code,omitempty"`
	Display           *string     `json:"display,omitempty"`
}

type OtherValueSet ValueSet

// on convert struct to json, automatically add resourceType=ValueSet
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}
func (r ValueSet) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ValueSet/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ValueSet"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ValueSet) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ValueSet) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ValueSet) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ValueSet) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ValueSet) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ValueSet) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ValueSet) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ValueSet) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ValueSet) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_Immutable(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("immutable", nil, htmlAttrs)
	}
	return BoolInput("immutable", resource.Immutable, htmlAttrs)
}
func (resource *ValueSet) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ValueSet) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ValueSet) T_ComposeLockedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("compose.lockedDate", nil, htmlAttrs)
	}
	return DateInput("compose.lockedDate", resource.Compose.LockedDate, htmlAttrs)
}
func (resource *ValueSet) T_ComposeInactive(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("compose.inactive", nil, htmlAttrs)
	}
	return BoolInput("compose.inactive", resource.Compose.Inactive, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeSystem(numInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].system", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].system", resource.Compose.Include[numInclude].System, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeVersion(numInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].version", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].version", resource.Compose.Include[numInclude].Version, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeValueSet(numInclude int, numValueSet int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numValueSet >= len(resource.Compose.Include[numInclude].ValueSet) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].valueSet["+strconv.Itoa(numValueSet)+"]", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].valueSet["+strconv.Itoa(numValueSet)+"]", &resource.Compose.Include[numInclude].ValueSet[numValueSet], htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptCode(numInclude int, numConcept int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) {
		return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].code", &resource.Compose.Include[numInclude].Concept[numConcept].Code, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDisplay(numInclude int, numConcept int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].display", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].display", resource.Compose.Include[numInclude].Concept[numConcept].Display, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationUse(numInclude int, numConcept int, numDesignation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) || numDesignation >= len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) {
		return CodingSelect("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].use", resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Use, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationValue(numInclude int, numConcept int, numDesignation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) || numDesignation >= len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].value", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].value", &resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Value, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterProperty(numInclude int, numFilter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].property", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].property", &resource.Compose.Include[numInclude].Filter[numFilter].Property, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterOp(numInclude int, numFilter int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].op", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].op", &resource.Compose.Include[numInclude].Filter[numFilter].Op, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterValue(numInclude int, numFilter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return StringInput("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].value", nil, htmlAttrs)
	}
	return StringInput("compose.include["+strconv.Itoa(numInclude)+"].filter["+strconv.Itoa(numFilter)+"].value", &resource.Compose.Include[numInclude].Filter[numFilter].Value, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionTimestamp(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("expansion.timestamp", nil, htmlAttrs)
	}
	return DateTimeInput("expansion.timestamp", &resource.Expansion.Timestamp, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionTotal(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("expansion.total", nil, htmlAttrs)
	}
	return IntInput("expansion.total", resource.Expansion.Total, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionOffset(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("expansion.offset", nil, htmlAttrs)
	}
	return IntInput("expansion.offset", resource.Expansion.Offset, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterName(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].name", nil, htmlAttrs)
	}
	return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].name", &resource.Expansion.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueString(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueString", resource.Expansion.Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueBoolean(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return BoolInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueBoolean", resource.Expansion.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueInteger(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return IntInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueInteger", resource.Expansion.Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueDecimal(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return Float64Input("expansion.parameter["+strconv.Itoa(numParameter)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("expansion.parameter["+strconv.Itoa(numParameter)+"].valueDecimal", resource.Expansion.Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueUri(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueUri", resource.Expansion.Parameter[numParameter].ValueUri, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueCode(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return CodeSelect("expansion.parameter["+strconv.Itoa(numParameter)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("expansion.parameter["+strconv.Itoa(numParameter)+"].valueCode", resource.Expansion.Parameter[numParameter].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueDateTime(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return DateTimeInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("expansion.parameter["+strconv.Itoa(numParameter)+"].valueDateTime", resource.Expansion.Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsSystem(numContains int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].system", nil, htmlAttrs)
	}
	return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].system", resource.Expansion.Contains[numContains].System, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsAbstract(numContains int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return BoolInput("expansion.contains["+strconv.Itoa(numContains)+"].abstract", nil, htmlAttrs)
	}
	return BoolInput("expansion.contains["+strconv.Itoa(numContains)+"].abstract", resource.Expansion.Contains[numContains].Abstract, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsInactive(numContains int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return BoolInput("expansion.contains["+strconv.Itoa(numContains)+"].inactive", nil, htmlAttrs)
	}
	return BoolInput("expansion.contains["+strconv.Itoa(numContains)+"].inactive", resource.Expansion.Contains[numContains].Inactive, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsVersion(numContains int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].version", nil, htmlAttrs)
	}
	return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].version", resource.Expansion.Contains[numContains].Version, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsCode(numContains int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return CodeSelect("expansion.contains["+strconv.Itoa(numContains)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("expansion.contains["+strconv.Itoa(numContains)+"].code", resource.Expansion.Contains[numContains].Code, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsDisplay(numContains int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].display", nil, htmlAttrs)
	}
	return StringInput("expansion.contains["+strconv.Itoa(numContains)+"].display", resource.Expansion.Contains[numContains].Display, htmlAttrs)
}
