//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/ContactPoint
type ContactPoint struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	System    *string     `json:"system,omitempty"`
	Value     *string     `json:"value,omitempty"`
	Use       *string     `json:"use,omitempty"`
	Rank      *int        `json:"rank,omitempty"`
	Period    *Period     `json:"period,omitempty"`
}
