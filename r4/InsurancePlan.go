//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlan struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	Type              []CodeableConcept       `json:"type,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Alias             []string                `json:"alias,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	OwnedBy           *Reference              `json:"ownedBy,omitempty"`
	AdministeredBy    *Reference              `json:"administeredBy,omitempty"`
	CoverageArea      []Reference             `json:"coverageArea,omitempty"`
	Contact           []InsurancePlanContact  `json:"contact,omitempty"`
	Endpoint          []Reference             `json:"endpoint,omitempty"`
	Network           []Reference             `json:"network,omitempty"`
	Coverage          []InsurancePlanCoverage `json:"coverage,omitempty"`
	Plan              []InsurancePlanPlan     `json:"plan,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanCoverage struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                `json:"type"`
	Network           []Reference                    `json:"network,omitempty"`
	Benefit           []InsurancePlanCoverageBenefit `json:"benefit"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanCoverageBenefit struct {
	Id                *string                             `json:"id,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                     `json:"type"`
	Requirement       *string                             `json:"requirement,omitempty"`
	Limit             []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanCoverageBenefitLimit struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Value             *Quantity        `json:"value,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanPlan struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                    `json:"identifier,omitempty"`
	Type              *CodeableConcept                `json:"type,omitempty"`
	CoverageArea      []Reference                     `json:"coverageArea,omitempty"`
	Network           []Reference                     `json:"network,omitempty"`
	GeneralCost       []InsurancePlanPlanGeneralCost  `json:"generalCost,omitempty"`
	SpecificCost      []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanPlanGeneralCost struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	GroupSize         *int             `json:"groupSize,omitempty"`
	Cost              *Money           `json:"cost,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanPlanSpecificCost struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Category          CodeableConcept                        `json:"category"`
	Benefit           []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanPlanSpecificCostBenefit struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                            `json:"type"`
	Cost              []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/InsurancePlan
type InsurancePlanPlanSpecificCostBenefitCost struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Applicability     *CodeableConcept  `json:"applicability,omitempty"`
	Qualifiers        []CodeableConcept `json:"qualifiers,omitempty"`
	Value             *Quantity         `json:"value,omitempty"`
}

type OtherInsurancePlan InsurancePlan

// on convert struct to json, automatically add resourceType=InsurancePlan
func (r InsurancePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInsurancePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherInsurancePlan: OtherInsurancePlan(r),
		ResourceType:       "InsurancePlan",
	})
}
