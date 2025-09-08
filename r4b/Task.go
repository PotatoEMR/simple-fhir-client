package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	AuthoredOn            *time.Time        `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	LastModified          *time.Time        `json:"lastModified,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate                time.Time           `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime            time.Time           `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate                time.Time           `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime            time.Time           `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Task) T_InstantiatesCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InstantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical", resource.InstantiatesCanonical, htmlAttrs)
}
func (resource *Task) T_InstantiatesUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InstantiatesUri", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Task) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSTask_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_BusinessStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("BusinessStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BusinessStatus", resource.BusinessStatus, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSTask_intent

	if resource == nil {
		return CodeSelect("Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *Task) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *Task) T_LastModified(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("LastModified", nil, htmlAttrs)
	}
	return DateTimeInput("LastModified", resource.LastModified, htmlAttrs)
}
func (resource *Task) T_PerformerType(numPerformerType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformerType >= len(resource.PerformerType) {
		return CodeableConceptSelect("PerformerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PerformerType["+strconv.Itoa(numPerformerType)+"]", &resource.PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Task) T_ReasonCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ReasonCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode", resource.ReasonCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Task) T_RestrictionRepetitions(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RestrictionRepetitions", nil, htmlAttrs)
	}
	return IntInput("RestrictionRepetitions", resource.Restriction.Repetitions, htmlAttrs)
}
func (resource *Task) T_InputType(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("Input["+strconv.Itoa(numInput)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Input["+strconv.Itoa(numInput)+"]Type", &resource.Input[numInput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueBase64Binary(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueBase64Binary", &resource.Input[numInput].ValueBase64Binary, htmlAttrs)
}
func (resource *Task) T_InputValueBoolean(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return BoolInput("Input["+strconv.Itoa(numInput)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Input["+strconv.Itoa(numInput)+"]ValueBoolean", &resource.Input[numInput].ValueBoolean, htmlAttrs)
}
func (resource *Task) T_InputValueCanonical(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueCanonical", &resource.Input[numInput].ValueCanonical, htmlAttrs)
}
func (resource *Task) T_InputValueCode(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeSelect("Input["+strconv.Itoa(numInput)+"]ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Input["+strconv.Itoa(numInput)+"]ValueCode", &resource.Input[numInput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueDate(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateInput("Input["+strconv.Itoa(numInput)+"]ValueDate", nil, htmlAttrs)
	}
	return DateInput("Input["+strconv.Itoa(numInput)+"]ValueDate", &resource.Input[numInput].ValueDate, htmlAttrs)
}
func (resource *Task) T_InputValueDateTime(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateTimeInput("Input["+strconv.Itoa(numInput)+"]ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Input["+strconv.Itoa(numInput)+"]ValueDateTime", &resource.Input[numInput].ValueDateTime, htmlAttrs)
}
func (resource *Task) T_InputValueDecimal(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Float64Input("Input["+strconv.Itoa(numInput)+"]ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Input["+strconv.Itoa(numInput)+"]ValueDecimal", &resource.Input[numInput].ValueDecimal, htmlAttrs)
}
func (resource *Task) T_InputValueId(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueId", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueId", &resource.Input[numInput].ValueId, htmlAttrs)
}
func (resource *Task) T_InputValueInstant(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueInstant", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueInstant", &resource.Input[numInput].ValueInstant, htmlAttrs)
}
func (resource *Task) T_InputValueInteger(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Input["+strconv.Itoa(numInput)+"]ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Input["+strconv.Itoa(numInput)+"]ValueInteger", &resource.Input[numInput].ValueInteger, htmlAttrs)
}
func (resource *Task) T_InputValueMarkdown(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueMarkdown", &resource.Input[numInput].ValueMarkdown, htmlAttrs)
}
func (resource *Task) T_InputValueOid(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueOid", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueOid", &resource.Input[numInput].ValueOid, htmlAttrs)
}
func (resource *Task) T_InputValuePositiveInt(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Input["+strconv.Itoa(numInput)+"]ValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("Input["+strconv.Itoa(numInput)+"]ValuePositiveInt", &resource.Input[numInput].ValuePositiveInt, htmlAttrs)
}
func (resource *Task) T_InputValueString(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueString", &resource.Input[numInput].ValueString, htmlAttrs)
}
func (resource *Task) T_InputValueTime(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueTime", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueTime", &resource.Input[numInput].ValueTime, htmlAttrs)
}
func (resource *Task) T_InputValueUnsignedInt(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Input["+strconv.Itoa(numInput)+"]ValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("Input["+strconv.Itoa(numInput)+"]ValueUnsignedInt", &resource.Input[numInput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Task) T_InputValueUri(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUri", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUri", &resource.Input[numInput].ValueUri, htmlAttrs)
}
func (resource *Task) T_InputValueUrl(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUrl", &resource.Input[numInput].ValueUrl, htmlAttrs)
}
func (resource *Task) T_InputValueUuid(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUuid", nil, htmlAttrs)
	}
	return StringInput("Input["+strconv.Itoa(numInput)+"]ValueUuid", &resource.Input[numInput].ValueUuid, htmlAttrs)
}
func (resource *Task) T_InputValueAnnotation(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AnnotationTextArea("Input["+strconv.Itoa(numInput)+"]ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Input["+strconv.Itoa(numInput)+"]ValueAnnotation", &resource.Input[numInput].ValueAnnotation, htmlAttrs)
}
func (resource *Task) T_InputValueCodeableConcept(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("Input["+strconv.Itoa(numInput)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Input["+strconv.Itoa(numInput)+"]ValueCodeableConcept", &resource.Input[numInput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_InputValueCoding(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodingSelect("Input["+strconv.Itoa(numInput)+"]ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Input["+strconv.Itoa(numInput)+"]ValueCoding", &resource.Input[numInput].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputType(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("Output["+strconv.Itoa(numOutput)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Output["+strconv.Itoa(numOutput)+"]Type", &resource.Output[numOutput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueBase64Binary(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueBase64Binary", &resource.Output[numOutput].ValueBase64Binary, htmlAttrs)
}
func (resource *Task) T_OutputValueBoolean(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return BoolInput("Output["+strconv.Itoa(numOutput)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Output["+strconv.Itoa(numOutput)+"]ValueBoolean", &resource.Output[numOutput].ValueBoolean, htmlAttrs)
}
func (resource *Task) T_OutputValueCanonical(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueCanonical", &resource.Output[numOutput].ValueCanonical, htmlAttrs)
}
func (resource *Task) T_OutputValueCode(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeSelect("Output["+strconv.Itoa(numOutput)+"]ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Output["+strconv.Itoa(numOutput)+"]ValueCode", &resource.Output[numOutput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueDate(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateInput("Output["+strconv.Itoa(numOutput)+"]ValueDate", nil, htmlAttrs)
	}
	return DateInput("Output["+strconv.Itoa(numOutput)+"]ValueDate", &resource.Output[numOutput].ValueDate, htmlAttrs)
}
func (resource *Task) T_OutputValueDateTime(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateTimeInput("Output["+strconv.Itoa(numOutput)+"]ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Output["+strconv.Itoa(numOutput)+"]ValueDateTime", &resource.Output[numOutput].ValueDateTime, htmlAttrs)
}
func (resource *Task) T_OutputValueDecimal(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Float64Input("Output["+strconv.Itoa(numOutput)+"]ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Output["+strconv.Itoa(numOutput)+"]ValueDecimal", &resource.Output[numOutput].ValueDecimal, htmlAttrs)
}
func (resource *Task) T_OutputValueId(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueId", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueId", &resource.Output[numOutput].ValueId, htmlAttrs)
}
func (resource *Task) T_OutputValueInstant(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueInstant", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueInstant", &resource.Output[numOutput].ValueInstant, htmlAttrs)
}
func (resource *Task) T_OutputValueInteger(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Output["+strconv.Itoa(numOutput)+"]ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Output["+strconv.Itoa(numOutput)+"]ValueInteger", &resource.Output[numOutput].ValueInteger, htmlAttrs)
}
func (resource *Task) T_OutputValueMarkdown(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueMarkdown", &resource.Output[numOutput].ValueMarkdown, htmlAttrs)
}
func (resource *Task) T_OutputValueOid(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueOid", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueOid", &resource.Output[numOutput].ValueOid, htmlAttrs)
}
func (resource *Task) T_OutputValuePositiveInt(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Output["+strconv.Itoa(numOutput)+"]ValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("Output["+strconv.Itoa(numOutput)+"]ValuePositiveInt", &resource.Output[numOutput].ValuePositiveInt, htmlAttrs)
}
func (resource *Task) T_OutputValueString(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueString", &resource.Output[numOutput].ValueString, htmlAttrs)
}
func (resource *Task) T_OutputValueTime(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueTime", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueTime", &resource.Output[numOutput].ValueTime, htmlAttrs)
}
func (resource *Task) T_OutputValueUnsignedInt(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Output["+strconv.Itoa(numOutput)+"]ValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("Output["+strconv.Itoa(numOutput)+"]ValueUnsignedInt", &resource.Output[numOutput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Task) T_OutputValueUri(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUri", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUri", &resource.Output[numOutput].ValueUri, htmlAttrs)
}
func (resource *Task) T_OutputValueUrl(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUrl", &resource.Output[numOutput].ValueUrl, htmlAttrs)
}
func (resource *Task) T_OutputValueUuid(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUuid", nil, htmlAttrs)
	}
	return StringInput("Output["+strconv.Itoa(numOutput)+"]ValueUuid", &resource.Output[numOutput].ValueUuid, htmlAttrs)
}
func (resource *Task) T_OutputValueAnnotation(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AnnotationTextArea("Output["+strconv.Itoa(numOutput)+"]ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Output["+strconv.Itoa(numOutput)+"]ValueAnnotation", &resource.Output[numOutput].ValueAnnotation, htmlAttrs)
}
func (resource *Task) T_OutputValueCodeableConcept(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("Output["+strconv.Itoa(numOutput)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Output["+strconv.Itoa(numOutput)+"]ValueCodeableConcept", &resource.Output[numOutput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Task) T_OutputValueCoding(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodingSelect("Output["+strconv.Itoa(numOutput)+"]ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Output["+strconv.Itoa(numOutput)+"]ValueCoding", &resource.Output[numOutput].ValueCoding, optionsValueSet, htmlAttrs)
}
