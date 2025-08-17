//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                        *string                             `json:"id,omitempty"`
	Meta                      *Meta                               `json:"meta,omitempty"`
	ImplicitRules             *string                             `json:"implicitRules,omitempty"`
	Language                  *string                             `json:"language,omitempty"`
	Text                      *Narrative                          `json:"text,omitempty"`
	Contained                 []Resource                          `json:"contained,omitempty"`
	Extension                 []Extension                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                         `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                        `json:"identifier,omitempty"`
	Instantiates              []string                            `json:"instantiates,omitempty"`
	PartOf                    []Reference                         `json:"partOf,omitempty"`
	Status                    string                              `json:"status"`
	StatusReason              []CodeableConcept                   `json:"statusReason,omitempty"`
	Category                  *CodeableConcept                    `json:"category,omitempty"`
	MedicationCodeableConcept CodeableConcept                     `json:"medicationCodeableConcept"`
	MedicationReference       Reference                           `json:"medicationReference"`
	Subject                   Reference                           `json:"subject"`
	Context                   *Reference                          `json:"context,omitempty"`
	SupportingInformation     []Reference                         `json:"supportingInformation,omitempty"`
	EffectiveDateTime         string                              `json:"effectiveDateTime"`
	EffectivePeriod           Period                              `json:"effectivePeriod"`
	Performer                 []MedicationAdministrationPerformer `json:"performer,omitempty"`
	ReasonCode                []CodeableConcept                   `json:"reasonCode,omitempty"`
	ReasonReference           []Reference                         `json:"reasonReference,omitempty"`
	Request                   *Reference                          `json:"request,omitempty"`
	Device                    []Reference                         `json:"device,omitempty"`
	Note                      []Annotation                        `json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory              []Reference                         `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
type MedicationAdministrationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
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
