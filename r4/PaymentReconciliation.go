package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Created           string                             `json:"created"`
	PaymentIssuer     *Reference                         `json:"paymentIssuer,omitempty"`
	Request           *Reference                         `json:"request,omitempty"`
	Requestor         *Reference                         `json:"requestor,omitempty"`
	Outcome           *string                            `json:"outcome,omitempty"`
	Disposition       *string                            `json:"disposition,omitempty"`
	PaymentDate       string                             `json:"paymentDate"`
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
	Date              *string         `json:"date,omitempty"`
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

func (resource *PaymentReconciliation) PaymentReconciliationLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *PaymentReconciliation) PaymentReconciliationStatus() templ.Component {
	optionsValueSet := VSFm_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *PaymentReconciliation) PaymentReconciliationOutcome() templ.Component {
	optionsValueSet := VSRemittance_outcome
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Outcome
	}
	return CodeSelect("outcome", currentVal, optionsValueSet)
}
func (resource *PaymentReconciliation) PaymentReconciliationProcessNoteType(numProcessNote int) templ.Component {
	optionsValueSet := VSNote_type
	currentVal := ""
	if resource != nil && len(resource.ProcessNote) >= numProcessNote {
		currentVal = *resource.ProcessNote[numProcessNote].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
