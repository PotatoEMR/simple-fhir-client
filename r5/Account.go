package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	DateOfDiagnosis   *string           `json:"dateOfDiagnosis,omitempty"`
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
	DateOfService     *string           `json:"dateOfService,omitempty"`
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

func (resource *Account) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Account.Id", nil)
	}
	return StringInput("Account.Id", resource.Id)
}
func (resource *Account) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Account.ImplicitRules", nil)
	}
	return StringInput("Account.ImplicitRules", resource.ImplicitRules)
}
func (resource *Account) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Account.Language", nil, optionsValueSet)
	}
	return CodeSelect("Account.Language", resource.Language, optionsValueSet)
}
func (resource *Account) T_Status() templ.Component {
	optionsValueSet := VSAccount_status

	if resource == nil {
		return CodeSelect("Account.Status", nil, optionsValueSet)
	}
	return CodeSelect("Account.Status", &resource.Status, optionsValueSet)
}
func (resource *Account) T_BillingStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.BillingStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.BillingStatus", resource.BillingStatus, optionsValueSet)
}
func (resource *Account) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Type", resource.Type, optionsValueSet)
}
func (resource *Account) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Account.Name", nil)
	}
	return StringInput("Account.Name", resource.Name)
}
func (resource *Account) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Account.Description", nil)
	}
	return StringInput("Account.Description", resource.Description)
}
func (resource *Account) T_Currency(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Account.Currency", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Currency", resource.Currency, optionsValueSet)
}
func (resource *Account) T_CalculatedAt() templ.Component {

	if resource == nil {
		return StringInput("Account.CalculatedAt", nil)
	}
	return StringInput("Account.CalculatedAt", resource.CalculatedAt)
}
func (resource *Account) T_CoverageId(numCoverage int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage {
		return StringInput("Account.Coverage["+strconv.Itoa(numCoverage)+"].Id", nil)
	}
	return StringInput("Account.Coverage["+strconv.Itoa(numCoverage)+"].Id", resource.Coverage[numCoverage].Id)
}
func (resource *Account) T_CoveragePriority(numCoverage int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage {
		return IntInput("Account.Coverage["+strconv.Itoa(numCoverage)+"].Priority", nil)
	}
	return IntInput("Account.Coverage["+strconv.Itoa(numCoverage)+"].Priority", resource.Coverage[numCoverage].Priority)
}
func (resource *Account) T_GuarantorId(numGuarantor int) templ.Component {

	if resource == nil || len(resource.Guarantor) >= numGuarantor {
		return StringInput("Account.Guarantor["+strconv.Itoa(numGuarantor)+"].Id", nil)
	}
	return StringInput("Account.Guarantor["+strconv.Itoa(numGuarantor)+"].Id", resource.Guarantor[numGuarantor].Id)
}
func (resource *Account) T_GuarantorOnHold(numGuarantor int) templ.Component {

	if resource == nil || len(resource.Guarantor) >= numGuarantor {
		return BoolInput("Account.Guarantor["+strconv.Itoa(numGuarantor)+"].OnHold", nil)
	}
	return BoolInput("Account.Guarantor["+strconv.Itoa(numGuarantor)+"].OnHold", resource.Guarantor[numGuarantor].OnHold)
}
func (resource *Account) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *Account) T_DiagnosisSequence(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return IntInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", nil)
	}
	return IntInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", resource.Diagnosis[numDiagnosis].Sequence)
}
func (resource *Account) T_DiagnosisDateOfDiagnosis(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].DateOfDiagnosis", nil)
	}
	return StringInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].DateOfDiagnosis", resource.Diagnosis[numDiagnosis].DateOfDiagnosis)
}
func (resource *Account) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis || len(resource.Diagnosis[numDiagnosis].Type) >= numType {
		return CodeableConceptSelect("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet)
}
func (resource *Account) T_DiagnosisOnAdmission(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return BoolInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", nil)
	}
	return BoolInput("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission)
}
func (resource *Account) T_DiagnosisPackageCode(numDiagnosis int, numPackageCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis || len(resource.Diagnosis[numDiagnosis].PackageCode) >= numPackageCode {
		return CodeableConceptSelect("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].PackageCode["+strconv.Itoa(numPackageCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Diagnosis["+strconv.Itoa(numDiagnosis)+"].PackageCode["+strconv.Itoa(numPackageCode)+"]", &resource.Diagnosis[numDiagnosis].PackageCode[numPackageCode], optionsValueSet)
}
func (resource *Account) T_ProcedureId(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].Id", nil)
	}
	return StringInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].Id", resource.Procedure[numProcedure].Id)
}
func (resource *Account) T_ProcedureSequence(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return IntInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", nil)
	}
	return IntInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", resource.Procedure[numProcedure].Sequence)
}
func (resource *Account) T_ProcedureDateOfService(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].DateOfService", nil)
	}
	return StringInput("Account.Procedure["+strconv.Itoa(numProcedure)+"].DateOfService", resource.Procedure[numProcedure].DateOfService)
}
func (resource *Account) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure || len(resource.Procedure[numProcedure].Type) >= numType {
		return CodeableConceptSelect("Account.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet)
}
func (resource *Account) T_ProcedurePackageCode(numProcedure int, numPackageCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure || len(resource.Procedure[numProcedure].PackageCode) >= numPackageCode {
		return CodeableConceptSelect("Account.Procedure["+strconv.Itoa(numProcedure)+"].PackageCode["+strconv.Itoa(numPackageCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Procedure["+strconv.Itoa(numProcedure)+"].PackageCode["+strconv.Itoa(numPackageCode)+"]", &resource.Procedure[numProcedure].PackageCode[numPackageCode], optionsValueSet)
}
func (resource *Account) T_RelatedAccountId(numRelatedAccount int) templ.Component {

	if resource == nil || len(resource.RelatedAccount) >= numRelatedAccount {
		return StringInput("Account.RelatedAccount["+strconv.Itoa(numRelatedAccount)+"].Id", nil)
	}
	return StringInput("Account.RelatedAccount["+strconv.Itoa(numRelatedAccount)+"].Id", resource.RelatedAccount[numRelatedAccount].Id)
}
func (resource *Account) T_RelatedAccountRelationship(numRelatedAccount int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RelatedAccount) >= numRelatedAccount {
		return CodeableConceptSelect("Account.RelatedAccount["+strconv.Itoa(numRelatedAccount)+"].Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.RelatedAccount["+strconv.Itoa(numRelatedAccount)+"].Relationship", resource.RelatedAccount[numRelatedAccount].Relationship, optionsValueSet)
}
func (resource *Account) T_BalanceId(numBalance int) templ.Component {

	if resource == nil || len(resource.Balance) >= numBalance {
		return StringInput("Account.Balance["+strconv.Itoa(numBalance)+"].Id", nil)
	}
	return StringInput("Account.Balance["+strconv.Itoa(numBalance)+"].Id", resource.Balance[numBalance].Id)
}
func (resource *Account) T_BalanceAggregate(numBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Balance) >= numBalance {
		return CodeableConceptSelect("Account.Balance["+strconv.Itoa(numBalance)+"].Aggregate", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Balance["+strconv.Itoa(numBalance)+"].Aggregate", resource.Balance[numBalance].Aggregate, optionsValueSet)
}
func (resource *Account) T_BalanceTerm(numBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Balance) >= numBalance {
		return CodeableConceptSelect("Account.Balance["+strconv.Itoa(numBalance)+"].Term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Account.Balance["+strconv.Itoa(numBalance)+"].Term", resource.Balance[numBalance].Term, optionsValueSet)
}
func (resource *Account) T_BalanceEstimate(numBalance int) templ.Component {

	if resource == nil || len(resource.Balance) >= numBalance {
		return BoolInput("Account.Balance["+strconv.Itoa(numBalance)+"].Estimate", nil)
	}
	return BoolInput("Account.Balance["+strconv.Itoa(numBalance)+"].Estimate", resource.Balance[numBalance].Estimate)
}
