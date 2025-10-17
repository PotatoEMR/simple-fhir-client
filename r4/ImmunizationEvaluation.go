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

// http://hl7.org/fhir/r4/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
	Id                     *string           `json:"id,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []Resource        `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Identifier             []Identifier      `json:"identifier,omitempty"`
	Status                 string            `json:"status"`
	Patient                Reference         `json:"patient"`
	Date                   *FhirDateTime     `json:"date,omitempty"`
	Authority              *Reference        `json:"authority,omitempty"`
	TargetDisease          CodeableConcept   `json:"targetDisease"`
	ImmunizationEvent      Reference         `json:"immunizationEvent"`
	DoseStatus             CodeableConcept   `json:"doseStatus"`
	DoseStatusReason       []CodeableConcept `json:"doseStatusReason,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Series                 *string           `json:"series,omitempty"`
	DoseNumberPositiveInt  *int              `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString       *string           `json:"doseNumberString,omitempty"`
	SeriesDosesPositiveInt *int              `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string           `json:"seriesDosesString,omitempty"`
}

type OtherImmunizationEvaluation ImmunizationEvaluation

// struct -> json, automatically add resourceType=Patient
func (r ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationEvaluation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationEvaluation: OtherImmunizationEvaluation(r),
		ResourceType:                "ImmunizationEvaluation",
	})
}

// json -> struct, first reject if resourceType != ImmunizationEvaluation
func (r *ImmunizationEvaluation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "ImmunizationEvaluation" {
		return errors.New("resourceType not ImmunizationEvaluation")
	}
	return json.Unmarshal(data, (*OtherImmunizationEvaluation)(r))
}

func (r ImmunizationEvaluation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ImmunizationEvaluation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ImmunizationEvaluation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImmunizationEvaluation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSImmunization_evaluation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Authority(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "authority", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "authority", resource.Authority, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_TargetDisease(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("targetDisease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("targetDisease", &resource.TargetDisease, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_ImmunizationEvent(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "immunizationEvent", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "immunizationEvent", &resource.ImmunizationEvent, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("doseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("doseStatus", &resource.DoseStatus, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatusReason(numDoseStatusReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDoseStatusReason >= len(resource.DoseStatusReason) {
		return CodeableConceptSelect("doseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("doseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", &resource.DoseStatusReason[numDoseStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Series(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("series", nil, htmlAttrs)
	}
	return StringInput("series", resource.Series, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseNumberPositiveInt(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("doseNumberPositiveInt", nil, htmlAttrs)
	}
	return IntInput("doseNumberPositiveInt", resource.DoseNumberPositiveInt, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseNumberString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("doseNumberString", nil, htmlAttrs)
	}
	return StringInput("doseNumberString", resource.DoseNumberString, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_SeriesDosesPositiveInt(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("seriesDosesPositiveInt", nil, htmlAttrs)
	}
	return IntInput("seriesDosesPositiveInt", resource.SeriesDosesPositiveInt, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_SeriesDosesString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("seriesDosesString", nil, htmlAttrs)
	}
	return StringInput("seriesDosesString", resource.SeriesDosesString, htmlAttrs)
}
