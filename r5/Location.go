package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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

// struct -> json, automatically add resourceType=Patient
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}

// json -> struct, first reject if resourceType != Location
func (r *Location) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Location" {
		return errors.New("resourceType not Location")
	}
	return json.Unmarshal(data, (*OtherLocation)(r))
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
func (resource *Location) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Location) T_Address(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AddressInput("address", nil, htmlAttrs)
	}
	return AddressInput("address", resource.Address, htmlAttrs)
}
func (resource *Location) T_Form(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("form", resource.Form, optionsValueSet, htmlAttrs)
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
func (resource *Location) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *Location) T_HoursOfOperation(numHoursOfOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHoursOfOperation >= len(resource.HoursOfOperation) {
		return AvailabilityInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"]", nil, htmlAttrs)
	}
	return AvailabilityInput("hoursOfOperation["+strconv.Itoa(numHoursOfOperation)+"]", &resource.HoursOfOperation[numHoursOfOperation], htmlAttrs)
}
func (resource *Location) T_VirtualService(numVirtualService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVirtualService >= len(resource.VirtualService) {
		return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", nil, htmlAttrs)
	}
	return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", &resource.VirtualService[numVirtualService], htmlAttrs)
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
