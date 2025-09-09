package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Parameters
type Parameters struct {
	Id            *string               `json:"id,omitempty"`
	Meta          *Meta                 `json:"meta,omitempty"`
	ImplicitRules *string               `json:"implicitRules,omitempty"`
	Language      *string               `json:"language,omitempty"`
	Parameter     []ParametersParameter `json:"parameter,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Parameters
type ParametersParameter struct {
	Id                       *string              `json:"id,omitempty"`
	Extension                []Extension          `json:"extension,omitempty"`
	ModifierExtension        []Extension          `json:"modifierExtension,omitempty"`
	Name                     string               `json:"name"`
	ValueBase64Binary        *string              `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                `json:"valueBoolean,omitempty"`
	ValueCanonical           *string              `json:"valueCanonical,omitempty"`
	ValueCode                *string              `json:"valueCode,omitempty"`
	ValueDate                *string              `json:"valueDate,omitempty"`
	ValueDateTime            *string              `json:"valueDateTime,omitempty"`
	ValueDecimal             *float64             `json:"valueDecimal,omitempty"`
	ValueId                  *string              `json:"valueId,omitempty"`
	ValueInstant             *string              `json:"valueInstant,omitempty"`
	ValueInteger             *int                 `json:"valueInteger,omitempty"`
	ValueMarkdown            *string              `json:"valueMarkdown,omitempty"`
	ValueOid                 *string              `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                 `json:"valuePositiveInt,omitempty"`
	ValueString              *string              `json:"valueString,omitempty"`
	ValueTime                *string              `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                 `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string              `json:"valueUri,omitempty"`
	ValueUrl                 *string              `json:"valueUrl,omitempty"`
	ValueUuid                *string              `json:"valueUuid,omitempty"`
	ValueAddress             *Address             `json:"valueAddress,omitempty"`
	ValueAge                 *Age                 `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation          `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment          `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint        `json:"valueContactPoint,omitempty"`
	ValueCount               *Count               `json:"valueCount,omitempty"`
	ValueDistance            *Distance            `json:"valueDistance,omitempty"`
	ValueDuration            *Duration            `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName           `json:"valueHumanName,omitempty"`
	ValueIdentifier          *Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money               `json:"valueMoney,omitempty"`
	ValuePeriod              *Period              `json:"valuePeriod,omitempty"`
	ValueQuantity            *Quantity            `json:"valueQuantity,omitempty"`
	ValueRange               *Range               `json:"valueRange,omitempty"`
	ValueRatio               *Ratio               `json:"valueRatio,omitempty"`
	ValueReference           *Reference           `json:"valueReference,omitempty"`
	ValueSampledData         *SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature           `json:"valueSignature,omitempty"`
	ValueTiming              *Timing              `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail       `json:"valueContactDetail,omitempty"`
	ValueContributor         *Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement     `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression          `json:"valueExpression,omitempty"`
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact     `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition   `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext        `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage              `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                `json:"valueMeta,omitempty"`
	Resource                 *Resource            `json:"resource,omitempty"`
}

func (resource *Parameters) T_ParameterName(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", &resource.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBase64Binary(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueBase64Binary", resource.Parameter[numParameter].ValueBase64Binary, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBoolean(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCanonical(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueCanonical", resource.Parameter[numParameter].ValueCanonical, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].valueCode", resource.Parameter[numParameter].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDate(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DateInput("parameter["+strconv.Itoa(numParameter)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("parameter["+strconv.Itoa(numParameter)+"].valueDate", resource.Parameter[numParameter].ValueDate, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDateTime(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DateTimeInput("parameter["+strconv.Itoa(numParameter)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("parameter["+strconv.Itoa(numParameter)+"].valueDateTime", resource.Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDecimal(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return Float64Input("parameter["+strconv.Itoa(numParameter)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("parameter["+strconv.Itoa(numParameter)+"].valueDecimal", resource.Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueId(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueId", resource.Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInstant(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueInstant", resource.Parameter[numParameter].ValueInstant, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInteger(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueInteger", resource.Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueMarkdown(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueMarkdown", resource.Parameter[numParameter].ValueMarkdown, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueOid(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueOid", resource.Parameter[numParameter].ValueOid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValuePositiveInt(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valuePositiveInt", resource.Parameter[numParameter].ValuePositiveInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueString(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueString", resource.Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueTime(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueTime", resource.Parameter[numParameter].ValueTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUnsignedInt(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueUnsignedInt", resource.Parameter[numParameter].ValueUnsignedInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUri(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUri", resource.Parameter[numParameter].ValueUri, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUrl(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUrl", resource.Parameter[numParameter].ValueUrl, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUuid(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUuid", resource.Parameter[numParameter].ValueUuid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAnnotation(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AnnotationTextArea("parameter["+strconv.Itoa(numParameter)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("parameter["+strconv.Itoa(numParameter)+"].valueAnnotation", resource.Parameter[numParameter].ValueAnnotation, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCoding(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodingSelect("parameter["+strconv.Itoa(numParameter)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("parameter["+strconv.Itoa(numParameter)+"].valueCoding", resource.Parameter[numParameter].ValueCoding, optionsValueSet, htmlAttrs)
}
