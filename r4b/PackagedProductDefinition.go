package r4b

//generated with command go run ./bultaoreune -nodownload
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
	StatusDate            *string                                        `json:"statusDate,omitempty"`
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
	ValueDate            *string          `json:"valueDate,omitempty"`
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

func (resource *PackagedProductDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Id", nil)
	}
	return StringInput("PackagedProductDefinition.Id", resource.Id)
}
func (resource *PackagedProductDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.ImplicitRules", nil)
	}
	return StringInput("PackagedProductDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *PackagedProductDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PackagedProductDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("PackagedProductDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Name", nil)
	}
	return StringInput("PackagedProductDefinition.Name", resource.Name)
}
func (resource *PackagedProductDefinition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Type", resource.Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Status", resource.Status, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.StatusDate", nil)
	}
	return StringInput("PackagedProductDefinition.StatusDate", resource.StatusDate)
}
func (resource *PackagedProductDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Description", nil)
	}
	return StringInput("PackagedProductDefinition.Description", resource.Description)
}
func (resource *PackagedProductDefinition) T_Characteristic(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("PackagedProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_CopackagedIndicator() templ.Component {

	if resource == nil {
		return BoolInput("PackagedProductDefinition.CopackagedIndicator", nil)
	}
	return BoolInput("PackagedProductDefinition.CopackagedIndicator", resource.CopackagedIndicator)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyId(numLegalStatusOfSupply int) templ.Component {

	if resource == nil || len(resource.LegalStatusOfSupply) >= numLegalStatusOfSupply {
		return StringInput("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Id", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Id)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyCode(numLegalStatusOfSupply int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LegalStatusOfSupply) >= numLegalStatusOfSupply {
		return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Code", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Code, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyJurisdiction(numLegalStatusOfSupply int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LegalStatusOfSupply) >= numLegalStatusOfSupply {
		return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"].Jurisdiction", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Jurisdiction, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageId() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Package.Id", nil)
	}
	return StringInput("PackagedProductDefinition.Package.Id", resource.Package.Id)
}
func (resource *PackagedProductDefinition) T_PackageType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Type", resource.Package.Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageQuantity() templ.Component {

	if resource == nil {
		return IntInput("PackagedProductDefinition.Package.Quantity", nil)
	}
	return IntInput("PackagedProductDefinition.Package.Quantity", resource.Package.Quantity)
}
func (resource *PackagedProductDefinition) T_PackageMaterial(numMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Package.Material) >= numMaterial {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Material["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Material["+strconv.Itoa(numMaterial)+"]", &resource.Package.Material[numMaterial], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageAlternateMaterial(numAlternateMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Package.AlternateMaterial) >= numAlternateMaterial {
		return CodeableConceptSelect("PackagedProductDefinition.Package.AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.Package.AlternateMaterial[numAlternateMaterial], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageId(numShelfLifeStorage int) templ.Component {

	if resource == nil || len(resource.Package.ShelfLifeStorage) >= numShelfLifeStorage {
		return StringInput("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].Id", resource.Package.ShelfLifeStorage[numShelfLifeStorage].Id)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageType(numShelfLifeStorage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Package.ShelfLifeStorage) >= numShelfLifeStorage {
		return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].Type", resource.Package.ShelfLifeStorage[numShelfLifeStorage].Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageSpecialPrecautionsForStorage(numShelfLifeStorage int, numSpecialPrecautionsForStorage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Package.ShelfLifeStorage) >= numShelfLifeStorage || len(resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage) >= numSpecialPrecautionsForStorage {
		return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].SpecialPrecautionsForStorage["+strconv.Itoa(numSpecialPrecautionsForStorage)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"].SpecialPrecautionsForStorage["+strconv.Itoa(numSpecialPrecautionsForStorage)+"]", &resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage[numSpecialPrecautionsForStorage], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackagePropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Package.Property) >= numProperty {
		return StringInput("PackagedProductDefinition.Package.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.Package.Property["+strconv.Itoa(numProperty)+"].Id", resource.Package.Property[numProperty].Id)
}
func (resource *PackagedProductDefinition) T_PackagePropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Package.Property) >= numProperty {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Package.Property[numProperty].Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackageContainedItemId(numContainedItem int) templ.Component {

	if resource == nil || len(resource.Package.ContainedItem) >= numContainedItem {
		return StringInput("PackagedProductDefinition.Package.ContainedItem["+strconv.Itoa(numContainedItem)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.Package.ContainedItem["+strconv.Itoa(numContainedItem)+"].Id", resource.Package.ContainedItem[numContainedItem].Id)
}
