package r5

//generated with command go run ./bultaoreune
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
func (resource *Medication) T_MarketingAuthorizationHolder(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("marketingAuthorizationHolder", nil, htmlAttrs)
	}
	return ReferenceInput("marketingAuthorizationHolder", resource.MarketingAuthorizationHolder, htmlAttrs)
}
func (resource *Medication) T_DoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("doseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("doseForm", resource.DoseForm, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_TotalVolume(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("totalVolume", nil, htmlAttrs)
	}
	return QuantityInput("totalVolume", resource.TotalVolume, htmlAttrs)
}
func (resource *Medication) T_Definition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("definition", nil, htmlAttrs)
	}
	return ReferenceInput("definition", resource.Definition, htmlAttrs)
}
func (resource *Medication) T_IngredientItem(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableReferenceInput("ingredient["+strconv.Itoa(numIngredient)+"].item", nil, htmlAttrs)
	}
	return CodeableReferenceInput("ingredient["+strconv.Itoa(numIngredient)+"].item", &resource.Ingredient[numIngredient].Item, htmlAttrs)
}
func (resource *Medication) T_IngredientIsActive(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", nil, htmlAttrs)
	}
	return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", resource.Ingredient[numIngredient].IsActive, htmlAttrs)
}
func (resource *Medication) T_IngredientStrengthRatio(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strengthRatio", nil, htmlAttrs)
	}
	return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strengthRatio", resource.Ingredient[numIngredient].StrengthRatio, htmlAttrs)
}
func (resource *Medication) T_IngredientStrengthCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].strengthCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].strengthCodeableConcept", resource.Ingredient[numIngredient].StrengthCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Medication) T_IngredientStrengthQuantity(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return QuantityInput("ingredient["+strconv.Itoa(numIngredient)+"].strengthQuantity", nil, htmlAttrs)
	}
	return QuantityInput("ingredient["+strconv.Itoa(numIngredient)+"].strengthQuantity", resource.Ingredient[numIngredient].StrengthQuantity, htmlAttrs)
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
