package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Account) AccountLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Account) AccountStatus() templ.Component {
	optionsValueSet := VSAccount_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Account) AccountBillingStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("billingStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("billingStatus", resource.BillingStatus, optionsValueSet)
}
func (resource *Account) AccountType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *Account) AccountCurrency(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("currency", nil, optionsValueSet)
	}
	return CodeableConceptSelect("currency", resource.Currency, optionsValueSet)
}
func (resource *Account) AccountDiagnosisType(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Diagnosis[numDiagnosis].Type[0], optionsValueSet)
}
func (resource *Account) AccountDiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("packageCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("packageCode", &resource.Diagnosis[numDiagnosis].PackageCode[0], optionsValueSet)
}
func (resource *Account) AccountProcedureType(numProcedure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Procedure) >= numProcedure {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Procedure[numProcedure].Type[0], optionsValueSet)
}
func (resource *Account) AccountProcedurePackageCode(numProcedure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Procedure) >= numProcedure {
		return CodeableConceptSelect("packageCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("packageCode", &resource.Procedure[numProcedure].PackageCode[0], optionsValueSet)
}
func (resource *Account) AccountRelatedAccountRelationship(numRelatedAccount int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.RelatedAccount) >= numRelatedAccount {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", resource.RelatedAccount[numRelatedAccount].Relationship, optionsValueSet)
}
func (resource *Account) AccountBalanceAggregate(numBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Balance) >= numBalance {
		return CodeableConceptSelect("aggregate", nil, optionsValueSet)
	}
	return CodeableConceptSelect("aggregate", resource.Balance[numBalance].Aggregate, optionsValueSet)
}
func (resource *Account) AccountBalanceTerm(numBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Balance) >= numBalance {
		return CodeableConceptSelect("term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("term", resource.Balance[numBalance].Term, optionsValueSet)
}
