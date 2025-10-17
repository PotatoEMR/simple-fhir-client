package r5

//generated with command go run ./bultaoreune
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

// struct -> json, automatically add resourceType=Patient
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
func (r SubstanceReferenceInformation) ResourceType() string {
	return "SubstanceReferenceInformation"
}

func (resource *SubstanceReferenceInformation) T_Comment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneGeneSequenceOrigin(numGene int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGene >= len(resource.Gene) {
		return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].geneSequenceOrigin", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].geneSequenceOrigin", resource.Gene[numGene].GeneSequenceOrigin, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneGene(numGene int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGene >= len(resource.Gene) {
		return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].gene", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("gene["+strconv.Itoa(numGene)+"].gene", resource.Gene[numGene].Gene, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneSource(frs []FhirResource, numGene int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGene >= len(resource.Gene) || numSource >= len(resource.Gene[numGene].Source) {
		return ReferenceInput(frs, "gene["+strconv.Itoa(numGene)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "gene["+strconv.Itoa(numGene)+"].source["+strconv.Itoa(numSource)+"]", &resource.Gene[numGene].Source[numSource], htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneElementType(numGeneElement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGeneElement >= len(resource.GeneElement) {
		return CodeableConceptSelect("geneElement["+strconv.Itoa(numGeneElement)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("geneElement["+strconv.Itoa(numGeneElement)+"].type", resource.GeneElement[numGeneElement].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneElementElement(numGeneElement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGeneElement >= len(resource.GeneElement) {
		return IdentifierInput("geneElement["+strconv.Itoa(numGeneElement)+"].element", nil, htmlAttrs)
	}
	return IdentifierInput("geneElement["+strconv.Itoa(numGeneElement)+"].element", resource.GeneElement[numGeneElement].Element, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_GeneElementSource(frs []FhirResource, numGeneElement int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGeneElement >= len(resource.GeneElement) || numSource >= len(resource.GeneElement[numGeneElement].Source) {
		return ReferenceInput(frs, "geneElement["+strconv.Itoa(numGeneElement)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "geneElement["+strconv.Itoa(numGeneElement)+"].source["+strconv.Itoa(numSource)+"]", &resource.GeneElement[numGeneElement].Source[numSource], htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetTarget(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return IdentifierInput("target["+strconv.Itoa(numTarget)+"].target", nil, htmlAttrs)
	}
	return IdentifierInput("target["+strconv.Itoa(numTarget)+"].target", resource.Target[numTarget].Target, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetType(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].type", resource.Target[numTarget].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetInteraction(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].interaction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].interaction", resource.Target[numTarget].Interaction, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganism(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organism", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organism", resource.Target[numTarget].Organism, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetOrganismType(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organismType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].organismType", resource.Target[numTarget].OrganismType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountQuantity(numTarget int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return QuantityInput("target["+strconv.Itoa(numTarget)+"].amountQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("target["+strconv.Itoa(numTarget)+"].amountQuantity", resource.Target[numTarget].AmountQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountRange(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return RangeInput("target["+strconv.Itoa(numTarget)+"].amountRange", nil, htmlAttrs)
	}
	return RangeInput("target["+strconv.Itoa(numTarget)+"].amountRange", resource.Target[numTarget].AmountRange, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountString(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return StringInput("target["+strconv.Itoa(numTarget)+"].amountString", nil, htmlAttrs)
	}
	return StringInput("target["+strconv.Itoa(numTarget)+"].amountString", resource.Target[numTarget].AmountString, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetAmountType(numTarget int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].amountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("target["+strconv.Itoa(numTarget)+"].amountType", resource.Target[numTarget].AmountType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceReferenceInformation) T_TargetSource(frs []FhirResource, numTarget int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) || numSource >= len(resource.Target[numTarget].Source) {
		return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"].source["+strconv.Itoa(numSource)+"]", &resource.Target[numTarget].Source[numSource], htmlAttrs)
}
