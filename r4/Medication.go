package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Medication
type Medication struct {
	Id                *string                `json:"id,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []Resource             `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier           `json:"identifier,omitempty"`
	Code              *CodeableConcept       `json:"code,omitempty"`
	Status            *string                `json:"status,omitempty"`
	Manufacturer      *Reference             `json:"manufacturer,omitempty"`
	Form              *CodeableConcept       `json:"form,omitempty"`
	Amount            *Ratio                 `json:"amount,omitempty"`
	Ingredient        []MedicationIngredient `json:"ingredient,omitempty"`
	Batch             *MedicationBatch       `json:"batch,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Medication
type MedicationIngredient struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Reference       `json:"itemReference"`
	IsActive            *bool           `json:"isActive,omitempty"`
	Strength            *Ratio          `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Medication
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
func (resource *Medication) T_Form(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Medication.Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Medication.Form", resource.Form, optionsValueSet)
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
