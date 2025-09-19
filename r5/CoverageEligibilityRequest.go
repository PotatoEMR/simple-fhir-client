package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
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
	Event             []CoverageEligibilityRequestEvent          `json:"event,omitempty"`
	ServicedDate      *FhirDate                                  `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                    `json:"servicedPeriod,omitempty"`
	Created           FhirDateTime                               `json:"created"`
	Enterer           *Reference                                 `json:"enterer,omitempty"`
	Provider          *Reference                                 `json:"provider,omitempty"`
	Insurer           Reference                                  `json:"insurer"`
	Facility          *Reference                                 `json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `json:"item,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestEvent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	WhenDateTime      FhirDateTime    `json:"whenDateTime"`
	WhenPeriod        Period          `json:"whenPeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestSupportingInfo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Sequence          int         `json:"sequence"`
	Information       Reference   `json:"information"`
	AppliesToAll      *bool       `json:"appliesToAll,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequestInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Focal               *bool       `json:"focal,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
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

// http://hl7.org/fhir/r5/StructureDefinition/CoverageEligibilityRequest
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
func (resource *CoverageEligibilityRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Purpose(numPurpose int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEligibilityrequest_purpose

	if resource == nil || numPurpose >= len(resource.Purpose) {
		return CodeSelect("purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("purpose["+strconv.Itoa(numPurpose)+"]", &resource.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", &resource.Patient, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ServicedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("servicedDate", resource.ServicedDate, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ServicedPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("servicedPeriod", resource.ServicedPeriod, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Enterer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("enterer", nil, htmlAttrs)
	}
	return ReferenceInput("enterer", resource.Enterer, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Provider(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("provider", nil, htmlAttrs)
	}
	return ReferenceInput("provider", resource.Provider, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Insurer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("insurer", nil, htmlAttrs)
	}
	return ReferenceInput("insurer", &resource.Insurer, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_Facility(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("facility", nil, htmlAttrs)
	}
	return ReferenceInput("facility", resource.Facility, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_EventWhenDateTime(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_EventWhenPeriod(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", nil, htmlAttrs)
	}
	return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", &resource.Event[numEvent].WhenPeriod, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoInformation(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].information", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].information", &resource.SupportingInfo[numSupportingInfo].Information, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_SupportingInfoAppliesToAll(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].appliesToAll", nil, htmlAttrs)
	}
	return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].appliesToAll", resource.SupportingInfo[numSupportingInfo].AppliesToAll, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_InsuranceFocal(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_InsuranceCoverage(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", &resource.Insurance[numInsurance].Coverage, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemSupportingInfoSequence(numItem int, numSupportingInfoSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numSupportingInfoSequence >= len(resource.Item[numItem].SupportingInfoSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].supportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].supportingInfoSequence["+strconv.Itoa(numSupportingInfoSequence)+"]", &resource.Item[numItem].SupportingInfoSequence[numSupportingInfoSequence], htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemProvider(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].provider", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].provider", resource.Item[numItem].Provider, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemQuantity(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", resource.Item[numItem].Quantity, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemUnitPrice(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", resource.Item[numItem].UnitPrice, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemFacility(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].facility", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].facility", resource.Item[numItem].Facility, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemDetail(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"]", &resource.Item[numItem].Detail[numDetail], htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemDiagnosisDiagnosisCodeableConcept(numItem int, numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosis >= len(resource.Item[numItem].Diagnosis) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", resource.Item[numItem].Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *CoverageEligibilityRequest) T_ItemDiagnosisDiagnosisReference(numItem int, numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosis >= len(resource.Item[numItem].Diagnosis) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", resource.Item[numItem].Diagnosis[numDiagnosis].DiagnosisReference, htmlAttrs)
}
