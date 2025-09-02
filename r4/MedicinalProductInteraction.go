package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductInteraction
type MedicinalProductInteraction struct {
	Id                *string                                  `json:"id,omitempty"`
	Meta              *Meta                                    `json:"meta,omitempty"`
	ImplicitRules     *string                                  `json:"implicitRules,omitempty"`
	Language          *string                                  `json:"language,omitempty"`
	Text              *Narrative                               `json:"text,omitempty"`
	Contained         []Resource                               `json:"contained,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Subject           []Reference                              `json:"subject,omitempty"`
	Description       *string                                  `json:"description,omitempty"`
	Interactant       []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`
	Type              *CodeableConcept                         `json:"type,omitempty"`
	Effect            *CodeableConcept                         `json:"effect,omitempty"`
	Incidence         *CodeableConcept                         `json:"incidence,omitempty"`
	Management        *CodeableConcept                         `json:"management,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductInteraction
type MedicinalProductInteractionInteractant struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

type OtherMedicinalProductInteraction MedicinalProductInteraction

// on convert struct to json, automatically add resourceType=MedicinalProductInteraction
func (r MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductInteraction
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductInteraction: OtherMedicinalProductInteraction(r),
		ResourceType:                     "MedicinalProductInteraction",
	})
}

func (resource *MedicinalProductInteraction) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Effect(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("effect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("effect", resource.Effect, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Incidence(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("incidence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("incidence", resource.Incidence, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Management(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("management", nil, optionsValueSet)
	}
	return CodeableConceptSelect("management", resource.Management, optionsValueSet)
}
