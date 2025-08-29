package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Age
type Age struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Value      *float64    `json:"value,omitempty"`
	Comparator *string     `json:"comparator,omitempty"`
	Unit       *string     `json:"unit,omitempty"`
	System     *string     `json:"system,omitempty"`
	Code       *string     `json:"code,omitempty"`
}
