package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	AmountQuantity    *Quantity        `json:"amountQuantity"`
	AmountString      *string          `json:"amountString"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSpecification
type SubstanceSpecificationProperty struct {
	Id                               *string          `json:"id,omitempty"`
	Extension                        []Extension      `json:"extension,omitempty"`
	ModifierExtension                []Extension      `json:"modifierExtension,omitempty"`
	Category                         *CodeableConcept `json:"category,omitempty"`
	Code                             *CodeableConcept `json:"code,omitempty"`
	Parameters                       *string          `json:"parameters,omitempty"`
	DefiningSubstanceReference       *Reference       `json:"definingSubstanceReference"`
	DefiningSubstanceCodeableConcept *CodeableConcept `json:"definingSubstanceCodeableConcept"`
	AmountQuantity                   *Quantity        `json:"amountQuantity"`
	AmountString                     *string          `json:"amountString"`
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
	SubstanceReference       *Reference       `json:"substanceReference"`
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept"`
	Relationship             *CodeableConcept `json:"relationship,omitempty"`
	IsDefining               *bool            `json:"isDefining,omitempty"`
	AmountQuantity           *Quantity        `json:"amountQuantity"`
	AmountRange              *Range           `json:"amountRange"`
	AmountRatio              *Ratio           `json:"amountRatio"`
	AmountString             *string          `json:"amountString"`
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

func (resource *SubstanceSpecification) SubstanceSpecificationLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationDomain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationMoietyRole(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Moiety[numMoiety].Role, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationMoietyStereochemistry(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("stereochemistry", resource.Moiety[numMoiety].Stereochemistry, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationMoietyOpticalActivity(numMoiety int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Moiety) >= numMoiety {
		return CodeableConceptSelect("opticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("opticalActivity", resource.Moiety[numMoiety].OpticalActivity, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationPropertyCategory(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Property[numProperty].Category, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationPropertyCode(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Property[numProperty].Code, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureStereochemistry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("stereochemistry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("stereochemistry", resource.Structure.Stereochemistry, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureOpticalActivity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("opticalActivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("opticalActivity", resource.Structure.OpticalActivity, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureIsotopeName(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("name", nil, optionsValueSet)
	}
	return CodeableConceptSelect("name", resource.Structure.Isotope[numIsotope].Name, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureIsotopeSubstitution(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("substitution", nil, optionsValueSet)
	}
	return CodeableConceptSelect("substitution", resource.Structure.Isotope[numIsotope].Substitution, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureIsotopeMolecularWeightMethod(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Structure.Isotope[numIsotope].MolecularWeight.Method, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureIsotopeMolecularWeightType(numIsotope int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Isotope) >= numIsotope {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Structure.Isotope[numIsotope].MolecularWeight.Type, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationStructureRepresentationType(numRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Structure.Representation) >= numRepresentation {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Structure.Representation[numRepresentation].Type, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationCodeCode(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Code) >= numCode {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code[numCode].Code, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationCodeStatus(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Code) >= numCode {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Code[numCode].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameStatus(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Name[numName].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameLanguage(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.Name[numName].Language[0], optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameDomain(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", &resource.Name[numName].Domain[0], optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameJurisdiction(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Name[numName].Jurisdiction[0], optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameOfficialAuthority(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("authority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authority", resource.Name[numName].Official[numOfficial].Authority, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationNameOfficialStatus(numName int, numOfficial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].Official) >= numOfficial {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Name[numName].Official[numOfficial].Status, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationRelationshipRelationship(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", resource.Relationship[numRelationship].Relationship, optionsValueSet)
}
func (resource *SubstanceSpecification) SubstanceSpecificationRelationshipAmountType(numRelationship int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relationship) >= numRelationship {
		return CodeableConceptSelect("amountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("amountType", resource.Relationship[numRelationship].AmountType, optionsValueSet)
}
