package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                *string              `json:"id,omitempty"`
	Meta              *Meta                `json:"meta,omitempty"`
	ImplicitRules     *string              `json:"implicitRules,omitempty"`
	Language          *string              `json:"language,omitempty"`
	Text              *Narrative           `json:"text,omitempty"`
	Contained         []Resource           `json:"contained,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Url               *string              `json:"url,omitempty"`
	Identifier        []Identifier         `json:"identifier,omitempty"`
	Version           *string              `json:"version,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Title             *string              `json:"title,omitempty"`
	Status            string               `json:"status"`
	Experimental      *bool                `json:"experimental,omitempty"`
	Date              *string              `json:"date,omitempty"`
	Publisher         *string              `json:"publisher,omitempty"`
	Contact           []ContactDetail      `json:"contact,omitempty"`
	Description       *string              `json:"description,omitempty"`
	UseContext        []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose           *string              `json:"purpose,omitempty"`
	Copyright         *string              `json:"copyright,omitempty"`
	CaseSensitive     *bool                `json:"caseSensitive,omitempty"`
	ValueSet          *string              `json:"valueSet,omitempty"`
	HierarchyMeaning  *string              `json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                `json:"compositional,omitempty"`
	VersionNeeded     *bool                `json:"versionNeeded,omitempty"`
	Content           string               `json:"content"`
	Supplements       *string              `json:"supplements,omitempty"`
	Count             *int                 `json:"count,omitempty"`
	Filter            []CodeSystemFilter   `json:"filter,omitempty"`
	Property          []CodeSystemProperty `json:"property,omitempty"`
	Concept           []CodeSystemConcept  `json:"concept,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystemFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Description       *string     `json:"description,omitempty"`
	Operator          []string    `json:"operator"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystemProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystemConcept struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              string                         `json:"code"`
	Display           *string                        `json:"display,omitempty"`
	Definition        *string                        `json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `json:"property,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystemConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CodeSystem
type CodeSystemConceptProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCode         string      `json:"valueCode"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueDecimal      float64     `json:"valueDecimal"`
}

type OtherCodeSystem CodeSystem

// on convert struct to json, automatically add resourceType=CodeSystem
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

