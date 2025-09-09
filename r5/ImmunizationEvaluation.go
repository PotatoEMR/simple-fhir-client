package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
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
	Patient           Reference         `json:"patient"`
	Date              *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Authority         *Reference        `json:"authority,omitempty"`
	TargetDisease     CodeableConcept   `json:"targetDisease"`
	ImmunizationEvent Reference         `json:"immunizationEvent"`
	DoseStatus        CodeableConcept   `json:"doseStatus"`
	DoseStatusReason  []CodeableConcept `json:"doseStatusReason,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Series            *string           `json:"series,omitempty"`
	DoseNumber        *string           `json:"doseNumber,omitempty"`
	SeriesDoses       *string           `json:"seriesDoses,omitempty"`
}

type OtherImmunizationEvaluation ImmunizationEvaluation

// on convert struct to json, automatically add resourceType=ImmunizationEvaluation
func (r ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationEvaluation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationEvaluation: OtherImmunizationEvaluation(r),
		ResourceType:                "ImmunizationEvaluation",
	})
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
func (resource *ImmunizationEvaluation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSImmunization_evaluation_status

	if resource == nil {
		return CodeSelect("ImmunizationEvaluation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImmunizationEvaluation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ImmunizationEvaluation.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ImmunizationEvaluation.Date", resource.Date, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_TargetDisease(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ImmunizationEvaluation.TargetDisease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.TargetDisease", &resource.TargetDisease, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ImmunizationEvaluation.DoseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.DoseStatus", &resource.DoseStatus, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatusReason(numDoseStatusReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDoseStatusReason >= len(resource.DoseStatusReason) {
		return CodeableConceptSelect("ImmunizationEvaluation.DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", &resource.DoseStatusReason[numDoseStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ImmunizationEvaluation.Description", nil, htmlAttrs)
	}
	return StringInput("ImmunizationEvaluation.Description", resource.Description, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Series(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ImmunizationEvaluation.Series", nil, htmlAttrs)
	}
	return StringInput("ImmunizationEvaluation.Series", resource.Series, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ImmunizationEvaluation.DoseNumber", nil, htmlAttrs)
	}
	return StringInput("ImmunizationEvaluation.DoseNumber", resource.DoseNumber, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_SeriesDoses(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ImmunizationEvaluation.SeriesDoses", nil, htmlAttrs)
	}
	return StringInput("ImmunizationEvaluation.SeriesDoses", resource.SeriesDoses, htmlAttrs)
}
