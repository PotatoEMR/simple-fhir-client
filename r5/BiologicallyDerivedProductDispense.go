package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *BiologicallyDerivedProductDispense) T_Id() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.Id", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.Id", resource.Id)
}
func (resource *BiologicallyDerivedProductDispense) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.ImplicitRules", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.ImplicitRules", resource.ImplicitRules)
}
func (resource *BiologicallyDerivedProductDispense) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProductDispense.Language", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProductDispense.Language", resource.Language, optionsValueSet)
}
func (resource *BiologicallyDerivedProductDispense) T_Status() templ.Component {
	optionsValueSet := VSBiologicallyderivedproductdispense_status

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProductDispense.Status", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProductDispense.Status", &resource.Status, optionsValueSet)
}
func (resource *BiologicallyDerivedProductDispense) T_OriginRelationshipType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.OriginRelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.OriginRelationshipType", resource.OriginRelationshipType, optionsValueSet)
}
func (resource *BiologicallyDerivedProductDispense) T_MatchStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.MatchStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.MatchStatus", resource.MatchStatus, optionsValueSet)
}
func (resource *BiologicallyDerivedProductDispense) T_PreparedDate() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.PreparedDate", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.PreparedDate", resource.PreparedDate)
}
func (resource *BiologicallyDerivedProductDispense) T_WhenHandedOver() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.WhenHandedOver", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.WhenHandedOver", resource.WhenHandedOver)
}
func (resource *BiologicallyDerivedProductDispense) T_UsageInstruction() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.UsageInstruction", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.UsageInstruction", resource.UsageInstruction)
}
func (resource *BiologicallyDerivedProductDispense) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("BiologicallyDerivedProductDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("BiologicallyDerivedProductDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *BiologicallyDerivedProductDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
