//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

// http://hl7.org/fhir/r4b/StructureDefinition/ProdCharacteristic
type ProdCharacteristic struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Height            *Quantity        `json:"height,omitempty"`
	Width             *Quantity        `json:"width,omitempty"`
	Depth             *Quantity        `json:"depth,omitempty"`
	Weight            *Quantity        `json:"weight,omitempty"`
	NominalVolume     *Quantity        `json:"nominalVolume,omitempty"`
	ExternalDiameter  *Quantity        `json:"externalDiameter,omitempty"`
	Shape             *string          `json:"shape,omitempty"`
	Color             []string         `json:"color,omitempty"`
	Imprint           []string         `json:"imprint,omitempty"`
	Image             []Attachment     `json:"image,omitempty"`
	Scoring           *CodeableConcept `json:"scoring,omitempty"`
}
