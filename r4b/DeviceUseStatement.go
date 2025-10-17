package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	BasedOn           []Reference       `json:"basedOn,omitempty"`
	Status            string            `json:"status"`
	Subject           Reference         `json:"subject"`
	DerivedFrom       []Reference       `json:"derivedFrom,omitempty"`
	TimingTiming      *Timing           `json:"timingTiming,omitempty"`
	TimingPeriod      *Period           `json:"timingPeriod,omitempty"`
	TimingDateTime    *FhirDateTime     `json:"timingDateTime,omitempty"`
	RecordedOn        *FhirDateTime     `json:"recordedOn,omitempty"`
	Source            *Reference        `json:"source,omitempty"`
	Device            Reference         `json:"device"`
	ReasonCode        []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference   []Reference       `json:"reasonReference,omitempty"`
	BodySite          *CodeableConcept  `json:"bodySite,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherDeviceUseStatement DeviceUseStatement

// struct -> json, automatically add resourceType=Patient
func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUseStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUseStatement: OtherDeviceUseStatement(r),
		ResourceType:            "DeviceUseStatement",
	})
}

func (r DeviceUseStatement) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceUseStatement/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceUseStatement"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r DeviceUseStatement) ResourceType() string {
	return "DeviceUseStatement"
}

func (resource *DeviceUseStatement) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DeviceUseStatement) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_statement_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *DeviceUseStatement) T_DerivedFrom(frs []FhirResource, numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *DeviceUseStatement) T_TimingTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("timingTiming", nil, htmlAttrs)
	}
	return TimingInput("timingTiming", resource.TimingTiming, htmlAttrs)
}
func (resource *DeviceUseStatement) T_TimingPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("timingPeriod", resource.TimingPeriod, htmlAttrs)
}
func (resource *DeviceUseStatement) T_TimingDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("timingDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("timingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *DeviceUseStatement) T_RecordedOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recordedOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recordedOn", resource.RecordedOn, htmlAttrs)
}
func (resource *DeviceUseStatement) T_Source(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source", resource.Source, htmlAttrs)
}
func (resource *DeviceUseStatement) T_Device(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "device", &resource.Device, htmlAttrs)
}
func (resource *DeviceUseStatement) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *DeviceUseStatement) T_BodySite(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
