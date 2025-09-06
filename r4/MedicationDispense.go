package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicationDispense
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
	WhenPrepared                *string                         `json:"whenPrepared,omitempty"`
	WhenHandedOver              *string                         `json:"whenHandedOver,omitempty"`
	Destination                 *Reference                      `json:"destination,omitempty"`
	Receiver                    []Reference                     `json:"receiver,omitempty"`
	Note                        []Annotation                    `json:"note,omitempty"`
	DosageInstruction           []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution                *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	DetectedIssue               []Reference                     `json:"detectedIssue,omitempty"`
	EventHistory                []Reference                     `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationDispense
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

func (resource *MedicationDispense) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.Id", nil)
	}
	return StringInput("MedicationDispense.Id", resource.Id)
}
func (resource *MedicationDispense) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.ImplicitRules", nil)
	}
	return StringInput("MedicationDispense.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicationDispense) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicationDispense.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicationDispense.Language", resource.Language, optionsValueSet)
}
func (resource *MedicationDispense) T_Status() templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource == nil {
		return CodeSelect("MedicationDispense.Status", nil, optionsValueSet)
	}
	return CodeSelect("MedicationDispense.Status", &resource.Status, optionsValueSet)
}
func (resource *MedicationDispense) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationDispense.Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationDispense.Category", resource.Category, optionsValueSet)
}
func (resource *MedicationDispense) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationDispense.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationDispense.Type", resource.Type, optionsValueSet)
}
func (resource *MedicationDispense) T_WhenPrepared() templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.WhenPrepared", nil)
	}
	return StringInput("MedicationDispense.WhenPrepared", resource.WhenPrepared)
}
func (resource *MedicationDispense) T_WhenHandedOver() templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.WhenHandedOver", nil)
	}
	return StringInput("MedicationDispense.WhenHandedOver", resource.WhenHandedOver)
}
func (resource *MedicationDispense) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("MedicationDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("MedicationDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *MedicationDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("MedicationDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *MedicationDispense) T_SubstitutionId() templ.Component {

	if resource == nil {
		return StringInput("MedicationDispense.Substitution.Id", nil)
	}
	return StringInput("MedicationDispense.Substitution.Id", resource.Substitution.Id)
}
func (resource *MedicationDispense) T_SubstitutionWasSubstituted() templ.Component {

	if resource == nil {
		return BoolInput("MedicationDispense.Substitution.WasSubstituted", nil)
	}
	return BoolInput("MedicationDispense.Substitution.WasSubstituted", &resource.Substitution.WasSubstituted)
}
func (resource *MedicationDispense) T_SubstitutionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationDispense.Substitution.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationDispense.Substitution.Type", resource.Substitution.Type, optionsValueSet)
}
func (resource *MedicationDispense) T_SubstitutionReason(numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Substitution.Reason) >= numReason {
		return CodeableConceptSelect("MedicationDispense.Substitution.Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationDispense.Substitution.Reason["+strconv.Itoa(numReason)+"]", &resource.Substitution.Reason[numReason], optionsValueSet)
}
