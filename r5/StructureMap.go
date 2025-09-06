package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMap struct {
	Id                     *string                 `json:"id,omitempty"`
	Meta                   *Meta                   `json:"meta,omitempty"`
	ImplicitRules          *string                 `json:"implicitRules,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Text                   *Narrative              `json:"text,omitempty"`
	Contained              []Resource              `json:"contained,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	Url                    string                  `json:"url"`
	Identifier             []Identifier            `json:"identifier,omitempty"`
	Version                *string                 `json:"version,omitempty"`
	VersionAlgorithmString *string                 `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                  `json:"name"`
	Title                  *string                 `json:"title,omitempty"`
	Status                 string                  `json:"status"`
	Experimental           *bool                   `json:"experimental,omitempty"`
	Date                   *string                 `json:"date,omitempty"`
	Publisher              *string                 `json:"publisher,omitempty"`
	Contact                []ContactDetail         `json:"contact,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	UseContext             []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                *string                 `json:"purpose,omitempty"`
	Copyright              *string                 `json:"copyright,omitempty"`
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty"`
	Structure              []StructureMapStructure `json:"structure,omitempty"`
	Import                 []string                `json:"import,omitempty"`
	Const                  []StructureMapConst     `json:"const,omitempty"`
	Group                  []StructureMapGroup     `json:"group"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapStructure struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Mode              string      `json:"mode"`
	Alias             *string     `json:"alias,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapConst struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Value             *string     `json:"value,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Name              string                   `json:"name"`
	Extends           *string                  `json:"extends,omitempty"`
	TypeMode          *string                  `json:"typeMode,omitempty"`
	Documentation     *string                  `json:"documentation,omitempty"`
	Input             []StructureMapGroupInput `json:"input"`
	Rule              []StructureMapGroupRule  `json:"rule,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupInput struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              *string     `json:"type,omitempty"`
	Mode              string      `json:"mode"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRule struct {
	Id                *string                          `json:"id,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Source            []StructureMapGroupRuleSource    `json:"source"`
	Target            []StructureMapGroupRuleTarget    `json:"target,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation     *string                          `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleSource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Context           string      `json:"context"`
	Min               *int        `json:"min,omitempty"`
	Max               *string     `json:"max,omitempty"`
	Type              *string     `json:"type,omitempty"`
	DefaultValue      *string     `json:"defaultValue,omitempty"`
	Element           *string     `json:"element,omitempty"`
	ListMode          *string     `json:"listMode,omitempty"`
	Variable          *string     `json:"variable,omitempty"`
	Condition         *string     `json:"condition,omitempty"`
	Check             *string     `json:"check,omitempty"`
	LogMessage        *string     `json:"logMessage,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleTarget struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Context           *string                                `json:"context,omitempty"`
	Element           *string                                `json:"element,omitempty"`
	Variable          *string                                `json:"variable,omitempty"`
	ListMode          []string                               `json:"listMode,omitempty"`
	ListRuleId        *string                                `json:"listRuleId,omitempty"`
	Transform         *string                                `json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueId           string      `json:"valueId"`
	ValueString       string      `json:"valueString"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueDate         string      `json:"valueDate"`
	ValueTime         string      `json:"valueTime"`
	ValueDateTime     string      `json:"valueDateTime"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleDependent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
}

type OtherStructureMap StructureMap

// on convert struct to json, automatically add resourceType=StructureMap
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

func (resource *StructureMap) T_Id() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Id", nil)
	}
	return StringInput("StructureMap.Id", resource.Id)
}
func (resource *StructureMap) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.ImplicitRules", nil)
	}
	return StringInput("StructureMap.ImplicitRules", resource.ImplicitRules)
}
func (resource *StructureMap) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("StructureMap.Language", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Language", resource.Language, optionsValueSet)
}
func (resource *StructureMap) T_Url() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Url", nil)
	}
	return StringInput("StructureMap.Url", &resource.Url)
}
func (resource *StructureMap) T_Version() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Version", nil)
	}
	return StringInput("StructureMap.Version", resource.Version)
}
func (resource *StructureMap) T_Name() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Name", nil)
	}
	return StringInput("StructureMap.Name", &resource.Name)
}
func (resource *StructureMap) T_Title() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Title", nil)
	}
	return StringInput("StructureMap.Title", resource.Title)
}
func (resource *StructureMap) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("StructureMap.Status", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Status", &resource.Status, optionsValueSet)
}
func (resource *StructureMap) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("StructureMap.Experimental", nil)
	}
	return BoolInput("StructureMap.Experimental", resource.Experimental)
}
func (resource *StructureMap) T_Date() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Date", nil)
	}
	return StringInput("StructureMap.Date", resource.Date)
}
func (resource *StructureMap) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Publisher", nil)
	}
	return StringInput("StructureMap.Publisher", resource.Publisher)
}
func (resource *StructureMap) T_Description() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Description", nil)
	}
	return StringInput("StructureMap.Description", resource.Description)
}
func (resource *StructureMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("StructureMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("StructureMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *StructureMap) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Purpose", nil)
	}
	return StringInput("StructureMap.Purpose", resource.Purpose)
}
func (resource *StructureMap) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Copyright", nil)
	}
	return StringInput("StructureMap.Copyright", resource.Copyright)
}
func (resource *StructureMap) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("StructureMap.CopyrightLabel", nil)
	}
	return StringInput("StructureMap.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *StructureMap) T_Import(numImport int) templ.Component {

	if resource == nil || len(resource.Import) >= numImport {
		return StringInput("StructureMap.Import["+strconv.Itoa(numImport)+"]", nil)
	}
	return StringInput("StructureMap.Import["+strconv.Itoa(numImport)+"]", &resource.Import[numImport])
}
func (resource *StructureMap) T_StructureId(numStructure int) templ.Component {

	if resource == nil || len(resource.Structure) >= numStructure {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Id", nil)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Id", resource.Structure[numStructure].Id)
}
func (resource *StructureMap) T_StructureUrl(numStructure int) templ.Component {

	if resource == nil || len(resource.Structure) >= numStructure {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Url", nil)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Url", &resource.Structure[numStructure].Url)
}
func (resource *StructureMap) T_StructureMode(numStructure int) templ.Component {
	optionsValueSet := VSMap_model_mode

	if resource == nil || len(resource.Structure) >= numStructure {
		return CodeSelect("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Mode", &resource.Structure[numStructure].Mode, optionsValueSet)
}
func (resource *StructureMap) T_StructureAlias(numStructure int) templ.Component {

	if resource == nil || len(resource.Structure) >= numStructure {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Alias", nil)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Alias", resource.Structure[numStructure].Alias)
}
func (resource *StructureMap) T_StructureDocumentation(numStructure int) templ.Component {

	if resource == nil || len(resource.Structure) >= numStructure {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Documentation", nil)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Documentation", resource.Structure[numStructure].Documentation)
}
func (resource *StructureMap) T_ConstId(numConst int) templ.Component {

	if resource == nil || len(resource.Const) >= numConst {
		return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Id", nil)
	}
	return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Id", resource.Const[numConst].Id)
}
func (resource *StructureMap) T_ConstName(numConst int) templ.Component {

	if resource == nil || len(resource.Const) >= numConst {
		return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Name", nil)
	}
	return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Name", resource.Const[numConst].Name)
}
func (resource *StructureMap) T_ConstValue(numConst int) templ.Component {

	if resource == nil || len(resource.Const) >= numConst {
		return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Value", nil)
	}
	return StringInput("StructureMap.Const["+strconv.Itoa(numConst)+"].Value", resource.Const[numConst].Value)
}
func (resource *StructureMap) T_GroupId(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Id", resource.Group[numGroup].Id)
}
func (resource *StructureMap) T_GroupName(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Name", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Name", &resource.Group[numGroup].Name)
}
func (resource *StructureMap) T_GroupExtends(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Extends", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Extends", resource.Group[numGroup].Extends)
}
func (resource *StructureMap) T_GroupTypeMode(numGroup int) templ.Component {
	optionsValueSet := VSMap_group_type_mode

	if resource == nil || len(resource.Group) >= numGroup {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].TypeMode", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].TypeMode", resource.Group[numGroup].TypeMode, optionsValueSet)
}
func (resource *StructureMap) T_GroupDocumentation(numGroup int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Documentation", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Documentation", resource.Group[numGroup].Documentation)
}
func (resource *StructureMap) T_GroupInputId(numGroup int, numInput int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Input) >= numInput {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Id", resource.Group[numGroup].Input[numInput].Id)
}
func (resource *StructureMap) T_GroupInputName(numGroup int, numInput int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Input) >= numInput {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Name", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Name", &resource.Group[numGroup].Input[numInput].Name)
}
func (resource *StructureMap) T_GroupInputType(numGroup int, numInput int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Input) >= numInput {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Type", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Type", resource.Group[numGroup].Input[numInput].Type)
}
func (resource *StructureMap) T_GroupInputMode(numGroup int, numInput int) templ.Component {
	optionsValueSet := VSMap_input_mode

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Input) >= numInput {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Mode", &resource.Group[numGroup].Input[numInput].Mode, optionsValueSet)
}
func (resource *StructureMap) T_GroupInputDocumentation(numGroup int, numInput int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Input) >= numInput {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Documentation", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Documentation", resource.Group[numGroup].Input[numInput].Documentation)
}
func (resource *StructureMap) T_GroupRuleId(numGroup int, numRule int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Id", resource.Group[numGroup].Rule[numRule].Id)
}
func (resource *StructureMap) T_GroupRuleName(numGroup int, numRule int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Name", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Name", resource.Group[numGroup].Rule[numRule].Name)
}
func (resource *StructureMap) T_GroupRuleDocumentation(numGroup int, numRule int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Documentation", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Documentation", resource.Group[numGroup].Rule[numRule].Documentation)
}
func (resource *StructureMap) T_GroupRuleSourceId(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Id", resource.Group[numGroup].Rule[numRule].Source[numSource].Id)
}
func (resource *StructureMap) T_GroupRuleSourceContext(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Context", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Context", &resource.Group[numGroup].Rule[numRule].Source[numSource].Context)
}
func (resource *StructureMap) T_GroupRuleSourceMin(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Min", nil)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Min", resource.Group[numGroup].Rule[numRule].Source[numSource].Min)
}
func (resource *StructureMap) T_GroupRuleSourceMax(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Max", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Max", resource.Group[numGroup].Rule[numRule].Source[numSource].Max)
}
func (resource *StructureMap) T_GroupRuleSourceType(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Type", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Type", resource.Group[numGroup].Rule[numRule].Source[numSource].Type)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValue(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValue", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValue", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValue)
}
func (resource *StructureMap) T_GroupRuleSourceElement(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Element", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Element", resource.Group[numGroup].Rule[numRule].Source[numSource].Element)
}
func (resource *StructureMap) T_GroupRuleSourceListMode(numGroup int, numRule int, numSource int) templ.Component {
	optionsValueSet := VSMap_source_list_mode

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].ListMode", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].ListMode", resource.Group[numGroup].Rule[numRule].Source[numSource].ListMode, optionsValueSet)
}
func (resource *StructureMap) T_GroupRuleSourceVariable(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Variable", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Variable", resource.Group[numGroup].Rule[numRule].Source[numSource].Variable)
}
func (resource *StructureMap) T_GroupRuleSourceCondition(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Condition", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Condition", resource.Group[numGroup].Rule[numRule].Source[numSource].Condition)
}
func (resource *StructureMap) T_GroupRuleSourceCheck(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Check", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Check", resource.Group[numGroup].Rule[numRule].Source[numSource].Check)
}
func (resource *StructureMap) T_GroupRuleSourceLogMessage(numGroup int, numRule int, numSource int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].LogMessage", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].LogMessage", resource.Group[numGroup].Rule[numRule].Source[numSource].LogMessage)
}
func (resource *StructureMap) T_GroupRuleTargetId(numGroup int, numRule int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Id", resource.Group[numGroup].Rule[numRule].Target[numTarget].Id)
}
func (resource *StructureMap) T_GroupRuleTargetContext(numGroup int, numRule int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Context", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Context", resource.Group[numGroup].Rule[numRule].Target[numTarget].Context)
}
func (resource *StructureMap) T_GroupRuleTargetElement(numGroup int, numRule int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Element", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Element", resource.Group[numGroup].Rule[numRule].Target[numTarget].Element)
}
func (resource *StructureMap) T_GroupRuleTargetVariable(numGroup int, numRule int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Variable", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Variable", resource.Group[numGroup].Rule[numRule].Target[numTarget].Variable)
}
func (resource *StructureMap) T_GroupRuleTargetListMode(numGroup int, numRule int, numTarget int, numListMode int) templ.Component {
	optionsValueSet := VSMap_target_list_mode

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget || len(resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode) >= numListMode {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListMode["+strconv.Itoa(numListMode)+"]", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListMode["+strconv.Itoa(numListMode)+"]", &resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode[numListMode], optionsValueSet)
}
func (resource *StructureMap) T_GroupRuleTargetListRuleId(numGroup int, numRule int, numTarget int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListRuleId", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListRuleId", resource.Group[numGroup].Rule[numRule].Target[numTarget].ListRuleId)
}
func (resource *StructureMap) T_GroupRuleTargetTransform(numGroup int, numRule int, numTarget int) templ.Component {
	optionsValueSet := VSMap_transform

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Transform", nil, optionsValueSet)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Transform", resource.Group[numGroup].Rule[numRule].Target[numTarget].Transform, optionsValueSet)
}
func (resource *StructureMap) T_GroupRuleTargetParameterId(numGroup int, numRule int, numTarget int, numParameter int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget || len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) >= numParameter {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].Id)
}
func (resource *StructureMap) T_GroupRuleDependentId(numGroup int, numRule int, numDependent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Dependent) >= numDependent {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Id", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Id", resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Id)
}
func (resource *StructureMap) T_GroupRuleDependentName(numGroup int, numRule int, numDependent int) templ.Component {

	if resource == nil || len(resource.Group) >= numGroup || len(resource.Group[numGroup].Rule) >= numRule || len(resource.Group[numGroup].Rule[numRule].Dependent) >= numDependent {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Name", nil)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Name", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Name)
}
