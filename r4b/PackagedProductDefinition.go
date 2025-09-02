package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	PeriodDuration               *Duration         `json:"periodDuration"`
	PeriodString                 *string           `json:"periodString"`
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackageProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
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

func (resource *PackagedProductDefinition) PackagedProductDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionCharacteristic(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("characteristic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("characteristic", &resource.Characteristic[0], optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionLegalStatusOfSupplyCode(numLegalStatusOfSupply int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.LegalStatusOfSupply) >= numLegalStatusOfSupply {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Code, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionLegalStatusOfSupplyJurisdiction(numLegalStatusOfSupply int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.LegalStatusOfSupply) >= numLegalStatusOfSupply {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Jurisdiction, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackageType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Package.Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackageMaterial(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("material", &resource.Package.Material[0], optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackageAlternateMaterial(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("alternateMaterial", nil, optionsValueSet)
	}
	return CodeableConceptSelect("alternateMaterial", &resource.Package.AlternateMaterial[0], optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackageShelfLifeStorageType(numShelfLifeStorage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Package.ShelfLifeStorage) >= numShelfLifeStorage {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Package.ShelfLifeStorage[numShelfLifeStorage].Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackageShelfLifeStorageSpecialPrecautionsForStorage(numShelfLifeStorage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Package.ShelfLifeStorage) >= numShelfLifeStorage {
		return CodeableConceptSelect("specialPrecautionsForStorage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialPrecautionsForStorage", &resource.Package.ShelfLifeStorage[numShelfLifeStorage].SpecialPrecautionsForStorage[0], optionsValueSet)
}
func (resource *PackagedProductDefinition) PackagedProductDefinitionPackagePropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Package.Property) >= numProperty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Package.Property[numProperty].Type, optionsValueSet)
}
