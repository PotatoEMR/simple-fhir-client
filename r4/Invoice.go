package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Invoice
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
	Date              *time.Time           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Participant       []InvoiceParticipant `json:"participant,omitempty"`
	Issuer            *Reference           `json:"issuer,omitempty"`
	Account           *Reference           `json:"account,omitempty"`
	LineItem          []InvoiceLineItem    `json:"lineItem,omitempty"`
	TotalNet          *Money               `json:"totalNet,omitempty"`
	TotalGross        *Money               `json:"totalGross,omitempty"`
	PaymentTerms      *string              `json:"paymentTerms,omitempty"`
	Note              []Annotation         `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Invoice
type InvoiceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Invoice
type InvoiceLineItem struct {
	Id                        *string                         `json:"id,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Sequence                  *int                            `json:"sequence,omitempty"`
	ChargeItemReference       Reference                       `json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept                 `json:"chargeItemCodeableConcept"`
	PriceComponent            []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Invoice
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
func (r Invoice) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Invoice/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Invoice"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Invoice) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSInvoice_status

	if resource == nil {
		return CodeSelect("Invoice.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Invoice.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_CancelledReason(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Invoice.CancelledReason", nil, htmlAttrs)
	}
	return StringInput("Invoice.CancelledReason", resource.CancelledReason, htmlAttrs)
}
func (resource *Invoice) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Invoice.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Invoice.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Invoice.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Invoice.Date", resource.Date, htmlAttrs)
}
func (resource *Invoice) T_PaymentTerms(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Invoice.PaymentTerms", nil, htmlAttrs)
	}
	return StringInput("Invoice.PaymentTerms", resource.PaymentTerms, htmlAttrs)
}
func (resource *Invoice) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Invoice.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Invoice.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Invoice) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemSequence(numLineItem int, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return IntInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Sequence", resource.LineItem[numLineItem].Sequence, htmlAttrs)
}
func (resource *Invoice) T_LineItemChargeItemCodeableConcept(numLineItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].ChargeItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].ChargeItemCodeableConcept", &resource.LineItem[numLineItem].ChargeItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentType(numLineItem int, numPriceComponent int, htmlAttrs string) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType

	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return CodeSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", &resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentCode(numLineItem int, numPriceComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentFactor(numLineItem int, numPriceComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return Float64Input("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Factor, htmlAttrs)
}
