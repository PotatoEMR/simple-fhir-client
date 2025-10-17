package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequest struct {
	Id                      *string                            `json:"id,omitempty"`
	Meta                    *Meta                              `json:"meta,omitempty"`
	ImplicitRules           *string                            `json:"implicitRules,omitempty"`
	Language                *string                            `json:"language,omitempty"`
	Text                    *Narrative                         `json:"text,omitempty"`
	Contained               []Resource                         `json:"contained,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                       `json:"identifier,omitempty"`
	InstantiatesCanonical   []string                           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string                           `json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference                        `json:"basedOn,omitempty"`
	Replaces                []Reference                        `json:"replaces,omitempty"`
	Requisition             *Identifier                        `json:"requisition,omitempty"`
	Status                  string                             `json:"status"`
	Intent                  string                             `json:"intent"`
	Category                []CodeableConcept                  `json:"category,omitempty"`
	Priority                *string                            `json:"priority,omitempty"`
	DoNotPerform            *bool                              `json:"doNotPerform,omitempty"`
	Code                    *CodeableReference                 `json:"code,omitempty"`
	OrderDetail             []ServiceRequestOrderDetail        `json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity                          `json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio                             `json:"quantityRatio,omitempty"`
	QuantityRange           *Range                             `json:"quantityRange,omitempty"`
	Subject                 Reference                          `json:"subject"`
	Focus                   []Reference                        `json:"focus,omitempty"`
	Encounter               *Reference                         `json:"encounter,omitempty"`
	OccurrenceDateTime      *FhirDateTime                      `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                            `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                            `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool                              `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept                   `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *FhirDateTime                      `json:"authoredOn,omitempty"`
	Requester               *Reference                         `json:"requester,omitempty"`
	PerformerType           *CodeableConcept                   `json:"performerType,omitempty"`
	Performer               []Reference                        `json:"performer,omitempty"`
	Location                []CodeableReference                `json:"location,omitempty"`
	Reason                  []CodeableReference                `json:"reason,omitempty"`
	Insurance               []Reference                        `json:"insurance,omitempty"`
	SupportingInfo          []CodeableReference                `json:"supportingInfo,omitempty"`
	Specimen                []Reference                        `json:"specimen,omitempty"`
	BodySite                []CodeableConcept                  `json:"bodySite,omitempty"`
	BodyStructure           *Reference                         `json:"bodyStructure,omitempty"`
	Note                    []Annotation                       `json:"note,omitempty"`
	PatientInstruction      []ServiceRequestPatientInstruction `json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference                        `json:"relevantHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestOrderDetail struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	ParameterFocus    *CodeableReference                   `json:"parameterFocus,omitempty"`
	Parameter         []ServiceRequestOrderDetailParameter `json:"parameter"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestOrderDetailParameter struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRatio           Ratio           `json:"valueRatio"`
	ValueRange           Range           `json:"valueRange"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueString          string          `json:"valueString"`
	ValuePeriod          Period          `json:"valuePeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestPatientInstruction struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	InstructionMarkdown  *string     `json:"instructionMarkdown,omitempty"`
	InstructionReference *Reference  `json:"instructionReference,omitempty"`
}

type OtherServiceRequest ServiceRequest

// struct -> json, automatically add resourceType=Patient
func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherServiceRequest: OtherServiceRequest(r),
		ResourceType:        "ServiceRequest",
	})
}

func (r ServiceRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ServiceRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ServiceRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r ServiceRequest) ResourceType() string {
	return "ServiceRequest"
}

func (resource *ServiceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *ServiceRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *ServiceRequest) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *ServiceRequest) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *ServiceRequest) T_Requisition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("requisition", nil, htmlAttrs)
	}
	return IdentifierInput("requisition", resource.Requisition, htmlAttrs)
}
func (resource *ServiceRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ServiceRequest) T_Code(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("code", nil, htmlAttrs)
	}
	return CodeableReferenceInput("code", resource.Code, htmlAttrs)
}
func (resource *ServiceRequest) T_QuantityQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantityQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantityQuantity", resource.QuantityQuantity, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_QuantityRatio(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RatioInput("quantityRatio", nil, htmlAttrs)
	}
	return RatioInput("quantityRatio", resource.QuantityRatio, htmlAttrs)
}
func (resource *ServiceRequest) T_QuantityRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("quantityRange", nil, htmlAttrs)
	}
	return RangeInput("quantityRange", resource.QuantityRange, htmlAttrs)
}
func (resource *ServiceRequest) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *ServiceRequest) T_Focus(frs []FhirResource, numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], htmlAttrs)
}
func (resource *ServiceRequest) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *ServiceRequest) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *ServiceRequest) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *ServiceRequest) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("asNeededBoolean", nil, htmlAttrs)
	}
	return BoolInput("asNeededBoolean", resource.AsNeededBoolean, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("asNeededCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("asNeededCodeableConcept", resource.AsNeededCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *ServiceRequest) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *ServiceRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Performer(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", &resource.Performer[numPerformer], htmlAttrs)
}
func (resource *ServiceRequest) T_Location(numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableReferenceInput("location["+strconv.Itoa(numLocation)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], htmlAttrs)
}
func (resource *ServiceRequest) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *ServiceRequest) T_Insurance(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", &resource.Insurance[numInsurance], htmlAttrs)
}
func (resource *ServiceRequest) T_SupportingInfo(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *ServiceRequest) T_Specimen(frs []FhirResource, numSpecimen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecimen >= len(resource.Specimen) {
		return ReferenceInput(frs, "specimen["+strconv.Itoa(numSpecimen)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "specimen["+strconv.Itoa(numSpecimen)+"]", &resource.Specimen[numSpecimen], htmlAttrs)
}
func (resource *ServiceRequest) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_BodyStructure(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "bodyStructure", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "bodyStructure", resource.BodyStructure, htmlAttrs)
}
func (resource *ServiceRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ServiceRequest) T_RelevantHistory(frs []FhirResource, numRelevantHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelevantHistory >= len(resource.RelevantHistory) {
		return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", &resource.RelevantHistory[numRelevantHistory], htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterFocus(numOrderDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) {
		return CodeableReferenceInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameterFocus", nil, htmlAttrs)
	}
	return CodeableReferenceInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameterFocus", resource.OrderDetail[numOrderDetail].ParameterFocus, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterCode(numOrderDetail int, numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].code", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueQuantity(numOrderDetail int, numParameter int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return QuantityInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueQuantity", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueRatio(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return RatioInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueRatio", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueRatio, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueRange(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return RangeInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueRange", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueRange, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueBoolean(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return BoolInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueCodeableConcept(numOrderDetail int, numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueString(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return StringInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValuePeriod(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return PeriodInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valuePeriod", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValuePeriod, htmlAttrs)
}
func (resource *ServiceRequest) T_PatientInstructionInstructionMarkdown(numPatientInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatientInstruction >= len(resource.PatientInstruction) {
		return StringInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionMarkdown", nil, htmlAttrs)
	}
	return StringInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionMarkdown", resource.PatientInstruction[numPatientInstruction].InstructionMarkdown, htmlAttrs)
}
func (resource *ServiceRequest) T_PatientInstructionInstructionReference(frs []FhirResource, numPatientInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatientInstruction >= len(resource.PatientInstruction) {
		return ReferenceInput(frs, "patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionReference", resource.PatientInstruction[numPatientInstruction].InstructionReference, htmlAttrs)
}
