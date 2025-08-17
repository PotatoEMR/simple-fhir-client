//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceAmount
type SubstanceAmount struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity"`
	AmountRange       *Range           `json:"amountRange"`
	AmountString      *string          `json:"amountString"`
	AmountType        *CodeableConcept `json:"amountType,omitempty"`
	AmountText        *string          `json:"amountText,omitempty"`
	ReferenceRange    *Element         `json:"referenceRange,omitempty"`
}
