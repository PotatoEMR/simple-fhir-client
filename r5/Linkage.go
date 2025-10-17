package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Linkage
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

// http://hl7.org/fhir/r5/StructureDefinition/Linkage
type LinkageItem struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Resource          Reference   `json:"resource"`
}

type OtherLinkage Linkage

// struct -> json, automatically add resourceType=Patient
func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType"`
	}{
		OtherLinkage: OtherLinkage(r),
		ResourceType: "Linkage",
	})
}

// json -> struct, first reject if resourceType != Linkage
func (r *Linkage) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Linkage" {
		return errors.New("resourceType not Linkage")
	}
	return json.Unmarshal(data, (*OtherLinkage)(r))
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
func (resource *Linkage) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
func (resource *Linkage) T_ItemType(numItem int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSLinkage_type

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("item["+strconv.Itoa(numItem)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("item["+strconv.Itoa(numItem)+"].type", &resource.Item[numItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *Linkage) T_ItemResource(frs []FhirResource, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].resource", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].resource", &resource.Item[numItem].Resource, htmlAttrs)
}
