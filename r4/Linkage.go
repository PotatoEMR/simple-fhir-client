package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Linkage
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

// http://hl7.org/fhir/r4/StructureDefinition/Linkage
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
func (r Linkage) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Linkage/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "Linkage"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Linkage) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Linkage) T_Author(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("author", nil, htmlAttrs)
	}
	return ReferenceInput("author", resource.Author, htmlAttrs)
}
func (resource *Linkage) T_ItemType(numItem int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSLinkage_type

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("item["+strconv.Itoa(numItem)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("item["+strconv.Itoa(numItem)+"].type", &resource.Item[numItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *Linkage) T_ItemResource(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].resource", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].resource", &resource.Item[numItem].Resource, htmlAttrs)
}
