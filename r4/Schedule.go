package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Schedule
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

func (resource *Schedule) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Schedule.Id", nil)
	}
	return StringInput("Schedule.Id", resource.Id)
}
func (resource *Schedule) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Schedule.ImplicitRules", nil)
	}
	return StringInput("Schedule.ImplicitRules", resource.ImplicitRules)
}
func (resource *Schedule) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Schedule.Language", nil, optionsValueSet)
	}
	return CodeSelect("Schedule.Language", resource.Language, optionsValueSet)
}
func (resource *Schedule) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Schedule.Active", nil)
	}
	return BoolInput("Schedule.Active", resource.Active)
}
func (resource *Schedule) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceCategory) >= numServiceCategory {
		return CodeableConceptSelect("Schedule.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Schedule.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet)
}
func (resource *Schedule) T_ServiceType(numServiceType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceType) >= numServiceType {
		return CodeableConceptSelect("Schedule.ServiceType["+strconv.Itoa(numServiceType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Schedule.ServiceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], optionsValueSet)
}
func (resource *Schedule) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("Schedule.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Schedule.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *Schedule) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("Schedule.Comment", nil)
	}
	return StringInput("Schedule.Comment", resource.Comment)
}
