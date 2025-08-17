//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
