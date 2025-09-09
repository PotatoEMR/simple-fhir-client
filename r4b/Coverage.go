package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Coverage
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
	Type              *CodeableConcept            `json:"type,omitempty"`
	PolicyHolder      *Reference                  `json:"policyHolder,omitempty"`
	Subscriber        *Reference                  `json:"subscriber,omitempty"`
	SubscriberId      *string                     `json:"subscriberId,omitempty"`
	Beneficiary       Reference                   `json:"beneficiary"`
	Dependent         *string                     `json:"dependent,omitempty"`
	Relationship      *CodeableConcept            `json:"relationship,omitempty"`
	Period            *Period                     `json:"period,omitempty"`
	Payor             []Reference                 `json:"payor"`
	Class             []CoverageClass             `json:"class,omitempty"`
	Order             *int                        `json:"order,omitempty"`
	Network           *string                     `json:"network,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`
	Subrogation       *bool                       `json:"subrogation,omitempty"`
	Contract          []Reference                 `json:"contract,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Coverage
type CoverageClass struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Value             string          `json:"value"`
	Name              *string         `json:"name,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Coverage
type CoverageCostToBeneficiary struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	ValueQuantity     Quantity                             `json:"valueQuantity"`
	ValueMoney        Money                                `json:"valueMoney"`
	Exception         []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Coverage
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
func (resource *Coverage) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_SubscriberId(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("subscriberId", nil, htmlAttrs)
	}
	return StringInput("subscriberId", resource.SubscriberId, htmlAttrs)
}
func (resource *Coverage) T_Dependent(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("dependent", nil, htmlAttrs)
	}
	return StringInput("dependent", resource.Dependent, htmlAttrs)
}
func (resource *Coverage) T_Relationship(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship", resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_Order(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("order", nil, htmlAttrs)
	}
	return IntInput("order", resource.Order, htmlAttrs)
}
func (resource *Coverage) T_Network(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("network", nil, htmlAttrs)
	}
	return StringInput("network", resource.Network, htmlAttrs)
}
func (resource *Coverage) T_Subrogation(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("subrogation", nil, htmlAttrs)
	}
	return BoolInput("subrogation", resource.Subrogation, htmlAttrs)
}
func (resource *Coverage) T_ClassType(numClass int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"].type", &resource.Class[numClass].Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_ClassValue(numClass int, htmlAttrs string) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return StringInput("class["+strconv.Itoa(numClass)+"].value", nil, htmlAttrs)
	}
	return StringInput("class["+strconv.Itoa(numClass)+"].value", &resource.Class[numClass].Value, htmlAttrs)
}
func (resource *Coverage) T_ClassName(numClass int, htmlAttrs string) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return StringInput("class["+strconv.Itoa(numClass)+"].name", nil, htmlAttrs)
	}
	return StringInput("class["+strconv.Itoa(numClass)+"].name", resource.Class[numClass].Name, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryType(numCostToBeneficiary int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].type", resource.CostToBeneficiary[numCostToBeneficiary].Type, optionsValueSet, htmlAttrs)
}
func (resource *Coverage) T_CostToBeneficiaryExceptionType(numCostToBeneficiary int, numException int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCostToBeneficiary >= len(resource.CostToBeneficiary) || numException >= len(resource.CostToBeneficiary[numCostToBeneficiary].Exception) {
		return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("costToBeneficiary["+strconv.Itoa(numCostToBeneficiary)+"].exception["+strconv.Itoa(numException)+"].type", &resource.CostToBeneficiary[numCostToBeneficiary].Exception[numException].Type, optionsValueSet, htmlAttrs)
}
