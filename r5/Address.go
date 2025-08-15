//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Address
type Address struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Use        *string     `json:"use,omitempty"`
	Type       *string     `json:"type,omitempty"`
	Text       *string     `json:"text,omitempty"`
	Line       []string    `json:"line,omitempty"`
	City       *string     `json:"city,omitempty"`
	District   *string     `json:"district,omitempty"`
	State      *string     `json:"state,omitempty"`
	PostalCode *string     `json:"postalCode,omitempty"`
	Country    *string     `json:"country,omitempty"`
	Period     *Period     `json:"period,omitempty"`
}
