package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	ValidTo                  *FhirDateTime              `json:"validTo,omitempty"`
	LastUpdated              *FhirDateTime              `json:"lastUpdated,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r CatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCatalogEntry
		ResourceType string `json:"resourceType"`
	}{
		OtherCatalogEntry: OtherCatalogEntry(r),
		ResourceType:      "CatalogEntry",
	})
}

// json -> struct, first reject if resourceType != CatalogEntry
func (r *CatalogEntry) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "CatalogEntry" {
		return errors.New("resourceType not CatalogEntry")
	}
	return json.Unmarshal(data, (*OtherCatalogEntry)(r))
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
func (resource *CatalogEntry) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_Orderable(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("orderable", nil, htmlAttrs)
	}
	return BoolInput("orderable", &resource.Orderable, htmlAttrs)
}
func (resource *CatalogEntry) T_ReferencedItem(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "referencedItem", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "referencedItem", &resource.ReferencedItem, htmlAttrs)
}
func (resource *CatalogEntry) T_AdditionalIdentifier(numAdditionalIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdditionalIdentifier >= len(resource.AdditionalIdentifier) {
		return IdentifierInput("additionalIdentifier["+strconv.Itoa(numAdditionalIdentifier)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("additionalIdentifier["+strconv.Itoa(numAdditionalIdentifier)+"]", &resource.AdditionalIdentifier[numAdditionalIdentifier], htmlAttrs)
}
func (resource *CatalogEntry) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_ValidityPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("validityPeriod", nil, htmlAttrs)
	}
	return PeriodInput("validityPeriod", resource.ValidityPeriod, htmlAttrs)
}
func (resource *CatalogEntry) T_ValidTo(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("validTo", nil, htmlAttrs)
	}
	return FhirDateTimeInput("validTo", resource.ValidTo, htmlAttrs)
}
func (resource *CatalogEntry) T_LastUpdated(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("lastUpdated", nil, htmlAttrs)
	}
	return FhirDateTimeInput("lastUpdated", resource.LastUpdated, htmlAttrs)
}
func (resource *CatalogEntry) T_AdditionalCharacteristic(numAdditionalCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdditionalCharacteristic >= len(resource.AdditionalCharacteristic) {
		return CodeableConceptSelect("additionalCharacteristic["+strconv.Itoa(numAdditionalCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("additionalCharacteristic["+strconv.Itoa(numAdditionalCharacteristic)+"]", &resource.AdditionalCharacteristic[numAdditionalCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_AdditionalClassification(numAdditionalClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdditionalClassification >= len(resource.AdditionalClassification) {
		return CodeableConceptSelect("additionalClassification["+strconv.Itoa(numAdditionalClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("additionalClassification["+strconv.Itoa(numAdditionalClassification)+"]", &resource.AdditionalClassification[numAdditionalClassification], optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_RelatedEntryRelationtype(numRelatedEntry int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRelation_type

	if resource == nil || numRelatedEntry >= len(resource.RelatedEntry) {
		return CodeSelect("relatedEntry["+strconv.Itoa(numRelatedEntry)+"].relationtype", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relatedEntry["+strconv.Itoa(numRelatedEntry)+"].relationtype", &resource.RelatedEntry[numRelatedEntry].Relationtype, optionsValueSet, htmlAttrs)
}
func (resource *CatalogEntry) T_RelatedEntryItem(frs []FhirResource, numRelatedEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedEntry >= len(resource.RelatedEntry) {
		return ReferenceInput(frs, "relatedEntry["+strconv.Itoa(numRelatedEntry)+"].item", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relatedEntry["+strconv.Itoa(numRelatedEntry)+"].item", &resource.RelatedEntry[numRelatedEntry].Item, htmlAttrs)
}
