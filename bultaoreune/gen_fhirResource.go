package main

func FhirResource(fhirVersion string) string {
	return "package " + fhirVersion + `

	type FhirResource interface {
		ToRef() Reference
}

var checkType struct {
					ResourceType string ` + "`json:\"resourceType\"`" + `
				}
	`
	//checkType should not be in here xd just lazy
}
