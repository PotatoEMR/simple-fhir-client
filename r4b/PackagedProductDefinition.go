package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
