package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Contributor
type Contributor struct {
	Id        *string         `json:"id,omitempty"`
	Extension []Extension     `json:"extension,omitempty"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Contact   []ContactDetail `json:"contact,omitempty"`
}
