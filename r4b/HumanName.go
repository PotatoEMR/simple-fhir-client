//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "strings"

// http://hl7.org/fhir/r4b/StructureDefinition/HumanName
type HumanName struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Use       *string     `json:"use,omitempty"`
	Text      *string     `json:"text,omitempty"`
	Family    *string     `json:"family,omitempty"`
	Given     []string    `json:"given,omitempty"`
	Prefix    []string    `json:"prefix,omitempty"`
	Suffix    []string    `json:"suffix,omitempty"`
	Period    *Period     `json:"period,omitempty"`
}

func (hn HumanName) String() string {
	if hn.Text != nil {
		return *hn.Text
	}
	var b strings.Builder
	b.WriteString(strings.Join(hn.Prefix, " "))
	b.WriteString(strings.Join(hn.Given, " "))
	b.WriteString(" ")
	if hn.Family != nil {
		b.WriteString(*hn.Family)
	}
	b.WriteString(" ")
	b.WriteString(strings.Join(hn.Suffix, ","))
	if hn.Use != nil {
		if *hn.Use == "nickname" || *hn.Use == "anonymous" {
			b.WriteString("(" + *hn.Use + ")")
		}
		if *hn.Use == "old" {
			b.WriteString("(old name)")
		}
	}
	return b.String()
}
