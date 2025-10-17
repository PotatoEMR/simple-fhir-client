package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicationRequest
type MedicationRequest struct {
	Id                        *string                           `json:"id,omitempty"`
	Meta                      *Meta                             `json:"meta,omitempty"`
	ImplicitRules             *string                           `json:"implicitRules,omitempty"`
	Language                  *string                           `json:"language,omitempty"`
	Text                      *Narrative                        `json:"text,omitempty"`
	Contained                 []Resource                        `json:"contained,omitempty"`
	Extension                 []Extension                       `json:"extension,omitempty"`
	ModifierExtension         []Extension                       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                      `json:"identifier,omitempty"`
	Status                    string                            `json:"status"`
	StatusReason              *CodeableConcept                  `json:"statusReason,omitempty"`
	Intent                    string                            `json:"intent"`
	Category                  []CodeableConcept                 `json:"category,omitempty"`
	Priority                  *string                           `json:"priority,omitempty"`
	DoNotPerform              *bool                             `json:"doNotPerform,omitempty"`
	ReportedBoolean           *bool                             `json:"reportedBoolean,omitempty"`
	ReportedReference         *Reference                        `json:"reportedReference,omitempty"`
	MedicationCodeableConcept CodeableConcept                   `json:"medicationCodeableConcept"`
	MedicationReference       Reference                         `json:"medicationReference"`
	Subject                   Reference                         `json:"subject"`
	Encounter                 *Reference                        `json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `json:"supportingInformation,omitempty"`
	AuthoredOn                *FhirDateTime                     `json:"authoredOn,omitempty"`
	Requester                 *Reference                        `json:"requester,omitempty"`
	Performer                 *Reference                        `json:"performer,omitempty"`
	PerformerType             *CodeableConcept                  `json:"performerType,omitempty"`
	Recorder                  *Reference                        `json:"recorder,omitempty"`
	ReasonCode                []CodeableConcept                 `json:"reasonCode,omitempty"`
	ReasonReference           []Reference                       `json:"reasonReference,omitempty"`
	InstantiatesCanonical     []string                          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri           []string                          `json:"instantiatesUri,omitempty"`
	BasedOn                   []Reference                       `json:"basedOn,omitempty"`
	GroupIdentifier           *Identifier                       `json:"groupIdentifier,omitempty"`
	CourseOfTherapyType       *CodeableConcept                  `json:"courseOfTherapyType,omitempty"`
	Insurance                 []Reference                       `json:"insurance,omitempty"`
	Note                      []Annotation                      `json:"note,omitempty"`
	DosageInstruction         []Dosage                          `json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `json:"substitution,omitempty"`
	PriorPrescription         *Reference                        `json:"priorPrescription,omitempty"`
	DetectedIssue             []Reference                       `json:"detectedIssue,omitempty"`
	EventHistory              []Reference                       `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationRequest
type MedicationRequestDispenseRequest struct {
	Id                     *string                                      `json:"id,omitempty"`
	Extension              []Extension                                  `json:"extension,omitempty"`
	ModifierExtension      []Extension                                  `json:"modifierExtension,omitempty"`
	InitialFill            *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`
	DispenseInterval       *Duration                                    `json:"dispenseInterval,omitempty"`
	ValidityPeriod         *Period                                      `json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *int                                         `json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity                                    `json:"quantity,omitempty"`
	ExpectedSupplyDuration *Duration                                    `json:"expectedSupplyDuration,omitempty"`
	Performer              *Reference                                   `json:"performer,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationRequest
type MedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	Duration          *Duration   `json:"duration,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationRequest
type MedicationRequestSubstitution struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AllowedBoolean         bool             `json:"allowedBoolean"`
	AllowedCodeableConcept CodeableConcept  `json:"allowedCodeableConcept"`
	Reason                 *CodeableConcept `json:"reason,omitempty"`
}

type OtherMedicationRequest MedicationRequest

// struct -> json, automatically add resourceType=Patient
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// json -> struct, first reject if resourceType != MedicationRequest
func (r *MedicationRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MedicationRequest" {
		return errors.New("resourceType not MedicationRequest")
	}
	return json.Unmarshal(data, (*OtherMedicationRequest)(r))
}

