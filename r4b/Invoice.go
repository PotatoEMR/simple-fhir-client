package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_CancelledReason(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("cancelledReason", nil, htmlAttrs)
	}
	return StringInput("cancelledReason", resource.CancelledReason, htmlAttrs)
}
func (resource *Invoice) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Invoice) T_PaymentTerms(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("paymentTerms", nil, htmlAttrs)
	}
	return StringInput("paymentTerms", resource.PaymentTerms, htmlAttrs)
}
func (resource *Invoice) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Invoice) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemSequence(numLineItem int, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return IntInput("lineItem["+strconv.Itoa(numLineItem)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("lineItem["+strconv.Itoa(numLineItem)+"].sequence", resource.LineItem[numLineItem].Sequence, htmlAttrs)
}
func (resource *Invoice) T_LineItemChargeItemCodeableConcept(numLineItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].chargeItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].chargeItemCodeableConcept", &resource.LineItem[numLineItem].ChargeItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentType(numLineItem int, numPriceComponent int, htmlAttrs string) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType

	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return CodeSelect("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].type", &resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentCode(numLineItem int, numPriceComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].code", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponentFactor(numLineItem int, numPriceComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return Float64Input("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].factor", resource.LineItem[numLineItem].PriceComponent[numPriceComponent].Factor, htmlAttrs)
}
