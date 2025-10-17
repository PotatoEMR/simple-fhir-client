package r5

type FhirResource interface {
	ToRef() Reference
}

var checkType struct {
	ResourceType string `json:"resourceType"`
}
