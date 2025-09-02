package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	DoseNumberPositiveInt  *int              `json:"doseNumberPositiveInt"`
	DoseNumberString       *string           `json:"doseNumberString"`
	SeriesDosesPositiveInt *int              `json:"seriesDosesPositiveInt"`
	SeriesDosesString      *string           `json:"seriesDosesString"`
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

func (resource *ImmunizationEvaluation) ImmunizationEvaluationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ImmunizationEvaluation) ImmunizationEvaluationStatus() templ.Component {
	optionsValueSet := VSImmunization_evaluation_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ImmunizationEvaluation) ImmunizationEvaluationTargetDisease(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("targetDisease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("targetDisease", &resource.TargetDisease, optionsValueSet)
}
func (resource *ImmunizationEvaluation) ImmunizationEvaluationDoseStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("doseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("doseStatus", &resource.DoseStatus, optionsValueSet)
}
func (resource *ImmunizationEvaluation) ImmunizationEvaluationDoseStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("doseStatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("doseStatusReason", &resource.DoseStatusReason[0], optionsValueSet)
}
