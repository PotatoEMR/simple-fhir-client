//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Meta
type Meta struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	VersionId   *string     `json:"versionId,omitempty"`
	LastUpdated *string     `json:"lastUpdated,omitempty"`
	Source      *string     `json:"source,omitempty"`
	Profile     []string    `json:"profile,omitempty"`
	Security    []Coding    `json:"security,omitempty"`
	Tag         []Coding    `json:"tag,omitempty"`
}
