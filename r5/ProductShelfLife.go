package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/ProductShelfLife
type ProductShelfLife struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Type                         *CodeableConcept  `json:"type,omitempty"`
	PeriodDuration               *Duration         `json:"periodDuration,omitempty"`
	PeriodString                 *string           `json:"periodString,omitempty"`
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}
