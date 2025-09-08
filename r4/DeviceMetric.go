package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r DeviceMetric) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceMetric/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceMetric"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DeviceMetric) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceMetric.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceMetric.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Unit(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceMetric.Unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceMetric.Unit", resource.Unit, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_OperationalStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSMetric_operational_status

	if resource == nil {
		return CodeSelect("DeviceMetric.OperationalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceMetric.OperationalStatus", resource.OperationalStatus, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Color(htmlAttrs string) templ.Component {
	optionsValueSet := VSMetric_color

	if resource == nil {
		return CodeSelect("DeviceMetric.Color", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceMetric.Color", resource.Color, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Category(htmlAttrs string) templ.Component {
	optionsValueSet := VSMetric_category

	if resource == nil {
		return CodeSelect("DeviceMetric.Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceMetric.Category", &resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationType(numCalibration int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMetric_calibration_type

	if resource == nil || numCalibration >= len(resource.Calibration) {
		return CodeSelect("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..Type", resource.Calibration[numCalibration].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationState(numCalibration int, htmlAttrs string) templ.Component {
	optionsValueSet := VSMetric_calibration_state

	if resource == nil || numCalibration >= len(resource.Calibration) {
		return CodeSelect("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..State", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..State", resource.Calibration[numCalibration].State, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationTime(numCalibration int, htmlAttrs string) templ.Component {

	if resource == nil || numCalibration >= len(resource.Calibration) {
		return StringInput("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..Time", nil, htmlAttrs)
	}
	return StringInput("DeviceMetric.Calibration."+strconv.Itoa(numCalibration)+"..Time", resource.Calibration[numCalibration].Time, htmlAttrs)
}
