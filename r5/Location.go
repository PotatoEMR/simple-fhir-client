package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Location) LocationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Location) LocationStatus() templ.Component {
	optionsValueSet := VSLocation_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *Location) LocationOperationalStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("operationalStatus", nil, optionsValueSet)
	}
	return CodingSelect("operationalStatus", resource.OperationalStatus, optionsValueSet)
}
func (resource *Location) LocationMode() templ.Component {
	optionsValueSet := VSLocation_mode

	if resource != nil {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", resource.Mode, optionsValueSet)
}
func (resource *Location) LocationType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *Location) LocationForm(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("form", resource.Form, optionsValueSet)
}
func (resource *Location) LocationCharacteristic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("characteristic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("characteristic", &resource.Characteristic[0], optionsValueSet)
}
