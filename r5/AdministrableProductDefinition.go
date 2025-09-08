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

// http://hl7.org/fhir/r5/StructureDefinition/AdministrableProductDefinition
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
	Description           *string                                               `json:"description,omitempty"`
	Property              []AdministrableProductDefinitionProperty              `json:"property,omitempty"`
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *time.Time       `json:"valueDate,omitempty,format:'2006-01-02'"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
	Status               *CodeableConcept `json:"status,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdministrableProductDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code"`
	WithdrawalPeriod  []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdministrableProductDefinition
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
func (resource *AdministrableProductDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("AdministrableProductDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AdministrableProductDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_AdministrableDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdministrableProductDefinition.AdministrableDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.AdministrableDoseForm", resource.AdministrableDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdministrableProductDefinition.UnitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("AdministrableProductDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("AdministrableProductDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("AdministrableProductDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueDate(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueMarkdown(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..ValueMarkdown", resource.Property[numProperty].ValueMarkdown, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyStatus(numProperty int, htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.Property."+strconv.Itoa(numProperty)+"..Status", resource.Property[numProperty].Status, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationCode(numRouteOfAdministration int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..Code", &resource.RouteOfAdministration[numRouteOfAdministration].Code, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesCode(numRouteOfAdministration int, numTargetSpecies int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..Code", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].Code, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodTissue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..WithdrawalPeriod."+strconv.Itoa(numWithdrawalPeriod)+"..Tissue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..WithdrawalPeriod."+strconv.Itoa(numWithdrawalPeriod)+"..Tissue", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Tissue, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs string) templ.Component {

	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return StringInput("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..WithdrawalPeriod."+strconv.Itoa(numWithdrawalPeriod)+"..SupportingInformation", nil, htmlAttrs)
	}
	return StringInput("AdministrableProductDefinition.RouteOfAdministration."+strconv.Itoa(numRouteOfAdministration)+"..TargetSpecies."+strconv.Itoa(numTargetSpecies)+"..WithdrawalPeriod."+strconv.Itoa(numWithdrawalPeriod)+"..SupportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation, htmlAttrs)
}
