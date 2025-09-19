package r5

//generated with command go run ./bultaoreune
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
	StartDate             *FhirDateTime          `json:"startDate,omitempty"`
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
	Date                  *FhirDateTime                   `json:"date,omitempty"`
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
func (r GenomicStudy) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "GenomicStudy/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "GenomicStudy"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *GenomicStudy) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGenomicstudy_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", &resource.Subject, htmlAttrs)
}
func (resource *GenomicStudy) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *GenomicStudy) T_StartDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("startDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("startDate", resource.StartDate, htmlAttrs)
}
func (resource *GenomicStudy) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *GenomicStudy) T_Referrer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("referrer", nil, htmlAttrs)
	}
	return ReferenceInput("referrer", resource.Referrer, htmlAttrs)
}
func (resource *GenomicStudy) T_Interpreter(numInterpreter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInterpreter >= len(resource.Interpreter) {
		return ReferenceInput("interpreter["+strconv.Itoa(numInterpreter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("interpreter["+strconv.Itoa(numInterpreter)+"]", &resource.Interpreter[numInterpreter], htmlAttrs)
}
func (resource *GenomicStudy) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *GenomicStudy) T_InstantiatesCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical", resource.InstantiatesCanonical, htmlAttrs)
}
func (resource *GenomicStudy) T_InstantiatesUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instantiatesUri", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *GenomicStudy) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *GenomicStudy) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisMethodType(numAnalysis int, numMethodType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numMethodType >= len(resource.Analysis[numAnalysis].MethodType) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].methodType["+strconv.Itoa(numMethodType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].methodType["+strconv.Itoa(numMethodType)+"]", &resource.Analysis[numAnalysis].MethodType[numMethodType], optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisChangeType(numAnalysis int, numChangeType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numChangeType >= len(resource.Analysis[numAnalysis].ChangeType) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].changeType["+strconv.Itoa(numChangeType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].changeType["+strconv.Itoa(numChangeType)+"]", &resource.Analysis[numAnalysis].ChangeType[numChangeType], optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisGenomeBuild(numAnalysis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].genomeBuild", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].genomeBuild", resource.Analysis[numAnalysis].GenomeBuild, optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInstantiatesCanonical(numAnalysis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].instantiatesCanonical", nil, htmlAttrs)
	}
	return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].instantiatesCanonical", resource.Analysis[numAnalysis].InstantiatesCanonical, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInstantiatesUri(numAnalysis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].instantiatesUri", nil, htmlAttrs)
	}
	return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].instantiatesUri", resource.Analysis[numAnalysis].InstantiatesUri, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisTitle(numAnalysis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].title", nil, htmlAttrs)
	}
	return StringInput("analysis["+strconv.Itoa(numAnalysis)+"].title", resource.Analysis[numAnalysis].Title, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisFocus(numAnalysis int, numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numFocus >= len(resource.Analysis[numAnalysis].Focus) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].focus["+strconv.Itoa(numFocus)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].focus["+strconv.Itoa(numFocus)+"]", &resource.Analysis[numAnalysis].Focus[numFocus], htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisSpecimen(numAnalysis int, numSpecimen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numSpecimen >= len(resource.Analysis[numAnalysis].Specimen) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].specimen["+strconv.Itoa(numSpecimen)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].specimen["+strconv.Itoa(numSpecimen)+"]", &resource.Analysis[numAnalysis].Specimen[numSpecimen], htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisDate(numAnalysis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return FhirDateTimeInput("analysis["+strconv.Itoa(numAnalysis)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("analysis["+strconv.Itoa(numAnalysis)+"].date", resource.Analysis[numAnalysis].Date, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisNote(numAnalysis int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numNote >= len(resource.Analysis[numAnalysis].Note) {
		return AnnotationTextArea("analysis["+strconv.Itoa(numAnalysis)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("analysis["+strconv.Itoa(numAnalysis)+"].note["+strconv.Itoa(numNote)+"]", &resource.Analysis[numAnalysis].Note[numNote], htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisProtocolPerformed(numAnalysis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].protocolPerformed", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].protocolPerformed", resource.Analysis[numAnalysis].ProtocolPerformed, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisRegionsStudied(numAnalysis int, numRegionsStudied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numRegionsStudied >= len(resource.Analysis[numAnalysis].RegionsStudied) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].regionsStudied["+strconv.Itoa(numRegionsStudied)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].regionsStudied["+strconv.Itoa(numRegionsStudied)+"]", &resource.Analysis[numAnalysis].RegionsStudied[numRegionsStudied], htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisRegionsCalled(numAnalysis int, numRegionsCalled int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numRegionsCalled >= len(resource.Analysis[numAnalysis].RegionsCalled) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].regionsCalled["+strconv.Itoa(numRegionsCalled)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].regionsCalled["+strconv.Itoa(numRegionsCalled)+"]", &resource.Analysis[numAnalysis].RegionsCalled[numRegionsCalled], htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInputFile(numAnalysis int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numInput >= len(resource.Analysis[numAnalysis].Input) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].file", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].file", resource.Analysis[numAnalysis].Input[numInput].File, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInputType(numAnalysis int, numInput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numInput >= len(resource.Analysis[numAnalysis].Input) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].type", resource.Analysis[numAnalysis].Input[numInput].Type, optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInputGeneratedByIdentifier(numAnalysis int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numInput >= len(resource.Analysis[numAnalysis].Input) {
		return IdentifierInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].generatedByIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].generatedByIdentifier", resource.Analysis[numAnalysis].Input[numInput].GeneratedByIdentifier, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisInputGeneratedByReference(numAnalysis int, numInput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numInput >= len(resource.Analysis[numAnalysis].Input) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].generatedByReference", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].input["+strconv.Itoa(numInput)+"].generatedByReference", resource.Analysis[numAnalysis].Input[numInput].GeneratedByReference, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisOutputFile(numAnalysis int, numOutput int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numOutput >= len(resource.Analysis[numAnalysis].Output) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].output["+strconv.Itoa(numOutput)+"].file", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].output["+strconv.Itoa(numOutput)+"].file", resource.Analysis[numAnalysis].Output[numOutput].File, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisOutputType(numAnalysis int, numOutput int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numOutput >= len(resource.Analysis[numAnalysis].Output) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].output["+strconv.Itoa(numOutput)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].output["+strconv.Itoa(numOutput)+"].type", resource.Analysis[numAnalysis].Output[numOutput].Type, optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisPerformerActor(numAnalysis int, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numPerformer >= len(resource.Analysis[numAnalysis].Performer) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].performer["+strconv.Itoa(numPerformer)+"].actor", resource.Analysis[numAnalysis].Performer[numPerformer].Actor, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisPerformerRole(numAnalysis int, numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numPerformer >= len(resource.Analysis[numAnalysis].Performer) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].performer["+strconv.Itoa(numPerformer)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].performer["+strconv.Itoa(numPerformer)+"].role", resource.Analysis[numAnalysis].Performer[numPerformer].Role, optionsValueSet, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisDeviceDevice(numAnalysis int, numDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numDevice >= len(resource.Analysis[numAnalysis].Device) {
		return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].device["+strconv.Itoa(numDevice)+"].device", nil, htmlAttrs)
	}
	return ReferenceInput("analysis["+strconv.Itoa(numAnalysis)+"].device["+strconv.Itoa(numDevice)+"].device", resource.Analysis[numAnalysis].Device[numDevice].Device, htmlAttrs)
}
func (resource *GenomicStudy) T_AnalysisDeviceFunction(numAnalysis int, numDevice int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAnalysis >= len(resource.Analysis) || numDevice >= len(resource.Analysis[numAnalysis].Device) {
		return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].device["+strconv.Itoa(numDevice)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("analysis["+strconv.Itoa(numAnalysis)+"].device["+strconv.Itoa(numDevice)+"].function", resource.Analysis[numAnalysis].Device[numDevice].Function, optionsValueSet, htmlAttrs)
}
