package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Basic
type Basic struct {
	Id                *string         `json:"id,omitempty"`
	Meta              *Meta           `json:"meta,omitempty"`
	ImplicitRules     *string         `json:"implicitRules,omitempty"`
	Language          *string         `json:"language,omitempty"`
	Text              *Narrative      `json:"text,omitempty"`
	Contained         []Resource      `json:"contained,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Subject           *Reference      `json:"subject,omitempty"`
	Created           *string         `json:"created,omitempty"`
	Author            *Reference      `json:"author,omitempty"`
}

type OtherBasic Basic

// on convert struct to json, automatically add resourceType=Basic
func (r Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

func (resource *Basic) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Basic.Id", nil)
	}
	return StringInput("Basic.Id", resource.Id)
}
func (resource *Basic) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Basic.ImplicitRules", nil)
	}
	return StringInput("Basic.ImplicitRules", resource.ImplicitRules)
}
func (resource *Basic) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Basic.Language", nil, optionsValueSet)
	}
	return CodeSelect("Basic.Language", resource.Language, optionsValueSet)
}
func (resource *Basic) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Basic.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Basic.Code", &resource.Code, optionsValueSet)
}
func (resource *Basic) T_Created() templ.Component {

	if resource == nil {
		return StringInput("Basic.Created", nil)
	}
	return StringInput("Basic.Created", resource.Created)
}
