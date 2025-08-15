//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	Type                   string           `json:"type"`
	Profile                []string         `json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept"`
	SubjectReference       *Reference       `json:"subjectReference"`
	MustSupport            []string         `json:"mustSupport,omitempty"`
	CodeFilter             []Element        `json:"codeFilter,omitempty"`
	DateFilter             []Element        `json:"dateFilter,omitempty"`
	Limit                  *int             `json:"limit,omitempty"`
	Sort                   []Element        `json:"sort,omitempty"`
}
