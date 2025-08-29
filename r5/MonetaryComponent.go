package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/MonetaryComponent
type MonetaryComponent struct {
	Id        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Type      string           `json:"type"`
	Code      *CodeableConcept `json:"code,omitempty"`
	Factor    *float64         `json:"factor,omitempty"`
	Amount    *Money           `json:"amount,omitempty"`
}
