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
	ReportedDateTime    time.Time                         `json:"reportedDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
	CountingDateTime  *time.Time                            `json:"countingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r InventoryReport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "InventoryReport/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "InventoryReport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *InventoryReport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSInventoryreport_status

	if resource == nil {
		return CodeSelect("InventoryReport.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("InventoryReport.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_CountType(htmlAttrs string) templ.Component {
	optionsValueSet := VSInventoryreport_counttype

	if resource == nil {
		return CodeSelect("InventoryReport.CountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("InventoryReport.CountType", &resource.CountType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_OperationType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("InventoryReport.OperationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InventoryReport.OperationType", resource.OperationType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_OperationTypeReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("InventoryReport.OperationTypeReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InventoryReport.OperationTypeReason", resource.OperationTypeReason, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_ReportedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("InventoryReport.ReportedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("InventoryReport.ReportedDateTime", &resource.ReportedDateTime, htmlAttrs)
}
func (resource *InventoryReport) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("InventoryReport.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("InventoryReport.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingItemStatus(numInventoryListing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) {
		return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].ItemStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].ItemStatus", resource.InventoryListing[numInventoryListing].ItemStatus, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingCountingDateTime(numInventoryListing int, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) {
		return DateTimeInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].CountingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].CountingDateTime", resource.InventoryListing[numInventoryListing].CountingDateTime, htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingItemCategory(numInventoryListing int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) || numItem >= len(resource.InventoryListing[numInventoryListing].Item) {
		return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InventoryReport.InventoryListing["+strconv.Itoa(numInventoryListing)+"].Item["+strconv.Itoa(numItem)+"].Category", resource.InventoryListing[numInventoryListing].Item[numItem].Category, optionsValueSet, htmlAttrs)
}
