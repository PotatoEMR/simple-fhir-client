package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r InsurancePlan) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "InsurancePlan/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "InsurancePlan"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *InsurancePlan) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("InsurancePlan.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("InsurancePlan.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("InsurancePlan.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InsurancePlan.Name", nil, htmlAttrs)
	}
	return StringInput("InsurancePlan.Name", resource.Name, htmlAttrs)
}
func (resource *InsurancePlan) T_Alias(numAlias int, htmlAttrs string) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("InsurancePlan.Alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("InsurancePlan.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *InsurancePlan) T_ContactPurpose(numContact int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Purpose", resource.Contact[numContact].Purpose, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageType(numCoverage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Type", &resource.Coverage[numCoverage].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitType(numCoverage int, numBenefit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Coverage[numCoverage].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitRequirement(numCoverage int, numBenefit int, htmlAttrs string) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) {
		return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Requirement", nil, htmlAttrs)
	}
	return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Requirement", resource.Coverage[numCoverage].Benefit[numBenefit].Requirement, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitLimitCode(numCoverage int, numBenefit int, numLimit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) || numLimit >= len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Code", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Code, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanType(numPlan int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Type", resource.Plan[numPlan].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostType(numPlan int, numGeneralCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Type", resource.Plan[numPlan].GeneralCost[numGeneralCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostGroupSize(numPlan int, numGeneralCost int, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return IntInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].GroupSize", nil, htmlAttrs)
	}
	return IntInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].GroupSize", resource.Plan[numPlan].GeneralCost[numGeneralCost].GroupSize, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostComment(numPlan int, numGeneralCost int, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Comment", resource.Plan[numPlan].GeneralCost[numGeneralCost].Comment, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostCategory(numPlan int, numSpecificCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Category", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Category, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitType(numPlan int, numSpecificCost int, numBenefit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostType(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostApplicability(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Applicability", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Applicability", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Applicability, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostQualifiers(numPlan int, numSpecificCost int, numBenefit int, numCost int, numQualifiers int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) || numQualifiers >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers) {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Qualifiers["+strconv.Itoa(numQualifiers)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Qualifiers["+strconv.Itoa(numQualifiers)+"]", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers[numQualifiers], optionsValueSet, htmlAttrs)
}
