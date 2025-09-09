package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	ValueDate            *string          `json:"valueDate,omitempty"`
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
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Domain(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Grade(numGrade int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGrade >= len(resource.Grade) {
		return CodeableConceptSelect("grade["+strconv.Itoa(numGrade)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("grade["+strconv.Itoa(numGrade)+"]", &resource.Grade[numGrade], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *SubstanceDefinition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyRole(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].role", resource.Moiety[numMoiety].Role, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyName(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].name", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].name", resource.Moiety[numMoiety].Name, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].opticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].opticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyMolecularFormula(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].molecularFormula", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].molecularFormula", resource.Moiety[numMoiety].MolecularFormula, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyAmountString(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].amountString", resource.Moiety[numMoiety].AmountString, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MoietyMeasurementType(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].measurementType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].measurementType", resource.Moiety[numMoiety].MeasurementType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueDate(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *SubstanceDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MolecularWeightMethod(numMolecularWeight int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMolecularWeight >= len(resource.MolecularWeight) {
		return CodeableConceptSelect("molecularWeight["+strconv.Itoa(numMolecularWeight)+"].method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("molecularWeight["+strconv.Itoa(numMolecularWeight)+"].method", resource.MolecularWeight[numMolecularWeight].Method, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_MolecularWeightType(numMolecularWeight int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMolecularWeight >= len(resource.MolecularWeight) {
		return CodeableConceptSelect("molecularWeight["+strconv.Itoa(numMolecularWeight)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("molecularWeight["+strconv.Itoa(numMolecularWeight)+"].type", resource.MolecularWeight[numMolecularWeight].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureStereochemistry(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("structure.stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.stereochemistry", resource.Structure.Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureOpticalActivity(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("structure.opticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.opticalActivity", resource.Structure.OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormula(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("structure.molecularFormula", nil, htmlAttrs)
	}
	return StringInput("structure.molecularFormula", resource.Structure.MolecularFormula, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureMolecularFormulaByMoiety(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("structure.molecularFormulaByMoiety", nil, htmlAttrs)
	}
	return StringInput("structure.molecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureTechnique(numTechnique int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTechnique >= len(resource.Structure.Technique) {
		return CodeableConceptSelect("structure.technique["+strconv.Itoa(numTechnique)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.technique["+strconv.Itoa(numTechnique)+"]", &resource.Structure.Technique[numTechnique], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationRepresentation(numRepresentation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return StringInput("structure.representation["+strconv.Itoa(numRepresentation)+"].representation", nil, htmlAttrs)
	}
	return StringInput("structure.representation["+strconv.Itoa(numRepresentation)+"].representation", resource.Structure.Representation[numRepresentation].Representation, htmlAttrs)
}
func (resource *SubstanceDefinition) T_StructureRepresentationFormat(numRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].format", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].format", resource.Structure.Representation[numRepresentation].Format, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeCode(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].code", resource.Code[numCode].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeStatus(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].status", resource.Code[numCode].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeStatusDate(numCode int, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return DateTimeInput("code["+strconv.Itoa(numCode)+"].statusDate", nil, htmlAttrs)
	}
	return DateTimeInput("code["+strconv.Itoa(numCode)+"].statusDate", resource.Code[numCode].StatusDate, htmlAttrs)
}
func (resource *SubstanceDefinition) T_CodeNote(numCode int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) || numNote >= len(resource.Code[numCode].Note) {
		return AnnotationTextArea("code["+strconv.Itoa(numCode)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("code["+strconv.Itoa(numCode)+"].note["+strconv.Itoa(numNote)+"]", &resource.Code[numCode].Note[numNote], htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameName(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].name", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameStatus(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].status", resource.Name[numName].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NamePreferred(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("name["+strconv.Itoa(numName)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("name["+strconv.Itoa(numName)+"].preferred", resource.Name[numName].Preferred, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numDomain >= len(resource.Name[numName].Domain) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].domain["+strconv.Itoa(numDomain)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].domain["+strconv.Itoa(numDomain)+"]", &resource.Name[numName].Domain[numDomain], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numJurisdiction >= len(resource.Name[numName].Jurisdiction) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].authority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_NameOfficialDate(numName int, numOfficial int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return DateTimeInput("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].date", nil, htmlAttrs)
	}
	return DateTimeInput("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].date", resource.Name[numName].Official[numOfficial].Date, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipSubstanceDefinitionCodeableConcept(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].substanceDefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].substanceDefinitionCodeableConcept", resource.Relationship[numRelationship].SubstanceDefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipType(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].type", &resource.Relationship[numRelationship].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipIsDefining(numRelationship int, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return BoolInput("relationship["+strconv.Itoa(numRelationship)+"].isDefining", nil, htmlAttrs)
	}
	return BoolInput("relationship["+strconv.Itoa(numRelationship)+"].isDefining", resource.Relationship[numRelationship].IsDefining, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipAmountString(numRelationship int, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return StringInput("relationship["+strconv.Itoa(numRelationship)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("relationship["+strconv.Itoa(numRelationship)+"].amountString", resource.Relationship[numRelationship].AmountString, htmlAttrs)
}
func (resource *SubstanceDefinition) T_RelationshipComparator(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].comparator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].comparator", resource.Relationship[numRelationship].Comparator, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterial.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterial.type", resource.SourceMaterial.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialGenus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterial.genus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterial.genus", resource.SourceMaterial.Genus, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialSpecies(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterial.species", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterial.species", resource.SourceMaterial.Species, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialPart(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterial.part", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterial.part", resource.SourceMaterial.Part, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceDefinition) T_SourceMaterialCountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCountryOfOrigin >= len(resource.SourceMaterial.CountryOfOrigin) {
		return CodeableConceptSelect("sourceMaterial.countryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterial.countryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", &resource.SourceMaterial.CountryOfOrigin[numCountryOfOrigin], optionsValueSet, htmlAttrs)
}
