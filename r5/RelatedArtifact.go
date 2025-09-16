package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	Type              string            `json:"type"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
	Label             *string           `json:"label,omitempty"`
	Display           *string           `json:"display,omitempty"`
	Citation          *string           `json:"citation,omitempty"`
	Document          *Attachment       `json:"document,omitempty"`
	Resource          *string           `json:"resource,omitempty"`
	ResourceReference *Reference        `json:"resourceReference,omitempty"`
	PublicationStatus *string           `json:"publicationStatus,omitempty"`
	PublicationDate   *FhirDate         `json:"publicationDate,omitempty"`
}
