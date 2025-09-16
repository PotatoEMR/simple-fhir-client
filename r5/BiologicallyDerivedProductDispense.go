package r5

//generated with command go run ./bultaoreune
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
	PreparedDate           *FhirDateTime                                 `json:"preparedDate,omitempty"`
	WhenHandedOver         *FhirDateTime                                 `json:"whenHandedOver,omitempty"`
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
func (r BiologicallyDerivedProductDispense) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "BiologicallyDerivedProductDispense/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "BiologicallyDerivedProductDispense"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *BiologicallyDerivedProductDispense) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSBiologicallyderivedproductdispense_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_OriginRelationshipType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("originRelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("originRelationshipType", resource.OriginRelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_MatchStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("matchStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("matchStatus", resource.MatchStatus, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PreparedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("preparedDate", nil, htmlAttrs)
	}
	return DateTimeInput("preparedDate", resource.PreparedDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_WhenHandedOver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("whenHandedOver", nil, htmlAttrs)
	}
	return DateTimeInput("whenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_UsageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usageInstruction", nil, htmlAttrs)
	}
	return StringInput("usageInstruction", resource.UsageInstruction, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
