package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	StatusDate            *time.Time                                     `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate            *time.Time       `json:"valueDate,omitempty,format:'2006-01-02'"`
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
func (resource *PackagedProductDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("PackagedProductDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_StatusDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("PackagedProductDefinition.StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("PackagedProductDefinition.StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("PackagedProductDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("PackagedProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+".", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_CopackagedIndicator(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("PackagedProductDefinition.CopackagedIndicator", nil, htmlAttrs)
	}
	return BoolInput("PackagedProductDefinition.CopackagedIndicator", resource.CopackagedIndicator, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyCode(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply."+strconv.Itoa(numLegalStatusOfSupply)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply."+strconv.Itoa(numLegalStatusOfSupply)+"..Code", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Code, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyJurisdiction(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply."+strconv.Itoa(numLegalStatusOfSupply)+"..Jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.LegalStatusOfSupply."+strconv.Itoa(numLegalStatusOfSupply)+"..Jurisdiction", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Type", resource.Package.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageQuantity(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("PackagedProductDefinition.Package.Quantity", nil, htmlAttrs)
	}
	return IntInput("PackagedProductDefinition.Package.Quantity", resource.Package.Quantity, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageMaterial(numMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMaterial >= len(resource.Package.Material) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Material."+strconv.Itoa(numMaterial)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Material."+strconv.Itoa(numMaterial)+".", &resource.Package.Material[numMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageAlternateMaterial(numAlternateMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numAlternateMaterial >= len(resource.Package.AlternateMaterial) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.AlternateMaterial."+strconv.Itoa(numAlternateMaterial)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.AlternateMaterial."+strconv.Itoa(numAlternateMaterial)+".", &resource.Package.AlternateMaterial[numAlternateMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageType(numShelfLifeStorage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..Type", resource.Package.ShelfLifeStorage[numShelfLifeStorage].Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStoragePeriodString(numShelfLifeStorage int, htmlAttrs string) templ.Component {

	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) {
		return StringInput("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..PeriodString", nil, htmlAttrs)
	}
	return StringInput("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..PeriodString", resource.Package.ShelfLifeStorage[numShelfLifeStorage].PeriodString, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackageShelfLifeStorageSpecialPrecautionsForStorage(numShelfLifeStorage int, numSpecialPrecautionsForStorage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numShelfLifeStorage >= len(resource.Package.ShelfLifeStorage) || numSpecialPrecautionsForStorage >= len(resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..SpecialPrecautionsForStorage."+strconv.Itoa(numSpecialPrecautionsForStorage)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.ShelfLifeStorage."+strconv.Itoa(numShelfLifeStorage)+"..SpecialPrecautionsForStorage."+strconv.Itoa(numSpecialPrecautionsForStorage)+".", &resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage[numSpecialPrecautionsForStorage], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Package.Property) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Package.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Package.Property) {
		return CodeableConceptSelect("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", resource.Package.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueDate(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Package.Property) {
		return DateInput("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueDate", resource.Package.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagePropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Package.Property) {
		return BoolInput("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("PackagedProductDefinition.Package.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", resource.Package.Property[numProperty].ValueBoolean, htmlAttrs)
}
