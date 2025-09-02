package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *CatalogEntry) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CatalogEntry) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *CatalogEntry) T_Classification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classification", &resource.Classification[0], optionsValueSet)
}
func (resource *CatalogEntry) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *CatalogEntry) T_AdditionalCharacteristic(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("additionalCharacteristic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("additionalCharacteristic", &resource.AdditionalCharacteristic[0], optionsValueSet)
}
func (resource *CatalogEntry) T_AdditionalClassification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("additionalClassification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("additionalClassification", &resource.AdditionalClassification[0], optionsValueSet)
}
func (resource *CatalogEntry) T_RelatedEntryRelationtype(numRelatedEntry int) templ.Component {
	optionsValueSet := VSRelation_type

	if resource == nil && len(resource.RelatedEntry) >= numRelatedEntry {
		return CodeSelect("relationtype", nil, optionsValueSet)
	}
	return CodeSelect("relationtype", &resource.RelatedEntry[numRelatedEntry].Relationtype, optionsValueSet)
}
