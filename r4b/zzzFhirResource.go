package r4b

type FhirResource interface {
	ToRef() Reference
}
