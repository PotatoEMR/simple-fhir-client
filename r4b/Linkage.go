package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *Linkage) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Linkage.Id", nil)
	}
	return StringInput("Linkage.Id", resource.Id)
}
func (resource *Linkage) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Linkage.ImplicitRules", nil)
	}
	return StringInput("Linkage.ImplicitRules", resource.ImplicitRules)
}
func (resource *Linkage) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Linkage.Language", nil, optionsValueSet)
	}
	return CodeSelect("Linkage.Language", resource.Language, optionsValueSet)
}
func (resource *Linkage) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Linkage.Active", nil)
	}
	return BoolInput("Linkage.Active", resource.Active)
}
func (resource *Linkage) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Linkage.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("Linkage.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *Linkage) T_ItemType(numItem int) templ.Component {
	optionsValueSet := VSLinkage_type

	if resource == nil || len(resource.Item) >= numItem {
		return CodeSelect("Linkage.Item["+strconv.Itoa(numItem)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Linkage.Item["+strconv.Itoa(numItem)+"].Type", &resource.Item[numItem].Type, optionsValueSet)
}
