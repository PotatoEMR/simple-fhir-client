package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterial struct {
	Id                   *string                                      `json:"id,omitempty"`
	Meta                 *Meta                                        `json:"meta,omitempty"`
	ImplicitRules        *string                                      `json:"implicitRules,omitempty"`
	Language             *string                                      `json:"language,omitempty"`
	Text                 *Narrative                                   `json:"text,omitempty"`
	Contained            []Resource                                   `json:"contained,omitempty"`
	Extension            []Extension                                  `json:"extension,omitempty"`
	ModifierExtension    []Extension                                  `json:"modifierExtension,omitempty"`
	SourceMaterialClass  *CodeableConcept                             `json:"sourceMaterialClass,omitempty"`
	SourceMaterialType   *CodeableConcept                             `json:"sourceMaterialType,omitempty"`
	SourceMaterialState  *CodeableConcept                             `json:"sourceMaterialState,omitempty"`
	OrganismId           *Identifier                                  `json:"organismId,omitempty"`
	OrganismName         *string                                      `json:"organismName,omitempty"`
	ParentSubstanceId    []Identifier                                 `json:"parentSubstanceId,omitempty"`
	ParentSubstanceName  []string                                     `json:"parentSubstanceName,omitempty"`
	CountryOfOrigin      []CodeableConcept                            `json:"countryOfOrigin,omitempty"`
	GeographicalLocation []string                                     `json:"geographicalLocation,omitempty"`
	DevelopmentStage     *CodeableConcept                             `json:"developmentStage,omitempty"`
	FractionDescription  []SubstanceSourceMaterialFractionDescription `json:"fractionDescription,omitempty"`
	Organism             *SubstanceSourceMaterialOrganism             `json:"organism,omitempty"`
	PartDescription      []SubstanceSourceMaterialPartDescription     `json:"partDescription,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialFractionDescription struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Fraction          *string          `json:"fraction,omitempty"`
	MaterialType      *CodeableConcept `json:"materialType,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganism struct {
	Id                       *string                                         `json:"id,omitempty"`
	Extension                []Extension                                     `json:"extension,omitempty"`
	ModifierExtension        []Extension                                     `json:"modifierExtension,omitempty"`
	Family                   *CodeableConcept                                `json:"family,omitempty"`
	Genus                    *CodeableConcept                                `json:"genus,omitempty"`
	Species                  *CodeableConcept                                `json:"species,omitempty"`
	IntraspecificType        *CodeableConcept                                `json:"intraspecificType,omitempty"`
	IntraspecificDescription *string                                         `json:"intraspecificDescription,omitempty"`
	Author                   []SubstanceSourceMaterialOrganismAuthor         `json:"author,omitempty"`
	Hybrid                   *SubstanceSourceMaterialOrganismHybrid          `json:"hybrid,omitempty"`
	OrganismGeneral          *SubstanceSourceMaterialOrganismOrganismGeneral `json:"organismGeneral,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganismAuthor struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	AuthorType        *CodeableConcept `json:"authorType,omitempty"`
	AuthorDescription *string          `json:"authorDescription,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganismHybrid struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	MaternalOrganismId   *string          `json:"maternalOrganismId,omitempty"`
	MaternalOrganismName *string          `json:"maternalOrganismName,omitempty"`
	PaternalOrganismId   *string          `json:"paternalOrganismId,omitempty"`
	PaternalOrganismName *string          `json:"paternalOrganismName,omitempty"`
	HybridType           *CodeableConcept `json:"hybridType,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Kingdom           *CodeableConcept `json:"kingdom,omitempty"`
	Phylum            *CodeableConcept `json:"phylum,omitempty"`
	Class             *CodeableConcept `json:"class,omitempty"`
	Order             *CodeableConcept `json:"order,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialPartDescription struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Part              *CodeableConcept `json:"part,omitempty"`
	PartLocation      *CodeableConcept `json:"partLocation,omitempty"`
}

type OtherSubstanceSourceMaterial SubstanceSourceMaterial

// on convert struct to json, automatically add resourceType=SubstanceSourceMaterial
func (r SubstanceSourceMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceSourceMaterial
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSourceMaterial: OtherSubstanceSourceMaterial(r),
		ResourceType:                 "SubstanceSourceMaterial",
	})
}

func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialSourceMaterialClass(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("sourceMaterialClass", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sourceMaterialClass", resource.SourceMaterialClass, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialSourceMaterialType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("sourceMaterialType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sourceMaterialType", resource.SourceMaterialType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialSourceMaterialState(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("sourceMaterialState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sourceMaterialState", resource.SourceMaterialState, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialCountryOfOrigin(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("countryOfOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("countryOfOrigin", &resource.CountryOfOrigin[0], optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialDevelopmentStage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("developmentStage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("developmentStage", resource.DevelopmentStage, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialFractionDescriptionMaterialType(numFractionDescription int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.FractionDescription) >= numFractionDescription {
		return CodeableConceptSelect("materialType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("materialType", resource.FractionDescription[numFractionDescription].MaterialType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismFamily(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("family", nil, optionsValueSet)
	}
	return CodeableConceptSelect("family", resource.Organism.Family, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismGenus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("genus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("genus", resource.Organism.Genus, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismSpecies(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("species", resource.Organism.Species, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismIntraspecificType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("intraspecificType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("intraspecificType", resource.Organism.IntraspecificType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismAuthorAuthorType(numAuthor int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Organism.Author) >= numAuthor {
		return CodeableConceptSelect("authorType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authorType", resource.Organism.Author[numAuthor].AuthorType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismHybridHybridType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("hybridType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("hybridType", resource.Organism.Hybrid.HybridType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismOrganismGeneralKingdom(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("kingdom", nil, optionsValueSet)
	}
	return CodeableConceptSelect("kingdom", resource.Organism.OrganismGeneral.Kingdom, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismOrganismGeneralPhylum(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("phylum", nil, optionsValueSet)
	}
	return CodeableConceptSelect("phylum", resource.Organism.OrganismGeneral.Phylum, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismOrganismGeneralClass(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("class", resource.Organism.OrganismGeneral.Class, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialOrganismOrganismGeneralOrder(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("order", nil, optionsValueSet)
	}
	return CodeableConceptSelect("order", resource.Organism.OrganismGeneral.Order, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialPartDescriptionPart(numPartDescription int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.PartDescription) >= numPartDescription {
		return CodeableConceptSelect("part", nil, optionsValueSet)
	}
	return CodeableConceptSelect("part", resource.PartDescription[numPartDescription].Part, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) SubstanceSourceMaterialPartDescriptionPartLocation(numPartDescription int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.PartDescription) >= numPartDescription {
		return CodeableConceptSelect("partLocation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("partLocation", resource.PartDescription[numPartDescription].PartLocation, optionsValueSet)
}
