package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	DateOfDiagnosis   *FhirDateTime     `json:"dateOfDiagnosis,omitempty"`
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
	DateOfService     *FhirDateTime     `json:"dateOfService,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}

// json -> struct, first reject if resourceType != Account
func (r *Account) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Account" {
		return errors.New("resourceType not Account")
	}
	return json.Unmarshal(data, (*OtherAccount)(r))
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
func (resource *Account) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAccount_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BillingStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("billingStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("billingStatus", resource.BillingStatus, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Account) T_Subject(frs []FhirResource, numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *Account) T_ServicePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("servicePeriod", nil, htmlAttrs)
	}
	return PeriodInput("servicePeriod", resource.ServicePeriod, htmlAttrs)
}
func (resource *Account) T_Owner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "owner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "owner", resource.Owner, htmlAttrs)
}
func (resource *Account) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Account) T_Currency(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("currency", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("currency", resource.Currency, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_CalculatedAt(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("calculatedAt", nil, htmlAttrs)
	}
	return StringInput("calculatedAt", resource.CalculatedAt, htmlAttrs)
}
func (resource *Account) T_CoverageCoverage(frs []FhirResource, numCoverage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].coverage", &resource.Coverage[numCoverage].Coverage, htmlAttrs)
}
func (resource *Account) T_CoveragePriority(numCoverage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return IntInput("coverage["+strconv.Itoa(numCoverage)+"].priority", nil, htmlAttrs)
	}
	return IntInput("coverage["+strconv.Itoa(numCoverage)+"].priority", resource.Coverage[numCoverage].Priority, htmlAttrs)
}
func (resource *Account) T_GuarantorParty(frs []FhirResource, numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return ReferenceInput(frs, "guarantor["+strconv.Itoa(numGuarantor)+"].party", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "guarantor["+strconv.Itoa(numGuarantor)+"].party", &resource.Guarantor[numGuarantor].Party, htmlAttrs)
}
func (resource *Account) T_GuarantorOnHold(numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return BoolInput("guarantor["+strconv.Itoa(numGuarantor)+"].onHold", nil, htmlAttrs)
	}
	return BoolInput("guarantor["+strconv.Itoa(numGuarantor)+"].onHold", resource.Guarantor[numGuarantor].OnHold, htmlAttrs)
}
func (resource *Account) T_GuarantorPeriod(numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return PeriodInput("guarantor["+strconv.Itoa(numGuarantor)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("guarantor["+strconv.Itoa(numGuarantor)+"].period", resource.Guarantor[numGuarantor].Period, htmlAttrs)
}
func (resource *Account) T_DiagnosisSequence(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *Account) T_DiagnosisCondition(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", nil, htmlAttrs)
	}
	return CodeableReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", &resource.Diagnosis[numDiagnosis].Condition, htmlAttrs)
}
func (resource *Account) T_DiagnosisDateOfDiagnosis(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return FhirDateTimeInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].dateOfDiagnosis", nil, htmlAttrs)
	}
	return FhirDateTimeInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].dateOfDiagnosis", resource.Diagnosis[numDiagnosis].DateOfDiagnosis, htmlAttrs)
}
func (resource *Account) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_DiagnosisOnAdmission(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return BoolInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", nil, htmlAttrs)
	}
	return BoolInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, htmlAttrs)
}
func (resource *Account) T_DiagnosisPackageCode(numDiagnosis int, numPackageCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numPackageCode >= len(resource.Diagnosis[numDiagnosis].PackageCode) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].packageCode["+strconv.Itoa(numPackageCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].packageCode["+strconv.Itoa(numPackageCode)+"]", &resource.Diagnosis[numDiagnosis].PackageCode[numPackageCode], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_ProcedureSequence(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *Account) T_ProcedureCode(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].code", nil, htmlAttrs)
	}
	return CodeableReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].code", &resource.Procedure[numProcedure].Code, htmlAttrs)
}
func (resource *Account) T_ProcedureDateOfService(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].dateOfService", nil, htmlAttrs)
	}
	return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].dateOfService", resource.Procedure[numProcedure].DateOfService, htmlAttrs)
}
func (resource *Account) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_ProcedurePackageCode(numProcedure int, numPackageCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numPackageCode >= len(resource.Procedure[numProcedure].PackageCode) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].packageCode["+strconv.Itoa(numPackageCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].packageCode["+strconv.Itoa(numPackageCode)+"]", &resource.Procedure[numProcedure].PackageCode[numPackageCode], optionsValueSet, htmlAttrs)
}
func (resource *Account) T_ProcedureDevice(frs []FhirResource, numProcedure int, numDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numDevice >= len(resource.Procedure[numProcedure].Device) {
		return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].device["+strconv.Itoa(numDevice)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].device["+strconv.Itoa(numDevice)+"]", &resource.Procedure[numProcedure].Device[numDevice], htmlAttrs)
}
func (resource *Account) T_RelatedAccountRelationship(numRelatedAccount int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedAccount >= len(resource.RelatedAccount) {
		return CodeableConceptSelect("relatedAccount["+strconv.Itoa(numRelatedAccount)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relatedAccount["+strconv.Itoa(numRelatedAccount)+"].relationship", resource.RelatedAccount[numRelatedAccount].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_RelatedAccountAccount(frs []FhirResource, numRelatedAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedAccount >= len(resource.RelatedAccount) {
		return ReferenceInput(frs, "relatedAccount["+strconv.Itoa(numRelatedAccount)+"].account", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relatedAccount["+strconv.Itoa(numRelatedAccount)+"].account", &resource.RelatedAccount[numRelatedAccount].Account, htmlAttrs)
}
func (resource *Account) T_BalanceAggregate(numBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBalance >= len(resource.Balance) {
		return CodeableConceptSelect("balance["+strconv.Itoa(numBalance)+"].aggregate", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("balance["+strconv.Itoa(numBalance)+"].aggregate", resource.Balance[numBalance].Aggregate, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BalanceTerm(numBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBalance >= len(resource.Balance) {
		return CodeableConceptSelect("balance["+strconv.Itoa(numBalance)+"].term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("balance["+strconv.Itoa(numBalance)+"].term", resource.Balance[numBalance].Term, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_BalanceEstimate(numBalance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBalance >= len(resource.Balance) {
		return BoolInput("balance["+strconv.Itoa(numBalance)+"].estimate", nil, htmlAttrs)
	}
	return BoolInput("balance["+strconv.Itoa(numBalance)+"].estimate", resource.Balance[numBalance].Estimate, htmlAttrs)
}
func (resource *Account) T_BalanceAmount(numBalance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBalance >= len(resource.Balance) {
		return MoneyInput("balance["+strconv.Itoa(numBalance)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("balance["+strconv.Itoa(numBalance)+"].amount", &resource.Balance[numBalance].Amount, htmlAttrs)
}
