//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Distance
type Distance struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Value      *float64    `json:"value,omitempty"`
	Comparator *string     `json:"comparator,omitempty"`
	Unit       *string     `json:"unit,omitempty"`
	System     *string     `json:"system,omitempty"`
	Code       *string     `json:"code,omitempty"`
}
