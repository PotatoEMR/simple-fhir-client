package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/CatalogEntry
type CatalogEntry struct {
	Id                       *string                    `json:"id,omitempty"`
	Meta                     *Meta                      `json:"meta,omitempty"`
	ImplicitRules            *string                    `json:"implicitRules,omitempty"`
	Language                 *string                    `json:"language,omitempty"`
	Text                     *Narrative                 `json:"text,omitempty"`
	Contained                []Resource                 `json:"contained,omitempty"`
	Extension                []Extension                `json:"extension,omitempty"`
	ModifierExtension        []Extension                `json:"modifierExtension,omitempty"`
	Identifier               []Identifier               `json:"identifier,omitempty"`
	Type                     *CodeableConcept           `json:"type,omitempty"`
	Orderable                bool                       `json:"orderable"`
	ReferencedItem           Reference                  `json:"referencedItem"`
	AdditionalIdentifier     []Identifier               `json:"additionalIdentifier,omitempty"`
	Classification           []CodeableConcept          `json:"classification,omitempty"`
	Status                   *string                    `json:"status,omitempty"`
	ValidityPeriod           *Period                    `json:"validityPeriod,omitempty"`
	ValidTo                  *time.Time                 `json:"validTo,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	LastUpdated              *time.Time                 `json:"lastUpdated,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	AdditionalCharacteristic []CodeableConcept          `json:"additionalCharacteristic,omitempty"`
	AdditionalClassification []CodeableConcept          `json:"additionalClassification,omitempty"`
	RelatedEntry             []CatalogEntryRelatedEntry `json:"relatedEntry,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CatalogEntry
type CatalogEntryRelatedEntry struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Relationtype      string      `json:"relationtype"`
	Item              Reference   `json:"item"`
}

type OtherCatalogEntry CatalogEntry

// on convert struct to json, automatically add resourceType=CatalogEntry
func (r CatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCatalogEntry
		ResourceType string `json:"resourceType"`
	}{
		OtherCatalogEntry: OtherCatalogEntry(r),
		ResourceType:      "CatalogEntry",
	})
}
func (r CatalogEntry) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CatalogEntry/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CatalogEntry"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CatalogEntry) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("CatalogEntry.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CatalogEntry.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_Orderable(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("CatalogEntry.Orderable", nil, htmlAttrs)
	}
	return BoolInput("CatalogEntry.Orderable", &resource.Orderable, htmlAttrs)
}
func (resource *CatalogEntry) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("CatalogEntry.Classification."+strconv.Itoa(numClassification)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CatalogEntry.Classification."+strconv.Itoa(numClassification)+".", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CatalogEntry.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CatalogEntry.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_ValidTo(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CatalogEntry.ValidTo", nil, htmlAttrs)
	}
	return DateTimeInput("CatalogEntry.ValidTo", resource.ValidTo, htmlAttrs)
}
func (resource *CatalogEntry) T_LastUpdated(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CatalogEntry.LastUpdated", nil, htmlAttrs)
	}
	return DateTimeInput("CatalogEntry.LastUpdated", resource.LastUpdated, htmlAttrs)
}
func (resource *CatalogEntry) T_AdditionalCharacteristic(numAdditionalCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numAdditionalCharacteristic >= len(resource.AdditionalCharacteristic) {
		return CodeableConceptSelect("CatalogEntry.AdditionalCharacteristic."+strconv.Itoa(numAdditionalCharacteristic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CatalogEntry.AdditionalCharacteristic."+strconv.Itoa(numAdditionalCharacteristic)+".", &resource.AdditionalCharacteristic[numAdditionalCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_AdditionalClassification(numAdditionalClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numAdditionalClassification >= len(resource.AdditionalClassification) {
		return CodeableConceptSelect("CatalogEntry.AdditionalClassification."+strconv.Itoa(numAdditionalClassification)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CatalogEntry.AdditionalClassification."+strconv.Itoa(numAdditionalClassification)+".", &resource.AdditionalClassification[numAdditionalClassification], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_RelatedEntryRelationtype(numRelatedEntry int, htmlAttrs string) templ.Component {
	optionsValueSet := VSRelation_type

	if resource == nil || numRelatedEntry >= len(resource.RelatedEntry) {
		return CodeSelect("CatalogEntry.RelatedEntry."+strconv.Itoa(numRelatedEntry)+"..Relationtype", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CatalogEntry.RelatedEntry."+strconv.Itoa(numRelatedEntry)+"..Relationtype", &resource.RelatedEntry[numRelatedEntry].Relationtype, optionsValueSet, htmlAttrs)
}
