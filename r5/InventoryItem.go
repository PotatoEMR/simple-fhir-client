package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItem struct {
	Id                      *string                                `json:"id,omitempty"`
	Meta                    *Meta                                  `json:"meta,omitempty"`
	ImplicitRules           *string                                `json:"implicitRules,omitempty"`
	Language                *string                                `json:"language,omitempty"`
	Text                    *Narrative                             `json:"text,omitempty"`
	Contained               []Resource                             `json:"contained,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                           `json:"identifier,omitempty"`
	Status                  string                                 `json:"status"`
	Category                []CodeableConcept                      `json:"category,omitempty"`
	Code                    []CodeableConcept                      `json:"code,omitempty"`
	Name                    []InventoryItemName                    `json:"name,omitempty"`
	ResponsibleOrganization []InventoryItemResponsibleOrganization `json:"responsibleOrganization,omitempty"`
	Description             *InventoryItemDescription              `json:"description,omitempty"`
	InventoryStatus         []CodeableConcept                      `json:"inventoryStatus,omitempty"`
	BaseUnit                *CodeableConcept                       `json:"baseUnit,omitempty"`
	NetContent              *Quantity                              `json:"netContent,omitempty"`
	Association             []InventoryItemAssociation             `json:"association,omitempty"`
	Characteristic          []InventoryItemCharacteristic          `json:"characteristic,omitempty"`
	Instance                *InventoryItemInstance                 `json:"instance,omitempty"`
	ProductReference        *Reference                             `json:"productReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NameType          Coding      `json:"nameType"`
	Language          string      `json:"language"`
	Name              string      `json:"name"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemResponsibleOrganization struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Role              CodeableConcept `json:"role"`
	Organization      Reference       `json:"organization"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemDescription struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemAssociation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	AssociationType   CodeableConcept `json:"associationType"`
	RelatedItem       Reference       `json:"relatedItem"`
	Quantity          Ratio           `json:"quantity"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemCharacteristic struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	CharacteristicType   CodeableConcept `json:"characteristicType"`
	ValueString          string          `json:"valueString"`
	ValueInteger         int             `json:"valueInteger"`
	ValueDecimal         float64         `json:"valueDecimal"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueUrl             string          `json:"valueUrl"`
	ValueDateTime        string          `json:"valueDateTime"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueRatio           Ratio           `json:"valueRatio"`
	ValueAnnotation      Annotation      `json:"valueAnnotation"`
	ValueAddress         Address         `json:"valueAddress"`
	ValueDuration        Duration        `json:"valueDuration"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/InventoryItem
type InventoryItemInstance struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier `json:"identifier,omitempty"`
	LotNumber         *string      `json:"lotNumber,omitempty"`
	Expiry            *string      `json:"expiry,omitempty"`
	Subject           *Reference   `json:"subject,omitempty"`
	Location          *Reference   `json:"location,omitempty"`
}

type OtherInventoryItem InventoryItem

// on convert struct to json, automatically add resourceType=InventoryItem
func (r InventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInventoryItem
		ResourceType string `json:"resourceType"`
	}{
		OtherInventoryItem: OtherInventoryItem(r),
		ResourceType:       "InventoryItem",
	})
}

func (resource *InventoryItem) InventoryItemLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *InventoryItem) InventoryItemStatus() templ.Component {
	optionsValueSet := VSInventoryitem_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *InventoryItem) InventoryItemNameLanguage(numName int) templ.Component {
	optionsValueSet := VSLanguages
	currentVal := ""
	if resource != nil && len(resource.Name) >= numName {
		currentVal = resource.Name[numName].Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *InventoryItem) InventoryItemDescriptionLanguage() templ.Component {
	optionsValueSet := VSLanguages
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Description.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
