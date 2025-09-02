package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *InsurancePlan) InsurancePlanLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanContactPurpose(numContact int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Contact) >= numContact {
		return CodeableConceptSelect("purpose", nil, optionsValueSet)
	}
	return CodeableConceptSelect("purpose", resource.Contact[numContact].Purpose, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanCoverageType(numCoverage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Coverage) >= numCoverage {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Coverage[numCoverage].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanCoverageBenefitType(numCoverage int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Coverage[numCoverage].Benefit) >= numBenefit {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Coverage[numCoverage].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanCoverageBenefitLimitCode(numCoverage int, numBenefit int, numLimit int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Coverage[numCoverage].Benefit[numBenefit].Limit) >= numLimit {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Coverage[numCoverage].Benefit[numBenefit].Limit[numLimit].Code, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanType(numPlan int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan) >= numPlan {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Plan[numPlan].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanGeneralCostType(numPlan int, numGeneralCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].GeneralCost) >= numGeneralCost {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Plan[numPlan].GeneralCost[numGeneralCost].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanSpecificCostCategory(numPlan int, numSpecificCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].SpecificCost) >= numSpecificCost {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Category, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanSpecificCostBenefitType(numPlan int, numSpecificCost int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit) >= numBenefit {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanSpecificCostBenefitCostType(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Type, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanSpecificCostBenefitCostApplicability(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return CodeableConceptSelect("applicability", nil, optionsValueSet)
	}
	return CodeableConceptSelect("applicability", resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Applicability, optionsValueSet)
}
func (resource *InsurancePlan) InsurancePlanPlanSpecificCostBenefitCostQualifiers(numPlan int, numSpecificCost int, numBenefit int, numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost) >= numCost {
		return CodeableConceptSelect("qualifiers", nil, optionsValueSet)
	}
	return CodeableConceptSelect("qualifiers", &resource.Plan[numPlan].SpecificCost[numSpecificCost].Benefit[numBenefit].Cost[numCost].Qualifiers[0], optionsValueSet)
}
