package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductContraindication) MedicinalProductContraindicationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductContraindication) MedicinalProductContraindicationDisease(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("disease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("disease", resource.Disease, optionsValueSet)
}
func (resource *MedicinalProductContraindication) MedicinalProductContraindicationDiseaseStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("diseaseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("diseaseStatus", resource.DiseaseStatus, optionsValueSet)
}
func (resource *MedicinalProductContraindication) MedicinalProductContraindicationComorbidity(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("comorbidity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("comorbidity", &resource.Comorbidity[0], optionsValueSet)
}
func (resource *MedicinalProductContraindication) MedicinalProductContraindicationOtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.OtherTherapy) >= numOtherTherapy {
		return CodeableConceptSelect("therapyRelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("therapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet)
}
