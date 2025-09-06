package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Dosage
type Dosage struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 *int              `json:"sequence,omitempty"`
	Text                     *string           `json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept `json:"additionalInstruction,omitempty"`
	PatientInstruction       *string           `json:"patientInstruction,omitempty"`
	Timing                   *Timing           `json:"timing,omitempty"`
	AsNeededBoolean          *bool             `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept  `json:"site,omitempty"`
	Route                    *CodeableConcept  `json:"route,omitempty"`
	Method                   *CodeableConcept  `json:"method,omitempty"`
	DoseAndRate              []Element         `json:"doseAndRate,omitempty"`
	MaxDosePerPeriod         *Ratio            `json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity         `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity         `json:"maxDosePerLifetime,omitempty"`
}
