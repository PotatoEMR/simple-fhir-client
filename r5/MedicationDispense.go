package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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
func (resource *MedicationDispense) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_StatusChanged(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("statusChanged", nil, htmlAttrs)
	}
	return DateTimeInput("statusChanged", resource.StatusChanged, htmlAttrs)
}
func (resource *MedicationDispense) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("recorded", nil, htmlAttrs)
	}
	return DateTimeInput("recorded", resource.Recorded, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenPrepared(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("whenPrepared", nil, htmlAttrs)
	}
	return DateTimeInput("whenPrepared", resource.WhenPrepared, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenHandedOver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("whenHandedOver", nil, htmlAttrs)
	}
	return DateTimeInput("whenHandedOver", resource.WhenHandedOver, htmlAttrs)
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
func (resource *MedicationDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
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
