package r4

//generated with command go run ./bultaoreune
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
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	LotNumber         *string       `json:"lotNumber,omitempty"`
	ExpirationDate    *FhirDateTime `json:"expirationDate,omitempty"`
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
func (resource *Medication) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_Manufacturer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "manufacturer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *Medication) T_Form(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("form", resource.Form, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_Amount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RatioInput("amount", nil, htmlAttrs)
	}
	return RatioInput("amount", resource.Amount, htmlAttrs)
}
func (resource *Medication) T_IngredientItemCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].itemCodeableConcept", &resource.Ingredient[numIngredient].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_IngredientItemReference(frs []FhirResource, numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"].itemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"].itemReference", &resource.Ingredient[numIngredient].ItemReference, htmlAttrs)
}
func (resource *Medication) T_IngredientIsActive(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", nil, htmlAttrs)
	}
	return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", resource.Ingredient[numIngredient].IsActive, htmlAttrs)
}
func (resource *Medication) T_IngredientStrength(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strength", nil, htmlAttrs)
	}
	return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strength", resource.Ingredient[numIngredient].Strength, htmlAttrs)
}
func (resource *Medication) T_BatchLotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("batch.lotNumber", nil, htmlAttrs)
	}
	return StringInput("batch.lotNumber", resource.Batch.LotNumber, htmlAttrs)
}
func (resource *Medication) T_BatchExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("batch.expirationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("batch.expirationDate", resource.Batch.ExpirationDate, htmlAttrs)
}
