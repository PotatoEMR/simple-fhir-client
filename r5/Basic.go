package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Basic
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

func (resource *Basic) BasicLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
