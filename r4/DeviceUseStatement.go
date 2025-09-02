package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/DeviceUseStatement
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
	TimingTiming      *Timing           `json:"timingTiming"`
	TimingPeriod      *Period           `json:"timingPeriod"`
	TimingDateTime    *string           `json:"timingDateTime"`
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

func (resource *DeviceUseStatement) DeviceUseStatementLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DeviceUseStatement) DeviceUseStatementStatus() templ.Component {
	optionsValueSet := VSDevice_statement_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *DeviceUseStatement) DeviceUseStatementReasonCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *DeviceUseStatement) DeviceUseStatementBodySite(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet)
}
