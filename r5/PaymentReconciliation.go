package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Created           string                             `json:"created"`
	Enterer           *Reference                         `json:"enterer,omitempty"`
	IssuerType        *CodeableConcept                   `json:"issuerType,omitempty"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	Date              string                             `json:"date"`
	Location          *Reference                         `json:"location,omitempty"`
	Method            *CodeableConcept                   `json:"method,omitempty"`
	CardBrand         *string                            `json:"cardBrand,omitempty"`
	AccountNumber     *string                            `json:"accountNumber,omitempty"`
	ExpirationDate    *string                            `json:"expirationDate,omitempty"`
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
	Date                  *string          `json:"date,omitempty"`
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

func (resource *PaymentReconciliation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Id", nil)
	}
	return StringInput("PaymentReconciliation.Id", resource.Id)
}
func (resource *PaymentReconciliation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.ImplicitRules", nil)
	}
	return StringInput("PaymentReconciliation.ImplicitRules", resource.ImplicitRules)
}
func (resource *PaymentReconciliation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Language", nil, optionsValueSet)
	}
	return CodeSelect("PaymentReconciliation.Language", resource.Language, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.Type", &resource.Type, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Status", nil, optionsValueSet)
	}
	return CodeSelect("PaymentReconciliation.Status", &resource.Status, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Kind(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Kind", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.Kind", resource.Kind, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Created() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Created", nil)
	}
	return StringInput("PaymentReconciliation.Created", &resource.Created)
}
func (resource *PaymentReconciliation) T_IssuerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.IssuerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.IssuerType", resource.IssuerType, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Outcome() templ.Component {
	optionsValueSet := VSPayment_outcome

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("PaymentReconciliation.Outcome", resource.Outcome, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Disposition() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Disposition", nil)
	}
	return StringInput("PaymentReconciliation.Disposition", resource.Disposition)
}
func (resource *PaymentReconciliation) T_Date() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Date", nil)
	}
	return StringInput("PaymentReconciliation.Date", &resource.Date)
}
func (resource *PaymentReconciliation) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.Method", resource.Method, optionsValueSet)
}
func (resource *PaymentReconciliation) T_CardBrand() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.CardBrand", nil)
	}
	return StringInput("PaymentReconciliation.CardBrand", resource.CardBrand)
}
func (resource *PaymentReconciliation) T_AccountNumber() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.AccountNumber", nil)
	}
	return StringInput("PaymentReconciliation.AccountNumber", resource.AccountNumber)
}
func (resource *PaymentReconciliation) T_ExpirationDate() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.ExpirationDate", nil)
	}
	return StringInput("PaymentReconciliation.ExpirationDate", resource.ExpirationDate)
}
func (resource *PaymentReconciliation) T_Processor() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Processor", nil)
	}
	return StringInput("PaymentReconciliation.Processor", resource.Processor)
}
func (resource *PaymentReconciliation) T_ReferenceNumber() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.ReferenceNumber", nil)
	}
	return StringInput("PaymentReconciliation.ReferenceNumber", resource.ReferenceNumber)
}
func (resource *PaymentReconciliation) T_Authorization() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Authorization", nil)
	}
	return StringInput("PaymentReconciliation.Authorization", resource.Authorization)
}
func (resource *PaymentReconciliation) T_FormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.FormCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.FormCode", resource.FormCode, optionsValueSet)
}
func (resource *PaymentReconciliation) T_AllocationId(numAllocation int) templ.Component {

	if resource == nil || len(resource.Allocation) >= numAllocation {
		return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Id", nil)
	}
	return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Id", resource.Allocation[numAllocation].Id)
}
func (resource *PaymentReconciliation) T_AllocationType(numAllocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Allocation) >= numAllocation {
		return CodeableConceptSelect("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Type", resource.Allocation[numAllocation].Type, optionsValueSet)
}
func (resource *PaymentReconciliation) T_AllocationDate(numAllocation int) templ.Component {

	if resource == nil || len(resource.Allocation) >= numAllocation {
		return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Date", nil)
	}
	return StringInput("PaymentReconciliation.Allocation["+strconv.Itoa(numAllocation)+"].Date", resource.Allocation[numAllocation].Date)
}
func (resource *PaymentReconciliation) T_ProcessNoteId(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", nil)
	}
	return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", resource.ProcessNote[numProcessNote].Id)
}
func (resource *PaymentReconciliation) T_ProcessNoteType(numProcessNote int) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return CodeSelect("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *PaymentReconciliation) T_ProcessNoteText(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", nil)
	}
	return StringInput("PaymentReconciliation.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", resource.ProcessNote[numProcessNote].Text)
}
