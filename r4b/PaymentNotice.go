package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/PaymentNotice
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
	Created           FhirDateTime     `json:"created"`
	Provider          *Reference       `json:"provider,omitempty"`
	Payment           Reference        `json:"payment"`
	PaymentDate       *FhirDate        `json:"paymentDate,omitempty"`
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
func (resource *PaymentNotice) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PaymentNotice) T_Request(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("request", nil, htmlAttrs)
	}
	return ReferenceInput("request", resource.Request, htmlAttrs)
}
func (resource *PaymentNotice) T_Response(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("response", nil, htmlAttrs)
	}
	return ReferenceInput("response", resource.Response, htmlAttrs)
}
func (resource *PaymentNotice) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *PaymentNotice) T_Provider(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("provider", nil, htmlAttrs)
	}
	return ReferenceInput("provider", resource.Provider, htmlAttrs)
}
func (resource *PaymentNotice) T_Payment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("payment", nil, htmlAttrs)
	}
	return ReferenceInput("payment", &resource.Payment, htmlAttrs)
}
func (resource *PaymentNotice) T_PaymentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("paymentDate", nil, htmlAttrs)
	}
	return FhirDateInput("paymentDate", resource.PaymentDate, htmlAttrs)
}
func (resource *PaymentNotice) T_Payee(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("payee", nil, htmlAttrs)
	}
	return ReferenceInput("payee", resource.Payee, htmlAttrs)
}
func (resource *PaymentNotice) T_Recipient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("recipient", nil, htmlAttrs)
	}
	return ReferenceInput("recipient", &resource.Recipient, htmlAttrs)
}
func (resource *PaymentNotice) T_Amount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("amount", nil, htmlAttrs)
	}
	return MoneyInput("amount", &resource.Amount, htmlAttrs)
}
func (resource *PaymentNotice) T_PaymentStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("paymentStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("paymentStatus", resource.PaymentStatus, optionsValueSet, htmlAttrs)
}
