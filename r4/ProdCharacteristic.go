package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/ProdCharacteristic
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
