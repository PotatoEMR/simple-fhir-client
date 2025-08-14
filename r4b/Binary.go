//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Binary
type Binary struct {
	Id              *string    `json:"id,omitempty"`
	Meta            *Meta      `json:"meta,omitempty"`
	ImplicitRules   *string    `json:"implicitRules,omitempty"`
	Language        *string    `json:"language,omitempty"`
	ContentType     string     `json:"contentType"`
	SecurityContext *Reference `json:"securityContext,omitempty"`
	Data            *string    `json:"data,omitempty"`
}
