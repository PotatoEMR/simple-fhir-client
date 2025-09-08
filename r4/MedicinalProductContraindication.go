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
func (resource *MedicinalProductContraindication) T_Disease(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductContraindication.Disease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.Disease", resource.Disease, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_DiseaseStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductContraindication.DiseaseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.DiseaseStatus", resource.DiseaseStatus, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_Comorbidity(numComorbidity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numComorbidity >= len(resource.Comorbidity) {
		return CodeableConceptSelect("MedicinalProductContraindication.Comorbidity."+strconv.Itoa(numComorbidity)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.Comorbidity."+strconv.Itoa(numComorbidity)+".", &resource.Comorbidity[numComorbidity], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy."+strconv.Itoa(numOtherTherapy)+"..TherapyRelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy."+strconv.Itoa(numOtherTherapy)+"..TherapyRelationshipType", &resource.OtherTherapy[numOtherTherapy].TherapyRelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductContraindication) T_OtherTherapyMedicationCodeableConcept(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numOtherTherapy >= len(resource.OtherTherapy) {
		return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy."+strconv.Itoa(numOtherTherapy)+"..MedicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductContraindication.OtherTherapy."+strconv.Itoa(numOtherTherapy)+"..MedicationCodeableConcept", &resource.OtherTherapy[numOtherTherapy].MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
