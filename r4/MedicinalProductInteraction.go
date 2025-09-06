package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *MedicinalProductInteraction) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductInteraction.Id", nil)
	}
	return StringInput("MedicinalProductInteraction.Id", resource.Id)
}
func (resource *MedicinalProductInteraction) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductInteraction.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductInteraction.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductInteraction) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductInteraction.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductInteraction.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Description() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductInteraction.Description", nil)
	}
	return StringInput("MedicinalProductInteraction.Description", resource.Description)
}
func (resource *MedicinalProductInteraction) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Effect(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Effect", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Effect", resource.Effect, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Incidence(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Incidence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Incidence", resource.Incidence, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_Management(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Management", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Management", resource.Management, optionsValueSet)
}
func (resource *MedicinalProductInteraction) T_InteractantId(numInteractant int) templ.Component {

	if resource == nil || len(resource.Interactant) >= numInteractant {
		return StringInput("MedicinalProductInteraction.Interactant["+strconv.Itoa(numInteractant)+"].Id", nil)
	}
	return StringInput("MedicinalProductInteraction.Interactant["+strconv.Itoa(numInteractant)+"].Id", resource.Interactant[numInteractant].Id)
}
