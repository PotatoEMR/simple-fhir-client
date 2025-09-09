package r4

//generated with command go run ./bultaoreune
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
func (r MedicinalProductIndication) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductIndication/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MedicinalProductIndication"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductIndication) T_DiseaseSymptomProcedure(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diseaseSymptomProcedure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diseaseSymptomProcedure", resource.DiseaseSymptomProcedure, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIndication) T_DiseaseStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diseaseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diseaseStatus", resource.DiseaseStatus, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIndication) T_Comorbidity(numComorbidity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComorbidity >= len(resource.Comorbidity) {
		return CodeableConceptSelect("comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Comorbidity[numComorbidity], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIndication) T_IntendedEffect(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("intendedEffect", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("intendedEffect", resource.IntendedEffect, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIndication) T_OtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapyRelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIndication) T_OtherTherapyMedicationCodeableConcept(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationCodeableConcept", &resource.OtherTherapy[numOtherTherapy].MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
