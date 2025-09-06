package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *InsurancePlan) T_Id() templ.Component {

	if resource == nil {
		return StringInput("InsurancePlan.Id", nil)
	}
	return StringInput("InsurancePlan.Id", resource.Id)
}
func (resource *InsurancePlan) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("InsurancePlan.ImplicitRules", nil)
	}
	return StringInput("InsurancePlan.ImplicitRules", resource.ImplicitRules)
}
func (resource *InsurancePlan) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("InsurancePlan.Language", nil, optionsValueSet)
	}
	return CodeSelect("InsurancePlan.Language", resource.Language, optionsValueSet)
}
func (resource *InsurancePlan) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("InsurancePlan.Status", nil, optionsValueSet)
	}
	return CodeSelect("InsurancePlan.Status", resource.Status, optionsValueSet)
}
func (resource *InsurancePlan) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("InsurancePlan.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *InsurancePlan) T_Name() templ.Component {

	if resource == nil {
		return StringInput("InsurancePlan.Name", nil)
	}
	return StringInput("InsurancePlan.Name", resource.Name)
}
func (resource *InsurancePlan) T_Alias(numAlias int) templ.Component {

	if resource == nil || len(resource.Alias) >= numAlias {
		return StringInput("InsurancePlan.Alias["+strconv.Itoa(numAlias)+"]", nil)
	}
	return StringInput("InsurancePlan.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias])
}
func (resource *InsurancePlan) T_ContactId(numContact int) templ.Component {

	if resource == nil || len(resource.Contact) >= numContact {
		return StringInput("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Id", resource.Contact[numContact].Id)
}
func (resource *InsurancePlan) T_ContactPurpose(numContact int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Contact) >= numContact {
		return CodeableConceptSelect("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Purpose", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Contact["+strconv.Itoa(numContact)+"].Purpose", resource.Contact[numContact].Purpose, optionsValueSet)
}
func (resource *InsurancePlan) T_CoverageId(numCoverage int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage {
		return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Id", resource.Coverage[numCoverage].Id)
}
func (resource *InsurancePlan) T_CoverageType(numCoverage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Type", &resource.Coverage[numCoverage].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_CoverageBenefitId(numCoverage int, numBenefit int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage || len(resource.Coverage[numCoverage].Benefit) >= numBenefit {
		return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", resource.Coverage[numCoverage].Benefit[numBenefit].Id)
}
func (resource *InsurancePlan) T_CoverageBenefitType(numCoverage int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage || len(resource.Coverage[numCoverage].Benefit) >= numBenefit {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Coverage[numCoverage].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_CoverageBenefitRequirement(numCoverage int, numBenefit int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage || len(resource.Coverage[numCoverage].Benefit) >= numBenefit {
		return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Requirement", nil)
	}
	return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Requirement", resource.Coverage[numCoverage].Benefit[numBenefit].Requirement)
}
func (resource *InsurancePlan) T_CoverageBenefitLimitId(numCoverage int, numBenefit int, numLimit int) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage || len(resource.Coverage[numCoverage].Benefit) >= numBenefit || len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) >= numLimit {
		return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Id", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Id)
}
func (resource *InsurancePlan) T_CoverageBenefitLimitCode(numCoverage int, numBenefit int, numLimit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Coverage) >= numCoverage || len(resource.Coverage[numCoverage].Benefit) >= numBenefit || len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) >= numLimit {
		return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Coverage["+strconv.Itoa(numCoverage)+"].Benefit["+strconv.Itoa(numBenefit)+"].Limit["+strconv.Itoa(numLimit)+"].Code", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Code, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanId(numPlan int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Id", resource.Plan[numPlan].Id)
}
func (resource *InsurancePlan) T_PlanType(numPlan int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].Type", resource.Plan[numPlan].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanGeneralCostId(numPlan int, numGeneralCost int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].GeneralCost) >= numGeneralCost {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Id", resource.Plan[numPlan].GeneralCost[numGeneralCost].Id)
}
func (resource *InsurancePlan) T_PlanGeneralCostType(numPlan int, numGeneralCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].GeneralCost) >= numGeneralCost {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Type", resource.Plan[numPlan].GeneralCost[numGeneralCost].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanGeneralCostGroupSize(numPlan int, numGeneralCost int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].GeneralCost) >= numGeneralCost {
		return IntInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].GroupSize", nil)
	}
	return IntInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].GroupSize", resource.Plan[numPlan].GeneralCost[numGeneralCost].GroupSize)
}
func (resource *InsurancePlan) T_PlanGeneralCostComment(numPlan int, numGeneralCost int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].GeneralCost) >= numGeneralCost {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Comment", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].GeneralCost["+strconv.Itoa(numGeneralCost)+"].Comment", resource.Plan[numPlan].GeneralCost[numGeneralCost].Comment)
}
func (resource *InsurancePlan) T_PlanSpecificCostId(numPlan int, numSpecificCost int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Id", resource.Plan[numPlan].SpecificCost[numSpecificCost].Id)
}
func (resource *InsurancePlan) T_PlanSpecificCostCategory(numPlan int, numSpecificCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Category", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Category, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitId(numPlan int, numSpecificCost int, numBenefit int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Id)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitType(numPlan int, numSpecificCost int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostId(numPlan int, numSpecificCost int, numBenefit int, numCost int) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Id", nil)
	}
	return StringInput("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Id", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Id)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostType(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Type, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostApplicability(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Applicability", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Applicability", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Applicability, optionsValueSet)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostQualifiers(numPlan int, numSpecificCost int, numBenefit int, numCost int, numQualifiers int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Plan) >= numPlan || len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost || len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers) >= numQualifiers {
		return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Qualifiers["+strconv.Itoa(numQualifiers)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InsurancePlan.Plan["+strconv.Itoa(numPlan)+"].SpecificCost["+strconv.Itoa(numSpecificCost)+"].Benefit["+strconv.Itoa(numBenefit)+"].Cost["+strconv.Itoa(numCost)+"].Qualifiers["+strconv.Itoa(numQualifiers)+"]", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers[numQualifiers], optionsValueSet)
}
