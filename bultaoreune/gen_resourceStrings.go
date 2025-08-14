package main

// string resource name -> its string rep function
// and resource list that need import string
func ResourceStrings() (map[string]string, []string) {
	rs := make(map[string]string)
	var ri []string //could probably just check if key in resourceString map, but idk maybe func doesnt need to import strings

	ri = append(ri, "HumanName")
	rs["HumanName"] = `
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
	}`

	ri = append(ri, "OperationOutcome")
	rs["OperationOutcome"] = `
	func (oo OperationOutcome) String() string {
		var b strings.Builder
		b.WriteString("OperationOutcome ")
		for _, v := range oo.Issue {
			b.WriteString(v.String())
			b.WriteString(" ")
		}
		return b.String()
	}`

	rs["OperationOutcomeIssue"] = `
	func (ooi OperationOutcomeIssue) String() string {
		ret := "Issue: "
		if ooi.Diagnostics != nil {
			ret += *ooi.Diagnostics
		}
		ret += " "
		errorPath := append(ooi.Location, ooi.Expression...) //maybe error fields? idk
		ret += strings.Join(errorPath, ", ")
		return ret
	}`

	ri = append(ri, "Patient")
	rs["Patient"] = `
		func (p Patient) String() string {
		var b strings.Builder
		lastNewline := len(p.Name) - 1
		for i, v := range p.Name {
			b.WriteString(v.String())
			if i != lastNewline {
				b.WriteString("\n")
			}
		}
		if b.String() == "" {
			return "Unnamed Patient"
		}
		return b.String()
	}`

	return rs, ri
}
