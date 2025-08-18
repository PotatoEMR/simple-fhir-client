//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Target            []Reference        `json:"target"`
	OccurredPeriod    *Period            `json:"occurredPeriod"`
	OccurredDateTime  *string            `json:"occurredDateTime"`
	Recorded          string             `json:"recorded"`
	Policy            []string           `json:"policy,omitempty"`
	Location          *Reference         `json:"location,omitempty"`
	Reason            []CodeableConcept  `json:"reason,omitempty"`
	Activity          *CodeableConcept   `json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `json:"agent"`
	Entity            []ProvenanceEntity `json:"entity,omitempty"`
	Signature         []Signature        `json:"signature,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
type ProvenanceAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
type ProvenanceEntity struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Role              string      `json:"role"`
	What              Reference   `json:"what"`
}

type OtherProvenance Provenance

// on convert struct to json, automatically add resourceType=Provenance
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}
