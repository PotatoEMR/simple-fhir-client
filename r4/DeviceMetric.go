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
func (resource *DeviceMetric) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Unit(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unit", resource.Unit, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Source(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source", resource.Source, htmlAttrs)
}
func (resource *DeviceMetric) T_Parent(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "parent", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "parent", resource.Parent, htmlAttrs)
}
func (resource *DeviceMetric) T_OperationalStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMetric_operational_status

	if resource == nil {
		return CodeSelect("operationalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("operationalStatus", resource.OperationalStatus, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Color(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMetric_color

	if resource == nil {
		return CodeSelect("color", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("color", resource.Color, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_Category(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMetric_category

	if resource == nil {
		return CodeSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("category", &resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_MeasurementPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("measurementPeriod", nil, htmlAttrs)
	}
	return TimingInput("measurementPeriod", resource.MeasurementPeriod, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationType(numCalibration int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMetric_calibration_type

	if resource == nil || numCalibration >= len(resource.Calibration) {
		return CodeSelect("calibration["+strconv.Itoa(numCalibration)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("calibration["+strconv.Itoa(numCalibration)+"].type", resource.Calibration[numCalibration].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationState(numCalibration int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMetric_calibration_state

	if resource == nil || numCalibration >= len(resource.Calibration) {
		return CodeSelect("calibration["+strconv.Itoa(numCalibration)+"].state", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("calibration["+strconv.Itoa(numCalibration)+"].state", resource.Calibration[numCalibration].State, optionsValueSet, htmlAttrs)
}
func (resource *DeviceMetric) T_CalibrationTime(numCalibration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCalibration >= len(resource.Calibration) {
		return StringInput("calibration["+strconv.Itoa(numCalibration)+"].time", nil, htmlAttrs)
	}
	return StringInput("calibration["+strconv.Itoa(numCalibration)+"].time", resource.Calibration[numCalibration].Time, htmlAttrs)
}
