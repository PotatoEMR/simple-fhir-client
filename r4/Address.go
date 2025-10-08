package r4

import "strings"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Address
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

func (a Address) String() string {
	parts := []string{}

	if a.Text != nil && *a.Text != "" {
		parts = append(parts, *a.Text)
	} else {
		if len(a.Line) > 0 {
			parts = append(parts, strings.Join(a.Line, ", "))
		}
		if a.City != nil && *a.City != "" {
			parts = append(parts, *a.City)
		}
		if a.District != nil && *a.District != "" {
			parts = append(parts, *a.District)
		}
		if a.State != nil && *a.State != "" {
			parts = append(parts, *a.State)
		}
		if a.PostalCode != nil && *a.PostalCode != "" {
			parts = append(parts, *a.PostalCode)
		}
		if a.Country != nil && *a.Country != "" {
			parts = append(parts, *a.Country)
		}
	}
	return strings.Join(parts, ", ")
}
