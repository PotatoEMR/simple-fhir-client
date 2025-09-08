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
func (resource *MedicinalProductInteraction) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductInteraction.Description", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductInteraction.Description", resource.Description, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Effect(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Effect", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Effect", resource.Effect, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Incidence(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Incidence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Incidence", resource.Incidence, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_Management(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductInteraction.Management", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Management", resource.Management, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductInteraction) T_InteractantItemCodeableConcept(numInteractant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numInteractant >= len(resource.Interactant) {
		return CodeableConceptSelect("MedicinalProductInteraction.Interactant."+strconv.Itoa(numInteractant)+"..ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductInteraction.Interactant."+strconv.Itoa(numInteractant)+"..ItemCodeableConcept", &resource.Interactant[numInteractant].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
