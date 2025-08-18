//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinition struct {
	Id                *string                              `json:"id,omitempty"`
	Meta              *Meta                                `json:"meta,omitempty"`
	ImplicitRules     *string                              `json:"implicitRules,omitempty"`
	Language          *string                              `json:"language,omitempty"`
	Text              *Narrative                           `json:"text,omitempty"`
	Contained         []Resource                           `json:"contained,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                         `json:"identifier,omitempty"`
	Version           *string                              `json:"version,omitempty"`
	Status            *CodeableConcept                     `json:"status,omitempty"`
	Classification    []CodeableConcept                    `json:"classification,omitempty"`
	Domain            *CodeableConcept                     `json:"domain,omitempty"`
	Grade             []CodeableConcept                    `json:"grade,omitempty"`
	Description       *string                              `json:"description,omitempty"`
	InformationSource []Reference                          `json:"informationSource,omitempty"`
	Note              []Annotation                         `json:"note,omitempty"`
	Manufacturer      []Reference                          `json:"manufacturer,omitempty"`
	Supplier          []Reference                          `json:"supplier,omitempty"`
	Moiety            []SubstanceDefinitionMoiety          `json:"moiety,omitempty"`
	Property          []SubstanceDefinitionProperty        `json:"property,omitempty"`
	MolecularWeight   []SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`
	Structure         *SubstanceDefinitionStructure        `json:"structure,omitempty"`
	Code              []SubstanceDefinitionCode            `json:"code,omitempty"`
	Name              []SubstanceDefinitionName            `json:"name,omitempty"`
	Relationship      []SubstanceDefinitionRelationship    `json:"relationship,omitempty"`
	SourceMaterial    *SubstanceDefinitionSourceMaterial   `json:"sourceMaterial,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionMoiety struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Stereochemistry   *CodeableConcept `json:"stereochemistry,omitempty"`
	OpticalActivity   *CodeableConcept `json:"opticalActivity,omitempty"`
	MolecularFormula  *string          `json:"molecularFormula,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity"`
	AmountString      *string          `json:"amountString"`
	MeasurementType   *CodeableConcept `json:"measurementType,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionProperty struct {
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

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionMolecularWeight struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            Quantity         `json:"amount"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionStructure struct {
	Id                       *string                                      `json:"id,omitempty"`
	Extension                []Extension                                  `json:"extension,omitempty"`
	ModifierExtension        []Extension                                  `json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                             `json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                             `json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                      `json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                      `json:"molecularFormulaByMoiety,omitempty"`
	Technique                []CodeableConcept                            `json:"technique,omitempty"`
	SourceDocument           []Reference                                  `json:"sourceDocument,omitempty"`
	Representation           []SubstanceDefinitionStructureRepresentation `json:"representation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionStructureRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Format            *CodeableConcept `json:"format,omitempty"`
	Document          *Reference       `json:"document,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionCode struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	StatusDate        *string          `json:"statusDate,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionName struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Name              string                            `json:"name"`
	Type              *CodeableConcept                  `json:"type,omitempty"`
	Status            *CodeableConcept                  `json:"status,omitempty"`
	Preferred         *bool                             `json:"preferred,omitempty"`
	Language          []CodeableConcept                 `json:"language,omitempty"`
	Domain            []CodeableConcept                 `json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                 `json:"jurisdiction,omitempty"`
	Official          []SubstanceDefinitionNameOfficial `json:"official,omitempty"`
	Source            []Reference                       `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionNameOfficial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionRelationship struct {
	Id                                 *string          `json:"id,omitempty"`
	Extension                          []Extension      `json:"extension,omitempty"`
	ModifierExtension                  []Extension      `json:"modifierExtension,omitempty"`
	SubstanceDefinitionReference       *Reference       `json:"substanceDefinitionReference"`
	SubstanceDefinitionCodeableConcept *CodeableConcept `json:"substanceDefinitionCodeableConcept"`
	Type                               CodeableConcept  `json:"type"`
	IsDefining                         *bool            `json:"isDefining,omitempty"`
	AmountQuantity                     *Quantity        `json:"amountQuantity"`
	AmountRatio                        *Ratio           `json:"amountRatio"`
	AmountString                       *string          `json:"amountString"`
	RatioHighLimitAmount               *Ratio           `json:"ratioHighLimitAmount,omitempty"`
	Comparator                         *CodeableConcept `json:"comparator,omitempty"`
	Source                             []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionSourceMaterial struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Genus             *CodeableConcept  `json:"genus,omitempty"`
	Species           *CodeableConcept  `json:"species,omitempty"`
	Part              *CodeableConcept  `json:"part,omitempty"`
	CountryOfOrigin   []CodeableConcept `json:"countryOfOrigin,omitempty"`
}

type OtherSubstanceDefinition SubstanceDefinition

// on convert struct to json, automatically add resourceType=SubstanceDefinition
func (r SubstanceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceDefinition: OtherSubstanceDefinition(r),
		ResourceType:             "SubstanceDefinition",
	})
}
