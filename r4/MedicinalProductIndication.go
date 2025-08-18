//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIndication
type MedicinalProductIndication struct {
	Id                      *string                                  `json:"id,omitempty"`
	Meta                    *Meta                                    `json:"meta,omitempty"`
	ImplicitRules           *string                                  `json:"implicitRules,omitempty"`
	Language                *string                                  `json:"language,omitempty"`
	Text                    *Narrative                               `json:"text,omitempty"`
	Contained               []Resource                               `json:"contained,omitempty"`
	Extension               []Extension                              `json:"extension,omitempty"`
	ModifierExtension       []Extension                              `json:"modifierExtension,omitempty"`
	Subject                 []Reference                              `json:"subject,omitempty"`
	DiseaseSymptomProcedure *CodeableConcept                         `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableConcept                         `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableConcept                        `json:"comorbidity,omitempty"`
	IntendedEffect          *CodeableConcept                         `json:"intendedEffect,omitempty"`
	Duration                *Quantity                                `json:"duration,omitempty"`
	OtherTherapy            []MedicinalProductIndicationOtherTherapy `json:"otherTherapy,omitempty"`
	UndesirableEffect       []Reference                              `json:"undesirableEffect,omitempty"`
	Population              []Population                             `json:"population,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIndication
type MedicinalProductIndicationOtherTherapy struct {
	Id                        *string         `json:"id,omitempty"`
	Extension                 []Extension     `json:"extension,omitempty"`
	ModifierExtension         []Extension     `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept `json:"therapyRelationshipType"`
	MedicationCodeableConcept CodeableConcept `json:"medicationCodeableConcept"`
	MedicationReference       Reference       `json:"medicationReference"`
}

type OtherMedicinalProductIndication MedicinalProductIndication

// on convert struct to json, automatically add resourceType=MedicinalProductIndication
func (r MedicinalProductIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductIndication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductIndication: OtherMedicinalProductIndication(r),
		ResourceType:                    "MedicinalProductIndication",
	})
}
