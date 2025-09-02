package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalAdministrableDoseForm(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("administrableDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("administrableDoseForm", &resource.AdministrableDoseForm, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalUnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalCharacteristicsCode(numCharacteristics int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Characteristics) >= numCharacteristics {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Characteristics[numCharacteristics].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalCharacteristicsStatus(numCharacteristics int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Characteristics) >= numCharacteristics {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Characteristics[numCharacteristics].Status, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalRouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.RouteOfAdministration) >= numRouteOfAdministration {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet)
}
func (resource *MedicinalProductPharmaceutical) MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return CodeableConceptSelect("tissue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet)
}
