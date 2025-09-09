package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
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
	Event             []CoverageEligibilityResponseEvent     `json:"event,omitempty"`
	ServicedDate      *time.Time                             `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod    *Period                                `json:"servicedPeriod,omitempty"`
	Created           time.Time                              `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseEvent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	WhenDateTime      time.Time       `json:"whenDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	WhenPeriod        Period          `json:"whenPeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsurance struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Coverage          Reference                                  `json:"coverage"`
	Inforce           *bool                                      `json:"inforce,omitempty"`
	BenefitPeriod     *Period                                    `json:"benefitPeriod,omitempty"`
	Item              []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
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

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
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

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Expression        []string        `json:"expression,omitempty"`
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
func (resource *CoverageEligibilityResponse) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("CoverageEligibilityResponse.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CoverageEligibilityResponse.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Purpose(numPurpose int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEligibilityresponse_purpose

	if resource == nil || numPurpose >= len(resource.Purpose) {
		return CodeSelect("CoverageEligibilityResponse.Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CoverageEligibilityResponse.Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ServicedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("CoverageEligibilityResponse.ServicedDate", nil, htmlAttrs)
	}
	return DateInput("CoverageEligibilityResponse.ServicedDate", resource.ServicedDate, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CoverageEligibilityResponse.Created", nil, htmlAttrs)
	}
	return DateTimeInput("CoverageEligibilityResponse.Created", &resource.Created, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSEligibility_outcome

	if resource == nil {
		return CodeSelect("CoverageEligibilityResponse.Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CoverageEligibilityResponse.Outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CoverageEligibilityResponse.Disposition", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Disposition", resource.Disposition, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_PreAuthRef(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CoverageEligibilityResponse.PreAuthRef", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.PreAuthRef", resource.PreAuthRef, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_Form(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CoverageEligibilityResponse.Form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Form", resource.Form, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Event["+strconv.Itoa(numEvent)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Event["+strconv.Itoa(numEvent)+"].Type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_EventWhenDateTime(numEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return DateTimeInput("CoverageEligibilityResponse.Event["+strconv.Itoa(numEvent)+"].WhenDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("CoverageEligibilityResponse.Event["+strconv.Itoa(numEvent)+"].WhenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceInforce(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Inforce", nil, htmlAttrs)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Inforce", resource.Insurance[numInsurance].Inforce, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemCategory(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Category", resource.Insurance[numInsurance].Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemProductOrService(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Insurance[numInsurance].Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemModifier(numInsurance int, numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numModifier >= len(resource.Insurance[numInsurance].Item[numItem].Modifier) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Insurance[numInsurance].Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemExcluded(numInsurance int, numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Excluded", nil, htmlAttrs)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Excluded", resource.Insurance[numInsurance].Item[numItem].Excluded, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemName(numInsurance int, numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Name", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Name", resource.Insurance[numInsurance].Item[numItem].Name, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemDescription(numInsurance int, numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Description", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Description", resource.Insurance[numInsurance].Item[numItem].Description, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemNetwork(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Network", resource.Insurance[numInsurance].Item[numItem].Network, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemUnit(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Unit", resource.Insurance[numInsurance].Item[numItem].Unit, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemTerm(numInsurance int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Term", resource.Insurance[numInsurance].Item[numItem].Term, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationRequired(numInsurance int, numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationRequired", nil, htmlAttrs)
	}
	return BoolInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationRequired", resource.Insurance[numInsurance].Item[numItem].AuthorizationRequired, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationSupporting(numInsurance int, numItem int, numAuthorizationSupporting int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numAuthorizationSupporting >= len(resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationSupporting["+strconv.Itoa(numAuthorizationSupporting)+"]", &resource.Insurance[numInsurance].Item[numItem].AuthorizationSupporting[numAuthorizationSupporting], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemAuthorizationUrl(numInsurance int, numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationUrl", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].AuthorizationUrl", resource.Insurance[numInsurance].Item[numItem].AuthorizationUrl, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitType(numInsurance int, numItem int, numBenefit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].Type", &resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].Type, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitAllowedUnsignedInt(numInsurance int, numItem int, numBenefit int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return IntInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].AllowedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].AllowedUnsignedInt", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].AllowedUnsignedInt, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitAllowedString(numInsurance int, numItem int, numBenefit int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].AllowedString", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].AllowedString", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].AllowedString, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitUsedUnsignedInt(numInsurance int, numItem int, numBenefit int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return IntInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].UsedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].UsedUnsignedInt", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].UsedUnsignedInt, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_InsuranceItemBenefitUsedString(numInsurance int, numItem int, numBenefit int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numItem >= len(resource.Insurance[numInsurance].Item) || numBenefit >= len(resource.Insurance[numInsurance].Item[numItem].Benefit) {
		return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].UsedString", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Insurance["+strconv.Itoa(numInsurance)+"].Item["+strconv.Itoa(numItem)+"].Benefit["+strconv.Itoa(numBenefit)+"].UsedString", resource.Insurance[numInsurance].Item[numItem].Benefit[numBenefit].UsedString, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ErrorCode(numError int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Code", &resource.Error[numError].Code, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityResponse) T_ErrorExpression(numError int, numExpression int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) || numExpression >= len(resource.Error[numError].Expression) {
		return StringInput("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Expression["+strconv.Itoa(numExpression)+"]", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityResponse.Error["+strconv.Itoa(numError)+"].Expression["+strconv.Itoa(numExpression)+"]", &resource.Error[numError].Expression[numExpression], htmlAttrs)
}
