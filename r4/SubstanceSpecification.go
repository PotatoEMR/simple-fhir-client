package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecification struct {
	Id                   *string                              `json:"id,omitempty"`
	Meta                 *Meta                                `json:"meta,omitempty"`
	ImplicitRules        *string                              `json:"implicitRules,omitempty"`
	Language             *string                              `json:"language,omitempty"`
	Text                 *Narrative                           `json:"text,omitempty"`
	Contained            []Resource                           `json:"contained,omitempty"`
	Extension            []Extension                          `json:"extension,omitempty"`
	ModifierExtension    []Extension                          `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                          `json:"identifier,omitempty"`
	Type                 *CodeableConcept                     `json:"type,omitempty"`
	Status               *CodeableConcept                     `json:"status,omitempty"`
	Domain               *CodeableConcept                     `json:"domain,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	Source               []Reference                          `json:"source,omitempty"`
	Comment              *string                              `json:"comment,omitempty"`
	Moiety               []SubstanceSpecificationMoiety       `json:"moiety,omitempty"`
	Property             []SubstanceSpecificationProperty     `json:"property,omitempty"`
	ReferenceInformation *Reference                           `json:"referenceInformation,omitempty"`
	Structure            *SubstanceSpecificationStructure     `json:"structure,omitempty"`
	Code                 []SubstanceSpecificationCode         `json:"code,omitempty"`
	Name                 []SubstanceSpecificationName         `json:"name,omitempty"`
	Relationship         []SubstanceSpecificationRelationship `json:"relationship,omitempty"`
	NucleicAcid          *Reference                           `json:"nucleicAcid,omitempty"`
	Polymer              *Reference                           `json:"polymer,omitempty"`
	Protein              *Reference                           `json:"protein,omitempty"`
	SourceMaterial       *Reference                           `json:"sourceMaterial,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationMoiety struct {
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
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationProperty struct {
	Id                               *string          `json:"id,omitempty"`
	Extension                        []Extension      `json:"extension,omitempty"`
	ModifierExtension                []Extension      `json:"modifierExtension,omitempty"`
	Category                         *CodeableConcept `json:"category,omitempty"`
	Code                             *CodeableConcept `json:"code,omitempty"`
	Parameters                       *string          `json:"parameters,omitempty"`
	DefiningSubstanceReference       *Reference       `json:"definingSubstanceReference,omitempty"`
	DefiningSubstanceCodeableConcept *CodeableConcept `json:"definingSubstanceCodeableConcept,omitempty"`
	AmountQuantity                   *Quantity        `json:"amountQuantity,omitempty"`
	AmountString                     *string          `json:"amountString,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationStructure struct {
	Id                       *string                                         `json:"id,omitempty"`
	Extension                []Extension                                     `json:"extension,omitempty"`
	ModifierExtension        []Extension                                     `json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                                `json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                                `json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                         `json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                         `json:"molecularFormulaByMoiety,omitempty"`
	Isotope                  []SubstanceSpecificationStructureIsotope        `json:"isotope,omitempty"`
	Source                   []Reference                                     `json:"source,omitempty"`
	Representation           []SubstanceSpecificationStructureRepresentation `json:"representation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationStructureIsotope struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                                            `json:"identifier,omitempty"`
	Name              *CodeableConcept                                       `json:"name,omitempty"`
	Substitution      *CodeableConcept                                       `json:"substitution,omitempty"`
	HalfLife          *Quantity                                              `json:"halfLife,omitempty"`
	MolecularWeight   *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationStructureRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationCode struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	StatusDate        *string          `json:"statusDate,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationName struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Name              string                               `json:"name"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	Status            *CodeableConcept                     `json:"status,omitempty"`
	Preferred         *bool                                `json:"preferred,omitempty"`
	Language          []CodeableConcept                    `json:"language,omitempty"`
	Domain            []CodeableConcept                    `json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Official          []SubstanceSpecificationNameOfficial `json:"official,omitempty"`
	Source            []Reference                          `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationNameOfficial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationRelationship struct {
	Id                       *string          `json:"id,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	SubstanceReference       *Reference       `json:"substanceReference,omitempty"`
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept,omitempty"`
	Relationship             *CodeableConcept `json:"relationship,omitempty"`
	IsDefining               *bool            `json:"isDefining,omitempty"`
	AmountQuantity           *Quantity        `json:"amountQuantity,omitempty"`
	AmountRange              *Range           `json:"amountRange,omitempty"`
	AmountRatio              *Ratio           `json:"amountRatio,omitempty"`
	AmountString             *string          `json:"amountString,omitempty"`
	AmountRatioLowLimit      *Ratio           `json:"amountRatioLowLimit,omitempty"`
	AmountType               *CodeableConcept `json:"amountType,omitempty"`
	Source                   []Reference      `json:"source,omitempty"`
}

type OtherSubstanceSpecification SubstanceSpecification

// on convert struct to json, automatically add resourceType=SubstanceSpecification
func (r SubstanceSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceSpecification
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSpecification: OtherSubstanceSpecification(r),
		ResourceType:                "SubstanceSpecification",
	})
}

