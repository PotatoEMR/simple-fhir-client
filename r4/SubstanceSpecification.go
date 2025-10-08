package r4

//generated with command go run ./bultaoreune
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
	StatusDate        *FhirDateTime    `json:"statusDate,omitempty"`
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
	Date              *FhirDateTime    `json:"date,omitempty"`
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
func (resource *SubstanceSpecification) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Domain(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Source(frs []FhirResource, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSource >= len(resource.Source) {
		return ReferenceInput(frs, "source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source["+strconv.Itoa(numSource)+"]", &resource.Source[numSource], htmlAttrs)
}
func (resource *SubstanceSpecification) T_Comment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *SubstanceSpecification) T_ReferenceInformation(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "referenceInformation", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "referenceInformation", resource.ReferenceInformation, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NucleicAcid(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "nucleicAcid", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "nucleicAcid", resource.NucleicAcid, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Polymer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "polymer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "polymer", resource.Polymer, htmlAttrs)
}
func (resource *SubstanceSpecification) T_Protein(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "protein", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "protein", resource.Protein, htmlAttrs)
}
func (resource *SubstanceSpecification) T_SourceMaterial(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "sourceMaterial", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "sourceMaterial", resource.SourceMaterial, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyRole(numMoiety int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].role", resource.Moiety[numMoiety].Role, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyName(numMoiety int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].name", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].name", resource.Moiety[numMoiety].Name, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyStereochemistry(numMoiety int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyOpticalActivity(numMoiety int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].opticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moiety["+strconv.Itoa(numMoiety)+"].opticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyMolecularFormula(numMoiety int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].molecularFormula", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].molecularFormula", resource.Moiety[numMoiety].MolecularFormula, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyAmountQuantity(numMoiety int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return QuantityInput("moiety["+strconv.Itoa(numMoiety)+"].amountQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("moiety["+strconv.Itoa(numMoiety)+"].amountQuantity", resource.Moiety[numMoiety].AmountQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_MoietyAmountString(numMoiety int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMoiety >= len(resource.Moiety) {
		return StringInput("moiety["+strconv.Itoa(numMoiety)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("moiety["+strconv.Itoa(numMoiety)+"].amountString", resource.Moiety[numMoiety].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyCategory(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].category", resource.Property[numProperty].Category, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].code", resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyParameters(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].parameters", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].parameters", resource.Property[numProperty].Parameters, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyDefiningSubstanceReference(frs []FhirResource, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return ReferenceInput(frs, "property["+strconv.Itoa(numProperty)+"].definingSubstanceReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "property["+strconv.Itoa(numProperty)+"].definingSubstanceReference", resource.Property[numProperty].DefiningSubstanceReference, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyDefiningSubstanceCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].definingSubstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].definingSubstanceCodeableConcept", resource.Property[numProperty].DefiningSubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyAmountQuantity(numProperty int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].amountQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].amountQuantity", resource.Property[numProperty].AmountQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_PropertyAmountString(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].amountString", resource.Property[numProperty].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureStereochemistry(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("structure.stereochemistry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.stereochemistry", resource.Structure.Stereochemistry, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureOpticalActivity(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("structure.opticalActivity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.opticalActivity", resource.Structure.OpticalActivity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormula(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("structure.molecularFormula", nil, htmlAttrs)
	}
	return StringInput("structure.molecularFormula", resource.Structure.MolecularFormula, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureMolecularFormulaByMoiety(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("structure.molecularFormulaByMoiety", nil, htmlAttrs)
	}
	return StringInput("structure.molecularFormulaByMoiety", resource.Structure.MolecularFormulaByMoiety, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureSource(frs []FhirResource, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSource >= len(resource.Structure.Source) {
		return ReferenceInput(frs, "structure.source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "structure.source["+strconv.Itoa(numSource)+"]", &resource.Structure.Source[numSource], htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeName(numIsotope int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].name", resource.Structure.Isotope[numIsotope].Name, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeSubstitution(numIsotope int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].substitution", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].substitution", resource.Structure.Isotope[numIsotope].Substitution, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeHalfLife(numIsotope int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return QuantityInput("structure.isotope["+strconv.Itoa(numIsotope)+"].halfLife", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("structure.isotope["+strconv.Itoa(numIsotope)+"].halfLife", resource.Structure.Isotope[numIsotope].HalfLife, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightMethod(numIsotope int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.method", resource.Structure.Isotope[numIsotope].MolecularWeight.Method, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightType(numIsotope int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.type", resource.Structure.Isotope[numIsotope].MolecularWeight.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureIsotopeMolecularWeightAmount(numIsotope int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numIsotope >= len(resource.Structure.Isotope) {
		return QuantityInput("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.amount", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("structure.isotope["+strconv.Itoa(numIsotope)+"].molecularWeight.amount", resource.Structure.Isotope[numIsotope].MolecularWeight.Amount, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureRepresentationType(numRepresentation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("structure.representation["+strconv.Itoa(numRepresentation)+"].type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureRepresentationRepresentation(numRepresentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return StringInput("structure.representation["+strconv.Itoa(numRepresentation)+"].representation", nil, htmlAttrs)
	}
	return StringInput("structure.representation["+strconv.Itoa(numRepresentation)+"].representation", resource.Structure.Representation[numRepresentation].Representation, htmlAttrs)
}
func (resource *SubstanceSpecification) T_StructureRepresentationAttachment(numRepresentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepresentation >= len(resource.Structure.Representation) {
		return AttachmentInput("structure.representation["+strconv.Itoa(numRepresentation)+"].attachment", nil, htmlAttrs)
	}
	return AttachmentInput("structure.representation["+strconv.Itoa(numRepresentation)+"].attachment", resource.Structure.Representation[numRepresentation].Attachment, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeCode(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].code", resource.Code[numCode].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeStatus(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"].status", resource.Code[numCode].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeStatusDate(numCode int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return FhirDateTimeInput("code["+strconv.Itoa(numCode)+"].statusDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("code["+strconv.Itoa(numCode)+"].statusDate", resource.Code[numCode].StatusDate, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeComment(numCode int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return StringInput("code["+strconv.Itoa(numCode)+"].comment", nil, htmlAttrs)
	}
	return StringInput("code["+strconv.Itoa(numCode)+"].comment", resource.Code[numCode].Comment, htmlAttrs)
}
func (resource *SubstanceSpecification) T_CodeSource(frs []FhirResource, numCode int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) || numSource >= len(resource.Code[numCode].Source) {
		return ReferenceInput(frs, "code["+strconv.Itoa(numCode)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "code["+strconv.Itoa(numCode)+"].source["+strconv.Itoa(numSource)+"]", &resource.Code[numCode].Source[numSource], htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameName(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].name", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].name", &resource.Name[numName].Name, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameStatus(numName int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].status", resource.Name[numName].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NamePreferred(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("name["+strconv.Itoa(numName)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("name["+strconv.Itoa(numName)+"].preferred", resource.Name[numName].Preferred, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameDomain(numName int, numDomain int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numDomain >= len(resource.Name[numName].Domain) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].domain["+strconv.Itoa(numDomain)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].domain["+strconv.Itoa(numDomain)+"]", &resource.Name[numName].Domain[numDomain], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameJurisdiction(numName int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numJurisdiction >= len(resource.Name[numName].Jurisdiction) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Name[numName].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameSource(frs []FhirResource, numName int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numSource >= len(resource.Name[numName].Source) {
		return ReferenceInput(frs, "name["+strconv.Itoa(numName)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "name["+strconv.Itoa(numName)+"].source["+strconv.Itoa(numSource)+"]", &resource.Name[numName].Source[numSource], htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].authority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_NameOfficialDate(numName int, numOfficial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numOfficial >= len(resource.Name[numName].Official) {
		return FhirDateTimeInput("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("name["+strconv.Itoa(numName)+"].official["+strconv.Itoa(numOfficial)+"].date", resource.Name[numName].Official[numOfficial].Date, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipSubstanceReference(frs []FhirResource, numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return ReferenceInput(frs, "relationship["+strconv.Itoa(numRelationship)+"].substanceReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relationship["+strconv.Itoa(numRelationship)+"].substanceReference", resource.Relationship[numRelationship].SubstanceReference, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipSubstanceCodeableConcept(numRelationship int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].substanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].substanceCodeableConcept", resource.Relationship[numRelationship].SubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipRelationship(numRelationship int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].relationship", resource.Relationship[numRelationship].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipIsDefining(numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return BoolInput("relationship["+strconv.Itoa(numRelationship)+"].isDefining", nil, htmlAttrs)
	}
	return BoolInput("relationship["+strconv.Itoa(numRelationship)+"].isDefining", resource.Relationship[numRelationship].IsDefining, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountQuantity(numRelationship int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return QuantityInput("relationship["+strconv.Itoa(numRelationship)+"].amountQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("relationship["+strconv.Itoa(numRelationship)+"].amountQuantity", resource.Relationship[numRelationship].AmountQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountRange(numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return RangeInput("relationship["+strconv.Itoa(numRelationship)+"].amountRange", nil, htmlAttrs)
	}
	return RangeInput("relationship["+strconv.Itoa(numRelationship)+"].amountRange", resource.Relationship[numRelationship].AmountRange, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountRatio(numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return RatioInput("relationship["+strconv.Itoa(numRelationship)+"].amountRatio", nil, htmlAttrs)
	}
	return RatioInput("relationship["+strconv.Itoa(numRelationship)+"].amountRatio", resource.Relationship[numRelationship].AmountRatio, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountString(numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return StringInput("relationship["+strconv.Itoa(numRelationship)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("relationship["+strconv.Itoa(numRelationship)+"].amountString", resource.Relationship[numRelationship].AmountString, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountRatioLowLimit(numRelationship int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return RatioInput("relationship["+strconv.Itoa(numRelationship)+"].amountRatioLowLimit", nil, htmlAttrs)
	}
	return RatioInput("relationship["+strconv.Itoa(numRelationship)+"].amountRatioLowLimit", resource.Relationship[numRelationship].AmountRatioLowLimit, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipAmountType(numRelationship int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) {
		return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].amountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship["+strconv.Itoa(numRelationship)+"].amountType", resource.Relationship[numRelationship].AmountType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSpecification) T_RelationshipSource(frs []FhirResource, numRelationship int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelationship >= len(resource.Relationship) || numSource >= len(resource.Relationship[numRelationship].Source) {
		return ReferenceInput(frs, "relationship["+strconv.Itoa(numRelationship)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relationship["+strconv.Itoa(numRelationship)+"].source["+strconv.Itoa(numSource)+"]", &resource.Relationship[numRelationship].Source[numSource], htmlAttrs)
}
