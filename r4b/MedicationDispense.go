package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                          *string                         `json:"id,omitempty"`
	Meta                        *Meta                           `json:"meta,omitempty"`
	ImplicitRules               *string                         `json:"implicitRules,omitempty"`
	Language                    *string                         `json:"language,omitempty"`
	Text                        *Narrative                      `json:"text,omitempty"`
	Contained                   []Resource                      `json:"contained,omitempty"`
	Extension                   []Extension                     `json:"extension,omitempty"`
	ModifierExtension           []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                  []Identifier                    `json:"identifier,omitempty"`
	PartOf                      []Reference                     `json:"partOf,omitempty"`
	Status                      string                          `json:"status"`
	StatusReasonCodeableConcept *CodeableConcept                `json:"statusReasonCodeableConcept,omitempty"`
	StatusReasonReference       *Reference                      `json:"statusReasonReference,omitempty"`
	Category                    *CodeableConcept                `json:"category,omitempty"`
	MedicationCodeableConcept   CodeableConcept                 `json:"medicationCodeableConcept"`
	MedicationReference         Reference                       `json:"medicationReference"`
	Subject                     *Reference                      `json:"subject,omitempty"`
	Context                     *Reference                      `json:"context,omitempty"`
	SupportingInformation       []Reference                     `json:"supportingInformation,omitempty"`
	Performer                   []MedicationDispensePerformer   `json:"performer,omitempty"`
	Location                    *Reference                      `json:"location,omitempty"`
	AuthorizingPrescription     []Reference                     `json:"authorizingPrescription,omitempty"`
	Type                        *CodeableConcept                `json:"type,omitempty"`
	Quantity                    *Quantity                       `json:"quantity,omitempty"`
	DaysSupply                  *Quantity                       `json:"daysSupply,omitempty"`
	WhenPrepared                *FhirDateTime                   `json:"whenPrepared,omitempty"`
	WhenHandedOver              *FhirDateTime                   `json:"whenHandedOver,omitempty"`
	Destination                 *Reference                      `json:"destination,omitempty"`
	Receiver                    []Reference                     `json:"receiver,omitempty"`
	Note                        []Annotation                    `json:"note,omitempty"`
	DosageInstruction           []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution                *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	DetectedIssue               []Reference                     `json:"detectedIssue,omitempty"`
	EventHistory                []Reference                     `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispenseSubstitution struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `json:"wasSubstituted"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	ResponsibleParty  []Reference       `json:"responsibleParty,omitempty"`
}

type OtherMedicationDispense MedicationDispense

// on convert struct to json, automatically add resourceType=MedicationDispense
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
}
func (r MedicationDispense) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationDispense/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicationDispense"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicationDispense) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *MedicationDispense) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_StatusReasonCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReasonCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReasonCodeableConcept", resource.StatusReasonCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_StatusReasonReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("statusReasonReference", nil, htmlAttrs)
	}
	return ReferenceInput("statusReasonReference", resource.StatusReasonReference, htmlAttrs)
}
func (resource *MedicationDispense) T_Category(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_MedicationCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicationCodeableConcept", &resource.MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_MedicationReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("medicationReference", nil, htmlAttrs)
	}
	return ReferenceInput("medicationReference", &resource.MedicationReference, htmlAttrs)
}
func (resource *MedicationDispense) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", resource.Subject, htmlAttrs)
}
func (resource *MedicationDispense) T_Context(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("context", nil, htmlAttrs)
	}
	return ReferenceInput("context", resource.Context, htmlAttrs)
}
func (resource *MedicationDispense) T_SupportingInformation(numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput("supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *MedicationDispense) T_Location(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("location", nil, htmlAttrs)
	}
	return ReferenceInput("location", resource.Location, htmlAttrs)
}
func (resource *MedicationDispense) T_AuthorizingPrescription(numAuthorizingPrescription int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthorizingPrescription >= len(resource.AuthorizingPrescription) {
		return ReferenceInput("authorizingPrescription["+strconv.Itoa(numAuthorizingPrescription)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("authorizingPrescription["+strconv.Itoa(numAuthorizingPrescription)+"]", &resource.AuthorizingPrescription[numAuthorizingPrescription], htmlAttrs)
}
func (resource *MedicationDispense) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Quantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *MedicationDispense) T_DaysSupply(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("daysSupply", nil, htmlAttrs)
	}
	return QuantityInput("daysSupply", resource.DaysSupply, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenPrepared(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("whenPrepared", nil, htmlAttrs)
	}
	return FhirDateTimeInput("whenPrepared", resource.WhenPrepared, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenHandedOver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("whenHandedOver", nil, htmlAttrs)
	}
	return FhirDateTimeInput("whenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *MedicationDispense) T_Destination(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("destination", nil, htmlAttrs)
	}
	return ReferenceInput("destination", resource.Destination, htmlAttrs)
}
func (resource *MedicationDispense) T_Receiver(numReceiver int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReceiver >= len(resource.Receiver) {
		return ReferenceInput("receiver["+strconv.Itoa(numReceiver)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("receiver["+strconv.Itoa(numReceiver)+"]", &resource.Receiver[numReceiver], htmlAttrs)
}
func (resource *MedicationDispense) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationDispense) T_DosageInstruction(numDosageInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosageInstruction >= len(resource.DosageInstruction) {
		return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", &resource.DosageInstruction[numDosageInstruction], htmlAttrs)
}
func (resource *MedicationDispense) T_DetectedIssue(numDetectedIssue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetectedIssue >= len(resource.DetectedIssue) {
		return ReferenceInput("detectedIssue["+strconv.Itoa(numDetectedIssue)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("detectedIssue["+strconv.Itoa(numDetectedIssue)+"]", &resource.DetectedIssue[numDetectedIssue], htmlAttrs)
}
func (resource *MedicationDispense) T_EventHistory(numEventHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventHistory >= len(resource.EventHistory) {
		return ReferenceInput("eventHistory["+strconv.Itoa(numEventHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("eventHistory["+strconv.Itoa(numEventHistory)+"]", &resource.EventHistory[numEventHistory], htmlAttrs)
}
func (resource *MedicationDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_PerformerActor(numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionWasSubstituted(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("substitution.wasSubstituted", nil, htmlAttrs)
	}
	return BoolInput("substitution.wasSubstituted", &resource.Substitution.WasSubstituted, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("substitution.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substitution.type", resource.Substitution.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionReason(numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Substitution.Reason) {
		return CodeableConceptSelect("substitution.reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substitution.reason["+strconv.Itoa(numReason)+"]", &resource.Substitution.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionResponsibleParty(numResponsibleParty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResponsibleParty >= len(resource.Substitution.ResponsibleParty) {
		return ReferenceInput("substitution.responsibleParty["+strconv.Itoa(numResponsibleParty)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("substitution.responsibleParty["+strconv.Itoa(numResponsibleParty)+"]", &resource.Substitution.ResponsibleParty[numResponsibleParty], htmlAttrs)
}
