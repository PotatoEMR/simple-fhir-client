package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *MolecularSequence) T_CoordinateSystem() templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.CoordinateSystem", nil)
	}
	return IntInput("MolecularSequence.CoordinateSystem", &resource.CoordinateSystem)
}
func (resource *MolecularSequence) T_ObservedSeq() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ObservedSeq", nil)
	}
	return StringInput("MolecularSequence.ObservedSeq", resource.ObservedSeq)
}
func (resource *MolecularSequence) T_ReadCoverage() templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReadCoverage", nil)
	}
	return IntInput("MolecularSequence.ReadCoverage", resource.ReadCoverage)
}
func (resource *MolecularSequence) T_ReferenceSeqId() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ReferenceSeq.Id", nil)
	}
	return StringInput("MolecularSequence.ReferenceSeq.Id", resource.ReferenceSeq.Id)
}
func (resource *MolecularSequence) T_ReferenceSeqChromosome(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MolecularSequence.ReferenceSeq.Chromosome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.ReferenceSeq.Chromosome", resource.ReferenceSeq.Chromosome, optionsValueSet)
}
func (resource *MolecularSequence) T_ReferenceSeqGenomeBuild() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ReferenceSeq.GenomeBuild", nil)
	}
	return StringInput("MolecularSequence.ReferenceSeq.GenomeBuild", resource.ReferenceSeq.GenomeBuild)
}
func (resource *MolecularSequence) T_ReferenceSeqOrientation() templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil {
		return CodeSelect("MolecularSequence.ReferenceSeq.Orientation", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.ReferenceSeq.Orientation", resource.ReferenceSeq.Orientation, optionsValueSet)
}
func (resource *MolecularSequence) T_ReferenceSeqReferenceSeqId(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MolecularSequence.ReferenceSeq.ReferenceSeqId", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.ReferenceSeq.ReferenceSeqId", resource.ReferenceSeq.ReferenceSeqId, optionsValueSet)
}
func (resource *MolecularSequence) T_ReferenceSeqReferenceSeqString() templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ReferenceSeq.ReferenceSeqString", nil)
	}
	return StringInput("MolecularSequence.ReferenceSeq.ReferenceSeqString", resource.ReferenceSeq.ReferenceSeqString)
}
func (resource *MolecularSequence) T_ReferenceSeqStrand() templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil {
		return CodeSelect("MolecularSequence.ReferenceSeq.Strand", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.ReferenceSeq.Strand", resource.ReferenceSeq.Strand, optionsValueSet)
}
func (resource *MolecularSequence) T_ReferenceSeqWindowStart() templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReferenceSeq.WindowStart", nil)
	}
	return IntInput("MolecularSequence.ReferenceSeq.WindowStart", resource.ReferenceSeq.WindowStart)
}
func (resource *MolecularSequence) T_ReferenceSeqWindowEnd() templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReferenceSeq.WindowEnd", nil)
	}
	return IntInput("MolecularSequence.ReferenceSeq.WindowEnd", resource.ReferenceSeq.WindowEnd)
}
func (resource *MolecularSequence) T_VariantId(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Id", nil)
	}
	return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Id", resource.Variant[numVariant].Id)
}
func (resource *MolecularSequence) T_VariantStart(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return IntInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Start", nil)
	}
	return IntInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Start", resource.Variant[numVariant].Start)
}
func (resource *MolecularSequence) T_VariantEnd(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return IntInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].End", nil)
	}
	return IntInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].End", resource.Variant[numVariant].End)
}
func (resource *MolecularSequence) T_VariantObservedAllele(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].ObservedAllele", nil)
	}
	return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].ObservedAllele", resource.Variant[numVariant].ObservedAllele)
}
func (resource *MolecularSequence) T_VariantReferenceAllele(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].ReferenceAllele", nil)
	}
	return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].ReferenceAllele", resource.Variant[numVariant].ReferenceAllele)
}
func (resource *MolecularSequence) T_VariantCigar(numVariant int) templ.Component {

	if resource == nil || len(resource.Variant) >= numVariant {
		return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Cigar", nil)
	}
	return StringInput("MolecularSequence.Variant["+strconv.Itoa(numVariant)+"].Cigar", resource.Variant[numVariant].Cigar)
}
func (resource *MolecularSequence) T_QualityId(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return StringInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Id", nil)
	}
	return StringInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Id", resource.Quality[numQuality].Id)
}
func (resource *MolecularSequence) T_QualityType(numQuality int) templ.Component {
	optionsValueSet := VSQuality_type

	if resource == nil || len(resource.Quality) >= numQuality {
		return CodeSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Type", &resource.Quality[numQuality].Type, optionsValueSet)
}
func (resource *MolecularSequence) T_QualityStandardSequence(numQuality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return CodeableConceptSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].StandardSequence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].StandardSequence", resource.Quality[numQuality].StandardSequence, optionsValueSet)
}
func (resource *MolecularSequence) T_QualityStart(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Start", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Start", resource.Quality[numQuality].Start)
}
func (resource *MolecularSequence) T_QualityEnd(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].End", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].End", resource.Quality[numQuality].End)
}
func (resource *MolecularSequence) T_QualityMethod(numQuality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return CodeableConceptSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Method", resource.Quality[numQuality].Method, optionsValueSet)
}
func (resource *MolecularSequence) T_QualityTruthTP(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].TruthTP", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].TruthTP", resource.Quality[numQuality].TruthTP)
}
func (resource *MolecularSequence) T_QualityQueryTP(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].QueryTP", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].QueryTP", resource.Quality[numQuality].QueryTP)
}
func (resource *MolecularSequence) T_QualityTruthFN(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].TruthFN", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].TruthFN", resource.Quality[numQuality].TruthFN)
}
func (resource *MolecularSequence) T_QualityQueryFP(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].QueryFP", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].QueryFP", resource.Quality[numQuality].QueryFP)
}
func (resource *MolecularSequence) T_QualityGtFP(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].GtFP", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].GtFP", resource.Quality[numQuality].GtFP)
}
func (resource *MolecularSequence) T_QualityPrecision(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Precision", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Precision", resource.Quality[numQuality].Precision)
}
func (resource *MolecularSequence) T_QualityRecall(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Recall", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Recall", resource.Quality[numQuality].Recall)
}
func (resource *MolecularSequence) T_QualityFScore(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].FScore", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].FScore", resource.Quality[numQuality].FScore)
}
func (resource *MolecularSequence) T_QualityRocId(numQuality int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality {
		return StringInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Id", nil)
	}
	return StringInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Id", resource.Quality[numQuality].Roc.Id)
}
func (resource *MolecularSequence) T_QualityRocScore(numQuality int, numScore int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.Score) >= numScore {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Score["+strconv.Itoa(numScore)+"]", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Score["+strconv.Itoa(numScore)+"]", &resource.Quality[numQuality].Roc.Score[numScore])
}
func (resource *MolecularSequence) T_QualityRocNumTP(numQuality int, numNumTP int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.NumTP) >= numNumTP {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumTP["+strconv.Itoa(numNumTP)+"]", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumTP["+strconv.Itoa(numNumTP)+"]", &resource.Quality[numQuality].Roc.NumTP[numNumTP])
}
func (resource *MolecularSequence) T_QualityRocNumFP(numQuality int, numNumFP int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.NumFP) >= numNumFP {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumFP["+strconv.Itoa(numNumFP)+"]", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumFP["+strconv.Itoa(numNumFP)+"]", &resource.Quality[numQuality].Roc.NumFP[numNumFP])
}
func (resource *MolecularSequence) T_QualityRocNumFN(numQuality int, numNumFN int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.NumFN) >= numNumFN {
		return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumFN["+strconv.Itoa(numNumFN)+"]", nil)
	}
	return IntInput("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.NumFN["+strconv.Itoa(numNumFN)+"]", &resource.Quality[numQuality].Roc.NumFN[numNumFN])
}
func (resource *MolecularSequence) T_QualityRocPrecision(numQuality int, numPrecision int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.Precision) >= numPrecision {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Precision["+strconv.Itoa(numPrecision)+"]", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Precision["+strconv.Itoa(numPrecision)+"]", &resource.Quality[numQuality].Roc.Precision[numPrecision])
}
func (resource *MolecularSequence) T_QualityRocSensitivity(numQuality int, numSensitivity int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.Sensitivity) >= numSensitivity {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Sensitivity["+strconv.Itoa(numSensitivity)+"]", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.Sensitivity["+strconv.Itoa(numSensitivity)+"]", &resource.Quality[numQuality].Roc.Sensitivity[numSensitivity])
}
func (resource *MolecularSequence) T_QualityRocFMeasure(numQuality int, numFMeasure int) templ.Component {

	if resource == nil || len(resource.Quality) >= numQuality || len(resource.Quality[numQuality].Roc.FMeasure) >= numFMeasure {
		return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.FMeasure["+strconv.Itoa(numFMeasure)+"]", nil)
	}
	return Float64Input("MolecularSequence.Quality["+strconv.Itoa(numQuality)+"].Roc.FMeasure["+strconv.Itoa(numFMeasure)+"]", &resource.Quality[numQuality].Roc.FMeasure[numFMeasure])
}
func (resource *MolecularSequence) T_RepositoryId(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Id", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Id", resource.Repository[numRepository].Id)
}
func (resource *MolecularSequence) T_RepositoryType(numRepository int) templ.Component {
	optionsValueSet := VSRepository_type

	if resource == nil || len(resource.Repository) >= numRepository {
		return CodeSelect("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Type", &resource.Repository[numRepository].Type, optionsValueSet)
}
func (resource *MolecularSequence) T_RepositoryUrl(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Url", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Url", resource.Repository[numRepository].Url)
}
func (resource *MolecularSequence) T_RepositoryName(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Name", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].Name", resource.Repository[numRepository].Name)
}
func (resource *MolecularSequence) T_RepositoryDatasetId(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].DatasetId", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].DatasetId", resource.Repository[numRepository].DatasetId)
}
func (resource *MolecularSequence) T_RepositoryVariantsetId(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].VariantsetId", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].VariantsetId", resource.Repository[numRepository].VariantsetId)
}
func (resource *MolecularSequence) T_RepositoryReadsetId(numRepository int) templ.Component {

	if resource == nil || len(resource.Repository) >= numRepository {
		return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].ReadsetId", nil)
	}
	return StringInput("MolecularSequence.Repository["+strconv.Itoa(numRepository)+"].ReadsetId", resource.Repository[numRepository].ReadsetId)
}
func (resource *MolecularSequence) T_StructureVariantId(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Id", nil)
	}
	return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Id", resource.StructureVariant[numStructureVariant].Id)
}
func (resource *MolecularSequence) T_StructureVariantVariantType(numStructureVariant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return CodeableConceptSelect("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].VariantType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].VariantType", resource.StructureVariant[numStructureVariant].VariantType, optionsValueSet)
}
func (resource *MolecularSequence) T_StructureVariantExact(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return BoolInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Exact", nil)
	}
	return BoolInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Exact", resource.StructureVariant[numStructureVariant].Exact)
}
func (resource *MolecularSequence) T_StructureVariantLength(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Length", nil)
	}
	return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Length", resource.StructureVariant[numStructureVariant].Length)
}
func (resource *MolecularSequence) T_StructureVariantOuterId(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.Id", nil)
	}
	return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.Id", resource.StructureVariant[numStructureVariant].Outer.Id)
}
func (resource *MolecularSequence) T_StructureVariantOuterStart(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.Start", nil)
	}
	return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.Start", resource.StructureVariant[numStructureVariant].Outer.Start)
}
func (resource *MolecularSequence) T_StructureVariantOuterEnd(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.End", nil)
	}
	return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Outer.End", resource.StructureVariant[numStructureVariant].Outer.End)
}
func (resource *MolecularSequence) T_StructureVariantInnerId(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.Id", nil)
	}
	return StringInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.Id", resource.StructureVariant[numStructureVariant].Inner.Id)
}
func (resource *MolecularSequence) T_StructureVariantInnerStart(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.Start", nil)
	}
	return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.Start", resource.StructureVariant[numStructureVariant].Inner.Start)
}
func (resource *MolecularSequence) T_StructureVariantInnerEnd(numStructureVariant int) templ.Component {

	if resource == nil || len(resource.StructureVariant) >= numStructureVariant {
		return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.End", nil)
	}
	return IntInput("MolecularSequence.StructureVariant["+strconv.Itoa(numStructureVariant)+"].Inner.End", resource.StructureVariant[numStructureVariant].Inner.End)
}
