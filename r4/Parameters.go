package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Parameters) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Parameters.Id", nil)
	}
	return StringInput("Parameters.Id", resource.Id)
}
func (resource *Parameters) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Parameters.ImplicitRules", nil)
	}
	return StringInput("Parameters.ImplicitRules", resource.ImplicitRules)
}
func (resource *Parameters) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Parameters.Language", nil, optionsValueSet)
	}
	return CodeSelect("Parameters.Language", resource.Language, optionsValueSet)
}
func (resource *Parameters) T_ParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("Parameters.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("Parameters.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Parameter[numParameter].Id)
}
func (resource *Parameters) T_ParameterName(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("Parameters.Parameter["+strconv.Itoa(numParameter)+"].Name", nil)
	}
	return StringInput("Parameters.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Parameter[numParameter].Name)
}
