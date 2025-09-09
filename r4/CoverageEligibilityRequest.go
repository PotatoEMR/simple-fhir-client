package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityRequest
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
	ServicedDate      *time.Time                                 `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod    *Period                                    `json:"servicedPeriod,omitempty"`
	Created           time.Time                                  `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Enterer           *Reference                                 `json:"enterer,omitempty"`
	Provider          *Reference                                 `json:"provider,omitempty"`
	Insurer           Reference                                  `json:"insurer"`
	Facility          *Reference                                 `json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestSupportingInfo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Sequence          int         `json:"sequence"`
	Information       Reference   `json:"information"`
	AppliesToAll      *bool       `json:"appliesToAll,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Focal               *bool       `json:"focal,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityRequest
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

// http://hl7.org/fhir/r4/StructureDefinition/CoverageEligibilityRequest
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
func (r CoverageEligibilityRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CoverageEligibilityRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CoverageEligibilityRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CoverageEligibilityRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("CoverageEligibilityRequest.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CoverageEligibilityRequest.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CoverageEligibilityRequest.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Purpose(numPurpose int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEligibilityrequest_purpose

	if resource == nil || numPurpose >= len(resource.Purpose) {
		return CodeSelect("CoverageEligibilityRequest.Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CoverageEligibilityRequest.Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ServicedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("CoverageEligibilityRequest.ServicedDate", nil, htmlAttrs)
	}
	return DateInput("CoverageEligibilityRequest.ServicedDate", resource.ServicedDate, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CoverageEligibilityRequest.Created", nil, htmlAttrs)
	}
	return DateTimeInput("CoverageEligibilityRequest.Created", &resource.Created, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoAppliesToAll(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].AppliesToAll", nil, htmlAttrs)
	}
	return BoolInput("CoverageEligibilityRequest.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].AppliesToAll", resource.SupportingInfo[numSupportingInfo].AppliesToAll, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil, htmlAttrs)
	}
	return BoolInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].Focal", resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", nil, htmlAttrs)
	}
	return StringInput("CoverageEligibilityRequest.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemSupportingInfoSequence(numItem int, numSupportingInfoSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numSupportingInfoSequence >= len(resource.Item[numItem].SupportingInfoSequence) {
		return IntInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].SupportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].SupportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", &resource.Item[numItem].SupportingInfoSequence[numSupportingInfoSequence], htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemDiagnosisDiagnosisCodeableConcept(numItem int, numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosis >= len(resource.Item[numItem].Diagnosis) {
		return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Diagnosis["+strconv.Itoa(numDiagnosis)+"].DiagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CoverageEligibilityRequest.Item["+strconv.Itoa(numItem)+"].Diagnosis["+strconv.Itoa(numDiagnosis)+"].DiagnosisCodeableConcept", resource.Item[numItem].Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
