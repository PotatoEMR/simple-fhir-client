package r5

//generated with command go run ./bultaoreune
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
	Date                   *FhirDateTime           `json:"date,omitempty"`
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
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	ValueId           string       `json:"valueId"`
	ValueString       string       `json:"valueString"`
	ValueBoolean      bool         `json:"valueBoolean"`
	ValueInteger      int          `json:"valueInteger"`
	ValueDecimal      float64      `json:"valueDecimal"`
	ValueDate         FhirDate     `json:"valueDate"`
	ValueTime         string       `json:"valueTime"`
	ValueDateTime     FhirDateTime `json:"valueDateTime"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleDependent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
}

type OtherStructureMap StructureMap

// struct -> json, automatically add resourceType=Patient
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
func (r StructureMap) ResourceType() string {
	return "StructureMap"
}

func (resource *StructureMap) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *StructureMap) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *StructureMap) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *StructureMap) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *StructureMap) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *StructureMap) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *StructureMap) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *StructureMap) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *StructureMap) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *StructureMap) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *StructureMap) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *StructureMap) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *StructureMap) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *StructureMap) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *StructureMap) T_Import(numImport int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImport >= len(resource.Import) {
		return StringInput("import["+strconv.Itoa(numImport)+"]", nil, htmlAttrs)
	}
	return StringInput("import["+strconv.Itoa(numImport)+"]", &resource.Import[numImport], htmlAttrs)
}
func (resource *StructureMap) T_StructureUrl(numStructure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("structure["+strconv.Itoa(numStructure)+"].url", nil, htmlAttrs)
	}
	return StringInput("structure["+strconv.Itoa(numStructure)+"].url", &resource.Structure[numStructure].Url, htmlAttrs)
}
func (resource *StructureMap) T_StructureMode(numStructure int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_model_mode

	if resource == nil || numStructure >= len(resource.Structure) {
		return CodeSelect("structure["+strconv.Itoa(numStructure)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("structure["+strconv.Itoa(numStructure)+"].mode", &resource.Structure[numStructure].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_StructureAlias(numStructure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("structure["+strconv.Itoa(numStructure)+"].alias", nil, htmlAttrs)
	}
	return StringInput("structure["+strconv.Itoa(numStructure)+"].alias", resource.Structure[numStructure].Alias, htmlAttrs)
}
func (resource *StructureMap) T_StructureDocumentation(numStructure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("structure["+strconv.Itoa(numStructure)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("structure["+strconv.Itoa(numStructure)+"].documentation", resource.Structure[numStructure].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_ConstName(numConst int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConst >= len(resource.Const) {
		return StringInput("const["+strconv.Itoa(numConst)+"].name", nil, htmlAttrs)
	}
	return StringInput("const["+strconv.Itoa(numConst)+"].name", resource.Const[numConst].Name, htmlAttrs)
}
func (resource *StructureMap) T_ConstValue(numConst int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConst >= len(resource.Const) {
		return StringInput("const["+strconv.Itoa(numConst)+"].value", nil, htmlAttrs)
	}
	return StringInput("const["+strconv.Itoa(numConst)+"].value", resource.Const[numConst].Value, htmlAttrs)
}
func (resource *StructureMap) T_GroupName(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].name", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].name", &resource.Group[numGroup].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupExtends(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].extends", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].extends", resource.Group[numGroup].Extends, htmlAttrs)
}
func (resource *StructureMap) T_GroupTypeMode(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_group_type_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].typeMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].typeMode", resource.Group[numGroup].TypeMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupDocumentation(numGroup int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].documentation", resource.Group[numGroup].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputName(numGroup int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].name", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].name", &resource.Group[numGroup].Input[numInput].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputType(numGroup int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].type", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].type", resource.Group[numGroup].Input[numInput].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputMode(numGroup int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_input_mode

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].mode", &resource.Group[numGroup].Input[numInput].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputDocumentation(numGroup int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].input["+strconv.Itoa(numInput)+"].documentation", resource.Group[numGroup].Input[numInput].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleName(numGroup int, numRule int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].name", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].name", resource.Group[numGroup].Rule[numRule].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDocumentation(numGroup int, numRule int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].documentation", resource.Group[numGroup].Rule[numRule].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceContext(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].context", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].context", &resource.Group[numGroup].Rule[numRule].Source[numSource].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMin(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].min", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].min", resource.Group[numGroup].Rule[numRule].Source[numSource].Min, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMax(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].max", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].max", resource.Group[numGroup].Rule[numRule].Source[numSource].Max, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceType(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].type", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].type", resource.Group[numGroup].Rule[numRule].Source[numSource].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValue(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValue", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValue", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValue, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceElement(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].element", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].element", resource.Group[numGroup].Rule[numRule].Source[numSource].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceListMode(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_source_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].listMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].listMode", resource.Group[numGroup].Rule[numRule].Source[numSource].ListMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceVariable(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].variable", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].variable", resource.Group[numGroup].Rule[numRule].Source[numSource].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCondition(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].condition", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].condition", resource.Group[numGroup].Rule[numRule].Source[numSource].Condition, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCheck(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].check", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].check", resource.Group[numGroup].Rule[numRule].Source[numSource].Check, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceLogMessage(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].logMessage", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].logMessage", resource.Group[numGroup].Rule[numRule].Source[numSource].LogMessage, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetContext(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].context", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].context", resource.Group[numGroup].Rule[numRule].Target[numTarget].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetElement(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].element", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].element", resource.Group[numGroup].Rule[numRule].Target[numTarget].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetVariable(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].variable", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].variable", resource.Group[numGroup].Rule[numRule].Target[numTarget].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListMode(numGroup int, numRule int, numTarget int, numListMode int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_target_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numListMode >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].listMode["+strconv.Itoa(numListMode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].listMode["+strconv.Itoa(numListMode)+"]", &resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode[numListMode], optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListRuleId(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].listRuleId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].listRuleId", resource.Group[numGroup].Rule[numRule].Target[numTarget].ListRuleId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetTransform(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_transform

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].transform", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].transform", resource.Group[numGroup].Rule[numRule].Target[numTarget].Transform, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueId(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueId", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueString(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueBoolean(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueInteger(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueInteger", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDecimal(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return Float64Input("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDecimal", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDate(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return FhirDateInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDate", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDate, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueTime(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueTime", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDateTime(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].parameter["+strconv.Itoa(numParameter)+"].valueDateTime", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDependentName(numGroup int, numRule int, numDependent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].name", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].name", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Name, htmlAttrs)
}
