package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r SubstanceReferenceInformation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceReferenceInformation/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubstanceReferenceInformation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceReferenceInformation) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneGeneSequenceOrigin(numGene int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGene >= len(resource.Gene) {
		return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].geneSequenceOrigin", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].geneSequenceOrigin", resource.Gene[numGene].GeneSequenceOrigin, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneGene(numGene int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGene >= len(resource.Gene) {
		return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].gene", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].gene", resource.Gene[numGene].Gene, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneElementType(numGeneElement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numGeneElement >= len(resource.GeneElement) {
		return CodeableConceptSelect("geneElement["+strconv.Itoa(numGeneElement)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("geneElement["+strconv.Itoa(numGeneElement)+"].type", resource.GeneElement[numGeneElement].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_ClassificationDomain(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].domain", resource.Classification[numClassification].Domain, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_ClassificationClassification(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].classification", resource.Classification[numClassification].Classification, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_ClassificationSubtype(numClassification int, numSubtype int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) || numSubtype >= len(resource.Classification[numClassification].Subtype) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].subtype["+strconv.Itoa(numSubtype)+"]", &resource.Classification[numClassification].Subtype[numSubtype], optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetType(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].type", resource.Target[numTarget].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetInteraction(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].interaction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].interaction", resource.Target[numTarget].Interaction, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganism(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organism", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organism", resource.Target[numTarget].Organism, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganismType(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organismType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organismType", resource.Target[numTarget].OrganismType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountString(numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return StringInput("target["+strconv.Itoa(numTarget)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("target["+strconv.Itoa(numTarget)+"].amountString", resource.Target[numTarget].AmountString, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountType(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].amountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].amountType", resource.Target[numTarget].AmountType, optionsValueSet, htmlAttrs)
}
