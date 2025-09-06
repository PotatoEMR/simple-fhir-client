package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *MolecularSequence) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.Id", nil)
	}
	return StringInput("MolecularSequence.Id", resource.Id)
}
func (resource *MolecularSequence) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ImplicitRules", nil)
	}
	return StringInput("MolecularSequence.ImplicitRules", resource.ImplicitRules)
}
func (resource *MolecularSequence) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MolecularSequence.Language", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Language", resource.Language, optionsValueSet)
}
func (resource *MolecularSequence) T_Type() templ.Component {
	optionsValueSet := VSSequence_type

	if resource == nil {
		return CodeSelect("MolecularSequence.Type", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Type", resource.Type, optionsValueSet)
}
func (resource *MolecularSequence) T_Literal() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.Literal", nil)
	}
	return StringInput("MolecularSequence.Literal", resource.Literal)
}
func (resource *MolecularSequence) T_RelativeId(numRelative int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Id", nil)
	}
	return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Id", resource.Relative[numRelative].Id)
}
func (resource *MolecularSequence) T_RelativeCoordinateSystem(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].CoordinateSystem", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].CoordinateSystem", &resource.Relative[numRelative].CoordinateSystem, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeOrdinalPosition(numRelative int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].OrdinalPosition", nil)
	}
	return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].OrdinalPosition", resource.Relative[numRelative].OrdinalPosition)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceId(numRelative int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Id", nil)
	}
	return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Id", resource.Relative[numRelative].StartingSequence.Id)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceGenomeAssembly(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.GenomeAssembly", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.GenomeAssembly", resource.Relative[numRelative].StartingSequence.GenomeAssembly, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceChromosome(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Chromosome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Chromosome", resource.Relative[numRelative].StartingSequence.Chromosome, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowStart(numRelative int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.WindowStart", nil)
	}
	return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.WindowStart", resource.Relative[numRelative].StartingSequence.WindowStart)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceWindowEnd(numRelative int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative {
		return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.WindowEnd", nil)
	}
	return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.WindowEnd", resource.Relative[numRelative].StartingSequence.WindowEnd)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceOrientation(numRelative int) templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil || len(resource.Relative) >= numRelative {
		return CodeSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Orientation", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Orientation", resource.Relative[numRelative].StartingSequence.Orientation, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceStrand(numRelative int) templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil || len(resource.Relative) >= numRelative {
		return CodeSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Strand", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].StartingSequence.Strand", resource.Relative[numRelative].StartingSequence.Strand, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeEditId(numRelative int, numEdit int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative || len(resource.Relative[numRelative].Edit) >= numEdit {
		return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].Id", nil)
	}
	return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].Id", resource.Relative[numRelative].Edit[numEdit].Id)
}
func (resource *MolecularSequence) T_RelativeEditStart(numRelative int, numEdit int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative || len(resource.Relative[numRelative].Edit) >= numEdit {
		return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].Start", nil)
	}
	return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].Start", resource.Relative[numRelative].Edit[numEdit].Start)
}
func (resource *MolecularSequence) T_RelativeEditEnd(numRelative int, numEdit int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative || len(resource.Relative[numRelative].Edit) >= numEdit {
		return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].End", nil)
	}
	return IntInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].End", resource.Relative[numRelative].Edit[numEdit].End)
}
func (resource *MolecularSequence) T_RelativeEditReplacementSequence(numRelative int, numEdit int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative || len(resource.Relative[numRelative].Edit) >= numEdit {
		return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].ReplacementSequence", nil)
	}
	return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].ReplacementSequence", resource.Relative[numRelative].Edit[numEdit].ReplacementSequence)
}
func (resource *MolecularSequence) T_RelativeEditReplacedSequence(numRelative int, numEdit int) templ.Component {

	if resource == nil || len(resource.Relative) >= numRelative || len(resource.Relative[numRelative].Edit) >= numEdit {
		return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].ReplacedSequence", nil)
	}
	return StringInput("MolecularSequence.Relative["+strconv.Itoa(numRelative)+"].Edit["+strconv.Itoa(numEdit)+"].ReplacedSequence", resource.Relative[numRelative].Edit[numEdit].ReplacedSequence)
}
