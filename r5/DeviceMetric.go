package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                   *string                   `json:"id,omitempty"`
	Meta                 *Meta                     `json:"meta,omitempty"`
	ImplicitRules        *string                   `json:"implicitRules,omitempty"`
	Language             *string                   `json:"language,omitempty"`
	Text                 *Narrative                `json:"text,omitempty"`
	Contained            []Resource                `json:"contained,omitempty"`
	Extension            []Extension               `json:"extension,omitempty"`
	ModifierExtension    []Extension               `json:"modifierExtension,omitempty"`
	Identifier           []Identifier              `json:"identifier,omitempty"`
	Type                 CodeableConcept           `json:"type"`
	Unit                 *CodeableConcept          `json:"unit,omitempty"`
	Device               Reference                 `json:"device"`
	OperationalStatus    *string                   `json:"operationalStatus,omitempty"`
	Color                *string                   `json:"color,omitempty"`
	Category             string                    `json:"category"`
	MeasurementFrequency *Quantity                 `json:"measurementFrequency,omitempty"`
	Calibration          []DeviceMetricCalibration `json:"calibration,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceMetric
type DeviceMetricCalibration struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              *string     `json:"type,omitempty"`
	State             *string     `json:"state,omitempty"`
	Time              *string     `json:"time,omitempty"`
}

type OtherDeviceMetric DeviceMetric

// on convert struct to json, automatically add resourceType=DeviceMetric
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetric
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceMetric: OtherDeviceMetric(r),
		ResourceType:      "DeviceMetric",
	})
}

func (resource *DeviceMetric) DeviceMetricLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unit", resource.Unit, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricOperationalStatus() templ.Component {
	optionsValueSet := VSMetric_operational_status

	if resource == nil {
		return CodeSelect("operationalStatus", nil, optionsValueSet)
	}
	return CodeSelect("operationalStatus", resource.OperationalStatus, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricColor(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("color", nil, optionsValueSet)
	}
	return CodeSelect("color", resource.Color, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCategory() templ.Component {
	optionsValueSet := VSMetric_category

	if resource == nil {
		return CodeSelect("category", nil, optionsValueSet)
	}
	return CodeSelect("category", &resource.Category, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCalibrationType(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_type

	if resource == nil && len(resource.Calibration) >= numCalibration {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Calibration[numCalibration].Type, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCalibrationState(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_state

	if resource == nil && len(resource.Calibration) >= numCalibration {
		return CodeSelect("state", nil, optionsValueSet)
	}
	return CodeSelect("state", resource.Calibration[numCalibration].State, optionsValueSet)
}
