package r5

type FhirResource interface {
	ResourceType() string
	ToRef() Reference
}
