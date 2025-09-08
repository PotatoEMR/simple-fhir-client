package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ImmunizationEvaluation
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
	Date                   *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_TargetDisease(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("TargetDisease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("TargetDisease", &resource.TargetDisease, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DoseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DoseStatus", &resource.DoseStatus, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseStatusReason(numDoseStatusReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDoseStatusReason >= len(resource.DoseStatusReason) {
		return CodeableConceptSelect("DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", &resource.DoseStatusReason[numDoseStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_Series(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Series", nil, htmlAttrs)
	}
	return StringInput("Series", resource.Series, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseNumberPositiveInt(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("DoseNumberPositiveInt", nil, htmlAttrs)
	}
	return IntInput("DoseNumberPositiveInt", resource.DoseNumberPositiveInt, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_DoseNumberString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DoseNumberString", nil, htmlAttrs)
	}
	return StringInput("DoseNumberString", resource.DoseNumberString, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_SeriesDosesPositiveInt(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SeriesDosesPositiveInt", nil, htmlAttrs)
	}
	return IntInput("SeriesDosesPositiveInt", resource.SeriesDosesPositiveInt, htmlAttrs)
}
func (resource *ImmunizationEvaluation) T_SeriesDosesString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SeriesDosesString", nil, htmlAttrs)
	}
	return StringInput("SeriesDosesString", resource.SeriesDosesString, htmlAttrs)
}
