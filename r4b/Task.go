package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Task) TaskLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Task) TaskStatus() templ.Component {
	optionsValueSet := VSTask_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Task) TaskStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Task) TaskBusinessStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("businessStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("businessStatus", resource.BusinessStatus, optionsValueSet)
}
func (resource *Task) TaskIntent() templ.Component {
	optionsValueSet := VSTask_intent

	if resource != nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *Task) TaskPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource != nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Task) TaskCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *Task) TaskPerformerType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerType", &resource.PerformerType[0], optionsValueSet)
}
func (resource *Task) TaskReasonCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", resource.ReasonCode, optionsValueSet)
}
func (resource *Task) TaskInputType(numInput int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Input) >= numInput {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Input[numInput].Type, optionsValueSet)
}
func (resource *Task) TaskOutputType(numOutput int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Output) >= numOutput {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Output[numOutput].Type, optionsValueSet)
}
