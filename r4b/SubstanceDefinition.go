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
	AmountQuantity    *Quantity        `json:"amountQuantity,omitempty"`
	AmountString      *string          `json:"amountString,omitempty"`
	MeasurementType   *CodeableConcept `json:"measurementType,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionProperty struct {
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
	StatusDate        *time.Time       `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubstanceDefinition
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
func (r SubstanceDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "SubstanceDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("SubstanceDefinition.Classification."+strconv.Itoa(numClassification)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Classification."+strconv.Itoa(numClassification)+".", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Domain(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Grade(numGrade int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numGrade >= len(resource.Grade) {
		return CodeableConceptSelect("SubstanceDefinition.Grade."+strconv.Itoa(numGrade)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Grade."+strconv.Itoa(numGrade)+".", &resource.Grade[numGrade], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("SubstanceDefinition.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("SubstanceDefinition.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyRole(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Role", resource.Moiety[numMoiety].Role, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyName(numMoiety int, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Name", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Name", resource.Moiety[numMoiety].Name, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..Stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..OpticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..OpticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyMolecularFormula(numMoiety int, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..MolecularFormula", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..MolecularFormula", resource.Moiety[numMoiety].MolecularFormula, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyAmountString(numMoiety int, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..AmountString", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..AmountString", resource.Moiety[numMoiety].AmountString, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyMeasurementType(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..MeasurementType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Moiety."+strconv.Itoa(numMoiety)+"..MeasurementType", resource.Moiety[numMoiety].MeasurementType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueDate(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("SubstanceDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MolecularWeightMethod(numMolecularWeight int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMolecularWeight >= len(resource.MolecularWeight) {
		return CodeableConceptSelect("SubstanceDefinition.MolecularWeight."+strconv.Itoa(numMolecularWeight)+"..Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.MolecularWeight."+strconv.Itoa(numMolecularWeight)+"..Method", resource.MolecularWeight[numMolecularWeight].Method, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MolecularWeightType(numMolecularWeight int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMolecularWeight >= len(resource.MolecularWeight) {
		return CodeableConceptSelect("SubstanceDefinition.MolecularWeight."+strconv.Itoa(numMolecularWeight)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.MolecularWeight."+strconv.Itoa(numMolecularWeight)+"..Type", resource.MolecularWeight[numMolecularWeight].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureStereochemistry(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Stereochemistry", resource.Structure.Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureOpticalActivity(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.Structure.OpticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.OpticalActivity", resource.Structure.OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormula(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Structure.MolecularFormula", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Structure.MolecularFormula", resource.Structure.MolecularFormula, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormulaByMoiety(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubstanceDefinition.Structure.MolecularFormulaByMoiety", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Structure.MolecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureTechnique(numTechnique int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTechnique >= len(resource.Structure.Technique) {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Technique."+strconv.Itoa(numTechnique)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Technique."+strconv.Itoa(numTechnique)+".", &resource.Structure.Technique[numTechnique], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationRepresentation(numRepresentation int, htmlAttrs string) templ.Component {

	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return StringInput("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Representation", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Representation", resource.Structure.Representation[numRepresentation].Representation, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationFormat(numRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Format", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Structure.Representation."+strconv.Itoa(numRepresentation)+"..Format", resource.Structure.Representation[numRepresentation].Format, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeCode(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Code", resource.Code[numCode].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeStatus(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Status", resource.Code[numCode].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeStatusDate(numCode int, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return DateTimeInput("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..StatusDate", resource.Code[numCode].StatusDate, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeNote(numCode int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) || numNote >= len(resource.Code[numCode].Note) {
		return AnnotationTextArea("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("SubstanceDefinition.Code."+strconv.Itoa(numCode)+"..Note."+strconv.Itoa(numNote)+".", &resource.Code[numCode].Note[numNote], htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameName(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return StringInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Name", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameStatus(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Status", resource.Name[numName].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NamePreferred(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Preferred", nil, htmlAttrs)
	}
	return BoolInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Preferred", resource.Name[numName].Preferred, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numDomain >= len(resource.Name[numName].Domain) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Domain."+strconv.Itoa(numDomain)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Domain."+strconv.Itoa(numDomain)+".", &resource.Name[numName].Domain[numDomain], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numJurisdiction >= len(resource.Name[numName].Jurisdiction) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Authority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialDate(numName int, numOfficial int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return DateTimeInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Date", nil, htmlAttrs)
	}
	return DateTimeInput("SubstanceDefinition.Name."+strconv.Itoa(numName)+"..Official."+strconv.Itoa(numOfficial)+"..Date", resource.Name[numName].Official[numOfficial].Date, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipSubstanceDefinitionCodeableConcept(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..SubstanceDefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..SubstanceDefinitionCodeableConcept", resource.Relationship[numRelationship].SubstanceDefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipType(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..Type", &resource.Relationship[numRelationship].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipIsDefining(numRelationship int, htmlAttrs string) templ.Component {

	if resource == nil || numRelationship >= len(resource.Relationship) {
		return BoolInput("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..IsDefining", nil, htmlAttrs)
	}
	return BoolInput("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..IsDefining", resource.Relationship[numRelationship].IsDefining, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipAmountString(numRelationship int, htmlAttrs string) templ.Component {

	if resource == nil || numRelationship >= len(resource.Relationship) {
		return StringInput("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..AmountString", nil, htmlAttrs)
	}
	return StringInput("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..AmountString", resource.Relationship[numRelationship].AmountString, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipComparator(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..Comparator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.Relationship."+strconv.Itoa(numRelationship)+"..Comparator", resource.Relationship[numRelationship].Comparator, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Type", resource.SourceMaterial.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialGenus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Genus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Genus", resource.SourceMaterial.Genus, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialSpecies(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Species", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Species", resource.SourceMaterial.Species, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialPart(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Part", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.Part", resource.SourceMaterial.Part, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialCountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCountryOfOrigin >= len(resource.SourceMaterial.CountryOfOrigin) {
		return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.CountryOfOrigin."+strconv.Itoa(numCountryOfOrigin)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceDefinition.SourceMaterial.CountryOfOrigin."+strconv.Itoa(numCountryOfOrigin)+".", &resource.SourceMaterial.CountryOfOrigin[numCountryOfOrigin], optionsValueSet, htmlAttrs)
}
