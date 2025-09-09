package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
	Id                 *string                          `json:"id,omitempty"`
	Meta               *Meta                            `json:"meta,omitempty"`
	ImplicitRules      *string                          `json:"implicitRules,omitempty"`
	Language           *string                          `json:"language,omitempty"`
	Text               *Narrative                       `json:"text,omitempty"`
	Contained          []Resource                       `json:"contained,omitempty"`
	Extension          []Extension                      `json:"extension,omitempty"`
	ModifierExtension  []Extension                      `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                     `json:"identifier,omitempty"`
	BasedOn            []Reference                      `json:"basedOn,omitempty"`
	Status             string                           `json:"status"`
	Category           []CodeableConcept                `json:"category,omitempty"`
	Code               CodeableConcept                  `json:"code"`
	Subject            *Reference                       `json:"subject,omitempty"`
	Encounter          *Reference                       `json:"encounter,omitempty"`
	EffectiveDateTime  *time.Time                       `json:"effectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	EffectivePeriod    *Period                          `json:"effectivePeriod,omitempty"`
	Issued             *string                          `json:"issued,omitempty"`
	Performer          []Reference                      `json:"performer,omitempty"`
	ResultsInterpreter []Reference                      `json:"resultsInterpreter,omitempty"`
	Specimen           []Reference                      `json:"specimen,omitempty"`
	Result             []Reference                      `json:"result,omitempty"`
	Note               []Annotation                     `json:"note,omitempty"`
	Study              []Reference                      `json:"study,omitempty"`
	SupportingInfo     []DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
	Media              []DiagnosticReportMedia          `json:"media,omitempty"`
	Composition        *Reference                       `json:"composition,omitempty"`
	Conclusion         *string                          `json:"conclusion,omitempty"`
	ConclusionCode     []CodeableConcept                `json:"conclusionCode,omitempty"`
	PresentedForm      []Attachment                     `json:"presentedForm,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DiagnosticReport
type DiagnosticReportSupportingInfo struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Reference         Reference       `json:"reference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DiagnosticReport
type DiagnosticReportMedia struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Link              Reference   `json:"link"`
}

type OtherDiagnosticReport DiagnosticReport

// on convert struct to json, automatically add resourceType=DiagnosticReport
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
func (resource *DiagnosticReport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDiagnostic_report_status

	if resource == nil {
		return CodeSelect("DiagnosticReport.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DiagnosticReport.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("DiagnosticReport.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DiagnosticReport.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DiagnosticReport.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DiagnosticReport.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_EffectiveDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("DiagnosticReport.EffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DiagnosticReport.EffectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *DiagnosticReport) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DiagnosticReport.Issued", nil, htmlAttrs)
	}
	return StringInput("DiagnosticReport.Issued", resource.Issued, htmlAttrs)
}
func (resource *DiagnosticReport) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DiagnosticReport.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("DiagnosticReport.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DiagnosticReport) T_Conclusion(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DiagnosticReport.Conclusion", nil, htmlAttrs)
	}
	return StringInput("DiagnosticReport.Conclusion", resource.Conclusion, htmlAttrs)
}
func (resource *DiagnosticReport) T_ConclusionCode(numConclusionCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConclusionCode >= len(resource.ConclusionCode) {
		return CodeableConceptSelect("DiagnosticReport.ConclusionCode["+strconv.Itoa(numConclusionCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DiagnosticReport.ConclusionCode["+strconv.Itoa(numConclusionCode)+"]", &resource.ConclusionCode[numConclusionCode], optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_SupportingInfoType(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("DiagnosticReport.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DiagnosticReport.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Type", &resource.SupportingInfo[numSupportingInfo].Type, optionsValueSet, htmlAttrs)
}
func (resource *DiagnosticReport) T_MediaComment(numMedia int, htmlAttrs string) templ.Component {
	if resource == nil || numMedia >= len(resource.Media) {
		return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Comment", resource.Media[numMedia].Comment, htmlAttrs)
}
