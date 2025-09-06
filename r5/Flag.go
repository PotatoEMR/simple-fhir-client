package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Flag
type Flag struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Status            string            `json:"status"`
	Category          []CodeableConcept `json:"category,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Subject           Reference         `json:"subject"`
	Period            *Period           `json:"period,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
}

type OtherFlag Flag

// on convert struct to json, automatically add resourceType=Flag
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

func (resource *Flag) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Flag.Id", nil)
	}
	return StringInput("Flag.Id", resource.Id)
}
func (resource *Flag) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Flag.ImplicitRules", nil)
	}
	return StringInput("Flag.ImplicitRules", resource.ImplicitRules)
}
func (resource *Flag) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Flag.Language", nil, optionsValueSet)
	}
	return CodeSelect("Flag.Language", resource.Language, optionsValueSet)
}
func (resource *Flag) T_Status() templ.Component {
	optionsValueSet := VSFlag_status

	if resource == nil {
		return CodeSelect("Flag.Status", nil, optionsValueSet)
	}
	return CodeSelect("Flag.Status", &resource.Status, optionsValueSet)
}
func (resource *Flag) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Flag.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Flag.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Flag) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Flag.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Flag.Code", &resource.Code, optionsValueSet)
}
