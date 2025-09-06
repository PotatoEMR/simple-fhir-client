package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Location
type Location struct {
	Id                     *string                    `json:"id,omitempty"`
	Meta                   *Meta                      `json:"meta,omitempty"`
	ImplicitRules          *string                    `json:"implicitRules,omitempty"`
	Language               *string                    `json:"language,omitempty"`
	Text                   *Narrative                 `json:"text,omitempty"`
	Contained              []Resource                 `json:"contained,omitempty"`
	Extension              []Extension                `json:"extension,omitempty"`
	ModifierExtension      []Extension                `json:"modifierExtension,omitempty"`
	Identifier             []Identifier               `json:"identifier,omitempty"`
	Status                 *string                    `json:"status,omitempty"`
	OperationalStatus      *Coding                    `json:"operationalStatus,omitempty"`
	Name                   *string                    `json:"name,omitempty"`
	Alias                  []string                   `json:"alias,omitempty"`
	Description            *string                    `json:"description,omitempty"`
	Mode                   *string                    `json:"mode,omitempty"`
	Type                   []CodeableConcept          `json:"type,omitempty"`
	Telecom                []ContactPoint             `json:"telecom,omitempty"`
	Address                *Address                   `json:"address,omitempty"`
	PhysicalType           *CodeableConcept           `json:"physicalType,omitempty"`
	Position               *LocationPosition          `json:"position,omitempty"`
	ManagingOrganization   *Reference                 `json:"managingOrganization,omitempty"`
	PartOf                 *Reference                 `json:"partOf,omitempty"`
	HoursOfOperation       []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`
	AvailabilityExceptions *string                    `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Location
type LocationPosition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Longitude         float64     `json:"longitude"`
	Latitude          float64     `json:"latitude"`
	Altitude          *float64    `json:"altitude,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Location
type LocationHoursOfOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DaysOfWeek        []string    `json:"daysOfWeek,omitempty"`
	AllDay            *bool       `json:"allDay,omitempty"`
	OpeningTime       *string     `json:"openingTime,omitempty"`
	ClosingTime       *string     `json:"closingTime,omitempty"`
}

type OtherLocation Location

// on convert struct to json, automatically add resourceType=Location
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}

func (resource *Location) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Location.Id", nil)
	}
	return StringInput("Location.Id", resource.Id)
}
func (resource *Location) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Location.ImplicitRules", nil)
	}
	return StringInput("Location.ImplicitRules", resource.ImplicitRules)
}
func (resource *Location) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Location.Language", nil, optionsValueSet)
	}
	return CodeSelect("Location.Language", resource.Language, optionsValueSet)
}
func (resource *Location) T_Status() templ.Component {
	optionsValueSet := VSLocation_status

	if resource == nil {
		return CodeSelect("Location.Status", nil, optionsValueSet)
	}
	return CodeSelect("Location.Status", resource.Status, optionsValueSet)
}
func (resource *Location) T_OperationalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("Location.OperationalStatus", nil, optionsValueSet)
	}
	return CodingSelect("Location.OperationalStatus", resource.OperationalStatus, optionsValueSet)
}
func (resource *Location) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Location.Name", nil)
	}
	return StringInput("Location.Name", resource.Name)
}
func (resource *Location) T_Alias(numAlias int) templ.Component {

	if resource == nil || len(resource.Alias) >= numAlias {
		return StringInput("Location.Alias["+strconv.Itoa(numAlias)+"]", nil)
	}
	return StringInput("Location.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias])
}
func (resource *Location) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Location.Description", nil)
	}
	return StringInput("Location.Description", resource.Description)
}
func (resource *Location) T_Mode() templ.Component {
	optionsValueSet := VSLocation_mode

	if resource == nil {
		return CodeSelect("Location.Mode", nil, optionsValueSet)
	}
	return CodeSelect("Location.Mode", resource.Mode, optionsValueSet)
}
func (resource *Location) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Location.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Location.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Location) T_PhysicalType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Location.PhysicalType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Location.PhysicalType", resource.PhysicalType, optionsValueSet)
}
func (resource *Location) T_AvailabilityExceptions() templ.Component {

	if resource == nil {
		return StringInput("Location.AvailabilityExceptions", nil)
	}
	return StringInput("Location.AvailabilityExceptions", resource.AvailabilityExceptions)
}
func (resource *Location) T_PositionId() templ.Component {

	if resource == nil {
		return StringInput("Location.Position.Id", nil)
	}
	return StringInput("Location.Position.Id", resource.Position.Id)
}
func (resource *Location) T_PositionLongitude() templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Longitude", nil)
	}
	return Float64Input("Location.Position.Longitude", &resource.Position.Longitude)
}
func (resource *Location) T_PositionLatitude() templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Latitude", nil)
	}
	return Float64Input("Location.Position.Latitude", &resource.Position.Latitude)
}
func (resource *Location) T_PositionAltitude() templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Altitude", nil)
	}
	return Float64Input("Location.Position.Altitude", resource.Position.Altitude)
}
func (resource *Location) T_HoursOfOperationId(numHoursOfOperation int) templ.Component {

	if resource == nil || len(resource.HoursOfOperation) >= numHoursOfOperation {
		return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].Id", nil)
	}
	return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].Id", resource.HoursOfOperation[numHoursOfOperation].Id)
}
func (resource *Location) T_HoursOfOperationDaysOfWeek(numHoursOfOperation int, numDaysOfWeek int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || len(resource.HoursOfOperation) >= numHoursOfOperation || len(resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek) >= numDaysOfWeek {
		return CodeSelect("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet)
	}
	return CodeSelect("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek[numDaysOfWeek], optionsValueSet)
}
func (resource *Location) T_HoursOfOperationAllDay(numHoursOfOperation int) templ.Component {

	if resource == nil || len(resource.HoursOfOperation) >= numHoursOfOperation {
		return BoolInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].AllDay", nil)
	}
	return BoolInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].AllDay", resource.HoursOfOperation[numHoursOfOperation].AllDay)
}
func (resource *Location) T_HoursOfOperationOpeningTime(numHoursOfOperation int) templ.Component {

	if resource == nil || len(resource.HoursOfOperation) >= numHoursOfOperation {
		return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].OpeningTime", nil)
	}
	return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].OpeningTime", resource.HoursOfOperation[numHoursOfOperation].OpeningTime)
}
func (resource *Location) T_HoursOfOperationClosingTime(numHoursOfOperation int) templ.Component {

	if resource == nil || len(resource.HoursOfOperation) >= numHoursOfOperation {
		return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].ClosingTime", nil)
	}
	return StringInput("Location.HoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].ClosingTime", resource.HoursOfOperation[numHoursOfOperation].ClosingTime)
}
