package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"

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
	Created           *FhirDate       `json:"created,omitempty"`
	Author            *Reference      `json:"author,omitempty"`
}

type OtherBasic Basic

// struct -> json, automatically add resourceType=Patient
func (r Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

// json -> struct, first reject if resourceType != Basic
func (r *Basic) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Basic" {
		return errors.New("resourceType not Basic")
	}
	return json.Unmarshal(data, (*OtherBasic)(r))
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
func (resource *Basic) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Basic) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Basic) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("created", nil, htmlAttrs)
	}
	return FhirDateInput("created", resource.Created, htmlAttrs)
}
func (resource *Basic) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
