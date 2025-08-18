//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/Account
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

// http://hl7.org/fhir/r4/StructureDefinition/Account
type AccountCoverage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Coverage          Reference   `json:"coverage"`
	Priority          *int        `json:"priority,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Account
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
