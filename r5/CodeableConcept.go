package r5

import "strings"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/CodeableConcept
type CodeableConcept struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Coding    []Coding    `json:"coding,omitempty"`
	Text      *string     `json:"text,omitempty"`
}

func (cc *CodeableConcept) String() string {
	if cc == nil {
		return ""
	}
	if cc.Text != nil {
		return *cc.Text
	}
	var b strings.Builder
	lastCoding := len(cc.Coding) - 1
	for i, v := range cc.Coding {
		b.WriteString(v.String())
		if i != lastCoding {
			b.WriteString(", ")
		}
	}
	if b.String() == "" {
		return "Unnamed CodeableConcept"
	}
	return b.String()
}
