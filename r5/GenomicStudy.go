package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *GenomicStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *GenomicStudy) T_Status() templ.Component {
	optionsValueSet := VSGenomicstudy_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *GenomicStudy) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisMethodType(numAnalysis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis) >= numAnalysis {
		return CodeableConceptSelect("methodType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("methodType", &resource.Analysis[numAnalysis].MethodType[0], optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisChangeType(numAnalysis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis) >= numAnalysis {
		return CodeableConceptSelect("changeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("changeType", &resource.Analysis[numAnalysis].ChangeType[0], optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisGenomeBuild(numAnalysis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis) >= numAnalysis {
		return CodeableConceptSelect("genomeBuild", nil, optionsValueSet)
	}
	return CodeableConceptSelect("genomeBuild", resource.Analysis[numAnalysis].GenomeBuild, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisInputType(numAnalysis int, numInput int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis[numAnalysis].Input) >= numInput {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Analysis[numAnalysis].Input[numInput].Type, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisOutputType(numAnalysis int, numOutput int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis[numAnalysis].Output) >= numOutput {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Analysis[numAnalysis].Output[numOutput].Type, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisPerformerRole(numAnalysis int, numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis[numAnalysis].Performer) >= numPerformer {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Analysis[numAnalysis].Performer[numPerformer].Role, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisDeviceFunction(numAnalysis int, numDevice int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Analysis[numAnalysis].Device) >= numDevice {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Analysis[numAnalysis].Device[numDevice].Function, optionsValueSet)
}
