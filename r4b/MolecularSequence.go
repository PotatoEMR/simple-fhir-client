package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
type MolecularSequenceStructureVariantOuter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MolecularSequence
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
		return CodeSelect("MolecularSequence.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MolecularSequence.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_CoordinateSystem(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.CoordinateSystem", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.CoordinateSystem", &resource.CoordinateSystem, htmlAttrs)
}
func (resource *MolecularSequence) T_ObservedSeq(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ObservedSeq", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.ObservedSeq", resource.ObservedSeq, htmlAttrs)
}
func (resource *MolecularSequence) T_ReadCoverage(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReadCoverage", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.ReadCoverage", resource.ReadCoverage, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqChromosome(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MolecularSequence.ReferenceSeq.Chromosome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MolecularSequence.ReferenceSeq.Chromosome", resource.ReferenceSeq.Chromosome, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqGenomeBuild(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ReferenceSeq.GenomeBuild", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.ReferenceSeq.GenomeBuild", resource.ReferenceSeq.GenomeBuild, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqOrientation(htmlAttrs string) templ.Component {
	optionsValueSet := VSOrientation_type

	if resource == nil {
		return CodeSelect("MolecularSequence.ReferenceSeq.Orientation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MolecularSequence.ReferenceSeq.Orientation", resource.ReferenceSeq.Orientation, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqReferenceSeqId(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MolecularSequence.ReferenceSeq.ReferenceSeqId", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MolecularSequence.ReferenceSeq.ReferenceSeqId", resource.ReferenceSeq.ReferenceSeqId, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqReferenceSeqString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MolecularSequence.ReferenceSeq.ReferenceSeqString", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.ReferenceSeq.ReferenceSeqString", resource.ReferenceSeq.ReferenceSeqString, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqStrand(htmlAttrs string) templ.Component {
	optionsValueSet := VSStrand_type

	if resource == nil {
		return CodeSelect("MolecularSequence.ReferenceSeq.Strand", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MolecularSequence.ReferenceSeq.Strand", resource.ReferenceSeq.Strand, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqWindowStart(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReferenceSeq.WindowStart", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.ReferenceSeq.WindowStart", resource.ReferenceSeq.WindowStart, htmlAttrs)
}
func (resource *MolecularSequence) T_ReferenceSeqWindowEnd(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("MolecularSequence.ReferenceSeq.WindowEnd", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.ReferenceSeq.WindowEnd", resource.ReferenceSeq.WindowEnd, htmlAttrs)
}
func (resource *MolecularSequence) T_VariantStart(numVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numVariant >= len(resource.Variant) {
		return IntInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..Start", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..Start", resource.Variant[numVariant].Start, htmlAttrs)
}
func (resource *MolecularSequence) T_VariantEnd(numVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numVariant >= len(resource.Variant) {
		return IntInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..End", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..End", resource.Variant[numVariant].End, htmlAttrs)
}
func (resource *MolecularSequence) T_VariantObservedAllele(numVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numVariant >= len(resource.Variant) {
		return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..ObservedAllele", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..ObservedAllele", resource.Variant[numVariant].ObservedAllele, htmlAttrs)
}
func (resource *MolecularSequence) T_VariantReferenceAllele(numVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numVariant >= len(resource.Variant) {
		return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..ReferenceAllele", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..ReferenceAllele", resource.Variant[numVariant].ReferenceAllele, htmlAttrs)
}
func (resource *MolecularSequence) T_VariantCigar(numVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numVariant >= len(resource.Variant) {
		return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..Cigar", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Variant."+strconv.Itoa(numVariant)+"..Cigar", resource.Variant[numVariant].Cigar, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityType(numQuality int, htmlAttrs string) templ.Component {
	optionsValueSet := VSQuality_type

	if resource == nil || numQuality >= len(resource.Quality) {
		return CodeSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Type", &resource.Quality[numQuality].Type, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityStandardSequence(numQuality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return CodeableConceptSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..StandardSequence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..StandardSequence", resource.Quality[numQuality].StandardSequence, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityStart(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Start", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Start", resource.Quality[numQuality].Start, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityEnd(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..End", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..End", resource.Quality[numQuality].End, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityMethod(numQuality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return CodeableConceptSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Method", resource.Quality[numQuality].Method, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityTruthTP(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..TruthTP", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..TruthTP", resource.Quality[numQuality].TruthTP, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityQueryTP(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..QueryTP", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..QueryTP", resource.Quality[numQuality].QueryTP, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityTruthFN(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..TruthFN", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..TruthFN", resource.Quality[numQuality].TruthFN, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityQueryFP(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..QueryFP", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..QueryFP", resource.Quality[numQuality].QueryFP, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityGtFP(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..GtFP", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..GtFP", resource.Quality[numQuality].GtFP, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityPrecision(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Precision", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Precision", resource.Quality[numQuality].Precision, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRecall(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Recall", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Recall", resource.Quality[numQuality].Recall, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityFScore(numQuality int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..FScore", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..FScore", resource.Quality[numQuality].FScore, htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocScore(numQuality int, numScore int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numScore >= len(resource.Quality[numQuality].Roc.Score) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Score."+strconv.Itoa(numScore)+".", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Score."+strconv.Itoa(numScore)+".", &resource.Quality[numQuality].Roc.Score[numScore], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocNumTP(numQuality int, numNumTP int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numNumTP >= len(resource.Quality[numQuality].Roc.NumTP) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumTP."+strconv.Itoa(numNumTP)+".", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumTP."+strconv.Itoa(numNumTP)+".", &resource.Quality[numQuality].Roc.NumTP[numNumTP], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocNumFP(numQuality int, numNumFP int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numNumFP >= len(resource.Quality[numQuality].Roc.NumFP) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumFP."+strconv.Itoa(numNumFP)+".", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumFP."+strconv.Itoa(numNumFP)+".", &resource.Quality[numQuality].Roc.NumFP[numNumFP], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocNumFN(numQuality int, numNumFN int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numNumFN >= len(resource.Quality[numQuality].Roc.NumFN) {
		return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumFN."+strconv.Itoa(numNumFN)+".", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.NumFN."+strconv.Itoa(numNumFN)+".", &resource.Quality[numQuality].Roc.NumFN[numNumFN], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocPrecision(numQuality int, numPrecision int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numPrecision >= len(resource.Quality[numQuality].Roc.Precision) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Precision."+strconv.Itoa(numPrecision)+".", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Precision."+strconv.Itoa(numPrecision)+".", &resource.Quality[numQuality].Roc.Precision[numPrecision], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocSensitivity(numQuality int, numSensitivity int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numSensitivity >= len(resource.Quality[numQuality].Roc.Sensitivity) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Sensitivity."+strconv.Itoa(numSensitivity)+".", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.Sensitivity."+strconv.Itoa(numSensitivity)+".", &resource.Quality[numQuality].Roc.Sensitivity[numSensitivity], htmlAttrs)
}
func (resource *MolecularSequence) T_QualityRocFMeasure(numQuality int, numFMeasure int, htmlAttrs string) templ.Component {

	if resource == nil || numQuality >= len(resource.Quality) || numFMeasure >= len(resource.Quality[numQuality].Roc.FMeasure) {
		return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.FMeasure."+strconv.Itoa(numFMeasure)+".", nil, htmlAttrs)
	}
	return Float64Input("MolecularSequence.Quality."+strconv.Itoa(numQuality)+"..Roc.FMeasure."+strconv.Itoa(numFMeasure)+".", &resource.Quality[numQuality].Roc.FMeasure[numFMeasure], htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryType(numRepository int, htmlAttrs string) templ.Component {
	optionsValueSet := VSRepository_type

	if resource == nil || numRepository >= len(resource.Repository) {
		return CodeSelect("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Type", &resource.Repository[numRepository].Type, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryUrl(numRepository int, htmlAttrs string) templ.Component {

	if resource == nil || numRepository >= len(resource.Repository) {
		return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Url", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Url", resource.Repository[numRepository].Url, htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryName(numRepository int, htmlAttrs string) templ.Component {

	if resource == nil || numRepository >= len(resource.Repository) {
		return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Name", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..Name", resource.Repository[numRepository].Name, htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryDatasetId(numRepository int, htmlAttrs string) templ.Component {

	if resource == nil || numRepository >= len(resource.Repository) {
		return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..DatasetId", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..DatasetId", resource.Repository[numRepository].DatasetId, htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryVariantsetId(numRepository int, htmlAttrs string) templ.Component {

	if resource == nil || numRepository >= len(resource.Repository) {
		return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..VariantsetId", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..VariantsetId", resource.Repository[numRepository].VariantsetId, htmlAttrs)
}
func (resource *MolecularSequence) T_RepositoryReadsetId(numRepository int, htmlAttrs string) templ.Component {

	if resource == nil || numRepository >= len(resource.Repository) {
		return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..ReadsetId", nil, htmlAttrs)
	}
	return StringInput("MolecularSequence.Repository."+strconv.Itoa(numRepository)+"..ReadsetId", resource.Repository[numRepository].ReadsetId, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantVariantType(numStructureVariant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return CodeableConceptSelect("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..VariantType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..VariantType", resource.StructureVariant[numStructureVariant].VariantType, optionsValueSet, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantExact(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return BoolInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Exact", nil, htmlAttrs)
	}
	return BoolInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Exact", resource.StructureVariant[numStructureVariant].Exact, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantLength(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Length", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Length", resource.StructureVariant[numStructureVariant].Length, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantOuterStart(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Outer.Start", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Outer.Start", resource.StructureVariant[numStructureVariant].Outer.Start, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantOuterEnd(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Outer.End", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Outer.End", resource.StructureVariant[numStructureVariant].Outer.End, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantInnerStart(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Inner.Start", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Inner.Start", resource.StructureVariant[numStructureVariant].Inner.Start, htmlAttrs)
}
func (resource *MolecularSequence) T_StructureVariantInnerEnd(numStructureVariant int, htmlAttrs string) templ.Component {

	if resource == nil || numStructureVariant >= len(resource.StructureVariant) {
		return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Inner.End", nil, htmlAttrs)
	}
	return IntInput("MolecularSequence.StructureVariant."+strconv.Itoa(numStructureVariant)+"..Inner.End", resource.StructureVariant[numStructureVariant].Inner.End, htmlAttrs)
}
