//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Type      string      `json:"type"`
	Label     *string     `json:"label,omitempty"`
	Display   *string     `json:"display,omitempty"`
	Citation  *string     `json:"citation,omitempty"`
	Url       *string     `json:"url,omitempty"`
	Document  *Attachment `json:"document,omitempty"`
	Resource  *string     `json:"resource,omitempty"`
}
