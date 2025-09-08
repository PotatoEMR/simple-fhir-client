package r4b

import "time"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/Period
type Period struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Start     *time.Time  `json:"start,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	End       *time.Time  `json:"end,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}
