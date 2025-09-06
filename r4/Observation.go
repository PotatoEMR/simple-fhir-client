package r4

//generated with command go run ./bultaoreune -nodownload
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
	EffectiveDateTime    *string                     `json:"effectiveDateTime,omitempty"`
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
	ValueDateTime        *string                     `json:"valueDateTime,omitempty"`
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
	ValueDateTime        *string           `json:"valueDateTime,omitempty"`
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

func (resource *Observation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Observation.Id", nil)
	}
	return StringInput("Observation.Id", resource.Id)
}
func (resource *Observation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Observation.ImplicitRules", nil)
	}
	return StringInput("Observation.ImplicitRules", resource.ImplicitRules)
}
func (resource *Observation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Observation.Language", nil, optionsValueSet)
	}
	return CodeSelect("Observation.Language", resource.Language, optionsValueSet)
}
func (resource *Observation) T_Status() templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("Observation.Status", nil, optionsValueSet)
	}
	return CodeSelect("Observation.Status", &resource.Status, optionsValueSet)
}
func (resource *Observation) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Observation.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Observation) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Observation.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Code", &resource.Code, optionsValueSet)
}
func (resource *Observation) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("Observation.Issued", nil)
	}
	return StringInput("Observation.Issued", resource.Issued)
}
func (resource *Observation) T_DataAbsentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Observation.DataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.DataAbsentReason", resource.DataAbsentReason, optionsValueSet)
}
func (resource *Observation) T_Interpretation(numInterpretation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Interpretation) >= numInterpretation {
		return CodeableConceptSelect("Observation.Interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Interpretation[numInterpretation], optionsValueSet)
}
func (resource *Observation) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Observation.BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.BodySite", resource.BodySite, optionsValueSet)
}
func (resource *Observation) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Observation.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Method", resource.Method, optionsValueSet)
}
func (resource *Observation) T_ReferenceRangeId(numReferenceRange int) templ.Component {

	if resource == nil || len(resource.ReferenceRange) >= numReferenceRange {
		return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Id", nil)
	}
	return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Id", resource.ReferenceRange[numReferenceRange].Id)
}
func (resource *Observation) T_ReferenceRangeType(numReferenceRange int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReferenceRange) >= numReferenceRange {
		return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Type", resource.ReferenceRange[numReferenceRange].Type, optionsValueSet)
}
func (resource *Observation) T_ReferenceRangeAppliesTo(numReferenceRange int, numAppliesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReferenceRange) >= numReferenceRange || len(resource.ReferenceRange[numReferenceRange].AppliesTo) >= numAppliesTo {
		return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.ReferenceRange[numReferenceRange].AppliesTo[numAppliesTo], optionsValueSet)
}
func (resource *Observation) T_ReferenceRangeText(numReferenceRange int) templ.Component {

	if resource == nil || len(resource.ReferenceRange) >= numReferenceRange {
		return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Text", nil)
	}
	return StringInput("Observation.ReferenceRange["+strconv.Itoa(numReferenceRange)+"].Text", resource.ReferenceRange[numReferenceRange].Text)
}
func (resource *Observation) T_ComponentId(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("Observation.Component["+strconv.Itoa(numComponent)+"].Id", resource.Component[numComponent].Id)
}
func (resource *Observation) T_ComponentCode(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Code", &resource.Component[numComponent].Code, optionsValueSet)
}
func (resource *Observation) T_ComponentDataAbsentReason(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].DataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].DataAbsentReason", resource.Component[numComponent].DataAbsentReason, optionsValueSet)
}
func (resource *Observation) T_ComponentInterpretation(numComponent int, numInterpretation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].Interpretation) >= numInterpretation {
		return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Interpretation["+strconv.Itoa(numInterpretation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Observation.Component["+strconv.Itoa(numComponent)+"].Interpretation["+strconv.Itoa(numInterpretation)+"]", &resource.Component[numComponent].Interpretation[numInterpretation], optionsValueSet)
}
