package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/DeviceUsage
type DeviceUsage struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	BasedOn           []Reference           `json:"basedOn,omitempty"`
	Status            string                `json:"status"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Patient           Reference             `json:"patient"`
	DerivedFrom       []Reference           `json:"derivedFrom,omitempty"`
	Context           *Reference            `json:"context,omitempty"`
	TimingTiming      *Timing               `json:"timingTiming"`
	TimingPeriod      *Period               `json:"timingPeriod"`
	TimingDateTime    *string               `json:"timingDateTime"`
	DateAsserted      *string               `json:"dateAsserted,omitempty"`
	UsageStatus       *CodeableConcept      `json:"usageStatus,omitempty"`
	UsageReason       []CodeableConcept     `json:"usageReason,omitempty"`
	Adherence         *DeviceUsageAdherence `json:"adherence,omitempty"`
	InformationSource *Reference            `json:"informationSource,omitempty"`
	Device            CodeableReference     `json:"device"`
	Reason            []CodeableReference   `json:"reason,omitempty"`
	BodySite          *CodeableReference    `json:"bodySite,omitempty"`
	Note              []Annotation          `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceUsage
type DeviceUsageAdherence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Reason            []CodeableConcept `json:"reason"`
}

type OtherDeviceUsage DeviceUsage

// on convert struct to json, automatically add resourceType=DeviceUsage
func (r DeviceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUsage
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUsage: OtherDeviceUsage(r),
		ResourceType:     "DeviceUsage",
	})
}
func (resource *DeviceUsage) DeviceUsageLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *DeviceUsage) DeviceUsageStatus() templ.Component {
	optionsValueSet := VSDeviceusage_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
