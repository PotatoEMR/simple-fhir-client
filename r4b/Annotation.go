//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `json:"id,omitempty"`
	Extension       []Extension `json:"extension,omitempty"`
	AuthorReference *Reference  `json:"authorReference"`
	AuthorString    *string     `json:"authorString"`
	Time            *string     `json:"time,omitempty"`
	Text            string      `json:"text"`
}
