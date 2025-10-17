package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Observation
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

// http://hl7.org/fhir/r4/StructureDefinition/Observation
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

// http://hl7.org/fhir/r4/StructureDefinition/Observation
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

// struct -> json, automatically add resourceType=Patient
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
func (r Observation) ResourceType() string {
	return "Observation"
}

func (resource *Observation) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Observation) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
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
func (resource *Observation) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Observation) T_Focus(frs []FhirResource, numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], htmlAttrs)
}
func (resource *Observation) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Observation) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *Observation) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *Observation) T_EffectiveTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("effectiveTiming", nil, htmlAttrs)
	}
	return TimingInput("effectiveTiming", resource.EffectiveTiming, htmlAttrs)
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
func (resource *Observation) T_Performer(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", &resource.Performer[numPerformer], htmlAttrs)
}
func (resource *Observation) T_ValueQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("valueQuantity", resource.ValueQuantity, optionsValueSet, htmlAttrs)
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
func (resource *Observation) T_ValueRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("valueRange", nil, htmlAttrs)
	}
	return RangeInput("valueRange", resource.ValueRange, htmlAttrs)
}
func (resource *Observation) T_ValueRatio(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RatioInput("valueRatio", nil, htmlAttrs)
	}
	return RatioInput("valueRatio", resource.ValueRatio, htmlAttrs)
}
func (resource *Observation) T_ValueSampledData(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return SampledDataInput("valueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("valueSampledData", resource.ValueSampledData, htmlAttrs)
}
func (resource *Observation) T_ValueTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("valueTime", nil, htmlAttrs)
	}
	return StringInput("valueTime", resource.ValueTime, htmlAttrs)
}
func (resource *Observation) T_ValueDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("valueDateTime", resource.ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_ValuePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("valuePeriod", resource.ValuePeriod, htmlAttrs)
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
func (resource *Observation) T_Specimen(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "specimen", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "specimen", resource.Specimen, htmlAttrs)
}
func (resource *Observation) T_Device(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "device", resource.Device, htmlAttrs)
}
func (resource *Observation) T_HasMember(frs []FhirResource, numHasMember int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHasMember >= len(resource.HasMember) {
		return ReferenceInput(frs, "hasMember["+strconv.Itoa(numHasMember)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "hasMember["+strconv.Itoa(numHasMember)+"]", &resource.HasMember[numHasMember], htmlAttrs)
}
func (resource *Observation) T_DerivedFrom(frs []FhirResource, numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeLow(numReferenceRange int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return QuantityInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].low", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].low", resource.ReferenceRange[numReferenceRange].Low, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ReferenceRangeHigh(numReferenceRange int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return QuantityInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].high", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].high", resource.ReferenceRange[numReferenceRange].High, optionsValueSet, htmlAttrs)
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
func (resource *Observation) T_ReferenceRangeAge(numReferenceRange int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferenceRange >= len(resource.ReferenceRange) {
		return RangeInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].age", nil, htmlAttrs)
	}
	return RangeInput("referenceRange["+strconv.Itoa(numReferenceRange)+"].age", resource.ReferenceRange[numReferenceRange].Age, htmlAttrs)
}
func (resource *Observation) T_ComponentCode(numComponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].code", &resource.Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *Observation) T_ComponentValueQuantity(numComponent int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return QuantityInput("component["+strconv.Itoa(numComponent)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("component["+strconv.Itoa(numComponent)+"].valueQuantity", resource.Component[numComponent].ValueQuantity, optionsValueSet, htmlAttrs)
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
func (resource *Observation) T_ComponentValueRange(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return RangeInput("component["+strconv.Itoa(numComponent)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("component["+strconv.Itoa(numComponent)+"].valueRange", resource.Component[numComponent].ValueRange, htmlAttrs)
}
func (resource *Observation) T_ComponentValueRatio(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return RatioInput("component["+strconv.Itoa(numComponent)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("component["+strconv.Itoa(numComponent)+"].valueRatio", resource.Component[numComponent].ValueRatio, htmlAttrs)
}
func (resource *Observation) T_ComponentValueSampledData(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return SampledDataInput("component["+strconv.Itoa(numComponent)+"].valueSampledData", nil, htmlAttrs)
	}
	return SampledDataInput("component["+strconv.Itoa(numComponent)+"].valueSampledData", resource.Component[numComponent].ValueSampledData, htmlAttrs)
}
func (resource *Observation) T_ComponentValueTime(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].valueTime", resource.Component[numComponent].ValueTime, htmlAttrs)
}
func (resource *Observation) T_ComponentValueDateTime(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return FhirDateTimeInput("component["+strconv.Itoa(numComponent)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("component["+strconv.Itoa(numComponent)+"].valueDateTime", resource.Component[numComponent].ValueDateTime, htmlAttrs)
}
func (resource *Observation) T_ComponentValuePeriod(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return PeriodInput("component["+strconv.Itoa(numComponent)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("component["+strconv.Itoa(numComponent)+"].valuePeriod", resource.Component[numComponent].ValuePeriod, htmlAttrs)
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
