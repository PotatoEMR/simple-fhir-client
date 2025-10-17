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

// struct -> json, automatically add resourceType=Patient
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
func (r MedicinalProductPharmaceutical) ResourceType() string {
	return "MedicinalProductPharmaceutical"
}

func (resource *MedicinalProductPharmaceutical) T_AdministrableDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("administrableDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("administrableDoseForm", &resource.AdministrableDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_Ingredient(frs []FhirResource, numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_Device(frs []FhirResource, numDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDevice >= len(resource.Device) {
		return ReferenceInput(frs, "device["+strconv.Itoa(numDevice)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "device["+strconv.Itoa(numDevice)+"]", &resource.Device[numDevice], htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsCode(numCharacteristics int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristics >= len(resource.Characteristics) {
		return CodeableConceptSelect("characteristics["+strconv.Itoa(numCharacteristics)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristics["+strconv.Itoa(numCharacteristics)+"].code", &resource.Characteristics[numCharacteristics].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_CharacteristicsStatus(numCharacteristics int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristics >= len(resource.Characteristics) {
		return CodeableConceptSelect("characteristics["+strconv.Itoa(numCharacteristics)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristics["+strconv.Itoa(numCharacteristics)+"].status", resource.Characteristics[numCharacteristics].Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationFirstDose(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].firstDose", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].firstDose", resource.RouteOfAdministration[numRouteOfAdministration].FirstDose, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationMaxSingleDose(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxSingleDose", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxSingleDose", resource.RouteOfAdministration[numRouteOfAdministration].MaxSingleDose, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationMaxDosePerDay(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerDay", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerDay", resource.RouteOfAdministration[numRouteOfAdministration].MaxDosePerDay, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationMaxDosePerTreatmentPeriod(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return RatioInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerTreatmentPeriod", nil, htmlAttrs)
	}
	return RatioInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerTreatmentPeriod", resource.RouteOfAdministration[numRouteOfAdministration].MaxDosePerTreatmentPeriod, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationMaxTreatmentPeriod(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return DurationInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxTreatmentPeriod", nil, htmlAttrs)
	}
	return DurationInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxTreatmentPeriod", resource.RouteOfAdministration[numRouteOfAdministration].MaxTreatmentPeriod, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].tissue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodValue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].value", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].value", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Value, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPharmaceutical) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", nil, htmlAttrs)
	}
	return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation, htmlAttrs)
}
