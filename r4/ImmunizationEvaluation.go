package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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
	Date                   *string           `json:"date,omitempty"`
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

func (resource *ImmunizationEvaluation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationEvaluation.Id", nil)
	}
	return StringInput("ImmunizationEvaluation.Id", resource.Id)
}
func (resource *ImmunizationEvaluation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationEvaluation.ImplicitRules", nil)
	}
	return StringInput("ImmunizationEvaluation.ImplicitRules", resource.ImplicitRules)
}
func (resource *ImmunizationEvaluation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ImmunizationEvaluation.Language", nil, optionsValueSet)
	}
	return CodeSelect("ImmunizationEvaluation.Language", resource.Language, optionsValueSet)
}
func (resource *ImmunizationEvaluation) T_Status() templ.Component {
	optionsValueSet := VSImmunization_evaluation_status

	if resource == nil {
		return CodeSelect("ImmunizationEvaluation.Status", nil, optionsValueSet)
	}
	return CodeSelect("ImmunizationEvaluation.Status", &resource.Status, optionsValueSet)
}
func (resource *ImmunizationEvaluation) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationEvaluation.Date", nil)
	}
	return StringInput("ImmunizationEvaluation.Date", resource.Date)
}
func (resource *ImmunizationEvaluation) T_TargetDisease(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ImmunizationEvaluation.TargetDisease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.TargetDisease", &resource.TargetDisease, optionsValueSet)
}
func (resource *ImmunizationEvaluation) T_DoseStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ImmunizationEvaluation.DoseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.DoseStatus", &resource.DoseStatus, optionsValueSet)
}
func (resource *ImmunizationEvaluation) T_DoseStatusReason(numDoseStatusReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.DoseStatusReason) >= numDoseStatusReason {
		return CodeableConceptSelect("ImmunizationEvaluation.DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationEvaluation.DoseStatusReason["+strconv.Itoa(numDoseStatusReason)+"]", &resource.DoseStatusReason[numDoseStatusReason], optionsValueSet)
}
func (resource *ImmunizationEvaluation) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationEvaluation.Description", nil)
	}
	return StringInput("ImmunizationEvaluation.Description", resource.Description)
}
func (resource *ImmunizationEvaluation) T_Series() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationEvaluation.Series", nil)
	}
	return StringInput("ImmunizationEvaluation.Series", resource.Series)
}
