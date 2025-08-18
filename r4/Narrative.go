//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/Narrative
type Narrative struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Status    string      `json:"status"`
	Div       string      `json:"div"`
}
