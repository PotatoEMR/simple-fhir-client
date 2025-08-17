//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
