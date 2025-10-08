package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Coverage
type Coverage struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Status            string                      `json:"status"`
	Kind              string                      `json:"kind"`
	PaymentBy         []CoveragePaymentBy         `json:"paymentBy,omitempty"`
	Type              *CodeableConcept            `json:"type,omitempty"`
	PolicyHolder      *Reference                  `json:"policyHolder,omitempty"`
	Subscriber        *Reference                  `json:"subscriber,omitempty"`
	SubscriberId      []Identifier                `json:"subscriberId,omitempty"`
	Beneficiary       Reference                   `json:"beneficiary"`
	Dependent         *string                     `json:"dependent,omitempty"`
	Relationship      *CodeableConcept            `json:"relationship,omitempty"`
	Period            *Period                     `json:"period,omitempty"`
	Insurer           *Reference                  `json:"insurer,omitempty"`
	Class             []CoverageClass             `json:"class,omitempty"`
	Order             *int                        `json:"order,omitempty"`
	Network           *string                     `json:"network,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`
	Subrogation       *bool                       `json:"subrogation,omitempty"`
	Contract          []Reference                 `json:"contract,omitempty"`
	InsurancePlan     *Reference                  `json:"insurancePlan,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Coverage
type CoveragePaymentBy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Party             Reference   `json:"party"`
	Responsibility    *string     `json:"responsibility,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Coverage
type CoverageClass struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Value             Identifier      `json:"value"`
	Name              *string         `json:"name,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Coverage
type CoverageCostToBeneficiary struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	Category          *CodeableConcept                     `json:"category,omitempty"`
	Network           *CodeableConcept                     `json:"network,omitempty"`
	Unit              *CodeableConcept                     `json:"unit,omitempty"`
	Term              *CodeableConcept                     `json:"term,omitempty"`
	ValueQuantity     *Quantity                            `json:"valueQuantity,omitempty"`
	ValueMoney        *Money                               `json:"valueMoney,omitempty"`
	Exception         []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Coverage
type CoverageCostToBeneficiaryException struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Period            *Period         `json:"period,omitempty"`
}

type OtherCoverage Coverage

// on convert struct to json, automatically add resourceType=Coverage
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
}
func (r Coverage) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Coverage/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Coverage"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Coverage) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCoverage_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_PolicyHolder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "policyHolder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "policyHolder", resource.PolicyHolder, htmlAttrs)
}
func (resource *Coverage) T_Subscriber(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subscriber", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subscriber", resource.Subscriber, htmlAttrs)
}
func (resource *Coverage) T_SubscriberId(numSubscriberId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubscriberId >= len(resource.SubscriberId) {
		return IdentifierInput("subscriberId["+strconv.Itoa(numSubscriberId)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("subscriberId["+strconv.Itoa(numSubscriberId)+"]", &resource.SubscriberId[numSubscriberId], htmlAttrs)
}
func (resource *Coverage) T_Beneficiary(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "beneficiary", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "beneficiary", &resource.Beneficiary, htmlAttrs)
}
func (resource *Coverage) T_Dependent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("dependent", nil, htmlAttrs)
	}
	return StringInput("dependent", resource.Dependent, htmlAttrs)
}
func (resource *Coverage) T_Relationship(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship", resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *Coverage) T_Insurer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "insurer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurer", resource.Insurer, htmlAttrs)
}
func (resource *Coverage) T_Order(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("order", nil, htmlAttrs)
	}
	return IntInput("order", resource.Order, htmlAttrs)
}
func (resource *Coverage) T_Network(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("network", nil, htmlAttrs)
	}
	return StringInput("network", resource.Network, htmlAttrs)
}
func (resource *Coverage) T_Subrogation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("subrogation", nil, htmlAttrs)
	}
	return BoolInput("subrogation", resource.Subrogation, htmlAttrs)
}
func (resource *Coverage) T_Contract(frs []FhirResource, numContract int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContract >= len(resource.Contract) {
		return ReferenceInput(frs, "contract["+strconv.Itoa(numContract)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contract["+strconv.Itoa(numContract)+"]", &resource.Contract[numContract], htmlAttrs)
}
func (resource *Coverage) T_InsurancePlan(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "insurancePlan", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurancePlan", resource.InsurancePlan, htmlAttrs)
}
func (resource *Coverage) T_PaymentByParty(frs []FhirResource, numPaymentBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPaymentBy >= len(resource.PaymentBy) {
		return ReferenceInput(frs, "paymentBy["+strconv.Itoa(numPaymentBy)+"].party", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "paymentBy["+strconv.Itoa(numPaymentBy)+"].party", &resource.PaymentBy[numPaymentBy].Party, htmlAttrs)
}
func (resource *Coverage) T_PaymentByResponsibility(numPaymentBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPaymentBy >= len(resource.PaymentBy) {
		return StringInput("paymentBy["+strconv.Itoa(numPaymentBy)+"].responsibility", nil, htmlAttrs)
	}
	return StringInput("paymentBy["+strconv.Itoa(numPaymentBy)+"].responsibility", resource.PaymentBy[numPaymentBy].Responsibility, htmlAttrs)
}
func (resource *Coverage) T_ClassType(numClass int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"].type", &resource.Class[numClass].Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_ClassValue(numClass int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return IdentifierInput("class["+strconv.Itoa(numClass)+"].value", nil, htmlAttrs)
	}
	return IdentifierInput("class["+strconv.Itoa(numClass)+"].value", &resource.Class[numClass].Value, htmlAttrs)
}
func (resource *Coverage) T_ClassName(numClass int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return StringInput("class["+strconv.Itoa(numClass)+"].name", nil, htmlAttrs)
	}
	return StringInput("class["+strconv.Itoa(numClass)+"].name", resource.Class[numClass].Name, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryType(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].type", resource.CostToBeneficiary[numCostToBeneficiary].Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryCategory(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].category", resource.CostToBeneficiary[numCostToBeneficiary].Category, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryNetwork(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].network", resource.CostToBeneficiary[numCostToBeneficiary].Network, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryUnit(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].unit", resource.CostToBeneficiary[numCostToBeneficiary].Unit, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryTerm(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].term", resource.CostToBeneficiary[numCostToBeneficiary].Term, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryValueQuantity(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return QuantityInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].valueQuantity", resource.CostToBeneficiary[numCostToBeneficiary].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryValueMoney(numCostToBeneficiary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return MoneyInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].valueMoney", nil, htmlAttrs)
	}
	return MoneyInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].valueMoney", resource.CostToBeneficiary[numCostToBeneficiary].ValueMoney, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryExceptionType(numCostToBeneficiary int, numException int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) || numException >= len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].type", &resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryExceptionPeriod(numCostToBeneficiary int, numException int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) || numException >= len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) {
		return PeriodInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].period", resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Period, htmlAttrs)
}
