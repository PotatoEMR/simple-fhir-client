package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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
	TimingDateTime    *string               `json:"timingDateTime,omitempty"`
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

func (resource *DeviceUsage) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceUsage.Id", nil)
	}
	return StringInput("DeviceUsage.Id", resource.Id)
}
func (resource *DeviceUsage) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceUsage.ImplicitRules", nil)
	}
	return StringInput("DeviceUsage.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceUsage) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceUsage.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceUsage.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceUsage) T_Status() templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeSelect("DeviceUsage.Status", nil, optionsValueSet)
	}
	return CodeSelect("DeviceUsage.Status", &resource.Status, optionsValueSet)
}
func (resource *DeviceUsage) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("DeviceUsage.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUsage.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *DeviceUsage) T_DateAsserted() templ.Component {

	if resource == nil {
		return StringInput("DeviceUsage.DateAsserted", nil)
	}
	return StringInput("DeviceUsage.DateAsserted", resource.DateAsserted)
}
func (resource *DeviceUsage) T_UsageStatus() templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeableConceptSelect("DeviceUsage.UsageStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUsage.UsageStatus", resource.UsageStatus, optionsValueSet)
}
func (resource *DeviceUsage) T_UsageReason(numUsageReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.UsageReason) >= numUsageReason {
		return CodeableConceptSelect("DeviceUsage.UsageReason["+strconv.Itoa(numUsageReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUsage.UsageReason["+strconv.Itoa(numUsageReason)+"]", &resource.UsageReason[numUsageReason], optionsValueSet)
}
func (resource *DeviceUsage) T_AdherenceId() templ.Component {

	if resource == nil {
		return StringInput("DeviceUsage.Adherence.Id", nil)
	}
	return StringInput("DeviceUsage.Adherence.Id", resource.Adherence.Id)
}
func (resource *DeviceUsage) T_AdherenceCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceUsage.Adherence.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUsage.Adherence.Code", &resource.Adherence.Code, optionsValueSet)
}
func (resource *DeviceUsage) T_AdherenceReason(numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Adherence.Reason) >= numReason {
		return CodeableConceptSelect("DeviceUsage.Adherence.Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUsage.Adherence.Reason["+strconv.Itoa(numReason)+"]", &resource.Adherence.Reason[numReason], optionsValueSet)
}
