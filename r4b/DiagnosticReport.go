package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/DiagnosticReport
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
	EffectiveDateTime  *string                 `json:"effectiveDateTime,omitempty"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/DiagnosticReport
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

func (resource *DiagnosticReport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DiagnosticReport.Id", nil)
	}
	return StringInput("DiagnosticReport.Id", resource.Id)
}
func (resource *DiagnosticReport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DiagnosticReport.ImplicitRules", nil)
	}
	return StringInput("DiagnosticReport.ImplicitRules", resource.ImplicitRules)
}
func (resource *DiagnosticReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DiagnosticReport.Language", nil, optionsValueSet)
	}
	return CodeSelect("DiagnosticReport.Language", resource.Language, optionsValueSet)
}
func (resource *DiagnosticReport) T_Status() templ.Component {
	optionsValueSet := VSDiagnostic_report_status

	if resource == nil {
		return CodeSelect("DiagnosticReport.Status", nil, optionsValueSet)
	}
	return CodeSelect("DiagnosticReport.Status", &resource.Status, optionsValueSet)
}
func (resource *DiagnosticReport) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("DiagnosticReport.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DiagnosticReport.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *DiagnosticReport) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DiagnosticReport.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DiagnosticReport.Code", &resource.Code, optionsValueSet)
}
func (resource *DiagnosticReport) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("DiagnosticReport.Issued", nil)
	}
	return StringInput("DiagnosticReport.Issued", resource.Issued)
}
func (resource *DiagnosticReport) T_Conclusion() templ.Component {

	if resource == nil {
		return StringInput("DiagnosticReport.Conclusion", nil)
	}
	return StringInput("DiagnosticReport.Conclusion", resource.Conclusion)
}
func (resource *DiagnosticReport) T_ConclusionCode(numConclusionCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConclusionCode) >= numConclusionCode {
		return CodeableConceptSelect("DiagnosticReport.ConclusionCode["+strconv.Itoa(numConclusionCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DiagnosticReport.ConclusionCode["+strconv.Itoa(numConclusionCode)+"]", &resource.ConclusionCode[numConclusionCode], optionsValueSet)
}
func (resource *DiagnosticReport) T_MediaId(numMedia int) templ.Component {

	if resource == nil || len(resource.Media) >= numMedia {
		return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Id", nil)
	}
	return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Id", resource.Media[numMedia].Id)
}
func (resource *DiagnosticReport) T_MediaComment(numMedia int) templ.Component {

	if resource == nil || len(resource.Media) >= numMedia {
		return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Comment", nil)
	}
	return StringInput("DiagnosticReport.Media["+strconv.Itoa(numMedia)+"].Comment", resource.Media[numMedia].Comment)
}
