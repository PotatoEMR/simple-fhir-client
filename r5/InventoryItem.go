package r5

//generated with command go run ./bultaoreune
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
	ValueDateTime        FhirDateTime    `json:"valueDateTime"`
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
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Identifier        []Identifier  `json:"identifier,omitempty"`
	LotNumber         *string       `json:"lotNumber,omitempty"`
	Expiry            *FhirDateTime `json:"expiry,omitempty"`
	Subject           *Reference    `json:"subject,omitempty"`
	Location          *Reference    `json:"location,omitempty"`
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
func (resource *InventoryItem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSInventoryitem_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_InventoryStatus(numInventoryStatus int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInventoryStatus >= len(resource.InventoryStatus) {
		return CodeableConceptSelect("inventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("inventoryStatus["+strconv.Itoa(numInventoryStatus)+"]", &resource.InventoryStatus[numInventoryStatus], optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_BaseUnit(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("baseUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("baseUnit", resource.BaseUnit, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_NetContent(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("netContent", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("netContent", resource.NetContent, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_ProductReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "productReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "productReference", resource.ProductReference, htmlAttrs)
}
func (resource *InventoryItem) T_NameNameType(numName int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodingSelect("name["+strconv.Itoa(numName)+"].nameType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("name["+strconv.Itoa(numName)+"].nameType", &resource.Name[numName].NameType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_NameName(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].name", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *InventoryItem) T_ResponsibleOrganizationRole(numResponsibleOrganization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResponsibleOrganization >= len(resource.ResponsibleOrganization) {
		return CodeableConceptSelect("responsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("responsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].role", &resource.ResponsibleOrganization[numResponsibleOrganization].Role, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_ResponsibleOrganizationOrganization(frs []FhirResource, numResponsibleOrganization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResponsibleOrganization >= len(resource.ResponsibleOrganization) {
		return ReferenceInput(frs, "responsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "responsibleOrganization["+strconv.Itoa(numResponsibleOrganization)+"].organization", &resource.ResponsibleOrganization[numResponsibleOrganization].Organization, htmlAttrs)
}
func (resource *InventoryItem) T_DescriptionDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description.description", nil, htmlAttrs)
	}
	return StringInput("description.description", resource.Description.Description, htmlAttrs)
}
func (resource *InventoryItem) T_AssociationAssociationType(numAssociation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAssociation >= len(resource.Association) {
		return CodeableConceptSelect("association["+strconv.Itoa(numAssociation)+"].associationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("association["+strconv.Itoa(numAssociation)+"].associationType", &resource.Association[numAssociation].AssociationType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_AssociationRelatedItem(frs []FhirResource, numAssociation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAssociation >= len(resource.Association) {
		return ReferenceInput(frs, "association["+strconv.Itoa(numAssociation)+"].relatedItem", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "association["+strconv.Itoa(numAssociation)+"].relatedItem", &resource.Association[numAssociation].RelatedItem, htmlAttrs)
}
func (resource *InventoryItem) T_AssociationQuantity(numAssociation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAssociation >= len(resource.Association) {
		return RatioInput("association["+strconv.Itoa(numAssociation)+"].quantity", nil, htmlAttrs)
	}
	return RatioInput("association["+strconv.Itoa(numAssociation)+"].quantity", &resource.Association[numAssociation].Quantity, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicCharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].characteristicType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].characteristicType", &resource.Characteristic[numCharacteristic].CharacteristicType, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueString(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueString", &resource.Characteristic[numCharacteristic].ValueString, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueInteger(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueInteger", &resource.Characteristic[numCharacteristic].ValueInteger, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueDecimal(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return Float64Input("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDecimal", &resource.Characteristic[numCharacteristic].ValueDecimal, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueUrl(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueUrl", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueUrl", &resource.Characteristic[numCharacteristic].ValueUrl, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueDateTime(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDateTime", &resource.Characteristic[numCharacteristic].ValueDateTime, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", &resource.Characteristic[numCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRange", &resource.Characteristic[numCharacteristic].ValueRange, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueRatio(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RatioInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRatio", &resource.Characteristic[numCharacteristic].ValueRatio, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueAnnotation(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAnnotation", nil, htmlAttrs)
	}
	return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAnnotation", &resource.Characteristic[numCharacteristic].ValueAnnotation, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueAddress(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return AddressInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAddress", nil, htmlAttrs)
	}
	return AddressInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAddress", &resource.Characteristic[numCharacteristic].ValueAddress, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueDuration(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDuration", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDuration", &resource.Characteristic[numCharacteristic].ValueDuration, htmlAttrs)
}
func (resource *InventoryItem) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceLotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instance.lotNumber", nil, htmlAttrs)
	}
	return StringInput("instance.lotNumber", resource.Instance.LotNumber, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceExpiry(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("instance.expiry", nil, htmlAttrs)
	}
	return FhirDateTimeInput("instance.expiry", resource.Instance.Expiry, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceSubject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "instance.subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "instance.subject", resource.Instance.Subject, htmlAttrs)
}
func (resource *InventoryItem) T_InstanceLocation(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "instance.location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "instance.location", resource.Instance.Location, htmlAttrs)
}
