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

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationRequest
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
	AuthoredOn                *time.Time                        `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationRequest
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

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationRequest
type MedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	Duration          *Duration   `json:"duration,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationRequest
type MedicationRequestSubstitution struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AllowedBoolean         bool             `json:"allowedBoolean"`
	AllowedCodeableConcept CodeableConcept  `json:"allowedCodeableConcept"`
	Reason                 *CodeableConcept `json:"reason,omitempty"`
}

type OtherMedicationRequest MedicationRequest

// on convert struct to json, automatically add resourceType=MedicationRequest
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
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
func (resource *MedicationRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedicationrequest_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedicationrequest_intent

	if resource == nil {
		return CodeSelect("Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *MedicationRequest) T_ReportedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ReportedBoolean", nil, htmlAttrs)
	}
	return BoolInput("ReportedBoolean", resource.ReportedBoolean, htmlAttrs)
}
func (resource *MedicationRequest) T_MedicationCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationCodeableConcept", &resource.MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *MedicationRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PerformerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PerformerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *MedicationRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *MedicationRequest) T_CourseOfTherapyType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CourseOfTherapyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CourseOfTherapyType", resource.CourseOfTherapyType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestNumberOfRepeatsAllowed(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("DispenseRequestNumberOfRepeatsAllowed", nil, htmlAttrs)
	}
	return IntInput("DispenseRequestNumberOfRepeatsAllowed", resource.DispenseRequest.NumberOfRepeatsAllowed, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionAllowedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("SubstitutionAllowedBoolean", nil, htmlAttrs)
	}
	return BoolInput("SubstitutionAllowedBoolean", &resource.Substitution.AllowedBoolean, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionAllowedCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstitutionAllowedCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstitutionAllowedCodeableConcept", &resource.Substitution.AllowedCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_SubstitutionReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstitutionReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstitutionReason", resource.Substitution.Reason, optionsValueSet, htmlAttrs)
}
