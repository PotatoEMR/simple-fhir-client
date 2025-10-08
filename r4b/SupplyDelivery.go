package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	OccurrenceDateTime *FhirDateTime               `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                     `json:"occurrenceTiming,omitempty"`
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
func (resource *SupplyDelivery) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *SupplyDelivery) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *SupplyDelivery) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSupplydelivery_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", resource.Patient, htmlAttrs)
}
func (resource *SupplyDelivery) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *SupplyDelivery) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *SupplyDelivery) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *SupplyDelivery) T_Supplier(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "supplier", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supplier", resource.Supplier, htmlAttrs)
}
func (resource *SupplyDelivery) T_Destination(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "destination", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination", resource.Destination, htmlAttrs)
}
func (resource *SupplyDelivery) T_Receiver(frs []FhirResource, numReceiver int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReceiver >= len(resource.Receiver) {
		return ReferenceInput(frs, "receiver["+strconv.Itoa(numReceiver)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "receiver["+strconv.Itoa(numReceiver)+"]", &resource.Receiver[numReceiver], htmlAttrs)
}
func (resource *SupplyDelivery) T_SuppliedItemQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("suppliedItem.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("suppliedItem.quantity", resource.SuppliedItem.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_SuppliedItemItemCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("suppliedItem.itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suppliedItem.itemCodeableConcept", resource.SuppliedItem.ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SupplyDelivery) T_SuppliedItemItemReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "suppliedItem.itemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "suppliedItem.itemReference", resource.SuppliedItem.ItemReference, htmlAttrs)
}
