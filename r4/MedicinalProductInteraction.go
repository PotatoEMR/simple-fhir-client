package r4

//generated with command go run ./bultaoreune
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
func (r MedicinalProductInteraction) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductInteraction/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "MedicinalProductInteraction"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductInteraction) T_Subject(numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Effect(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("effect", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("effect", resource.Effect, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Incidence(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("incidence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("incidence", resource.Incidence, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Management(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("management", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("management", resource.Management, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_InteractantItemReference(numInteractant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInteractant >= len(resource.Interactant) {
		return ReferenceInput("interactant["+strconv.Itoa(numInteractant)+"].itemReference", nil, htmlAttrs)
	}
	return ReferenceInput("interactant["+strconv.Itoa(numInteractant)+"].itemReference", &resource.Interactant[numInteractant].ItemReference, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_InteractantItemCodeableConcept(numInteractant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInteractant >= len(resource.Interactant) {
		return CodeableConceptSelect("interactant["+strconv.Itoa(numInteractant)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interactant["+strconv.Itoa(numInteractant)+"].itemCodeableConcept", &resource.Interactant[numInteractant].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
