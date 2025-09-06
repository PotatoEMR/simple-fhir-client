package r4b

//generated with command go run ./bultaoreune -nodownload
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
	ServicedDate      *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                `json:"servicedPeriod,omitempty"`
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

func (resource *CoverageEligibilityResponse) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityResponse.Id", nil)
	}
	return StringInput("CoverageEligibilityResponse.Id", resource.Id)
}
func (resource *CoverageEligibilityResponse) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityResponse.ImplicitRules", nil)
	}
	return StringInput("CoverageEligibilityResponse.ImplicitRules", resource.ImplicitRules)
}
func (resource *CoverageEligibilityResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CoverageEligibilityResponse.Language", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityResponse.Language", resource.Language, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("CoverageEligibilityResponse.Status", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityResponse.Status", &resource.Status, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Purpose(numPurpose int) templ.Component {
	optionsValueSet := VSEligibilityresponse_purpose

	if resource == nil || len(resource.Purpose) >= numPurpose {
		return CodeSelect("CoverageEligibilityResponse.Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityResponse.Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Created() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityResponse.Created", nil)
	}
	return StringInput("CoverageEligibilityResponse.Created", &resource.Created)
}
func (resource *CoverageEligibilityResponse) T_Outcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("CoverageEligibilityResponse.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityResponse.Outcome", &resource.Outcome, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_Disposition() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityResponse.Disposition", nil)
	}
	return StringInput("CoverageEligibilityResponse.Disposition", resource.Disposition)
}
func (resource *CoverageEligibilityResponse) T_PreAuthRef() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityResponse.PreAuthRef", nil)
	}
	return StringInput("CoverageEligibilityResponse.PreAuthRef", resource.PreAuthRef)
}
func (resource *CoverageEligibilityResponse) T_Form(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("CoverageEligibilityResponse.Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Form", resource.Form, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceId(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Id", resource.Insurance[numInsurance].Id)
}
func (resource *CoverageEligibilityResponse) T_InsuranceInforce(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Inforce", nil)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Inforce", resource.Insurance[numInsurance].Inforce)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemId(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Id", resource.Insurance[numInsurance].Item[numItem].Id)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemCategory(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Category", resource.Insurance[numInsurance].Item[numItem].Category, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemProductOrService(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Insurance[numInsurance].Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemModifier(numInsurance int, numItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem || len(resource.Insurance[numInsurance].Item[numItem].Modifier) >= numModifier {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Insurance[numInsurance].Item[numItem].Modifier[numModifier], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemExcluded(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Excluded", nil)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Excluded", resource.Insurance[numInsurance].Item[numItem].Excluded)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemName(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Name", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Name", resource.Insurance[numInsurance].Item[numItem].Name)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemDescription(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Description", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Description", resource.Insurance[numInsurance].Item[numItem].Description)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemNetwork(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Network", resource.Insurance[numInsurance].Item[numItem].Network, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemUnit(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Unit", resource.Insurance[numInsurance].Item[numItem].Unit, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemTerm(numInsurance int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Term", resource.Insurance[numInsurance].Item[numItem].Term, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationRequired(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationRequired", nil)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationRequired", resource.Insurance[numInsurance].Item[numItem].AuthorizationRequired)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationSupporting(numInsurance int, numItem int, numAuthorizationSupporting int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem || len(resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting) >= numAuthorizationSupporting {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", &resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting[numAuthorizationSupporting], optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationUrl(numInsurance int, numItem int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationUrl", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationUrl", resource.Insurance[numInsurance].Item[numItem].AuthorizationUrl)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitId(numInsurance int, numItem int, numBenefit int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem || len(resource.Insurance[numInsurance].Item[numItem].Benefit) >= numBenefit {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Id", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].Id)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitType(numInsurance int, numItem int, numBenefit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].Item) >= numItem || len(resource.Insurance[numInsurance].Item[numItem].Benefit) >= numBenefit {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].Type, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) T_ErrorId(numError int) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return StringInput("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Id", resource.Error[numError].Id)
}
func (resource *CoverageEligibilityResponse) T_ErrorCode(numError int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return CodeableConceptSelect("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Code", &resource.Error[numError].Code, optionsValueSet)
}
