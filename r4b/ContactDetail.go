package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/ContactDetail
type ContactDetail struct {
	Id        *string        `json:"id,omitempty"`
	Extension []Extension    `json:"extension,omitempty"`
	Name      *string        `json:"name,omitempty"`
	Telecom   []ContactPoint `json:"telecom,omitempty"`
}