func (resource *SubstanceSpecification) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Id", nil)
	}
	return StringInput("SubstanceSpecification.Id", resource.Id)
}
func (resource *SubstanceSpecification) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.ImplicitRules", nil)
	}
	return StringInput("SubstanceSpecification.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceSpecification) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceSpecification.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceSpecification.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceSpecification) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Type", resource.Type, optionsValueSet)
}
func (resource *SubstanceSpecification) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Status", resource.Status, optionsValueSet)
}
func (resource *SubstanceSpecification) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Domain", resource.Domain, optionsValueSet)
}
func (resource *SubstanceSpecification) T_Description() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Description", nil)
	}
	return StringInput("SubstanceSpecification.Description", resource.Description)
}
func (resource *SubstanceSpecification) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Comment", nil)
	}
	return StringInput("SubstanceSpecification.Comment", resource.Comment)
}
func (resource *SubstanceSpecification) T_MoietyId(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Id", resource.Moiety[numMoiety].Id)
}
func (resource *SubstanceSpecification) T_MoietyRole(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Role", resource.Moiety[numMoiety].Role, optionsValueSet)
}
func (resource *SubstanceSpecification) T_MoietyName(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Name", nil)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Name", resource.Moiety[numMoiety].Name)
}
func (resource *SubstanceSpecification) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet)
}
func (resource *SubstanceSpecification) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet)
}
func (resource *SubstanceSpecification) T_MoietyMolecularFormula(numMoiety int) templ.Component {

	if resource == nil || len(resource.Moiety) >= numMoiety {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", nil)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", resource.Moiety[numMoiety].MolecularFormula)
}
func (resource *SubstanceSpecification) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *SubstanceSpecification) T_PropertyCategory(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Category", resource.Property[numProperty].Category, optionsValueSet)
}
func (resource *SubstanceSpecification) T_PropertyCode(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Code", resource.Property[numProperty].Code, optionsValueSet)
}
func (resource *SubstanceSpecification) T_PropertyParameters(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Parameters", nil)
	}
	return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Parameters", resource.Property[numProperty].Parameters)
}
func (resource *SubstanceSpecification) T_StructureId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Structure.Id", nil)
	}
	return StringInput("SubstanceSpecification.Structure.Id", resource.Structure.Id)
}
func (resource *SubstanceSpecification) T_StructureStereochemistry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Stereochemistry", resource.Structure.Stereochemistry, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureOpticalActivity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Structure.OpticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.OpticalActivity", resource.Structure.OpticalActivity, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormula() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Structure.MolecularFormula", nil)
	}
	return StringInput("SubstanceSpecification.Structure.MolecularFormula", resource.Structure.MolecularFormula)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormulaByMoiety() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSpecification.Structure.MolecularFormulaByMoiety", nil)
	}
	return StringInput("SubstanceSpecification.Structure.MolecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety)
}
func (resource *SubstanceSpecification) T_StructureIsotopeId(numIsotope int) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return StringInput("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Id", resource.Structure.Isotope[numIsotope].Id)
}
func (resource *SubstanceSpecification) T_StructureIsotopeName(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Name", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Name", resource.Structure.Isotope[numIsotope].Name, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureIsotopeSubstitution(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Substitution", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Substitution", resource.Structure.Isotope[numIsotope].Substitution, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightId(numIsotope int) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return StringInput("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Id", nil)
	}
	return StringInput("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Id", resource.Structure.Isotope[numIsotope].MolecularWeight.Id)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightMethod(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Method", resource.Structure.Isotope[numIsotope].MolecularWeight.Method, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightType(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Type", resource.Structure.Isotope[numIsotope].MolecularWeight.Type, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureRepresentationId(numRepresentation int) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Id", resource.Structure.Representation[numRepresentation].Id)
}
func (resource *SubstanceSpecification) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet)
}
func (resource *SubstanceSpecification) T_StructureRepresentationRepresentation(numRepresentation int) templ.Component {

	if resource == nil || len(resource.Structure.Representation) >= numRepresentation {
		return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", nil)
	}
	return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", resource.Structure.Representation[numRepresentation].Representation)
}
func (resource *SubstanceSpecification) T_CodeId(numCode int) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Id", resource.Code[numCode].Id)
}
func (resource *SubstanceSpecification) T_CodeCode(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Code", resource.Code[numCode].Code, optionsValueSet)
}
func (resource *SubstanceSpecification) T_CodeStatus(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Status", resource.Code[numCode].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) T_CodeStatusDate(numCode int) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].StatusDate", nil)
	}
	return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].StatusDate", resource.Code[numCode].StatusDate)
}
func (resource *SubstanceSpecification) T_CodeComment(numCode int) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Comment", nil)
	}
	return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Comment", resource.Code[numCode].Comment)
}
func (resource *SubstanceSpecification) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *SubstanceSpecification) T_NameName(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Name", nil)
	}
	return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Name", &resource.Name[numName].Name)
}
func (resource *SubstanceSpecification) T_NameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameStatus(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Status", resource.Name[numName].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) T_NamePreferred(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return BoolInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Preferred", nil)
	}
	return BoolInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Preferred", resource.Name[numName].Preferred)
}
func (resource *SubstanceSpecification) T_NameLanguage(numName int, numLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Language) >= numLanguage {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Language["+strconv.Itoa(numLanguage)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Language["+strconv.Itoa(numLanguage)+"]", &resource.Name[numName].Language[numLanguage], optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Domain) >= numDomain {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", &resource.Name[numName].Domain[numDomain], optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameOfficialId(numName int, numOfficial int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Id", resource.Name[numName].Official[numOfficial].Id)
}
func (resource *SubstanceSpecification) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) T_NameOfficialDate(numName int, numOfficial int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Official) >= numOfficial {
		return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", nil)
	}
	return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", resource.Name[numName].Official[numOfficial].Date)
}
func (resource *SubstanceSpecification) T_RelationshipId(numRelationship int) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return StringInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Id", nil)
	}
	return StringInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Id", resource.Relationship[numRelationship].Id)
}
func (resource *SubstanceSpecification) T_RelationshipRelationship(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Relationship", resource.Relationship[numRelationship].Relationship, optionsValueSet)
}
func (resource *SubstanceSpecification) T_RelationshipIsDefining(numRelationship int) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return BoolInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", nil)
	}
	return BoolInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", resource.Relationship[numRelationship].IsDefining)
}
func (resource *SubstanceSpecification) T_RelationshipAmountType(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountType", resource.Relationship[numRelationship].AmountType, optionsValueSet)
}
