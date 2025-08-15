//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string     `json:"id,omitempty"`
	Meta              *Meta       `json:"meta,omitempty"`
	ImplicitRules     *string     `json:"implicitRules,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Text              *Narrative  `json:"text,omitempty"`
	Contained         []Resource  `json:"contained,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}
