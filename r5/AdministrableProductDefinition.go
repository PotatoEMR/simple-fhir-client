package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	ValueDate            *FhirDate        `json:"valueDate,omitempty"`
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
func (resource *AdministrableProductDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_FormOf(numFormOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFormOf >= len(resource.FormOf) {
		return ReferenceInput("formOf["+strconv.Itoa(numFormOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("formOf["+strconv.Itoa(numFormOf)+"]", &resource.FormOf[numFormOf], htmlAttrs)
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
func (resource *AdministrableProductDefinition) T_ProducedFrom(numProducedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProducedFrom >= len(resource.ProducedFrom) {
		return ReferenceInput("producedFrom["+strconv.Itoa(numProducedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("producedFrom["+strconv.Itoa(numProducedFrom)+"]", &resource.ProducedFrom[numProducedFrom], htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Device(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("device", nil, htmlAttrs)
	}
	return ReferenceInput("device", resource.Device, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
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
func (resource *AdministrableProductDefinition) T_PropertyValueQuantity(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", resource.Property[numProperty].ValueQuantity, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueDate(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return FhirDateInput("property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("property["+strconv.Itoa(numProperty)+"].valueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueMarkdown(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueMarkdown", resource.Property[numProperty].ValueMarkdown, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueAttachment(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", resource.Property[numProperty].ValueAttachment, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_PropertyValueReference(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return ReferenceInput("property["+strconv.Itoa(numProperty)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput("property["+strconv.Itoa(numProperty)+"].valueReference", resource.Property[numProperty].ValueReference, htmlAttrs)
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
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationFirstDose(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].firstDose", nil, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].firstDose", resource.RouteOfAdministration[numRouteOfAdministration].FirstDose, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationMaxSingleDose(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxSingleDose", nil, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxSingleDose", resource.RouteOfAdministration[numRouteOfAdministration].MaxSingleDose, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationMaxDosePerDay(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerDay", nil, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerDay", resource.RouteOfAdministration[numRouteOfAdministration].MaxDosePerDay, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationMaxDosePerTreatmentPeriod(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return RatioInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerTreatmentPeriod", nil, htmlAttrs)
	}
	return RatioInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxDosePerTreatmentPeriod", resource.RouteOfAdministration[numRouteOfAdministration].MaxDosePerTreatmentPeriod, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationMaxTreatmentPeriod(numRouteOfAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) {
		return DurationInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxTreatmentPeriod", nil, htmlAttrs)
	}
	return DurationInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].maxTreatmentPeriod", resource.RouteOfAdministration[numRouteOfAdministration].MaxTreatmentPeriod, htmlAttrs)
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
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodValue(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].value", nil, htmlAttrs)
	}
	return QuantityInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].value", &resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].Value, htmlAttrs)
}
func (resource *AdministrableProductDefinition) T_RouteOfAdministrationTargetSpeciesWithdrawalPeriodSupportingInformation(numRouteOfAdministration int, numTargetSpecies int, numWithdrawalPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRouteOfAdministration >= len(resource.RouteOfAdministration) || numTargetSpecies >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies) || numWithdrawalPeriod >= len(resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod) {
		return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", nil, htmlAttrs)
	}
	return StringInput("routeOfAdministration["+strconv.Itoa(numRouteOfAdministration)+"].targetSpecies["+strconv.Itoa(numTargetSpecies)+"].withdrawalPeriod["+strconv.Itoa(numWithdrawalPeriod)+"].supportingInformation", resource.RouteOfAdministration[numRouteOfAdministration].TargetSpecies[numTargetSpecies].WithdrawalPeriod[numWithdrawalPeriod].SupportingInformation, htmlAttrs)
}
