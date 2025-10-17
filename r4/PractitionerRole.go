package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []Resource                      `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `json:"identifier,omitempty"`
	Active                 *bool                           `json:"active,omitempty"`
	Period                 *Period                         `json:"period,omitempty"`
	Practitioner           *Reference                      `json:"practitioner,omitempty"`
	Organization           *Reference                      `json:"organization,omitempty"`
	Code                   []CodeableConcept               `json:"code,omitempty"`
	Specialty              []CodeableConcept               `json:"specialty,omitempty"`
	Location               []Reference                     `json:"location,omitempty"`
	HealthcareService      []Reference                     `json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                  `json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                         `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                     `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
type PractitionerRoleAvailableTime struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	DaysOfWeek         []string    `json:"daysOfWeek,omitempty"`
	AllDay             *bool       `json:"allDay,omitempty"`
	AvailableStartTime *string     `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string     `json:"availableEndTime,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
type PractitionerRoleNotAvailable struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	During            *Period     `json:"during,omitempty"`
}

type OtherPractitionerRole PractitionerRole

// struct -> json, automatically add resourceType=Patient
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

func (r PractitionerRole) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PractitionerRole/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PractitionerRole"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r PractitionerRole) ResourceType() string {
	return "PractitionerRole"
}

func (resource *PractitionerRole) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *PractitionerRole) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *PractitionerRole) T_Practitioner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "practitioner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "practitioner", resource.Practitioner, htmlAttrs)
}
func (resource *PractitionerRole) T_Organization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "organization", resource.Organization, htmlAttrs)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Location(frs []FhirResource, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], htmlAttrs)
}
func (resource *PractitionerRole) T_HealthcareService(frs []FhirResource, numHealthcareService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHealthcareService >= len(resource.HealthcareService) {
		return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", &resource.HealthcareService[numHealthcareService], htmlAttrs)
}
func (resource *PractitionerRole) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *PractitionerRole) T_AvailabilityExceptions(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("availabilityExceptions", nil, htmlAttrs)
	}
	return StringInput("availabilityExceptions", resource.AvailabilityExceptions, htmlAttrs)
}
func (resource *PractitionerRole) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeDaysOfWeek(numAvailableTime int, numDaysOfWeek int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numAvailableTime >= len(resource.AvailableTime) || numDaysOfWeek >= len(resource.AvailableTime[numAvailableTime].DaysOfWeek) {
		return CodeSelect("availableTime["+strconv.Itoa(numAvailableTime)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("availableTime["+strconv.Itoa(numAvailableTime)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.AvailableTime[numAvailableTime].DaysOfWeek[numDaysOfWeek], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAllDay(numAvailableTime int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return BoolInput("availableTime["+strconv.Itoa(numAvailableTime)+"].allDay", nil, htmlAttrs)
	}
	return BoolInput("availableTime["+strconv.Itoa(numAvailableTime)+"].allDay", resource.AvailableTime[numAvailableTime].AllDay, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableStartTime(numAvailableTime int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableStartTime", nil, htmlAttrs)
	}
	return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableStartTime", resource.AvailableTime[numAvailableTime].AvailableStartTime, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableEndTime(numAvailableTime int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableEndTime", nil, htmlAttrs)
	}
	return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableEndTime", resource.AvailableTime[numAvailableTime].AvailableEndTime, htmlAttrs)
}
func (resource *PractitionerRole) T_NotAvailableDescription(numNotAvailable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNotAvailable >= len(resource.NotAvailable) {
		return StringInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].description", nil, htmlAttrs)
	}
	return StringInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].description", &resource.NotAvailable[numNotAvailable].Description, htmlAttrs)
}
func (resource *PractitionerRole) T_NotAvailableDuring(numNotAvailable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNotAvailable >= len(resource.NotAvailable) {
		return PeriodInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].during", nil, htmlAttrs)
	}
	return PeriodInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].during", resource.NotAvailable[numNotAvailable].During, htmlAttrs)
}
