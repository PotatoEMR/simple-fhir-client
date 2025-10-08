package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/Identifier
type Identifier struct {
	Id        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Use       *string          `json:"use,omitempty"`
	Type      *CodeableConcept `json:"type,omitempty"`
	System    *string          `json:"system,omitempty"`
	Value     *string          `json:"value,omitempty"`
	Period    *Period          `json:"period,omitempty"`
	Assigner  *Reference       `json:"assigner,omitempty"`
}

func (r Identifier) String() string {
	ret := ""
	if r.Value != nil {
		ret = ret + *r.Value
	}
	if r.System != nil {
		ret = ret + " (" + *r.System + ")"
	}
	if r.Period != nil {
		ret = ret + " (" + "period string todo" + ")"
	}
	if r.Use != nil {
		ret = ret + " (Use: " + *r.Use + ")"
	}
	return ret
}
