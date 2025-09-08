package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Parameters
type Parameters struct {
	Id            *string               `json:"id,omitempty"`
	Meta          *Meta                 `json:"meta,omitempty"`
	ImplicitRules *string               `json:"implicitRules,omitempty"`
	Language      *string               `json:"language,omitempty"`
	Parameter     []ParametersParameter `json:"parameter,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Parameters
type ParametersParameter struct {
	Id                         *string                `json:"id,omitempty"`
	Extension                  []Extension            `json:"extension,omitempty"`
	ModifierExtension          []Extension            `json:"modifierExtension,omitempty"`
	Name                       string                 `json:"name"`
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty"`
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty"`
	ValueCanonical             *string                `json:"valueCanonical,omitempty"`
	ValueCode                  *string                `json:"valueCode,omitempty"`
	ValueDate                  *time.Time             `json:"valueDate,omitempty,format:'2006-01-02'"`
	ValueDateTime              *time.Time             `json:"valueDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ValueDecimal               *float64               `json:"valueDecimal,omitempty"`
	ValueId                    *string                `json:"valueId,omitempty"`
	ValueInstant               *string                `json:"valueInstant,omitempty"`
	ValueInteger               *int                   `json:"valueInteger,omitempty"`
	ValueInteger64             *int64                 `json:"valueInteger64,omitempty"`
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty"`
	ValueOid                   *string                `json:"valueOid,omitempty"`
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty"`
	ValueString                *string                `json:"valueString,omitempty"`
	ValueTime                  *string                `json:"valueTime,omitempty"`
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty"`
	ValueUri                   *string                `json:"valueUri,omitempty"`
	ValueUrl                   *string                `json:"valueUrl,omitempty"`
	ValueUuid                  *string                `json:"valueUuid,omitempty"`
	ValueAddress               *Address               `json:"valueAddress,omitempty"`
	ValueAge                   *Age                   `json:"valueAge,omitempty"`
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty"`
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty"`
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty"`
	ValueCoding                *Coding                `json:"valueCoding,omitempty"`
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty"`
	ValueCount                 *Count                 `json:"valueCount,omitempty"`
	ValueDistance              *Distance              `json:"valueDistance,omitempty"`
	ValueDuration              *Duration              `json:"valueDuration,omitempty"`
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty"`
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty"`
	ValueMoney                 *Money                 `json:"valueMoney,omitempty"`
	ValuePeriod                *Period                `json:"valuePeriod,omitempty"`
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty"`
	ValueRange                 *Range                 `json:"valueRange,omitempty"`
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty"`
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty"`
	ValueReference             *Reference             `json:"valueReference,omitempty"`
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty"`
	ValueSignature             *Signature             `json:"valueSignature,omitempty"`
	ValueTiming                *Timing                `json:"valueTiming,omitempty"`
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty"`
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty"`
	ValueExpression            *Expression            `json:"valueExpression,omitempty"`
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty"`
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                *Dosage                `json:"valueDosage,omitempty"`
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty"`
	Resource                   *Resource              `json:"resource,omitempty"`
}

func (resource *Parameters) T_ParameterName(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Name", &resource.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBase64Binary(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBase64Binary", resource.Parameter[numParameter].ValueBase64Binary, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBoolean(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCanonical(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueCanonical", resource.Parameter[numParameter].ValueCanonical, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCode", resource.Parameter[numParameter].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDate(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DateInput("Parameter["+strconv.Itoa(numParameter)+"]ValueDate", nil, htmlAttrs)
	}
	return DateInput("Parameter["+strconv.Itoa(numParameter)+"]ValueDate", resource.Parameter[numParameter].ValueDate, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDateTime(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DateTimeInput("Parameter["+strconv.Itoa(numParameter)+"]ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Parameter["+strconv.Itoa(numParameter)+"]ValueDateTime", resource.Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDecimal(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return Float64Input("Parameter["+strconv.Itoa(numParameter)+"]ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Parameter["+strconv.Itoa(numParameter)+"]ValueDecimal", resource.Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueId(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueId", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueId", resource.Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInstant(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueInstant", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueInstant", resource.Parameter[numParameter].ValueInstant, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInteger(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValueInteger", resource.Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInteger64(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return Int64Input("Parameter["+strconv.Itoa(numParameter)+"]ValueInteger64", nil, htmlAttrs)
	}
	return Int64Input("Parameter["+strconv.Itoa(numParameter)+"]ValueInteger64", resource.Parameter[numParameter].ValueInteger64, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueMarkdown(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueMarkdown", resource.Parameter[numParameter].ValueMarkdown, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueOid(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueOid", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueOid", resource.Parameter[numParameter].ValueOid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValuePositiveInt(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValuePositiveInt", resource.Parameter[numParameter].ValuePositiveInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueString(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueString", resource.Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueTime(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueTime", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueTime", resource.Parameter[numParameter].ValueTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUnsignedInt(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUnsignedInt", resource.Parameter[numParameter].ValueUnsignedInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUri(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUri", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUri", resource.Parameter[numParameter].ValueUri, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUrl(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUrl", resource.Parameter[numParameter].ValueUrl, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUuid(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUuid", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ValueUuid", resource.Parameter[numParameter].ValueUuid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAnnotation(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AnnotationTextArea("Parameter["+strconv.Itoa(numParameter)+"]ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Parameter["+strconv.Itoa(numParameter)+"]ValueAnnotation", resource.Parameter[numParameter].ValueAnnotation, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCoding(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodingSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCoding", resource.Parameter[numParameter].ValueCoding, optionsValueSet, htmlAttrs)
}
