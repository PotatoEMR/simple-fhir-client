//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
	Id                     *string                             `json:"id,omitempty"`
	Meta                   *Meta                               `json:"meta,omitempty"`
	ImplicitRules          *string                             `json:"implicitRules,omitempty"`
	Language               *string                             `json:"language,omitempty"`
	Text                   *Narrative                          `json:"text,omitempty"`
	Contained              []Resource                          `json:"contained,omitempty"`
	Extension              []Extension                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                         `json:"modifierExtension,omitempty"`
	Url                    *string                             `json:"url,omitempty"`
	Identifier             []Identifier                        `json:"identifier,omitempty"`
	Version                *string                             `json:"version,omitempty"`
	VersionAlgorithmString *string                             `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                             `json:"versionAlgorithmCoding"`
	Name                   *string                             `json:"name,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	DerivedFromUri         []string                            `json:"derivedFromUri,omitempty"`
	PartOf                 []string                            `json:"partOf,omitempty"`
	Replaces               []string                            `json:"replaces,omitempty"`
	Status                 string                              `json:"status"`
	Experimental           *bool                               `json:"experimental,omitempty"`
	Date                   *string                             `json:"date,omitempty"`
	Publisher              *string                             `json:"publisher,omitempty"`
	Contact                []ContactDetail                     `json:"contact,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	UseContext             []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                   `json:"jurisdiction,omitempty"`
	Purpose                *string                             `json:"purpose,omitempty"`
	Copyright              *string                             `json:"copyright,omitempty"`
	CopyrightLabel         *string                             `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                             `json:"approvalDate,omitempty"`
	LastReviewDate         *string                             `json:"lastReviewDate,omitempty"`
	Code                   *CodeableConcept                    `json:"code,omitempty"`
	Instance               []Reference                         `json:"instance,omitempty"`
	Applicability          []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	PropertyGroup          []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionApplicability struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Condition         *Expression      `json:"condition,omitempty"`
	EffectivePeriod   *Period          `json:"effectivePeriod,omitempty"`
	RelatedArtifact   *RelatedArtifact `json:"relatedArtifact,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionPropertyGroup struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	PriceComponent    []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherChargeItemDefinition ChargeItemDefinition

// on convert struct to json, automatically add resourceType=ChargeItemDefinition
func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItemDefinition: OtherChargeItemDefinition(r),
		ResourceType:              "ChargeItemDefinition",
	})
}
