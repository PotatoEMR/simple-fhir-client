package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/DocumentManifest
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
	Created           *time.Time                `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Author            []Reference               `json:"author,omitempty"`
	Recipient         []Reference               `json:"recipient,omitempty"`
	Source            *string                   `json:"source,omitempty"`
	Description       *string                   `json:"description,omitempty"`
	Content           []Reference               `json:"content"`
	Related           []DocumentManifestRelated `json:"related,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DocumentManifest
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
func (r DocumentManifest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DocumentManifest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DocumentManifest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DocumentManifest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("DocumentManifest.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DocumentManifest.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DocumentManifest) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentManifest.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DocumentManifest.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DocumentManifest) T_Created(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DocumentManifest.Created", nil, htmlAttrs)
	}
	return DateTimeInput("DocumentManifest.Created", resource.Created, htmlAttrs)
}
func (resource *DocumentManifest) T_Source(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Source", nil, htmlAttrs)
	}
	return StringInput("DocumentManifest.Source", resource.Source, htmlAttrs)
}
func (resource *DocumentManifest) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Description", nil, htmlAttrs)
	}
	return StringInput("DocumentManifest.Description", resource.Description, htmlAttrs)
}
