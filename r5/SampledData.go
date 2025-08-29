package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/SampledData
type SampledData struct {
	Id           *string     `json:"id,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	Origin       Quantity    `json:"origin"`
	Interval     *float64    `json:"interval,omitempty"`
	IntervalUnit string      `json:"intervalUnit"`
	Factor       *float64    `json:"factor,omitempty"`
	LowerLimit   *float64    `json:"lowerLimit,omitempty"`
	UpperLimit   *float64    `json:"upperLimit,omitempty"`
	Dimensions   int         `json:"dimensions"`
	CodeMap      *string     `json:"codeMap,omitempty"`
	Offsets      *string     `json:"offsets,omitempty"`
	Data         *string     `json:"data,omitempty"`
}
