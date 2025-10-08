package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/Reference
type Reference struct {
	Id         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Reference  *string     `json:"reference,omitempty"`
	Type       *string     `json:"type,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display    *string     `json:"display,omitempty"`
}

func (r Reference) String() string {
	if r.Display != nil {
		return *r.Display
	} else if r.Identifier != nil {
		return r.Identifier.String()
	} else if r.Id != nil {
		return *r.Id
	} else if r.Reference != nil {
		return *r.Reference
	} else if r.Type != nil {
		return *r.Type
	}
	return ""
}
