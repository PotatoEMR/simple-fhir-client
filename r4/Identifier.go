package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Identifier
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
