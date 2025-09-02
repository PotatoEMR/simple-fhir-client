package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	EffectiveDateTime    *string                     `json:"effectiveDateTime"`
	EffectivePeriod      *Period                     `json:"effectivePeriod"`
	EffectiveTiming      *Timing                     `json:"effectiveTiming"`
	EffectiveInstant     *string                     `json:"effectiveInstant"`
	Issued               *string                     `json:"issued,omitempty"`
	Performer            []Reference                 `json:"performer,omitempty"`
	ValueQuantity        *Quantity                   `json:"valueQuantity"`
	ValueCodeableConcept *CodeableConcept            `json:"valueCodeableConcept"`
	ValueString          *string                     `json:"valueString"`
	ValueBoolean         *bool                       `json:"valueBoolean"`
	ValueInteger         *int                        `json:"valueInteger"`
	ValueRange           *Range                      `json:"valueRange"`
	ValueRatio           *Ratio                      `json:"valueRatio"`
	ValueSampledData     *SampledData                `json:"valueSampledData"`
	ValueTime            *string                     `json:"valueTime"`
	ValueDateTime        *string                     `json:"valueDateTime"`
	ValuePeriod          *Period                     `json:"valuePeriod"`
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
	ValueQuantity        *Quantity         `json:"valueQuantity"`
	ValueCodeableConcept *CodeableConcept  `json:"valueCodeableConcept"`
	ValueString          *string           `json:"valueString"`
	ValueBoolean         *bool             `json:"valueBoolean"`
	ValueInteger         *int              `json:"valueInteger"`
	ValueRange           *Range            `json:"valueRange"`
	ValueRatio           *Ratio            `json:"valueRatio"`
	ValueSampledData     *SampledData      `json:"valueSampledData"`
	ValueTime            *string           `json:"valueTime"`
	ValueDateTime        *string           `json:"valueDateTime"`
	ValuePeriod          *Period           `json:"valuePeriod"`
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

func (resource *Observation) ObservationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Observation) ObservationStatus() templ.Component {
	optionsValueSet := VSObservation_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Observation) ObservationCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Observation) ObservationCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *Observation) ObservationDataAbsentReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.DataAbsentReason, optionsValueSet)
}
func (resource *Observation) ObservationInterpretation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("interpretation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("interpretation", &resource.Interpretation[0], optionsValueSet)
}
func (resource *Observation) ObservationBodySite(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet)
}
func (resource *Observation) ObservationMethod(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet)
}
func (resource *Observation) ObservationReferenceRangeType(numReferenceRange int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.ReferenceRange) >= numReferenceRange {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.ReferenceRange[numReferenceRange].Type, optionsValueSet)
}
func (resource *Observation) ObservationReferenceRangeAppliesTo(numReferenceRange int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.ReferenceRange) >= numReferenceRange {
		return CodeableConceptSelect("appliesTo", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appliesTo", &resource.ReferenceRange[numReferenceRange].AppliesTo[0], optionsValueSet)
}
func (resource *Observation) ObservationComponentCode(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Component) >= numComponent {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Component[numComponent].Code, optionsValueSet)
}
func (resource *Observation) ObservationComponentDataAbsentReason(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Component) >= numComponent {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.Component[numComponent].DataAbsentReason, optionsValueSet)
}
func (resource *Observation) ObservationComponentInterpretation(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Component) >= numComponent {
		return CodeableConceptSelect("interpretation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("interpretation", &resource.Component[numComponent].Interpretation[0], optionsValueSet)
}
