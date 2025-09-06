package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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
	CompletionTime        *string               `json:"completionTime,omitempty"`
	AuthoredOn            *string               `json:"authoredOn,omitempty"`
	LastModified          *string               `json:"lastModified,omitempty"`
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
	ValueDate                  string                `json:"valueDate"`
	ValueDateTime              string                `json:"valueDateTime"`
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
	ValueDate                  string                `json:"valueDate"`
	ValueDateTime              string                `json:"valueDateTime"`
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

func (resource *Transport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Transport.Id", nil)
	}
	return StringInput("Transport.Id", resource.Id)
}
func (resource *Transport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Transport.ImplicitRules", nil)
	}
	return StringInput("Transport.ImplicitRules", resource.ImplicitRules)
}
func (resource *Transport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Transport.Language", nil, optionsValueSet)
	}
	return CodeSelect("Transport.Language", resource.Language, optionsValueSet)
}
func (resource *Transport) T_InstantiatesCanonical() templ.Component {

	if resource == nil {
		return StringInput("Transport.InstantiatesCanonical", nil)
	}
	return StringInput("Transport.InstantiatesCanonical", resource.InstantiatesCanonical)
}
func (resource *Transport) T_InstantiatesUri() templ.Component {

	if resource == nil {
		return StringInput("Transport.InstantiatesUri", nil)
	}
	return StringInput("Transport.InstantiatesUri", resource.InstantiatesUri)
}
func (resource *Transport) T_Status() templ.Component {
	optionsValueSet := VSTransport_status

	if resource == nil {
		return CodeSelect("Transport.Status", nil, optionsValueSet)
	}
	return CodeSelect("Transport.Status", resource.Status, optionsValueSet)
}
func (resource *Transport) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Transport.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Transport.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Transport) T_Intent() templ.Component {
	optionsValueSet := VSTransport_intent

	if resource == nil {
		return CodeSelect("Transport.Intent", nil, optionsValueSet)
	}
	return CodeSelect("Transport.Intent", &resource.Intent, optionsValueSet)
}
func (resource *Transport) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Transport.Priority", nil, optionsValueSet)
	}
	return CodeSelect("Transport.Priority", resource.Priority, optionsValueSet)
}
func (resource *Transport) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Transport.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Transport.Code", resource.Code, optionsValueSet)
}
func (resource *Transport) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Transport.Description", nil)
	}
	return StringInput("Transport.Description", resource.Description)
}
func (resource *Transport) T_CompletionTime() templ.Component {

	if resource == nil {
		return StringInput("Transport.CompletionTime", nil)
	}
	return StringInput("Transport.CompletionTime", resource.CompletionTime)
}
func (resource *Transport) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("Transport.AuthoredOn", nil)
	}
	return StringInput("Transport.AuthoredOn", resource.AuthoredOn)
}
func (resource *Transport) T_LastModified() templ.Component {

	if resource == nil {
		return StringInput("Transport.LastModified", nil)
	}
	return StringInput("Transport.LastModified", resource.LastModified)
}
func (resource *Transport) T_PerformerType(numPerformerType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PerformerType) >= numPerformerType {
		return CodeableConceptSelect("Transport.PerformerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Transport.PerformerType["+strconv.Itoa(numPerformerType)+"]", &resource.PerformerType[numPerformerType], optionsValueSet)
}
func (resource *Transport) T_RestrictionId() templ.Component {

	if resource == nil {
		return StringInput("Transport.Restriction.Id", nil)
	}
	return StringInput("Transport.Restriction.Id", resource.Restriction.Id)
}
func (resource *Transport) T_RestrictionRepetitions() templ.Component {

	if resource == nil {
		return IntInput("Transport.Restriction.Repetitions", nil)
	}
	return IntInput("Transport.Restriction.Repetitions", resource.Restriction.Repetitions)
}
func (resource *Transport) T_InputId(numInput int) templ.Component {

	if resource == nil || len(resource.Input) >= numInput {
		return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].Id", nil)
	}
	return StringInput("Transport.Input["+strconv.Itoa(numInput)+"].Id", resource.Input[numInput].Id)
}
func (resource *Transport) T_InputType(numInput int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Input) >= numInput {
		return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Transport.Input["+strconv.Itoa(numInput)+"].Type", &resource.Input[numInput].Type, optionsValueSet)
}
func (resource *Transport) T_OutputId(numOutput int) templ.Component {

	if resource == nil || len(resource.Output) >= numOutput {
		return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].Id", nil)
	}
	return StringInput("Transport.Output["+strconv.Itoa(numOutput)+"].Id", resource.Output[numOutput].Id)
}
func (resource *Transport) T_OutputType(numOutput int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Output) >= numOutput {
		return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Transport.Output["+strconv.Itoa(numOutput)+"].Type", &resource.Output[numOutput].Type, optionsValueSet)
}
