package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Coverage) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Coverage.Id", nil)
	}
	return StringInput("Coverage.Id", resource.Id)
}
func (resource *Coverage) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Coverage.ImplicitRules", nil)
	}
	return StringInput("Coverage.ImplicitRules", resource.ImplicitRules)
}
func (resource *Coverage) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Coverage.Language", nil, optionsValueSet)
	}
	return CodeSelect("Coverage.Language", resource.Language, optionsValueSet)
}
func (resource *Coverage) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Coverage.Status", nil, optionsValueSet)
	}
	return CodeSelect("Coverage.Status", &resource.Status, optionsValueSet)
}
func (resource *Coverage) T_Kind() templ.Component {
	optionsValueSet := VSCoverage_kind

	if resource == nil {
		return CodeSelect("Coverage.Kind", nil, optionsValueSet)
	}
	return CodeSelect("Coverage.Kind", &resource.Kind, optionsValueSet)
}
func (resource *Coverage) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Coverage.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.Type", resource.Type, optionsValueSet)
}
func (resource *Coverage) T_Dependent() templ.Component {

	if resource == nil {
		return StringInput("Coverage.Dependent", nil)
	}
	return StringInput("Coverage.Dependent", resource.Dependent)
}
func (resource *Coverage) T_Relationship(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Coverage.Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.Relationship", resource.Relationship, optionsValueSet)
}
func (resource *Coverage) T_Order() templ.Component {

	if resource == nil {
		return IntInput("Coverage.Order", nil)
	}
	return IntInput("Coverage.Order", resource.Order)
}
func (resource *Coverage) T_Network() templ.Component {

	if resource == nil {
		return StringInput("Coverage.Network", nil)
	}
	return StringInput("Coverage.Network", resource.Network)
}
func (resource *Coverage) T_Subrogation() templ.Component {

	if resource == nil {
		return BoolInput("Coverage.Subrogation", nil)
	}
	return BoolInput("Coverage.Subrogation", resource.Subrogation)
}
func (resource *Coverage) T_PaymentById(numPaymentBy int) templ.Component {

	if resource == nil || len(resource.PaymentBy) >= numPaymentBy {
		return StringInput("Coverage.PaymentBy["+strconv.Itoa(numPaymentBy)+"].Id", nil)
	}
	return StringInput("Coverage.PaymentBy["+strconv.Itoa(numPaymentBy)+"].Id", resource.PaymentBy[numPaymentBy].Id)
}
func (resource *Coverage) T_PaymentByResponsibility(numPaymentBy int) templ.Component {

	if resource == nil || len(resource.PaymentBy) >= numPaymentBy {
		return StringInput("Coverage.PaymentBy["+strconv.Itoa(numPaymentBy)+"].Responsibility", nil)
	}
	return StringInput("Coverage.PaymentBy["+strconv.Itoa(numPaymentBy)+"].Responsibility", resource.PaymentBy[numPaymentBy].Responsibility)
}
func (resource *Coverage) T_ClassId(numClass int) templ.Component {

	if resource == nil || len(resource.Class) >= numClass {
		return StringInput("Coverage.Class["+strconv.Itoa(numClass)+"].Id", nil)
	}
	return StringInput("Coverage.Class["+strconv.Itoa(numClass)+"].Id", resource.Class[numClass].Id)
}
func (resource *Coverage) T_ClassType(numClass int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Class) >= numClass {
		return CodeableConceptSelect("Coverage.Class["+strconv.Itoa(numClass)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.Class["+strconv.Itoa(numClass)+"].Type", &resource.Class[numClass].Type, optionsValueSet)
}
func (resource *Coverage) T_ClassName(numClass int) templ.Component {

	if resource == nil || len(resource.Class) >= numClass {
		return StringInput("Coverage.Class["+strconv.Itoa(numClass)+"].Name", nil)
	}
	return StringInput("Coverage.Class["+strconv.Itoa(numClass)+"].Name", resource.Class[numClass].Name)
}
func (resource *Coverage) T_CostToBeneficiaryId(numCostToBeneficiary int) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return StringInput("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Id", nil)
	}
	return StringInput("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Id", resource.CostToBeneficiary[numCostToBeneficiary].Id)
}
func (resource *Coverage) T_CostToBeneficiaryType(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Type", resource.CostToBeneficiary[numCostToBeneficiary].Type, optionsValueSet)
}
func (resource *Coverage) T_CostToBeneficiaryCategory(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Category", resource.CostToBeneficiary[numCostToBeneficiary].Category, optionsValueSet)
}
func (resource *Coverage) T_CostToBeneficiaryNetwork(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Network", resource.CostToBeneficiary[numCostToBeneficiary].Network, optionsValueSet)
}
func (resource *Coverage) T_CostToBeneficiaryUnit(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Unit", resource.CostToBeneficiary[numCostToBeneficiary].Unit, optionsValueSet)
}
func (resource *Coverage) T_CostToBeneficiaryTerm(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Term", resource.CostToBeneficiary[numCostToBeneficiary].Term, optionsValueSet)
}
func (resource *Coverage) T_CostToBeneficiaryExceptionId(numCostToBeneficiary int, numException int) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary || len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) >= numException {
		return StringInput("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Exception["+strconv.Itoa(numException)+"].Id", nil)
	}
	return StringInput("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Exception["+strconv.Itoa(numException)+"].Id", resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Id)
}
func (resource *Coverage) T_CostToBeneficiaryExceptionType(numCostToBeneficiary int, numException int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CostToBeneficiary) >= numCostToBeneficiary || len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) >= numException {
		return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Exception["+strconv.Itoa(numException)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Coverage.CostToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].Exception["+strconv.Itoa(numException)+"].Type", &resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Type, optionsValueSet)
}
