//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Identifier
type Identifier struct {
	Id        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Use       *string          `json:"use,omitempty"`
	Type      *CodeableConcept `json:"type,omitempty"`
	System    *string          `json:"system,omitempty"`
	Value     *string          `json:"value,omitempty"`
	Period    *Period          `json:"period,omitempty"`
	Assigner  *Reference       `json:"assigner,omitempty"`
}
