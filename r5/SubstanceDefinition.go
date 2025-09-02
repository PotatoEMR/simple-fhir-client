package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
type SubstanceDefinition struct {
	Id                   *string                               `json:"id,omitempty"`
	Meta                 *Meta                                 `json:"meta,omitempty"`
	ImplicitRules        *string                               `json:"implicitRules,omitempty"`
	Language             *string                               `json:"language,omitempty"`
	Text                 *Narrative                            `json:"text,omitempty"`
	Contained            []Resource                            `json:"contained,omitempty"`
	Extension            []Extension                           `json:"extension,omitempty"`
	ModifierExtension    []Extension                           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                          `json:"identifier,omitempty"`
	Version              *string                               `json:"version,omitempty"`
	Status               *CodeableConcept                      `json:"status,omitempty"`
	Classification       []CodeableConcept                     `json:"classification,omitempty"`
	Domain               *CodeableConcept                      `json:"domain,omitempty"`
	Grade                []CodeableConcept                     `json:"grade,omitempty"`
	Description          *string                               `json:"description,omitempty"`
	InformationSource    []Reference                           `json:"informationSource,omitempty"`
	Note                 []Annotation                          `json:"note,omitempty"`
	Manufacturer         []Reference                           `json:"manufacturer,omitempty"`
	Supplier             []Reference                           `json:"supplier,omitempty"`
	Moiety               []SubstanceDefinitionMoiety           `json:"moiety,omitempty"`
	Characterization     []SubstanceDefinitionCharacterization `json:"characterization,omitempty"`
	Property             []SubstanceDefinitionProperty         `json:"property,omitempty"`
	ReferenceInformation *Reference                            `json:"referenceInformation,omitempty"`
	MolecularWeight      []SubstanceDefinitionMolecularWeight  `json:"molecularWeight,omitempty"`
	Structure            *SubstanceDefinitionStructure         `json:"structure,omitempty"`
	Code                 []SubstanceDefinitionCode             `json:"code,omitempty"`
	Name                 []SubstanceDefinitionName             `json:"name,omitempty"`
	Relationship         []SubstanceDefinitionRelationship     `json:"relationship,omitempty"`
	NucleicAcid          *Reference                            `json:"nucleicAcid,omitempty"`
	Polymer              *Reference                            `json:"polymer,omitempty"`
	Protein              *Reference                            `json:"protein,omitempty"`
	SourceMaterial       *SubstanceDefinitionSourceMaterial    `json:"sourceMaterial,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionCharacterization struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Technique         *CodeableConcept `json:"technique,omitempty"`
	Form              *CodeableConcept `json:"form,omitempty"`
	Description       *string          `json:"description,omitempty"`
	File              []Attachment     `json:"file,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionMolecularWeight struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            Quantity         `json:"amount"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionStructureRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Format            *CodeableConcept `json:"format,omitempty"`
	Document          *Reference       `json:"document,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionNameOfficial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceDefinition
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

func (resource *SubstanceDefinition) SubstanceDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionClassification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classification", &resource.Classification[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionDomain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionGrade(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("grade", nil, optionsValueSet)
	}
	return CodeableConceptSelect("grade", &resource.Grade[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMoietyRole(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Moiety[numMoiety].Role, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMoietyStereochemistry(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMoietyOpticalActivity(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("opticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("opticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMoietyMeasurementType(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("measurementType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("measurementType", resource.Moiety[numMoiety].MeasurementType, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionCharacterizationTechnique(numCharacterization int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Characterization) >= numCharacterization {
		return CodeableConceptSelect("technique", nil, optionsValueSet)
	}
	return CodeableConceptSelect("technique", resource.Characterization[numCharacterization].Technique, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionCharacterizationForm(numCharacterization int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Characterization) >= numCharacterization {
		return CodeableConceptSelect("form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("form", resource.Characterization[numCharacterization].Form, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionPropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMolecularWeightMethod(numMolecularWeight int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.MolecularWeight) >= numMolecularWeight {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.MolecularWeight[numMolecularWeight].Method, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionMolecularWeightType(numMolecularWeight int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.MolecularWeight) >= numMolecularWeight {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.MolecularWeight[numMolecularWeight].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStructureStereochemistry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("stereochemistry", resource.Structure.Stereochemistry, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStructureOpticalActivity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("opticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("opticalActivity", resource.Structure.OpticalActivity, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStructureTechnique(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("technique", nil, optionsValueSet)
	}
	return CodeableConceptSelect("technique", &resource.Structure.Technique[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStructureRepresentationType(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionStructureRepresentationFormat(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("format", nil, optionsValueSet)
	}
	return CodeableConceptSelect("format", resource.Structure.Representation[numRepresentation].Format, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionCodeCode(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Code) >= numCode {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code[numCode].Code, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionCodeStatus(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Code) >= numCode {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Code[numCode].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameStatus(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Name[numName].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameLanguage(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.Name[numName].Language[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameDomain(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", &resource.Name[numName].Domain[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameJurisdiction(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Name[numName].Jurisdiction[0], optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("authority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionNameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionRelationshipType(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Relationship[numRelationship].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionRelationshipComparator(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("comparator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("comparator", resource.Relationship[numRelationship].Comparator, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionSourceMaterialType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.SourceMaterial.Type, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionSourceMaterialGenus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("genus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("genus", resource.SourceMaterial.Genus, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionSourceMaterialSpecies(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("species", resource.SourceMaterial.Species, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionSourceMaterialPart(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("part", nil, optionsValueSet)
	}
	return CodeableConceptSelect("part", resource.SourceMaterial.Part, optionsValueSet)
}
func (resource *SubstanceDefinition) SubstanceDefinitionSourceMaterialCountryOfOrigin(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("countryOfOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("countryOfOrigin", &resource.SourceMaterial.CountryOfOrigin[0], optionsValueSet)
}
