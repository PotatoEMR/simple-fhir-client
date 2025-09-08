package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Availability
type Availability struct {
	Id               *string     `json:"id,omitempty"`
	Extension        []Extension `json:"extension,omitempty"`
	AvailableTime    []Element   `json:"availableTime,omitempty"`
	NotAvailableTime []Element   `json:"notAvailableTime,omitempty"`
}
