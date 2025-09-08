package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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

// on convert struct to json, automatically add resourceType=MedicinalProductManufactured
func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductManufactured
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductManufactured: OtherMedicinalProductManufactured(r),
		ResourceType:                      "MedicinalProductManufactured",
	})
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
func (resource *MedicinalProductManufactured) T_ManufacturedDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductManufactured.ManufacturedDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductManufactured.ManufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductManufactured.UnitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductManufactured.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductManufactured) T_OtherCharacteristics(numOtherCharacteristics int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numOtherCharacteristics >= len(resource.OtherCharacteristics) {
		return CodeableConceptSelect("MedicinalProductManufactured.OtherCharacteristics."+strconv.Itoa(numOtherCharacteristics)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductManufactured.OtherCharacteristics."+strconv.Itoa(numOtherCharacteristics)+".", &resource.OtherCharacteristics[numOtherCharacteristics], optionsValueSet, htmlAttrs)
}
