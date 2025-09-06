package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

// http://hl7.org/fhir/r4b/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
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
	Timestamp         string                       `json:"timestamp"`
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
	ValueDateTime     *string     `json:"valueDateTime,omitempty"`
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

func (resource *ValueSet) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Id", nil)
	}
	return StringInput("ValueSet.Id", resource.Id)
}
func (resource *ValueSet) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.ImplicitRules", nil)
	}
	return StringInput("ValueSet.ImplicitRules", resource.ImplicitRules)
}
func (resource *ValueSet) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ValueSet.Language", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Language", resource.Language, optionsValueSet)
}
func (resource *ValueSet) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Url", nil)
	}
	return StringInput("ValueSet.Url", resource.Url)
}
func (resource *ValueSet) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Version", nil)
	}
	return StringInput("ValueSet.Version", resource.Version)
}
func (resource *ValueSet) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Name", nil)
	}
	return StringInput("ValueSet.Name", resource.Name)
}
func (resource *ValueSet) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Title", nil)
	}
	return StringInput("ValueSet.Title", resource.Title)
}
func (resource *ValueSet) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ValueSet.Status", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Status", &resource.Status, optionsValueSet)
}
func (resource *ValueSet) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ValueSet.Experimental", nil)
	}
	return BoolInput("ValueSet.Experimental", resource.Experimental)
}
func (resource *ValueSet) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Date", nil)
	}
	return StringInput("ValueSet.Date", resource.Date)
}
func (resource *ValueSet) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Publisher", nil)
	}
	return StringInput("ValueSet.Publisher", resource.Publisher)
}
func (resource *ValueSet) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Description", nil)
	}
	return StringInput("ValueSet.Description", resource.Description)
}
func (resource *ValueSet) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ValueSet.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ValueSet.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ValueSet) T_Immutable() templ.Component {

	if resource == nil {
		return BoolInput("ValueSet.Immutable", nil)
	}
	return BoolInput("ValueSet.Immutable", resource.Immutable)
}
func (resource *ValueSet) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Purpose", nil)
	}
	return StringInput("ValueSet.Purpose", resource.Purpose)
}
func (resource *ValueSet) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Copyright", nil)
	}
	return StringInput("ValueSet.Copyright", resource.Copyright)
}
func (resource *ValueSet) T_ComposeId() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Compose.Id", nil)
	}
	return StringInput("ValueSet.Compose.Id", resource.Compose.Id)
}
func (resource *ValueSet) T_ComposeLockedDate() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Compose.LockedDate", nil)
	}
	return StringInput("ValueSet.Compose.LockedDate", resource.Compose.LockedDate)
}
func (resource *ValueSet) T_ComposeInactive() templ.Component {

	if resource == nil {
		return BoolInput("ValueSet.Compose.Inactive", nil)
	}
	return BoolInput("ValueSet.Compose.Inactive", resource.Compose.Inactive)
}
func (resource *ValueSet) T_ComposeIncludeId(numInclude int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Id", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Id", resource.Compose.Include[numInclude].Id)
}
func (resource *ValueSet) T_ComposeIncludeSystem(numInclude int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].System", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].System", resource.Compose.Include[numInclude].System)
}
func (resource *ValueSet) T_ComposeIncludeVersion(numInclude int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Version", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Version", resource.Compose.Include[numInclude].Version)
}
func (resource *ValueSet) T_ComposeIncludeValueSet(numInclude int, numValueSet int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].ValueSet) >= numValueSet {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].ValueSet["+strconv.Itoa(numValueSet)+"]", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].ValueSet["+strconv.Itoa(numValueSet)+"]", &resource.Compose.Include[numInclude].ValueSet[numValueSet])
}
func (resource *ValueSet) T_ComposeIncludeConceptId(numInclude int, numConcept int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Id", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Id", resource.Compose.Include[numInclude].Concept[numConcept].Id)
}
func (resource *ValueSet) T_ComposeIncludeConceptCode(numInclude int, numConcept int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept {
		return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Code", &resource.Compose.Include[numInclude].Concept[numConcept].Code, optionsValueSet)
}
func (resource *ValueSet) T_ComposeIncludeConceptDisplay(numInclude int, numConcept int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Display", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Display", resource.Compose.Include[numInclude].Concept[numConcept].Display)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationId(numInclude int, numConcept int, numDesignation int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept || len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) >= numDesignation {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Id", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Id", resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Id)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationLanguage(numInclude int, numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept || len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) >= numDesignation {
		return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Language", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Language", resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Language, optionsValueSet)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationUse(numInclude int, numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept || len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) >= numDesignation {
		return CodingSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", nil, optionsValueSet)
	}
	return CodingSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Use, optionsValueSet)
}
func (resource *ValueSet) T_ComposeIncludeConceptDesignationValue(numInclude int, numConcept int, numDesignation int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Concept) >= numConcept || len(resource.Compose.Include[numInclude].Concept[numConcept].Designation) >= numDesignation {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", &resource.Compose.Include[numInclude].Concept[numConcept].Designation[numDesignation].Value)
}
func (resource *ValueSet) T_ComposeIncludeFilterId(numInclude int, numFilter int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Id", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Id", resource.Compose.Include[numInclude].Filter[numFilter].Id)
}
func (resource *ValueSet) T_ComposeIncludeFilterProperty(numInclude int, numFilter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Property", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Property", &resource.Compose.Include[numInclude].Filter[numFilter].Property, optionsValueSet)
}
func (resource *ValueSet) T_ComposeIncludeFilterOp(numInclude int, numFilter int) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Op", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Op", &resource.Compose.Include[numInclude].Filter[numFilter].Op, optionsValueSet)
}
func (resource *ValueSet) T_ComposeIncludeFilterValue(numInclude int, numFilter int) templ.Component {

	if resource == nil || len(resource.Compose.Include) >= numInclude || len(resource.Compose.Include[numInclude].Filter) >= numFilter {
		return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Value", nil)
	}
	return StringInput("ValueSet.Compose.Include["+strconv.Itoa(numInclude)+"].Filter["+strconv.Itoa(numFilter)+"].Value", &resource.Compose.Include[numInclude].Filter[numFilter].Value)
}
func (resource *ValueSet) T_ExpansionId() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Expansion.Id", nil)
	}
	return StringInput("ValueSet.Expansion.Id", resource.Expansion.Id)
}
func (resource *ValueSet) T_ExpansionIdentifier() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Expansion.Identifier", nil)
	}
	return StringInput("ValueSet.Expansion.Identifier", resource.Expansion.Identifier)
}
func (resource *ValueSet) T_ExpansionTimestamp() templ.Component {

	if resource == nil {
		return StringInput("ValueSet.Expansion.Timestamp", nil)
	}
	return StringInput("ValueSet.Expansion.Timestamp", &resource.Expansion.Timestamp)
}
func (resource *ValueSet) T_ExpansionTotal() templ.Component {

	if resource == nil {
		return IntInput("ValueSet.Expansion.Total", nil)
	}
	return IntInput("ValueSet.Expansion.Total", resource.Expansion.Total)
}
func (resource *ValueSet) T_ExpansionOffset() templ.Component {

	if resource == nil {
		return IntInput("ValueSet.Expansion.Offset", nil)
	}
	return IntInput("ValueSet.Expansion.Offset", resource.Expansion.Offset)
}
func (resource *ValueSet) T_ExpansionParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Expansion.Parameter) >= numParameter {
		return StringInput("ValueSet.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("ValueSet.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Expansion.Parameter[numParameter].Id)
}
func (resource *ValueSet) T_ExpansionParameterName(numParameter int) templ.Component {

	if resource == nil || len(resource.Expansion.Parameter) >= numParameter {
		return StringInput("ValueSet.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", nil)
	}
	return StringInput("ValueSet.Expansion.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Expansion.Parameter[numParameter].Name)
}
func (resource *ValueSet) T_ExpansionContainsId(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Id", nil)
	}
	return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Id", resource.Expansion.Contains[numContains].Id)
}
func (resource *ValueSet) T_ExpansionContainsSystem(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].System", nil)
	}
	return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].System", resource.Expansion.Contains[numContains].System)
}
func (resource *ValueSet) T_ExpansionContainsAbstract(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return BoolInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Abstract", nil)
	}
	return BoolInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Abstract", resource.Expansion.Contains[numContains].Abstract)
}
func (resource *ValueSet) T_ExpansionContainsInactive(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return BoolInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Inactive", nil)
	}
	return BoolInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Inactive", resource.Expansion.Contains[numContains].Inactive)
}
func (resource *ValueSet) T_ExpansionContainsVersion(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Version", nil)
	}
	return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Version", resource.Expansion.Contains[numContains].Version)
}
func (resource *ValueSet) T_ExpansionContainsCode(numContains int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return CodeSelect("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Code", resource.Expansion.Contains[numContains].Code, optionsValueSet)
}
func (resource *ValueSet) T_ExpansionContainsDisplay(numContains int) templ.Component {

	if resource == nil || len(resource.Expansion.Contains) >= numContains {
		return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Display", nil)
	}
	return StringInput("ValueSet.Expansion.Contains["+strconv.Itoa(numContains)+"].Display", resource.Expansion.Contains[numContains].Display)
}
