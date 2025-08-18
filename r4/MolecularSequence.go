//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequence struct {
	Id                *string                             `json:"id,omitempty"`
	Meta              *Meta                               `json:"meta,omitempty"`
	ImplicitRules     *string                             `json:"implicitRules,omitempty"`
	Language          *string                             `json:"language,omitempty"`
	Text              *Narrative                          `json:"text,omitempty"`
	Contained         []Resource                          `json:"contained,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                        `json:"identifier,omitempty"`
	Type              *string                             `json:"type,omitempty"`
	CoordinateSystem  int                                 `json:"coordinateSystem"`
	Patient           *Reference                          `json:"patient,omitempty"`
	Specimen          *Reference                          `json:"specimen,omitempty"`
	Device            *Reference                          `json:"device,omitempty"`
	Performer         *Reference                          `json:"performer,omitempty"`
	Quantity          *Quantity                           `json:"quantity,omitempty"`
	ReferenceSeq      *MolecularSequenceReferenceSeq      `json:"referenceSeq,omitempty"`
	Variant           []MolecularSequenceVariant          `json:"variant,omitempty"`
	ObservedSeq       *string                             `json:"observedSeq,omitempty"`
	Quality           []MolecularSequenceQuality          `json:"quality,omitempty"`
	ReadCoverage      *int                                `json:"readCoverage,omitempty"`
	Repository        []MolecularSequenceRepository       `json:"repository,omitempty"`
	Pointer           []Reference                         `json:"pointer,omitempty"`
	StructureVariant  []MolecularSequenceStructureVariant `json:"structureVariant,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceReferenceSeq struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Chromosome          *CodeableConcept `json:"chromosome,omitempty"`
	GenomeBuild         *string          `json:"genomeBuild,omitempty"`
	Orientation         *string          `json:"orientation,omitempty"`
	ReferenceSeqId      *CodeableConcept `json:"referenceSeqId,omitempty"`
	ReferenceSeqPointer *Reference       `json:"referenceSeqPointer,omitempty"`
	ReferenceSeqString  *string          `json:"referenceSeqString,omitempty"`
	Strand              *string          `json:"strand,omitempty"`
	WindowStart         *int             `json:"windowStart,omitempty"`
	WindowEnd           *int             `json:"windowEnd,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceVariant struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
	ObservedAllele    *string     `json:"observedAllele,omitempty"`
	ReferenceAllele   *string     `json:"referenceAllele,omitempty"`
	Cigar             *string     `json:"cigar,omitempty"`
	VariantPointer    *Reference  `json:"variantPointer,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceQuality struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Type              string                       `json:"type"`
	StandardSequence  *CodeableConcept             `json:"standardSequence,omitempty"`
	Start             *int                         `json:"start,omitempty"`
	End               *int                         `json:"end,omitempty"`
	Score             *Quantity                    `json:"score,omitempty"`
	Method            *CodeableConcept             `json:"method,omitempty"`
	TruthTP           *float64                     `json:"truthTP,omitempty"`
	QueryTP           *float64                     `json:"queryTP,omitempty"`
	TruthFN           *float64                     `json:"truthFN,omitempty"`
	QueryFP           *float64                     `json:"queryFP,omitempty"`
	GtFP              *float64                     `json:"gtFP,omitempty"`
	Precision         *float64                     `json:"precision,omitempty"`
	Recall            *float64                     `json:"recall,omitempty"`
	FScore            *float64                     `json:"fScore,omitempty"`
	Roc               *MolecularSequenceQualityRoc `json:"roc,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceQualityRoc struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Score             []int       `json:"score,omitempty"`
	NumTP             []int       `json:"numTP,omitempty"`
	NumFP             []int       `json:"numFP,omitempty"`
	NumFN             []int       `json:"numFN,omitempty"`
	Precision         []float64   `json:"precision,omitempty"`
	Sensitivity       []float64   `json:"sensitivity,omitempty"`
	FMeasure          []float64   `json:"fMeasure,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceRepository struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Url               *string     `json:"url,omitempty"`
	Name              *string     `json:"name,omitempty"`
	DatasetId         *string     `json:"datasetId,omitempty"`
	VariantsetId      *string     `json:"variantsetId,omitempty"`
	ReadsetId         *string     `json:"readsetId,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceStructureVariant struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	VariantType       *CodeableConcept                        `json:"variantType,omitempty"`
	Exact             *bool                                   `json:"exact,omitempty"`
	Length            *int                                    `json:"length,omitempty"`
	Outer             *MolecularSequenceStructureVariantOuter `json:"outer,omitempty"`
	Inner             *MolecularSequenceStructureVariantInner `json:"inner,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceStructureVariantOuter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MolecularSequence
type MolecularSequenceStructureVariantInner struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
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
