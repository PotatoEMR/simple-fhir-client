package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityResponse
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
	ServicedDate      *string                                `json:"servicedDate"`
	ServicedPeriod    *Period                                `json:"servicedPeriod"`
	Created           string                                 `json:"created"`
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

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsurance struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Coverage          Reference                                  `json:"coverage"`
	Inforce           *bool                                      `json:"inforce,omitempty"`
	BenefitPeriod     *Period                                    `json:"benefitPeriod,omitempty"`
	Item              []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityResponse
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

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	Id                 *string         `json:"id,omitempty"`
	Extension          []Extension     `json:"extension,omitempty"`
	ModifierExtension  []Extension     `json:"modifierExtension,omitempty"`
	Type               CodeableConcept `json:"type"`
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt"`
	AllowedString      *string         `json:"allowedString"`
	AllowedMoney       *Money          `json:"allowedMoney"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt"`
	UsedString         *string         `json:"usedString"`
	UsedMoney          *Money          `json:"usedMoney"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
}

type OtherCoverageEligibilityResponse CoverageEligibilityResponse

// on convert struct to json, automatically add resourceType=CoverageEligibilityResponse
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityResponse: OtherCoverageEligibilityResponse(r),
		ResourceType:                     "CoverageEligibilityResponse",
	})
}

func (resource *CoverageEligibilityResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Purpose() templ.Component {
	optionsValueSet := VSEligibilityresponse_purpose

	if resource == nil {
		return CodeSelect("purpose", nil, optionsValueSet)
	}
	return CodeSelect("purpose", &resource.Purpose[0], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Outcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Form(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("form", resource.Form, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemCategory(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Insurance[numInsurance].Item[numItem].Category, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemProductOrService(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.Insurance[numInsurance].Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemModifier(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Insurance[numInsurance].Item[numItem].Modifier[0], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemNetwork(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("network", resource.Insurance[numInsurance].Item[numItem].Network, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemUnit(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unit", resource.Insurance[numInsurance].Item[numItem].Unit, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemTerm(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("term", resource.Insurance[numInsurance].Item[numItem].Term, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationSupporting(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("authorizationSupporting", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authorizationSupporting", &resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting[0], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitType(numInsurance int, numItem int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Insurance[numInsurance].Item[numItem].Benefit) >= numBenefit {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_ErrorCode(numError int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Error) >= numError {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Error[numError].Code, optionsValueSet)
}
