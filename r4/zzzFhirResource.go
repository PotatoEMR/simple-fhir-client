package r4

type FhirResource interface {
	ResourceType() string
	ToRef() Reference
}
