package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *CoverageEligibilityResponse) CoverageEligibilityResponseLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) CoverageEligibilityResponseStatus() templ.Component {
	optionsValueSet := VSFm_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) CoverageEligibilityResponsePurpose() templ.Component {
	optionsValueSet := VSEligibilityresponse_purpose
	currentVal := ""
	if resource != nil {
		currentVal = resource.Purpose[0]
	}
	return CodeSelect("purpose", currentVal, optionsValueSet)
}
func (resource *CoverageEligibilityResponse) CoverageEligibilityResponseOutcome() templ.Component {
	optionsValueSet := VSRemittance_outcome
	currentVal := ""
	if resource != nil {
		currentVal = resource.Outcome
	}
	return CodeSelect("outcome", currentVal, optionsValueSet)
}
