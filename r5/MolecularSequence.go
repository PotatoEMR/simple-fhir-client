package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MolecularSequence
type MolecularSequence struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Type              *string                     `json:"type,omitempty"`
	Subject           *Reference                  `json:"subject,omitempty"`
	Focus             []Reference                 `json:"focus,omitempty"`
	Specimen          *Reference                  `json:"specimen,omitempty"`
	Device            *Reference                  `json:"device,omitempty"`
	Performer         *Reference                  `json:"performer,omitempty"`
	Literal           *string                     `json:"literal,omitempty"`
	Formatted         []Attachment                `json:"formatted,omitempty"`
	Relative          []MolecularSequenceRelative `json:"relative,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MolecularSequence
type MolecularSequenceRelative struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	CoordinateSystem  CodeableConcept                            `json:"coordinateSystem"`
	OrdinalPosition   *int                                       `json:"ordinalPosition,omitempty"`
	SequenceRange     *Range                                     `json:"sequenceRange,omitempty"`
	StartingSequence  *MolecularSequenceRelativeStartingSequence `json:"startingSequence,omitempty"`
	Edit              []MolecularSequenceRelativeEdit            `json:"edit,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MolecularSequence
type MolecularSequenceRelativeStartingSequence struct {
	Id                      *string          `json:"id,omitempty"`
	Extension               []Extension      `json:"extension,omitempty"`
	ModifierExtension       []Extension      `json:"modifierExtension,omitempty"`
	GenomeAssembly          *CodeableConcept `json:"genomeAssembly,omitempty"`
	Chromosome              *CodeableConcept `json:"chromosome,omitempty"`
	SequenceCodeableConcept *CodeableConcept `json:"sequenceCodeableConcept,omitempty"`
	SequenceString          *string          `json:"sequenceString,omitempty"`
	SequenceReference       *Reference       `json:"sequenceReference,omitempty"`
	WindowStart             *int             `json:"windowStart,omitempty"`
	WindowEnd               *int             `json:"windowEnd,omitempty"`
	Orientation             *string          `json:"orientation,omitempty"`
	Strand                  *string          `json:"strand,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MolecularSequence
type MolecularSequenceRelativeEdit struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Start               *int        `json:"start,omitempty"`
	End                 *int        `json:"end,omitempty"`
	ReplacementSequence *string     `json:"replacementSequence,omitempty"`
	ReplacedSequence    *string     `json:"replacedSequence,omitempty"`
}

type OtherMolecularSequence MolecularSequence

// on convert struct to json, automatically add resourceType=MolecularSequence
func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMolecularSequence
		ResourceType string `json:"resourceType"`
	}{
		OtherMolecularSequence: OtherMolecularSequence(r),
		ResourceType:           "MolecularSequence",
	})
}
func (r MolecularSequence) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MolecularSequence/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MolecularSequence"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MolecularSequence) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSSequence_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_Literal(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Literal", nil, htmlAttrs)
	}
	return StringInput("Literal", resource.Literal, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeCoordinateSystem(numRelative int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]CoordinateSystem", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]CoordinateSystem", &resource.Relative[numRelative].CoordinateSystem, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeOrdinalPosition(numRelative int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("Relative["+strconv.Itoa(numRelative)+"]OrdinalPosition", nil, htmlAttrs)
	}
	return IntInput("Relative["+strconv.Itoa(numRelative)+"]OrdinalPosition", resource.Relative[numRelative].OrdinalPosition, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceGenomeAssembly(numRelative int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.GenomeAssembly", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.GenomeAssembly", resource.Relative[numRelative].StartingSequence.GenomeAssembly, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceChromosome(numRelative int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Chromosome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Chromosome", resource.Relative[numRelative].StartingSequence.Chromosome, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceSequenceCodeableConcept(numRelative int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.SequenceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.SequenceCodeableConcept", resource.Relative[numRelative].StartingSequence.SequenceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceSequenceString(numRelative int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return StringInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.SequenceString", nil, htmlAttrs)
	}
	return StringInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.SequenceString", resource.Relative[numRelative].StartingSequence.SequenceString, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowStart(numRelative int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.WindowStart", nil, htmlAttrs)
	}
	return IntInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.WindowStart", resource.Relative[numRelative].StartingSequence.WindowStart, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowEnd(numRelative int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.WindowEnd", nil, htmlAttrs)
	}
	return IntInput("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.WindowEnd", resource.Relative[numRelative].StartingSequence.WindowEnd, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceOrientation(numRelative int, htmlAttrs string) templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Orientation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Orientation", resource.Relative[numRelative].StartingSequence.Orientation, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceStrand(numRelative int, htmlAttrs string) templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Strand", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Relative["+strconv.Itoa(numRelative)+"]StartingSequence.Strand", resource.Relative[numRelative].StartingSequence.Strand, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditStart(numRelative int, numEdit int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return IntInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].Start", nil, htmlAttrs)
	}
	return IntInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].Start", resource.Relative[numRelative].Edit[numEdit].Start, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditEnd(numRelative int, numEdit int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return IntInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].End", nil, htmlAttrs)
	}
	return IntInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].End", resource.Relative[numRelative].Edit[numEdit].End, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditReplacementSequence(numRelative int, numEdit int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return StringInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].ReplacementSequence", nil, htmlAttrs)
	}
	return StringInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].ReplacementSequence", resource.Relative[numRelative].Edit[numEdit].ReplacementSequence, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditReplacedSequence(numRelative int, numEdit int, htmlAttrs string) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return StringInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].ReplacedSequence", nil, htmlAttrs)
	}
	return StringInput("Relative["+strconv.Itoa(numRelative)+"]Edit["+strconv.Itoa(numEdit)+"].ReplacedSequence", resource.Relative[numRelative].Edit[numEdit].ReplacedSequence, htmlAttrs)
}
