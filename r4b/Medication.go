package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Medication
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

// http://hl7.org/fhir/r4b/StructureDefinition/Medication
type MedicationIngredient struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Reference       `json:"itemReference"`
	IsActive            *bool           `json:"isActive,omitempty"`
	Strength            *Ratio          `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Medication
type MedicationBatch struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LotNumber         *string     `json:"lotNumber,omitempty"`
	ExpirationDate    *time.Time  `json:"expirationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Medication) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Medication/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Medication"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Medication) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Medication.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Medication.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedication_status

	if resource == nil {
		return CodeSelect("Medication.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Medication.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_Form(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Medication.Form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Medication.Form", resource.Form, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_IngredientItemCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("Medication.Ingredient."+strconv.Itoa(numIngredient)+"..ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Medication.Ingredient."+strconv.Itoa(numIngredient)+"..ItemCodeableConcept", &resource.Ingredient[numIngredient].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_IngredientIsActive(numIngredient int, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return BoolInput("Medication.Ingredient."+strconv.Itoa(numIngredient)+"..IsActive", nil, htmlAttrs)
	}
	return BoolInput("Medication.Ingredient."+strconv.Itoa(numIngredient)+"..IsActive", resource.Ingredient[numIngredient].IsActive, htmlAttrs)
}
func (resource *Medication) T_BatchLotNumber(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Medication.Batch.LotNumber", nil, htmlAttrs)
	}
	return StringInput("Medication.Batch.LotNumber", resource.Batch.LotNumber, htmlAttrs)
}
func (resource *Medication) T_BatchExpirationDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Medication.Batch.ExpirationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Medication.Batch.ExpirationDate", resource.Batch.ExpirationDate, htmlAttrs)
}
