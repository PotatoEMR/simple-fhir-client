package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/Coding
type Coding struct {
	Id           *string     `json:"id,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	System       *string     `json:"system,omitempty"`
	Version      *string     `json:"version,omitempty"`
	Code         *string     `json:"code,omitempty"`
	Display      *string     `json:"display,omitempty"`
	UserSelected *bool       `json:"userSelected,omitempty"`
}

func (c *Coding) String() string {
	if c.Display != nil {
		return *c.Display
	} else if c.Code != nil {
		return *c.Code
	} else {
		return "Unnamed Code"
	}
}
