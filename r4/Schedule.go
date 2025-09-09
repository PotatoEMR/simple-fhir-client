package r4

//generated with command go run ./bultaoreune
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
func (r Schedule) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Schedule/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Schedule"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Schedule) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Schedule) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Schedule) T_ServiceType(numServiceType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableConceptSelect("serviceType["+strconv.Itoa(numServiceType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], optionsValueSet, htmlAttrs)
}
func (resource *Schedule) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Schedule) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
