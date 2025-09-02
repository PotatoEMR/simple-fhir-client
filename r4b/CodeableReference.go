package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/CodeableReference
type CodeableReference struct {
	Id        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Concept   *CodeableConcept `json:"concept,omitempty"`
	Reference *Reference       `json:"reference,omitempty"`
}
