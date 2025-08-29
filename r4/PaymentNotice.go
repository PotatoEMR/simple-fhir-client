package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *PaymentNotice) PaymentNoticeLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *PaymentNotice) PaymentNoticeStatus() templ.Component {
	optionsValueSet := VSFm_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
