package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Account
type Account struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Status            string                  `json:"status"`
	BillingStatus     *CodeableConcept        `json:"billingStatus,omitempty"`
	Type              *CodeableConcept        `json:"type,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Subject           []Reference             `json:"subject,omitempty"`
	ServicePeriod     *Period                 `json:"servicePeriod,omitempty"`
	Coverage          []AccountCoverage       `json:"coverage,omitempty"`
	Owner             *Reference              `json:"owner,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Guarantor         []AccountGuarantor      `json:"guarantor,omitempty"`
	Diagnosis         []AccountDiagnosis      `json:"diagnosis,omitempty"`
	Procedure         []AccountProcedure      `json:"procedure,omitempty"`
	RelatedAccount    []AccountRelatedAccount `json:"relatedAccount,omitempty"`
	Currency          *CodeableConcept        `json:"currency,omitempty"`
	Balance           []AccountBalance        `json:"balance,omitempty"`
	CalculatedAt      *string                 `json:"calculatedAt,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountCoverage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Coverage          Reference   `json:"coverage"`
	Priority          *int        `json:"priority,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountGuarantor struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Party             Reference   `json:"party"`
	OnHold            *bool       `json:"onHold,omitempty"`
	Period            *Period     `json:"period,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountDiagnosis struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Sequence          *int              `json:"sequence,omitempty"`
	Condition         CodeableReference `json:"condition"`
	DateOfDiagnosis   *time.Time        `json:"dateOfDiagnosis,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Type              []CodeableConcept `json:"type,omitempty"`
	OnAdmission       *bool             `json:"onAdmission,omitempty"`
	PackageCode       []CodeableConcept `json:"packageCode,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountProcedure struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Sequence          *int              `json:"sequence,omitempty"`
	Code              CodeableReference `json:"code"`
	DateOfService     *time.Time        `json:"dateOfService,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Type              []CodeableConcept `json:"type,omitempty"`
	PackageCode       []CodeableConcept `json:"packageCode,omitempty"`
	Device            []Reference       `json:"device,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountRelatedAccount struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Account           Reference        `json:"account"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Account
type AccountBalance struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Aggregate         *CodeableConcept `json:"aggregate,omitempty"`
	Term              *CodeableConcept `json:"term,omitempty"`
	Estimate          *bool            `json:"estimate,omitempty"`
	Amount            Money            `json:"amount"`
}

type OtherAccount Account

