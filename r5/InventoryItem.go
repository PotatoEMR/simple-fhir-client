package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *InventoryItem) T_Id() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Id", nil)
	}
	return StringInput("InventoryItem.Id", resource.Id)
}
func (resource *InventoryItem) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.ImplicitRules", nil)
	}
	return StringInput("InventoryItem.ImplicitRules", resource.ImplicitRules)
}
func (resource *InventoryItem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("InventoryItem.Language", nil, optionsValueSet)
	}
	return CodeSelect("InventoryItem.Language", resource.Language, optionsValueSet)
}
func (resource *InventoryItem) T_Status() templ.Component {
	optionsValueSet := VSInventoryitem_status

	if resource == nil {
		return CodeSelect("InventoryItem.Status", nil, optionsValueSet)
	}
	return CodeSelect("InventoryItem.Status", &resource.Status, optionsValueSet)
}
func (resource *InventoryItem) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("InventoryItem.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *InventoryItem) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("InventoryItem.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *InventoryItem) T_InventoryStatus(numInventoryStatus int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.InventoryStatus) >= numInventoryStatus {
		return CodeableConceptSelect("InventoryItem.InventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.InventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", &resource.InventoryStatus[numInventoryStatus], optionsValueSet)
}
func (resource *InventoryItem) T_BaseUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("InventoryItem.BaseUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.BaseUnit", resource.BaseUnit, optionsValueSet)
}
func (resource *InventoryItem) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("InventoryItem.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("InventoryItem.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *InventoryItem) T_NameNameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodingSelect("InventoryItem.Name["+strconv.Itoa(numName)+"].NameType", nil, optionsValueSet)
	}
	return CodingSelect("InventoryItem.Name["+strconv.Itoa(numName)+"].NameType", &resource.Name[numName].NameType, optionsValueSet)
}
func (resource *InventoryItem) T_NameLanguage(numName int) templ.Component {
	optionsValueSet := VSLanguages

	if resource == nil || len(resource.Name) >= numName {
		return CodeSelect("InventoryItem.Name["+strconv.Itoa(numName)+"].Language", nil, optionsValueSet)
	}
	return CodeSelect("InventoryItem.Name["+strconv.Itoa(numName)+"].Language", &resource.Name[numName].Language, optionsValueSet)
}
func (resource *InventoryItem) T_NameName(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("InventoryItem.Name["+strconv.Itoa(numName)+"].Name", nil)
	}
	return StringInput("InventoryItem.Name["+strconv.Itoa(numName)+"].Name", &resource.Name[numName].Name)
}
func (resource *InventoryItem) T_ResponsibleOrganizationId(numResponsibleOrganization int) templ.Component {

	if resource == nil || len(resource.ResponsibleOrganization) >= numResponsibleOrganization {
		return StringInput("InventoryItem.ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].Id", nil)
	}
	return StringInput("InventoryItem.ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].Id", resource.ResponsibleOrganization[numResponsibleOrganization].Id)
}
func (resource *InventoryItem) T_ResponsibleOrganizationRole(numResponsibleOrganization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ResponsibleOrganization) >= numResponsibleOrganization {
		return CodeableConceptSelect("InventoryItem.ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].Role", &resource.ResponsibleOrganization[numResponsibleOrganization].Role, optionsValueSet)
}
func (resource *InventoryItem) T_DescriptionId() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Description.Id", nil)
	}
	return StringInput("InventoryItem.Description.Id", resource.Description.Id)
}
func (resource *InventoryItem) T_DescriptionLanguage() templ.Component {
	optionsValueSet := VSLanguages

	if resource == nil {
		return CodeSelect("InventoryItem.Description.Language", nil, optionsValueSet)
	}
	return CodeSelect("InventoryItem.Description.Language", resource.Description.Language, optionsValueSet)
}
func (resource *InventoryItem) T_DescriptionDescription() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Description.Description", nil)
	}
	return StringInput("InventoryItem.Description.Description", resource.Description.Description)
}
func (resource *InventoryItem) T_AssociationId(numAssociation int) templ.Component {

	if resource == nil || len(resource.Association) >= numAssociation {
		return StringInput("InventoryItem.Association["+strconv.Itoa(numAssociation)+"].Id", nil)
	}
	return StringInput("InventoryItem.Association["+strconv.Itoa(numAssociation)+"].Id", resource.Association[numAssociation].Id)
}
func (resource *InventoryItem) T_AssociationAssociationType(numAssociation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Association) >= numAssociation {
		return CodeableConceptSelect("InventoryItem.Association["+strconv.Itoa(numAssociation)+"].AssociationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.Association["+strconv.Itoa(numAssociation)+"].AssociationType", &resource.Association[numAssociation].AssociationType, optionsValueSet)
}
func (resource *InventoryItem) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("InventoryItem.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("InventoryItem.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *InventoryItem) T_CharacteristicCharacteristicType(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("InventoryItem.Characteristic["+strconv.Itoa(numCharacteristic)+"].CharacteristicType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("InventoryItem.Characteristic["+strconv.Itoa(numCharacteristic)+"].CharacteristicType", &resource.Characteristic[numCharacteristic].CharacteristicType, optionsValueSet)
}
func (resource *InventoryItem) T_InstanceId() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Instance.Id", nil)
	}
	return StringInput("InventoryItem.Instance.Id", resource.Instance.Id)
}
func (resource *InventoryItem) T_InstanceLotNumber() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Instance.LotNumber", nil)
	}
	return StringInput("InventoryItem.Instance.LotNumber", resource.Instance.LotNumber)
}
func (resource *InventoryItem) T_InstanceExpiry() templ.Component {

	if resource == nil {
		return StringInput("InventoryItem.Instance.Expiry", nil)
	}
	return StringInput("InventoryItem.Instance.Expiry", resource.Instance.Expiry)
}
