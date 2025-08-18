//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/Timing
type Timing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Event             []string         `json:"event,omitempty"`
	Repeat            *Element         `json:"repeat,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}
