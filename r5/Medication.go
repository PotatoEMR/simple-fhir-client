package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	StrengthRatio           *Ratio            `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *CodeableConcept  `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *Quantity         `json:"strengthQuantity,omitempty"`
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

func (resource *Medication) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Medication.Id", nil)
	}
	return StringInput("Medication.Id", resource.Id)
}
func (resource *Medication) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Medication.ImplicitRules", nil)
	}
	return StringInput("Medication.ImplicitRules", resource.ImplicitRules)
}
func (resource *Medication) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Medication.Language", nil, optionsValueSet)
	}
	return CodeSelect("Medication.Language", resource.Language, optionsValueSet)
}
func (resource *Medication) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Medication.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Medication.Code", resource.Code, optionsValueSet)
}
func (resource *Medication) T_Status() templ.Component {
	optionsValueSet := VSMedication_status

	if resource == nil {
		return CodeSelect("Medication.Status", nil, optionsValueSet)
	}
	return CodeSelect("Medication.Status", resource.Status, optionsValueSet)
}
func (resource *Medication) T_DoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Medication.DoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Medication.DoseForm", resource.DoseForm, optionsValueSet)
}
func (resource *Medication) T_IngredientId(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return StringInput("Medication.Ingredient["+strconv.Itoa(numIngredient)+"].Id", nil)
	}
	return StringInput("Medication.Ingredient["+strconv.Itoa(numIngredient)+"].Id", resource.Ingredient[numIngredient].Id)
}
func (resource *Medication) T_IngredientIsActive(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return BoolInput("Medication.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", nil)
	}
	return BoolInput("Medication.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", resource.Ingredient[numIngredient].IsActive)
}
func (resource *Medication) T_BatchId() templ.Component {

	if resource == nil {
		return StringInput("Medication.Batch.Id", nil)
	}
	return StringInput("Medication.Batch.Id", resource.Batch.Id)
}
func (resource *Medication) T_BatchLotNumber() templ.Component {

	if resource == nil {
		return StringInput("Medication.Batch.LotNumber", nil)
	}
	return StringInput("Medication.Batch.LotNumber", resource.Batch.LotNumber)
}
func (resource *Medication) T_BatchExpirationDate() templ.Component {

	if resource == nil {
		return StringInput("Medication.Batch.ExpirationDate", nil)
	}
	return StringInput("Medication.Batch.ExpirationDate", resource.Batch.ExpirationDate)
}
