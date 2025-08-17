//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Reference
type Reference struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Reference  *string     `json:"reference,omitempty"`
	Type       *string     `json:"type,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display    *string     `json:"display,omitempty"`
}