func (r MedicationRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicationRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicationRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationrequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationrequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *MedicationRequest) T_ReportedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("reportedBoolean", nil, htmlAttrs)
	}
	return BoolInput("reportedBoolean", resource.ReportedBoolean, htmlAttrs)
}
func (resource *MedicationRequest) T_ReportedReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "reportedReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reportedReference", resource.ReportedReference, htmlAttrs)
}
func (resource *MedicationRequest) T_MedicationCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicationCodeableConcept", &resource.MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_MedicationReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "medicationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "medicationReference", &resource.MedicationReference, htmlAttrs)
}
func (resource *MedicationRequest) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationRequest) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *MedicationRequest) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *MedicationRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *MedicationRequest) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *MedicationRequest) T_Performer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "performer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer", resource.Performer, htmlAttrs)
}
func (resource *MedicationRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Recorder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recorder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recorder", resource.Recorder, htmlAttrs)
}
func (resource *MedicationRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *MedicationRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *MedicationRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *MedicationRequest) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *MedicationRequest) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
}
func (resource *MedicationRequest) T_CourseOfTherapyType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("courseOfTherapyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("courseOfTherapyType", resource.CourseOfTherapyType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Insurance(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", &resource.Insurance[numInsurance], htmlAttrs)
}
func (resource *MedicationRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationRequest) T_DosageInstruction(numDosageInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosageInstruction >= len(resource.DosageInstruction) {
		return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", &resource.DosageInstruction[numDosageInstruction], htmlAttrs)
}
func (resource *MedicationRequest) T_PriorPrescription(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "priorPrescription", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "priorPrescription", resource.PriorPrescription, htmlAttrs)
}
func (resource *MedicationRequest) T_DetectedIssue(frs []FhirResource, numDetectedIssue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetectedIssue >= len(resource.DetectedIssue) {
		return ReferenceInput(frs, "detectedIssue["+strconv.Itoa(numDetectedIssue)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detectedIssue["+strconv.Itoa(numDetectedIssue)+"]", &resource.DetectedIssue[numDetectedIssue], htmlAttrs)
}
func (resource *MedicationRequest) T_EventHistory(frs []FhirResource, numEventHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventHistory >= len(resource.EventHistory) {
		return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", &resource.EventHistory[numEventHistory], htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestDispenseInterval(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("dispenseRequest.dispenseInterval", nil, htmlAttrs)
	}
	return DurationInput("dispenseRequest.dispenseInterval", resource.DispenseRequest.DispenseInterval, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestValidityPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("dispenseRequest.validityPeriod", nil, htmlAttrs)
	}
	return PeriodInput("dispenseRequest.validityPeriod", resource.DispenseRequest.ValidityPeriod, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestNumberOfRepeatsAllowed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("dispenseRequest.numberOfRepeatsAllowed", nil, htmlAttrs)
	}
	return IntInput("dispenseRequest.numberOfRepeatsAllowed", resource.DispenseRequest.NumberOfRepeatsAllowed, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("dispenseRequest.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("dispenseRequest.quantity", resource.DispenseRequest.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestExpectedSupplyDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("dispenseRequest.expectedSupplyDuration", nil, htmlAttrs)
	}
	return DurationInput("dispenseRequest.expectedSupplyDuration", resource.DispenseRequest.ExpectedSupplyDuration, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestPerformer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "dispenseRequest.performer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "dispenseRequest.performer", resource.DispenseRequest.Performer, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestInitialFillQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("dispenseRequest.initialFill.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("dispenseRequest.initialFill.quantity", resource.DispenseRequest.InitialFill.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestInitialFillDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("dispenseRequest.initialFill.duration", nil, htmlAttrs)
	}
	return DurationInput("dispenseRequest.initialFill.duration", resource.DispenseRequest.InitialFill.Duration, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionAllowedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("substitution.allowedBoolean", nil, htmlAttrs)
	}
	return BoolInput("substitution.allowedBoolean", &resource.Substitution.AllowedBoolean, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionAllowedCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("substitution.allowedCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substitution.allowedCodeableConcept", &resource.Substitution.AllowedCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("substitution.reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substitution.reason", resource.Substitution.Reason, optionsValueSet, htmlAttrs)
}
