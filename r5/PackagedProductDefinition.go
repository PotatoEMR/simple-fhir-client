package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PackagedProductDefinition
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
	CopackagedIndicator   *bool                                          `json:"copackagedIndicator,omitempty"`
	Manufacturer          []Reference                                    `json:"manufacturer,omitempty"`
	AttachedDocument      []Reference                                    `json:"attachedDocument,omitempty"`
	Packaging             *PackagedProductDefinitionPackaging            `json:"packaging,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionLegalStatusOfSupply struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackaging struct {
	Id                *string                                           `json:"id,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                                      `json:"identifier,omitempty"`
	Type              *CodeableConcept                                  `json:"type,omitempty"`
	ComponentPart     *bool                                             `json:"componentPart,omitempty"`
	Quantity          *int                                              `json:"quantity,omitempty"`
	Material          []CodeableConcept                                 `json:"material,omitempty"`
	AlternateMaterial []CodeableConcept                                 `json:"alternateMaterial,omitempty"`
	ShelfLifeStorage  []ProductShelfLife                                `json:"shelfLifeStorage,omitempty"`
	Manufacturer      []Reference                                       `json:"manufacturer,omitempty"`
	Property          []PackagedProductDefinitionPackagingProperty      `json:"property,omitempty"`
	ContainedItem     []PackagedProductDefinitionPackagingContainedItem `json:"containedItem,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackagingProperty struct {
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

// http://hl7.org/fhir/r5/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackagingContainedItem struct {
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
func (resource *PackagedProductDefinition) T_PackagingId() templ.Component {

	if resource == nil {
		return StringInput("PackagedProductDefinition.Packaging.Id", nil)
	}
	return StringInput("PackagedProductDefinition.Packaging.Id", resource.Packaging.Id)
}
func (resource *PackagedProductDefinition) T_PackagingType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("PackagedProductDefinition.Packaging.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Packaging.Type", resource.Packaging.Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackagingComponentPart() templ.Component {

	if resource == nil {
		return BoolInput("PackagedProductDefinition.Packaging.ComponentPart", nil)
	}
	return BoolInput("PackagedProductDefinition.Packaging.ComponentPart", resource.Packaging.ComponentPart)
}
func (resource *PackagedProductDefinition) T_PackagingQuantity() templ.Component {

	if resource == nil {
		return IntInput("PackagedProductDefinition.Packaging.Quantity", nil)
	}
	return IntInput("PackagedProductDefinition.Packaging.Quantity", resource.Packaging.Quantity)
}
func (resource *PackagedProductDefinition) T_PackagingMaterial(numMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Packaging.Material) >= numMaterial {
		return CodeableConceptSelect("PackagedProductDefinition.Packaging.Material["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Packaging.Material["+strconv.Itoa(numMaterial)+"]", &resource.Packaging.Material[numMaterial], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackagingAlternateMaterial(numAlternateMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Packaging.AlternateMaterial) >= numAlternateMaterial {
		return CodeableConceptSelect("PackagedProductDefinition.Packaging.AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Packaging.AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.Packaging.AlternateMaterial[numAlternateMaterial], optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Packaging.Property) >= numProperty {
		return StringInput("PackagedProductDefinition.Packaging.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.Packaging.Property["+strconv.Itoa(numProperty)+"].Id", resource.Packaging.Property[numProperty].Id)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Packaging.Property) >= numProperty {
		return CodeableConceptSelect("PackagedProductDefinition.Packaging.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PackagedProductDefinition.Packaging.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Packaging.Property[numProperty].Type, optionsValueSet)
}
func (resource *PackagedProductDefinition) T_PackagingContainedItemId(numContainedItem int) templ.Component {

	if resource == nil || len(resource.Packaging.ContainedItem) >= numContainedItem {
		return StringInput("PackagedProductDefinition.Packaging.ContainedItem["+strconv.Itoa(numContainedItem)+"].Id", nil)
	}
	return StringInput("PackagedProductDefinition.Packaging.ContainedItem["+strconv.Itoa(numContainedItem)+"].Id", resource.Packaging.ContainedItem[numContainedItem].Id)
}
