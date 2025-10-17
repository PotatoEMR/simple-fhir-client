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

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                        *string                         `json:"id,omitempty"`
	Meta                      *Meta                           `json:"meta,omitempty"`
	ImplicitRules             *string                         `json:"implicitRules,omitempty"`
	Language                  *string                         `json:"language,omitempty"`
	Text                      *Narrative                      `json:"text,omitempty"`
	Contained                 []Resource                      `json:"contained,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                    `json:"identifier,omitempty"`
	BasedOn                   []Reference                     `json:"basedOn,omitempty"`
	PartOf                    []Reference                     `json:"partOf,omitempty"`
	Status                    string                          `json:"status"`
	NotPerformedReason        *CodeableReference              `json:"notPerformedReason,omitempty"`
	StatusChanged             *FhirDateTime                   `json:"statusChanged,omitempty"`
	Category                  []CodeableConcept               `json:"category,omitempty"`
	Medication                CodeableReference               `json:"medication"`
	Subject                   Reference                       `json:"subject"`
	Encounter                 *Reference                      `json:"encounter,omitempty"`
	SupportingInformation     []Reference                     `json:"supportingInformation,omitempty"`
	Performer                 []MedicationDispensePerformer   `json:"performer,omitempty"`
	Location                  *Reference                      `json:"location,omitempty"`
	AuthorizingPrescription   []Reference                     `json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                `json:"type,omitempty"`
	Quantity                  *Quantity                       `json:"quantity,omitempty"`
	DaysSupply                *Quantity                       `json:"daysSupply,omitempty"`
	Recorded                  *FhirDateTime                   `json:"recorded,omitempty"`
	WhenPrepared              *FhirDateTime                   `json:"whenPrepared,omitempty"`
	WhenHandedOver            *FhirDateTime                   `json:"whenHandedOver,omitempty"`
	Destination               *Reference                      `json:"destination,omitempty"`
	Receiver                  []Reference                     `json:"receiver,omitempty"`
	Note                      []Annotation                    `json:"note,omitempty"`
	RenderedDosageInstruction *string                         `json:"renderedDosageInstruction,omitempty"`
	DosageInstruction         []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	EventHistory              []Reference                     `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispenseSubstitution struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `json:"wasSubstituted"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	ResponsibleParty  *Reference        `json:"responsibleParty,omitempty"`
}

type OtherMedicationDispense MedicationDispense

// struct -> json, automatically add resourceType=Patient
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
}

// json -> struct, first reject if resourceType != MedicationDispense
func (r *MedicationDispense) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MedicationDispense" {
		return errors.New("resourceType not MedicationDispense")
	}
	return json.Unmarshal(data, (*OtherMedicationDispense)(r))
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
func (resource *MedicationDispense) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *MedicationDispense) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *MedicationDispense) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_NotPerformedReason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("notPerformedReason", nil, htmlAttrs)
	}
	return CodeableReferenceInput("notPerformedReason", resource.NotPerformedReason, htmlAttrs)
}
func (resource *MedicationDispense) T_StatusChanged(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("statusChanged", nil, htmlAttrs)
	}
	return FhirDateTimeInput("statusChanged", resource.StatusChanged, htmlAttrs)
}
func (resource *MedicationDispense) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Medication(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("medication", nil, htmlAttrs)
	}
	return CodeableReferenceInput("medication", &resource.Medication, htmlAttrs)
}
func (resource *MedicationDispense) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationDispense) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *MedicationDispense) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *MedicationDispense) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *MedicationDispense) T_AuthorizingPrescription(frs []FhirResource, numAuthorizingPrescription int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthorizingPrescription >= len(resource.AuthorizingPrescription) {
		return ReferenceInput(frs, "authorizingPrescription["+strconv.Itoa(numAuthorizingPrescription)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "authorizingPrescription["+strconv.Itoa(numAuthorizingPrescription)+"]", &resource.AuthorizingPrescription[numAuthorizingPrescription], htmlAttrs)
}
func (resource *MedicationDispense) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_DaysSupply(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("daysSupply", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("daysSupply", resource.DaysSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recorded", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recorded", resource.Recorded, htmlAttrs)
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
func (resource *MedicationDispense) T_Destination(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "destination", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination", resource.Destination, htmlAttrs)
}
func (resource *MedicationDispense) T_Receiver(frs []FhirResource, numReceiver int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReceiver >= len(resource.Receiver) {
		return ReferenceInput(frs, "receiver["+strconv.Itoa(numReceiver)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "receiver["+strconv.Itoa(numReceiver)+"]", &resource.Receiver[numReceiver], htmlAttrs)
}
func (resource *MedicationDispense) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationDispense) T_RenderedDosageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("renderedDosageInstruction", nil, htmlAttrs)
	}
	return StringInput("renderedDosageInstruction", resource.RenderedDosageInstruction, htmlAttrs)
}
func (resource *MedicationDispense) T_DosageInstruction(numDosageInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosageInstruction >= len(resource.DosageInstruction) {
		return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosageInstruction["+strconv.Itoa(numDosageInstruction)+"]", &resource.DosageInstruction[numDosageInstruction], htmlAttrs)
}
func (resource *MedicationDispense) T_EventHistory(frs []FhirResource, numEventHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventHistory >= len(resource.EventHistory) {
		return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", &resource.EventHistory[numEventHistory], htmlAttrs)
}
func (resource *MedicationDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
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
func (resource *MedicationDispense) T_SubstitutionResponsibleParty(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "substitution.responsibleParty", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "substitution.responsibleParty", resource.Substitution.ResponsibleParty, htmlAttrs)
}
