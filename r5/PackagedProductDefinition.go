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
	StatusDate            *time.Time                                     `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate            *time.Time       `json:"valueDate,omitempty,format:'2006-01-02'"`
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
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_StatusDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_CopackagedIndicator(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("CopackagedIndicator", nil, htmlAttrs)
	}
	return BoolInput("CopackagedIndicator", resource.CopackagedIndicator, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyCode(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"]Code", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Code, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_LegalStatusOfSupplyJurisdiction(numLegalStatusOfSupply int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLegalStatusOfSupply >= len(resource.LegalStatusOfSupply) {
		return CodeableConceptSelect("LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"]Jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("LegalStatusOfSupply["+strconv.Itoa(numLegalStatusOfSupply)+"]Jurisdiction", resource.LegalStatusOfSupply[numLegalStatusOfSupply].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PackagingType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagingType", resource.Packaging.Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingComponentPart(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("PackagingComponentPart", nil, htmlAttrs)
	}
	return BoolInput("PackagingComponentPart", resource.Packaging.ComponentPart, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingQuantity(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("PackagingQuantity", nil, htmlAttrs)
	}
	return IntInput("PackagingQuantity", resource.Packaging.Quantity, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingMaterial(numMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Packaging.Material) {
		return CodeableConceptSelect("PackagingMaterial["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagingMaterial["+strconv.Itoa(numMaterial)+"]", &resource.Packaging.Material[numMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingAlternateMaterial(numAlternateMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAlternateMaterial >= len(resource.Packaging.AlternateMaterial) {
		return CodeableConceptSelect("PackagingAlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagingAlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.Packaging.AlternateMaterial[numAlternateMaterial], optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Packaging.Property) {
		return CodeableConceptSelect("PackagingProperty["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagingProperty["+strconv.Itoa(numProperty)+"].Type", &resource.Packaging.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Packaging.Property) {
		return CodeableConceptSelect("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", resource.Packaging.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyValueDate(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Packaging.Property) {
		return DateInput("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueDate", nil, htmlAttrs)
	}
	return DateInput("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueDate", resource.Packaging.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *PackagedProductDefinition) T_PackagingPropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Packaging.Property) {
		return BoolInput("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("PackagingProperty["+strconv.Itoa(numProperty)+"].ValueBoolean", resource.Packaging.Property[numProperty].ValueBoolean, htmlAttrs)
}
