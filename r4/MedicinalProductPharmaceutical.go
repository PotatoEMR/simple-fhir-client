package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceutical struct {
	Id                    *string                                               `json:"id,omitempty"`
	Meta                  *Meta                                                 `json:"meta,omitempty"`
	ImplicitRules         *string                                               `json:"implicitRules,omitempty"`
	Language              *string                                               `json:"language,omitempty"`
	Text                  *Narrative                                            `json:"text,omitempty"`
	Contained             []Resource                                            `json:"contained,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `json:"identifier,omitempty"`
	AdministrableDoseForm CodeableConcept                                       `json:"administrableDoseForm"`
	UnitOfPresentation    *CodeableConcept                                      `json:"unitOfPresentation,omitempty"`
	Ingredient            []Reference                                           `json:"ingredient,omitempty"`
	Device                []Reference                                           `json:"device,omitempty"`
	Characteristics       []MedicinalProductPharmaceuticalCharacteristics       `json:"characteristics,omitempty"`
	RouteOfAdministration []MedicinalProductPharmaceuticalRouteOfAdministration `json:"routeOfAdministration"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceuticalCharacteristics struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `json:"code"`
	Status            *CodeableConcept `json:"status,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceuticalRouteOfAdministration struct {
	Id                        *string                                                            `json:"id,omitempty"`
	Extension                 []Extension                                                        `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `json:"code"`
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code"`
	WithdrawalPeriod  []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                    *string         `json:"id,omitempty"`
	Extension             []Extension     `json:"extension,omitempty"`
	ModifierExtension     []Extension     `json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `json:"tissue"`
	Value                 Quantity        `json:"value"`
	SupportingInformation *string         `json:"supportingInformation,omitempty"`
}

type OtherMedicinalProductPharmaceutical MedicinalProductPharmaceutical

// on convert struct to json, automatically add resourceType=MedicinalProductPharmaceutical
func (r MedicinalProductPharmaceutical) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPharmaceutical
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPharmaceutical: OtherMedicinalProductPharmaceutical(r),
		ResourceType:                        "MedicinalProductPharmaceutical",
	})
}

func (resource *MedicinalProductPharmaceutical) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductPharmaceutical.Id", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.Id", resource.Id)
}
func (resource *MedicinalProductPharmaceutical) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductPharmaceutical.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductPharmaceutical) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductPharmaceutical.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductPharmaceutical.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_AdministrableDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.AdministrableDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.AdministrableDoseForm", &resource.AdministrableDoseForm, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_UnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.UnitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsId(numCharacteristics int) templ.Component {

	if resource == nil || len(resource.Characteristics) >= numCharacteristics {
		return StringInput("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Id", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Id", resource.Characteristics[numCharacteristics].Id)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsCode(numCharacteristics int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristics) >= numCharacteristics {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Code", &resource.Characteristics[numCharacteristics].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsStatus(numCharacteristics int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristics) >= numCharacteristics {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Status", resource.Characteristics[numCharacteristics].Status, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationId(numRouteOfAdministration int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration {
		return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Id", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].Id)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesId(numRouteOfAdministration int, numTargetSpecies int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies {
		return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Id", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Id)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodId(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Id", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Id)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", nil)
	}
	return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation)
}
