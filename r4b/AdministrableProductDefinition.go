package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
	Id                    *string                                               `json:"id,omitempty"`
	Meta                  *Meta                                                 `json:"meta,omitempty"`
	ImplicitRules         *string                                               `json:"implicitRules,omitempty"`
	Language              *string                                               `json:"language,omitempty"`
	Text                  *Narrative                                            `json:"text,omitempty"`
	Contained             []Resource                                            `json:"contained,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `json:"identifier,omitempty"`
	Status                string                                                `json:"status"`
	FormOf                []Reference                                           `json:"formOf,omitempty"`
	AdministrableDoseForm *CodeableConcept                                      `json:"administrableDoseForm,omitempty"`
	UnitOfPresentation    *CodeableConcept                                      `json:"unitOfPresentation,omitempty"`
	ProducedFrom          []Reference                                           `json:"producedFrom,omitempty"`
	Ingredient            []CodeableConcept                                     `json:"ingredient,omitempty"`
	Device                *Reference                                            `json:"device,omitempty"`
	Property              []AdministrableProductDefinitionProperty              `json:"property,omitempty"`
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	Status               *CodeableConcept `json:"status,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionRouteOfAdministration struct {
	Id                        *string                                                            `json:"id,omitempty"`
	Extension                 []Extension                                                        `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `json:"code"`
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code"`
	WithdrawalPeriod  []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                    *string         `json:"id,omitempty"`
	Extension             []Extension     `json:"extension,omitempty"`
	ModifierExtension     []Extension     `json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `json:"tissue"`
	Value                 Quantity        `json:"value"`
	SupportingInformation *string         `json:"supportingInformation,omitempty"`
}

type OtherAdministrableProductDefinition AdministrableProductDefinition

// on convert struct to json, automatically add resourceType=AdministrableProductDefinition
func (r AdministrableProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdministrableProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherAdministrableProductDefinition: OtherAdministrableProductDefinition(r),
		ResourceType:                        "AdministrableProductDefinition",
	})
}

func (resource *AdministrableProductDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("AdministrableProductDefinition.Id", nil)
	}
	return StringInput("AdministrableProductDefinition.Id", resource.Id)
}
func (resource *AdministrableProductDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("AdministrableProductDefinition.ImplicitRules", nil)
	}
	return StringInput("AdministrableProductDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *AdministrableProductDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("AdministrableProductDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("AdministrableProductDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("AdministrableProductDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("AdministrableProductDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_AdministrableDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdministrableProductDefinition.AdministrableDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.AdministrableDoseForm", resource.AdministrableDoseForm, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_UnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdministrableProductDefinition.UnitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return CodeableConceptSelect("AdministrableProductDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *AdministrableProductDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_PropertyStatus(numProperty int) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Property["+strconv.Itoa(numProperty)+"].Status", resource.Property[numProperty].Status, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationId(numRouteOfAdministration int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration {
		return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Id", nil)
	}
	return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].Id)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesId(numRouteOfAdministration int, numTargetSpecies int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies {
		return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Id", nil)
	}
	return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Id)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].Code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodId(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Id", nil)
	}
	return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Id", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Id)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].Tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int) templ.Component {

	if resource == nil || len(resource.RouteOfAdministration) >= numRouteOfAdministration || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) >= numTargetSpecies || len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) >= numWithdrawalPeriod {
		return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", nil)
	}
	return StringInput("AdministrableProductDefinition.RouteOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].TargetSpecies["+strconv.Itoa(numTargetSpecies)+"].WithdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].SupportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation)
}
