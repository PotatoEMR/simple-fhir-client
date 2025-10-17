package r4b

type FhirResource interface {
	ToRef() Reference
}

var checkType struct {
	ResourceType string `json:"resourceType"`
}
