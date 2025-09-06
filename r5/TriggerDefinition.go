package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	Type              string            `json:"type"`
	Name              *string           `json:"name,omitempty"`
	Code              *CodeableConcept  `json:"code,omitempty"`
	SubscriptionTopic *string           `json:"subscriptionTopic,omitempty"`
	TimingTiming      *Timing           `json:"timingTiming,omitempty"`
	TimingReference   *Reference        `json:"timingReference,omitempty"`
	TimingDate        *string           `json:"timingDate,omitempty"`
	TimingDateTime    *string           `json:"timingDateTime,omitempty"`
	Data              []DataRequirement `json:"data,omitempty"`
	Condition         *Expression       `json:"condition,omitempty"`
}
