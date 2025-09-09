package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	OccurrenceDateTime *string                      `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                      `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                      `json:"occurrenceTiming,omitempty"`
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
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
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
func (r SupplyDelivery) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SupplyDelivery/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "SupplyDelivery"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SupplyDelivery) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSupplydelivery_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSSupplydelivery_supplyitemtype

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *SupplyDelivery) T_SuppliedItemItemCodeableConcept(numSuppliedItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSuppliedItem >= len(resource.SuppliedItem) {
		return CodeableConceptSelect("suppliedItem["+strconv.Itoa(numSuppliedItem)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suppliedItem["+strconv.Itoa(numSuppliedItem)+"].itemCodeableConcept", resource.SuppliedItem[numSuppliedItem].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
