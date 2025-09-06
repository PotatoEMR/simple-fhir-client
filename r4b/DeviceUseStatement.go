package r4b

//generated with command go run ./bultaoreune -nodownload
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
	TimingDateTime    *string           `json:"timingDateTime,omitempty"`
	RecordedOn        *string           `json:"recordedOn,omitempty"`
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

func (resource *DeviceUseStatement) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceUseStatement.Id", nil)
	}
	return StringInput("DeviceUseStatement.Id", resource.Id)
}
func (resource *DeviceUseStatement) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceUseStatement.ImplicitRules", nil)
	}
	return StringInput("DeviceUseStatement.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceUseStatement) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceUseStatement.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceUseStatement.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceUseStatement) T_Status() templ.Component {
	optionsValueSet := VSDevice_statement_status

	if resource == nil {
		return CodeSelect("DeviceUseStatement.Status", nil, optionsValueSet)
	}
	return CodeSelect("DeviceUseStatement.Status", &resource.Status, optionsValueSet)
}
func (resource *DeviceUseStatement) T_RecordedOn() templ.Component {

	if resource == nil {
		return StringInput("DeviceUseStatement.RecordedOn", nil)
	}
	return StringInput("DeviceUseStatement.RecordedOn", resource.RecordedOn)
}
func (resource *DeviceUseStatement) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("DeviceUseStatement.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUseStatement.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *DeviceUseStatement) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceUseStatement.BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceUseStatement.BodySite", resource.BodySite, optionsValueSet)
}
