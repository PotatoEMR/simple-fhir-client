//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Availability
type Availability struct {
	Id               *string     `json:"id,omitempty"`
	Extension        []Extension `json:"extension,omitempty"`
	AvailableTime    []Element   `json:"availableTime,omitempty"`
	NotAvailableTime []Element   `json:"notAvailableTime,omitempty"`
}
