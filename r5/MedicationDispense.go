package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	StatusChanged             *time.Time                      `json:"statusChanged,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	Recorded                  *time.Time                      `json:"recorded,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	WhenPrepared              *time.Time                      `json:"whenPrepared,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	WhenHandedOver            *time.Time                      `json:"whenHandedOver,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *MedicationDispense) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource == nil {
		return CodeSelect("MedicationDispense.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MedicationDispense.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_StatusChanged(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MedicationDispense.StatusChanged", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationDispense.StatusChanged", resource.StatusChanged, htmlAttrs)
}
func (resource *MedicationDispense) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("MedicationDispense.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationDispense.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationDispense.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationDispense.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_Recorded(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MedicationDispense.Recorded", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationDispense.Recorded", resource.Recorded, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenPrepared(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MedicationDispense.WhenPrepared", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationDispense.WhenPrepared", resource.WhenPrepared, htmlAttrs)
}
func (resource *MedicationDispense) T_WhenHandedOver(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MedicationDispense.WhenHandedOver", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationDispense.WhenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *MedicationDispense) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("MedicationDispense.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("MedicationDispense.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationDispense) T_RenderedDosageInstruction(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.RenderedDosageInstruction", nil, htmlAttrs)
	}
	return StringInput("MedicationDispense.RenderedDosageInstruction", resource.RenderedDosageInstruction, htmlAttrs)
}
func (resource *MedicationDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("MedicationDispense.Performer."+strconv.Itoa(numPerformer)+"..Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationDispense.Performer."+strconv.Itoa(numPerformer)+"..Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionWasSubstituted(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("MedicationDispense.Substitution.WasSubstituted", nil, htmlAttrs)
	}
	return BoolInput("MedicationDispense.Substitution.WasSubstituted", &resource.Substitution.WasSubstituted, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationDispense.Substitution.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationDispense.Substitution.Type", resource.Substitution.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationDispense) T_SubstitutionReason(numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReason >= len(resource.Substitution.Reason) {
		return CodeableConceptSelect("MedicationDispense.Substitution.Reason."+strconv.Itoa(numReason)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationDispense.Substitution.Reason."+strconv.Itoa(numReason)+".", &resource.Substitution.Reason[numReason], optionsValueSet, htmlAttrs)
}
