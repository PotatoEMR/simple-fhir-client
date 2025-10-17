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

// struct -> json, automatically add resourceType=Patient
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
func (r BiologicallyDerivedProductDispense) ResourceType() string {
	return "BiologicallyDerivedProductDispense"
}

func (resource *BiologicallyDerivedProductDispense) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
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
func (resource *BiologicallyDerivedProductDispense) T_Product(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "product", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "product", &resource.Product, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_MatchStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("matchStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("matchStatus", resource.MatchStatus, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PreparedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("preparedDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("preparedDate", resource.PreparedDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_WhenHandedOver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("whenHandedOver", nil, htmlAttrs)
	}
	return FhirDateTimeInput("whenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Destination(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "destination", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "destination", resource.Destination, htmlAttrs)
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
func (resource *BiologicallyDerivedProductDispense) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
