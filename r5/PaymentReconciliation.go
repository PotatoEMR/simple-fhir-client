package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
	Id                *string                            `json:"id,omitempty"`
	Meta              *Meta                              `json:"meta,omitempty"`
	ImplicitRules     *string                            `json:"implicitRules,omitempty"`
	Language          *string                            `json:"language,omitempty"`
	Text              *Narrative                         `json:"text,omitempty"`
	Contained         []Resource                         `json:"contained,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                       `json:"identifier,omitempty"`
	Type              CodeableConcept                    `json:"type"`
	Status            string                             `json:"status"`
	Kind              *CodeableConcept                   `json:"kind,omitempty"`
	Period            *Period                            `json:"period,omitempty"`
	Created           time.Time                          `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Enterer           *Reference                         `json:"enterer,omitempty"`
	IssuerType        *CodeableConcept                   `json:"issuerType,omitempty"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	Date              time.Time                          `json:"date,format:'2006-01-02'"`
	Location          *Reference                         `json:"location,omitempty"`
	Method            *CodeableConcept                   `json:"method,omitempty"`
	CardBrand         *string                            `json:"cardBrand,omitempty"`
	AccountNumber     *string                            `json:"accountNumber,omitempty"`
	ExpirationDate    *time.Time                         `json:"expirationDate,omitempty,format:'2006-01-02'"`
	Processor         *string                            `json:"processor,omitempty"`
	ReferenceNumber   *string                            `json:"referenceNumber,omitempty"`
	Authorization     *string                            `json:"authorization,omitempty"`
	TenderedAmount    *Money                             `json:"tenderedAmount,omitempty"`
	ReturnedAmount    *Money                             `json:"returnedAmount,omitempty"`
	Amount            Money                              `json:"amount"`
	PaymentIdentifier *Identifier                        `json:"paymentIdentifier,omitempty"`
	Allocation        []PaymentReconciliationAllocation  `json:"allocation,omitempty"`
	FormCode          *CodeableConcept                   `json:"formCode,omitempty"`
	ProcessNote       []PaymentReconciliationProcessNote `json:"processNote,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PaymentReconciliation
type PaymentReconciliationAllocation struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Identifier            *Identifier      `json:"identifier,omitempty"`
	Predecessor           *Identifier      `json:"predecessor,omitempty"`
	Target                *Reference       `json:"target,omitempty"`
	TargetItemString      *string          `json:"targetItemString,omitempty"`
	TargetItemIdentifier  *Identifier      `json:"targetItemIdentifier,omitempty"`
	TargetItemPositiveInt *int             `json:"targetItemPositiveInt,omitempty"`
	Encounter             *Reference       `json:"encounter,omitempty"`
	Account               *Reference       `json:"account,omitempty"`
	Type                  *CodeableConcept `json:"type,omitempty"`
	Submitter             *Reference       `json:"submitter,omitempty"`
	Response              *Reference       `json:"response,omitempty"`
	Date                  *time.Time       `json:"date,omitempty,format:'2006-01-02'"`
	Responsible           *Reference       `json:"responsible,omitempty"`
	Payee                 *Reference       `json:"payee,omitempty"`
	Amount                *Money           `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PaymentReconciliation
type PaymentReconciliationProcessNote struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              *string     `json:"type,omitempty"`
	Text              *string     `json:"text,omitempty"`
}

type OtherPaymentReconciliation PaymentReconciliation

// on convert struct to json, automatically add resourceType=PaymentReconciliation
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
}
func (r PaymentReconciliation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PaymentReconciliation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PaymentReconciliation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PaymentReconciliation) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PaymentReconciliation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Kind(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.Kind", resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PaymentReconciliation.Created", nil, htmlAttrs)
	}
	return DateTimeInput("PaymentReconciliation.Created", &resource.Created, htmlAttrs)
}
func (resource *PaymentReconciliation) T_IssuerType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.IssuerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.IssuerType", resource.IssuerType, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSPayment_outcome

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PaymentReconciliation.Outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.Disposition", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.Disposition", resource.Disposition, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentReconciliation.Date", nil, htmlAttrs)
	}
	return DateInput("PaymentReconciliation.Date", &resource.Date, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Method(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.Method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_CardBrand(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.CardBrand", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.CardBrand", resource.CardBrand, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AccountNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.AccountNumber", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.AccountNumber", resource.AccountNumber, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ExpirationDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentReconciliation.ExpirationDate", nil, htmlAttrs)
	}
	return DateInput("PaymentReconciliation.ExpirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Processor(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.Processor", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.Processor", resource.Processor, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ReferenceNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.ReferenceNumber", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.ReferenceNumber", resource.ReferenceNumber, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Authorization(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PaymentReconciliation.Authorization", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.Authorization", resource.Authorization, htmlAttrs)
}
func (resource *PaymentReconciliation) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.FormCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.FormCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTargetItemString(numAllocation int, htmlAttrs string) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].TargetItemString", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].TargetItemString", resource.Allocation[numAllocation].TargetItemString, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTargetItemPositiveInt(numAllocation int, htmlAttrs string) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return IntInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].TargetItemPositiveInt", nil, htmlAttrs)
	}
	return IntInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].TargetItemPositiveInt", resource.Allocation[numAllocation].TargetItemPositiveInt, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationType(numAllocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return CodeableConceptSelect("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Type", resource.Allocation[numAllocation].Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationDate(numAllocation int, htmlAttrs string) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return DateInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Date", nil, htmlAttrs)
	}
	return DateInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Date", resource.Allocation[numAllocation].Date, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ProcessNoteType(numProcessNote int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ProcessNoteText(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", nil, htmlAttrs)
	}
	return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", resource.ProcessNote[numProcessNote].Text, htmlAttrs)
}
