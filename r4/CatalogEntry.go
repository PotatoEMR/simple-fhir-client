//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
