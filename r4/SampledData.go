//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

// http://hl7.org/fhir/r4/StructureDefinition/SampledData
type SampledData struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Origin     Quantity    `json:"origin"`
	Period     float64     `json:"period"`
	Factor     *float64    `json:"factor,omitempty"`
	LowerLimit *float64    `json:"lowerLimit,omitempty"`
	UpperLimit *float64    `json:"upperLimit,omitempty"`
	Dimensions int         `json:"dimensions"`
	Data       *string     `json:"data,omitempty"`
}
