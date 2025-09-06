package r4b

//generated with command go run ./bultaoreune -nodownload
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
	Date              *string         `json:"date,omitempty"`
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
func (resource *PaymentReconciliation) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("PaymentReconciliation.Status", nil, optionsValueSet)
	}
	return CodeSelect("PaymentReconciliation.Status", &resource.Status, optionsValueSet)
}
func (resource *PaymentReconciliation) T_Created() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.Created", nil)
	}
	return StringInput("PaymentReconciliation.Created", &resource.Created)
}
func (resource *PaymentReconciliation) T_Outcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

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
func (resource *PaymentReconciliation) T_PaymentDate() templ.Component {

	if resource == nil {
		return StringInput("PaymentReconciliation.PaymentDate", nil)
	}
	return StringInput("PaymentReconciliation.PaymentDate", &resource.PaymentDate)
}
func (resource *PaymentReconciliation) T_FormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PaymentReconciliation.FormCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.FormCode", resource.FormCode, optionsValueSet)
}
func (resource *PaymentReconciliation) T_DetailId(numDetail int) templ.Component {

	if resource == nil || len(resource.Detail) >= numDetail {
		return StringInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Id", resource.Detail[numDetail].Id)
}
func (resource *PaymentReconciliation) T_DetailType(numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Detail) >= numDetail {
		return CodeableConceptSelect("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Detail[numDetail].Type, optionsValueSet)
}
func (resource *PaymentReconciliation) T_DetailDate(numDetail int) templ.Component {

	if resource == nil || len(resource.Detail) >= numDetail {
		return StringInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Date", nil)
	}
	return StringInput("PaymentReconciliation.Detail["+strconv.Itoa(numDetail)+"].Date", resource.Detail[numDetail].Date)
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
