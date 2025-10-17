package r4b

type FhirResource interface {
	ResourceType() string
	ToRef() Reference
}
