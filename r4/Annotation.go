package r4

import "time"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `json:"id,omitempty"`
	Extension       []Extension `json:"extension,omitempty"`
	AuthorReference *Reference  `json:"authorReference,omitempty"`
	AuthorString    *string     `json:"authorString,omitempty"`
	Time            *time.Time  `json:"time,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Text            string      `json:"text"`
}
