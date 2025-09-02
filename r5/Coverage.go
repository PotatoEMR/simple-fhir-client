package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ValueQuantity     *Quantity                            `json:"valueQuantity"`
	ValueMoney        *Money                               `json:"valueMoney"`
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

func (resource *Coverage) CoverageLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Coverage) CoverageStatus() templ.Component {
	optionsValueSet := VSFm_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Coverage) CoverageKind() templ.Component {
	optionsValueSet := VSCoverage_kind

	if resource != nil {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet)
}
func (resource *Coverage) CoverageType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *Coverage) CoverageRelationship(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", resource.Relationship, optionsValueSet)
}
func (resource *Coverage) CoverageClassType(numClass int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Class) >= numClass {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Class[numClass].Type, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryType(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CostToBeneficiary[numCostToBeneficiary].Type, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryCategory(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.CostToBeneficiary[numCostToBeneficiary].Category, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryNetwork(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("network", resource.CostToBeneficiary[numCostToBeneficiary].Network, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryUnit(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unit", resource.CostToBeneficiary[numCostToBeneficiary].Unit, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryTerm(numCostToBeneficiary int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary) >= numCostToBeneficiary {
		return CodeableConceptSelect("term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("term", resource.CostToBeneficiary[numCostToBeneficiary].Term, optionsValueSet)
}
func (resource *Coverage) CoverageCostToBeneficiaryExceptionType(numCostToBeneficiary int, numException int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) >= numException {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Type, optionsValueSet)
}
