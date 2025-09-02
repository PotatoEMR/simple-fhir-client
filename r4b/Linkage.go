package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Linkage
type Linkage struct {
	Id                *string       `json:"id,omitempty"`
	Meta              *Meta         `json:"meta,omitempty"`
	ImplicitRules     *string       `json:"implicitRules,omitempty"`
	Language          *string       `json:"language,omitempty"`
	Text              *Narrative    `json:"text,omitempty"`
	Contained         []Resource    `json:"contained,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Active            *bool         `json:"active,omitempty"`
	Author            *Reference    `json:"author,omitempty"`
	Item              []LinkageItem `json:"item"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Linkage
type LinkageItem struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Resource          Reference   `json:"resource"`
}

type OtherLinkage Linkage

// on convert struct to json, automatically add resourceType=Linkage
func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType"`
	}{
		OtherLinkage: OtherLinkage(r),
		ResourceType: "Linkage",
	})
}

func (resource *Linkage) LinkageLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Linkage) LinkageItemType(numItem int) templ.Component {
	optionsValueSet := VSLinkage_type

	if resource == nil && len(resource.Item) >= numItem {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Item[numItem].Type, optionsValueSet)
}
