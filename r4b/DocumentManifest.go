package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Created           *FhirDateTime             `json:"created,omitempty"`
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
func (resource *DocumentManifest) T_MasterIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("masterIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("masterIdentifier", resource.MasterIdentifier, htmlAttrs)
}
func (resource *DocumentManifest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DocumentManifest) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DocumentManifest) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", resource.Subject, htmlAttrs)
}
func (resource *DocumentManifest) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", resource.Created, htmlAttrs)
}
func (resource *DocumentManifest) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ReferenceInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *DocumentManifest) T_Recipient(numRecipient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecipient >= len(resource.Recipient) {
		return ReferenceInput("recipient["+strconv.Itoa(numRecipient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("recipient["+strconv.Itoa(numRecipient)+"]", &resource.Recipient[numRecipient], htmlAttrs)
}
func (resource *DocumentManifest) T_Source(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source", nil, htmlAttrs)
	}
	return StringInput("source", resource.Source, htmlAttrs)
}
func (resource *DocumentManifest) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *DocumentManifest) T_Content(numContent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return ReferenceInput("content["+strconv.Itoa(numContent)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("content["+strconv.Itoa(numContent)+"]", &resource.Content[numContent], htmlAttrs)
}
func (resource *DocumentManifest) T_RelatedRef(numRelated int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return ReferenceInput("related["+strconv.Itoa(numRelated)+"].ref", nil, htmlAttrs)
	}
	return ReferenceInput("related["+strconv.Itoa(numRelated)+"].ref", resource.Related[numRelated].Ref, htmlAttrs)
}
