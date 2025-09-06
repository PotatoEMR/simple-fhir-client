package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/InventoryReport
type InventoryReport struct {
	Id                  *string                           `json:"id,omitempty"`
	Meta                *Meta                             `json:"meta,omitempty"`
	ImplicitRules       *string                           `json:"implicitRules,omitempty"`
	Language            *string                           `json:"language,omitempty"`
	Text                *Narrative                        `json:"text,omitempty"`
	Contained           []Resource                        `json:"contained,omitempty"`
	Extension           []Extension                       `json:"extension,omitempty"`
	ModifierExtension   []Extension                       `json:"modifierExtension,omitempty"`
	Identifier          []Identifier                      `json:"identifier,omitempty"`
	Status              string                            `json:"status"`
	CountType           string                            `json:"countType"`
	OperationType       *CodeableConcept                  `json:"operationType,omitempty"`
	OperationTypeReason *CodeableConcept                  `json:"operationTypeReason,omitempty"`
	ReportedDateTime    string                            `json:"reportedDateTime"`
	Reporter            *Reference                        `json:"reporter,omitempty"`
	ReportingPeriod     *Period                           `json:"reportingPeriod,omitempty"`
	InventoryListing    []InventoryReportInventoryListing `json:"inventoryListing,omitempty"`
	Note                []Annotation                      `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryReport
type InventoryReportInventoryListing struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Location          *Reference                            `json:"location,omitempty"`
	ItemStatus        *CodeableConcept                      `json:"itemStatus,omitempty"`
	CountingDateTime  *string                               `json:"countingDateTime,omitempty"`
	Item              []InventoryReportInventoryListingItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryReport
type InventoryReportInventoryListingItem struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept  `json:"category,omitempty"`
	Quantity          Quantity          `json:"quantity"`
	Item              CodeableReference `json:"item"`
}

type OtherInventoryReport InventoryReport

// on convert struct to json, automatically add resourceType=InventoryReport
func (r InventoryReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInventoryReport
		ResourceType string `json:"resourceType"`
	}{
		OtherInventoryReport: OtherInventoryReport(r),
		ResourceType:         "InventoryReport",
	})
}

func (resource *InventoryReport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("InventoryReport.Id", nil)
	}
	return StringInput("InventoryReport.Id", resource.Id)
}
func (resource *InventoryReport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("InventoryReport.ImplicitRules", nil)
	}
	return StringInput("InventoryReport.ImplicitRules", resource.ImplicitRules)
}
func (resource *InventoryReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("InventoryReport.Language", nil, optionsValueSet)
	}
	return CodeSelect("InventoryReport.Language", resource.Language, optionsValueSet)
}
func (resource *InventoryReport) T_Status() templ.Component {
	optionsValueSet := VSInventoryreport_status

	if resource == nil {
		return CodeSelect("InventoryReport.Status", nil, optionsValueSet)
	}
	return CodeSelect("InventoryReport.Status", &resource.Status, optionsValueSet)
}
func (resource *InventoryReport) T_CountType() templ.Component {
	optionsValueSet := VSInventoryreport_counttype

	if resource == nil {
		return CodeSelect("InventoryReport.CountType", nil, optionsValueSet)
	}
	return CodeSelect("InventoryReport.CountType", &resource.CountType, optionsValueSet)
}
func (resource *InventoryReport) T_OperationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("InventoryReport.OperationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryReport.OperationType", resource.OperationType, optionsValueSet)
}
func (resource *InventoryReport) T_OperationTypeReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("InventoryReport.OperationTypeReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryReport.OperationTypeReason", resource.OperationTypeReason, optionsValueSet)
}
func (resource *InventoryReport) T_ReportedDateTime() templ.Component {

	if resource == nil {
		return StringInput("InventoryReport.ReportedDateTime", nil)
	}
	return StringInput("InventoryReport.ReportedDateTime", &resource.ReportedDateTime)
}
func (resource *InventoryReport) T_InventoryListingId(numInventoryListing int) templ.Component {

	if resource == nil || len(resource.InventoryListing) >= numInventoryListing {
		return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Id", nil)
	}
	return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Id", resource.InventoryListing[numInventoryListing].Id)
}
func (resource *InventoryReport) T_InventoryListingItemStatus(numInventoryListing int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.InventoryListing) >= numInventoryListing {
		return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].ItemStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].ItemStatus", resource.InventoryListing[numInventoryListing].ItemStatus, optionsValueSet)
}
func (resource *InventoryReport) T_InventoryListingCountingDateTime(numInventoryListing int) templ.Component {

	if resource == nil || len(resource.InventoryListing) >= numInventoryListing {
		return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].CountingDateTime", nil)
	}
	return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].CountingDateTime", resource.InventoryListing[numInventoryListing].CountingDateTime)
}
func (resource *InventoryReport) T_InventoryListingItemId(numInventoryListing int, numItem int) templ.Component {

	if resource == nil || len(resource.InventoryListing) >= numInventoryListing || len(resource.InventoryListing[numInventoryListing].Item) >= numItem {
		return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Id", resource.InventoryListing[numInventoryListing].Item[numItem].Id)
}
func (resource *InventoryReport) T_InventoryListingItemCategory(numInventoryListing int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.InventoryListing) >= numInventoryListing || len(resource.InventoryListing[numInventoryListing].Item) >= numItem {
		return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Category", resource.InventoryListing[numInventoryListing].Item[numItem].Category, optionsValueSet)
}
