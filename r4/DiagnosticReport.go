package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
	Id                 *string                 `json:"id,omitempty"`
	Meta               *Meta                   `json:"meta,omitempty"`
	ImplicitRules      *string                 `json:"implicitRules,omitempty"`
	Language           *string                 `json:"language,omitempty"`
	Text               *Narrative              `json:"text,omitempty"`
	Contained          []Resource              `json:"contained,omitempty"`
	Extension          []Extension             `json:"extension,omitempty"`
	ModifierExtension  []Extension             `json:"modifierExtension,omitempty"`
	Identifier         []Identifier            `json:"identifier,omitempty"`
	BasedOn            []Reference             `json:"basedOn,omitempty"`
	Status             string                  `json:"status"`
	Category           []CodeableConcept       `json:"category,omitempty"`
	Code               CodeableConcept         `json:"code"`
	Subject            *Reference              `json:"subject,omitempty"`
	Encounter          *Reference              `json:"encounter,omitempty"`
	EffectiveDateTime  *FhirDateTime           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod    *Period                 `json:"effectivePeriod,omitempty"`
	Issued             *string                 `json:"issued,omitempty"`
	Performer          []Reference             `json:"performer,omitempty"`
	ResultsInterpreter []Reference             `json:"resultsInterpreter,omitempty"`
	Specimen           []Reference             `json:"specimen,omitempty"`
	Result             []Reference             `json:"result,omitempty"`
	ImagingStudy       []Reference             `json:"imagingStudy,omitempty"`
	Media              []DiagnosticReportMedia `json:"media,omitempty"`
	Conclusion         *string                 `json:"conclusion,omitempty"`
	ConclusionCode     []CodeableConcept       `json:"conclusionCode,omitempty"`
	PresentedForm      []Attachment            `json:"presentedForm,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DiagnosticReport
type DiagnosticReportMedia struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Link              Reference   `json:"link"`
}

type OtherDiagnosticReport DiagnosticReport

// struct -> json, automatically add resourceType=Patient
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
}

func (r DiagnosticReport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DiagnosticReport/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DiagnosticReport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r DiagnosticReport) ResourceType() string {
	return "DiagnosticReport"
}

func (resource *DiagnosticReport) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DiagnosticReport) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDiagnostic_report_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *DiagnosticReport) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *DiagnosticReport) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *DiagnosticReport) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *DiagnosticReport) T_Issued(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("issued", nil, htmlAttrs)
	}
	return StringInput("issued", resource.Issued, htmlAttrs)
}
func (resource *DiagnosticReport) T_Performer(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", &resource.Performer[numPerformer], htmlAttrs)
}
func (resource *DiagnosticReport) T_ResultsInterpreter(frs []FhirResource, numResultsInterpreter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResultsInterpreter >= len(resource.ResultsInterpreter) {
		return ReferenceInput(frs, "resultsInterpreter["+strconv.Itoa(numResultsInterpreter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "resultsInterpreter["+strconv.Itoa(numResultsInterpreter)+"]", &resource.ResultsInterpreter[numResultsInterpreter], htmlAttrs)
}
func (resource *DiagnosticReport) T_Specimen(frs []FhirResource, numSpecimen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecimen >= len(resource.Specimen) {
		return ReferenceInput(frs, "specimen["+strconv.Itoa(numSpecimen)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "specimen["+strconv.Itoa(numSpecimen)+"]", &resource.Specimen[numSpecimen], htmlAttrs)
}
func (resource *DiagnosticReport) T_Result(frs []FhirResource, numResult int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResult >= len(resource.Result) {
		return ReferenceInput(frs, "result["+strconv.Itoa(numResult)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "result["+strconv.Itoa(numResult)+"]", &resource.Result[numResult], htmlAttrs)
}
func (resource *DiagnosticReport) T_ImagingStudy(frs []FhirResource, numImagingStudy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImagingStudy >= len(resource.ImagingStudy) {
		return ReferenceInput(frs, "imagingStudy["+strconv.Itoa(numImagingStudy)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "imagingStudy["+strconv.Itoa(numImagingStudy)+"]", &resource.ImagingStudy[numImagingStudy], htmlAttrs)
}
func (resource *DiagnosticReport) T_Conclusion(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("conclusion", nil, htmlAttrs)
	}
	return StringInput("conclusion", resource.Conclusion, htmlAttrs)
}
func (resource *DiagnosticReport) T_ConclusionCode(numConclusionCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConclusionCode >= len(resource.ConclusionCode) {
		return CodeableConceptSelect("conclusionCode["+strconv.Itoa(numConclusionCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conclusionCode["+strconv.Itoa(numConclusionCode)+"]", &resource.ConclusionCode[numConclusionCode], optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_PresentedForm(numPresentedForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPresentedForm >= len(resource.PresentedForm) {
		return AttachmentInput("presentedForm["+strconv.Itoa(numPresentedForm)+"]", nil, htmlAttrs)
	}
	return AttachmentInput("presentedForm["+strconv.Itoa(numPresentedForm)+"]", &resource.PresentedForm[numPresentedForm], htmlAttrs)
}
func (resource *DiagnosticReport) T_MediaComment(numMedia int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedia >= len(resource.Media) {
		return StringInput("media["+strconv.Itoa(numMedia)+"].comment", nil, htmlAttrs)
	}
	return StringInput("media["+strconv.Itoa(numMedia)+"].comment", resource.Media[numMedia].Comment, htmlAttrs)
}
func (resource *DiagnosticReport) T_MediaLink(frs []FhirResource, numMedia int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedia >= len(resource.Media) {
		return ReferenceInput(frs, "media["+strconv.Itoa(numMedia)+"].link", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "media["+strconv.Itoa(numMedia)+"].link", &resource.Media[numMedia].Link, htmlAttrs)
}
