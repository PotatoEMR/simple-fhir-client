//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispense struct {
	Id                     *string                                       `json:"id,omitempty"`
	Meta                   *Meta                                         `json:"meta,omitempty"`
	ImplicitRules          *string                                       `json:"implicitRules,omitempty"`
	Language               *string                                       `json:"language,omitempty"`
	Text                   *Narrative                                    `json:"text,omitempty"`
	Contained              []Resource                                    `json:"contained,omitempty"`
	Extension              []Extension                                   `json:"extension,omitempty"`
	ModifierExtension      []Extension                                   `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                                  `json:"identifier,omitempty"`
	BasedOn                []Reference                                   `json:"basedOn,omitempty"`
	PartOf                 []Reference                                   `json:"partOf,omitempty"`
	Status                 string                                        `json:"status"`
	OriginRelationshipType *CodeableConcept                              `json:"originRelationshipType,omitempty"`
	Product                Reference                                     `json:"product"`
	Patient                Reference                                     `json:"patient"`
	MatchStatus            *CodeableConcept                              `json:"matchStatus,omitempty"`
	Performer              []BiologicallyDerivedProductDispensePerformer `json:"performer,omitempty"`
	Location               *Reference                                    `json:"location,omitempty"`
	Quantity               *Quantity                                     `json:"quantity,omitempty"`
	PreparedDate           *string                                       `json:"preparedDate,omitempty"`
	WhenHandedOver         *string                                       `json:"whenHandedOver,omitempty"`
	Destination            *Reference                                    `json:"destination,omitempty"`
	Note                   []Annotation                                  `json:"note,omitempty"`
	UsageInstruction       *string                                       `json:"usageInstruction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherBiologicallyDerivedProductDispense BiologicallyDerivedProductDispense

// on convert struct to json, automatically add resourceType=BiologicallyDerivedProductDispense
func (r BiologicallyDerivedProductDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProductDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProductDispense: OtherBiologicallyDerivedProductDispense(r),
		ResourceType:                            "BiologicallyDerivedProductDispense",
	})
}
