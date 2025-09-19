package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"strconv"

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
	ValueDate                  *FhirDate              `json:"valueDate,omitempty"`
	ValueDateTime              *FhirDateTime          `json:"valueDateTime,omitempty"`
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

func (resource *Parameters) T_ParameterName(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", &resource.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBase64Binary(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueBase64Binary", resource.Parameter[numParameter].ValueBase64Binary, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueBoolean(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCanonical(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueCanonical", resource.Parameter[numParameter].ValueCanonical, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCode(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].valueCode", resource.Parameter[numParameter].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDate(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return FhirDateInput("parameter["+strconv.Itoa(numParameter)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("parameter["+strconv.Itoa(numParameter)+"].valueDate", resource.Parameter[numParameter].ValueDate, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDateTime(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return FhirDateTimeInput("parameter["+strconv.Itoa(numParameter)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("parameter["+strconv.Itoa(numParameter)+"].valueDateTime", resource.Parameter[numParameter].ValueDateTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDecimal(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return Float64Input("parameter["+strconv.Itoa(numParameter)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("parameter["+strconv.Itoa(numParameter)+"].valueDecimal", resource.Parameter[numParameter].ValueDecimal, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueId(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueId", resource.Parameter[numParameter].ValueId, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInstant(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueInstant", resource.Parameter[numParameter].ValueInstant, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInteger(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueInteger", resource.Parameter[numParameter].ValueInteger, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueInteger64(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return Int64Input("parameter["+strconv.Itoa(numParameter)+"].valueInteger64", nil, htmlAttrs)
	}
	return Int64Input("parameter["+strconv.Itoa(numParameter)+"].valueInteger64", resource.Parameter[numParameter].ValueInteger64, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueMarkdown(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueMarkdown", resource.Parameter[numParameter].ValueMarkdown, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueOid(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueOid", resource.Parameter[numParameter].ValueOid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValuePositiveInt(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valuePositiveInt", resource.Parameter[numParameter].ValuePositiveInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueString(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueString", resource.Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueTime(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueTime", resource.Parameter[numParameter].ValueTime, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUnsignedInt(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].valueUnsignedInt", resource.Parameter[numParameter].ValueUnsignedInt, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUri(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUri", resource.Parameter[numParameter].ValueUri, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUrl(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUrl", resource.Parameter[numParameter].ValueUrl, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUuid(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].valueUuid", resource.Parameter[numParameter].ValueUuid, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAddress(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AddressInput("parameter["+strconv.Itoa(numParameter)+"].valueAddress", nil, htmlAttrs)
	}
	return AddressInput("parameter["+strconv.Itoa(numParameter)+"].valueAddress", resource.Parameter[numParameter].ValueAddress, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAge(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AgeInput("parameter["+strconv.Itoa(numParameter)+"].valueAge", nil, htmlAttrs)
	}
	return AgeInput("parameter["+strconv.Itoa(numParameter)+"].valueAge", resource.Parameter[numParameter].ValueAge, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAnnotation(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AnnotationTextArea("parameter["+strconv.Itoa(numParameter)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("parameter["+strconv.Itoa(numParameter)+"].valueAnnotation", resource.Parameter[numParameter].ValueAnnotation, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAttachment(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AttachmentInput("parameter["+strconv.Itoa(numParameter)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("parameter["+strconv.Itoa(numParameter)+"].valueAttachment", resource.Parameter[numParameter].ValueAttachment, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCodeableReference(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableReferenceInput("parameter["+strconv.Itoa(numParameter)+"].valueCodeableReference", nil, htmlAttrs)
	}
	return CodeableReferenceInput("parameter["+strconv.Itoa(numParameter)+"].valueCodeableReference", resource.Parameter[numParameter].ValueCodeableReference, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCoding(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodingSelect("parameter["+strconv.Itoa(numParameter)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("parameter["+strconv.Itoa(numParameter)+"].valueCoding", resource.Parameter[numParameter].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueContactPoint(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ContactPointInput("parameter["+strconv.Itoa(numParameter)+"].valueContactPoint", nil, htmlAttrs)
	}
	return ContactPointInput("parameter["+strconv.Itoa(numParameter)+"].valueContactPoint", resource.Parameter[numParameter].ValueContactPoint, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueCount(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CountInput("parameter["+strconv.Itoa(numParameter)+"].valueCount", nil, htmlAttrs)
	}
	return CountInput("parameter["+strconv.Itoa(numParameter)+"].valueCount", resource.Parameter[numParameter].ValueCount, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDistance(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DistanceInput("parameter["+strconv.Itoa(numParameter)+"].valueDistance", nil, htmlAttrs)
	}
	return DistanceInput("parameter["+strconv.Itoa(numParameter)+"].valueDistance", resource.Parameter[numParameter].ValueDistance, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDuration(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DurationInput("parameter["+strconv.Itoa(numParameter)+"].valueDuration", nil, htmlAttrs)
	}
	return DurationInput("parameter["+strconv.Itoa(numParameter)+"].valueDuration", resource.Parameter[numParameter].ValueDuration, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueHumanName(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return HumanNameInput("parameter["+strconv.Itoa(numParameter)+"].valueHumanName", nil, htmlAttrs)
	}
	return HumanNameInput("parameter["+strconv.Itoa(numParameter)+"].valueHumanName", resource.Parameter[numParameter].ValueHumanName, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueIdentifier(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IdentifierInput("parameter["+strconv.Itoa(numParameter)+"].valueIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("parameter["+strconv.Itoa(numParameter)+"].valueIdentifier", resource.Parameter[numParameter].ValueIdentifier, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueMoney(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return MoneyInput("parameter["+strconv.Itoa(numParameter)+"].valueMoney", nil, htmlAttrs)
	}
	return MoneyInput("parameter["+strconv.Itoa(numParameter)+"].valueMoney", resource.Parameter[numParameter].ValueMoney, htmlAttrs)
}
func (resource *Parameters) T_ParameterValuePeriod(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return PeriodInput("parameter["+strconv.Itoa(numParameter)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("parameter["+strconv.Itoa(numParameter)+"].valuePeriod", resource.Parameter[numParameter].ValuePeriod, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueQuantity(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", resource.Parameter[numParameter].ValueQuantity, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueRange(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", resource.Parameter[numParameter].ValueRange, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueRatio(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RatioInput("parameter["+strconv.Itoa(numParameter)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("parameter["+strconv.Itoa(numParameter)+"].valueRatio", resource.Parameter[numParameter].ValueRatio, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueRatioRange(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RatioRangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRatioRange", resource.Parameter[numParameter].ValueRatioRange, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueReference(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ReferenceInput("parameter["+strconv.Itoa(numParameter)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput("parameter["+strconv.Itoa(numParameter)+"].valueReference", resource.Parameter[numParameter].ValueReference, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueSampledData(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return SampledDataInput("parameter["+strconv.Itoa(numParameter)+"].valueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("parameter["+strconv.Itoa(numParameter)+"].valueSampledData", resource.Parameter[numParameter].ValueSampledData, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueSignature(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return SignatureInput("parameter["+strconv.Itoa(numParameter)+"].valueSignature", nil, htmlAttrs)
	}
	return SignatureInput("parameter["+strconv.Itoa(numParameter)+"].valueSignature", resource.Parameter[numParameter].ValueSignature, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueTiming(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return TimingInput("parameter["+strconv.Itoa(numParameter)+"].valueTiming", nil, htmlAttrs)
	}
	return TimingInput("parameter["+strconv.Itoa(numParameter)+"].valueTiming", resource.Parameter[numParameter].ValueTiming, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueContactDetail(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ContactDetailInput("parameter["+strconv.Itoa(numParameter)+"].valueContactDetail", nil, htmlAttrs)
	}
	return ContactDetailInput("parameter["+strconv.Itoa(numParameter)+"].valueContactDetail", resource.Parameter[numParameter].ValueContactDetail, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDataRequirement(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DataRequirementInput("parameter["+strconv.Itoa(numParameter)+"].valueDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("parameter["+strconv.Itoa(numParameter)+"].valueDataRequirement", resource.Parameter[numParameter].ValueDataRequirement, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueExpression(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ExpressionInput("parameter["+strconv.Itoa(numParameter)+"].valueExpression", nil, htmlAttrs)
	}
	return ExpressionInput("parameter["+strconv.Itoa(numParameter)+"].valueExpression", resource.Parameter[numParameter].ValueExpression, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueParameterDefinition(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ParameterDefinitionInput("parameter["+strconv.Itoa(numParameter)+"].valueParameterDefinition", nil, htmlAttrs)
	}
	return ParameterDefinitionInput("parameter["+strconv.Itoa(numParameter)+"].valueParameterDefinition", resource.Parameter[numParameter].ValueParameterDefinition, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueRelatedArtifact(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RelatedArtifactInput("parameter["+strconv.Itoa(numParameter)+"].valueRelatedArtifact", nil, htmlAttrs)
	}
	return RelatedArtifactInput("parameter["+strconv.Itoa(numParameter)+"].valueRelatedArtifact", resource.Parameter[numParameter].ValueRelatedArtifact, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueTriggerDefinition(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return TriggerDefinitionInput("parameter["+strconv.Itoa(numParameter)+"].valueTriggerDefinition", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("parameter["+strconv.Itoa(numParameter)+"].valueTriggerDefinition", resource.Parameter[numParameter].ValueTriggerDefinition, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueUsageContext(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return UsageContextInput("parameter["+strconv.Itoa(numParameter)+"].valueUsageContext", nil, htmlAttrs)
	}
	return UsageContextInput("parameter["+strconv.Itoa(numParameter)+"].valueUsageContext", resource.Parameter[numParameter].ValueUsageContext, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueAvailability(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return AvailabilityInput("parameter["+strconv.Itoa(numParameter)+"].valueAvailability", nil, htmlAttrs)
	}
	return AvailabilityInput("parameter["+strconv.Itoa(numParameter)+"].valueAvailability", resource.Parameter[numParameter].ValueAvailability, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueExtendedContactDetail(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return ExtendedContactDetailInput("parameter["+strconv.Itoa(numParameter)+"].valueExtendedContactDetail", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("parameter["+strconv.Itoa(numParameter)+"].valueExtendedContactDetail", resource.Parameter[numParameter].ValueExtendedContactDetail, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueDosage(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return DosageInput("parameter["+strconv.Itoa(numParameter)+"].valueDosage", nil, htmlAttrs)
	}
	return DosageInput("parameter["+strconv.Itoa(numParameter)+"].valueDosage", resource.Parameter[numParameter].ValueDosage, htmlAttrs)
}
func (resource *Parameters) T_ParameterValueMeta(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return MetaInput("parameter["+strconv.Itoa(numParameter)+"].valueMeta", nil, htmlAttrs)
	}
	return MetaInput("parameter["+strconv.Itoa(numParameter)+"].valueMeta", resource.Parameter[numParameter].ValueMeta, htmlAttrs)
}
