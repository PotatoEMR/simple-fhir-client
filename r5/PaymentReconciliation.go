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
	Created           FhirDateTime                       `json:"created"`
	Enterer           *Reference                         `json:"enterer,omitempty"`
	IssuerType        *CodeableConcept                   `json:"issuerType,omitempty"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	Date              FhirDate                           `json:"date"`
	Location          *Reference                         `json:"location,omitempty"`
	Method            *CodeableConcept                   `json:"method,omitempty"`
	CardBrand         *string                            `json:"cardBrand,omitempty"`
	AccountNumber     *string                            `json:"accountNumber,omitempty"`
	ExpirationDate    *FhirDate                          `json:"expirationDate,omitempty"`
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
	Date                  *FhirDate        `json:"date,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
}

// json -> struct, first reject if resourceType != PaymentReconciliation
func (r *PaymentReconciliation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "PaymentReconciliation" {
		return errors.New("resourceType not PaymentReconciliation")
	}
	return json.Unmarshal(data, (*OtherPaymentReconciliation)(r))
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
func (resource *PaymentReconciliation) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Kind(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("kind", resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Enterer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "enterer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "enterer", resource.Enterer, htmlAttrs)
}
func (resource *PaymentReconciliation) T_IssuerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("issuerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("issuerType", resource.IssuerType, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_PaymentIssuer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "paymentIssuer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "paymentIssuer", resource.PaymentIssuer, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Request(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request", resource.Request, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Requestor(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requestor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requestor", resource.Requestor, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPayment_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("date", nil, htmlAttrs)
	}
	return FhirDateInput("date", &resource.Date, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Method(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_CardBrand(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("cardBrand", nil, htmlAttrs)
	}
	return StringInput("cardBrand", resource.CardBrand, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AccountNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("accountNumber", nil, htmlAttrs)
	}
	return StringInput("accountNumber", resource.AccountNumber, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("expirationDate", nil, htmlAttrs)
	}
	return FhirDateInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Processor(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("processor", nil, htmlAttrs)
	}
	return StringInput("processor", resource.Processor, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ReferenceNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("referenceNumber", nil, htmlAttrs)
	}
	return StringInput("referenceNumber", resource.ReferenceNumber, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Authorization(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("authorization", nil, htmlAttrs)
	}
	return StringInput("authorization", resource.Authorization, htmlAttrs)
}
func (resource *PaymentReconciliation) T_TenderedAmount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("tenderedAmount", nil, htmlAttrs)
	}
	return MoneyInput("tenderedAmount", resource.TenderedAmount, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ReturnedAmount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("returnedAmount", nil, htmlAttrs)
	}
	return MoneyInput("returnedAmount", resource.ReturnedAmount, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Amount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("amount", nil, htmlAttrs)
	}
	return MoneyInput("amount", &resource.Amount, htmlAttrs)
}
func (resource *PaymentReconciliation) T_PaymentIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("paymentIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("paymentIdentifier", resource.PaymentIdentifier, htmlAttrs)
}
func (resource *PaymentReconciliation) T_FormCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationPredecessor(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return IdentifierInput("allocation["+strconv.Itoa(numAllocation)+"].predecessor", nil, htmlAttrs)
	}
	return IdentifierInput("allocation["+strconv.Itoa(numAllocation)+"].predecessor", resource.Allocation[numAllocation].Predecessor, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTarget(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].target", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].target", resource.Allocation[numAllocation].Target, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTargetItemString(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return StringInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemString", nil, htmlAttrs)
	}
	return StringInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemString", resource.Allocation[numAllocation].TargetItemString, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTargetItemIdentifier(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return IdentifierInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemIdentifier", resource.Allocation[numAllocation].TargetItemIdentifier, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationTargetItemPositiveInt(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return IntInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemPositiveInt", nil, htmlAttrs)
	}
	return IntInput("allocation["+strconv.Itoa(numAllocation)+"].targetItemPositiveInt", resource.Allocation[numAllocation].TargetItemPositiveInt, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationEncounter(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].encounter", resource.Allocation[numAllocation].Encounter, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationAccount(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].account", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].account", resource.Allocation[numAllocation].Account, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationType(numAllocation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return CodeableConceptSelect("allocation["+strconv.Itoa(numAllocation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("allocation["+strconv.Itoa(numAllocation)+"].type", resource.Allocation[numAllocation].Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationSubmitter(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].submitter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].submitter", resource.Allocation[numAllocation].Submitter, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationResponse(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].response", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].response", resource.Allocation[numAllocation].Response, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationDate(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return FhirDateInput("allocation["+strconv.Itoa(numAllocation)+"].date", nil, htmlAttrs)
	}
	return FhirDateInput("allocation["+strconv.Itoa(numAllocation)+"].date", resource.Allocation[numAllocation].Date, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationResponsible(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].responsible", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].responsible", resource.Allocation[numAllocation].Responsible, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationPayee(frs []FhirResource, numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].payee", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allocation["+strconv.Itoa(numAllocation)+"].payee", resource.Allocation[numAllocation].Payee, htmlAttrs)
}
func (resource *PaymentReconciliation) T_AllocationAmount(numAllocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllocation >= len(resource.Allocation) {
		return MoneyInput("allocation["+strconv.Itoa(numAllocation)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("allocation["+strconv.Itoa(numAllocation)+"].amount", resource.Allocation[numAllocation].Amount, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ProcessNoteType(numProcessNote int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
