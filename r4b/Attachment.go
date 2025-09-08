package r4b

import "time"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/Attachment
type Attachment struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	ContentType *string     `json:"contentType,omitempty"`
	Language    *string     `json:"language,omitempty"`
	Data        *string     `json:"data,omitempty"`
	Url         *string     `json:"url,omitempty"`
	Size        *int        `json:"size,omitempty"`
	Hash        *string     `json:"hash,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Creation    *time.Time  `json:"creation,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}
