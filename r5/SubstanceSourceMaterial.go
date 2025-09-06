package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialFractionDescription struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Fraction          *string          `json:"fraction,omitempty"`
	MaterialType      *CodeableConcept `json:"materialType,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganismAuthor struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	AuthorType        *CodeableConcept `json:"authorType,omitempty"`
	AuthorDescription *string          `json:"authorDescription,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
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

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Kingdom           *CodeableConcept `json:"kingdom,omitempty"`
	Phylum            *CodeableConcept `json:"phylum,omitempty"`
	Class             *CodeableConcept `json:"class,omitempty"`
	Order             *CodeableConcept `json:"order,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceSourceMaterial
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

func (resource *SubstanceSourceMaterial) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.Id", resource.Id)
}
func (resource *SubstanceSourceMaterial) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.ImplicitRules", nil)
	}
	return StringInput("SubstanceSourceMaterial.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceSourceMaterial) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceSourceMaterial.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceSourceMaterial.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialClass(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialClass", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialClass", resource.SourceMaterialClass, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialType", resource.SourceMaterialType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialState(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.SourceMaterialState", resource.SourceMaterialState, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismName() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.OrganismName", nil)
	}
	return StringInput("SubstanceSourceMaterial.OrganismName", resource.OrganismName)
}
func (resource *SubstanceSourceMaterial) T_ParentSubstanceName(numParentSubstanceName int) templ.Component {

	if resource == nil || len(resource.ParentSubstanceName) >= numParentSubstanceName {
		return StringInput("SubstanceSourceMaterial.ParentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", nil)
	}
	return StringInput("SubstanceSourceMaterial.ParentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", &resource.ParentSubstanceName[numParentSubstanceName])
}
func (resource *SubstanceSourceMaterial) T_CountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CountryOfOrigin) >= numCountryOfOrigin {
		return CodeableConceptSelect("SubstanceSourceMaterial.CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", &resource.CountryOfOrigin[numCountryOfOrigin], optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_GeographicalLocation(numGeographicalLocation int) templ.Component {

	if resource == nil || len(resource.GeographicalLocation) >= numGeographicalLocation {
		return StringInput("SubstanceSourceMaterial.GeographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", nil)
	}
	return StringInput("SubstanceSourceMaterial.GeographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", &resource.GeographicalLocation[numGeographicalLocation])
}
func (resource *SubstanceSourceMaterial) T_DevelopmentStage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.DevelopmentStage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.DevelopmentStage", resource.DevelopmentStage, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionId(numFractionDescription int) templ.Component {

	if resource == nil || len(resource.FractionDescription) >= numFractionDescription {
		return StringInput("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].Id", resource.FractionDescription[numFractionDescription].Id)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionFraction(numFractionDescription int) templ.Component {

	if resource == nil || len(resource.FractionDescription) >= numFractionDescription {
		return StringInput("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].Fraction", nil)
	}
	return StringInput("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].Fraction", resource.FractionDescription[numFractionDescription].Fraction)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionMaterialType(numFractionDescription int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.FractionDescription) >= numFractionDescription {
		return CodeableConceptSelect("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].MaterialType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.FractionDescription["+strconv.Itoa(numFractionDescription)+"].MaterialType", resource.FractionDescription[numFractionDescription].MaterialType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Id", resource.Organism.Id)
}
func (resource *SubstanceSourceMaterial) T_OrganismFamily(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Family", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Family", resource.Organism.Family, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismGenus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Genus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Genus", resource.Organism.Genus, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismSpecies(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Species", resource.Organism.Species, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.IntraspecificType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.IntraspecificType", resource.Organism.IntraspecificType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificDescription() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.IntraspecificDescription", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.IntraspecificDescription", resource.Organism.IntraspecificDescription)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorId(numAuthor int) templ.Component {

	if resource == nil || len(resource.Organism.Author) >= numAuthor {
		return StringInput("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].Id", resource.Organism.Author[numAuthor].Id)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorType(numAuthor int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Organism.Author) >= numAuthor {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].AuthorType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].AuthorType", resource.Organism.Author[numAuthor].AuthorType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorDescription(numAuthor int) templ.Component {

	if resource == nil || len(resource.Organism.Author) >= numAuthor {
		return StringInput("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].AuthorDescription", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Author["+strconv.Itoa(numAuthor)+"].AuthorDescription", resource.Organism.Author[numAuthor].AuthorDescription)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Hybrid.Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Hybrid.Id", resource.Organism.Hybrid.Id)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Hybrid.MaternalOrganismId", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Hybrid.MaternalOrganismId", resource.Organism.Hybrid.MaternalOrganismId)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismName() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Hybrid.MaternalOrganismName", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Hybrid.MaternalOrganismName", resource.Organism.Hybrid.MaternalOrganismName)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Hybrid.PaternalOrganismId", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Hybrid.PaternalOrganismId", resource.Organism.Hybrid.PaternalOrganismId)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismName() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.Hybrid.PaternalOrganismName", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.Hybrid.PaternalOrganismName", resource.Organism.Hybrid.PaternalOrganismName)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridHybridType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Hybrid.HybridType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.Hybrid.HybridType", resource.Organism.Hybrid.HybridType, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralId() templ.Component {

	if resource == nil {
		return StringInput("SubstanceSourceMaterial.Organism.OrganismGeneral.Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.Organism.OrganismGeneral.Id", resource.Organism.OrganismGeneral.Id)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralKingdom(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Kingdom", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Kingdom", resource.Organism.OrganismGeneral.Kingdom, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralPhylum(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Phylum", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Phylum", resource.Organism.OrganismGeneral.Phylum, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralClass(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Class", resource.Organism.OrganismGeneral.Class, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralOrder(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Order", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.Organism.OrganismGeneral.Order", resource.Organism.OrganismGeneral.Order, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionId(numPartDescription int) templ.Component {

	if resource == nil || len(resource.PartDescription) >= numPartDescription {
		return StringInput("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].Id", nil)
	}
	return StringInput("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].Id", resource.PartDescription[numPartDescription].Id)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPart(numPartDescription int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PartDescription) >= numPartDescription {
		return CodeableConceptSelect("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].Part", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].Part", resource.PartDescription[numPartDescription].Part, optionsValueSet)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPartLocation(numPartDescription int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PartDescription) >= numPartDescription {
		return CodeableConceptSelect("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].PartLocation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceSourceMaterial.PartDescription["+strconv.Itoa(numPartDescription)+"].PartLocation", resource.PartDescription[numPartDescription].PartLocation, optionsValueSet)
}
