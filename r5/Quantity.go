package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Quantity
type Quantity struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Value      *float64    `json:"value,omitempty"`
	Comparator *string     `json:"comparator,omitempty"`
	Unit       *string     `json:"unit,omitempty"`
	System     *string     `json:"system,omitempty"`
	Code       *string     `json:"code,omitempty"`
}
