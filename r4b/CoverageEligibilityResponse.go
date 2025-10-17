package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponse struct {
	Id                *string                                `json:"id,omitempty"`
	Meta              *Meta                                  `json:"meta,omitempty"`
	ImplicitRules     *string                                `json:"implicitRules,omitempty"`
	Language          *string                                `json:"language,omitempty"`
	Text              *Narrative                             `json:"text,omitempty"`
	Contained         []Resource                             `json:"contained,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `json:"identifier,omitempty"`
	Status            string                                 `json:"status"`
	Purpose           []string                               `json:"purpose"`
	Patient           Reference                              `json:"patient"`
	ServicedDate      *FhirDate                              `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                `json:"servicedPeriod,omitempty"`
	Created           FhirDateTime                           `json:"created"`
	Requestor         *Reference                             `json:"requestor,omitempty"`
	Request           Reference                              `json:"request"`
	Outcome           string                                 `json:"outcome"`
	Disposition       *string                                `json:"disposition,omitempty"`
	Insurer           Reference                              `json:"insurer"`
	Insurance         []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	PreAuthRef        *string                                `json:"preAuthRef,omitempty"`
	Form              *CodeableConcept                       `json:"form,omitempty"`
	Error             []CoverageEligibilityResponseError     `json:"error,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsurance struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Coverage          Reference                                  `json:"coverage"`
	Inforce           *bool                                      `json:"inforce,omitempty"`
	BenefitPeriod     *Period                                    `json:"benefitPeriod,omitempty"`
	Item              []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsuranceItem struct {
	Id                      *string                                           `json:"id,omitempty"`
	Extension               []Extension                                       `json:"extension,omitempty"`
	ModifierExtension       []Extension                                       `json:"modifierExtension,omitempty"`
	Category                *CodeableConcept                                  `json:"category,omitempty"`
	ProductOrService        *CodeableConcept                                  `json:"productOrService,omitempty"`
	Modifier                []CodeableConcept                                 `json:"modifier,omitempty"`
	Provider                *Reference                                        `json:"provider,omitempty"`
	Excluded                *bool                                             `json:"excluded,omitempty"`
	Name                    *string                                           `json:"name,omitempty"`
	Description             *string                                           `json:"description,omitempty"`
	Network                 *CodeableConcept                                  `json:"network,omitempty"`
	Unit                    *CodeableConcept                                  `json:"unit,omitempty"`
	Term                    *CodeableConcept                                  `json:"term,omitempty"`
	Benefit                 []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`
	AuthorizationRequired   *bool                                             `json:"authorizationRequired,omitempty"`
	AuthorizationSupporting []CodeableConcept                                 `json:"authorizationSupporting,omitempty"`
	AuthorizationUrl        *string                                           `json:"authorizationUrl,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	Id                 *string         `json:"id,omitempty"`
	Extension          []Extension     `json:"extension,omitempty"`
	ModifierExtension  []Extension     `json:"modifierExtension,omitempty"`
	Type               CodeableConcept `json:"type"`
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `json:"allowedString,omitempty"`
	AllowedMoney       *Money          `json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt,omitempty"`
	UsedString         *string         `json:"usedString,omitempty"`
	UsedMoney          *Money          `json:"usedMoney,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
}

type OtherCoverageEligibilityResponse CoverageEligibilityResponse

// struct -> json, automatically add resourceType=Patient
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityResponse: OtherCoverageEligibilityResponse(r),
		ResourceType:                     "CoverageEligibilityResponse",
	})
}

func (r CoverageEligibilityResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CoverageEligibilityResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CoverageEligibilityResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r CoverageEligibilityResponse) ResourceType() string {
	return "CoverageEligibilityResponse"
}

func (resource *CoverageEligibilityResponse) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Purpose(numPurpose int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEligibilityresponse_purpose

	if resource == nil || numPurpose >= len(resource.Purpose) {
		return CodeSelect("purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ServicedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("servicedDate", resource.ServicedDate, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ServicedPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("servicedPeriod", resource.ServicedPeriod, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Requestor(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requestor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requestor", resource.Requestor, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Request(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request", &resource.Request, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Insurer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "insurer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurer", &resource.Insurer, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_PreAuthRef(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("preAuthRef", nil, htmlAttrs)
	}
	return StringInput("preAuthRef", resource.PreAuthRef, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Form(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("form", resource.Form, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceCoverage(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].coverage", &resource.Insurance[numInsurance].Coverage, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceInforce(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].inforce", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].inforce", resource.Insurance[numInsurance].Inforce, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceBenefitPeriod(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return PeriodInput("insurance["+strconv.Itoa(numInsurance)+"].benefitPeriod", nil, htmlAttrs)
	}
	return PeriodInput("insurance["+strconv.Itoa(numInsurance)+"].benefitPeriod", resource.Insurance[numInsurance].BenefitPeriod, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemCategory(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].category", resource.Insurance[numInsurance].Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemProductOrService(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].productOrService", resource.Insurance[numInsurance].Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemModifier(numInsurance int, numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numModifier >= len(resource.Insurance[numInsurance].Item[numItem].Modifier) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Insurance[numInsurance].Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemProvider(frs []FhirResource, numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].provider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].provider", resource.Insurance[numInsurance].Item[numItem].Provider, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemExcluded(numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].excluded", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].excluded", resource.Insurance[numInsurance].Item[numItem].Excluded, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemName(numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].name", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].name", resource.Insurance[numInsurance].Item[numItem].Name, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemDescription(numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].description", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].description", resource.Insurance[numInsurance].Item[numItem].Description, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemNetwork(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].network", resource.Insurance[numInsurance].Item[numItem].Network, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemUnit(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].unit", resource.Insurance[numInsurance].Item[numItem].Unit, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemTerm(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].term", resource.Insurance[numInsurance].Item[numItem].Term, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationRequired(numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationRequired", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationRequired", resource.Insurance[numInsurance].Item[numItem].AuthorizationRequired, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationSupporting(numInsurance int, numItem int, numAuthorizationSupporting int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numAuthorizationSupporting >= len(resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", &resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting[numAuthorizationSupporting], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationUrl(numInsurance int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationUrl", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].authorizationUrl", resource.Insurance[numInsurance].Item[numItem].AuthorizationUrl, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitType(numInsurance int, numItem int, numBenefit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].type", &resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitAllowedUnsignedInt(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return IntInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedUnsignedInt", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].AllowedUnsignedInt, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitAllowedString(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedString", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedString", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].AllowedString, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitAllowedMoney(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return MoneyInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedMoney", nil, htmlAttrs)
	}
	return MoneyInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].allowedMoney", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].AllowedMoney, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitUsedUnsignedInt(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return IntInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedUnsignedInt", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].UsedUnsignedInt, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitUsedString(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedString", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedString", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].UsedString, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitUsedMoney(numInsurance int, numItem int, numBenefit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return MoneyInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedMoney", nil, htmlAttrs)
	}
	return MoneyInput("insurance["+strconv.Itoa(numInsurance)+"].item["+strconv.Itoa(numItem)+"].benefit["+strconv.Itoa(numBenefit)+"].usedMoney", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].UsedMoney, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ErrorCode(numError int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", &resource.Error[numError].Code, optionsValueSet, htmlAttrs)
}
