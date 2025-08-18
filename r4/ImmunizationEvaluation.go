//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
