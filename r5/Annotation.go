package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `json:"id,omitempty"`
	Extension       []Extension `json:"extension,omitempty"`
	AuthorReference *Reference  `json:"authorReference,omitempty"`
	AuthorString    *string     `json:"authorString,omitempty"`
	Time            *string     `json:"time,omitempty"`
	Text            string      `json:"text"`
}
