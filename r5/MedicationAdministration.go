package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                    *string                             `json:"id,omitempty"`
	Meta                  *Meta                               `json:"meta,omitempty"`
	ImplicitRules         *string                             `json:"implicitRules,omitempty"`
	Language              *string                             `json:"language,omitempty"`
	Text                  *Narrative                          `json:"text,omitempty"`
	Contained             []Resource                          `json:"contained,omitempty"`
	Extension             []Extension                         `json:"extension,omitempty"`
	ModifierExtension     []Extension                         `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                        `json:"identifier,omitempty"`
	BasedOn               []Reference                         `json:"basedOn,omitempty"`
	PartOf                []Reference                         `json:"partOf,omitempty"`
	Status                string                              `json:"status"`
	StatusReason          []CodeableConcept                   `json:"statusReason,omitempty"`
	Category              []CodeableConcept                   `json:"category,omitempty"`
	Medication            CodeableReference                   `json:"medication"`
	Subject               Reference                           `json:"subject"`
	Encounter             *Reference                          `json:"encounter,omitempty"`
	SupportingInformation []Reference                         `json:"supportingInformation,omitempty"`
	OccurenceDateTime     string                              `json:"occurenceDateTime"`
	OccurencePeriod       Period                              `json:"occurencePeriod"`
	OccurenceTiming       Timing                              `json:"occurenceTiming"`
	Recorded              *string                             `json:"recorded,omitempty"`
	IsSubPotent           *bool                               `json:"isSubPotent,omitempty"`
	SubPotentReason       []CodeableConcept                   `json:"subPotentReason,omitempty"`
	Performer             []MedicationAdministrationPerformer `json:"performer,omitempty"`
	Reason                []CodeableReference                 `json:"reason,omitempty"`
	Request               *Reference                          `json:"request,omitempty"`
	Device                []CodeableReference                 `json:"device,omitempty"`
	Note                  []Annotation                        `json:"note,omitempty"`
	Dosage                *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory          []Reference                         `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
type MedicationAdministrationPerformer struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept  `json:"function,omitempty"`
	Actor             CodeableReference `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
type MedicationAdministrationDosage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Site              *CodeableConcept `json:"site,omitempty"`
	Route             *CodeableConcept `json:"route,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Dose              *Quantity        `json:"dose,omitempty"`
	RateRatio         *Ratio           `json:"rateRatio,omitempty"`
	RateQuantity      *Quantity        `json:"rateQuantity,omitempty"`
}

type OtherMedicationAdministration MedicationAdministration

// on convert struct to json, automatically add resourceType=MedicationAdministration
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
}

func (resource *MedicationAdministration) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicationAdministration.Id", nil)
	}
	return StringInput("MedicationAdministration.Id", resource.Id)
}
func (resource *MedicationAdministration) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicationAdministration.ImplicitRules", nil)
	}
	return StringInput("MedicationAdministration.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicationAdministration) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicationAdministration.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicationAdministration.Language", resource.Language, optionsValueSet)
}
func (resource *MedicationAdministration) T_Status() templ.Component {
	optionsValueSet := VSMedication_admin_status

	if resource == nil {
		return CodeSelect("MedicationAdministration.Status", nil, optionsValueSet)
	}
	return CodeSelect("MedicationAdministration.Status", &resource.Status, optionsValueSet)
}
func (resource *MedicationAdministration) T_StatusReason(numStatusReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StatusReason) >= numStatusReason {
		return CodeableConceptSelect("MedicationAdministration.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet)
}
func (resource *MedicationAdministration) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("MedicationAdministration.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *MedicationAdministration) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("MedicationAdministration.Recorded", nil)
	}
	return StringInput("MedicationAdministration.Recorded", resource.Recorded)
}
func (resource *MedicationAdministration) T_IsSubPotent() templ.Component {

	if resource == nil {
		return BoolInput("MedicationAdministration.IsSubPotent", nil)
	}
	return BoolInput("MedicationAdministration.IsSubPotent", resource.IsSubPotent)
}
func (resource *MedicationAdministration) T_SubPotentReason(numSubPotentReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SubPotentReason) >= numSubPotentReason {
		return CodeableConceptSelect("MedicationAdministration.SubPotentReason["+strconv.Itoa(numSubPotentReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.SubPotentReason["+strconv.Itoa(numSubPotentReason)+"]", &resource.SubPotentReason[numSubPotentReason], optionsValueSet)
}
func (resource *MedicationAdministration) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *MedicationAdministration) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageId() templ.Component {

	if resource == nil {
		return StringInput("MedicationAdministration.Dosage.Id", nil)
	}
	return StringInput("MedicationAdministration.Dosage.Id", resource.Dosage.Id)
}
func (resource *MedicationAdministration) T_DosageText() templ.Component {

	if resource == nil {
		return StringInput("MedicationAdministration.Dosage.Text", nil)
	}
	return StringInput("MedicationAdministration.Dosage.Text", resource.Dosage.Text)
}
func (resource *MedicationAdministration) T_DosageSite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Site", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Site", resource.Dosage.Site, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageRoute(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Route", resource.Dosage.Route, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageMethod(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Method", resource.Dosage.Method, optionsValueSet)
}