// on convert struct to json, automatically add resourceType=Account
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}
func (r Account) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Account/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Account"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Account) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSAccount_status

	if resource == nil {
		return CodeSelect("Account.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Account.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BillingStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.BillingStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.BillingStatus", resource.BillingStatus, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Account.Name", nil, htmlAttrs)
	}
	return StringInput("Account.Name", resource.Name, htmlAttrs)
}
func (resource *Account) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Account.Description", nil, htmlAttrs)
	}
	return StringInput("Account.Description", resource.Description, htmlAttrs)
}
func (resource *Account) T_Currency(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.Currency", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Currency", resource.Currency, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_CalculatedAt(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Account.CalculatedAt", nil, htmlAttrs)
	}
	return StringInput("Account.CalculatedAt", resource.CalculatedAt, htmlAttrs)
}
func (resource *Account) T_CoveragePriority(numCoverage int, htmlAttrs string) templ.Component {

	if resource == nil || numCoverage >= len(resource.Coverage) {
		return IntInput("Account.Coverage."+strconv.Itoa(numCoverage)+"..Priority", nil, htmlAttrs)
	}
	return IntInput("Account.Coverage."+strconv.Itoa(numCoverage)+"..Priority", resource.Coverage[numCoverage].Priority, htmlAttrs)
}
func (resource *Account) T_GuarantorOnHold(numGuarantor int, htmlAttrs string) templ.Component {

	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return BoolInput("Account.Guarantor."+strconv.Itoa(numGuarantor)+"..OnHold", nil, htmlAttrs)
	}
	return BoolInput("Account.Guarantor."+strconv.Itoa(numGuarantor)+"..OnHold", resource.Guarantor[numGuarantor].OnHold, htmlAttrs)
}
func (resource *Account) T_DiagnosisSequence(numDiagnosis int, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Sequence", nil, htmlAttrs)
	}
	return IntInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Sequence", resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *Account) T_DiagnosisDateOfDiagnosis(numDiagnosis int, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return DateTimeInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..DateOfDiagnosis", nil, htmlAttrs)
	}
	return DateTimeInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..DateOfDiagnosis", resource.Diagnosis[numDiagnosis].DateOfDiagnosis, htmlAttrs)
}
func (resource *Account) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Type."+strconv.Itoa(numType)+".", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_DiagnosisOnAdmission(numDiagnosis int, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return BoolInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..OnAdmission", nil, htmlAttrs)
	}
	return BoolInput("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, htmlAttrs)
}
func (resource *Account) T_DiagnosisPackageCode(numDiagnosis int, numPackageCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numPackageCode >= len(resource.Diagnosis[numDiagnosis].PackageCode) {
		return CodeableConceptSelect("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..PackageCode."+strconv.Itoa(numPackageCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Diagnosis."+strconv.Itoa(numDiagnosis)+"..PackageCode."+strconv.Itoa(numPackageCode)+".", &resource.Diagnosis[numDiagnosis].PackageCode[numPackageCode], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_ProcedureSequence(numProcedure int, htmlAttrs string) templ.Component {

	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("Account.Procedure."+strconv.Itoa(numProcedure)+"..Sequence", nil, htmlAttrs)
	}
	return IntInput("Account.Procedure."+strconv.Itoa(numProcedure)+"..Sequence", resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *Account) T_ProcedureDateOfService(numProcedure int, htmlAttrs string) templ.Component {

	if resource == nil || numProcedure >= len(resource.Procedure) {
		return DateTimeInput("Account.Procedure."+strconv.Itoa(numProcedure)+"..DateOfService", nil, htmlAttrs)
	}
	return DateTimeInput("Account.Procedure."+strconv.Itoa(numProcedure)+"..DateOfService", resource.Procedure[numProcedure].DateOfService, htmlAttrs)
}
func (resource *Account) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("Account.Procedure."+strconv.Itoa(numProcedure)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Procedure."+strconv.Itoa(numProcedure)+"..Type."+strconv.Itoa(numType)+".", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_ProcedurePackageCode(numProcedure int, numPackageCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProcedure >= len(resource.Procedure) || numPackageCode >= len(resource.Procedure[numProcedure].PackageCode) {
		return CodeableConceptSelect("Account.Procedure."+strconv.Itoa(numProcedure)+"..PackageCode."+strconv.Itoa(numPackageCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Procedure."+strconv.Itoa(numProcedure)+"..PackageCode."+strconv.Itoa(numPackageCode)+".", &resource.Procedure[numProcedure].PackageCode[numPackageCode], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_RelatedAccountRelationship(numRelatedAccount int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRelatedAccount >= len(resource.RelatedAccount) {
		return CodeableConceptSelect("Account.RelatedAccount."+strconv.Itoa(numRelatedAccount)+"..Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.RelatedAccount."+strconv.Itoa(numRelatedAccount)+"..Relationship", resource.RelatedAccount[numRelatedAccount].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BalanceAggregate(numBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numBalance >= len(resource.Balance) {
		return CodeableConceptSelect("Account.Balance."+strconv.Itoa(numBalance)+"..Aggregate", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Balance."+strconv.Itoa(numBalance)+"..Aggregate", resource.Balance[numBalance].Aggregate, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BalanceTerm(numBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numBalance >= len(resource.Balance) {
		return CodeableConceptSelect("Account.Balance."+strconv.Itoa(numBalance)+"..Term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Account.Balance."+strconv.Itoa(numBalance)+"..Term", resource.Balance[numBalance].Term, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BalanceEstimate(numBalance int, htmlAttrs string) templ.Component {

	if resource == nil || numBalance >= len(resource.Balance) {
		return BoolInput("Account.Balance."+strconv.Itoa(numBalance)+"..Estimate", nil, htmlAttrs)
	}
	return BoolInput("Account.Balance."+strconv.Itoa(numBalance)+"..Estimate", resource.Balance[numBalance].Estimate, htmlAttrs)
}
