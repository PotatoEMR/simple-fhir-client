package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Task
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

// http://hl7.org/fhir/r4/StructureDefinition/Task
type TaskRestriction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Repetitions       *int        `json:"repetitions,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Recipient         []Reference `json:"recipient,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Task
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

// http://hl7.org/fhir/r4/StructureDefinition/Task
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
func (r Task) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Task/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Task"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Task) T_InstantiatesCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical", resource.InstantiatesCanonical, htmlAttrs)
}
func (resource *Task) T_InstantiatesUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesUri", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Task) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSTask_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_BusinessStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("businessStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("businessStatus", resource.BusinessStatus, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSTask_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Task) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *Task) T_LastModified(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("lastModified", nil, htmlAttrs)
	}
	return DateTimeInput("lastModified", resource.LastModified, htmlAttrs)
}
func (resource *Task) T_PerformerType(numPerformerType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformerType >= len(resource.PerformerType) {
		return CodeableConceptSelect("performerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType["+strconv.Itoa(numPerformerType)+"]", &resource.PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Task) T_ReasonCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode", resource.ReasonCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Task) T_RestrictionRepetitions(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("restriction.repetitions", nil, htmlAttrs)
	}
	return IntInput("restriction.repetitions", resource.Restriction.Repetitions, htmlAttrs)
}
func (resource *Task) T_InputType(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].type", &resource.Input[numInput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueBase64Binary(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueBase64Binary", &resource.Input[numInput].ValueBase64Binary, htmlAttrs)
}
func (resource *Task) T_InputValueBoolean(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return BoolInput("input["+strconv.Itoa(numInput)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("input["+strconv.Itoa(numInput)+"].valueBoolean", &resource.Input[numInput].ValueBoolean, htmlAttrs)
}
func (resource *Task) T_InputValueCanonical(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueCanonical", &resource.Input[numInput].ValueCanonical, htmlAttrs)
}
func (resource *Task) T_InputValueCode(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeSelect("input["+strconv.Itoa(numInput)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("input["+strconv.Itoa(numInput)+"].valueCode", &resource.Input[numInput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueDate(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateInput("input["+strconv.Itoa(numInput)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("input["+strconv.Itoa(numInput)+"].valueDate", &resource.Input[numInput].ValueDate, htmlAttrs)
}
func (resource *Task) T_InputValueDateTime(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateTimeInput("input["+strconv.Itoa(numInput)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("input["+strconv.Itoa(numInput)+"].valueDateTime", &resource.Input[numInput].ValueDateTime, htmlAttrs)
}
func (resource *Task) T_InputValueDecimal(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Float64Input("input["+strconv.Itoa(numInput)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("input["+strconv.Itoa(numInput)+"].valueDecimal", &resource.Input[numInput].ValueDecimal, htmlAttrs)
}
func (resource *Task) T_InputValueId(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueId", &resource.Input[numInput].ValueId, htmlAttrs)
}
func (resource *Task) T_InputValueInstant(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueInstant", &resource.Input[numInput].ValueInstant, htmlAttrs)
}
func (resource *Task) T_InputValueInteger(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valueInteger", &resource.Input[numInput].ValueInteger, htmlAttrs)
}
func (resource *Task) T_InputValueMarkdown(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueMarkdown", &resource.Input[numInput].ValueMarkdown, htmlAttrs)
}
func (resource *Task) T_InputValueOid(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueOid", &resource.Input[numInput].ValueOid, htmlAttrs)
}
func (resource *Task) T_InputValuePositiveInt(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valuePositiveInt", &resource.Input[numInput].ValuePositiveInt, htmlAttrs)
}
func (resource *Task) T_InputValueString(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueString", &resource.Input[numInput].ValueString, htmlAttrs)
}
func (resource *Task) T_InputValueTime(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueTime", &resource.Input[numInput].ValueTime, htmlAttrs)
}
func (resource *Task) T_InputValueUnsignedInt(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valueUnsignedInt", &resource.Input[numInput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Task) T_InputValueUri(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUri", &resource.Input[numInput].ValueUri, htmlAttrs)
}
func (resource *Task) T_InputValueUrl(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUrl", &resource.Input[numInput].ValueUrl, htmlAttrs)
}
func (resource *Task) T_InputValueUuid(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUuid", &resource.Input[numInput].ValueUuid, htmlAttrs)
}
func (resource *Task) T_InputValueAnnotation(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AnnotationTextArea("input["+strconv.Itoa(numInput)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("input["+strconv.Itoa(numInput)+"].valueAnnotation", &resource.Input[numInput].ValueAnnotation, htmlAttrs)
}
func (resource *Task) T_InputValueCodeableConcept(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].valueCodeableConcept", &resource.Input[numInput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueCoding(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodingSelect("input["+strconv.Itoa(numInput)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("input["+strconv.Itoa(numInput)+"].valueCoding", &resource.Input[numInput].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputType(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].type", &resource.Output[numOutput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueBase64Binary(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueBase64Binary", &resource.Output[numOutput].ValueBase64Binary, htmlAttrs)
}
func (resource *Task) T_OutputValueBoolean(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return BoolInput("output["+strconv.Itoa(numOutput)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("output["+strconv.Itoa(numOutput)+"].valueBoolean", &resource.Output[numOutput].ValueBoolean, htmlAttrs)
}
func (resource *Task) T_OutputValueCanonical(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueCanonical", &resource.Output[numOutput].ValueCanonical, htmlAttrs)
}
func (resource *Task) T_OutputValueCode(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeSelect("output["+strconv.Itoa(numOutput)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("output["+strconv.Itoa(numOutput)+"].valueCode", &resource.Output[numOutput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueDate(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateInput("output["+strconv.Itoa(numOutput)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("output["+strconv.Itoa(numOutput)+"].valueDate", &resource.Output[numOutput].ValueDate, htmlAttrs)
}
func (resource *Task) T_OutputValueDateTime(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateTimeInput("output["+strconv.Itoa(numOutput)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("output["+strconv.Itoa(numOutput)+"].valueDateTime", &resource.Output[numOutput].ValueDateTime, htmlAttrs)
}
func (resource *Task) T_OutputValueDecimal(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Float64Input("output["+strconv.Itoa(numOutput)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("output["+strconv.Itoa(numOutput)+"].valueDecimal", &resource.Output[numOutput].ValueDecimal, htmlAttrs)
}
func (resource *Task) T_OutputValueId(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueId", &resource.Output[numOutput].ValueId, htmlAttrs)
}
func (resource *Task) T_OutputValueInstant(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueInstant", &resource.Output[numOutput].ValueInstant, htmlAttrs)
}
func (resource *Task) T_OutputValueInteger(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valueInteger", &resource.Output[numOutput].ValueInteger, htmlAttrs)
}
func (resource *Task) T_OutputValueMarkdown(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueMarkdown", &resource.Output[numOutput].ValueMarkdown, htmlAttrs)
}
func (resource *Task) T_OutputValueOid(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueOid", &resource.Output[numOutput].ValueOid, htmlAttrs)
}
func (resource *Task) T_OutputValuePositiveInt(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valuePositiveInt", &resource.Output[numOutput].ValuePositiveInt, htmlAttrs)
}
func (resource *Task) T_OutputValueString(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueString", &resource.Output[numOutput].ValueString, htmlAttrs)
}
func (resource *Task) T_OutputValueTime(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueTime", &resource.Output[numOutput].ValueTime, htmlAttrs)
}
func (resource *Task) T_OutputValueUnsignedInt(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valueUnsignedInt", &resource.Output[numOutput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Task) T_OutputValueUri(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUri", &resource.Output[numOutput].ValueUri, htmlAttrs)
}
func (resource *Task) T_OutputValueUrl(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUrl", &resource.Output[numOutput].ValueUrl, htmlAttrs)
}
func (resource *Task) T_OutputValueUuid(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUuid", &resource.Output[numOutput].ValueUuid, htmlAttrs)
}
func (resource *Task) T_OutputValueAnnotation(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AnnotationTextArea("output["+strconv.Itoa(numOutput)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("output["+strconv.Itoa(numOutput)+"].valueAnnotation", &resource.Output[numOutput].ValueAnnotation, htmlAttrs)
}
func (resource *Task) T_OutputValueCodeableConcept(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].valueCodeableConcept", &resource.Output[numOutput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueCoding(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodingSelect("output["+strconv.Itoa(numOutput)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("output["+strconv.Itoa(numOutput)+"].valueCoding", &resource.Output[numOutput].ValueCoding, optionsValueSet, htmlAttrs)
}
