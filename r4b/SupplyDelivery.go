package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                     `json:"id,omitempty"`
	Meta               *Meta                       `json:"meta,omitempty"`
	ImplicitRules      *string                     `json:"implicitRules,omitempty"`
	Language           *string                     `json:"language,omitempty"`
	Text               *Narrative                  `json:"text,omitempty"`
	Contained          []Resource                  `json:"contained,omitempty"`
	Extension          []Extension                 `json:"extension,omitempty"`
	ModifierExtension  []Extension                 `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                `json:"identifier,omitempty"`
	BasedOn            []Reference                 `json:"basedOn,omitempty"`
	PartOf             []Reference                 `json:"partOf,omitempty"`
	Status             *string                     `json:"status,omitempty"`
	Patient            *Reference                  `json:"patient,omitempty"`
	Type               *CodeableConcept            `json:"type,omitempty"`
	SuppliedItem       *SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	OccurrenceDateTime *string                     `json:"occurrenceDateTime"`
	OccurrencePeriod   *Period                     `json:"occurrencePeriod"`
	OccurrenceTiming   *Timing                     `json:"occurrenceTiming"`
	Supplier           *Reference                  `json:"supplier,omitempty"`
	Destination        *Reference                  `json:"destination,omitempty"`
	Receiver           []Reference                 `json:"receiver,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SupplyDelivery
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

func (resource *SupplyDelivery) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SupplyDelivery) T_Status() templ.Component {
	optionsValueSet := VSSupplydelivery_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *SupplyDelivery) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
