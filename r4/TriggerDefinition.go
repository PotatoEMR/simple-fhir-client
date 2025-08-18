//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id              *string           `json:"id,omitempty"`
	Extension       []Extension       `json:"extension,omitempty"`
	Type            string            `json:"type"`
	Name            *string           `json:"name,omitempty"`
	TimingTiming    *Timing           `json:"timingTiming"`
	TimingReference *Reference        `json:"timingReference"`
	TimingDate      *string           `json:"timingDate"`
	TimingDateTime  *string           `json:"timingDateTime"`
	Data            []DataRequirement `json:"data,omitempty"`
	Condition       *Expression       `json:"condition,omitempty"`
}
