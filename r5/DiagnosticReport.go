//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	EffectiveDateTime  *string                          `json:"effectiveDateTime"`
	EffectivePeriod    *Period                          `json:"effectivePeriod"`
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
