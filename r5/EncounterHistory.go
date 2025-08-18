//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistory struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Encounter         *Reference                 `json:"encounter,omitempty"`
	Identifier        []Identifier               `json:"identifier,omitempty"`
	Status            string                     `json:"status"`
	Class             CodeableConcept            `json:"class"`
	Type              []CodeableConcept          `json:"type,omitempty"`
	ServiceType       []CodeableReference        `json:"serviceType,omitempty"`
	Subject           *Reference                 `json:"subject,omitempty"`
	SubjectStatus     *CodeableConcept           `json:"subjectStatus,omitempty"`
	ActualPeriod      *Period                    `json:"actualPeriod,omitempty"`
	PlannedStartDate  *string                    `json:"plannedStartDate,omitempty"`
	PlannedEndDate    *string                    `json:"plannedEndDate,omitempty"`
	Length            *Duration                  `json:"length,omitempty"`
	Location          []EncounterHistoryLocation `json:"location,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistoryLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Location          Reference        `json:"location"`
	Form              *CodeableConcept `json:"form,omitempty"`
}

type OtherEncounterHistory EncounterHistory

// on convert struct to json, automatically add resourceType=EncounterHistory
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounterHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounterHistory: OtherEncounterHistory(r),
		ResourceType:          "EncounterHistory",
	})
}
