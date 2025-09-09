package r5

//generated with command go run ./bultaoreune
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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_CountType(htmlAttrs string) templ.Component {
	optionsValueSet := VSInventoryreport_counttype

	if resource == nil {
		return CodeSelect("countType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("countType", &resource.CountType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_OperationType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("operationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("operationType", resource.OperationType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_OperationTypeReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("operationTypeReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("operationTypeReason", resource.OperationTypeReason, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_ReportedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("reportedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("reportedDateTime", &resource.ReportedDateTime, htmlAttrs)
}
func (resource *InventoryReport) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingItemStatus(numInventoryListing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) {
		return CodeableConceptSelect("inventoryListing["+strconv.Itoa(numInventoryListing)+"].itemStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("inventoryListing["+strconv.Itoa(numInventoryListing)+"].itemStatus", resource.InventoryListing[numInventoryListing].ItemStatus, optionsValueSet, htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingCountingDateTime(numInventoryListing int, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) {
		return DateTimeInput("inventoryListing["+strconv.Itoa(numInventoryListing)+"].countingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("inventoryListing["+strconv.Itoa(numInventoryListing)+"].countingDateTime", resource.InventoryListing[numInventoryListing].CountingDateTime, htmlAttrs)
}
func (resource *InventoryReport) T_InventoryListingItemCategory(numInventoryListing int, numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryListing >= len(resource.InventoryListing) || numItem >= len(resource.InventoryListing[numInventoryListing].Item) {
		return CodeableConceptSelect("inventoryListing["+strconv.Itoa(numInventoryListing)+"].item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("inventoryListing["+strconv.Itoa(numInventoryListing)+"].item["+strconv.Itoa(numItem)+"].category", resource.InventoryListing[numInventoryListing].Item[numItem].Category, optionsValueSet, htmlAttrs)
}
