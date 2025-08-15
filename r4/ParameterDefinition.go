//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	Id            *string     `json:"id,omitempty"`
	Extension     []Extension `json:"extension,omitempty"`
	Name          *string     `json:"name,omitempty"`
	Use           string      `json:"use"`
	Min           *int        `json:"min,omitempty"`
	Max           *string     `json:"max,omitempty"`
	Documentation *string     `json:"documentation,omitempty"`
	Type          string      `json:"type"`
	Profile       *string     `json:"profile,omitempty"`
}
