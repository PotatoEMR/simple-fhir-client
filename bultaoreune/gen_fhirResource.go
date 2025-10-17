package main

func FhirResource(fhirVersion string) string {
	return "package " + fhirVersion + `

	type FhirResource interface {
		ResourceType() string
		ToRef() Reference
	}
	`
}
