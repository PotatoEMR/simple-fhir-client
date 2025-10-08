package main

// string resource name -> its string rep function
// and resource list that need import string
func ResourceStrings() (map[string]string, []string) {
	rs := make(map[string]string)
	var ri []string //could probably just check if key in resourceString map, but idk maybe func doesnt need to import strings

	rs["Address"] = `
	func (a Address) String() string {
		parts := []string{}

		if a.Text != nil && *a.Text != "" {
			parts = append(parts, *a.Text)
		} else {
			if len(a.Line) > 0 {
				parts = append(parts, strings.Join(a.Line, ", "))
			}
			if a.City != nil && *a.City != "" {
				parts = append(parts, *a.City)
			}
			if a.District != nil && *a.District != "" {
				parts = append(parts, *a.District)
			}
			if a.State != nil && *a.State != "" {
				parts = append(parts, *a.State)
			}
			if a.PostalCode != nil && *a.PostalCode != "" {
				parts = append(parts, *a.PostalCode)
			}
			if a.Country != nil && *a.Country != "" {
				parts = append(parts, *a.Country)
			}
		}
		return strings.Join(parts, ", ")
		}
	`

	ri = append(ri, "CodeableConcept")
	rs["CodeableConcept"] = `
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
	}`

	//don't need strings
	rs["Coding"] = `
	func (c *Coding) String() string {
		if c == nil {
			return ""
		}
		if c.Display != nil {
			return *c.Display
		} else if c.Code != nil {
			return *c.Code
		} else {
			return "Unnamed Code"
		}
	}`

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

	rs["Identifier"] = `
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
	`

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

	rs["Reference"] = `
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
	}`

	return rs, ri
}
