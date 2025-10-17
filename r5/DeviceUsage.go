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
	TimingTiming      *Timing               `json:"timingTiming,omitempty"`
	TimingPeriod      *Period               `json:"timingPeriod,omitempty"`
	TimingDateTime    *FhirDateTime         `json:"timingDateTime,omitempty"`
	DateAsserted      *FhirDateTime         `json:"dateAsserted,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r DeviceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUsage
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUsage: OtherDeviceUsage(r),
		ResourceType:     "DeviceUsage",
	})
}

// json -> struct, first reject if resourceType != DeviceUsage
func (r *DeviceUsage) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "DeviceUsage" {
		return errors.New("resourceType not DeviceUsage")
	}
	return json.Unmarshal(data, (*OtherDeviceUsage)(r))
}

func (r DeviceUsage) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceUsage/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceUsage"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DeviceUsage) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DeviceUsage) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *DeviceUsage) T_DerivedFrom(frs []FhirResource, numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *DeviceUsage) T_Context(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "context", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "context", resource.Context, htmlAttrs)
}
func (resource *DeviceUsage) T_TimingTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("timingTiming", nil, htmlAttrs)
	}
	return TimingInput("timingTiming", resource.TimingTiming, htmlAttrs)
}
func (resource *DeviceUsage) T_TimingPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("timingPeriod", resource.TimingPeriod, htmlAttrs)
}
func (resource *DeviceUsage) T_TimingDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("timingDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("timingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *DeviceUsage) T_DateAsserted(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateAsserted", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateAsserted", resource.DateAsserted, htmlAttrs)
}
func (resource *DeviceUsage) T_UsageStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeableConceptSelect("usageStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("usageStatus", resource.UsageStatus, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_UsageReason(numUsageReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUsageReason >= len(resource.UsageReason) {
		return CodeableConceptSelect("usageReason["+strconv.Itoa(numUsageReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("usageReason["+strconv.Itoa(numUsageReason)+"]", &resource.UsageReason[numUsageReason], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_InformationSource(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "informationSource", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "informationSource", resource.InformationSource, htmlAttrs)
}
func (resource *DeviceUsage) T_Device(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("device", nil, htmlAttrs)
	}
	return CodeableReferenceInput("device", &resource.Device, htmlAttrs)
}
func (resource *DeviceUsage) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *DeviceUsage) T_BodySite(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("bodySite", nil, htmlAttrs)
	}
	return CodeableReferenceInput("bodySite", resource.BodySite, htmlAttrs)
}
func (resource *DeviceUsage) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceUsage) T_AdherenceCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("adherence.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("adherence.code", &resource.Adherence.Code, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_AdherenceReason(numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Adherence.Reason) {
		return CodeableConceptSelect("adherence.reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("adherence.reason["+strconv.Itoa(numReason)+"]", &resource.Adherence.Reason[numReason], optionsValueSet, htmlAttrs)
}
