//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudy struct {
	Id                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []Resource             `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	Status                string                 `json:"status"`
	Type                  []CodeableConcept      `json:"type,omitempty"`
	Subject               Reference              `json:"subject"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	StartDate             *string                `json:"startDate,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	Referrer              *Reference             `json:"referrer,omitempty"`
	Interpreter           []Reference            `json:"interpreter,omitempty"`
	Reason                []CodeableReference    `json:"reason,omitempty"`
	InstantiatesCanonical *string                `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string                `json:"instantiatesUri,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	Description           *string                `json:"description,omitempty"`
	Analysis              []GenomicStudyAnalysis `json:"analysis,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudyAnalysis struct {
	Id                    *string                         `json:"id,omitempty"`
	Extension             []Extension                     `json:"extension,omitempty"`
	ModifierExtension     []Extension                     `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                    `json:"identifier,omitempty"`
	MethodType            []CodeableConcept               `json:"methodType,omitempty"`
	ChangeType            []CodeableConcept               `json:"changeType,omitempty"`
	GenomeBuild           *CodeableConcept                `json:"genomeBuild,omitempty"`
	InstantiatesCanonical *string                         `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string                         `json:"instantiatesUri,omitempty"`
	Title                 *string                         `json:"title,omitempty"`
	Focus                 []Reference                     `json:"focus,omitempty"`
	Specimen              []Reference                     `json:"specimen,omitempty"`
	Date                  *string                         `json:"date,omitempty"`
	Note                  []Annotation                    `json:"note,omitempty"`
	ProtocolPerformed     *Reference                      `json:"protocolPerformed,omitempty"`
	RegionsStudied        []Reference                     `json:"regionsStudied,omitempty"`
	RegionsCalled         []Reference                     `json:"regionsCalled,omitempty"`
	Input                 []GenomicStudyAnalysisInput     `json:"input,omitempty"`
	Output                []GenomicStudyAnalysisOutput    `json:"output,omitempty"`
	Performer             []GenomicStudyAnalysisPerformer `json:"performer,omitempty"`
	Device                []GenomicStudyAnalysisDevice    `json:"device,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudyAnalysisInput struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	File                  *Reference       `json:"file,omitempty"`
	Type                  *CodeableConcept `json:"type,omitempty"`
	GeneratedByIdentifier *Identifier      `json:"generatedByIdentifier"`
	GeneratedByReference  *Reference       `json:"generatedByReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudyAnalysisOutput struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	File              *Reference       `json:"file,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudyAnalysisPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Actor             *Reference       `json:"actor,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GenomicStudy
type GenomicStudyAnalysisDevice struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Device            *Reference       `json:"device,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
}

type OtherGenomicStudy GenomicStudy

// on convert struct to json, automatically add resourceType=GenomicStudy
func (r GenomicStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGenomicStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherGenomicStudy: OtherGenomicStudy(r),
		ResourceType:      "GenomicStudy",
	})
}
