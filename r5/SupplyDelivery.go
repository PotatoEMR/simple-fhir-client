//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                      `json:"id,omitempty"`
	Meta               *Meta                        `json:"meta,omitempty"`
	ImplicitRules      *string                      `json:"implicitRules,omitempty"`
	Language           *string                      `json:"language,omitempty"`
	Text               *Narrative                   `json:"text,omitempty"`
	Contained          []Resource                   `json:"contained,omitempty"`
	Extension          []Extension                  `json:"extension,omitempty"`
	ModifierExtension  []Extension                  `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                 `json:"identifier,omitempty"`
	BasedOn            []Reference                  `json:"basedOn,omitempty"`
	PartOf             []Reference                  `json:"partOf,omitempty"`
	Status             *string                      `json:"status,omitempty"`
	Patient            *Reference                   `json:"patient,omitempty"`
	Type               *CodeableConcept             `json:"type,omitempty"`
	SuppliedItem       []SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	OccurrenceDateTime *string                      `json:"occurrenceDateTime"`
	OccurrencePeriod   *Period                      `json:"occurrencePeriod"`
	OccurrenceTiming   *Timing                      `json:"occurrenceTiming"`
	Supplier           *Reference                   `json:"supplier,omitempty"`
	Destination        *Reference                   `json:"destination,omitempty"`
	Receiver           []Reference                  `json:"receiver,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SupplyDelivery
type SupplyDeliverySuppliedItem struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Quantity            *Quantity        `json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       *Reference       `json:"itemReference"`
}

type OtherSupplyDelivery SupplyDelivery

// on convert struct to json, automatically add resourceType=SupplyDelivery
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}
