//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `json:"id,omitempty"`
	Extension       []Extension `json:"extension,omitempty"`
	AuthorReference *Reference  `json:"authorReference"`
	AuthorString    *string     `json:"authorString"`
	Time            *string     `json:"time,omitempty"`
	Text            string      `json:"text"`
}
