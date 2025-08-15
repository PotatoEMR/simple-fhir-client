//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/ContactPoint
type ContactPoint struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	System    *string     `json:"system,omitempty"`
	Value     *string     `json:"value,omitempty"`
	Use       *string     `json:"use,omitempty"`
	Rank      *int        `json:"rank,omitempty"`
	Period    *Period     `json:"period,omitempty"`
}
