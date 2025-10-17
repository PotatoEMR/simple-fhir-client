package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/PaymentReconciliation
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
	Status            string                             `json:"status"`
	Period            *Period                            `json:"period,omitempty"`
	Created           FhirDateTime                       `json:"created"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	PaymentDate       FhirDate                           `json:"paymentDate"`
	PaymentAmount     Money                              `json:"paymentAmount"`
	PaymentIdentifier *Identifier                        `json:"paymentIdentifier,omitempty"`
	Detail            []PaymentReconciliationDetail      `json:"detail,omitempty"`
	FormCode          *CodeableConcept                   `json:"formCode,omitempty"`
	ProcessNote       []PaymentReconciliationProcessNote `json:"processNote,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PaymentReconciliation
type PaymentReconciliationDetail struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        *Identifier     `json:"identifier,omitempty"`
	Predecessor       *Identifier     `json:"predecessor,omitempty"`
	Type              CodeableConcept `json:"type"`
	Request           *Reference      `json:"request,omitempty"`
	Submitter         *Reference      `json:"submitter,omitempty"`
	Response          *Reference      `json:"response,omitempty"`
	Date              *FhirDate       `json:"date,omitempty"`
	Responsible       *Reference      `json:"responsible,omitempty"`
	Payee             *Reference      `json:"payee,omitempty"`
	Amount            *Money          `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PaymentReconciliation
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
func (r PaymentReconciliation) ResourceType() string {
	return "PaymentReconciliation"
}

func (resource *PaymentReconciliation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
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
	optionsValueSet := VSRemittance_outcome

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
func (resource *PaymentReconciliation) T_PaymentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("paymentDate", nil, htmlAttrs)
	}
	return FhirDateInput("paymentDate", &resource.PaymentDate, htmlAttrs)
}
func (resource *PaymentReconciliation) T_PaymentAmount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("paymentAmount", nil, htmlAttrs)
	}
	return MoneyInput("paymentAmount", &resource.PaymentAmount, htmlAttrs)
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
func (resource *PaymentReconciliation) T_DetailPredecessor(numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return IdentifierInput("detail["+strconv.Itoa(numDetail)+"].predecessor", nil, htmlAttrs)
	}
	return IdentifierInput("detail["+strconv.Itoa(numDetail)+"].predecessor", resource.Detail[numDetail].Predecessor, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailType(numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return CodeableConceptSelect("detail["+strconv.Itoa(numDetail)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("detail["+strconv.Itoa(numDetail)+"].type", &resource.Detail[numDetail].Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailRequest(frs []FhirResource, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].request", resource.Detail[numDetail].Request, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailSubmitter(frs []FhirResource, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].submitter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].submitter", resource.Detail[numDetail].Submitter, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailResponse(frs []FhirResource, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].response", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].response", resource.Detail[numDetail].Response, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailDate(numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return FhirDateInput("detail["+strconv.Itoa(numDetail)+"].date", nil, htmlAttrs)
	}
	return FhirDateInput("detail["+strconv.Itoa(numDetail)+"].date", resource.Detail[numDetail].Date, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailResponsible(frs []FhirResource, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].responsible", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].responsible", resource.Detail[numDetail].Responsible, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailPayee(frs []FhirResource, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].payee", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "detail["+strconv.Itoa(numDetail)+"].payee", resource.Detail[numDetail].Payee, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailAmount(numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return MoneyInput("detail["+strconv.Itoa(numDetail)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("detail["+strconv.Itoa(numDetail)+"].amount", resource.Detail[numDetail].Amount, htmlAttrs)
}
func (resource *PaymentReconciliation) T_ProcessNoteType(numProcessNote int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
