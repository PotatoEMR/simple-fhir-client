package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
	Id                *string                                       `json:"id,omitempty"`
	Meta              *Meta                                         `json:"meta,omitempty"`
	ImplicitRules     *string                                       `json:"implicitRules,omitempty"`
	Language          *string                                       `json:"language,omitempty"`
	Text              *Narrative                                    `json:"text,omitempty"`
	Contained         []Resource                                    `json:"contained,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Comment           *string                                       `json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene           `json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement    `json:"geneElement,omitempty"`
	Classification    []SubstanceReferenceInformationClassification `json:"classification,omitempty"`
	Target            []SubstanceReferenceInformationTarget         `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationGene struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `json:"gene,omitempty"`
	Source             []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationGeneElement struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Element           *Identifier      `json:"element,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Domain            *CodeableConcept  `json:"domain,omitempty"`
	Classification    *CodeableConcept  `json:"classification,omitempty"`
	Subtype           []CodeableConcept `json:"subtype,omitempty"`
	Source            []Reference       `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationTarget struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Target            *Identifier      `json:"target,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Interaction       *CodeableConcept `json:"interaction,omitempty"`
	Organism          *CodeableConcept `json:"organism,omitempty"`
	OrganismType      *CodeableConcept `json:"organismType,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity"`
	AmountRange       *Range           `json:"amountRange"`
	AmountString      *string          `json:"amountString"`
	AmountType        *CodeableConcept `json:"amountType,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

type OtherSubstanceReferenceInformation SubstanceReferenceInformation

// on convert struct to json, automatically add resourceType=SubstanceReferenceInformation
func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(r),
		ResourceType:                       "SubstanceReferenceInformation",
	})
}

func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationGeneGeneSequenceOrigin(numGene int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Gene) >= numGene {
		return CodeableConceptSelect("geneSequenceOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("geneSequenceOrigin", resource.Gene[numGene].GeneSequenceOrigin, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationGeneGene(numGene int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Gene) >= numGene {
		return CodeableConceptSelect("gene", nil, optionsValueSet)
	}
	return CodeableConceptSelect("gene", resource.Gene[numGene].Gene, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationGeneElementType(numGeneElement int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.GeneElement) >= numGeneElement {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.GeneElement[numGeneElement].Type, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationClassificationDomain(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", resource.Classification[numClassification].Domain, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationClassificationClassification(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classification", resource.Classification[numClassification].Classification, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationClassificationSubtype(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("subtype", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subtype", &resource.Classification[numClassification].Subtype[0], optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationTargetType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Target[numTarget].Type, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationTargetInteraction(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("interaction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("interaction", resource.Target[numTarget].Interaction, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationTargetOrganism(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("organism", nil, optionsValueSet)
	}
	return CodeableConceptSelect("organism", resource.Target[numTarget].Organism, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationTargetOrganismType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("organismType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("organismType", resource.Target[numTarget].OrganismType, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) SubstanceReferenceInformationTargetAmountType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Target) >= numTarget {
		return CodeableConceptSelect("amountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("amountType", resource.Target[numTarget].AmountType, optionsValueSet)
}
