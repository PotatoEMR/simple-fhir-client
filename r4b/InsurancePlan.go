package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
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

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanCoverage struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                `json:"type"`
	Network           []Reference                    `json:"network,omitempty"`
	Benefit           []InsurancePlanCoverageBenefit `json:"benefit"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanCoverageBenefit struct {
	Id                *string                             `json:"id,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                     `json:"type"`
	Requirement       *string                             `json:"requirement,omitempty"`
	Limit             []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanCoverageBenefitLimit struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Value             *Quantity        `json:"value,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
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

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanPlanGeneralCost struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	GroupSize         *int             `json:"groupSize,omitempty"`
	Cost              *Money           `json:"cost,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanPlanSpecificCost struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Category          CodeableConcept                        `json:"category"`
	Benefit           []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
type InsurancePlanPlanSpecificCostBenefit struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                            `json:"type"`
	Cost              []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/InsurancePlan
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

// struct -> json, automatically add resourceType=Patient
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
func (r InsurancePlan) ResourceType() string {
	return "InsurancePlan"
}

func (resource *InsurancePlan) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *InsurancePlan) T_Alias(numAlias int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *InsurancePlan) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *InsurancePlan) T_OwnedBy(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "ownedBy", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "ownedBy", resource.OwnedBy, htmlAttrs)
}
func (resource *InsurancePlan) T_AdministeredBy(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "administeredBy", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "administeredBy", resource.AdministeredBy, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageArea(frs []FhirResource, numCoverageArea int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverageArea >= len(resource.CoverageArea) {
		return ReferenceInput(frs, "coverageArea["+strconv.Itoa(numCoverageArea)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "coverageArea["+strconv.Itoa(numCoverageArea)+"]", &resource.CoverageArea[numCoverageArea], htmlAttrs)
}
func (resource *InsurancePlan) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
func (resource *InsurancePlan) T_Network(frs []FhirResource, numNetwork int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNetwork >= len(resource.Network) {
		return ReferenceInput(frs, "network["+strconv.Itoa(numNetwork)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "network["+strconv.Itoa(numNetwork)+"]", &resource.Network[numNetwork], htmlAttrs)
}
func (resource *InsurancePlan) T_ContactPurpose(numContact int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].purpose", resource.Contact[numContact].Purpose, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_ContactName(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", nil, htmlAttrs)
	}
	return HumanNameInput("contact["+strconv.Itoa(numContact)+"].name", resource.Contact[numContact].Name, htmlAttrs)
}
func (resource *InsurancePlan) T_ContactTelecom(numContact int, numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) || numTelecom >= len(resource.Contact[numContact].Telecom) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"].telecom["+strconv.Itoa(numTelecom)+"]", &resource.Contact[numContact].Telecom[numTelecom], htmlAttrs)
}
func (resource *InsurancePlan) T_ContactAddress(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return AddressInput("contact["+strconv.Itoa(numContact)+"].address", nil, htmlAttrs)
	}
	return AddressInput("contact["+strconv.Itoa(numContact)+"].address", resource.Contact[numContact].Address, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageType(numCoverage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].type", &resource.Coverage[numCoverage].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageNetwork(frs []FhirResource, numCoverage int, numNetwork int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numNetwork >= len(resource.Coverage[numCoverage].Network) {
		return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].network["+strconv.Itoa(numNetwork)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].network["+strconv.Itoa(numNetwork)+"]", &resource.Coverage[numCoverage].Network[numNetwork], htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitType(numCoverage int, numBenefit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) {
		return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].type", &resource.Coverage[numCoverage].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitRequirement(numCoverage int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) {
		return StringInput("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].requirement", nil, htmlAttrs)
	}
	return StringInput("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].requirement", resource.Coverage[numCoverage].Benefit[numBenefit].Requirement, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitLimitValue(numCoverage int, numBenefit int, numLimit int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) || numLimit >= len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) {
		return QuantityInput("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].limit["+strconv.Itoa(numLimit)+"].value", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].limit["+strconv.Itoa(numLimit)+"].value", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Value, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_CoverageBenefitLimitCode(numCoverage int, numBenefit int, numLimit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) || numBenefit >= len(resource.Coverage[numCoverage].Benefit) || numLimit >= len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) {
		return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].limit["+strconv.Itoa(numLimit)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("coverage["+strconv.Itoa(numCoverage)+"].benefit["+strconv.Itoa(numBenefit)+"].limit["+strconv.Itoa(numLimit)+"].code", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Code, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanType(numPlan int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].type", resource.Plan[numPlan].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanCoverageArea(frs []FhirResource, numPlan int, numCoverageArea int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numCoverageArea >= len(resource.Plan[numPlan].CoverageArea) {
		return ReferenceInput(frs, "plan["+strconv.Itoa(numPlan)+"].coverageArea["+strconv.Itoa(numCoverageArea)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "plan["+strconv.Itoa(numPlan)+"].coverageArea["+strconv.Itoa(numCoverageArea)+"]", &resource.Plan[numPlan].CoverageArea[numCoverageArea], htmlAttrs)
}
func (resource *InsurancePlan) T_PlanNetwork(frs []FhirResource, numPlan int, numNetwork int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numNetwork >= len(resource.Plan[numPlan].Network) {
		return ReferenceInput(frs, "plan["+strconv.Itoa(numPlan)+"].network["+strconv.Itoa(numNetwork)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "plan["+strconv.Itoa(numPlan)+"].network["+strconv.Itoa(numNetwork)+"]", &resource.Plan[numPlan].Network[numNetwork], htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostType(numPlan int, numGeneralCost int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].type", resource.Plan[numPlan].GeneralCost[numGeneralCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostGroupSize(numPlan int, numGeneralCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return IntInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].groupSize", nil, htmlAttrs)
	}
	return IntInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].groupSize", resource.Plan[numPlan].GeneralCost[numGeneralCost].GroupSize, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostCost(numPlan int, numGeneralCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return MoneyInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].cost", nil, htmlAttrs)
	}
	return MoneyInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].cost", resource.Plan[numPlan].GeneralCost[numGeneralCost].Cost, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanGeneralCostComment(numPlan int, numGeneralCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numGeneralCost >= len(resource.Plan[numPlan].GeneralCost) {
		return StringInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].comment", nil, htmlAttrs)
	}
	return StringInput("plan["+strconv.Itoa(numPlan)+"].generalCost["+strconv.Itoa(numGeneralCost)+"].comment", resource.Plan[numPlan].GeneralCost[numGeneralCost].Comment, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostCategory(numPlan int, numSpecificCost int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].category", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Category, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitType(numPlan int, numSpecificCost int, numBenefit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostType(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostApplicability(numPlan int, numSpecificCost int, numBenefit int, numCost int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSInsuranceplan_applicability

	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].applicability", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].applicability", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Applicability, optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostQualifiers(numPlan int, numSpecificCost int, numBenefit int, numCost int, numQualifiers int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) || numQualifiers >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].qualifiers["+strconv.Itoa(numQualifiers)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].qualifiers["+strconv.Itoa(numQualifiers)+"]", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers[numQualifiers], optionsValueSet, htmlAttrs)
}
func (resource *InsurancePlan) T_PlanSpecificCostBenefitCostValue(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) || numSpecificCost >= len(resource.Plan[numPlan].SpecificCost) || numBenefit >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) || numCost >= len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) {
		return QuantityInput("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].value", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("plan["+strconv.Itoa(numPlan)+"].specificCost["+strconv.Itoa(numSpecificCost)+"].benefit["+strconv.Itoa(numBenefit)+"].cost["+strconv.Itoa(numCost)+"].value", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Value, optionsValueSet, htmlAttrs)
}
