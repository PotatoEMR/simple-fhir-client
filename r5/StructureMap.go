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
	Date                   *time.Time              `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate         time.Time   `json:"valueDate,format:'2006-01-02'"`
	ValueTime         string      `json:"valueTime"`
	ValueDateTime     time.Time   `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r StructureMap) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "StructureMap/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "StructureMap"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *StructureMap) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Url", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Url", &resource.Url, htmlAttrs)
}
func (resource *StructureMap) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Version", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Version", resource.Version, htmlAttrs)
}
func (resource *StructureMap) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("StructureMap.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *StructureMap) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("StructureMap.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("StructureMap.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Name", &resource.Name, htmlAttrs)
}
func (resource *StructureMap) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Title", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Title", resource.Title, htmlAttrs)
}
func (resource *StructureMap) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("StructureMap.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("StructureMap.Experimental", nil, htmlAttrs)
	}
	return BoolInput("StructureMap.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *StructureMap) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("StructureMap.Date", nil, htmlAttrs)
	}
	return DateTimeInput("StructureMap.Date", resource.Date, htmlAttrs)
}
func (resource *StructureMap) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Publisher", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *StructureMap) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Description", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Description", resource.Description, htmlAttrs)
}
func (resource *StructureMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("StructureMap.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StructureMap.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Purpose", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *StructureMap) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.Copyright", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *StructureMap) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureMap.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("StructureMap.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *StructureMap) T_Import(numImport int, htmlAttrs string) templ.Component {

	if resource == nil || numImport >= len(resource.Import) {
		return StringInput("StructureMap.Import."+strconv.Itoa(numImport)+".", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Import."+strconv.Itoa(numImport)+".", &resource.Import[numImport], htmlAttrs)
}
func (resource *StructureMap) T_StructureUrl(numStructure int, htmlAttrs string) templ.Component {

	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Url", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Url", &resource.Structure[numStructure].Url, htmlAttrs)
}
func (resource *StructureMap) T_StructureMode(numStructure int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_model_mode

	if resource == nil || numStructure >= len(resource.Structure) {
		return CodeSelect("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Mode", &resource.Structure[numStructure].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_StructureAlias(numStructure int, htmlAttrs string) templ.Component {

	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Alias", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Alias", resource.Structure[numStructure].Alias, htmlAttrs)
}
func (resource *StructureMap) T_StructureDocumentation(numStructure int, htmlAttrs string) templ.Component {

	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure."+strconv.Itoa(numStructure)+"..Documentation", resource.Structure[numStructure].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_ConstName(numConst int, htmlAttrs string) templ.Component {

	if resource == nil || numConst >= len(resource.Const) {
		return StringInput("StructureMap.Const."+strconv.Itoa(numConst)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Const."+strconv.Itoa(numConst)+"..Name", resource.Const[numConst].Name, htmlAttrs)
}
func (resource *StructureMap) T_ConstValue(numConst int, htmlAttrs string) templ.Component {

	if resource == nil || numConst >= len(resource.Const) {
		return StringInput("StructureMap.Const."+strconv.Itoa(numConst)+"..Value", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Const."+strconv.Itoa(numConst)+"..Value", resource.Const[numConst].Value, htmlAttrs)
}
func (resource *StructureMap) T_GroupName(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Name", &resource.Group[numGroup].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupExtends(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Extends", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Extends", resource.Group[numGroup].Extends, htmlAttrs)
}
func (resource *StructureMap) T_GroupTypeMode(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_group_type_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..TypeMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..TypeMode", resource.Group[numGroup].TypeMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupDocumentation(numGroup int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Documentation", resource.Group[numGroup].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputName(numGroup int, numInput int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Name", &resource.Group[numGroup].Input[numInput].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputType(numGroup int, numInput int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Type", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Type", resource.Group[numGroup].Input[numInput].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputMode(numGroup int, numInput int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_input_mode

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Mode", &resource.Group[numGroup].Input[numInput].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputDocumentation(numGroup int, numInput int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Input."+strconv.Itoa(numInput)+"..Documentation", resource.Group[numGroup].Input[numInput].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleName(numGroup int, numRule int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Name", resource.Group[numGroup].Rule[numRule].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDocumentation(numGroup int, numRule int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Documentation", resource.Group[numGroup].Rule[numRule].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceContext(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Context", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Context", &resource.Group[numGroup].Rule[numRule].Source[numSource].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMin(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Min", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Min", resource.Group[numGroup].Rule[numRule].Source[numSource].Min, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMax(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Max", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Max", resource.Group[numGroup].Rule[numRule].Source[numSource].Max, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceType(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Type", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Type", resource.Group[numGroup].Rule[numRule].Source[numSource].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValue(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..DefaultValue", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..DefaultValue", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValue, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceElement(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Element", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Element", resource.Group[numGroup].Rule[numRule].Source[numSource].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceListMode(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_source_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..ListMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..ListMode", resource.Group[numGroup].Rule[numRule].Source[numSource].ListMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceVariable(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Variable", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Variable", resource.Group[numGroup].Rule[numRule].Source[numSource].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCondition(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Condition", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Condition", resource.Group[numGroup].Rule[numRule].Source[numSource].Condition, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCheck(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Check", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..Check", resource.Group[numGroup].Rule[numRule].Source[numSource].Check, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceLogMessage(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..LogMessage", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Source."+strconv.Itoa(numSource)+"..LogMessage", resource.Group[numGroup].Rule[numRule].Source[numSource].LogMessage, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetContext(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Context", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Context", resource.Group[numGroup].Rule[numRule].Target[numTarget].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetElement(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Element", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Element", resource.Group[numGroup].Rule[numRule].Target[numTarget].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetVariable(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Variable", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Variable", resource.Group[numGroup].Rule[numRule].Target[numTarget].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListMode(numGroup int, numRule int, numTarget int, numListMode int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_target_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numListMode >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode) {
		return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..ListMode."+strconv.Itoa(numListMode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..ListMode."+strconv.Itoa(numListMode)+".", &resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode[numListMode], optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListRuleId(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..ListRuleId", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..ListRuleId", resource.Group[numGroup].Rule[numRule].Target[numTarget].ListRuleId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetTransform(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_transform

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Transform", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Transform", resource.Group[numGroup].Rule[numRule].Target[numTarget].Transform, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueId(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueId", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueId", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueString(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueString", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueBoolean(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return BoolInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueBoolean", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueInteger(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return IntInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueInteger", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDecimal(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return Float64Input("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDecimal", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDate(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return DateInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDate", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDate, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueTime(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueTime", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueTime", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDateTime(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return DateTimeInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Target."+strconv.Itoa(numTarget)+"..Parameter."+strconv.Itoa(numParameter)+"..ValueDateTime", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDependentName(numGroup int, numRule int, numDependent int, htmlAttrs string) templ.Component {

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) {
		return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Dependent."+strconv.Itoa(numDependent)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group."+strconv.Itoa(numGroup)+"..Rule."+strconv.Itoa(numRule)+"..Dependent."+strconv.Itoa(numDependent)+"..Name", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Name, htmlAttrs)
}
