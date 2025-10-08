package main

func FhirResource(fhirVersion string) string {
	return "package " + fhirVersion + `

	type FhirResource interface {
		ToRef() Reference
}
	`
}
