package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *MedicinalProductContraindication) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductContraindication.Id", nil)
	}
	return StringInput("MedicinalProductContraindication.Id", resource.Id)
}
func (resource *MedicinalProductContraindication) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductContraindication.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductContraindication.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductContraindication) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductContraindication.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductContraindication.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductContraindication) T_Disease(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductContraindication.Disease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.Disease", resource.Disease, optionsValueSet)
}
func (resource *MedicinalProductContraindication) T_DiseaseStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductContraindication.DiseaseStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.DiseaseStatus", resource.DiseaseStatus, optionsValueSet)
}
func (resource *MedicinalProductContraindication) T_Comorbidity(numComorbidity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Comorbidity) >= numComorbidity {
		return CodeableConceptSelect("MedicinalProductContraindication.Comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.Comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Comorbidity[numComorbidity], optionsValueSet)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyId(numOtherTherapy int) templ.Component {

	if resource == nil || len(resource.OtherTherapy) >= numOtherTherapy {
		return StringInput("MedicinalProductContraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", nil)
	}
	return StringInput("MedicinalProductContraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", resource.OtherTherapy[numOtherTherapy].Id)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OtherTherapy) >= numOtherTherapy {
		return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].TherapyRelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].TherapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet)
}
