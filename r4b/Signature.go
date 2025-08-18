//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Signature
type Signature struct {
	Id           *string     `json:"id,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	Type         []Coding    `json:"type"`
	When         string      `json:"when"`
	Who          Reference   `json:"who"`
	OnBehalfOf   *Reference  `json:"onBehalfOf,omitempty"`
	TargetFormat *string     `json:"targetFormat,omitempty"`
	SigFormat    *string     `json:"sigFormat,omitempty"`
	Data         *string     `json:"data,omitempty"`
}
