package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

func (resource *Invoice) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Invoice.Id", nil)
	}
	return StringInput("Invoice.Id", resource.Id)
}
func (resource *Invoice) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Invoice.ImplicitRules", nil)
	}
	return StringInput("Invoice.ImplicitRules", resource.ImplicitRules)
}
func (resource *Invoice) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Invoice.Language", nil, optionsValueSet)
	}
	return CodeSelect("Invoice.Language", resource.Language, optionsValueSet)
}
func (resource *Invoice) T_Status() templ.Component {
	optionsValueSet := VSInvoice_status

	if resource == nil {
		return CodeSelect("Invoice.Status", nil, optionsValueSet)
	}
	return CodeSelect("Invoice.Status", &resource.Status, optionsValueSet)
}
func (resource *Invoice) T_CancelledReason() templ.Component {

	if resource == nil {
		return StringInput("Invoice.CancelledReason", nil)
	}
	return StringInput("Invoice.CancelledReason", resource.CancelledReason)
}
func (resource *Invoice) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Invoice.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Invoice.Type", resource.Type, optionsValueSet)
}
func (resource *Invoice) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Invoice.Date", nil)
	}
	return StringInput("Invoice.Date", resource.Date)
}
func (resource *Invoice) T_PaymentTerms() templ.Component {

	if resource == nil {
		return StringInput("Invoice.PaymentTerms", nil)
	}
	return StringInput("Invoice.PaymentTerms", resource.PaymentTerms)
}
func (resource *Invoice) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *Invoice) T_ParticipantRole(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Invoice.Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Participant[numParticipant].Role, optionsValueSet)
}
func (resource *Invoice) T_LineItemId(numLineItem int) templ.Component {

	if resource == nil || len(resource.LineItem) >= numLineItem {
		return StringInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Id", nil)
	}
	return StringInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Id", resource.LineItem[numLineItem].Id)
}
func (resource *Invoice) T_LineItemSequence(numLineItem int) templ.Component {

	if resource == nil || len(resource.LineItem) >= numLineItem {
		return IntInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Sequence", nil)
	}
	return IntInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].Sequence", resource.LineItem[numLineItem].Sequence)
}
func (resource *Invoice) T_LineItemPriceComponentId(numLineItem int, numPriceComponent int) templ.Component {

	if resource == nil || len(resource.LineItem) >= numLineItem || len(resource.LineItem[numLineItem].PriceComponent) >= numPriceComponent {
		return StringInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Id", nil)
	}
	return StringInput("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Id", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Id)
}
func (resource *Invoice) T_LineItemPriceComponentType(numLineItem int, numPriceComponent int) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType

	if resource == nil || len(resource.LineItem) >= numLineItem || len(resource.LineItem[numLineItem].PriceComponent) >= numPriceComponent {
		return CodeSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", &resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Type, optionsValueSet)
}
func (resource *Invoice) T_LineItemPriceComponentCode(numLineItem int, numPriceComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LineItem) >= numLineItem || len(resource.LineItem[numLineItem].PriceComponent) >= numPriceComponent {
		return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Code, optionsValueSet)
}
func (resource *Invoice) T_LineItemPriceComponentFactor(numLineItem int, numPriceComponent int) templ.Component {

	if resource == nil || len(resource.LineItem) >= numLineItem || len(resource.LineItem[numLineItem].PriceComponent) >= numPriceComponent {
		return Float64Input("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", nil)
	}
	return Float64Input("Invoice.LineItem["+strconv.Itoa(numLineItem)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Factor)
}
