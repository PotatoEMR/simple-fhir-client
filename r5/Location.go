package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Location) LocationStatus() templ.Component {
	optionsValueSet := VSLocation_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Location) LocationMode() templ.Component {
	optionsValueSet := VSLocation_mode
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
