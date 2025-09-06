package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	Type                   string           `json:"type"`
	Profile                []string         `json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `json:"subjectReference,omitempty"`
	MustSupport            []string         `json:"mustSupport,omitempty"`
	CodeFilter             []Element        `json:"codeFilter,omitempty"`
	DateFilter             []Element        `json:"dateFilter,omitempty"`
	Limit                  *int             `json:"limit,omitempty"`
	Sort                   []Element        `json:"sort,omitempty"`
}
