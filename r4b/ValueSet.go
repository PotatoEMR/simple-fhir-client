package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
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
	Date              *time.Time         `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *time.Time               `json:"lockedDate,omitempty,format:'2006-01-02'"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
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

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetComposeIncludeFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	Op                string      `json:"op"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetExpansion struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Timestamp         time.Time                    `json:"timestamp,format:'2006-01-02T15:04:05Z07:00'"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
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
	ValueDateTime     *time.Time  `json:"valueDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
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
func (resource *ValueSet) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *ValueSet) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ValueSet) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *ValueSet) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ValueSet) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ValueSet) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ValueSet) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ValueSet) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ValueSet) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_Immutable(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Immutable", nil, htmlAttrs)
	}
	return BoolInput("Immutable", resource.Immutable, htmlAttrs)
}
func (resource *ValueSet) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ValueSet) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ValueSet) T_ComposeLockedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ComposeLockedDate", nil, htmlAttrs)
	}
	return DateInput("ComposeLockedDate", resource.Compose.LockedDate, htmlAttrs)
}
func (resource *ValueSet) T_ComposeInactive(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ComposeInactive", nil, htmlAttrs)
	}
	return BoolInput("ComposeInactive", resource.Compose.Inactive, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeSystem(numInclude int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].System", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].System", resource.Compose.Include[numInclude].System, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeVersion(numInclude int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Version", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Version", resource.Compose.Include[numInclude].Version, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeValueSet(numInclude int, numValueSet int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numValueSet >= len(resource.Compose.Include[numInclude].ValueSet) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].ValueSet["+strconv.Itoa(numValueSet)+"]", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].ValueSet["+strconv.Itoa(numValueSet)+"]", &resource.Compose.Include[numInclude].ValueSet[numValueSet], htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptCode(numInclude int, numConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) {
		return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Code", &resource.Compose.Include[numInclude].Concept[numConcept].Code, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDisplay(numInclude int, numConcept int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Display", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Display", resource.Compose.Include[numInclude].Concept[numConcept].Display, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationUse(numInclude int, numConcept int, numDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) || numDesignation >= len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) {
		return CodingSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Use, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationValue(numInclude int, numConcept int, numDesignation int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numConcept >= len(resource.Compose.Include[numInclude].Concept) || numDesignation >= len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", &resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Value, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterProperty(numInclude int, numFilter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Property", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Property", &resource.Compose.Include[numInclude].Filter[numFilter].Property, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterOp(numInclude int, numFilter int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Op", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Op", &resource.Compose.Include[numInclude].Filter[numFilter].Op, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ComposeIncludeFilterValue(numInclude int, numFilter int, htmlAttrs string) templ.Component {
	if resource == nil || numInclude >= len(resource.Compose.Include) || numFilter >= len(resource.Compose.Include[numInclude].Filter) {
		return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Value", nil, htmlAttrs)
	}
	return StringInput("ComposeInclude["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Value", &resource.Compose.Include[numInclude].Filter[numFilter].Value, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionTimestamp(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ExpansionTimestamp", nil, htmlAttrs)
	}
	return DateTimeInput("ExpansionTimestamp", &resource.Expansion.Timestamp, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionTotal(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("ExpansionTotal", nil, htmlAttrs)
	}
	return IntInput("ExpansionTotal", resource.Expansion.Total, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionOffset(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("ExpansionOffset", nil, htmlAttrs)
	}
	return IntInput("ExpansionOffset", resource.Expansion.Offset, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterName(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].Name", &resource.Expansion.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueString(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueString", resource.Expansion.Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueBoolean(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return BoolInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueBoolean", resource.Expansion.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueInteger(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return IntInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueInteger", resource.Expansion.Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueDecimal(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return Float64Input("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueDecimal", resource.Expansion.Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueUri(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueUri", nil, htmlAttrs)
	}
	return StringInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueUri", resource.Expansion.Parameter[numParameter].ValueUri, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return CodeSelect("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueCode", resource.Expansion.Parameter[numParameter].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionParameterValueDateTime(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Expansion.Parameter) {
		return DateTimeInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("ExpansionParameter["+strconv.Itoa(numParameter)+"].ValueDateTime", resource.Expansion.Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsSystem(numContains int, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].System", nil, htmlAttrs)
	}
	return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].System", resource.Expansion.Contains[numContains].System, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsAbstract(numContains int, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return BoolInput("ExpansionContains["+strconv.Itoa(numContains)+"].Abstract", nil, htmlAttrs)
	}
	return BoolInput("ExpansionContains["+strconv.Itoa(numContains)+"].Abstract", resource.Expansion.Contains[numContains].Abstract, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsInactive(numContains int, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return BoolInput("ExpansionContains["+strconv.Itoa(numContains)+"].Inactive", nil, htmlAttrs)
	}
	return BoolInput("ExpansionContains["+strconv.Itoa(numContains)+"].Inactive", resource.Expansion.Contains[numContains].Inactive, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsVersion(numContains int, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].Version", nil, htmlAttrs)
	}
	return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].Version", resource.Expansion.Contains[numContains].Version, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsCode(numContains int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return CodeSelect("ExpansionContains["+strconv.Itoa(numContains)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExpansionContains["+strconv.Itoa(numContains)+"].Code", resource.Expansion.Contains[numContains].Code, optionsValueSet, htmlAttrs)
}
func (resource *ValueSet) T_ExpansionContainsDisplay(numContains int, htmlAttrs string) templ.Component {
	if resource == nil || numContains >= len(resource.Expansion.Contains) {
		return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].Display", nil, htmlAttrs)
	}
	return StringInput("ExpansionContains["+strconv.Itoa(numContains)+"].Display", resource.Expansion.Contains[numContains].Display, htmlAttrs)
}
