package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinition struct {
	Id                    *string                                        `json:"id,omitempty"`
	Meta                  *Meta                                          `json:"meta,omitempty"`
	ImplicitRules         *string                                        `json:"implicitRules,omitempty"`
	Language              *string                                        `json:"language,omitempty"`
	Text                  *Narrative                                     `json:"text,omitempty"`
	Contained             []Resource                                     `json:"contained,omitempty"`
	Extension             []Extension                                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                                    `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                   `json:"identifier,omitempty"`
	Name                  *string                                        `json:"name,omitempty"`
	Type                  *CodeableConcept                               `json:"type,omitempty"`
	PackageFor            []Reference                                    `json:"packageFor,omitempty"`
	Status                *CodeableConcept                               `json:"status,omitempty"`
	StatusDate            *FhirDateTime                                  `json:"statusDate,omitempty"`
	ContainedItemQuantity []Quantity                                     `json:"containedItemQuantity,omitempty"`
	Description           *string                                        `json:"description,omitempty"`
	LegalStatusOfSupply   []PackagedProductDefinitionLegalStatusOfSupply `json:"legalStatusOfSupply,omitempty"`
	MarketingStatus       []MarketingStatus                              `json:"marketingStatus,omitempty"`
	Characteristic        []CodeableConcept                              `json:"characteristic,omitempty"`
	CopackagedIndicator   *bool                                          `json:"copackagedIndicator,omitempty"`
	Manufacturer          []Reference                                    `json:"manufacturer,omitempty"`
	Package               *PackagedProductDefinitionPackage              `json:"package,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionLegalStatusOfSupply struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackage struct {
	Id                *string                                            `json:"id,omitempty"`
	Extension         []Extension                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                        `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                                       `json:"identifier,omitempty"`
	Type              *CodeableConcept                                   `json:"type,omitempty"`
	Quantity          *int                                               `json:"quantity,omitempty"`
	Material          []CodeableConcept                                  `json:"material,omitempty"`
	AlternateMaterial []CodeableConcept                                  `json:"alternateMaterial,omitempty"`
	ShelfLifeStorage  []PackagedProductDefinitionPackageShelfLifeStorage `json:"shelfLifeStorage,omitempty"`
	Manufacturer      []Reference                                        `json:"manufacturer,omitempty"`
	Property          []PackagedProductDefinitionPackageProperty         `json:"property,omitempty"`
	ContainedItem     []PackagedProductDefinitionPackageContainedItem    `json:"containedItem,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackageShelfLifeStorage struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Type                         *CodeableConcept  `json:"type,omitempty"`
	PeriodDuration               *Duration         `json:"periodDuration,omitempty"`
	PeriodString                 *string           `json:"periodString,omitempty"`
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackageProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *FhirDate        `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackageContainedItem struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Item              CodeableReference `json:"item"`
	Amount            *Quantity         `json:"amount,omitempty"`
}

type OtherPackagedProductDefinition PackagedProductDefinition

// on convert struct to json, automatically add resourceType=PackagedProductDefinition
func (r PackagedProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPackagedProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPackagedProductDefinition: OtherPackagedProductDefinition(r),
		ResourceType:                   "PackagedProductDefinition",
	})
}
func (r PackagedProductDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PackagedProductDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PackagedProductDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PackagedProductDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageFor(numPackageFor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackageFor >= len(resource.PackageFor) {
		return ReferenceInput("packageFor["+strconv.Itoa(numPackageFor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("packageFor["+strconv.Itoa(numPackageFor)+"]", &resource.PackageFor[numPackageFor], htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_StatusDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("statusDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("statusDate", resource.StatusDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_ContainedItemQuantity(numContainedItemQuantity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContainedItemQuantity >= len(resource.ContainedItemQuantity) {
		return QuantityInput("containedItemQuantity["+strconv.Itoa(numContainedItemQuantity)+"]", nil, htmlAttrs)
	}
	return QuantityInput("containedItemQuantity["+strconv.Itoa(numContainedItemQuantity)+"]", &resource.ContainedItemQuantity[numContainedItemQuantity], htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_MarketingStatus(numMarketingStatus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMarketingStatus >= len(resource.MarketingStatus) {
		return MarketingStatusInput("marketingStatus["+strconv.Itoa(numMarketingStatus)+"]", nil, htmlAttrs)
	}
	return MarketingStatusInput("marketingStatus["+strconv.Itoa(numMarketingStatus)+"]", &resource.MarketingStatus[numMarketingStatus], htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_CopackagedIndicator(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("copackagedIndicator", nil, htmlAttrs)
	}
	return BoolInput("copackagedIndicator", resource.CopackagedIndicator, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Manufacturer(numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return ReferenceInput("manufacturer["+strconv.Itoa(numManufacturer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("manufacturer["+strconv.Itoa(numManufacturer)+"]", &resource.Manufacturer[numManufacturer], htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyCode(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("legalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].code", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Code, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyJurisdiction(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("legalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].jurisdiction", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("package.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.type", resource.Package.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageQuantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("package.quantity", nil, htmlAttrs)
	}
	return IntInput("package.quantity", resource.Package.Quantity, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageMaterial(numMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMaterial >= len(resource.Package.Material) {
		return CodeableConceptSelect("package.material["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.material["+strconv.Itoa(numMaterial)+"]", &resource.Package.Material[numMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageAlternateMaterial(numAlternateMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAlternateMaterial >= len(resource.Package.AlternateMaterial) {
		return CodeableConceptSelect("package.alternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.alternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.Package.AlternateMaterial[numAlternateMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageManufacturer(numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Package.Manufacturer) {
		return ReferenceInput("package.manufacturer["+strconv.Itoa(numManufacturer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("package.manufacturer["+strconv.Itoa(numManufacturer)+"]", &resource.Package.Manufacturer[numManufacturer], htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageType(numShelfLifeStorage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) {
		return CodeableConceptSelect("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].type", resource.Package.ShelfLifeStorage[numShelfLifeStorage].Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStoragePeriodDuration(numShelfLifeStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) {
		return DurationInput("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].periodDuration", nil, htmlAttrs)
	}
	return DurationInput("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].periodDuration", resource.Package.ShelfLifeStorage[numShelfLifeStorage].PeriodDuration, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStoragePeriodString(numShelfLifeStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) {
		return StringInput("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].periodString", nil, htmlAttrs)
	}
	return StringInput("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].periodString", resource.Package.ShelfLifeStorage[numShelfLifeStorage].PeriodString, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageSpecialPrecautionsForStorage(numShelfLifeStorage int, numSpecialPrecautionsForStorage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) || numSpecialPrecautionsForStorage >= len(resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage) {
		return CodeableConceptSelect("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].specialPrecautionsForStorage["+strconv.Itoa(numSpecialPrecautionsForStorage)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].specialPrecautionsForStorage["+strconv.Itoa(numSpecialPrecautionsForStorage)+"]", &resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage[numSpecialPrecautionsForStorage], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return CodeableConceptSelect("package.property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.property["+strconv.Itoa(numProperty)+"].type", &resource.Package.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return CodeableConceptSelect("package.property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("package.property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", resource.Package.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueQuantity(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return QuantityInput("package.property["+strconv.Itoa(numProperty)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("package.property["+strconv.Itoa(numProperty)+"].valueQuantity", resource.Package.Property[numProperty].ValueQuantity, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueDate(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return FhirDateInput("package.property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("package.property["+strconv.Itoa(numProperty)+"].valueDate", resource.Package.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return BoolInput("package.property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("package.property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Package.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueAttachment(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Package.Property) {
		return AttachmentInput("package.property["+strconv.Itoa(numProperty)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("package.property["+strconv.Itoa(numProperty)+"].valueAttachment", resource.Package.Property[numProperty].ValueAttachment, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageContainedItemItem(numContainedItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContainedItem >= len(resource.Package.ContainedItem) {
		return CodeableReferenceInput("package.containedItem["+strconv.Itoa(numContainedItem)+"].item", nil, htmlAttrs)
	}
	return CodeableReferenceInput("package.containedItem["+strconv.Itoa(numContainedItem)+"].item", &resource.Package.ContainedItem[numContainedItem].Item, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageContainedItemAmount(numContainedItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContainedItem >= len(resource.Package.ContainedItem) {
		return QuantityInput("package.containedItem["+strconv.Itoa(numContainedItem)+"].amount", nil, htmlAttrs)
	}
	return QuantityInput("package.containedItem["+strconv.Itoa(numContainedItem)+"].amount", resource.Package.ContainedItem[numContainedItem].Amount, htmlAttrs)
}
