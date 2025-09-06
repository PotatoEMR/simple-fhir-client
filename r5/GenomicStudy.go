package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	GeneratedByIdentifier *Identifier      `json:"generatedByIdentifier,omitempty"`
	GeneratedByReference  *Reference       `json:"generatedByReference,omitempty"`
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

func (resource *GenomicStudy) T_Id() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.Id", nil)
	}
	return StringInput("GenomicStudy.Id", resource.Id)
}
func (resource *GenomicStudy) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.ImplicitRules", nil)
	}
	return StringInput("GenomicStudy.ImplicitRules", resource.ImplicitRules)
}
func (resource *GenomicStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("GenomicStudy.Language", nil, optionsValueSet)
	}
	return CodeSelect("GenomicStudy.Language", resource.Language, optionsValueSet)
}
func (resource *GenomicStudy) T_Status() templ.Component {
	optionsValueSet := VSGenomicstudy_status

	if resource == nil {
		return CodeSelect("GenomicStudy.Status", nil, optionsValueSet)
	}
	return CodeSelect("GenomicStudy.Status", &resource.Status, optionsValueSet)
}
func (resource *GenomicStudy) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("GenomicStudy.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *GenomicStudy) T_StartDate() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.StartDate", nil)
	}
	return StringInput("GenomicStudy.StartDate", resource.StartDate)
}
func (resource *GenomicStudy) T_InstantiatesCanonical() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.InstantiatesCanonical", nil)
	}
	return StringInput("GenomicStudy.InstantiatesCanonical", resource.InstantiatesCanonical)
}
func (resource *GenomicStudy) T_InstantiatesUri() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.InstantiatesUri", nil)
	}
	return StringInput("GenomicStudy.InstantiatesUri", resource.InstantiatesUri)
}
func (resource *GenomicStudy) T_Description() templ.Component {

	if resource == nil {
		return StringInput("GenomicStudy.Description", nil)
	}
	return StringInput("GenomicStudy.Description", resource.Description)
}
func (resource *GenomicStudy) T_AnalysisId(numAnalysis int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Id", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Id", resource.Analysis[numAnalysis].Id)
}
func (resource *GenomicStudy) T_AnalysisMethodType(numAnalysis int, numMethodType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].MethodType) >= numMethodType {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].MethodType["+strconv.Itoa(numMethodType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].MethodType["+strconv.Itoa(numMethodType)+"]", &resource.Analysis[numAnalysis].MethodType[numMethodType], optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisChangeType(numAnalysis int, numChangeType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].ChangeType) >= numChangeType {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].ChangeType["+strconv.Itoa(numChangeType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].ChangeType["+strconv.Itoa(numChangeType)+"]", &resource.Analysis[numAnalysis].ChangeType[numChangeType], optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisGenomeBuild(numAnalysis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].GenomeBuild", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].GenomeBuild", resource.Analysis[numAnalysis].GenomeBuild, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisInstantiatesCanonical(numAnalysis int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].InstantiatesCanonical", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].InstantiatesCanonical", resource.Analysis[numAnalysis].InstantiatesCanonical)
}
func (resource *GenomicStudy) T_AnalysisInstantiatesUri(numAnalysis int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].InstantiatesUri", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].InstantiatesUri", resource.Analysis[numAnalysis].InstantiatesUri)
}
func (resource *GenomicStudy) T_AnalysisTitle(numAnalysis int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Title", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Title", resource.Analysis[numAnalysis].Title)
}
func (resource *GenomicStudy) T_AnalysisDate(numAnalysis int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Date", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Date", resource.Analysis[numAnalysis].Date)
}
func (resource *GenomicStudy) T_AnalysisInputId(numAnalysis int, numInput int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Input) >= numInput {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Input["+strconv.Itoa(numInput)+"].Id", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Input["+strconv.Itoa(numInput)+"].Id", resource.Analysis[numAnalysis].Input[numInput].Id)
}
func (resource *GenomicStudy) T_AnalysisInputType(numAnalysis int, numInput int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Input) >= numInput {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Input["+strconv.Itoa(numInput)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Input["+strconv.Itoa(numInput)+"].Type", resource.Analysis[numAnalysis].Input[numInput].Type, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisOutputId(numAnalysis int, numOutput int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Output) >= numOutput {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Output["+strconv.Itoa(numOutput)+"].Id", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Output["+strconv.Itoa(numOutput)+"].Id", resource.Analysis[numAnalysis].Output[numOutput].Id)
}
func (resource *GenomicStudy) T_AnalysisOutputType(numAnalysis int, numOutput int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Output) >= numOutput {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Output["+strconv.Itoa(numOutput)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Output["+strconv.Itoa(numOutput)+"].Type", resource.Analysis[numAnalysis].Output[numOutput].Type, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisPerformerId(numAnalysis int, numPerformer int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Performer) >= numPerformer {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Analysis[numAnalysis].Performer[numPerformer].Id)
}
func (resource *GenomicStudy) T_AnalysisPerformerRole(numAnalysis int, numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Performer) >= numPerformer {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Performer["+strconv.Itoa(numPerformer)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Performer["+strconv.Itoa(numPerformer)+"].Role", resource.Analysis[numAnalysis].Performer[numPerformer].Role, optionsValueSet)
}
func (resource *GenomicStudy) T_AnalysisDeviceId(numAnalysis int, numDevice int) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Device) >= numDevice {
		return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Device["+strconv.Itoa(numDevice)+"].Id", nil)
	}
	return StringInput("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Device["+strconv.Itoa(numDevice)+"].Id", resource.Analysis[numAnalysis].Device[numDevice].Id)
}
func (resource *GenomicStudy) T_AnalysisDeviceFunction(numAnalysis int, numDevice int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Analysis) >= numAnalysis || len(resource.Analysis[numAnalysis].Device) >= numDevice {
		return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Device["+strconv.Itoa(numDevice)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GenomicStudy.Analysis["+strconv.Itoa(numAnalysis)+"].Device["+strconv.Itoa(numDevice)+"].Function", resource.Analysis[numAnalysis].Device[numDevice].Function, optionsValueSet)
}
