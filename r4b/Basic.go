package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Basic
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
	Created           *time.Time      `json:"created,omitempty,format:'2006-01-02'"`
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
func (r Basic) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Basic/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Basic"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Basic) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Basic) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("Created", nil, htmlAttrs)
	}
	return DateInput("Created", resource.Created, htmlAttrs)
}
