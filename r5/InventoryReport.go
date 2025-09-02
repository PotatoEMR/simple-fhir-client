package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *InventoryReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *InventoryReport) T_Status() templ.Component {
	optionsValueSet := VSInventoryreport_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *InventoryReport) T_CountType() templ.Component {
	optionsValueSet := VSInventoryreport_counttype

	if resource == nil {
		return CodeSelect("countType", nil, optionsValueSet)
	}
	return CodeSelect("countType", &resource.CountType, optionsValueSet)
}
func (resource *InventoryReport) T_OperationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("operationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("operationType", resource.OperationType, optionsValueSet)
}
func (resource *InventoryReport) T_OperationTypeReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("operationTypeReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("operationTypeReason", resource.OperationTypeReason, optionsValueSet)
}
func (resource *InventoryReport) T_InventoryListingItemStatus(numInventoryListing int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.InventoryListing) >= numInventoryListing {
		return CodeableConceptSelect("itemStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("itemStatus", resource.InventoryListing[numInventoryListing].ItemStatus, optionsValueSet)
}
func (resource *InventoryReport) T_InventoryListingItemCategory(numInventoryListing int, numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.InventoryListing[numInventoryListing].Item) >= numItem {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.InventoryListing[numInventoryListing].Item[numItem].Category, optionsValueSet)
}
