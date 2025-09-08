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
	ValueDateTime        time.Time       `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
	Expiry            *time.Time   `json:"expiry,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r InventoryItem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "InventoryItem/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "InventoryItem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *InventoryItem) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSInventoryitem_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_InventoryStatus(numInventoryStatus int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInventoryStatus >= len(resource.InventoryStatus) {
		return CodeableConceptSelect("InventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("InventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", &resource.InventoryStatus[numInventoryStatus], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_BaseUnit(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("BaseUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BaseUnit", resource.BaseUnit, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_NameNameType(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodingSelect("Name["+strconv.Itoa(numName)+"]NameType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Name["+strconv.Itoa(numName)+"]NameType", &resource.Name[numName].NameType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_NameName(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("Name["+strconv.Itoa(numName)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Name["+strconv.Itoa(numName)+"]Name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *InventoryItem) T_ResponsibleOrganizationRole(numResponsibleOrganization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numResponsibleOrganization >= len(resource.ResponsibleOrganization) {
		return CodeableConceptSelect("ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResponsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"]Role", &resource.ResponsibleOrganization[numResponsibleOrganization].Role, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_DescriptionDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DescriptionDescription", nil, htmlAttrs)
	}
	return StringInput("DescriptionDescription", resource.Description.Description, htmlAttrs)
}
func (resource *InventoryItem) T_AssociationAssociationType(numAssociation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAssociation >= len(resource.Association) {
		return CodeableConceptSelect("Association["+strconv.Itoa(numAssociation)+"]AssociationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Association["+strconv.Itoa(numAssociation)+"]AssociationType", &resource.Association[numAssociation].AssociationType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicCharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]CharacteristicType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]CharacteristicType", &resource.Characteristic[numCharacteristic].CharacteristicType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueString(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueString", &resource.Characteristic[numCharacteristic].ValueString, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueInteger(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueInteger", &resource.Characteristic[numCharacteristic].ValueInteger, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueDecimal(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return Float64Input("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueDecimal", &resource.Characteristic[numCharacteristic].ValueDecimal, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueUrl(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueUrl", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueUrl", &resource.Characteristic[numCharacteristic].ValueUrl, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueDateTime(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueDateTime", &resource.Characteristic[numCharacteristic].ValueDateTime, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueAnnotation(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueAnnotation", &resource.Characteristic[numCharacteristic].ValueAnnotation, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceLotNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InstanceLotNumber", nil, htmlAttrs)
	}
	return StringInput("InstanceLotNumber", resource.Instance.LotNumber, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceExpiry(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("InstanceExpiry", nil, htmlAttrs)
	}
	return DateTimeInput("InstanceExpiry", resource.Instance.Expiry, htmlAttrs)
}
