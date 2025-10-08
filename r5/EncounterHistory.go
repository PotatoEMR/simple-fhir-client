package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistory struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Encounter         *Reference                 `json:"encounter,omitempty"`
	Identifier        []Identifier               `json:"identifier,omitempty"`
	Status            string                     `json:"status"`
	Class             CodeableConcept            `json:"class"`
	Type              []CodeableConcept          `json:"type,omitempty"`
	ServiceType       []CodeableReference        `json:"serviceType,omitempty"`
	Subject           *Reference                 `json:"subject,omitempty"`
	SubjectStatus     *CodeableConcept           `json:"subjectStatus,omitempty"`
	ActualPeriod      *Period                    `json:"actualPeriod,omitempty"`
	PlannedStartDate  *FhirDateTime              `json:"plannedStartDate,omitempty"`
	PlannedEndDate    *FhirDateTime              `json:"plannedEndDate,omitempty"`
	Length            *Duration                  `json:"length,omitempty"`
	Location          []EncounterHistoryLocation `json:"location,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistoryLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Location          Reference        `json:"location"`
	Form              *CodeableConcept `json:"form,omitempty"`
}

type OtherEncounterHistory EncounterHistory

// on convert struct to json, automatically add resourceType=EncounterHistory
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounterHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounterHistory: OtherEncounterHistory(r),
		ResourceType:          "EncounterHistory",
	})
}
func (r EncounterHistory) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EncounterHistory/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EncounterHistory"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EncounterHistory) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *EncounterHistory) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_Class(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class", &resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_ServiceType(numServiceType int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], htmlAttrs)
}
func (resource *EncounterHistory) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *EncounterHistory) T_SubjectStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectStatus", resource.SubjectStatus, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_ActualPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("actualPeriod", nil, htmlAttrs)
	}
	return PeriodInput("actualPeriod", resource.ActualPeriod, htmlAttrs)
}
func (resource *EncounterHistory) T_PlannedStartDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("plannedStartDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("plannedStartDate", resource.PlannedStartDate, htmlAttrs)
}
func (resource *EncounterHistory) T_PlannedEndDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("plannedEndDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("plannedEndDate", resource.PlannedEndDate, htmlAttrs)
}
func (resource *EncounterHistory) T_Length(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("length", nil, htmlAttrs)
	}
	return DurationInput("length", resource.Length, htmlAttrs)
}
func (resource *EncounterHistory) T_LocationLocation(frs []FhirResource, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"].location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"].location", &resource.Location[numLocation].Location, htmlAttrs)
}
func (resource *EncounterHistory) T_LocationForm(numLocation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].form", resource.Location[numLocation].Form, optionsValueSet, htmlAttrs)
}
