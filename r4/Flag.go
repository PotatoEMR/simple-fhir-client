package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Flag
type Flag struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Status            string            `json:"status"`
	Category          []CodeableConcept `json:"category,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Subject           Reference         `json:"subject"`
	Period            *Period           `json:"period,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
}

type OtherFlag Flag

// struct -> json, automatically add resourceType=Patient
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

// json -> struct, first reject if resourceType != Flag
func (r *Flag) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Flag" {
		return errors.New("resourceType not Flag")
	}
	return json.Unmarshal(data, (*OtherFlag)(r))
}

func (r Flag) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Flag/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Flag"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Flag) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFlag_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Flag) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Flag) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Flag) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *Flag) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *Flag) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Flag) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
