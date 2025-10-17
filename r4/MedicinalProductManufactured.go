package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductManufactured
type MedicinalProductManufactured struct {
	Id                      *string             `json:"id,omitempty"`
	Meta                    *Meta               `json:"meta,omitempty"`
	ImplicitRules           *string             `json:"implicitRules,omitempty"`
	Language                *string             `json:"language,omitempty"`
	Text                    *Narrative          `json:"text,omitempty"`
	Contained               []Resource          `json:"contained,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	ManufacturedDoseForm    CodeableConcept     `json:"manufacturedDoseForm"`
	UnitOfPresentation      *CodeableConcept    `json:"unitOfPresentation,omitempty"`
	Quantity                Quantity            `json:"quantity"`
	Manufacturer            []Reference         `json:"manufacturer,omitempty"`
	Ingredient              []Reference         `json:"ingredient,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept   `json:"otherCharacteristics,omitempty"`
}

type OtherMedicinalProductManufactured MedicinalProductManufactured

// struct -> json, automatically add resourceType=Patient
func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductManufactured
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductManufactured: OtherMedicinalProductManufactured(r),
		ResourceType:                      "MedicinalProductManufactured",
	})
}

// json -> struct, first reject if resourceType != MedicinalProductManufactured
func (r *MedicinalProductManufactured) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MedicinalProductManufactured" {
		return errors.New("resourceType not MedicinalProductManufactured")
	}
	return json.Unmarshal(data, (*OtherMedicinalProductManufactured)(r))
}

func (r MedicinalProductManufactured) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductManufactured/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MedicinalProductManufactured"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductManufactured) T_ManufacturedDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("manufacturedDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("manufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", &resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_Manufacturer(frs []FhirResource, numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", &resource.Manufacturer[numManufacturer], htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_Ingredient(frs []FhirResource, numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_PhysicalCharacteristics(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ProdCharacteristicInput("physicalCharacteristics", nil, htmlAttrs)
	}
	return ProdCharacteristicInput("physicalCharacteristics", resource.PhysicalCharacteristics, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_OtherCharacteristics(numOtherCharacteristics int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherCharacteristics >= len(resource.OtherCharacteristics) {
		return CodeableConceptSelect("otherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("otherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", &resource.OtherCharacteristics[numOtherCharacteristics], optionsValueSet, htmlAttrs)
}
