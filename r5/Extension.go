package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Extension
type Extension struct {
	Id                         *string                `json:"id,omitempty"`
	Extension                  []Extension            `json:"extension,omitempty"`
	Url                        string                 `json:"url"`
	ValueBase64Binary          *string                `json:"valueBase64Binary"`
	ValueBoolean               *bool                  `json:"valueBoolean"`
	ValueCanonical             *string                `json:"valueCanonical"`
	ValueCode                  *string                `json:"valueCode"`
	ValueDate                  *string                `json:"valueDate"`
	ValueDateTime              *string                `json:"valueDateTime"`
	ValueDecimal               *float64               `json:"valueDecimal"`
	ValueId                    *string                `json:"valueId"`
	ValueInstant               *string                `json:"valueInstant"`
	ValueInteger               *int                   `json:"valueInteger"`
	ValueInteger64             *int64                 `json:"valueInteger64"`
	ValueMarkdown              *string                `json:"valueMarkdown"`
	ValueOid                   *string                `json:"valueOid"`
	ValuePositiveInt           *int                   `json:"valuePositiveInt"`
	ValueString                *string                `json:"valueString"`
	ValueTime                  *string                `json:"valueTime"`
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt"`
	ValueUri                   *string                `json:"valueUri"`
	ValueUrl                   *string                `json:"valueUrl"`
	ValueUuid                  *string                `json:"valueUuid"`
	ValueAddress               *Address               `json:"valueAddress"`
	ValueAge                   *Age                   `json:"valueAge"`
	ValueAnnotation            *Annotation            `json:"valueAnnotation"`
	ValueAttachment            *Attachment            `json:"valueAttachment"`
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept"`
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference"`
	ValueCoding                *Coding                `json:"valueCoding"`
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint"`
	ValueCount                 *Count                 `json:"valueCount"`
	ValueDistance              *Distance              `json:"valueDistance"`
	ValueDuration              *Duration              `json:"valueDuration"`
	ValueHumanName             *HumanName             `json:"valueHumanName"`
	ValueIdentifier            *Identifier            `json:"valueIdentifier"`
	ValueMoney                 *Money                 `json:"valueMoney"`
	ValuePeriod                *Period                `json:"valuePeriod"`
	ValueQuantity              *Quantity              `json:"valueQuantity"`
	ValueRange                 *Range                 `json:"valueRange"`
	ValueRatio                 *Ratio                 `json:"valueRatio"`
	ValueRatioRange            *RatioRange            `json:"valueRatioRange"`
	ValueReference             *Reference             `json:"valueReference"`
	ValueSampledData           *SampledData           `json:"valueSampledData"`
	ValueSignature             *Signature             `json:"valueSignature"`
	ValueTiming                *Timing                `json:"valueTiming"`
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail"`
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement"`
	ValueExpression            *Expression            `json:"valueExpression"`
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition"`
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact"`
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition"`
	ValueUsageContext          *UsageContext          `json:"valueUsageContext"`
	ValueAvailability          *Availability          `json:"valueAvailability"`
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail"`
	ValueDosage                *Dosage                `json:"valueDosage"`
	ValueMeta                  *Meta                  `json:"valueMeta"`
}
