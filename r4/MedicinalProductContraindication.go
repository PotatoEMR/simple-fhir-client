package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductContraindication
type MedicinalProductContraindication struct {
	Id                    *string                                        `json:"id,omitempty"`
	Meta                  *Meta                                          `json:"meta,omitempty"`
	ImplicitRules         *string                                        `json:"implicitRules,omitempty"`
	Language              *string                                        `json:"language,omitempty"`
	Text                  *Narrative                                     `json:"text,omitempty"`
	Contained             []Resource                                     `json:"contained,omitempty"`
	Extension             []Extension                                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                                    `json:"modifierExtension,omitempty"`
	Subject               []Reference                                    `json:"subject,omitempty"`
	Disease               *CodeableConcept                               `json:"disease,omitempty"`
	DiseaseStatus         *CodeableConcept                               `json:"diseaseStatus,omitempty"`
	Comorbidity           []CodeableConcept                              `json:"comorbidity,omitempty"`
	TherapeuticIndication []Reference                                    `json:"therapeuticIndication,omitempty"`
	OtherTherapy          []MedicinalProductContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
	Population            []Population                                   `json:"population,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductContraindication
type MedicinalProductContraindicationOtherTherapy struct {
	Id                        *string         `json:"id,omitempty"`
	Extension                 []Extension     `json:"extension,omitempty"`
	ModifierExtension         []Extension     `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept `json:"therapyRelationshipType"`
	MedicationCodeableConcept CodeableConcept `json:"medicationCodeableConcept"`
	MedicationReference       Reference       `json:"medicationReference"`
}

type OtherMedicinalProductContraindication MedicinalProductContraindication

// on convert struct to json, automatically add resourceType=MedicinalProductContraindication
func (r MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductContraindication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductContraindication: OtherMedicinalProductContraindication(r),
		ResourceType:                          "MedicinalProductContraindication",
	})
}
func (r MedicinalProductContraindication) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductContraindication/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MedicinalProductContraindication"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductContraindication) T_Subject(numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_Disease(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("disease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("disease", resource.Disease, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_DiseaseStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diseaseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diseaseStatus", resource.DiseaseStatus, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_Comorbidity(numComorbidity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComorbidity >= len(resource.Comorbidity) {
		return CodeableConceptSelect("comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Comorbidity[numComorbidity], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_TherapeuticIndication(numTherapeuticIndication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTherapeuticIndication >= len(resource.TherapeuticIndication) {
		return ReferenceInput("therapeuticIndication["+strconv.Itoa(numTherapeuticIndication)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("therapeuticIndication["+strconv.Itoa(numTherapeuticIndication)+"]", &resource.TherapeuticIndication[numTherapeuticIndication], htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_Population(numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPopulation >= len(resource.Population) {
		return PopulationInput("population["+strconv.Itoa(numPopulation)+"]", nil, htmlAttrs)
	}
	return PopulationInput("population["+strconv.Itoa(numPopulation)+"]", &resource.Population[numPopulation], htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapyRelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyMedicationCodeableConcept(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationCodeableConcept", &resource.OtherTherapy[numOtherTherapy].MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyMedicationReference(numOtherTherapy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return ReferenceInput("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationReference", nil, htmlAttrs)
	}
	return ReferenceInput("otherTherapy["+strconv.Itoa(numOtherTherapy)+"].medicationReference", &resource.OtherTherapy[numOtherTherapy].MedicationReference, htmlAttrs)
}
