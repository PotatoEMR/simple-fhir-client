package r5

type FhirResource interface {
	ToRef() Reference
}
