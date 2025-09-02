package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	SequenceCodeableConcept *CodeableConcept `json:"sequenceCodeableConcept"`
	SequenceString          *string          `json:"sequenceString"`
	SequenceReference       *Reference       `json:"sequenceReference"`
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

func (resource *MolecularSequence) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MolecularSequence) T_Type() templ.Component {
	optionsValueSet := VSSequence_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Type, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeCoordinateSystem(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("coordinateSystem", nil, optionsValueSet)
	}
	return CodeableConceptSelect("coordinateSystem", &resource.Relative[numRelative].CoordinateSystem, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceGenomeAssembly(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("genomeAssembly", nil, optionsValueSet)
	}
	return CodeableConceptSelect("genomeAssembly", resource.Relative[numRelative].StartingSequence.GenomeAssembly, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceChromosome(numRelative int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Relative) >= numRelative {
		return CodeableConceptSelect("chromosome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("chromosome", resource.Relative[numRelative].StartingSequence.Chromosome, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceOrientation(numRelative int) templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil && len(resource.Relative) >= numRelative {
		return CodeSelect("orientation", nil, optionsValueSet)
	}
	return CodeSelect("orientation", resource.Relative[numRelative].StartingSequence.Orientation, optionsValueSet)
}
func (resource *MolecularSequence) T_RelativeStartingSequenceStrand(numRelative int) templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil && len(resource.Relative) >= numRelative {
		return CodeSelect("strand", nil, optionsValueSet)
	}
	return CodeSelect("strand", resource.Relative[numRelative].StartingSequence.Strand, optionsValueSet)
}
