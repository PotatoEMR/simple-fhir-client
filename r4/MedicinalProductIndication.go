package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *MedicinalProductIndication) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductIndication.Id", nil)
	}
	return StringInput("MedicinalProductIndication.Id", resource.Id)
}
func (resource *MedicinalProductIndication) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductIndication.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductIndication.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductIndication) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductIndication.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductIndication.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductIndication) T_DiseaseSymptomProcedure(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductIndication.DiseaseSymptomProcedure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIndication.DiseaseSymptomProcedure", resource.DiseaseSymptomProcedure, optionsValueSet)
}
func (resource *MedicinalProductIndication) T_DiseaseStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductIndication.DiseaseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIndication.DiseaseStatus", resource.DiseaseStatus, optionsValueSet)
}
func (resource *MedicinalProductIndication) T_Comorbidity(numComorbidity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Comorbidity) >= numComorbidity {
		return CodeableConceptSelect("MedicinalProductIndication.Comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIndication.Comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Comorbidity[numComorbidity], optionsValueSet)
}
func (resource *MedicinalProductIndication) T_IntendedEffect(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductIndication.IntendedEffect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIndication.IntendedEffect", resource.IntendedEffect, optionsValueSet)
}
func (resource *MedicinalProductIndication) T_OtherTherapyId(numOtherTherapy int) templ.Component {

	if resource == nil || len(resource.OtherTherapy) >= numOtherTherapy {
		return StringInput("MedicinalProductIndication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", nil)
	}
	return StringInput("MedicinalProductIndication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", resource.OtherTherapy[numOtherTherapy].Id)
}
func (resource *MedicinalProductIndication) T_OtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OtherTherapy) >= numOtherTherapy {
		return CodeableConceptSelect("MedicinalProductIndication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].TherapyRelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIndication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].TherapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet)
}
