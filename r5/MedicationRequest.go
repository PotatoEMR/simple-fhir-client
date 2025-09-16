package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MedicationRequest
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
	BasedOn                   []Reference                       `json:"basedOn,omitempty"`
	PriorPrescription         *Reference                        `json:"priorPrescription,omitempty"`
	GroupIdentifier           *Identifier                       `json:"groupIdentifier,omitempty"`
	Status                    string                            `json:"status"`
	StatusReason              *CodeableConcept                  `json:"statusReason,omitempty"`
	StatusChanged             *FhirDateTime                     `json:"statusChanged,omitempty"`
	Intent                    string                            `json:"intent"`
	Category                  []CodeableConcept                 `json:"category,omitempty"`
	Priority                  *string                           `json:"priority,omitempty"`
	DoNotPerform              *bool                             `json:"doNotPerform,omitempty"`
	Medication                CodeableReference                 `json:"medication"`
	Subject                   Reference                         `json:"subject"`
	InformationSource         []Reference                       `json:"informationSource,omitempty"`
	Encounter                 *Reference                        `json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `json:"supportingInformation,omitempty"`
	AuthoredOn                *FhirDateTime                     `json:"authoredOn,omitempty"`
	Requester                 *Reference                        `json:"requester,omitempty"`
	Reported                  *bool                             `json:"reported,omitempty"`
	PerformerType             *CodeableConcept                  `json:"performerType,omitempty"`
	Performer                 []Reference                       `json:"performer,omitempty"`
	Device                    []CodeableReference               `json:"device,omitempty"`
	Recorder                  *Reference                        `json:"recorder,omitempty"`
	Reason                    []CodeableReference               `json:"reason,omitempty"`
	CourseOfTherapyType       *CodeableConcept                  `json:"courseOfTherapyType,omitempty"`
	Insurance                 []Reference                       `json:"insurance,omitempty"`
	Note                      []Annotation                      `json:"note,omitempty"`
	RenderedDosageInstruction *string                           `json:"renderedDosageInstruction,omitempty"`
	EffectiveDosePeriod       *Period                           `json:"effectiveDosePeriod,omitempty"`
	DosageInstruction         []Dosage                          `json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `json:"substitution,omitempty"`
	EventHistory              []Reference                       `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationRequest
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
	Dispenser              *Reference                                   `json:"dispenser,omitempty"`
	DispenserInstruction   []Annotation                                 `json:"dispenserInstruction,omitempty"`
	DoseAdministrationAid  *CodeableConcept                             `json:"doseAdministrationAid,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationRequest
type MedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	Duration          *Duration   `json:"duration,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationRequest
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
func (resource *MedicationRequest) T_StatusChanged(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("statusChanged", nil, htmlAttrs)
	}
	return DateTimeInput("statusChanged", resource.StatusChanged, htmlAttrs)
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
func (resource *MedicationRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *MedicationRequest) T_Reported(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("reported", nil, htmlAttrs)
	}
	return BoolInput("reported", resource.Reported, htmlAttrs)
}
func (resource *MedicationRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_CourseOfTherapyType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("courseOfTherapyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("courseOfTherapyType", resource.CourseOfTherapyType, optionsValueSet, htmlAttrs)
}
func (resource *MedicationRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationRequest) T_RenderedDosageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("renderedDosageInstruction", nil, htmlAttrs)
	}
	return StringInput("renderedDosageInstruction", resource.RenderedDosageInstruction, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestNumberOfRepeatsAllowed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("dispenseRequest.numberOfRepeatsAllowed", nil, htmlAttrs)
	}
	return IntInput("dispenseRequest.numberOfRepeatsAllowed", resource.DispenseRequest.NumberOfRepeatsAllowed, htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestDispenserInstruction(numDispenserInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDispenserInstruction >= len(resource.DispenseRequest.DispenserInstruction) {
		return AnnotationTextArea("dispenseRequest.dispenserInstruction["+strconv.Itoa(numDispenserInstruction)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("dispenseRequest.dispenserInstruction["+strconv.Itoa(numDispenserInstruction)+"]", &resource.DispenseRequest.DispenserInstruction[numDispenserInstruction], htmlAttrs)
}
func (resource *MedicationRequest) T_DispenseRequestDoseAdministrationAid(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dispenseRequest.doseAdministrationAid", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dispenseRequest.doseAdministrationAid", resource.DispenseRequest.DoseAdministrationAid, optionsValueSet, htmlAttrs)
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
