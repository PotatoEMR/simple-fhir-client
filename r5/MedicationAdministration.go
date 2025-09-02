package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	RateRatio         *Ratio           `json:"rateRatio"`
	RateQuantity      *Quantity        `json:"rateQuantity"`
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

func (resource *MedicationAdministration) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicationAdministration) T_Status() templ.Component {
	optionsValueSet := VSMedication_admin_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *MedicationAdministration) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", &resource.StatusReason[0], optionsValueSet)
}
func (resource *MedicationAdministration) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *MedicationAdministration) T_SubPotentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subPotentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subPotentReason", &resource.SubPotentReason[0], optionsValueSet)
}
func (resource *MedicationAdministration) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageSite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("site", nil, optionsValueSet)
	}
	return CodeableConceptSelect("site", resource.Dosage.Site, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageRoute(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("route", resource.Dosage.Route, optionsValueSet)
}
func (resource *MedicationAdministration) T_DosageMethod(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Dosage.Method, optionsValueSet)
}
