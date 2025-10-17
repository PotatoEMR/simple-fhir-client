package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

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
	CompletionTime        *FhirDateTime         `json:"completionTime,omitempty"`
	AuthoredOn            *FhirDateTime         `json:"authoredOn,omitempty"`
	LastModified          *FhirDateTime         `json:"lastModified,omitempty"`
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
	ValueDate                  FhirDate              `json:"valueDate"`
	ValueDateTime              FhirDateTime          `json:"valueDateTime"`
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
	ValueDate                  FhirDate              `json:"valueDate"`
	ValueDateTime              FhirDateTime          `json:"valueDateTime"`
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

// struct -> json, automatically add resourceType=Patient
func (r Transport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTransport
		ResourceType string `json:"resourceType"`
	}{
		OtherTransport: OtherTransport(r),
		ResourceType:   "Transport",
	})
}

// json -> struct, first reject if resourceType != Transport
func (r *Transport) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Transport" {
		return errors.New("resourceType not Transport")
	}
	return json.Unmarshal(data, (*OtherTransport)(r))
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
func (resource *Transport) T_InstantiatesCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical", resource.InstantiatesCanonical, htmlAttrs)
}
func (resource *Transport) T_InstantiatesUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesUri", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Transport) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Transport) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
}
func (resource *Transport) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *Transport) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSTransport_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSTransport_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Transport) T_Focus(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "focus", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focus", resource.Focus, htmlAttrs)
}
func (resource *Transport) T_For(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "for", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "for", resource.For, htmlAttrs)
}
func (resource *Transport) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Transport) T_CompletionTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("completionTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("completionTime", resource.CompletionTime, htmlAttrs)
}
func (resource *Transport) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *Transport) T_LastModified(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("lastModified", nil, htmlAttrs)
	}
	return FhirDateTimeInput("lastModified", resource.LastModified, htmlAttrs)
}
func (resource *Transport) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *Transport) T_PerformerType(numPerformerType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformerType >= len(resource.PerformerType) {
		return CodeableConceptSelect("performerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType["+strconv.Itoa(numPerformerType)+"]", &resource.PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_Owner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "owner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "owner", resource.Owner, htmlAttrs)
}
func (resource *Transport) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *Transport) T_Insurance(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", &resource.Insurance[numInsurance], htmlAttrs)
}
func (resource *Transport) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Transport) T_RelevantHistory(frs []FhirResource, numRelevantHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelevantHistory >= len(resource.RelevantHistory) {
		return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", &resource.RelevantHistory[numRelevantHistory], htmlAttrs)
}
func (resource *Transport) T_RequestedLocation(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requestedLocation", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requestedLocation", &resource.RequestedLocation, htmlAttrs)
}
func (resource *Transport) T_CurrentLocation(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "currentLocation", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "currentLocation", &resource.CurrentLocation, htmlAttrs)
}
func (resource *Transport) T_Reason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("reason", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason", resource.Reason, htmlAttrs)
}
func (resource *Transport) T_History(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "history", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "history", resource.History, htmlAttrs)
}
func (resource *Transport) T_RestrictionRepetitions(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("restriction.repetitions", nil, htmlAttrs)
	}
	return IntInput("restriction.repetitions", resource.Restriction.Repetitions, htmlAttrs)
}
func (resource *Transport) T_RestrictionPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("restriction.period", nil, htmlAttrs)
	}
	return PeriodInput("restriction.period", resource.Restriction.Period, htmlAttrs)
}
func (resource *Transport) T_RestrictionRecipient(frs []FhirResource, numRecipient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecipient >= len(resource.Restriction.Recipient) {
		return ReferenceInput(frs, "restriction.recipient["+strconv.Itoa(numRecipient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "restriction.recipient["+strconv.Itoa(numRecipient)+"]", &resource.Restriction.Recipient[numRecipient], htmlAttrs)
}
func (resource *Transport) T_InputType(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].type", &resource.Input[numInput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueBase64Binary(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueBase64Binary", &resource.Input[numInput].ValueBase64Binary, htmlAttrs)
}
func (resource *Transport) T_InputValueBoolean(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return BoolInput("input["+strconv.Itoa(numInput)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("input["+strconv.Itoa(numInput)+"].valueBoolean", &resource.Input[numInput].ValueBoolean, htmlAttrs)
}
func (resource *Transport) T_InputValueCanonical(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueCanonical", &resource.Input[numInput].ValueCanonical, htmlAttrs)
}
func (resource *Transport) T_InputValueCode(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeSelect("input["+strconv.Itoa(numInput)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("input["+strconv.Itoa(numInput)+"].valueCode", &resource.Input[numInput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueDate(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return FhirDateInput("input["+strconv.Itoa(numInput)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("input["+strconv.Itoa(numInput)+"].valueDate", &resource.Input[numInput].ValueDate, htmlAttrs)
}
func (resource *Transport) T_InputValueDateTime(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return FhirDateTimeInput("input["+strconv.Itoa(numInput)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("input["+strconv.Itoa(numInput)+"].valueDateTime", &resource.Input[numInput].ValueDateTime, htmlAttrs)
}
func (resource *Transport) T_InputValueDecimal(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Float64Input("input["+strconv.Itoa(numInput)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("input["+strconv.Itoa(numInput)+"].valueDecimal", &resource.Input[numInput].ValueDecimal, htmlAttrs)
}
func (resource *Transport) T_InputValueId(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueId", &resource.Input[numInput].ValueId, htmlAttrs)
}
func (resource *Transport) T_InputValueInstant(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueInstant", &resource.Input[numInput].ValueInstant, htmlAttrs)
}
func (resource *Transport) T_InputValueInteger(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valueInteger", &resource.Input[numInput].ValueInteger, htmlAttrs)
}
func (resource *Transport) T_InputValueInteger64(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return Int64Input("input["+strconv.Itoa(numInput)+"].valueInteger64", nil, htmlAttrs)
	}
	return Int64Input("input["+strconv.Itoa(numInput)+"].valueInteger64", &resource.Input[numInput].ValueInteger64, htmlAttrs)
}
func (resource *Transport) T_InputValueMarkdown(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueMarkdown", &resource.Input[numInput].ValueMarkdown, htmlAttrs)
}
func (resource *Transport) T_InputValueOid(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueOid", &resource.Input[numInput].ValueOid, htmlAttrs)
}
func (resource *Transport) T_InputValuePositiveInt(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valuePositiveInt", &resource.Input[numInput].ValuePositiveInt, htmlAttrs)
}
func (resource *Transport) T_InputValueString(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueString", &resource.Input[numInput].ValueString, htmlAttrs)
}
func (resource *Transport) T_InputValueTime(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueTime", &resource.Input[numInput].ValueTime, htmlAttrs)
}
func (resource *Transport) T_InputValueUnsignedInt(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IntInput("input["+strconv.Itoa(numInput)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("input["+strconv.Itoa(numInput)+"].valueUnsignedInt", &resource.Input[numInput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Transport) T_InputValueUri(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUri", &resource.Input[numInput].ValueUri, htmlAttrs)
}
func (resource *Transport) T_InputValueUrl(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUrl", &resource.Input[numInput].ValueUrl, htmlAttrs)
}
func (resource *Transport) T_InputValueUuid(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return StringInput("input["+strconv.Itoa(numInput)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("input["+strconv.Itoa(numInput)+"].valueUuid", &resource.Input[numInput].ValueUuid, htmlAttrs)
}
func (resource *Transport) T_InputValueAddress(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AddressInput("input["+strconv.Itoa(numInput)+"].valueAddress", nil, htmlAttrs)
	}
	return AddressInput("input["+strconv.Itoa(numInput)+"].valueAddress", &resource.Input[numInput].ValueAddress, htmlAttrs)
}
func (resource *Transport) T_InputValueAge(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AgeInput("input["+strconv.Itoa(numInput)+"].valueAge", nil, htmlAttrs)
	}
	return AgeInput("input["+strconv.Itoa(numInput)+"].valueAge", &resource.Input[numInput].ValueAge, htmlAttrs)
}
func (resource *Transport) T_InputValueAnnotation(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AnnotationTextArea("input["+strconv.Itoa(numInput)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("input["+strconv.Itoa(numInput)+"].valueAnnotation", &resource.Input[numInput].ValueAnnotation, htmlAttrs)
}
func (resource *Transport) T_InputValueAttachment(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AttachmentInput("input["+strconv.Itoa(numInput)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("input["+strconv.Itoa(numInput)+"].valueAttachment", &resource.Input[numInput].ValueAttachment, htmlAttrs)
}
func (resource *Transport) T_InputValueCodeableConcept(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("input["+strconv.Itoa(numInput)+"].valueCodeableConcept", &resource.Input[numInput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueCodeableReference(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodeableReferenceInput("input["+strconv.Itoa(numInput)+"].valueCodeableReference", nil, htmlAttrs)
	}
	return CodeableReferenceInput("input["+strconv.Itoa(numInput)+"].valueCodeableReference", &resource.Input[numInput].ValueCodeableReference, htmlAttrs)
}
func (resource *Transport) T_InputValueCoding(numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CodingSelect("input["+strconv.Itoa(numInput)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("input["+strconv.Itoa(numInput)+"].valueCoding", &resource.Input[numInput].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueContactPoint(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ContactPointInput("input["+strconv.Itoa(numInput)+"].valueContactPoint", nil, htmlAttrs)
	}
	return ContactPointInput("input["+strconv.Itoa(numInput)+"].valueContactPoint", &resource.Input[numInput].ValueContactPoint, htmlAttrs)
}
func (resource *Transport) T_InputValueCount(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return CountInput("input["+strconv.Itoa(numInput)+"].valueCount", nil, htmlAttrs)
	}
	return CountInput("input["+strconv.Itoa(numInput)+"].valueCount", &resource.Input[numInput].ValueCount, htmlAttrs)
}
func (resource *Transport) T_InputValueDistance(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DistanceInput("input["+strconv.Itoa(numInput)+"].valueDistance", nil, htmlAttrs)
	}
	return DistanceInput("input["+strconv.Itoa(numInput)+"].valueDistance", &resource.Input[numInput].ValueDistance, htmlAttrs)
}
func (resource *Transport) T_InputValueDuration(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DurationInput("input["+strconv.Itoa(numInput)+"].valueDuration", nil, htmlAttrs)
	}
	return DurationInput("input["+strconv.Itoa(numInput)+"].valueDuration", &resource.Input[numInput].ValueDuration, htmlAttrs)
}
func (resource *Transport) T_InputValueHumanName(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return HumanNameInput("input["+strconv.Itoa(numInput)+"].valueHumanName", nil, htmlAttrs)
	}
	return HumanNameInput("input["+strconv.Itoa(numInput)+"].valueHumanName", &resource.Input[numInput].ValueHumanName, htmlAttrs)
}
func (resource *Transport) T_InputValueIdentifier(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return IdentifierInput("input["+strconv.Itoa(numInput)+"].valueIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("input["+strconv.Itoa(numInput)+"].valueIdentifier", &resource.Input[numInput].ValueIdentifier, htmlAttrs)
}
func (resource *Transport) T_InputValueMoney(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return MoneyInput("input["+strconv.Itoa(numInput)+"].valueMoney", nil, htmlAttrs)
	}
	return MoneyInput("input["+strconv.Itoa(numInput)+"].valueMoney", &resource.Input[numInput].ValueMoney, htmlAttrs)
}
func (resource *Transport) T_InputValuePeriod(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return PeriodInput("input["+strconv.Itoa(numInput)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("input["+strconv.Itoa(numInput)+"].valuePeriod", &resource.Input[numInput].ValuePeriod, htmlAttrs)
}
func (resource *Transport) T_InputValueQuantity(numInput int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return QuantityInput("input["+strconv.Itoa(numInput)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("input["+strconv.Itoa(numInput)+"].valueQuantity", &resource.Input[numInput].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_InputValueRange(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return RangeInput("input["+strconv.Itoa(numInput)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("input["+strconv.Itoa(numInput)+"].valueRange", &resource.Input[numInput].ValueRange, htmlAttrs)
}
func (resource *Transport) T_InputValueRatio(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return RatioInput("input["+strconv.Itoa(numInput)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("input["+strconv.Itoa(numInput)+"].valueRatio", &resource.Input[numInput].ValueRatio, htmlAttrs)
}
func (resource *Transport) T_InputValueRatioRange(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return RatioRangeInput("input["+strconv.Itoa(numInput)+"].valueRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("input["+strconv.Itoa(numInput)+"].valueRatioRange", &resource.Input[numInput].ValueRatioRange, htmlAttrs)
}
func (resource *Transport) T_InputValueReference(frs []FhirResource, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ReferenceInput(frs, "input["+strconv.Itoa(numInput)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "input["+strconv.Itoa(numInput)+"].valueReference", &resource.Input[numInput].ValueReference, htmlAttrs)
}
func (resource *Transport) T_InputValueSampledData(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return SampledDataInput("input["+strconv.Itoa(numInput)+"].valueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("input["+strconv.Itoa(numInput)+"].valueSampledData", &resource.Input[numInput].ValueSampledData, htmlAttrs)
}
func (resource *Transport) T_InputValueSignature(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return SignatureInput("input["+strconv.Itoa(numInput)+"].valueSignature", nil, htmlAttrs)
	}
	return SignatureInput("input["+strconv.Itoa(numInput)+"].valueSignature", &resource.Input[numInput].ValueSignature, htmlAttrs)
}
func (resource *Transport) T_InputValueTiming(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return TimingInput("input["+strconv.Itoa(numInput)+"].valueTiming", nil, htmlAttrs)
	}
	return TimingInput("input["+strconv.Itoa(numInput)+"].valueTiming", &resource.Input[numInput].ValueTiming, htmlAttrs)
}
func (resource *Transport) T_InputValueContactDetail(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ContactDetailInput("input["+strconv.Itoa(numInput)+"].valueContactDetail", nil, htmlAttrs)
	}
	return ContactDetailInput("input["+strconv.Itoa(numInput)+"].valueContactDetail", &resource.Input[numInput].ValueContactDetail, htmlAttrs)
}
func (resource *Transport) T_InputValueDataRequirement(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DataRequirementInput("input["+strconv.Itoa(numInput)+"].valueDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("input["+strconv.Itoa(numInput)+"].valueDataRequirement", &resource.Input[numInput].ValueDataRequirement, htmlAttrs)
}
func (resource *Transport) T_InputValueExpression(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ExpressionInput("input["+strconv.Itoa(numInput)+"].valueExpression", nil, htmlAttrs)
	}
	return ExpressionInput("input["+strconv.Itoa(numInput)+"].valueExpression", &resource.Input[numInput].ValueExpression, htmlAttrs)
}
func (resource *Transport) T_InputValueParameterDefinition(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ParameterDefinitionInput("input["+strconv.Itoa(numInput)+"].valueParameterDefinition", nil, htmlAttrs)
	}
	return ParameterDefinitionInput("input["+strconv.Itoa(numInput)+"].valueParameterDefinition", &resource.Input[numInput].ValueParameterDefinition, htmlAttrs)
}
func (resource *Transport) T_InputValueRelatedArtifact(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return RelatedArtifactInput("input["+strconv.Itoa(numInput)+"].valueRelatedArtifact", nil, htmlAttrs)
	}
	return RelatedArtifactInput("input["+strconv.Itoa(numInput)+"].valueRelatedArtifact", &resource.Input[numInput].ValueRelatedArtifact, htmlAttrs)
}
func (resource *Transport) T_InputValueTriggerDefinition(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return TriggerDefinitionInput("input["+strconv.Itoa(numInput)+"].valueTriggerDefinition", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("input["+strconv.Itoa(numInput)+"].valueTriggerDefinition", &resource.Input[numInput].ValueTriggerDefinition, htmlAttrs)
}
func (resource *Transport) T_InputValueUsageContext(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return UsageContextInput("input["+strconv.Itoa(numInput)+"].valueUsageContext", nil, htmlAttrs)
	}
	return UsageContextInput("input["+strconv.Itoa(numInput)+"].valueUsageContext", &resource.Input[numInput].ValueUsageContext, htmlAttrs)
}
func (resource *Transport) T_InputValueAvailability(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return AvailabilityInput("input["+strconv.Itoa(numInput)+"].valueAvailability", nil, htmlAttrs)
	}
	return AvailabilityInput("input["+strconv.Itoa(numInput)+"].valueAvailability", &resource.Input[numInput].ValueAvailability, htmlAttrs)
}
func (resource *Transport) T_InputValueExtendedContactDetail(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return ExtendedContactDetailInput("input["+strconv.Itoa(numInput)+"].valueExtendedContactDetail", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("input["+strconv.Itoa(numInput)+"].valueExtendedContactDetail", &resource.Input[numInput].ValueExtendedContactDetail, htmlAttrs)
}
func (resource *Transport) T_InputValueDosage(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return DosageInput("input["+strconv.Itoa(numInput)+"].valueDosage", nil, htmlAttrs)
	}
	return DosageInput("input["+strconv.Itoa(numInput)+"].valueDosage", &resource.Input[numInput].ValueDosage, htmlAttrs)
}
func (resource *Transport) T_InputValueMeta(numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInput >= len(resource.Input) {
		return MetaInput("input["+strconv.Itoa(numInput)+"].valueMeta", nil, htmlAttrs)
	}
	return MetaInput("input["+strconv.Itoa(numInput)+"].valueMeta", &resource.Input[numInput].ValueMeta, htmlAttrs)
}
func (resource *Transport) T_OutputType(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].type", &resource.Output[numOutput].Type, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueBase64Binary(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueBase64Binary", &resource.Output[numOutput].ValueBase64Binary, htmlAttrs)
}
func (resource *Transport) T_OutputValueBoolean(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return BoolInput("output["+strconv.Itoa(numOutput)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("output["+strconv.Itoa(numOutput)+"].valueBoolean", &resource.Output[numOutput].ValueBoolean, htmlAttrs)
}
func (resource *Transport) T_OutputValueCanonical(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueCanonical", &resource.Output[numOutput].ValueCanonical, htmlAttrs)
}
func (resource *Transport) T_OutputValueCode(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeSelect("output["+strconv.Itoa(numOutput)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("output["+strconv.Itoa(numOutput)+"].valueCode", &resource.Output[numOutput].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueDate(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return FhirDateInput("output["+strconv.Itoa(numOutput)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("output["+strconv.Itoa(numOutput)+"].valueDate", &resource.Output[numOutput].ValueDate, htmlAttrs)
}
func (resource *Transport) T_OutputValueDateTime(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return FhirDateTimeInput("output["+strconv.Itoa(numOutput)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("output["+strconv.Itoa(numOutput)+"].valueDateTime", &resource.Output[numOutput].ValueDateTime, htmlAttrs)
}
func (resource *Transport) T_OutputValueDecimal(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Float64Input("output["+strconv.Itoa(numOutput)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("output["+strconv.Itoa(numOutput)+"].valueDecimal", &resource.Output[numOutput].ValueDecimal, htmlAttrs)
}
func (resource *Transport) T_OutputValueId(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueId", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueId", &resource.Output[numOutput].ValueId, htmlAttrs)
}
func (resource *Transport) T_OutputValueInstant(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueInstant", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueInstant", &resource.Output[numOutput].ValueInstant, htmlAttrs)
}
func (resource *Transport) T_OutputValueInteger(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valueInteger", &resource.Output[numOutput].ValueInteger, htmlAttrs)
}
func (resource *Transport) T_OutputValueInteger64(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return Int64Input("output["+strconv.Itoa(numOutput)+"].valueInteger64", nil, htmlAttrs)
	}
	return Int64Input("output["+strconv.Itoa(numOutput)+"].valueInteger64", &resource.Output[numOutput].ValueInteger64, htmlAttrs)
}
func (resource *Transport) T_OutputValueMarkdown(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueMarkdown", &resource.Output[numOutput].ValueMarkdown, htmlAttrs)
}
func (resource *Transport) T_OutputValueOid(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueOid", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueOid", &resource.Output[numOutput].ValueOid, htmlAttrs)
}
func (resource *Transport) T_OutputValuePositiveInt(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valuePositiveInt", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valuePositiveInt", &resource.Output[numOutput].ValuePositiveInt, htmlAttrs)
}
func (resource *Transport) T_OutputValueString(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueString", &resource.Output[numOutput].ValueString, htmlAttrs)
}
func (resource *Transport) T_OutputValueTime(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueTime", &resource.Output[numOutput].ValueTime, htmlAttrs)
}
func (resource *Transport) T_OutputValueUnsignedInt(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IntInput("output["+strconv.Itoa(numOutput)+"].valueUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("output["+strconv.Itoa(numOutput)+"].valueUnsignedInt", &resource.Output[numOutput].ValueUnsignedInt, htmlAttrs)
}
func (resource *Transport) T_OutputValueUri(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUri", &resource.Output[numOutput].ValueUri, htmlAttrs)
}
func (resource *Transport) T_OutputValueUrl(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUrl", &resource.Output[numOutput].ValueUrl, htmlAttrs)
}
func (resource *Transport) T_OutputValueUuid(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return StringInput("output["+strconv.Itoa(numOutput)+"].valueUuid", nil, htmlAttrs)
	}
	return StringInput("output["+strconv.Itoa(numOutput)+"].valueUuid", &resource.Output[numOutput].ValueUuid, htmlAttrs)
}
func (resource *Transport) T_OutputValueAddress(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AddressInput("output["+strconv.Itoa(numOutput)+"].valueAddress", nil, htmlAttrs)
	}
	return AddressInput("output["+strconv.Itoa(numOutput)+"].valueAddress", &resource.Output[numOutput].ValueAddress, htmlAttrs)
}
func (resource *Transport) T_OutputValueAge(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AgeInput("output["+strconv.Itoa(numOutput)+"].valueAge", nil, htmlAttrs)
	}
	return AgeInput("output["+strconv.Itoa(numOutput)+"].valueAge", &resource.Output[numOutput].ValueAge, htmlAttrs)
}
func (resource *Transport) T_OutputValueAnnotation(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AnnotationTextArea("output["+strconv.Itoa(numOutput)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("output["+strconv.Itoa(numOutput)+"].valueAnnotation", &resource.Output[numOutput].ValueAnnotation, htmlAttrs)
}
func (resource *Transport) T_OutputValueAttachment(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AttachmentInput("output["+strconv.Itoa(numOutput)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("output["+strconv.Itoa(numOutput)+"].valueAttachment", &resource.Output[numOutput].ValueAttachment, htmlAttrs)
}
func (resource *Transport) T_OutputValueCodeableConcept(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("output["+strconv.Itoa(numOutput)+"].valueCodeableConcept", &resource.Output[numOutput].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueCodeableReference(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodeableReferenceInput("output["+strconv.Itoa(numOutput)+"].valueCodeableReference", nil, htmlAttrs)
	}
	return CodeableReferenceInput("output["+strconv.Itoa(numOutput)+"].valueCodeableReference", &resource.Output[numOutput].ValueCodeableReference, htmlAttrs)
}
func (resource *Transport) T_OutputValueCoding(numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CodingSelect("output["+strconv.Itoa(numOutput)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("output["+strconv.Itoa(numOutput)+"].valueCoding", &resource.Output[numOutput].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueContactPoint(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ContactPointInput("output["+strconv.Itoa(numOutput)+"].valueContactPoint", nil, htmlAttrs)
	}
	return ContactPointInput("output["+strconv.Itoa(numOutput)+"].valueContactPoint", &resource.Output[numOutput].ValueContactPoint, htmlAttrs)
}
func (resource *Transport) T_OutputValueCount(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return CountInput("output["+strconv.Itoa(numOutput)+"].valueCount", nil, htmlAttrs)
	}
	return CountInput("output["+strconv.Itoa(numOutput)+"].valueCount", &resource.Output[numOutput].ValueCount, htmlAttrs)
}
func (resource *Transport) T_OutputValueDistance(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DistanceInput("output["+strconv.Itoa(numOutput)+"].valueDistance", nil, htmlAttrs)
	}
	return DistanceInput("output["+strconv.Itoa(numOutput)+"].valueDistance", &resource.Output[numOutput].ValueDistance, htmlAttrs)
}
func (resource *Transport) T_OutputValueDuration(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DurationInput("output["+strconv.Itoa(numOutput)+"].valueDuration", nil, htmlAttrs)
	}
	return DurationInput("output["+strconv.Itoa(numOutput)+"].valueDuration", &resource.Output[numOutput].ValueDuration, htmlAttrs)
}
func (resource *Transport) T_OutputValueHumanName(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return HumanNameInput("output["+strconv.Itoa(numOutput)+"].valueHumanName", nil, htmlAttrs)
	}
	return HumanNameInput("output["+strconv.Itoa(numOutput)+"].valueHumanName", &resource.Output[numOutput].ValueHumanName, htmlAttrs)
}
func (resource *Transport) T_OutputValueIdentifier(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return IdentifierInput("output["+strconv.Itoa(numOutput)+"].valueIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("output["+strconv.Itoa(numOutput)+"].valueIdentifier", &resource.Output[numOutput].ValueIdentifier, htmlAttrs)
}
func (resource *Transport) T_OutputValueMoney(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return MoneyInput("output["+strconv.Itoa(numOutput)+"].valueMoney", nil, htmlAttrs)
	}
	return MoneyInput("output["+strconv.Itoa(numOutput)+"].valueMoney", &resource.Output[numOutput].ValueMoney, htmlAttrs)
}
func (resource *Transport) T_OutputValuePeriod(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return PeriodInput("output["+strconv.Itoa(numOutput)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("output["+strconv.Itoa(numOutput)+"].valuePeriod", &resource.Output[numOutput].ValuePeriod, htmlAttrs)
}
func (resource *Transport) T_OutputValueQuantity(numOutput int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return QuantityInput("output["+strconv.Itoa(numOutput)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("output["+strconv.Itoa(numOutput)+"].valueQuantity", &resource.Output[numOutput].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Transport) T_OutputValueRange(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return RangeInput("output["+strconv.Itoa(numOutput)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("output["+strconv.Itoa(numOutput)+"].valueRange", &resource.Output[numOutput].ValueRange, htmlAttrs)
}
func (resource *Transport) T_OutputValueRatio(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return RatioInput("output["+strconv.Itoa(numOutput)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("output["+strconv.Itoa(numOutput)+"].valueRatio", &resource.Output[numOutput].ValueRatio, htmlAttrs)
}
func (resource *Transport) T_OutputValueRatioRange(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return RatioRangeInput("output["+strconv.Itoa(numOutput)+"].valueRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("output["+strconv.Itoa(numOutput)+"].valueRatioRange", &resource.Output[numOutput].ValueRatioRange, htmlAttrs)
}
func (resource *Transport) T_OutputValueReference(frs []FhirResource, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ReferenceInput(frs, "output["+strconv.Itoa(numOutput)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "output["+strconv.Itoa(numOutput)+"].valueReference", &resource.Output[numOutput].ValueReference, htmlAttrs)
}
func (resource *Transport) T_OutputValueSampledData(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return SampledDataInput("output["+strconv.Itoa(numOutput)+"].valueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("output["+strconv.Itoa(numOutput)+"].valueSampledData", &resource.Output[numOutput].ValueSampledData, htmlAttrs)
}
func (resource *Transport) T_OutputValueSignature(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return SignatureInput("output["+strconv.Itoa(numOutput)+"].valueSignature", nil, htmlAttrs)
	}
	return SignatureInput("output["+strconv.Itoa(numOutput)+"].valueSignature", &resource.Output[numOutput].ValueSignature, htmlAttrs)
}
func (resource *Transport) T_OutputValueTiming(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return TimingInput("output["+strconv.Itoa(numOutput)+"].valueTiming", nil, htmlAttrs)
	}
	return TimingInput("output["+strconv.Itoa(numOutput)+"].valueTiming", &resource.Output[numOutput].ValueTiming, htmlAttrs)
}
func (resource *Transport) T_OutputValueContactDetail(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ContactDetailInput("output["+strconv.Itoa(numOutput)+"].valueContactDetail", nil, htmlAttrs)
	}
	return ContactDetailInput("output["+strconv.Itoa(numOutput)+"].valueContactDetail", &resource.Output[numOutput].ValueContactDetail, htmlAttrs)
}
func (resource *Transport) T_OutputValueDataRequirement(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DataRequirementInput("output["+strconv.Itoa(numOutput)+"].valueDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("output["+strconv.Itoa(numOutput)+"].valueDataRequirement", &resource.Output[numOutput].ValueDataRequirement, htmlAttrs)
}
func (resource *Transport) T_OutputValueExpression(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ExpressionInput("output["+strconv.Itoa(numOutput)+"].valueExpression", nil, htmlAttrs)
	}
	return ExpressionInput("output["+strconv.Itoa(numOutput)+"].valueExpression", &resource.Output[numOutput].ValueExpression, htmlAttrs)
}
func (resource *Transport) T_OutputValueParameterDefinition(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ParameterDefinitionInput("output["+strconv.Itoa(numOutput)+"].valueParameterDefinition", nil, htmlAttrs)
	}
	return ParameterDefinitionInput("output["+strconv.Itoa(numOutput)+"].valueParameterDefinition", &resource.Output[numOutput].ValueParameterDefinition, htmlAttrs)
}
func (resource *Transport) T_OutputValueRelatedArtifact(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return RelatedArtifactInput("output["+strconv.Itoa(numOutput)+"].valueRelatedArtifact", nil, htmlAttrs)
	}
	return RelatedArtifactInput("output["+strconv.Itoa(numOutput)+"].valueRelatedArtifact", &resource.Output[numOutput].ValueRelatedArtifact, htmlAttrs)
}
func (resource *Transport) T_OutputValueTriggerDefinition(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return TriggerDefinitionInput("output["+strconv.Itoa(numOutput)+"].valueTriggerDefinition", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("output["+strconv.Itoa(numOutput)+"].valueTriggerDefinition", &resource.Output[numOutput].ValueTriggerDefinition, htmlAttrs)
}
func (resource *Transport) T_OutputValueUsageContext(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return UsageContextInput("output["+strconv.Itoa(numOutput)+"].valueUsageContext", nil, htmlAttrs)
	}
	return UsageContextInput("output["+strconv.Itoa(numOutput)+"].valueUsageContext", &resource.Output[numOutput].ValueUsageContext, htmlAttrs)
}
func (resource *Transport) T_OutputValueAvailability(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return AvailabilityInput("output["+strconv.Itoa(numOutput)+"].valueAvailability", nil, htmlAttrs)
	}
	return AvailabilityInput("output["+strconv.Itoa(numOutput)+"].valueAvailability", &resource.Output[numOutput].ValueAvailability, htmlAttrs)
}
func (resource *Transport) T_OutputValueExtendedContactDetail(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return ExtendedContactDetailInput("output["+strconv.Itoa(numOutput)+"].valueExtendedContactDetail", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("output["+strconv.Itoa(numOutput)+"].valueExtendedContactDetail", &resource.Output[numOutput].ValueExtendedContactDetail, htmlAttrs)
}
func (resource *Transport) T_OutputValueDosage(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return DosageInput("output["+strconv.Itoa(numOutput)+"].valueDosage", nil, htmlAttrs)
	}
	return DosageInput("output["+strconv.Itoa(numOutput)+"].valueDosage", &resource.Output[numOutput].ValueDosage, htmlAttrs)
}
func (resource *Transport) T_OutputValueMeta(numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutput >= len(resource.Output) {
		return MetaInput("output["+strconv.Itoa(numOutput)+"].valueMeta", nil, htmlAttrs)
	}
	return MetaInput("output["+strconv.Itoa(numOutput)+"].valueMeta", &resource.Output[numOutput].ValueMeta, htmlAttrs)
}
