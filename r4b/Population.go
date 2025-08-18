//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/Population
type Population struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AgeRange               *Range           `json:"ageRange"`
	AgeCodeableConcept     *CodeableConcept `json:"ageCodeableConcept"`
	Gender                 *CodeableConcept `json:"gender,omitempty"`
	Race                   *CodeableConcept `json:"race,omitempty"`
	PhysiologicalCondition *CodeableConcept `json:"physiologicalCondition,omitempty"`
}
