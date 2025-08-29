package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Type              CodeableConcept           `json:"type"`
	Unit              *CodeableConcept          `json:"unit,omitempty"`
	Source            *Reference                `json:"source,omitempty"`
	Parent            *Reference                `json:"parent,omitempty"`
	OperationalStatus *string                   `json:"operationalStatus,omitempty"`
	Color             *string                   `json:"color,omitempty"`
	Category          string                    `json:"category"`
	MeasurementPeriod *Timing                   `json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibration `json:"calibration,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceMetric
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricOperationalStatus() templ.Component {
	optionsValueSet := VSMetric_operational_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.OperationalStatus
	}
	return CodeSelect("operationalStatus", currentVal, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricColor() templ.Component {
	optionsValueSet := VSMetric_color
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Color
	}
	return CodeSelect("color", currentVal, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCategory() templ.Component {
	optionsValueSet := VSMetric_category
	currentVal := ""
	if resource != nil {
		currentVal = resource.Category
	}
	return CodeSelect("category", currentVal, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCalibrationType(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_type
	currentVal := ""
	if resource != nil && len(resource.Calibration) >= numCalibration {
		currentVal = *resource.Calibration[numCalibration].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *DeviceMetric) DeviceMetricCalibrationState(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_state
	currentVal := ""
	if resource != nil && len(resource.Calibration) >= numCalibration {
		currentVal = *resource.Calibration[numCalibration].State
	}
	return CodeSelect("state", currentVal, optionsValueSet)
}
