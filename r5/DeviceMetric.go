package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *DeviceMetric) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceMetric.Id", nil)
	}
	return StringInput("DeviceMetric.Id", resource.Id)
}
func (resource *DeviceMetric) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceMetric.ImplicitRules", nil)
	}
	return StringInput("DeviceMetric.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceMetric) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceMetric.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceMetric) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceMetric.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceMetric.Type", &resource.Type, optionsValueSet)
}
func (resource *DeviceMetric) T_Unit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceMetric.Unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceMetric.Unit", resource.Unit, optionsValueSet)
}
func (resource *DeviceMetric) T_OperationalStatus() templ.Component {
	optionsValueSet := VSMetric_operational_status

	if resource == nil {
		return CodeSelect("DeviceMetric.OperationalStatus", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.OperationalStatus", resource.OperationalStatus, optionsValueSet)
}
func (resource *DeviceMetric) T_Color(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceMetric.Color", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.Color", resource.Color, optionsValueSet)
}
func (resource *DeviceMetric) T_Category() templ.Component {
	optionsValueSet := VSMetric_category

	if resource == nil {
		return CodeSelect("DeviceMetric.Category", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.Category", &resource.Category, optionsValueSet)
}
func (resource *DeviceMetric) T_CalibrationId(numCalibration int) templ.Component {

	if resource == nil || len(resource.Calibration) >= numCalibration {
		return StringInput("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Id", nil)
	}
	return StringInput("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Id", resource.Calibration[numCalibration].Id)
}
func (resource *DeviceMetric) T_CalibrationType(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_type

	if resource == nil || len(resource.Calibration) >= numCalibration {
		return CodeSelect("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Type", resource.Calibration[numCalibration].Type, optionsValueSet)
}
func (resource *DeviceMetric) T_CalibrationState(numCalibration int) templ.Component {
	optionsValueSet := VSMetric_calibration_state

	if resource == nil || len(resource.Calibration) >= numCalibration {
		return CodeSelect("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].State", nil, optionsValueSet)
	}
	return CodeSelect("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].State", resource.Calibration[numCalibration].State, optionsValueSet)
}
func (resource *DeviceMetric) T_CalibrationTime(numCalibration int) templ.Component {

	if resource == nil || len(resource.Calibration) >= numCalibration {
		return StringInput("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Time", nil)
	}
	return StringInput("DeviceMetric.Calibration["+strconv.Itoa(numCalibration)+"].Time", resource.Calibration[numCalibration].Time)
}
