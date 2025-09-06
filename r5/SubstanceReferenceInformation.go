package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
	Id                *string                                    `json:"id,omitempty"`
	Meta              *Meta                                      `json:"meta,omitempty"`
	ImplicitRules     *string                                    `json:"implicitRules,omitempty"`
	Language          *string                                    `json:"language,omitempty"`
	Text              *Narrative                                 `json:"text,omitempty"`
	Contained         []Resource                                 `json:"contained,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Comment           *string                                    `json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene        `json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement `json:"geneElement,omitempty"`
	Target            []SubstanceReferenceInformationTarget      `json:"target,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationGene struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `json:"gene,omitempty"`
	Source             []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationGeneElement struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Element           *Identifier      `json:"element,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationTarget struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Target            *Identifier      `json:"target,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Interaction       *CodeableConcept `json:"interaction,omitempty"`
	Organism          *CodeableConcept `json:"organism,omitempty"`
	OrganismType      *CodeableConcept `json:"organismType,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity,omitempty"`
	AmountRange       *Range           `json:"amountRange,omitempty"`
	AmountString      *string          `json:"amountString,omitempty"`
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

func (resource *SubstanceReferenceInformation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceReferenceInformation.Id", nil)
	}
	return StringInput("SubstanceReferenceInformation.Id", resource.Id)
}
func (resource *SubstanceReferenceInformation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceReferenceInformation.ImplicitRules", nil)
	}
	return StringInput("SubstanceReferenceInformation.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceReferenceInformation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceReferenceInformation.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceReferenceInformation.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("SubstanceReferenceInformation.Comment", nil)
	}
	return StringInput("SubstanceReferenceInformation.Comment", resource.Comment)
}
func (resource *SubstanceReferenceInformation) T_GeneId(numGene int) templ.Component {

	if resource == nil || len(resource.Gene) >= numGene {
		return StringInput("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].Id", nil)
	}
	return StringInput("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].Id", resource.Gene[numGene].Id)
}
func (resource *SubstanceReferenceInformation) T_GeneGeneSequenceOrigin(numGene int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Gene) >= numGene {
		return CodeableConceptSelect("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].GeneSequenceOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].GeneSequenceOrigin", resource.Gene[numGene].GeneSequenceOrigin, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_GeneGene(numGene int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Gene) >= numGene {
		return CodeableConceptSelect("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].Gene", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Gene["+strconv.Itoa(numGene)+"].Gene", resource.Gene[numGene].Gene, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_GeneElementId(numGeneElement int) templ.Component {

	if resource == nil || len(resource.GeneElement) >= numGeneElement {
		return StringInput("SubstanceReferenceInformation.GeneElement["+strconv.Itoa(numGeneElement)+"].Id", nil)
	}
	return StringInput("SubstanceReferenceInformation.GeneElement["+strconv.Itoa(numGeneElement)+"].Id", resource.GeneElement[numGeneElement].Id)
}
func (resource *SubstanceReferenceInformation) T_GeneElementType(numGeneElement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.GeneElement) >= numGeneElement {
		return CodeableConceptSelect("SubstanceReferenceInformation.GeneElement["+strconv.Itoa(numGeneElement)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.GeneElement["+strconv.Itoa(numGeneElement)+"].Type", resource.GeneElement[numGeneElement].Type, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_TargetId(numTarget int) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return StringInput("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Id", resource.Target[numTarget].Id)
}
func (resource *SubstanceReferenceInformation) T_TargetType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Type", resource.Target[numTarget].Type, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_TargetInteraction(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Interaction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Interaction", resource.Target[numTarget].Interaction, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganism(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Organism", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].Organism", resource.Target[numTarget].Organism, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganismType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].OrganismType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].OrganismType", resource.Target[numTarget].OrganismType, optionsValueSet)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountType(numTarget int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].AmountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceReferenceInformation.Target["+strconv.Itoa(numTarget)+"].AmountType", resource.Target[numTarget].AmountType, optionsValueSet)
}
