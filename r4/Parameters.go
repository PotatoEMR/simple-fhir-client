package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "github.com/a-h/templ"

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
	ValueBase64Binary        *string              `json:"valueBase64Binary"`
	ValueBoolean             *bool                `json:"valueBoolean"`
	ValueCanonical           *string              `json:"valueCanonical"`
	ValueCode                *string              `json:"valueCode"`
	ValueDate                *string              `json:"valueDate"`
	ValueDateTime            *string              `json:"valueDateTime"`
	ValueDecimal             *float64             `json:"valueDecimal"`
	ValueId                  *string              `json:"valueId"`
	ValueInstant             *string              `json:"valueInstant"`
	ValueInteger             *int                 `json:"valueInteger"`
	ValueMarkdown            *string              `json:"valueMarkdown"`
	ValueOid                 *string              `json:"valueOid"`
	ValuePositiveInt         *int                 `json:"valuePositiveInt"`
	ValueString              *string              `json:"valueString"`
	ValueTime                *string              `json:"valueTime"`
	ValueUnsignedInt         *int                 `json:"valueUnsignedInt"`
	ValueUri                 *string              `json:"valueUri"`
	ValueUrl                 *string              `json:"valueUrl"`
	ValueUuid                *string              `json:"valueUuid"`
	ValueAddress             *Address             `json:"valueAddress"`
	ValueAge                 *Age                 `json:"valueAge"`
	ValueAnnotation          *Annotation          `json:"valueAnnotation"`
	ValueAttachment          *Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     *CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              *Coding              `json:"valueCoding"`
	ValueContactPoint        *ContactPoint        `json:"valueContactPoint"`
	ValueCount               *Count               `json:"valueCount"`
	ValueDistance            *Distance            `json:"valueDistance"`
	ValueDuration            *Duration            `json:"valueDuration"`
	ValueHumanName           *HumanName           `json:"valueHumanName"`
	ValueIdentifier          *Identifier          `json:"valueIdentifier"`
	ValueMoney               *Money               `json:"valueMoney"`
	ValuePeriod              *Period              `json:"valuePeriod"`
	ValueQuantity            *Quantity            `json:"valueQuantity"`
	ValueRange               *Range               `json:"valueRange"`
	ValueRatio               *Ratio               `json:"valueRatio"`
	ValueReference           *Reference           `json:"valueReference"`
	ValueSampledData         *SampledData         `json:"valueSampledData"`
	ValueSignature           *Signature           `json:"valueSignature"`
	ValueTiming              *Timing              `json:"valueTiming"`
	ValueContactDetail       *ContactDetail       `json:"valueContactDetail"`
	ValueContributor         *Contributor         `json:"valueContributor"`
	ValueDataRequirement     *DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          *Expression          `json:"valueExpression"`
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     *RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   *TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        *UsageContext        `json:"valueUsageContext"`
	ValueDosage              *Dosage              `json:"valueDosage"`
	ValueMeta                *Meta                `json:"valueMeta"`
	Resource                 *Resource            `json:"resource,omitempty"`
}

func (resource *Parameters) ParametersLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
