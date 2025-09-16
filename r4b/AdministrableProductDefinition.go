package r4b

//generated with command go run ./bultaoreune
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
	ValueDate            *FhirDate        `json:"valueDate,omitempty"`
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
func (r AdministrableProductDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AdministrableProductDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "AdministrableProductDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AdministrableProductDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_AdministrableDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("administrableDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("administrableDoseForm", resource.AdministrableDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueDate(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyStatus(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].status", resource.Property[numProperty].Status, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].tissue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", nil, htmlAttrs)
	}
	return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation, htmlAttrs)
}
