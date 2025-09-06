package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/CatalogEntry
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
	ValidTo                  *string                    `json:"validTo,omitempty"`
	LastUpdated              *string                    `json:"lastUpdated,omitempty"`
	AdditionalCharacteristic []CodeableConcept          `json:"additionalCharacteristic,omitempty"`
	AdditionalClassification []CodeableConcept          `json:"additionalClassification,omitempty"`
	RelatedEntry             []CatalogEntryRelatedEntry `json:"relatedEntry,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CatalogEntry
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

func (resource *CatalogEntry) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CatalogEntry.Id", nil)
	}
	return StringInput("CatalogEntry.Id", resource.Id)
}
func (resource *CatalogEntry) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CatalogEntry.ImplicitRules", nil)
	}
	return StringInput("CatalogEntry.ImplicitRules", resource.ImplicitRules)
}
func (resource *CatalogEntry) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CatalogEntry.Language", nil, optionsValueSet)
	}
	return CodeSelect("CatalogEntry.Language", resource.Language, optionsValueSet)
}
func (resource *CatalogEntry) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("CatalogEntry.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CatalogEntry.Type", resource.Type, optionsValueSet)
}
func (resource *CatalogEntry) T_Orderable() templ.Component {

	if resource == nil {
		return BoolInput("CatalogEntry.Orderable", nil)
	}
	return BoolInput("CatalogEntry.Orderable", &resource.Orderable)
}
func (resource *CatalogEntry) T_Classification(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("CatalogEntry.Classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CatalogEntry.Classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet)
}
func (resource *CatalogEntry) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CatalogEntry.Status", nil, optionsValueSet)
	}
	return CodeSelect("CatalogEntry.Status", resource.Status, optionsValueSet)
}
func (resource *CatalogEntry) T_ValidTo() templ.Component {

	if resource == nil {
		return StringInput("CatalogEntry.ValidTo", nil)
	}
	return StringInput("CatalogEntry.ValidTo", resource.ValidTo)
}
func (resource *CatalogEntry) T_LastUpdated() templ.Component {

	if resource == nil {
		return StringInput("CatalogEntry.LastUpdated", nil)
	}
	return StringInput("CatalogEntry.LastUpdated", resource.LastUpdated)
}
func (resource *CatalogEntry) T_AdditionalCharacteristic(numAdditionalCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AdditionalCharacteristic) >= numAdditionalCharacteristic {
		return CodeableConceptSelect("CatalogEntry.AdditionalCharacteristic["+strconv.Itoa(numAdditionalCharacteristic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CatalogEntry.AdditionalCharacteristic["+strconv.Itoa(numAdditionalCharacteristic)+"]", &resource.AdditionalCharacteristic[numAdditionalCharacteristic], optionsValueSet)
}
func (resource *CatalogEntry) T_AdditionalClassification(numAdditionalClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AdditionalClassification) >= numAdditionalClassification {
		return CodeableConceptSelect("CatalogEntry.AdditionalClassification["+strconv.Itoa(numAdditionalClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CatalogEntry.AdditionalClassification["+strconv.Itoa(numAdditionalClassification)+"]", &resource.AdditionalClassification[numAdditionalClassification], optionsValueSet)
}
func (resource *CatalogEntry) T_RelatedEntryId(numRelatedEntry int) templ.Component {

	if resource == nil || len(resource.RelatedEntry) >= numRelatedEntry {
		return StringInput("CatalogEntry.RelatedEntry["+strconv.Itoa(numRelatedEntry)+"].Id", nil)
	}
	return StringInput("CatalogEntry.RelatedEntry["+strconv.Itoa(numRelatedEntry)+"].Id", resource.RelatedEntry[numRelatedEntry].Id)
}
func (resource *CatalogEntry) T_RelatedEntryRelationtype(numRelatedEntry int) templ.Component {
	optionsValueSet := VSRelation_type

	if resource == nil || len(resource.RelatedEntry) >= numRelatedEntry {
		return CodeSelect("CatalogEntry.RelatedEntry["+strconv.Itoa(numRelatedEntry)+"].Relationtype", nil, optionsValueSet)
	}
	return CodeSelect("CatalogEntry.RelatedEntry["+strconv.Itoa(numRelatedEntry)+"].Relationtype", &resource.RelatedEntry[numRelatedEntry].Relationtype, optionsValueSet)
}
