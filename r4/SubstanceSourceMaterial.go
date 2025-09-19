package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *SubstanceSourceMaterial) T_SourceMaterialClass(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterialClass", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterialClass", resource.SourceMaterialClass, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterialType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterialType", resource.SourceMaterialType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_SourceMaterialState(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sourceMaterialState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sourceMaterialState", resource.SourceMaterialState, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("organismId", nil, htmlAttrs)
	}
	return IdentifierInput("organismId", resource.OrganismId, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organismName", nil, htmlAttrs)
	}
	return StringInput("organismName", resource.OrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_ParentSubstanceId(numParentSubstanceId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParentSubstanceId >= len(resource.ParentSubstanceId) {
		return IdentifierInput("parentSubstanceId["+strconv.Itoa(numParentSubstanceId)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("parentSubstanceId["+strconv.Itoa(numParentSubstanceId)+"]", &resource.ParentSubstanceId[numParentSubstanceId], htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_ParentSubstanceName(numParentSubstanceName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParentSubstanceName >= len(resource.ParentSubstanceName) {
		return StringInput("parentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", nil, htmlAttrs)
	}
	return StringInput("parentSubstanceName["+strconv.Itoa(numParentSubstanceName)+"]", &resource.ParentSubstanceName[numParentSubstanceName], htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_CountryOfOrigin(numCountryOfOrigin int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCountryOfOrigin >= len(resource.CountryOfOrigin) {
		return CodeableConceptSelect("countryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("countryOfOrigin["+strconv.Itoa(numCountryOfOrigin)+"]", &resource.CountryOfOrigin[numCountryOfOrigin], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_GeographicalLocation(numGeographicalLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGeographicalLocation >= len(resource.GeographicalLocation) {
		return StringInput("geographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", nil, htmlAttrs)
	}
	return StringInput("geographicalLocation["+strconv.Itoa(numGeographicalLocation)+"]", &resource.GeographicalLocation[numGeographicalLocation], htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_DevelopmentStage(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("developmentStage", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("developmentStage", resource.DevelopmentStage, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionFraction(numFractionDescription int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFractionDescription >= len(resource.FractionDescription) {
		return StringInput("fractionDescription["+strconv.Itoa(numFractionDescription)+"].fraction", nil, htmlAttrs)
	}
	return StringInput("fractionDescription["+strconv.Itoa(numFractionDescription)+"].fraction", resource.FractionDescription[numFractionDescription].Fraction, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_FractionDescriptionMaterialType(numFractionDescription int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFractionDescription >= len(resource.FractionDescription) {
		return CodeableConceptSelect("fractionDescription["+strconv.Itoa(numFractionDescription)+"].materialType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fractionDescription["+strconv.Itoa(numFractionDescription)+"].materialType", resource.FractionDescription[numFractionDescription].MaterialType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismFamily(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.family", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.family", resource.Organism.Family, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismGenus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.genus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.genus", resource.Organism.Genus, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismSpecies(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.species", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.species", resource.Organism.Species, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.intraspecificType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.intraspecificType", resource.Organism.IntraspecificType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismIntraspecificDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organism.intraspecificDescription", nil, htmlAttrs)
	}
	return StringInput("organism.intraspecificDescription", resource.Organism.IntraspecificDescription, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorType(numAuthor int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Organism.Author) {
		return CodeableConceptSelect("organism.author["+strconv.Itoa(numAuthor)+"].authorType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.author["+strconv.Itoa(numAuthor)+"].authorType", resource.Organism.Author[numAuthor].AuthorType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismAuthorAuthorDescription(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Organism.Author) {
		return StringInput("organism.author["+strconv.Itoa(numAuthor)+"].authorDescription", nil, htmlAttrs)
	}
	return StringInput("organism.author["+strconv.Itoa(numAuthor)+"].authorDescription", resource.Organism.Author[numAuthor].AuthorDescription, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organism.hybrid.maternalOrganismId", nil, htmlAttrs)
	}
	return StringInput("organism.hybrid.maternalOrganismId", resource.Organism.Hybrid.MaternalOrganismId, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridMaternalOrganismName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organism.hybrid.maternalOrganismName", nil, htmlAttrs)
	}
	return StringInput("organism.hybrid.maternalOrganismName", resource.Organism.Hybrid.MaternalOrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organism.hybrid.paternalOrganismId", nil, htmlAttrs)
	}
	return StringInput("organism.hybrid.paternalOrganismId", resource.Organism.Hybrid.PaternalOrganismId, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridPaternalOrganismName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("organism.hybrid.paternalOrganismName", nil, htmlAttrs)
	}
	return StringInput("organism.hybrid.paternalOrganismName", resource.Organism.Hybrid.PaternalOrganismName, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismHybridHybridType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.hybrid.hybridType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.hybrid.hybridType", resource.Organism.Hybrid.HybridType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralKingdom(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.organismGeneral.kingdom", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.organismGeneral.kingdom", resource.Organism.OrganismGeneral.Kingdom, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralPhylum(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.organismGeneral.phylum", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.organismGeneral.phylum", resource.Organism.OrganismGeneral.Phylum, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralClass(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.organismGeneral.class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.organismGeneral.class", resource.Organism.OrganismGeneral.Class, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_OrganismOrganismGeneralOrder(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("organism.organismGeneral.order", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("organism.organismGeneral.order", resource.Organism.OrganismGeneral.Order, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPart(numPartDescription int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartDescription >= len(resource.PartDescription) {
		return CodeableConceptSelect("partDescription["+strconv.Itoa(numPartDescription)+"].part", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("partDescription["+strconv.Itoa(numPartDescription)+"].part", resource.PartDescription[numPartDescription].Part, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceSourceMaterial) T_PartDescriptionPartLocation(numPartDescription int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartDescription >= len(resource.PartDescription) {
		return CodeableConceptSelect("partDescription["+strconv.Itoa(numPartDescription)+"].partLocation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("partDescription["+strconv.Itoa(numPartDescription)+"].partLocation", resource.PartDescription[numPartDescription].PartLocation, optionsValueSet, htmlAttrs)
}
