package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	AmountQuantity    *Quantity        `json:"amountQuantity,omitempty"`
	AmountString      *string          `json:"amountString,omitempty"`
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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
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
	SubstanceDefinitionReference       *Reference       `json:"substanceDefinitionReference,omitempty"`
	SubstanceDefinitionCodeableConcept *CodeableConcept `json:"substanceDefinitionCodeableConcept,omitempty"`
	Type                               CodeableConcept  `json:"type"`
	IsDefining                         *bool            `json:"isDefining,omitempty"`
	AmountQuantity                     *Quantity        `json:"amountQuantity,omitempty"`
	AmountRatio                        *Ratio           `json:"amountRatio,omitempty"`
	AmountString                       *string          `json:"amountString,omitempty"`
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

func (resource *SubstanceDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Id", nil)
	}
	return StringInput("SubstanceDefinition.Id", resource.Id)
}
func (resource *SubstanceDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.ImplicitRules", nil)
	}
	return StringInput("SubstanceDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Version", nil)
	}
	return StringInput("SubstanceDefinition.Version", resource.Version)
}
func (resource *SubstanceDefinition) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Status", resource.Status, optionsValueSet)
}
func (resource *SubstanceDefinition) T_Classification(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("SubstanceDefinition.Classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet)
}
func (resource *SubstanceDefinition) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Domain", resource.Domain, optionsValueSet)
}
func (resource *SubstanceDefinition) T_Grade(numGrade int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Grade) >= numGrade {
		return CodeableConceptSelect("SubstanceDefinition.Grade["+strconv.Itoa(numGrade)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Grade["+strconv.Itoa(numGrade)+"]", &resource.Grade[numGrade], optionsValueSet)
}
func (resource *SubstanceDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Description", nil)
	}
	return StringInput("SubstanceDefinition.Description", resource.Description)
}
func (resource *SubstanceDefinition) T_MoietyId(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Id", resource.Moiety[numMoiety].Id)
}
func (resource *SubstanceDefinition) T_MoietyRole(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Role", resource.Moiety[numMoiety].Role, optionsValueSet)
}
func (resource *SubstanceDefinition) T_MoietyName(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Name", nil)
	}
	return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Name", resource.Moiety[numMoiety].Name)
}
func (resource *SubstanceDefinition) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet)
}
func (resource *SubstanceDefinition) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet)
}
func (resource *SubstanceDefinition) T_MoietyMolecularFormula(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", nil)
	}
	return StringInput("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", resource.Moiety[numMoiety].MolecularFormula)
}
func (resource *SubstanceDefinition) T_MoietyMeasurementType(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].MeasurementType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety["+strconv.Itoa(numMoiety)+"].MeasurementType", resource.Moiety[numMoiety].MeasurementType, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CharacterizationId(numCharacterization int) templ.Component {

	if resource == nil || len(resource.Characterization) >= numCharacterization {
		return StringInput("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Id", resource.Characterization[numCharacterization].Id)
}
func (resource *SubstanceDefinition) T_CharacterizationTechnique(numCharacterization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characterization) >= numCharacterization {
		return CodeableConceptSelect("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Technique", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Technique", resource.Characterization[numCharacterization].Technique, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CharacterizationForm(numCharacterization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characterization) >= numCharacterization {
		return CodeableConceptSelect("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Form", resource.Characterization[numCharacterization].Form, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CharacterizationDescription(numCharacterization int) templ.Component {

	if resource == nil || len(resource.Characterization) >= numCharacterization {
		return StringInput("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Description", nil)
	}
	return StringInput("SubstanceDefinition.Characterization["+strconv.Itoa(numCharacterization)+"].Description", resource.Characterization[numCharacterization].Description)
}
func (resource *SubstanceDefinition) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("SubstanceDefinition.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *SubstanceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("SubstanceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_MolecularWeightId(numMolecularWeight int) templ.Component {

	if resource == nil || len(resource.MolecularWeight) >= numMolecularWeight {
		return StringInput("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Id", resource.MolecularWeight[numMolecularWeight].Id)
}
func (resource *SubstanceDefinition) T_MolecularWeightMethod(numMolecularWeight int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MolecularWeight) >= numMolecularWeight {
		return CodeableConceptSelect("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Method", resource.MolecularWeight[numMolecularWeight].Method, optionsValueSet)
}
func (resource *SubstanceDefinition) T_MolecularWeightType(numMolecularWeight int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MolecularWeight) >= numMolecularWeight {
		return CodeableConceptSelect("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.MolecularWeight["+strconv.Itoa(numMolecularWeight)+"].Type", resource.MolecularWeight[numMolecularWeight].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_StructureId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Structure.Id", nil)
	}
	return StringInput("SubstanceDefinition.Structure.Id", resource.Structure.Id)
}
func (resource *SubstanceDefinition) T_StructureStereochemistry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Stereochemistry", resource.Structure.Stereochemistry, optionsValueSet)
}
func (resource *SubstanceDefinition) T_StructureOpticalActivity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Structure.OpticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.OpticalActivity", resource.Structure.OpticalActivity, optionsValueSet)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormula() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Structure.MolecularFormula", nil)
	}
	return StringInput("SubstanceDefinition.Structure.MolecularFormula", resource.Structure.MolecularFormula)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormulaByMoiety() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Structure.MolecularFormulaByMoiety", nil)
	}
	return StringInput("SubstanceDefinition.Structure.MolecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety)
}
func (resource *SubstanceDefinition) T_StructureTechnique(numTechnique int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Technique) >= numTechnique {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Technique["+strconv.Itoa(numTechnique)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Technique["+strconv.Itoa(numTechnique)+"]", &resource.Structure.Technique[numTechnique], optionsValueSet)
}
func (resource *SubstanceDefinition) T_StructureRepresentationId(numRepresentation int) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return StringInput("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Id", resource.Structure.Representation[numRepresentation].Id)
}
func (resource *SubstanceDefinition) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_StructureRepresentationRepresentation(numRepresentation int) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return StringInput("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", nil)
	}
	return StringInput("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", resource.Structure.Representation[numRepresentation].Representation)
}
func (resource *SubstanceDefinition) T_StructureRepresentationFormat(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Format", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Format", resource.Structure.Representation[numRepresentation].Format, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CodeId(numCode int) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return StringInput("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Id", resource.Code[numCode].Id)
}
func (resource *SubstanceDefinition) T_CodeCode(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Code", resource.Code[numCode].Code, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CodeStatus(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].Status", resource.Code[numCode].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) T_CodeStatusDate(numCode int) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return StringInput("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].StatusDate", nil)
	}
	return StringInput("SubstanceDefinition.Code["+strconv.Itoa(numCode)+"].StatusDate", resource.Code[numCode].StatusDate)
}
func (resource *SubstanceDefinition) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *SubstanceDefinition) T_NameName(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Name", nil)
	}
	return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Name", &resource.Name[numName].Name)
}
func (resource *SubstanceDefinition) T_NameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameStatus(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Status", resource.Name[numName].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) T_NamePreferred(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return BoolInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Preferred", nil)
	}
	return BoolInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Preferred", resource.Name[numName].Preferred)
}
func (resource *SubstanceDefinition) T_NameLanguage(numName int, numLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Language) >= numLanguage {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Language["+strconv.Itoa(numLanguage)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Language["+strconv.Itoa(numLanguage)+"]", &resource.Name[numName].Language[numLanguage], optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Domain) >= numDomain {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", &resource.Name[numName].Domain[numDomain], optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameOfficialId(numName int, numOfficial int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Id", resource.Name[numName].Official[numOfficial].Id)
}
func (resource *SubstanceDefinition) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet)
}
func (resource *SubstanceDefinition) T_NameOfficialDate(numName int, numOfficial int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", nil)
	}
	return StringInput("SubstanceDefinition.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", resource.Name[numName].Official[numOfficial].Date)
}
func (resource *SubstanceDefinition) T_RelationshipId(numRelationship int) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return StringInput("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Id", nil)
	}
	return StringInput("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Id", resource.Relationship[numRelationship].Id)
}
func (resource *SubstanceDefinition) T_RelationshipType(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Type", &resource.Relationship[numRelationship].Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_RelationshipIsDefining(numRelationship int) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return BoolInput("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", nil)
	}
	return BoolInput("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", resource.Relationship[numRelationship].IsDefining)
}
func (resource *SubstanceDefinition) T_RelationshipComparator(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Comparator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.Relationship["+strconv.Itoa(numRelationship)+"].Comparator", resource.Relationship[numRelationship].Comparator, optionsValueSet)
}
func (resource *SubstanceDefinition) T_SourceMaterialId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.SourceMaterial.Id", nil)
	}
	return StringInput("SubstanceDefinition.SourceMaterial.Id", resource.SourceMaterial.Id)
}
func (resource *SubstanceDefinition) T_SourceMaterialType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Type", resource.SourceMaterial.Type, optionsValueSet)
}
func (resource *SubstanceDefinition) T_SourceMaterialGenus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Genus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Genus", resource.SourceMaterial.Genus, optionsValueSet)
}
func (resource *SubstanceDefinition) T_SourceMaterialSpecies(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Species", resource.SourceMaterial.Species, optionsValueSet)
}
func (resource *SubstanceDefinition) T_SourceMaterialPart(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Part", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Part", resource.SourceMaterial.Part, optionsValueSet)
}
func (resource *SubstanceDefinition) T_SourceMaterialCountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SourceMaterial.CountryOfOrigin) >= numCountryOfOrigin {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", &resource.SourceMaterial.CountryOfOrigin[numCountryOfOrigin], optionsValueSet)
}
