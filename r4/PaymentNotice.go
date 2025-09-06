package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/PaymentNotice
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
	Created           string           `json:"created"`
	Provider          *Reference       `json:"provider,omitempty"`
	Payment           Reference        `json:"payment"`
	PaymentDate       *string          `json:"paymentDate,omitempty"`
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

func (resource *PaymentNotice) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PaymentNotice.Id", nil)
	}
	return StringInput("PaymentNotice.Id", resource.Id)
}
func (resource *PaymentNotice) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PaymentNotice.ImplicitRules", nil)
	}
	return StringInput("PaymentNotice.ImplicitRules", resource.ImplicitRules)
}
func (resource *PaymentNotice) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PaymentNotice.Language", nil, optionsValueSet)
	}
	return CodeSelect("PaymentNotice.Language", resource.Language, optionsValueSet)
}
func (resource *PaymentNotice) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("PaymentNotice.Status", nil, optionsValueSet)
	}
	return CodeSelect("PaymentNotice.Status", &resource.Status, optionsValueSet)
}
func (resource *PaymentNotice) T_Created() templ.Component {

	if resource == nil {
		return StringInput("PaymentNotice.Created", nil)
	}
	return StringInput("PaymentNotice.Created", &resource.Created)
}
func (resource *PaymentNotice) T_PaymentDate() templ.Component {

	if resource == nil {
		return StringInput("PaymentNotice.PaymentDate", nil)
	}
	return StringInput("PaymentNotice.PaymentDate", resource.PaymentDate)
}
func (resource *PaymentNotice) T_PaymentStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentNotice.PaymentStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentNotice.PaymentStatus", resource.PaymentStatus, optionsValueSet)
}
