//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

// http://hl7.org/fhir/r5/StructureDefinition/Dosage
type Dosage struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 *int              `json:"sequence,omitempty"`
	Text                     *string           `json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept `json:"additionalInstruction,omitempty"`
	PatientInstruction       *string           `json:"patientInstruction,omitempty"`
	Timing                   *Timing           `json:"timing,omitempty"`
	AsNeeded                 *bool             `json:"asNeeded,omitempty"`
	AsNeededFor              []CodeableConcept `json:"asNeededFor,omitempty"`
	Site                     *CodeableConcept  `json:"site,omitempty"`
	Route                    *CodeableConcept  `json:"route,omitempty"`
	Method                   *CodeableConcept  `json:"method,omitempty"`
	DoseAndRate              []Element         `json:"doseAndRate,omitempty"`
	MaxDosePerPeriod         []Ratio           `json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity         `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity         `json:"maxDosePerLifetime,omitempty"`
}
