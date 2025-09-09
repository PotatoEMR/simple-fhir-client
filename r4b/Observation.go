package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	EffectiveDateTime    *time.Time                  `json:"effectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDateTime        *time.Time                  `json:"valueDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDateTime        *time.Time        `json:"valueDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Observation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("Observation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Observation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Observation.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Observation.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_EffectiveDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Observation.EffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Observation.EffectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *Observation) T_EffectiveInstant(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Observation.EffectiveInstant", nil, htmlAttrs)
	}
	return StringInput("Observation.EffectiveInstant", resource.EffectiveInstant, htmlAttrs)
}
func (resource *Observation) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Observation.Issued", nil, htmlAttrs)
	}
	return StringInput("Observation.Issued", resource.Issued, htmlAttrs)
}
func (resource *Observation) T_ValueCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Observation.ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.ValueCodeableConcept", resource.ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ValueString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Observation.ValueString", nil, htmlAttrs)
	}
	return StringInput("Observation.ValueString", resource.ValueString, htmlAttrs)
}
func (resource *Observation) T_ValueBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Observation.ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Observation.ValueBoolean", resource.ValueBoolean, htmlAttrs)
}
func (resource *Observation) T_ValueInteger(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Observation.ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Observation.ValueInteger", resource.ValueInteger, htmlAttrs)
}
func (resource *Observation) T_ValueTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Observation.ValueTime", nil, htmlAttrs)
	}
	return StringInput("Observation.ValueTime", resource.ValueTime, htmlAttrs)
}
func (resource *Observation) T_ValueDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Observation.ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Observation.ValueDateTime", resource.ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Observation.DataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.DataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Interpretation(numInterpretation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInterpretation >= len(resource.Interpretation) {
		return CodeableConceptSelect("Observation.Interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Interpretation[numInterpretation], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Observation.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Observation.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Observation) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Observation.BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_Method(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Observation.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeType(numReferenceRange int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Type", resource.ReferenceRange[numReferenceRange].Type, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeAppliesTo(numReferenceRange int, numAppliesTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) || numAppliesTo >= len(resource.ReferenceRange[numReferenceRange].AppliesTo) {
		return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.ReferenceRange[numReferenceRange].AppliesTo[numAppliesTo], optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeText(numReferenceRange int, htmlAttrs string) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Text", nil, htmlAttrs)
	}
	return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Text", resource.ReferenceRange[numReferenceRange].Text, htmlAttrs)
}
func (resource *Observation) T_ComponentCode(numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Code", &resource.Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentValueCodeableConcept(numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].ValueCodeableConcept", resource.Component[numComponent].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentValueString(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueString", resource.Component[numComponent].ValueString, htmlAttrs)
}
func (resource *Observation) T_ComponentValueBoolean(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return BoolInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueBoolean", resource.Component[numComponent].ValueBoolean, htmlAttrs)
}
func (resource *Observation) T_ComponentValueInteger(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return IntInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueInteger", resource.Component[numComponent].ValueInteger, htmlAttrs)
}
func (resource *Observation) T_ComponentValueTime(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueTime", nil, htmlAttrs)
	}
	return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueTime", resource.Component[numComponent].ValueTime, htmlAttrs)
}
func (resource *Observation) T_ComponentValueDateTime(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return DateTimeInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Observation.Component["+strconv.Itoa(numComponent)+"].ValueDateTime", resource.Component[numComponent].ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_ComponentDataAbsentReason(numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].DataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].DataAbsentReason", resource.Component[numComponent].DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentInterpretation(numComponent int, numInterpretation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) || numInterpretation >= len(resource.Component[numComponent].Interpretation) {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Component[numComponent].Interpretation[numInterpretation], optionsValueSet, htmlAttrs)
}
