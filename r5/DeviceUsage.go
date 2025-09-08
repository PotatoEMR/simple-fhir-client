package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	TimingDateTime    *time.Time            `json:"timingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	DateAsserted      *time.Time            `json:"dateAsserted,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *DeviceUsage) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeSelect("DeviceUsage.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceUsage.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("DeviceUsage.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUsage.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_TimingDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceUsage.TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceUsage.TimingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *DeviceUsage) T_DateAsserted(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceUsage.DateAsserted", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceUsage.DateAsserted", resource.DateAsserted, htmlAttrs)
}
func (resource *DeviceUsage) T_UsageStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSDeviceusage_status

	if resource == nil {
		return CodeableConceptSelect("DeviceUsage.UsageStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUsage.UsageStatus", resource.UsageStatus, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_UsageReason(numUsageReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numUsageReason >= len(resource.UsageReason) {
		return CodeableConceptSelect("DeviceUsage.UsageReason."+strconv.Itoa(numUsageReason)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUsage.UsageReason."+strconv.Itoa(numUsageReason)+".", &resource.UsageReason[numUsageReason], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DeviceUsage.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("DeviceUsage.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceUsage) T_AdherenceCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceUsage.Adherence.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUsage.Adherence.Code", &resource.Adherence.Code, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUsage) T_AdherenceReason(numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReason >= len(resource.Adherence.Reason) {
		return CodeableConceptSelect("DeviceUsage.Adherence.Reason."+strconv.Itoa(numReason)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUsage.Adherence.Reason."+strconv.Itoa(numReason)+".", &resource.Adherence.Reason[numReason], optionsValueSet, htmlAttrs)
}
