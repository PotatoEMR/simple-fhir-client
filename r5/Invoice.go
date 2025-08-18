//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/Invoice
type Invoice struct {
	Id                  *string              `json:"id,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	ImplicitRules       *string              `json:"implicitRules,omitempty"`
	Language            *string              `json:"language,omitempty"`
	Text                *Narrative           `json:"text,omitempty"`
	Contained           []Resource           `json:"contained,omitempty"`
	Extension           []Extension          `json:"extension,omitempty"`
	ModifierExtension   []Extension          `json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `json:"identifier,omitempty"`
	Status              string               `json:"status"`
	CancelledReason     *string              `json:"cancelledReason,omitempty"`
	Type                *CodeableConcept     `json:"type,omitempty"`
	Subject             *Reference           `json:"subject,omitempty"`
	Recipient           *Reference           `json:"recipient,omitempty"`
	Date                *string              `json:"date,omitempty"`
	Creation            *string              `json:"creation,omitempty"`
	PeriodDate          *string              `json:"periodDate"`
	PeriodPeriod        *Period              `json:"periodPeriod"`
	Participant         []InvoiceParticipant `json:"participant,omitempty"`
	Issuer              *Reference           `json:"issuer,omitempty"`
	Account             *Reference           `json:"account,omitempty"`
	LineItem            []InvoiceLineItem    `json:"lineItem,omitempty"`
	TotalPriceComponent []MonetaryComponent  `json:"totalPriceComponent,omitempty"`
	TotalNet            *Money               `json:"totalNet,omitempty"`
	TotalGross          *Money               `json:"totalGross,omitempty"`
	PaymentTerms        *string              `json:"paymentTerms,omitempty"`
	Note                []Annotation         `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Invoice
type InvoiceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Invoice
type InvoiceLineItem struct {
	Id                        *string             `json:"id,omitempty"`
	Extension                 []Extension         `json:"extension,omitempty"`
	ModifierExtension         []Extension         `json:"modifierExtension,omitempty"`
	Sequence                  *int                `json:"sequence,omitempty"`
	ServicedDate              *string             `json:"servicedDate"`
	ServicedPeriod            *Period             `json:"servicedPeriod"`
	ChargeItemReference       Reference           `json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept     `json:"chargeItemCodeableConcept"`
	PriceComponent            []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherInvoice Invoice

// on convert struct to json, automatically add resourceType=Invoice
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}
