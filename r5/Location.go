package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Location
type Location struct {
	Id                   *string                 `json:"id,omitempty"`
	Meta                 *Meta                   `json:"meta,omitempty"`
	ImplicitRules        *string                 `json:"implicitRules,omitempty"`
	Language             *string                 `json:"language,omitempty"`
	Text                 *Narrative              `json:"text,omitempty"`
	Contained            []Resource              `json:"contained,omitempty"`
	Extension            []Extension             `json:"extension,omitempty"`
	ModifierExtension    []Extension             `json:"modifierExtension,omitempty"`
	Identifier           []Identifier            `json:"identifier,omitempty"`
	Status               *string                 `json:"status,omitempty"`
	OperationalStatus    *Coding                 `json:"operationalStatus,omitempty"`
	Name                 *string                 `json:"name,omitempty"`
	Alias                []string                `json:"alias,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	Mode                 *string                 `json:"mode,omitempty"`
	Type                 []CodeableConcept       `json:"type,omitempty"`
	Contact              []ExtendedContactDetail `json:"contact,omitempty"`
	Address              *Address                `json:"address,omitempty"`
	Form                 *CodeableConcept        `json:"form,omitempty"`
	Position             *LocationPosition       `json:"position,omitempty"`
	ManagingOrganization *Reference              `json:"managingOrganization,omitempty"`
	PartOf               *Reference              `json:"partOf,omitempty"`
	Characteristic       []CodeableConcept       `json:"characteristic,omitempty"`
	HoursOfOperation     []Availability          `json:"hoursOfOperation,omitempty"`
	VirtualService       []VirtualServiceDetail  `json:"virtualService,omitempty"`
	Endpoint             []Reference             `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Location
type LocationPosition struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Longitude         float64     `json:"longitude"`
	Latitude          float64     `json:"latitude"`
	Altitude          *float64    `json:"altitude,omitempty"`
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
func (resource *Location) T_Form(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Location.Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Location.Form", resource.Form, optionsValueSet)
}
func (resource *Location) T_Characteristic(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("Location.Characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Location.Characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet)
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
