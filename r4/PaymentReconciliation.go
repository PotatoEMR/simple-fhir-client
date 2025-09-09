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

// http://hl7.org/fhir/r4/StructureDefinition/PaymentReconciliation
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
	Created           time.Time                          `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	PaymentDate       time.Time                          `json:"paymentDate,format:'2006-01-02'"`
	PaymentAmount     Money                              `json:"paymentAmount"`
	PaymentIdentifier *Identifier                        `json:"paymentIdentifier,omitempty"`
	Detail            []PaymentReconciliationDetail      `json:"detail,omitempty"`
	FormCode          *CodeableConcept                   `json:"formCode,omitempty"`
	ProcessNote       []PaymentReconciliationProcessNote `json:"processNote,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PaymentReconciliation
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
	Date              *time.Time      `json:"date,omitempty,format:'2006-01-02'"`
	Responsible       *Reference      `json:"responsible,omitempty"`
	Payee             *Reference      `json:"payee,omitempty"`
	Amount            *Money          `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PaymentReconciliation
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
func (resource *PaymentReconciliation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PaymentReconciliation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PaymentReconciliation.Created", nil, htmlAttrs)
	}
	return DateTimeInput("PaymentReconciliation.Created", &resource.Created, htmlAttrs)
}
func (resource *PaymentReconciliation) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSRemittance_outcome

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
func (resource *PaymentReconciliation) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentReconciliation.PaymentDate", nil, htmlAttrs)
	}
	return DateInput("PaymentReconciliation.PaymentDate", &resource.PaymentDate, htmlAttrs)
}
func (resource *PaymentReconciliation) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.FormCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.FormCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailType(numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return CodeableConceptSelect("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Detail[numDetail].Type, optionsValueSet, htmlAttrs)
}
func (resource *PaymentReconciliation) T_DetailDate(numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numDetail >= len(resource.Detail) {
		return DateInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Date", nil, htmlAttrs)
	}
	return DateInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Date", resource.Detail[numDetail].Date, htmlAttrs)
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
