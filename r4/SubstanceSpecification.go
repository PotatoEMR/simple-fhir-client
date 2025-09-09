package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	StatusDate        *time.Time       `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r SubstanceSpecification) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceSpecification/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "SubstanceSpecification"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceSpecification) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Domain(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SubstanceSpecification.Description", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Description", resource.Description, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SubstanceSpecification.Comment", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Comment", resource.Comment, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyRole(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Role", resource.Moiety[numMoiety].Role, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyName(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Name", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Name", resource.Moiety[numMoiety].Name, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].Stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].OpticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyMolecularFormula(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].MolecularFormula", resource.Moiety[numMoiety].MolecularFormula, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyAmountString(numMoiety int, htmlAttrs string) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].AmountString", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Moiety["+strconv.Itoa(numMoiety)+"].AmountString", resource.Moiety[numMoiety].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyCategory(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Category", resource.Property[numProperty].Category, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Code", resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyParameters(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Parameters", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].Parameters", resource.Property[numProperty].Parameters, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyDefiningSubstanceCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].DefiningSubstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].DefiningSubstanceCodeableConcept", resource.Property[numProperty].DefiningSubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyAmountString(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].AmountString", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Property["+strconv.Itoa(numProperty)+"].AmountString", resource.Property[numProperty].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureStereochemistry(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Stereochemistry", resource.Structure.Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureOpticalActivity(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceSpecification.Structure.OpticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.OpticalActivity", resource.Structure.OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormula(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SubstanceSpecification.Structure.MolecularFormula", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Structure.MolecularFormula", resource.Structure.MolecularFormula, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormulaByMoiety(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SubstanceSpecification.Structure.MolecularFormulaByMoiety", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Structure.MolecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeName(numIsotope int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Name", resource.Structure.Isotope[numIsotope].Name, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeSubstitution(numIsotope int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Substitution", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].Substitution", resource.Structure.Isotope[numIsotope].Substitution, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightMethod(numIsotope int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Method", resource.Structure.Isotope[numIsotope].MolecularWeight.Method, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightType(numIsotope int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Isotope["+strconv.Itoa(numIsotope)+"].MolecularWeight.Type", resource.Structure.Isotope[numIsotope].MolecularWeight.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureRepresentationRepresentation(numRepresentation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Structure.Representation["+strconv.Itoa(numRepresentation)+"].Representation", resource.Structure.Representation[numRepresentation].Representation, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeCode(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Code", resource.Code[numCode].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeStatus(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Status", resource.Code[numCode].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeStatusDate(numCode int, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return DateTimeInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].StatusDate", resource.Code[numCode].StatusDate, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeComment(numCode int, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Code["+strconv.Itoa(numCode)+"].Comment", resource.Code[numCode].Comment, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameName(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Name", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameStatus(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Status", resource.Name[numName].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NamePreferred(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Preferred", nil, htmlAttrs)
	}
	return BoolInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Preferred", resource.Name[numName].Preferred, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numDomain >= len(resource.Name[numName].Domain) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Domain["+strconv.Itoa(numDomain)+"]", &resource.Name[numName].Domain[numDomain], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numJurisdiction >= len(resource.Name[numName].Jurisdiction) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialDate(numName int, numOfficial int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return DateTimeInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", nil, htmlAttrs)
	}
	return DateTimeInput("SubstanceSpecification.Name["+strconv.Itoa(numName)+"].Official["+strconv.Itoa(numOfficial)+"].Date", resource.Name[numName].Official[numOfficial].Date, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipSubstanceCodeableConcept(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].SubstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].SubstanceCodeableConcept", resource.Relationship[numRelationship].SubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipRelationship(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].Relationship", resource.Relationship[numRelationship].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipIsDefining(numRelationship int, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return BoolInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", nil, htmlAttrs)
	}
	return BoolInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].IsDefining", resource.Relationship[numRelationship].IsDefining, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountString(numRelationship int, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return StringInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountString", nil, htmlAttrs)
	}
	return StringInput("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountString", resource.Relationship[numRelationship].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountType(numRelationship int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceSpecification.Relationship["+strconv.Itoa(numRelationship)+"].AmountType", resource.Relationship[numRelationship].AmountType, optionsValueSet, htmlAttrs)
}
