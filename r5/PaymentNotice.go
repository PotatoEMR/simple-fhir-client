package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string          `json:"id,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	ImplicitRules     *string          `json:"implicitRules,omitempty"`
	Language          *string          `json:"language,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	Contained         []Resource       `json:"contained,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `json:"identifier,omitempty"`
	Status            string           `json:"status"`
	Request           *Reference       `json:"request,omitempty"`
	Response          *Reference       `json:"response,omitempty"`
	Created           time.Time        `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Reporter          *Reference       `json:"reporter,omitempty"`
	Payment           *Reference       `json:"payment,omitempty"`
	PaymentDate       *time.Time       `json:"paymentDate,omitempty,format:'2006-01-02'"`
	Payee             *Reference       `json:"payee,omitempty"`
	Recipient         Reference        `json:"recipient"`
	Amount            Money            `json:"amount"`
	PaymentStatus     *CodeableConcept `json:"paymentStatus,omitempty"`
}

type OtherPaymentNotice PaymentNotice

// on convert struct to json, automatically add resourceType=PaymentNotice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
}
func (r PaymentNotice) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PaymentNotice/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PaymentNotice"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PaymentNotice) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PaymentNotice) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Created", nil, htmlAttrs)
	}
	return DateTimeInput("Created", &resource.Created, htmlAttrs)
}
func (resource *PaymentNotice) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentDate", nil, htmlAttrs)
	}
	return DateInput("PaymentDate", resource.PaymentDate, htmlAttrs)
}
func (resource *PaymentNotice) T_PaymentStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentStatus", resource.PaymentStatus, optionsValueSet, htmlAttrs)
}
