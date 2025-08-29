package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
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

func (resource *AdministrableProductDefinition) AdministrableProductDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *AdministrableProductDefinition) AdministrableProductDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
