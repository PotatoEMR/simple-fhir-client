package r4b

//generated with command go run ./bultaoreune -nodownload
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
	Created           *string                   `json:"created,omitempty"`
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

func (resource *DocumentManifest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Id", nil)
	}
	return StringInput("DocumentManifest.Id", resource.Id)
}
func (resource *DocumentManifest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.ImplicitRules", nil)
	}
	return StringInput("DocumentManifest.ImplicitRules", resource.ImplicitRules)
}
func (resource *DocumentManifest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DocumentManifest.Language", nil, optionsValueSet)
	}
	return CodeSelect("DocumentManifest.Language", resource.Language, optionsValueSet)
}
func (resource *DocumentManifest) T_Status() templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("DocumentManifest.Status", nil, optionsValueSet)
	}
	return CodeSelect("DocumentManifest.Status", &resource.Status, optionsValueSet)
}
func (resource *DocumentManifest) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentManifest.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentManifest.Type", resource.Type, optionsValueSet)
}
func (resource *DocumentManifest) T_Created() templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Created", nil)
	}
	return StringInput("DocumentManifest.Created", resource.Created)
}
func (resource *DocumentManifest) T_Source() templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Source", nil)
	}
	return StringInput("DocumentManifest.Source", resource.Source)
}
func (resource *DocumentManifest) T_Description() templ.Component {

	if resource == nil {
		return StringInput("DocumentManifest.Description", nil)
	}
	return StringInput("DocumentManifest.Description", resource.Description)
}
func (resource *DocumentManifest) T_RelatedId(numRelated int) templ.Component {

	if resource == nil || len(resource.Related) >= numRelated {
		return StringInput("DocumentManifest.Related["+strconv.Itoa(numRelated)+"].Id", nil)
	}
	return StringInput("DocumentManifest.Related["+strconv.Itoa(numRelated)+"].Id", resource.Related[numRelated].Id)
}
