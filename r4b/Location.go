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
func (resource *Location) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSLocation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_OperationalStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("operationalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("operationalStatus", resource.OperationalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Location) T_Alias(numAlias int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Location) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Location) T_Mode(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSLocation_mode

	if resource == nil {
		return CodeSelect("mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("mode", resource.Mode, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Location) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *Location) T_Address(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AddressInput("address", nil, htmlAttrs)
	}
	return AddressInput("address", resource.Address, htmlAttrs)
}
func (resource *Location) T_PhysicalType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("physicalType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("physicalType", resource.PhysicalType, optionsValueSet, htmlAttrs)
}
func (resource *Location) T_ManagingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization", resource.ManagingOrganization, htmlAttrs)
}
func (resource *Location) T_PartOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "partOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf", resource.PartOf, htmlAttrs)
}
func (resource *Location) T_AvailabilityExceptions(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("availabilityExceptions", nil, htmlAttrs)
	}
	return StringInput("availabilityExceptions", resource.AvailabilityExceptions, htmlAttrs)
}
func (resource *Location) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
func (resource *Location) T_PositionLongitude(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("position.longitude", nil, htmlAttrs)
	}
	return Float64Input("position.longitude", &resource.Position.Longitude, htmlAttrs)
}
func (resource *Location) T_PositionLatitude(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("position.latitude", nil, htmlAttrs)
	}
	return Float64Input("position.latitude", &resource.Position.Latitude, htmlAttrs)
}
func (resource *Location) T_PositionAltitude(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("position.altitude", nil, htmlAttrs)
	}
	return Float64Input("position.altitude", resource.Position.Altitude, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationDaysOfWeek(numHoursOfOperation int, numDaysOfWeek int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) || numDaysOfWeek >= len(resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek) {
		return CodeSelect("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.HoursOfOperation[numHoursOfOperation].DaysOfWeek[numDaysOfWeek], optionsValueSet, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationAllDay(numHoursOfOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return BoolInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].allDay", nil, htmlAttrs)
	}
	return BoolInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].allDay", resource.HoursOfOperation[numHoursOfOperation].AllDay, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationOpeningTime(numHoursOfOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return StringInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].openingTime", nil, htmlAttrs)
	}
	return StringInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].openingTime", resource.HoursOfOperation[numHoursOfOperation].OpeningTime, htmlAttrs)
}
func (resource *Location) T_HoursOfOperationClosingTime(numHoursOfOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return StringInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].closingTime", nil, htmlAttrs)
	}
	return StringInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"].closingTime", resource.HoursOfOperation[numHoursOfOperation].ClosingTime, htmlAttrs)
}
