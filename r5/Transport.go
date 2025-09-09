package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Transport
type Transport struct {
	Id                    *string               `json:"id,omitempty"`
	Meta                  *Meta                 `json:"meta,omitempty"`
	ImplicitRules         *string               `json:"implicitRules,omitempty"`
	Language              *string               `json:"language,omitempty"`
	Text                  *Narrative            `json:"text,omitempty"`
	Contained             []Resource            `json:"contained,omitempty"`
	Extension             []Extension           `json:"extension,omitempty"`
	ModifierExtension     []Extension           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier          `json:"identifier,omitempty"`
	InstantiatesCanonical *string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference           `json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier           `json:"groupIdentifier,omitempty"`
	PartOf                []Reference           `json:"partOf,omitempty"`
	Status                *string               `json:"status,omitempty"`
	StatusReason          *CodeableConcept      `json:"statusReason,omitempty"`
	Intent                string                `json:"intent"`
	Priority              *string               `json:"priority,omitempty"`
	Code                  *CodeableConcept      `json:"code,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Focus                 *Reference            `json:"focus,omitempty"`
	For                   *Reference            `json:"for,omitempty"`
	Encounter             *Reference            `json:"encounter,omitempty"`
	CompletionTime        *time.Time            `json:"completionTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	AuthoredOn            *time.Time            `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	LastModified          *time.Time            `json:"lastModified,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Requester             *Reference            `json:"requester,omitempty"`
	PerformerType         []CodeableConcept     `json:"performerType,omitempty"`
	Owner                 *Reference            `json:"owner,omitempty"`
	Location              *Reference            `json:"location,omitempty"`
	Insurance             []Reference           `json:"insurance,omitempty"`
	Note                  []Annotation          `json:"note,omitempty"`
	RelevantHistory       []Reference           `json:"relevantHistory,omitempty"`
	Restriction           *TransportRestriction `json:"restriction,omitempty"`
	Input                 []TransportInput      `json:"input,omitempty"`
	Output                []TransportOutput     `json:"output,omitempty"`
	RequestedLocation     Reference             `json:"requestedLocation"`
	CurrentLocation       Reference             `json:"currentLocation"`
	Reason                *CodeableReference    `json:"reason,omitempty"`
	History               *Reference            `json:"history,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Transport
type TransportRestriction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Repetitions       *int        `json:"repetitions,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Recipient         []Reference `json:"recipient,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Transport
type TransportInput struct {
	Id                         *string               `json:"id,omitempty"`
	Extension                  []Extension           `json:"extension,omitempty"`
	ModifierExtension          []Extension           `json:"modifierExtension,omitempty"`
	Type                       CodeableConcept       `json:"type"`
	ValueBase64Binary          string                `json:"valueBase64Binary"`
	ValueBoolean               bool                  `json:"valueBoolean"`
	ValueCanonical             string                `json:"valueCanonical"`
	ValueCode                  string                `json:"valueCode"`
	ValueDate                  time.Time             `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime              time.Time             `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	ValueDecimal               float64               `json:"valueDecimal"`
	ValueId                    string                `json:"valueId"`
	ValueInstant               string                `json:"valueInstant"`
	ValueInteger               int                   `json:"valueInteger"`
	ValueInteger64             int64                 `json:"valueInteger64"`
	ValueMarkdown              string                `json:"valueMarkdown"`
	ValueOid                   string                `json:"valueOid"`
	ValuePositiveInt           int                   `json:"valuePositiveInt"`
	ValueString                string                `json:"valueString"`
	ValueTime                  string                `json:"valueTime"`
	ValueUnsignedInt           int                   `json:"valueUnsignedInt"`
	ValueUri                   string                `json:"valueUri"`
	ValueUrl                   string                `json:"valueUrl"`
	ValueUuid                  string                `json:"valueUuid"`
	ValueAddress               Address               `json:"valueAddress"`
	ValueAge                   Age                   `json:"valueAge"`
	ValueAnnotation            Annotation            `json:"valueAnnotation"`
	ValueAttachment            Attachment            `json:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `json:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `json:"valueCodeableReference"`
	ValueCoding                Coding                `json:"valueCoding"`
	ValueContactPoint          ContactPoint          `json:"valueContactPoint"`
	ValueCount                 Count                 `json:"valueCount"`
	ValueDistance              Distance              `json:"valueDistance"`
	ValueDuration              Duration              `json:"valueDuration"`
	ValueHumanName             HumanName             `json:"valueHumanName"`
	ValueIdentifier            Identifier            `json:"valueIdentifier"`
	ValueMoney                 Money                 `json:"valueMoney"`
	ValuePeriod                Period                `json:"valuePeriod"`
	ValueQuantity              Quantity              `json:"valueQuantity"`
	ValueRange                 Range                 `json:"valueRange"`
	ValueRatio                 Ratio                 `json:"valueRatio"`
	ValueRatioRange            RatioRange            `json:"valueRatioRange"`
	ValueReference             Reference             `json:"valueReference"`
	ValueSampledData           SampledData           `json:"valueSampledData"`
	ValueSignature             Signature             `json:"valueSignature"`
	ValueTiming                Timing                `json:"valueTiming"`
	ValueContactDetail         ContactDetail         `json:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `json:"valueDataRequirement"`
	ValueExpression            Expression            `json:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `json:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `json:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `json:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `json:"valueUsageContext"`
	ValueAvailability          Availability          `json:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `json:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `json:"valueDosage"`
	ValueMeta                  Meta                  `json:"valueMeta"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Transport
type TransportOutput struct {
	Id                         *string               `json:"id,omitempty"`
	Extension                  []Extension           `json:"extension,omitempty"`
	ModifierExtension          []Extension           `json:"modifierExtension,omitempty"`
	Type                       CodeableConcept       `json:"type"`
	ValueBase64Binary          string                `json:"valueBase64Binary"`
	ValueBoolean               bool                  `json:"valueBoolean"`
	ValueCanonical             string                `json:"valueCanonical"`
	ValueCode                  string                `json:"valueCode"`
	ValueDate                  time.Time             `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime              time.Time             `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	ValueDecimal               float64               `json:"valueDecimal"`
	ValueId                    string                `json:"valueId"`
	ValueInstant               string                `json:"valueInstant"`
	ValueInteger               int                   `json:"valueInteger"`
	ValueInteger64             int64                 `json:"valueInteger64"`
	ValueMarkdown              string                `json:"valueMarkdown"`
	ValueOid                   string                `json:"valueOid"`
	ValuePositiveInt           int                   `json:"valuePositiveInt"`
	ValueString                string                `json:"valueString"`
	ValueTime                  string                `json:"valueTime"`
	ValueUnsignedInt           int                   `json:"valueUnsignedInt"`
	ValueUri                   string                `json:"valueUri"`
	ValueUrl                   string                `json:"valueUrl"`
	ValueUuid                  string                `json:"valueUuid"`
	ValueAddress               Address               `json:"valueAddress"`
	ValueAge                   Age                   `json:"valueAge"`
	ValueAnnotation            Annotation            `json:"valueAnnotation"`
	ValueAttachment            Attachment            `json:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `json:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `json:"valueCodeableReference"`
	ValueCoding                Coding                `json:"valueCoding"`
	ValueContactPoint          ContactPoint          `json:"valueContactPoint"`
	ValueCount                 Count                 `json:"valueCount"`
	ValueDistance              Distance              `json:"valueDistance"`
	ValueDuration              Duration              `json:"valueDuration"`
	ValueHumanName             HumanName             `json:"valueHumanName"`
	ValueIdentifier            Identifier            `json:"valueIdentifier"`
	ValueMoney                 Money                 `json:"valueMoney"`
	ValuePeriod                Period                `json:"valuePeriod"`
	ValueQuantity              Quantity              `json:"valueQuantity"`
	ValueRange                 Range                 `json:"valueRange"`
	ValueRatio                 Ratio                 `json:"valueRatio"`
	ValueRatioRange            RatioRange            `json:"valueRatioRange"`
	ValueReference             Reference             `json:"valueReference"`
	ValueSampledData           SampledData           `json:"valueSampledData"`
	ValueSignature             Signature             `json:"valueSignature"`
	ValueTiming                Timing                `json:"valueTiming"`
	ValueContactDetail         ContactDetail         `json:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `json:"valueDataRequirement"`
	ValueExpression            Expression            `json:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `json:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `json:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `json:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `json:"valueUsageContext"`
	ValueAvailability          Availability          `json:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `json:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `json:"valueDosage"`
	ValueMeta                  Meta                  `json:"valueMeta"`
}

type OtherTransport Transport

// on convert struct to json, automatically add resourceType=Transport
func (r Transport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTransport
		ResourceType string `json:"resourceType"`
	}{
		OtherTransport: OtherTransport(r),
		ResourceType:   "Transport",
	})
}
func (r Transport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Transport/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Transport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Transport) T_InstantiatesCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Transport.InstantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("Transport.InstantiatesCanonical", resource.InstantiatesCanonical, htmlAttrs)
}
func (resource *Transport) T_InstantiatesUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Transport.InstantiatesUri", nil, htmlAttrs)
	}
	return StringInput("Transport.InstantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Transport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSTransport_status

	if resource == nil {
		return CodeSelect("Transport.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Transport.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Transport.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSTransport_intent

	if resource == nil {
		return CodeSelect("Transport.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Transport.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Transport.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Transport.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Transport.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Transport.Description", nil, htmlAttrs)
	}
	return StringInput("Transport.Description", resource.Description, htmlAttrs)
}
func (resource *Transport) T_CompletionTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Transport.CompletionTime", nil, htmlAttrs)
	}
	return DateTimeInput("Transport.CompletionTime", resource.CompletionTime, htmlAttrs)
}
func (resource *Transport) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Transport.AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("Transport.AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *Transport) T_LastModified(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Transport.LastModified", nil, htmlAttrs)
	}
	return DateTimeInput("Transport.LastModified", resource.LastModified, htmlAttrs)
}
func (resource *Transport) T_PerformerType(numPerformerType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformerType >= len(resource.PerformerType) {
		return CodeableConceptSelect("Transport.PerformerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.PerformerType["+strconv.Itoa(numPerformerType)+"]", &resource.PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Transport.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Transport.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Transport) T_RestrictionRepetitions(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Transport.Restriction.Repetitions", nil, htmlAttrs)
	}
	return IntInput("Transport.Restriction.Repetitions", resource.Restriction.Repetitions, htmlAttrs)
}
func (resource *Transport) T_InputType(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].Type", &resource.Input[numInput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueBase64Binary(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueBase64Binary", &resource.Input[numInput].ValueBase64Binary, htmlAttrs)
}
func (resource *Transport) T_InputValueBoolean(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return BoolInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueBoolean", &resource.Input[numInput].ValueBoolean, htmlAttrs)
}
func (resource *Transport) T_InputValueCanonical(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueCanonical", &resource.Input[numInput].ValueCanonical, htmlAttrs)
}
func (resource *Transport) T_InputValueCode(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCode", &resource.Input[numInput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueDate(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueDate", nil, htmlAttrs)
	}
	return DateInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueDate", &resource.Input[numInput].ValueDate, htmlAttrs)
}
func (resource *Transport) T_InputValueDateTime(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DateTimeInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueDateTime", &resource.Input[numInput].ValueDateTime, htmlAttrs)
}
func (resource *Transport) T_InputValueDecimal(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Float64Input("Transport.Input["+strconv.Itoa(numInput)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Transport.Input["+strconv.Itoa(numInput)+"].ValueDecimal", &resource.Input[numInput].ValueDecimal, htmlAttrs)
}
func (resource *Transport) T_InputValueId(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueId", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueId", &resource.Input[numInput].ValueId, htmlAttrs)
}
func (resource *Transport) T_InputValueInstant(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueInstant", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueInstant", &resource.Input[numInput].ValueInstant, htmlAttrs)
}
func (resource *Transport) T_InputValueInteger(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueInteger", &resource.Input[numInput].ValueInteger, htmlAttrs)
}
func (resource *Transport) T_InputValueInteger64(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Int64Input("Transport.Input["+strconv.Itoa(numInput)+"].ValueInteger64", nil, htmlAttrs)
	}
	return Int64Input("Transport.Input["+strconv.Itoa(numInput)+"].ValueInteger64", &resource.Input[numInput].ValueInteger64, htmlAttrs)
}
func (resource *Transport) T_InputValueMarkdown(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueMarkdown", &resource.Input[numInput].ValueMarkdown, htmlAttrs)
}
func (resource *Transport) T_InputValueOid(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueOid", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueOid", &resource.Input[numInput].ValueOid, htmlAttrs)
}
func (resource *Transport) T_InputValuePositiveInt(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValuePositiveInt", &resource.Input[numInput].ValuePositiveInt, htmlAttrs)
}
func (resource *Transport) T_InputValueString(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueString", &resource.Input[numInput].ValueString, htmlAttrs)
}
func (resource *Transport) T_InputValueTime(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueTime", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueTime", &resource.Input[numInput].ValueTime, htmlAttrs)
}
func (resource *Transport) T_InputValueUnsignedInt(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUnsignedInt", &resource.Input[numInput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Transport) T_InputValueUri(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUri", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUri", &resource.Input[numInput].ValueUri, htmlAttrs)
}
func (resource *Transport) T_InputValueUrl(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUrl", &resource.Input[numInput].ValueUrl, htmlAttrs)
}
func (resource *Transport) T_InputValueUuid(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUuid", nil, htmlAttrs)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].ValueUuid", &resource.Input[numInput].ValueUuid, htmlAttrs)
}
func (resource *Transport) T_InputValueAnnotation(numInput int, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AnnotationTextArea("Transport.Input["+strconv.Itoa(numInput)+"].ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Transport.Input["+strconv.Itoa(numInput)+"].ValueAnnotation", &resource.Input[numInput].ValueAnnotation, htmlAttrs)
}
func (resource *Transport) T_InputValueCodeableConcept(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCodeableConcept", &resource.Input[numInput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueCoding(numInput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodingSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Transport.Input["+strconv.Itoa(numInput)+"].ValueCoding", &resource.Input[numInput].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputType(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].Type", &resource.Output[numOutput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueBase64Binary(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueBase64Binary", &resource.Output[numOutput].ValueBase64Binary, htmlAttrs)
}
func (resource *Transport) T_OutputValueBoolean(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return BoolInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueBoolean", &resource.Output[numOutput].ValueBoolean, htmlAttrs)
}
func (resource *Transport) T_OutputValueCanonical(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCanonical", &resource.Output[numOutput].ValueCanonical, htmlAttrs)
}
func (resource *Transport) T_OutputValueCode(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCode", &resource.Output[numOutput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueDate(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDate", nil, htmlAttrs)
	}
	return DateInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDate", &resource.Output[numOutput].ValueDate, htmlAttrs)
}
func (resource *Transport) T_OutputValueDateTime(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DateTimeInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDateTime", &resource.Output[numOutput].ValueDateTime, htmlAttrs)
}
func (resource *Transport) T_OutputValueDecimal(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Float64Input("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Transport.Output["+strconv.Itoa(numOutput)+"].ValueDecimal", &resource.Output[numOutput].ValueDecimal, htmlAttrs)
}
func (resource *Transport) T_OutputValueId(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueId", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueId", &resource.Output[numOutput].ValueId, htmlAttrs)
}
func (resource *Transport) T_OutputValueInstant(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInstant", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInstant", &resource.Output[numOutput].ValueInstant, htmlAttrs)
}
func (resource *Transport) T_OutputValueInteger(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInteger", &resource.Output[numOutput].ValueInteger, htmlAttrs)
}
func (resource *Transport) T_OutputValueInteger64(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Int64Input("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInteger64", nil, htmlAttrs)
	}
	return Int64Input("Transport.Output["+strconv.Itoa(numOutput)+"].ValueInteger64", &resource.Output[numOutput].ValueInteger64, htmlAttrs)
}
func (resource *Transport) T_OutputValueMarkdown(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueMarkdown", &resource.Output[numOutput].ValueMarkdown, htmlAttrs)
}
func (resource *Transport) T_OutputValueOid(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueOid", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueOid", &resource.Output[numOutput].ValueOid, htmlAttrs)
}
func (resource *Transport) T_OutputValuePositiveInt(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValuePositiveInt", &resource.Output[numOutput].ValuePositiveInt, htmlAttrs)
}
func (resource *Transport) T_OutputValueString(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueString", &resource.Output[numOutput].ValueString, htmlAttrs)
}
func (resource *Transport) T_OutputValueTime(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueTime", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueTime", &resource.Output[numOutput].ValueTime, htmlAttrs)
}
func (resource *Transport) T_OutputValueUnsignedInt(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUnsignedInt", &resource.Output[numOutput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Transport) T_OutputValueUri(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUri", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUri", &resource.Output[numOutput].ValueUri, htmlAttrs)
}
func (resource *Transport) T_OutputValueUrl(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUrl", &resource.Output[numOutput].ValueUrl, htmlAttrs)
}
func (resource *Transport) T_OutputValueUuid(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUuid", nil, htmlAttrs)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].ValueUuid", &resource.Output[numOutput].ValueUuid, htmlAttrs)
}
func (resource *Transport) T_OutputValueAnnotation(numOutput int, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AnnotationTextArea("Transport.Output["+strconv.Itoa(numOutput)+"].ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Transport.Output["+strconv.Itoa(numOutput)+"].ValueAnnotation", &resource.Output[numOutput].ValueAnnotation, htmlAttrs)
}
func (resource *Transport) T_OutputValueCodeableConcept(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCodeableConcept", &resource.Output[numOutput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueCoding(numOutput int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodingSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Transport.Output["+strconv.Itoa(numOutput)+"].ValueCoding", &resource.Output[numOutput].ValueCoding, optionsValueSet, htmlAttrs)
}
