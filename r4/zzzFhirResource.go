package r4

type FhirResource interface {
	ToRef() Reference
}
