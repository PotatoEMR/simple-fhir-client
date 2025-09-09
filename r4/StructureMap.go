package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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
	Date              *time.Time              `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
type StructureMapStructure struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Mode              string      `json:"mode"`
	Alias             *string     `json:"alias,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
type StructureMapGroupInput struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              *string     `json:"type,omitempty"`
	Mode              string      `json:"mode"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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
	DefaultValueDate                *time.Time           `json:"defaultValueDate,omitempty,format:'2006-01-02'"`
	DefaultValueDateTime            *time.Time           `json:"defaultValueDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
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

// http://hl7.org/fhir/r4/StructureDefinition/StructureMap
type StructureMapGroupRuleDependent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Variable          []string    `json:"variable"`
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
		return CodeableConceptSelect("StructureMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StructureMap.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
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
func (resource *StructureMap) T_Import(numImport int, htmlAttrs string) templ.Component {
	if resource == nil || numImport >= len(resource.Import) {
		return StringInput("StructureMap.Import["+strconv.Itoa(numImport)+"]", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Import["+strconv.Itoa(numImport)+"]", &resource.Import[numImport], htmlAttrs)
}
func (resource *StructureMap) T_StructureUrl(numStructure int, htmlAttrs string) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Url", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Url", &resource.Structure[numStructure].Url, htmlAttrs)
}
func (resource *StructureMap) T_StructureMode(numStructure int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_model_mode

	if resource == nil || numStructure >= len(resource.Structure) {
		return CodeSelect("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Mode", &resource.Structure[numStructure].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_StructureAlias(numStructure int, htmlAttrs string) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Alias", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Alias", resource.Structure[numStructure].Alias, htmlAttrs)
}
func (resource *StructureMap) T_StructureDocumentation(numStructure int, htmlAttrs string) templ.Component {
	if resource == nil || numStructure >= len(resource.Structure) {
		return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Structure["+strconv.Itoa(numStructure)+"].Documentation", resource.Structure[numStructure].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupName(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Name", &resource.Group[numGroup].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupExtends(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Extends", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Extends", resource.Group[numGroup].Extends, htmlAttrs)
}
func (resource *StructureMap) T_GroupTypeMode(numGroup int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_group_type_mode

	if resource == nil || numGroup >= len(resource.Group) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].TypeMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].TypeMode", &resource.Group[numGroup].TypeMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupDocumentation(numGroup int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Documentation", resource.Group[numGroup].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputName(numGroup int, numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Name", &resource.Group[numGroup].Input[numInput].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputType(numGroup int, numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Type", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Type", resource.Group[numGroup].Input[numInput].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputMode(numGroup int, numInput int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_input_mode

	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Mode", &resource.Group[numGroup].Input[numInput].Mode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupInputDocumentation(numGroup int, numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numInput >= len(resource.Group[numGroup].Input) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Input["+strconv.Itoa(numInput)+"].Documentation", resource.Group[numGroup].Input[numInput].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleName(numGroup int, numRule int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Name", &resource.Group[numGroup].Rule[numRule].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDocumentation(numGroup int, numRule int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Documentation", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Documentation", resource.Group[numGroup].Rule[numRule].Documentation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceContext(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Context", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Context", &resource.Group[numGroup].Rule[numRule].Source[numSource].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMin(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Min", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Min", resource.Group[numGroup].Rule[numRule].Source[numSource].Min, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceMax(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Max", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Max", resource.Group[numGroup].Rule[numRule].Source[numSource].Max, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceType(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Type", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Type", resource.Group[numGroup].Rule[numRule].Source[numSource].Type, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueBase64Binary(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueBase64Binary", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueBase64Binary, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueBoolean(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return BoolInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueBoolean", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueBoolean, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCanonical(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCanonical", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCanonical", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCanonical, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCode(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCode", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDate(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DateInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDate", nil, htmlAttrs)
	}
	return DateInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDate", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDate, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDateTime(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return DateTimeInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDateTime", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDateTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueDecimal(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return Float64Input("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueDecimal", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueDecimal, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueId(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueId", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueId", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueInstant(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueInstant", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueInstant", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueInstant, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueInteger(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueInteger", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueInteger", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueInteger, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueMarkdown(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueMarkdown", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueMarkdown, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueOid(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueOid", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueOid", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueOid, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValuePositiveInt(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValuePositiveInt", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValuePositiveInt, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueString(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueString", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueString", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueString, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueTime(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueTime", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueTime", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueTime, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUnsignedInt(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUnsignedInt", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUnsignedInt, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUri(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUri", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUri", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUri, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUrl(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUrl", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUrl", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUrl, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueUuid(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUuid", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueUuid", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueUuid, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueAnnotation(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return AnnotationTextArea("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueAnnotation", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueAnnotation, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCodeableConcept(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeableConceptSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCodeableConcept", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceDefaultValueCoding(numGroup int, numRule int, numSource int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodingSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].DefaultValueCoding", resource.Group[numGroup].Rule[numRule].Source[numSource].DefaultValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceElement(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Element", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Element", resource.Group[numGroup].Rule[numRule].Source[numSource].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceListMode(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_source_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].ListMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].ListMode", resource.Group[numGroup].Rule[numRule].Source[numSource].ListMode, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceVariable(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Variable", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Variable", resource.Group[numGroup].Rule[numRule].Source[numSource].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCondition(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Condition", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Condition", resource.Group[numGroup].Rule[numRule].Source[numSource].Condition, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceCheck(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Check", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].Check", resource.Group[numGroup].Rule[numRule].Source[numSource].Check, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleSourceLogMessage(numGroup int, numRule int, numSource int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numSource >= len(resource.Group[numGroup].Rule[numRule].Source) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].LogMessage", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Source["+strconv.Itoa(numSource)+"].LogMessage", resource.Group[numGroup].Rule[numRule].Source[numSource].LogMessage, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetContext(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Context", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Context", resource.Group[numGroup].Rule[numRule].Target[numTarget].Context, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetContextType(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_context_type

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ContextType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ContextType", resource.Group[numGroup].Rule[numRule].Target[numTarget].ContextType, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetElement(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Element", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Element", resource.Group[numGroup].Rule[numRule].Target[numTarget].Element, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetVariable(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Variable", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Variable", resource.Group[numGroup].Rule[numRule].Target[numTarget].Variable, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListMode(numGroup int, numRule int, numTarget int, numListMode int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_target_list_mode

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numListMode >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListMode["+strconv.Itoa(numListMode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListMode["+strconv.Itoa(numListMode)+"]", &resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode[numListMode], optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetListRuleId(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListRuleId", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].ListRuleId", resource.Group[numGroup].Rule[numRule].Target[numTarget].ListRuleId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetTransform(numGroup int, numRule int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMap_transform

	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) {
		return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Transform", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Transform", resource.Group[numGroup].Rule[numRule].Target[numTarget].Transform, optionsValueSet, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueId(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueId", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueId", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueString(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueString", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueBoolean(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return BoolInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueBoolean", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueInteger(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueInteger", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleTargetParameterValueDecimal(numGroup int, numRule int, numTarget int, numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numTarget >= len(resource.Group[numGroup].Rule[numRule].Target) || numParameter >= len(resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter) {
		return Float64Input("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Target["+strconv.Itoa(numTarget)+"].Parameter["+strconv.Itoa(numParameter)+"].ValueDecimal", &resource.Group[numGroup].Rule[numRule].Target[numTarget].Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDependentName(numGroup int, numRule int, numDependent int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Name", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Name", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Name, htmlAttrs)
}
func (resource *StructureMap) T_GroupRuleDependentVariable(numGroup int, numRule int, numDependent int, numVariable int, htmlAttrs string) templ.Component {
	if resource == nil || numGroup >= len(resource.Group) || numRule >= len(resource.Group[numGroup].Rule) || numDependent >= len(resource.Group[numGroup].Rule[numRule].Dependent) || numVariable >= len(resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Variable) {
		return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Variable["+strconv.Itoa(numVariable)+"]", nil, htmlAttrs)
	}
	return StringInput("StructureMap.Group["+strconv.Itoa(numGroup)+"].Rule["+strconv.Itoa(numRule)+"].Dependent["+strconv.Itoa(numDependent)+"].Variable["+strconv.Itoa(numVariable)+"]", &resource.Group[numGroup].Rule[numRule].Dependent[numDependent].Variable[numVariable], htmlAttrs)
}
