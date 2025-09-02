package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Medication
type Medication struct {
	Id                           *string                `json:"id,omitempty"`
	Meta                         *Meta                  `json:"meta,omitempty"`
	ImplicitRules                *string                `json:"implicitRules,omitempty"`
	Language                     *string                `json:"language,omitempty"`
	Text                         *Narrative             `json:"text,omitempty"`
	Contained                    []Resource             `json:"contained,omitempty"`
	Extension                    []Extension            `json:"extension,omitempty"`
	ModifierExtension            []Extension            `json:"modifierExtension,omitempty"`
	Identifier                   []Identifier           `json:"identifier,omitempty"`
	Code                         *CodeableConcept       `json:"code,omitempty"`
	Status                       *string                `json:"status,omitempty"`
	MarketingAuthorizationHolder *Reference             `json:"marketingAuthorizationHolder,omitempty"`
	DoseForm                     *CodeableConcept       `json:"doseForm,omitempty"`
	TotalVolume                  *Quantity              `json:"totalVolume,omitempty"`
	Ingredient                   []MedicationIngredient `json:"ingredient,omitempty"`
	Batch                        *MedicationBatch       `json:"batch,omitempty"`
	Definition                   *Reference             `json:"definition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Medication
type MedicationIngredient struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Item                    CodeableReference `json:"item"`
	IsActive                *bool             `json:"isActive,omitempty"`
	StrengthRatio           *Ratio            `json:"strengthRatio"`
	StrengthCodeableConcept *CodeableConcept  `json:"strengthCodeableConcept"`
	StrengthQuantity        *Quantity         `json:"strengthQuantity"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Medication
type MedicationBatch struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LotNumber         *string     `json:"lotNumber,omitempty"`
	ExpirationDate    *string     `json:"expirationDate,omitempty"`
}

type OtherMedication Medication

// on convert struct to json, automatically add resourceType=Medication
func (r Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

func (resource *Medication) MedicationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Medication) MedicationCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *Medication) MedicationStatus() templ.Component {
	optionsValueSet := VSMedication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *Medication) MedicationDoseForm(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("doseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("doseForm", resource.DoseForm, optionsValueSet)
}
