package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

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
	Date                *FhirDateTime        `json:"date,omitempty"`
	Creation            *FhirDateTime        `json:"creation,omitempty"`
	PeriodDate          *FhirDate            `json:"periodDate,omitempty"`
	PeriodPeriod        *Period              `json:"periodPeriod,omitempty"`
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
	ServicedDate              *FhirDate           `json:"servicedDate,omitempty"`
	ServicedPeriod            *Period             `json:"servicedPeriod,omitempty"`
	ChargeItemReference       Reference           `json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept     `json:"chargeItemCodeableConcept"`
	PriceComponent            []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherInvoice Invoice

// struct -> json, automatically add resourceType=Patient
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}

// json -> struct, first reject if resourceType != Invoice
func (r *Invoice) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Invoice" {
		return errors.New("resourceType not Invoice")
	}
	return json.Unmarshal(data, (*OtherInvoice)(r))
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
func (resource *Invoice) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSInvoice_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_CancelledReason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("cancelledReason", nil, htmlAttrs)
	}
	return StringInput("cancelledReason", resource.CancelledReason, htmlAttrs)
}
func (resource *Invoice) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Invoice) T_Recipient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recipient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recipient", resource.Recipient, htmlAttrs)
}
func (resource *Invoice) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Invoice) T_Creation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("creation", nil, htmlAttrs)
	}
	return FhirDateTimeInput("creation", resource.Creation, htmlAttrs)
}
func (resource *Invoice) T_PeriodDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("periodDate", nil, htmlAttrs)
	}
	return FhirDateInput("periodDate", resource.PeriodDate, htmlAttrs)
}
func (resource *Invoice) T_PeriodPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("periodPeriod", nil, htmlAttrs)
	}
	return PeriodInput("periodPeriod", resource.PeriodPeriod, htmlAttrs)
}
func (resource *Invoice) T_Issuer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "issuer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "issuer", resource.Issuer, htmlAttrs)
}
func (resource *Invoice) T_Account(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "account", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account", resource.Account, htmlAttrs)
}
func (resource *Invoice) T_TotalPriceComponent(numTotalPriceComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTotalPriceComponent >= len(resource.TotalPriceComponent) {
		return MonetaryComponentInput("totalPriceComponent["+strconv.Itoa(numTotalPriceComponent)+"]", nil, htmlAttrs)
	}
	return MonetaryComponentInput("totalPriceComponent["+strconv.Itoa(numTotalPriceComponent)+"]", &resource.TotalPriceComponent[numTotalPriceComponent], htmlAttrs)
}
func (resource *Invoice) T_TotalNet(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("totalNet", nil, htmlAttrs)
	}
	return MoneyInput("totalNet", resource.TotalNet, htmlAttrs)
}
func (resource *Invoice) T_TotalGross(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("totalGross", nil, htmlAttrs)
	}
	return MoneyInput("totalGross", resource.TotalGross, htmlAttrs)
}
func (resource *Invoice) T_PaymentTerms(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("paymentTerms", nil, htmlAttrs)
	}
	return StringInput("paymentTerms", resource.PaymentTerms, htmlAttrs)
}
func (resource *Invoice) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Invoice) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_ParticipantActor(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", &resource.Participant[numParticipant].Actor, htmlAttrs)
}
func (resource *Invoice) T_LineItemSequence(numLineItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return IntInput("lineItem["+strconv.Itoa(numLineItem)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("lineItem["+strconv.Itoa(numLineItem)+"].sequence", resource.LineItem[numLineItem].Sequence, htmlAttrs)
}
func (resource *Invoice) T_LineItemServicedDate(numLineItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return FhirDateInput("lineItem["+strconv.Itoa(numLineItem)+"].servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("lineItem["+strconv.Itoa(numLineItem)+"].servicedDate", resource.LineItem[numLineItem].ServicedDate, htmlAttrs)
}
func (resource *Invoice) T_LineItemServicedPeriod(numLineItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return PeriodInput("lineItem["+strconv.Itoa(numLineItem)+"].servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("lineItem["+strconv.Itoa(numLineItem)+"].servicedPeriod", resource.LineItem[numLineItem].ServicedPeriod, htmlAttrs)
}
func (resource *Invoice) T_LineItemChargeItemReference(frs []FhirResource, numLineItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return ReferenceInput(frs, "lineItem["+strconv.Itoa(numLineItem)+"].chargeItemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "lineItem["+strconv.Itoa(numLineItem)+"].chargeItemReference", &resource.LineItem[numLineItem].ChargeItemReference, htmlAttrs)
}
func (resource *Invoice) T_LineItemChargeItemCodeableConcept(numLineItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) {
		return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].chargeItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("lineItem["+strconv.Itoa(numLineItem)+"].chargeItemCodeableConcept", &resource.LineItem[numLineItem].ChargeItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Invoice) T_LineItemPriceComponent(numLineItem int, numPriceComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLineItem >= len(resource.LineItem) || numPriceComponent >= len(resource.LineItem[numLineItem].PriceComponent) {
		return MonetaryComponentInput("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"]", nil, htmlAttrs)
	}
	return MonetaryComponentInput("lineItem["+strconv.Itoa(numLineItem)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"]", &resource.LineItem[numLineItem].PriceComponent[numPriceComponent], htmlAttrs)
}
