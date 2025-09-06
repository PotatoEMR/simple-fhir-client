package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type Account struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Status            string             `json:"status"`
	Type              *CodeableConcept   `json:"type,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Subject           []Reference        `json:"subject,omitempty"`
	ServicePeriod     *Period            `json:"servicePeriod,omitempty"`
	Coverage          []AccountCoverage  `json:"coverage,omitempty"`
	Owner             *Reference         `json:"owner,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Guarantor         []AccountGuarantor `json:"guarantor,omitempty"`
	PartOf            *Reference         `json:"partOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type AccountCoverage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Coverage          Reference   `json:"coverage"`
	Priority          *int        `json:"priority,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type AccountGuarantor struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Party             Reference   `json:"party"`
	OnHold            *bool       `json:"onHold,omitempty"`
	Period            *Period     `json:"period,omitempty"`
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
