package r4b

//generated with command go run ./bultaoreune
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
func (r Location) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Location/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Location"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Location) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSLocation_status

	if resource == nil {
		return CodeSelect("Location.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Location.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_OperationalStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("Location.OperationalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Location.OperationalStatus", resource.OperationalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Location.Name", nil, htmlAttrs)
	}
	return StringInput("Location.Name", resource.Name, htmlAttrs)
}
func (resource *Location) T_Alias(numAlias int, htmlAttrs string) templ.Component {

	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("Location.Alias."+strconv.Itoa(numAlias)+".", nil, htmlAttrs)
	}
	return StringInput("Location.Alias."+strconv.Itoa(numAlias)+".", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Location) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Location.Description", nil, htmlAttrs)
	}
	return StringInput("Location.Description", resource.Description, htmlAttrs)
}
func (resource *Location) T_Mode(htmlAttrs string) templ.Component {
	optionsValueSet := VSLocation_mode

	if resource == nil {
		return CodeSelect("Location.Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Location.Mode", resource.Mode, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Location.Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Location.Type."+strconv.Itoa(numType)+".", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Location) T_PhysicalType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Location.PhysicalType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Location.PhysicalType", resource.PhysicalType, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_AvailabilityExceptions(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Location.AvailabilityExceptions", nil, htmlAttrs)
	}
	return StringInput("Location.AvailabilityExceptions", resource.AvailabilityExceptions, htmlAttrs)
}
func (resource *Location) T_PositionLongitude(htmlAttrs string) templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Longitude", nil, htmlAttrs)
	}
	return Float64Input("Location.Position.Longitude", &resource.Position.Longitude, htmlAttrs)
}
func (resource *Location) T_PositionLatitude(htmlAttrs string) templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Latitude", nil, htmlAttrs)
	}
	return Float64Input("Location.Position.Latitude", &resource.Position.Latitude, htmlAttrs)
}
func (resource *Location) T_PositionAltitude(htmlAttrs string) templ.Component {

	if resource == nil {
		return Float64Input("Location.Position.Altitude", nil, htmlAttrs)
	}
	return Float64Input("Location.Position.Altitude", resource.Position.Altitude, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationDaysOfWeek(numHoursOfOperation int, numDaysOfWeek int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) || numDaysOfWeek >= len(resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek) {
		return CodeSelect("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..DaysOfWeek."+strconv.Itoa(numDaysOfWeek)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..DaysOfWeek."+strconv.Itoa(numDaysOfWeek)+".", &resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek[numDaysOfWeek], optionsValueSet, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationAllDay(numHoursOfOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return BoolInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..AllDay", nil, htmlAttrs)
	}
	return BoolInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..AllDay", resource.HoursOfOperation[numHoursOfOperation].AllDay, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationOpeningTime(numHoursOfOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return StringInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..OpeningTime", nil, htmlAttrs)
	}
	return StringInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..OpeningTime", resource.HoursOfOperation[numHoursOfOperation].OpeningTime, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationClosingTime(numHoursOfOperation int, htmlAttrs string) templ.Component {

	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return StringInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..ClosingTime", nil, htmlAttrs)
	}
	return StringInput("Location.HoursOfOperation."+strconv.Itoa(numHoursOfOperation)+"..ClosingTime", resource.HoursOfOperation[numHoursOfOperation].ClosingTime, htmlAttrs)
}
