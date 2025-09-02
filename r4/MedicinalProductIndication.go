package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductIndication) MedicinalProductIndicationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductIndication) MedicinalProductIndicationDiseaseSymptomProcedure(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("diseaseSymptomProcedure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("diseaseSymptomProcedure", resource.DiseaseSymptomProcedure, optionsValueSet)
}
func (resource *MedicinalProductIndication) MedicinalProductIndicationDiseaseStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("diseaseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("diseaseStatus", resource.DiseaseStatus, optionsValueSet)
}
func (resource *MedicinalProductIndication) MedicinalProductIndicationComorbidity(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("comorbidity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("comorbidity", &resource.Comorbidity[0], optionsValueSet)
}
func (resource *MedicinalProductIndication) MedicinalProductIndicationIntendedEffect(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("intendedEffect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("intendedEffect", resource.IntendedEffect, optionsValueSet)
}
func (resource *MedicinalProductIndication) MedicinalProductIndicationOtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.OtherTherapy) >= numOtherTherapy {
		return CodeableConceptSelect("therapyRelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("therapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet)
}
