//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/DocumentManifest
type DocumentManifest struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier               `json:"masterIdentifier,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Status            string                    `json:"status"`
	Type              *CodeableConcept          `json:"type,omitempty"`
	Subject           *Reference                `json:"subject,omitempty"`
	Created           *string                   `json:"created,omitempty"`
	Author            []Reference               `json:"author,omitempty"`
	Recipient         []Reference               `json:"recipient,omitempty"`
	Source            *string                   `json:"source,omitempty"`
	Description       *string                   `json:"description,omitempty"`
	Content           []Reference               `json:"content"`
	Related           []DocumentManifestRelated `json:"related,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DocumentManifest
type DocumentManifestRelated struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Ref               *Reference  `json:"ref,omitempty"`
}

type OtherDocumentManifest DocumentManifest

// on convert struct to json, automatically add resourceType=DocumentManifest
func (r DocumentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentManifest: OtherDocumentManifest(r),
		ResourceType:          "DocumentManifest",
	})
}
