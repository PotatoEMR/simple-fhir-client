package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	TimingDateTime    *time.Time        `json:"timingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	RecordedOn        *time.Time        `json:"recordedOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Source            *Reference        `json:"source,omitempty"`
	Device            Reference         `json:"device"`
	ReasonCode        []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference   []Reference       `json:"reasonReference,omitempty"`
	BodySite          *CodeableConcept  `json:"bodySite,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherDeviceUseStatement DeviceUseStatement

// on convert struct to json, automatically add resourceType=DeviceUseStatement
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
func (resource *DeviceUseStatement) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_statement_status

	if resource == nil {
		return CodeSelect("DeviceUseStatement.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceUseStatement.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_TimingDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceUseStatement.TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceUseStatement.TimingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *DeviceUseStatement) T_RecordedOn(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceUseStatement.RecordedOn", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceUseStatement.RecordedOn", resource.RecordedOn, htmlAttrs)
}
func (resource *DeviceUseStatement) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("DeviceUseStatement.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUseStatement.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceUseStatement.BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceUseStatement.BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *DeviceUseStatement) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DeviceUseStatement.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("DeviceUseStatement.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
