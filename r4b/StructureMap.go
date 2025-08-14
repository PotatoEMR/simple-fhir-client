//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

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
	Date              *string                 `json:"date,omitempty"`
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
	DefaultValueBase64Binary        *string              `json:"defaultValueBase64Binary"`
	DefaultValueBoolean             *bool                `json:"defaultValueBoolean"`
	DefaultValueCanonical           *string              `json:"defaultValueCanonical"`
	DefaultValueCode                *string              `json:"defaultValueCode"`
	DefaultValueDate                *string              `json:"defaultValueDate"`
	DefaultValueDateTime            *string              `json:"defaultValueDateTime"`
	DefaultValueDecimal             *float64             `json:"defaultValueDecimal"`
	DefaultValueId                  *string              `json:"defaultValueId"`
	DefaultValueInstant             *string              `json:"defaultValueInstant"`
	DefaultValueInteger             *int                 `json:"defaultValueInteger"`
	DefaultValueMarkdown            *string              `json:"defaultValueMarkdown"`
	DefaultValueOid                 *string              `json:"defaultValueOid"`
	DefaultValuePositiveInt         *int                 `json:"defaultValuePositiveInt"`
	DefaultValueString              *string              `json:"defaultValueString"`
	DefaultValueTime                *string              `json:"defaultValueTime"`
	DefaultValueUnsignedInt         *int                 `json:"defaultValueUnsignedInt"`
	DefaultValueUri                 *string              `json:"defaultValueUri"`
	DefaultValueUrl                 *string              `json:"defaultValueUrl"`
	DefaultValueUuid                *string              `json:"defaultValueUuid"`
	DefaultValueAddress             *Address             `json:"defaultValueAddress"`
	DefaultValueAge                 *Age                 `json:"defaultValueAge"`
	DefaultValueAnnotation          *Annotation          `json:"defaultValueAnnotation"`
	DefaultValueAttachment          *Attachment          `json:"defaultValueAttachment"`
	DefaultValueCodeableConcept     *CodeableConcept     `json:"defaultValueCodeableConcept"`
	DefaultValueCoding              *Coding              `json:"defaultValueCoding"`
	DefaultValueContactPoint        *ContactPoint        `json:"defaultValueContactPoint"`
	DefaultValueCount               *Count               `json:"defaultValueCount"`
	DefaultValueDistance            *Distance            `json:"defaultValueDistance"`
	DefaultValueDuration            *Duration            `json:"defaultValueDuration"`
	DefaultValueHumanName           *HumanName           `json:"defaultValueHumanName"`
	DefaultValueIdentifier          *Identifier          `json:"defaultValueIdentifier"`
	DefaultValueMoney               *Money               `json:"defaultValueMoney"`
	DefaultValuePeriod              *Period              `json:"defaultValuePeriod"`
	DefaultValueQuantity            *Quantity            `json:"defaultValueQuantity"`
	DefaultValueRange               *Range               `json:"defaultValueRange"`
	DefaultValueRatio               *Ratio               `json:"defaultValueRatio"`
	DefaultValueReference           *Reference           `json:"defaultValueReference"`
	DefaultValueSampledData         *SampledData         `json:"defaultValueSampledData"`
	DefaultValueSignature           *Signature           `json:"defaultValueSignature"`
	DefaultValueTiming              *Timing              `json:"defaultValueTiming"`
	DefaultValueContactDetail       *ContactDetail       `json:"defaultValueContactDetail"`
	DefaultValueContributor         *Contributor         `json:"defaultValueContributor"`
	DefaultValueDataRequirement     *DataRequirement     `json:"defaultValueDataRequirement"`
	DefaultValueExpression          *Expression          `json:"defaultValueExpression"`
	DefaultValueParameterDefinition *ParameterDefinition `json:"defaultValueParameterDefinition"`
	DefaultValueRelatedArtifact     *RelatedArtifact     `json:"defaultValueRelatedArtifact"`
	DefaultValueTriggerDefinition   *TriggerDefinition   `json:"defaultValueTriggerDefinition"`
	DefaultValueUsageContext        *UsageContext        `json:"defaultValueUsageContext"`
	DefaultValueDosage              *Dosage              `json:"defaultValueDosage"`
	DefaultValueMeta                *Meta                `json:"defaultValueMeta"`
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
