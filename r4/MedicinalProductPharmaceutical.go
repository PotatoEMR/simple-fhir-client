package r4

//generated with command go run ./bultaoreune
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
func (r MedicinalProductPharmaceutical) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductPharmaceutical/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicinalProductPharmaceutical"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductPharmaceutical) T_AdministrableDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.AdministrableDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.AdministrableDoseForm", &resource.AdministrableDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.UnitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsCode(numCharacteristics int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristics >= len(resource.Characteristics) {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Code", &resource.Characteristics[numCharacteristics].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsStatus(numCharacteristics int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristics >= len(resource.Characteristics) {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.Characteristics["+strconv.Itoa(numCharacteristics)+"].Status", resource.Characteristics[numCharacteristics].Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs string) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductPharmaceutical.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation, htmlAttrs)
}
