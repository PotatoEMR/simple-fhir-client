package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/FormularyItem
type FormularyItem struct {
	Id                *string          `json:"id,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	ImplicitRules     *string          `json:"implicitRules,omitempty"`
	Language          *string          `json:"language,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	Contained         []Resource       `json:"contained,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `json:"identifier,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *string          `json:"status,omitempty"`
}

type OtherFormularyItem FormularyItem

// struct -> json, automatically add resourceType=Patient
func (r FormularyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFormularyItem
		ResourceType string `json:"resourceType"`
	}{
		OtherFormularyItem: OtherFormularyItem(r),
		ResourceType:       "FormularyItem",
	})
}

// json -> struct, first reject if resourceType != FormularyItem
func (r *FormularyItem) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "FormularyItem" {
		return errors.New("resourceType not FormularyItem")
	}
	return json.Unmarshal(data, (*OtherFormularyItem)(r))
}

func (r FormularyItem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "FormularyItem/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "FormularyItem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *FormularyItem) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *FormularyItem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFormularyitem_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
