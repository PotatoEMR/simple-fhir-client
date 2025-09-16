package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Observation
type Observation struct {
	Id                   *string                     `json:"id,omitempty"`
	Meta                 *Meta                       `json:"meta,omitempty"`
	ImplicitRules        *string                     `json:"implicitRules,omitempty"`
	Language             *string                     `json:"language,omitempty"`
	Text                 *Narrative                  `json:"text,omitempty"`
	Contained            []Resource                  `json:"contained,omitempty"`
	Extension            []Extension                 `json:"extension,omitempty"`
	ModifierExtension    []Extension                 `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                `json:"identifier,omitempty"`
	BasedOn              []Reference                 `json:"basedOn,omitempty"`
	PartOf               []Reference                 `json:"partOf,omitempty"`
	Status               string                      `json:"status"`
	Category             []CodeableConcept           `json:"category,omitempty"`
	Code                 CodeableConcept             `json:"code"`
	Subject              *Reference                  `json:"subject,omitempty"`
	Focus                []Reference                 `json:"focus,omitempty"`
	Encounter            *Reference                  `json:"encounter,omitempty"`
	EffectiveDateTime    *FhirDateTime               `json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                     `json:"effectivePeriod,omitempty"`
	EffectiveTiming      *Timing                     `json:"effectiveTiming,omitempty"`
	EffectiveInstant     *string                     `json:"effectiveInstant,omitempty"`
	Issued               *string                     `json:"issued,omitempty"`
	Performer            []Reference                 `json:"performer,omitempty"`
	ValueQuantity        *Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                     `json:"valueString,omitempty"`
	ValueBoolean         *bool                       `json:"valueBoolean,omitempty"`
	ValueInteger         *int                        `json:"valueInteger,omitempty"`
	ValueRange           *Range                      `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `json:"valueSampledData,omitempty"`
	ValueTime            *string                     `json:"valueTime,omitempty"`
	ValueDateTime        *FhirDateTime               `json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept           `json:"interpretation,omitempty"`
	Note                 []Annotation                `json:"note,omitempty"`
	BodySite             *CodeableConcept            `json:"bodySite,omitempty"`
	Method               *CodeableConcept            `json:"method,omitempty"`
	Specimen             *Reference                  `json:"specimen,omitempty"`
	Device               *Reference                  `json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRange `json:"referenceRange,omitempty"`
	HasMember            []Reference                 `json:"hasMember,omitempty"`
	DerivedFrom          []Reference                 `json:"derivedFrom,omitempty"`
	Component            []ObservationComponent      `json:"component,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Observation
type ObservationReferenceRange struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Low               *Quantity         `json:"low,omitempty"`
	High              *Quantity         `json:"high,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	AppliesTo         []CodeableConcept `json:"appliesTo,omitempty"`
	Age               *Range            `json:"age,omitempty"`
	Text              *string           `json:"text,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Observation
type ObservationComponent struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept   `json:"code"`
	ValueQuantity        *Quantity         `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept  `json:"valueCodeableConcept,omitempty"`
	ValueString          *string           `json:"valueString,omitempty"`
	ValueBoolean         *bool             `json:"valueBoolean,omitempty"`
	ValueInteger         *int              `json:"valueInteger,omitempty"`
	ValueRange           *Range            `json:"valueRange,omitempty"`
	ValueRatio           *Ratio            `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData      `json:"valueSampledData,omitempty"`
	ValueTime            *string           `json:"valueTime,omitempty"`
	ValueDateTime        *FhirDateTime     `json:"valueDateTime,omitempty"`
	ValuePeriod          *Period           `json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept  `json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept `json:"interpretation,omitempty"`
}

type OtherObservation Observation

// on convert struct to json, automatically add resourceType=Observation
func (r Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
}
func (r Observation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Observation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Observation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Observation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *Observation) T_EffectiveInstant(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("effectiveInstant", nil, htmlAttrs)
	}
	return StringInput("effectiveInstant", resource.EffectiveInstant, htmlAttrs)
}
func (resource *Observation) T_Issued(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("issued", nil, htmlAttrs)
	}
	return StringInput("issued", resource.Issued, htmlAttrs)
}
func (resource *Observation) T_ValueCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("valueCodeableConcept", resource.ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ValueString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("valueString", nil, htmlAttrs)
	}
	return StringInput("valueString", resource.ValueString, htmlAttrs)
}
func (resource *Observation) T_ValueBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("valueBoolean", resource.ValueBoolean, htmlAttrs)
}
func (resource *Observation) T_ValueInteger(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("valueInteger", nil, htmlAttrs)
	}
	return IntInput("valueInteger", resource.ValueInteger, htmlAttrs)
}
func (resource *Observation) T_ValueTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("valueTime", nil, htmlAttrs)
	}
	return StringInput("valueTime", resource.ValueTime, htmlAttrs)
}
func (resource *Observation) T_ValueDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("valueDateTime", resource.ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Interpretation(numInterpretation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInterpretation >= len(resource.Interpretation) {
		return CodeableConceptSelect("interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Interpretation[numInterpretation], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Observation) T_BodySite(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Method(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeType(numReferenceRange int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return CodeableConceptSelect("referenceRange["+strconv.Itoa(numReferenceRange)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("referenceRange["+strconv.Itoa(numReferenceRange)+"].type", resource.ReferenceRange[numReferenceRange].Type, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeAppliesTo(numReferenceRange int, numAppliesTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) || numAppliesTo >= len(resource.ReferenceRange[numReferenceRange].AppliesTo) {
		return CodeableConceptSelect("referenceRange["+strconv.Itoa(numReferenceRange)+"].appliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("referenceRange["+strconv.Itoa(numReferenceRange)+"].appliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.ReferenceRange[numReferenceRange].AppliesTo[numAppliesTo], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeText(numReferenceRange int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return StringInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].text", nil, htmlAttrs)
	}
	return StringInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].text", resource.ReferenceRange[numReferenceRange].Text, htmlAttrs)
}
func (resource *Observation) T_ComponentCode(numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].code", &resource.Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentValueCodeableConcept(numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].valueCodeableConcept", resource.Component[numComponent].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentValueString(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].valueString", resource.Component[numComponent].ValueString, htmlAttrs)
}
func (resource *Observation) T_ComponentValueBoolean(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return BoolInput("component["+strconv.Itoa(numComponent)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("component["+strconv.Itoa(numComponent)+"].valueBoolean", resource.Component[numComponent].ValueBoolean, htmlAttrs)
}
func (resource *Observation) T_ComponentValueInteger(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return IntInput("component["+strconv.Itoa(numComponent)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("component["+strconv.Itoa(numComponent)+"].valueInteger", resource.Component[numComponent].ValueInteger, htmlAttrs)
}
func (resource *Observation) T_ComponentValueTime(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].valueTime", resource.Component[numComponent].ValueTime, htmlAttrs)
}
func (resource *Observation) T_ComponentValueDateTime(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return DateTimeInput("component["+strconv.Itoa(numComponent)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("component["+strconv.Itoa(numComponent)+"].valueDateTime", resource.Component[numComponent].ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_ComponentDataAbsentReason(numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].dataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].dataAbsentReason", resource.Component[numComponent].DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentInterpretation(numComponent int, numInterpretation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) || numInterpretation >= len(resource.Component[numComponent].Interpretation) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Component[numComponent].Interpretation[numInterpretation], optionsValueSet, htmlAttrs)
}
