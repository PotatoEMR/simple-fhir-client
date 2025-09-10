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
func (resource *MolecularSequence) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSequence_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_Literal(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("literal", nil, htmlAttrs)
	}
	return StringInput("literal", resource.Literal, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeCoordinateSystem(numRelative int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].coordinateSystem", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].coordinateSystem", &resource.Relative[numRelative].CoordinateSystem, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeOrdinalPosition(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("relative["+strconv.Itoa(numRelative)+"].ordinalPosition", nil, htmlAttrs)
	}
	return IntInput("relative["+strconv.Itoa(numRelative)+"].ordinalPosition", resource.Relative[numRelative].OrdinalPosition, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceGenomeAssembly(numRelative int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.genomeAssembly", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.genomeAssembly", resource.Relative[numRelative].StartingSequence.GenomeAssembly, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceChromosome(numRelative int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.chromosome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.chromosome", resource.Relative[numRelative].StartingSequence.Chromosome, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceSequenceCodeableConcept(numRelative int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.sequenceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.sequenceCodeableConcept", resource.Relative[numRelative].StartingSequence.SequenceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceSequenceString(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return StringInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.sequenceString", nil, htmlAttrs)
	}
	return StringInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.sequenceString", resource.Relative[numRelative].StartingSequence.SequenceString, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowStart(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.windowStart", nil, htmlAttrs)
	}
	return IntInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.windowStart", resource.Relative[numRelative].StartingSequence.WindowStart, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowEnd(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) {
		return IntInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.windowEnd", nil, htmlAttrs)
	}
	return IntInput("relative["+strconv.Itoa(numRelative)+"].startingSequence.windowEnd", resource.Relative[numRelative].StartingSequence.WindowEnd, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceOrientation(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.orientation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.orientation", resource.Relative[numRelative].StartingSequence.Orientation, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceStrand(numRelative int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil || numRelative >= len(resource.Relative) {
		return CodeSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.strand", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relative["+strconv.Itoa(numRelative)+"].startingSequence.strand", resource.Relative[numRelative].StartingSequence.Strand, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditStart(numRelative int, numEdit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return IntInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].start", nil, htmlAttrs)
	}
	return IntInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].start", resource.Relative[numRelative].Edit[numEdit].Start, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditEnd(numRelative int, numEdit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return IntInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].end", nil, htmlAttrs)
	}
	return IntInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].end", resource.Relative[numRelative].Edit[numEdit].End, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditReplacementSequence(numRelative int, numEdit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return StringInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].replacementSequence", nil, htmlAttrs)
	}
	return StringInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].replacementSequence", resource.Relative[numRelative].Edit[numEdit].ReplacementSequence, htmlAttrs)
}
func (resource *MolecularSequence) T_RelativeEditReplacedSequence(numRelative int, numEdit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelative >= len(resource.Relative) || numEdit >= len(resource.Relative[numRelative].Edit) {
		return StringInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].replacedSequence", nil, htmlAttrs)
	}
	return StringInput("relative["+strconv.Itoa(numRelative)+"].edit["+strconv.Itoa(numEdit)+"].replacedSequence", resource.Relative[numRelative].Edit[numEdit].ReplacedSequence, htmlAttrs)
}
