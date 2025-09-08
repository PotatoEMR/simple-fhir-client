package r5

//generated with command go run ./bultaoreune
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
func (r SubstanceSourceMaterial) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceSourceMaterial/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubstanceSourceMaterial"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialClass(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SourceMaterialClass", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SourceMaterialClass", resource.SourceMaterialClass, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SourceMaterialType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SourceMaterialType", resource.SourceMaterialType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialState(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SourceMaterialState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SourceMaterialState", resource.SourceMaterialState, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismName", nil, htmlAttrs)
	}
	return StringInput("OrganismName", resource.OrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_ParentSubstanceName(numParentSubstanceName int, htmlAttrs string) templ.Component {
	if resource == nil || numParentSubstanceName >= len(resource.ParentSubstanceName) {
		return StringInput("ParentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", nil, htmlAttrs)
	}
	return StringInput("ParentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", &resource.ParentSubstanceName[numParentSubstanceName], htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_CountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCountryOfOrigin >= len(resource.CountryOfOrigin) {
		return CodeableConceptSelect("CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CountryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", &resource.CountryOfOrigin[numCountryOfOrigin], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_GeographicalLocation(numGeographicalLocation int, htmlAttrs string) templ.Component {
	if resource == nil || numGeographicalLocation >= len(resource.GeographicalLocation) {
		return StringInput("GeographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", nil, htmlAttrs)
	}
	return StringInput("GeographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", &resource.GeographicalLocation[numGeographicalLocation], htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_DevelopmentStage(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DevelopmentStage", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DevelopmentStage", resource.DevelopmentStage, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionFraction(numFractionDescription int, htmlAttrs string) templ.Component {
	if resource == nil || numFractionDescription >= len(resource.FractionDescription) {
		return StringInput("FractionDescription["+strconv.Itoa(numFractionDescription)+"]Fraction", nil, htmlAttrs)
	}
	return StringInput("FractionDescription["+strconv.Itoa(numFractionDescription)+"]Fraction", resource.FractionDescription[numFractionDescription].Fraction, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionMaterialType(numFractionDescription int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFractionDescription >= len(resource.FractionDescription) {
		return CodeableConceptSelect("FractionDescription["+strconv.Itoa(numFractionDescription)+"]MaterialType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FractionDescription["+strconv.Itoa(numFractionDescription)+"]MaterialType", resource.FractionDescription[numFractionDescription].MaterialType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismFamily(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismFamily", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismFamily", resource.Organism.Family, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismGenus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismGenus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismGenus", resource.Organism.Genus, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismSpecies(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismSpecies", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismSpecies", resource.Organism.Species, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismIntraspecificType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismIntraspecificType", resource.Organism.IntraspecificType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismIntraspecificDescription", nil, htmlAttrs)
	}
	return StringInput("OrganismIntraspecificDescription", resource.Organism.IntraspecificDescription, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorType(numAuthor int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAuthor >= len(resource.Organism.Author) {
		return CodeableConceptSelect("OrganismAuthor["+strconv.Itoa(numAuthor)+"].AuthorType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismAuthor["+strconv.Itoa(numAuthor)+"].AuthorType", resource.Organism.Author[numAuthor].AuthorType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorDescription(numAuthor int, htmlAttrs string) templ.Component {
	if resource == nil || numAuthor >= len(resource.Organism.Author) {
		return StringInput("OrganismAuthor["+strconv.Itoa(numAuthor)+"].AuthorDescription", nil, htmlAttrs)
	}
	return StringInput("OrganismAuthor["+strconv.Itoa(numAuthor)+"].AuthorDescription", resource.Organism.Author[numAuthor].AuthorDescription, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismId(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismHybrid.MaternalOrganismId", nil, htmlAttrs)
	}
	return StringInput("OrganismHybrid.MaternalOrganismId", resource.Organism.Hybrid.MaternalOrganismId, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismHybrid.MaternalOrganismName", nil, htmlAttrs)
	}
	return StringInput("OrganismHybrid.MaternalOrganismName", resource.Organism.Hybrid.MaternalOrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismId(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismHybrid.PaternalOrganismId", nil, htmlAttrs)
	}
	return StringInput("OrganismHybrid.PaternalOrganismId", resource.Organism.Hybrid.PaternalOrganismId, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OrganismHybrid.PaternalOrganismName", nil, htmlAttrs)
	}
	return StringInput("OrganismHybrid.PaternalOrganismName", resource.Organism.Hybrid.PaternalOrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridHybridType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismHybrid.HybridType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismHybrid.HybridType", resource.Organism.Hybrid.HybridType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralKingdom(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismOrganismGeneral.Kingdom", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismOrganismGeneral.Kingdom", resource.Organism.OrganismGeneral.Kingdom, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralPhylum(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismOrganismGeneral.Phylum", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismOrganismGeneral.Phylum", resource.Organism.OrganismGeneral.Phylum, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralClass(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismOrganismGeneral.Class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismOrganismGeneral.Class", resource.Organism.OrganismGeneral.Class, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralOrder(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrganismOrganismGeneral.Order", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganismOrganismGeneral.Order", resource.Organism.OrganismGeneral.Order, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPart(numPartDescription int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPartDescription >= len(resource.PartDescription) {
		return CodeableConceptSelect("PartDescription["+strconv.Itoa(numPartDescription)+"]Part", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PartDescription["+strconv.Itoa(numPartDescription)+"]Part", resource.PartDescription[numPartDescription].Part, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPartLocation(numPartDescription int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPartDescription >= len(resource.PartDescription) {
		return CodeableConceptSelect("PartDescription["+strconv.Itoa(numPartDescription)+"]PartLocation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PartDescription["+strconv.Itoa(numPartDescription)+"]PartLocation", resource.PartDescription[numPartDescription].PartLocation, optionsValueSet, htmlAttrs)
}
