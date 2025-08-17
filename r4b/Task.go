//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/Task
type Task struct {
	Id                    *string           `json:"id,omitempty"`
	Meta                  *Meta             `json:"meta,omitempty"`
	ImplicitRules         *string           `json:"implicitRules,omitempty"`
	Language              *string           `json:"language,omitempty"`
	Text                  *Narrative        `json:"text,omitempty"`
	Contained             []Resource        `json:"contained,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Identifier            []Identifier      `json:"identifier,omitempty"`
	InstantiatesCanonical *string           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string           `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference       `json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier       `json:"groupIdentifier,omitempty"`
	PartOf                []Reference       `json:"partOf,omitempty"`
	Status                string            `json:"status"`
	StatusReason          *CodeableConcept  `json:"statusReason,omitempty"`
	BusinessStatus        *CodeableConcept  `json:"businessStatus,omitempty"`
	Intent                string            `json:"intent"`
	Priority              *string           `json:"priority,omitempty"`
	Code                  *CodeableConcept  `json:"code,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Focus                 *Reference        `json:"focus,omitempty"`
	For                   *Reference        `json:"for,omitempty"`
	Encounter             *Reference        `json:"encounter,omitempty"`
	ExecutionPeriod       *Period           `json:"executionPeriod,omitempty"`
	AuthoredOn            *string           `json:"authoredOn,omitempty"`
	LastModified          *string           `json:"lastModified,omitempty"`
	Requester             *Reference        `json:"requester,omitempty"`
	PerformerType         []CodeableConcept `json:"performerType,omitempty"`
	Owner                 *Reference        `json:"owner,omitempty"`
	Location              *Reference        `json:"location,omitempty"`
	ReasonCode            *CodeableConcept  `json:"reasonCode,omitempty"`
	ReasonReference       *Reference        `json:"reasonReference,omitempty"`
	Insurance             []Reference       `json:"insurance,omitempty"`
	Note                  []Annotation      `json:"note,omitempty"`
	RelevantHistory       []Reference       `json:"relevantHistory,omitempty"`
	Restriction           *TaskRestriction  `json:"restriction,omitempty"`
	Input                 []TaskInput       `json:"input,omitempty"`
	Output                []TaskOutput      `json:"output,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Task
type TaskRestriction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Repetitions       *int        `json:"repetitions,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Recipient         []Reference `json:"recipient,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Task
type TaskInput struct {
	Id                       *string             `json:"id,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Type                     CodeableConcept     `json:"type"`
	ValueBase64Binary        string              `json:"valueBase64Binary"`
	ValueBoolean             bool                `json:"valueBoolean"`
	ValueCanonical           string              `json:"valueCanonical"`
	ValueCode                string              `json:"valueCode"`
	ValueDate                string              `json:"valueDate"`
	ValueDateTime            string              `json:"valueDateTime"`
	ValueDecimal             float64             `json:"valueDecimal"`
	ValueId                  string              `json:"valueId"`
	ValueInstant             string              `json:"valueInstant"`
	ValueInteger             int                 `json:"valueInteger"`
	ValueMarkdown            string              `json:"valueMarkdown"`
	ValueOid                 string              `json:"valueOid"`
	ValuePositiveInt         int                 `json:"valuePositiveInt"`
	ValueString              string              `json:"valueString"`
	ValueTime                string              `json:"valueTime"`
	ValueUnsignedInt         int                 `json:"valueUnsignedInt"`
	ValueUri                 string              `json:"valueUri"`
	ValueUrl                 string              `json:"valueUrl"`
	ValueUuid                string              `json:"valueUuid"`
	ValueAddress             Address             `json:"valueAddress"`
	ValueAge                 Age                 `json:"valueAge"`
	ValueAnnotation          Annotation          `json:"valueAnnotation"`
	ValueAttachment          Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              Coding              `json:"valueCoding"`
	ValueContactPoint        ContactPoint        `json:"valueContactPoint"`
	ValueCount               Count               `json:"valueCount"`
	ValueDistance            Distance            `json:"valueDistance"`
	ValueDuration            Duration            `json:"valueDuration"`
	ValueHumanName           HumanName           `json:"valueHumanName"`
	ValueIdentifier          Identifier          `json:"valueIdentifier"`
	ValueMoney               Money               `json:"valueMoney"`
	ValuePeriod              Period              `json:"valuePeriod"`
	ValueQuantity            Quantity            `json:"valueQuantity"`
	ValueRange               Range               `json:"valueRange"`
	ValueRatio               Ratio               `json:"valueRatio"`
	ValueReference           Reference           `json:"valueReference"`
	ValueSampledData         SampledData         `json:"valueSampledData"`
	ValueSignature           Signature           `json:"valueSignature"`
	ValueTiming              Timing              `json:"valueTiming"`
	ValueContactDetail       ContactDetail       `json:"valueContactDetail"`
	ValueContributor         Contributor         `json:"valueContributor"`
	ValueDataRequirement     DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          Expression          `json:"valueExpression"`
	ValueParameterDefinition ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        UsageContext        `json:"valueUsageContext"`
	ValueDosage              Dosage              `json:"valueDosage"`
	ValueMeta                Meta                `json:"valueMeta"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Task
type TaskOutput struct {
	Id                       *string             `json:"id,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Type                     CodeableConcept     `json:"type"`
	ValueBase64Binary        string              `json:"valueBase64Binary"`
	ValueBoolean             bool                `json:"valueBoolean"`
	ValueCanonical           string              `json:"valueCanonical"`
	ValueCode                string              `json:"valueCode"`
	ValueDate                string              `json:"valueDate"`
	ValueDateTime            string              `json:"valueDateTime"`
	ValueDecimal             float64             `json:"valueDecimal"`
	ValueId                  string              `json:"valueId"`
	ValueInstant             string              `json:"valueInstant"`
	ValueInteger             int                 `json:"valueInteger"`
	ValueMarkdown            string              `json:"valueMarkdown"`
	ValueOid                 string              `json:"valueOid"`
	ValuePositiveInt         int                 `json:"valuePositiveInt"`
	ValueString              string              `json:"valueString"`
	ValueTime                string              `json:"valueTime"`
	ValueUnsignedInt         int                 `json:"valueUnsignedInt"`
	ValueUri                 string              `json:"valueUri"`
	ValueUrl                 string              `json:"valueUrl"`
	ValueUuid                string              `json:"valueUuid"`
	ValueAddress             Address             `json:"valueAddress"`
	ValueAge                 Age                 `json:"valueAge"`
	ValueAnnotation          Annotation          `json:"valueAnnotation"`
	ValueAttachment          Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              Coding              `json:"valueCoding"`
	ValueContactPoint        ContactPoint        `json:"valueContactPoint"`
	ValueCount               Count               `json:"valueCount"`
	ValueDistance            Distance            `json:"valueDistance"`
	ValueDuration            Duration            `json:"valueDuration"`
	ValueHumanName           HumanName           `json:"valueHumanName"`
	ValueIdentifier          Identifier          `json:"valueIdentifier"`
	ValueMoney               Money               `json:"valueMoney"`
	ValuePeriod              Period              `json:"valuePeriod"`
	ValueQuantity            Quantity            `json:"valueQuantity"`
	ValueRange               Range               `json:"valueRange"`
	ValueRatio               Ratio               `json:"valueRatio"`
	ValueReference           Reference           `json:"valueReference"`
	ValueSampledData         SampledData         `json:"valueSampledData"`
	ValueSignature           Signature           `json:"valueSignature"`
	ValueTiming              Timing              `json:"valueTiming"`
	ValueContactDetail       ContactDetail       `json:"valueContactDetail"`
	ValueContributor         Contributor         `json:"valueContributor"`
	ValueDataRequirement     DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          Expression          `json:"valueExpression"`
	ValueParameterDefinition ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        UsageContext        `json:"valueUsageContext"`
	ValueDosage              Dosage              `json:"valueDosage"`
	ValueMeta                Meta                `json:"valueMeta"`
}

type OtherTask Task

// on convert struct to json, automatically add resourceType=Task
func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTask
		ResourceType string `json:"resourceType"`
	}{
		OtherTask:    OtherTask(r),
		ResourceType: "Task",
	})
}
