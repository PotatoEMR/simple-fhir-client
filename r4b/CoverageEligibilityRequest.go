package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequest struct {
	Id                *string                                    `json:"id,omitempty"`
	Meta              *Meta                                      `json:"meta,omitempty"`
	ImplicitRules     *string                                    `json:"implicitRules,omitempty"`
	Language          *string                                    `json:"language,omitempty"`
	Text              *Narrative                                 `json:"text,omitempty"`
	Contained         []Resource                                 `json:"contained,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `json:"identifier,omitempty"`
	Status            string                                     `json:"status"`
	Priority          *CodeableConcept                           `json:"priority,omitempty"`
	Purpose           []string                                   `json:"purpose"`
	Patient           Reference                                  `json:"patient"`
	ServicedDate      *string                                    `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                    `json:"servicedPeriod,omitempty"`
	Created           string                                     `json:"created"`
	Enterer           *Reference                                 `json:"enterer,omitempty"`
	Provider          *Reference                                 `json:"provider,omitempty"`
	Insurer           Reference                                  `json:"insurer"`
	Facility          *Reference                                 `json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestSupportingInfo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Sequence          int         `json:"sequence"`
	Information       Reference   `json:"information"`
	AppliesToAll      *bool       `json:"appliesToAll,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Focal               *bool       `json:"focal,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestItem struct {
	Id                     *string                                   `json:"id,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	SupportingInfoSequence []int                                     `json:"supportingInfoSequence,omitempty"`
	Category               *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService       *CodeableConcept                          `json:"productOrService,omitempty"`
	Modifier               []CodeableConcept                         `json:"modifier,omitempty"`
	Provider               *Reference                                `json:"provider,omitempty"`
	Quantity               *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice              *Money                                    `json:"unitPrice,omitempty"`
	Facility               *Reference                                `json:"facility,omitempty"`
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`
	Detail                 []Reference                               `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestItemDiagnosis struct {
	Id                       *string          `json:"id,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference       `json:"diagnosisReference,omitempty"`
}

type OtherCoverageEligibilityRequest CoverageEligibilityRequest

// on convert struct to json, automatically add resourceType=CoverageEligibilityRequest
func (r CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityRequest: OtherCoverageEligibilityRequest(r),
		ResourceType:                    "CoverageEligibilityRequest",
	})
}

func (resource *CoverageEligibilityRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityRequest.Id", nil)
	}
	return StringInput("CoverageEligibilityRequest.Id", resource.Id)
}
func (resource *CoverageEligibilityRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityRequest.ImplicitRules", nil)
	}
	return StringInput("CoverageEligibilityRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *CoverageEligibilityRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CoverageEligibilityRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityRequest.Language", resource.Language, optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("CoverageEligibilityRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityRequest.Status", &resource.Status, optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("CoverageEligibilityRequest.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_Purpose(numPurpose int) templ.Component {
	optionsValueSet := VSEligibilityrequest_purpose

	if resource == nil || len(resource.Purpose) >= numPurpose {
		return CodeSelect("CoverageEligibilityRequest.Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet)
	}
	return CodeSelect("CoverageEligibilityRequest.Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_Created() templ.Component {

	if resource == nil {
		return StringInput("CoverageEligibilityRequest.Created", nil)
	}
	return StringInput("CoverageEligibilityRequest.Created", &resource.Created)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoId(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return StringInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", resource.SupportingInfo[numSupportingInfo].Id)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoSequence(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return IntInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", nil)
	}
	return IntInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoAppliesToAll(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return BoolInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].AppliesToAll", nil)
	}
	return BoolInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].AppliesToAll", resource.SupportingInfo[numSupportingInfo].AppliesToAll)
}
func (resource *CoverageEligibilityRequest) T_InsuranceId(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Id", resource.Insurance[numInsurance].Id)
}
func (resource *CoverageEligibilityRequest) T_InsuranceFocal(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return BoolInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil)
	}
	return BoolInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Focal", resource.Insurance[numInsurance].Focal)
}
func (resource *CoverageEligibilityRequest) T_InsuranceBusinessArrangement(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", nil)
	}
	return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement)
}
func (resource *CoverageEligibilityRequest) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *CoverageEligibilityRequest) T_ItemSupportingInfoSequence(numItem int, numSupportingInfoSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].SupportingInfoSequence) >= numSupportingInfoSequence {
		return IntInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].SupportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", nil)
	}
	return IntInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].SupportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", &resource.Item[numItem].SupportingInfoSequence[numSupportingInfoSequence])
}
func (resource *CoverageEligibilityRequest) T_ItemCategory(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Category", resource.Item[numItem].Category, optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_ItemProductOrService(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Modifier) >= numModifier {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet)
}
func (resource *CoverageEligibilityRequest) T_ItemDiagnosisId(numItem int, numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Diagnosis) >= numDiagnosis {
		return StringInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Item[numItem].Diagnosis[numDiagnosis].Id)
}
