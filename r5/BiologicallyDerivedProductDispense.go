package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	PreparedDate           *time.Time                                    `json:"preparedDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	WhenHandedOver         *time.Time                                    `json:"whenHandedOver,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *BiologicallyDerivedProductDispense) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSBiologicallyderivedproductdispense_status

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProductDispense.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("BiologicallyDerivedProductDispense.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_OriginRelationshipType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.OriginRelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.OriginRelationshipType", resource.OriginRelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_MatchStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.MatchStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.MatchStatus", resource.MatchStatus, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PreparedDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("BiologicallyDerivedProductDispense.PreparedDate", nil, htmlAttrs)
	}
	return DateTimeInput("BiologicallyDerivedProductDispense.PreparedDate", resource.PreparedDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_WhenHandedOver(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("BiologicallyDerivedProductDispense.WhenHandedOver", nil, htmlAttrs)
	}
	return DateTimeInput("BiologicallyDerivedProductDispense.WhenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("BiologicallyDerivedProductDispense.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("BiologicallyDerivedProductDispense.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_UsageInstruction(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProductDispense.UsageInstruction", nil, htmlAttrs)
	}
	return StringInput("BiologicallyDerivedProductDispense.UsageInstruction", resource.UsageInstruction, htmlAttrs)
}
func (resource *BiologicallyDerivedProductDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("BiologicallyDerivedProductDispense.Performer."+strconv.Itoa(numPerformer)+"..Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProductDispense.Performer."+strconv.Itoa(numPerformer)+"..Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
