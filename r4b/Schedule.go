package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Schedule
type Schedule struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	ServiceCategory   []CodeableConcept `json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `json:"specialty,omitempty"`
	Actor             []Reference       `json:"actor"`
	PlanningHorizon   *Period           `json:"planningHorizon,omitempty"`
	Comment           *string           `json:"comment,omitempty"`
}

type OtherSchedule Schedule

// on convert struct to json, automatically add resourceType=Schedule
func (r Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSchedule
		ResourceType string `json:"resourceType"`
	}{
		OtherSchedule: OtherSchedule(r),
		ResourceType:  "Schedule",
	})
}

func (resource *Schedule) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Schedule) T_ServiceCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("serviceCategory", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceCategory", &resource.ServiceCategory[0], optionsValueSet)
}
func (resource *Schedule) T_ServiceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceType", &resource.ServiceType[0], optionsValueSet)
}
func (resource *Schedule) T_Specialty(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
