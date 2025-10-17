package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMap struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Url               string                  `json:"url"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Version           *string                 `json:"version,omitempty"`
	Name              string                  `json:"name"`
	Title             *string                 `json:"title,omitempty"`
	Status            string                  `json:"status"`
	Experimental      *bool                   `json:"experimental,omitempty"`
	Date              *FhirDateTime           `json:"date,omitempty"`
	Publisher         *string                 `json:"publisher,omitempty"`
	Contact           []ContactDetail         `json:"contact,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	UseContext        []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose           *string                 `json:"purpose,omitempty"`
	Copyright         *string                 `json:"copyright,omitempty"`
	Structure         []StructureMapStructure `json:"structure,omitempty"`
	Import            []string                `json:"import,omitempty"`
	Group             []StructureMapGroup     `json:"group"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapStructure struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Mode              string      `json:"mode"`
	Alias             *string     `json:"alias,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Name              string                   `json:"name"`
	Extends           *string                  `json:"extends,omitempty"`
	TypeMode          string                   `json:"typeMode"`
	Documentation     *string                  `json:"documentation,omitempty"`
	Input             []StructureMapGroupInput `json:"input"`
	Rule              []StructureMapGroupRule  `json:"rule"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupInput struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              *string     `json:"type,omitempty"`
	Mode              string      `json:"mode"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupRule struct {
	Id                *string                          `json:"id,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Name              string                           `json:"name"`
	Source            []StructureMapGroupRuleSource    `json:"source"`
	Target            []StructureMapGroupRuleTarget    `json:"target,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation     *string                          `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupRuleSource struct {
	Id                              *string              `json:"id,omitempty"`
	Extension                       []Extension          `json:"extension,omitempty"`
	ModifierExtension               []Extension          `json:"modifierExtension,omitempty"`
	Context                         string               `json:"context"`
	Min                             *int                 `json:"min,omitempty"`
	Max                             *string              `json:"max,omitempty"`
	Type                            *string              `json:"type,omitempty"`
	DefaultValueBase64Binary        *string              `json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean             *bool                `json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical           *string              `json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                *string              `json:"defaultValueCode,omitempty"`
	DefaultValueDate                *FhirDate            `json:"defaultValueDate,omitempty"`
	DefaultValueDateTime            *FhirDateTime        `json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal             *float64             `json:"defaultValueDecimal,omitempty"`
	DefaultValueId                  *string              `json:"defaultValueId,omitempty"`
	DefaultValueInstant             *string              `json:"defaultValueInstant,omitempty"`
	DefaultValueInteger             *int                 `json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown            *string              `json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                 *string              `json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt         *int                 `json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString              *string              `json:"defaultValueString,omitempty"`
	DefaultValueTime                *string              `json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt         *int                 `json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                 *string              `json:"defaultValueUri,omitempty"`
	DefaultValueUrl                 *string              `json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                *string              `json:"defaultValueUuid,omitempty"`
	DefaultValueAddress             *Address             `json:"defaultValueAddress,omitempty"`
	DefaultValueAge                 *Age                 `json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation          *Annotation          `json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment          *Attachment          `json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept     *CodeableConcept     `json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding              *Coding              `json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint        *ContactPoint        `json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount               *Count               `json:"defaultValueCount,omitempty"`
	DefaultValueDistance            *Distance            `json:"defaultValueDistance,omitempty"`
	DefaultValueDuration            *Duration            `json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName           *HumanName           `json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier          *Identifier          `json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney               *Money               `json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod              *Period              `json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity            *Quantity            `json:"defaultValueQuantity,omitempty"`
	DefaultValueRange               *Range               `json:"defaultValueRange,omitempty"`
	DefaultValueRatio               *Ratio               `json:"defaultValueRatio,omitempty"`
	DefaultValueReference           *Reference           `json:"defaultValueReference,omitempty"`
	DefaultValueSampledData         *SampledData         `json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature           *Signature           `json:"defaultValueSignature,omitempty"`
	DefaultValueTiming              *Timing              `json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail       *ContactDetail       `json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor         *Contributor         `json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement     *DataRequirement     `json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression          *Expression          `json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition *ParameterDefinition `json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact     *RelatedArtifact     `json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition   *TriggerDefinition   `json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext        *UsageContext        `json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage              *Dosage              `json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                *Meta                `json:"defaultValueMeta,omitempty"`
	Element                         *string              `json:"element,omitempty"`
	ListMode                        *string              `json:"listMode,omitempty"`
	Variable                        *string              `json:"variable,omitempty"`
	Condition                       *string              `json:"condition,omitempty"`
	Check                           *string              `json:"check,omitempty"`
	LogMessage                      *string              `json:"logMessage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupRuleTarget struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Context           *string                                `json:"context,omitempty"`
	ContextType       *string                                `json:"contextType,omitempty"`
	Element           *string                                `json:"element,omitempty"`
	Variable          *string                                `json:"variable,omitempty"`
	ListMode          []string                               `json:"listMode,omitempty"`
	ListRuleId        *string                                `json:"listRuleId,omitempty"`
	Transform         *string                                `json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueId           string      `json:"valueId"`
	ValueString       string      `json:"valueString"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDecimal      float64     `json:"valueDecimal"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureMap
type StructureMapGroupRuleDependent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Variable          []string    `json:"variable"`
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
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].typeMode", &resource.Group[numGroup].TypeMode, optionsValueSet, htmlAttrs)
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
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].name", &resource.Group[numGroup].Rule[numRule].Name, htmlAttrs)
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
func (resource *StructureMap) T_GroupRuleSourceDefaultValueBase64Binary(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueBase64Binary", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueBase64Binary, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueBoolean(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return BoolInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueBoolean", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueBoolean, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCanonical(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCanonical", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCanonical", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCanonical, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCode(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCode", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDate(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return FhirDateInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDate", nil, htmlAttrs)
	}
	return FhirDateInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDate", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDate, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDateTime(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDateTime", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDateTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDecimal(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return Float64Input("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDecimal", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDecimal, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueId(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueId", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueId", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueInstant(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueInstant", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueInstant", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueInstant, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueInteger(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueInteger", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueInteger", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueInteger, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueMarkdown(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMarkdown", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueMarkdown, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueOid(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueOid", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueOid", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueOid, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValuePositiveInt(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValuePositiveInt", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValuePositiveInt, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueString(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueString", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueString", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueString, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueTime(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTime", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTime", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUnsignedInt(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUnsignedInt", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUnsignedInt, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUri(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUri", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUri", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUri, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUrl(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUrl", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUrl", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUrl, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUuid(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUuid", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUuid", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUuid, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueAddress(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return AddressInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAddress", nil, htmlAttrs)
	}
	return AddressInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAddress", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueAddress, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueAge(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return AgeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAge", nil, htmlAttrs)
	}
	return AgeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAge", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueAge, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueAnnotation(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return AnnotationTextArea("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAnnotation", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueAnnotation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueAttachment(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return AttachmentInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueAttachment", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueAttachment, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCodeableConcept(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCodeableConcept", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCoding(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodingSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCoding", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueContactPoint(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ContactPointInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContactPoint", nil, htmlAttrs)
	}
	return ContactPointInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContactPoint", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueContactPoint, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCount(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CountInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCount", nil, htmlAttrs)
	}
	return CountInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueCount", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCount, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDistance(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DistanceInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDistance", nil, htmlAttrs)
	}
	return DistanceInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDistance", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDistance, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDuration(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DurationInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDuration", nil, htmlAttrs)
	}
	return DurationInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDuration", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDuration, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueHumanName(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return HumanNameInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueHumanName", nil, htmlAttrs)
	}
	return HumanNameInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueHumanName", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueHumanName, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueIdentifier(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IdentifierInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueIdentifier", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueIdentifier, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueMoney(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return MoneyInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMoney", nil, htmlAttrs)
	}
	return MoneyInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMoney", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueMoney, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValuePeriod(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return PeriodInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValuePeriod", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValuePeriod, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueQuantity(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return QuantityInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueQuantity", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueRange(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return RangeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRange", nil, htmlAttrs)
	}
	return RangeInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRange", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueRange, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueRatio(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return RatioInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRatio", nil, htmlAttrs)
	}
	return RatioInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRatio", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueRatio, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueReference(frs []FhirResource, numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueReference", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueReference, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueSampledData(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return SampledDataInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueSampledData", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueSampledData, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueSignature(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return SignatureInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueSignature", nil, htmlAttrs)
	}
	return SignatureInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueSignature", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueSignature, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueTiming(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return TimingInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTiming", nil, htmlAttrs)
	}
	return TimingInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTiming", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueTiming, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueContactDetail(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ContactDetailInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContactDetail", nil, htmlAttrs)
	}
	return ContactDetailInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContactDetail", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueContactDetail, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueContributor(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ContributorInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContributor", nil, htmlAttrs)
	}
	return ContributorInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueContributor", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueContributor, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDataRequirement(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DataRequirementInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDataRequirement", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDataRequirement, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueExpression(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ExpressionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueExpression", nil, htmlAttrs)
	}
	return ExpressionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueExpression", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueExpression, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueParameterDefinition(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return ParameterDefinitionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueParameterDefinition", nil, htmlAttrs)
	}
	return ParameterDefinitionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueParameterDefinition", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueParameterDefinition, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueRelatedArtifact(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return RelatedArtifactInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRelatedArtifact", nil, htmlAttrs)
	}
	return RelatedArtifactInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueRelatedArtifact", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueRelatedArtifact, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueTriggerDefinition(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return TriggerDefinitionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTriggerDefinition", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueTriggerDefinition", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueTriggerDefinition, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUsageContext(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return UsageContextInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUsageContext", nil, htmlAttrs)
	}
	return UsageContextInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueUsageContext", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUsageContext, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDosage(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DosageInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDosage", nil, htmlAttrs)
	}
	return DosageInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueDosage", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDosage, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueMeta(numGroup int, numRule int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return MetaInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMeta", nil, htmlAttrs)
	}
	return MetaInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].source["+strconv.Itoa(numSource)+"].defaultValueMeta", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueMeta, htmlAttrs)
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
func (resource *StructureMap) T_GroupRuleTargetContextType(numGroup int, numRule int, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMap_context_type

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].contextType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].target["+strconv.Itoa(numTarget)+"].contextType", resource.Group[numGroup].Rule[numRule].Target[numTarget].ContextType, optionsValueSet, htmlAttrs)
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
func (resource *StructureMap) T_GroupRuleDependentName(numGroup int, numRule int, numDependent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].name", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].name", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDependentVariable(numGroup int, numRule int, numDependent int, numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) || numVariable >= len(resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Variable) {
		return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].variable["+strconv.Itoa(numVariable)+"]", nil, htmlAttrs)
	}
	return StringInput("group["+strconv.Itoa(numGroup)+"].rule["+strconv.Itoa(numRule)+"].dependent["+strconv.Itoa(numDependent)+"].variable["+strconv.Itoa(numVariable)+"]", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Variable[numVariable], htmlAttrs)
}
