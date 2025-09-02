package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductManufactured) MedicinalProductManufacturedLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductManufactured) MedicinalProductManufacturedManufacturedDoseForm(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("manufacturedDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("manufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet)
}
func (resource *MedicinalProductManufactured) MedicinalProductManufacturedUnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *MedicinalProductManufactured) MedicinalProductManufacturedOtherCharacteristics(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("otherCharacteristics", nil, optionsValueSet)
	}
	return CodeableConceptSelect("otherCharacteristics", &resource.OtherCharacteristics[0], optionsValueSet)
}
