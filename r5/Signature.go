//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Signature
type Signature struct {
	Id           *string     `json:"id,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	Type         []Coding    `json:"type,omitempty"`
	When         *string     `json:"when,omitempty"`
	Who          *Reference  `json:"who,omitempty"`
	OnBehalfOf   *Reference  `json:"onBehalfOf,omitempty"`
	TargetFormat *string     `json:"targetFormat,omitempty"`
	SigFormat    *string     `json:"sigFormat,omitempty"`
	Data         *string     `json:"data,omitempty"`
}
