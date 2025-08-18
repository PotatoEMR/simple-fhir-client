//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Distance
type Distance struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Value      *float64    `json:"value,omitempty"`
	Comparator *string     `json:"comparator,omitempty"`
	Unit       *string     `json:"unit,omitempty"`
	System     *string     `json:"system,omitempty"`
	Code       *string     `json:"code,omitempty"`
}
