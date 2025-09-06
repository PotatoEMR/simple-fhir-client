package r5

//generated with command go run ./bultaoreune -nodownload
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
	StatusChanged             *string                           `json:"statusChanged,omitempty"`
	Intent                    string                            `json:"intent"`
	Category                  []CodeableConcept                 `json:"category,omitempty"`
	Priority                  *string                           `json:"priority,omitempty"`
	DoNotPerform              *bool                             `json:"doNotPerform,omitempty"`
	Medication                CodeableReference                 `json:"medication"`
	Subject                   Reference                         `json:"subject"`
	InformationSource         []Reference                       `json:"informationSource,omitempty"`
	Encounter                 *Reference                        `json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `json:"supportingInformation,omitempty"`
	AuthoredOn                *string                           `json:"authoredOn,omitempty"`
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

func (resource *MedicationRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.Id", nil)
	}
	return StringInput("MedicationRequest.Id", resource.Id)
}
func (resource *MedicationRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.ImplicitRules", nil)
	}
	return StringInput("MedicationRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicationRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicationRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicationRequest.Language", resource.Language, optionsValueSet)
}
func (resource *MedicationRequest) T_Status() templ.Component {
	optionsValueSet := VSMedicationrequest_status

	if resource == nil {
		return CodeSelect("MedicationRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("MedicationRequest.Status", &resource.Status, optionsValueSet)
}
func (resource *MedicationRequest) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationRequest.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *MedicationRequest) T_StatusChanged() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.StatusChanged", nil)
	}
	return StringInput("MedicationRequest.StatusChanged", resource.StatusChanged)
}
func (resource *MedicationRequest) T_Intent() templ.Component {
	optionsValueSet := VSMedicationrequest_intent

	if resource == nil {
		return CodeSelect("MedicationRequest.Intent", nil, optionsValueSet)
	}
	return CodeSelect("MedicationRequest.Intent", &resource.Intent, optionsValueSet)
}
func (resource *MedicationRequest) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("MedicationRequest.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *MedicationRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("MedicationRequest.Priority", nil, optionsValueSet)
	}
	return CodeSelect("MedicationRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *MedicationRequest) T_DoNotPerform() templ.Component {

	if resource == nil {
		return BoolInput("MedicationRequest.DoNotPerform", nil)
	}
	return BoolInput("MedicationRequest.DoNotPerform", resource.DoNotPerform)
}
func (resource *MedicationRequest) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.AuthoredOn", nil)
	}
	return StringInput("MedicationRequest.AuthoredOn", resource.AuthoredOn)
}
func (resource *MedicationRequest) T_Reported() templ.Component {

	if resource == nil {
		return BoolInput("MedicationRequest.Reported", nil)
	}
	return BoolInput("MedicationRequest.Reported", resource.Reported)
}
func (resource *MedicationRequest) T_PerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationRequest.PerformerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.PerformerType", resource.PerformerType, optionsValueSet)
}
func (resource *MedicationRequest) T_CourseOfTherapyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationRequest.CourseOfTherapyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.CourseOfTherapyType", resource.CourseOfTherapyType, optionsValueSet)
}
func (resource *MedicationRequest) T_RenderedDosageInstruction() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.RenderedDosageInstruction", nil)
	}
	return StringInput("MedicationRequest.RenderedDosageInstruction", resource.RenderedDosageInstruction)
}
func (resource *MedicationRequest) T_DispenseRequestId() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.DispenseRequest.Id", nil)
	}
	return StringInput("MedicationRequest.DispenseRequest.Id", resource.DispenseRequest.Id)
}
func (resource *MedicationRequest) T_DispenseRequestNumberOfRepeatsAllowed() templ.Component {

	if resource == nil {
		return IntInput("MedicationRequest.DispenseRequest.NumberOfRepeatsAllowed", nil)
	}
	return IntInput("MedicationRequest.DispenseRequest.NumberOfRepeatsAllowed", resource.DispenseRequest.NumberOfRepeatsAllowed)
}
func (resource *MedicationRequest) T_DispenseRequestDoseAdministrationAid(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationRequest.DispenseRequest.DoseAdministrationAid", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.DispenseRequest.DoseAdministrationAid", resource.DispenseRequest.DoseAdministrationAid, optionsValueSet)
}
func (resource *MedicationRequest) T_DispenseRequestInitialFillId() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.DispenseRequest.InitialFill.Id", nil)
	}
	return StringInput("MedicationRequest.DispenseRequest.InitialFill.Id", resource.DispenseRequest.InitialFill.Id)
}
func (resource *MedicationRequest) T_SubstitutionId() templ.Component {

	if resource == nil {
		return StringInput("MedicationRequest.Substitution.Id", nil)
	}
	return StringInput("MedicationRequest.Substitution.Id", resource.Substitution.Id)
}
func (resource *MedicationRequest) T_SubstitutionReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationRequest.Substitution.Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationRequest.Substitution.Reason", resource.Substitution.Reason, optionsValueSet)
}
