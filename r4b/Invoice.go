package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Invoice
type Invoice struct {
	Id                *string              `json:"id,omitempty"`
	Meta              *Meta                `json:"meta,omitempty"`
	ImplicitRules     *string              `json:"implicitRules,omitempty"`
	Language          *string              `json:"language,omitempty"`
	Text              *Narrative           `json:"text,omitempty"`
	Contained         []Resource           `json:"contained,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Identifier        []Identifier         `json:"identifier,omitempty"`
	Status            string               `json:"status"`
	CancelledReason   *string              `json:"cancelledReason,omitempty"`
	Type              *CodeableConcept     `json:"type,omitempty"`
	Subject           *Reference           `json:"subject,omitempty"`
	Recipient         *Reference           `json:"recipient,omitempty"`
	Date              *string              `json:"date,omitempty"`
	Participant       []InvoiceParticipant `json:"participant,omitempty"`
	Issuer            *Reference           `json:"issuer,omitempty"`
	Account           *Reference           `json:"account,omitempty"`
	LineItem          []InvoiceLineItem    `json:"lineItem,omitempty"`
	TotalNet          *Money               `json:"totalNet,omitempty"`
	TotalGross        *Money               `json:"totalGross,omitempty"`
	PaymentTerms      *string              `json:"paymentTerms,omitempty"`
	Note              []Annotation         `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Invoice
type InvoiceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Invoice
type InvoiceLineItem struct {
	Id                        *string                         `json:"id,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Sequence                  *int                            `json:"sequence,omitempty"`
	ChargeItemReference       Reference                       `json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept                 `json:"chargeItemCodeableConcept"`
	PriceComponent            []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Invoice
type InvoiceLineItemPriceComponent struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              string           `json:"type"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Factor            *float64         `json:"factor,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
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

func (resource *Invoice) InvoiceLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Invoice) InvoiceStatus() templ.Component {
	optionsValueSet := VSInvoice_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Invoice) InvoiceLineItemPriceComponentType(numLineItem int, numPriceComponent int) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType
	currentVal := ""
	if resource != nil && len(resource.LineItem[numLineItem].PriceComponent) >= numPriceComponent {
		currentVal = resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