func (resource *CodeSystem) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Id", nil)
	}
	return StringInput("CodeSystem.Id", resource.Id)
}
func (resource *CodeSystem) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.ImplicitRules", nil)
	}
	return StringInput("CodeSystem.ImplicitRules", resource.ImplicitRules)
}
func (resource *CodeSystem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CodeSystem.Language", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Language", resource.Language, optionsValueSet)
}
func (resource *CodeSystem) T_Url() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Url", nil)
	}
	return StringInput("CodeSystem.Url", resource.Url)
}
func (resource *CodeSystem) T_Version() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Version", nil)
	}
	return StringInput("CodeSystem.Version", resource.Version)
}
func (resource *CodeSystem) T_Name() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Name", nil)
	}
	return StringInput("CodeSystem.Name", resource.Name)
}
func (resource *CodeSystem) T_Title() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Title", nil)
	}
	return StringInput("CodeSystem.Title", resource.Title)
}
func (resource *CodeSystem) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CodeSystem.Status", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Status", &resource.Status, optionsValueSet)
}
func (resource *CodeSystem) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("CodeSystem.Experimental", nil)
	}
	return BoolInput("CodeSystem.Experimental", resource.Experimental)
}
func (resource *CodeSystem) T_Date() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Date", nil)
	}
	return StringInput("CodeSystem.Date", resource.Date)
}
func (resource *CodeSystem) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Publisher", nil)
	}
	return StringInput("CodeSystem.Publisher", resource.Publisher)
}
func (resource *CodeSystem) T_Description() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Description", nil)
	}
	return StringInput("CodeSystem.Description", resource.Description)
}
func (resource *CodeSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("CodeSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CodeSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *CodeSystem) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Purpose", nil)
	}
	return StringInput("CodeSystem.Purpose", resource.Purpose)
}
func (resource *CodeSystem) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Copyright", nil)
	}
	return StringInput("CodeSystem.Copyright", resource.Copyright)
}
func (resource *CodeSystem) T_CaseSensitive() templ.Component {

	if resource == nil {
		return BoolInput("CodeSystem.CaseSensitive", nil)
	}
	return BoolInput("CodeSystem.CaseSensitive", resource.CaseSensitive)
}
func (resource *CodeSystem) T_ValueSet() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.ValueSet", nil)
	}
	return StringInput("CodeSystem.ValueSet", resource.ValueSet)
}
func (resource *CodeSystem) T_HierarchyMeaning() templ.Component {
	optionsValueSet := VSCodesystem_hierarchy_meaning

	if resource == nil {
		return CodeSelect("CodeSystem.HierarchyMeaning", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.HierarchyMeaning", resource.HierarchyMeaning, optionsValueSet)
}
func (resource *CodeSystem) T_Compositional() templ.Component {

	if resource == nil {
		return BoolInput("CodeSystem.Compositional", nil)
	}
	return BoolInput("CodeSystem.Compositional", resource.Compositional)
}
func (resource *CodeSystem) T_VersionNeeded() templ.Component {

	if resource == nil {
		return BoolInput("CodeSystem.VersionNeeded", nil)
	}
	return BoolInput("CodeSystem.VersionNeeded", resource.VersionNeeded)
}
func (resource *CodeSystem) T_Content() templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil {
		return CodeSelect("CodeSystem.Content", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Content", &resource.Content, optionsValueSet)
}
func (resource *CodeSystem) T_Supplements() templ.Component {

	if resource == nil {
		return StringInput("CodeSystem.Supplements", nil)
	}
	return StringInput("CodeSystem.Supplements", resource.Supplements)
}
func (resource *CodeSystem) T_Count() templ.Component {

	if resource == nil {
		return IntInput("CodeSystem.Count", nil)
	}
	return IntInput("CodeSystem.Count", resource.Count)
}
func (resource *CodeSystem) T_FilterId(numFilter int) templ.Component {

	if resource == nil || len(resource.Filter) >= numFilter {
		return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Id", nil)
	}
	return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Id", resource.Filter[numFilter].Id)
}
func (resource *CodeSystem) T_FilterCode(numFilter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Filter) >= numFilter {
		return CodeSelect("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Code", &resource.Filter[numFilter].Code, optionsValueSet)
}
func (resource *CodeSystem) T_FilterDescription(numFilter int) templ.Component {

	if resource == nil || len(resource.Filter) >= numFilter {
		return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Description", nil)
	}
	return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Description", resource.Filter[numFilter].Description)
}
func (resource *CodeSystem) T_FilterOperator(numFilter int, numOperator int) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || len(resource.Filter) >= numFilter || len(resource.Filter[numFilter].Operator) >= numOperator {
		return CodeSelect("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Operator["+strconv.Itoa(numOperator)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Operator["+strconv.Itoa(numOperator)+"]", &resource.Filter[numFilter].Operator[numOperator], optionsValueSet)
}
func (resource *CodeSystem) T_FilterValue(numFilter int) templ.Component {

	if resource == nil || len(resource.Filter) >= numFilter {
		return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Value", nil)
	}
	return StringInput("CodeSystem.Filter["+strconv.Itoa(numFilter)+"].Value", &resource.Filter[numFilter].Value)
}
func (resource *CodeSystem) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *CodeSystem) T_PropertyCode(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeSelect("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Code", &resource.Property[numProperty].Code, optionsValueSet)
}
func (resource *CodeSystem) T_PropertyUri(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Uri", nil)
	}
	return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Uri", resource.Property[numProperty].Uri)
}
func (resource *CodeSystem) T_PropertyDescription(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Description", nil)
	}
	return StringInput("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Description", resource.Property[numProperty].Description)
}
func (resource *CodeSystem) T_PropertyType(numProperty int) templ.Component {
	optionsValueSet := VSConcept_property_type

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeSelect("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *CodeSystem) T_ConceptId(numConcept int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Id", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Id", resource.Concept[numConcept].Id)
}
func (resource *CodeSystem) T_ConceptCode(numConcept int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept {
		return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Code", &resource.Concept[numConcept].Code, optionsValueSet)
}
func (resource *CodeSystem) T_ConceptDisplay(numConcept int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Display", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Display", resource.Concept[numConcept].Display)
}
func (resource *CodeSystem) T_ConceptDefinition(numConcept int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Definition", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Definition", resource.Concept[numConcept].Definition)
}
func (resource *CodeSystem) T_ConceptDesignationId(numConcept int, numDesignation int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Designation) >= numDesignation {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Id", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Id", resource.Concept[numConcept].Designation[numDesignation].Id)
}
func (resource *CodeSystem) T_ConceptDesignationLanguage(numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Designation) >= numDesignation {
		return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Language", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Language", resource.Concept[numConcept].Designation[numDesignation].Language, optionsValueSet)
}
func (resource *CodeSystem) T_ConceptDesignationUse(numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Designation) >= numDesignation {
		return CodingSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", nil, optionsValueSet)
	}
	return CodingSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Use", resource.Concept[numConcept].Designation[numDesignation].Use, optionsValueSet)
}
func (resource *CodeSystem) T_ConceptDesignationValue(numConcept int, numDesignation int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Designation) >= numDesignation {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Designation["+strconv.Itoa(numDesignation)+"].Value", &resource.Concept[numConcept].Designation[numDesignation].Value)
}
func (resource *CodeSystem) T_ConceptPropertyId(numConcept int, numProperty int) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Property) >= numProperty {
		return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Property["+strconv.Itoa(numProperty)+"].Id", resource.Concept[numConcept].Property[numProperty].Id)
}
func (resource *CodeSystem) T_ConceptPropertyCode(numConcept int, numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Concept) >= numConcept || len(resource.Concept[numConcept].Property) >= numProperty {
		return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CodeSystem.Concept["+strconv.Itoa(numConcept)+"].Property["+strconv.Itoa(numProperty)+"].Code", &resource.Concept[numConcept].Property[numProperty].Code, optionsValueSet)
}
