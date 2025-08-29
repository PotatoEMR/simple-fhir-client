package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Observation
type Observation struct {
	Id                    *string                     `json:"id,omitempty"`
	Meta                  *Meta                       `json:"meta,omitempty"`
	ImplicitRules         *string                     `json:"implicitRules,omitempty"`
	Language              *string                     `json:"language,omitempty"`
	Text                  *Narrative                  `json:"text,omitempty"`
	Contained             []Resource                  `json:"contained,omitempty"`
	Extension             []Extension                 `json:"extension,omitempty"`
	ModifierExtension     []Extension                 `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                `json:"identifier,omitempty"`
	InstantiatesCanonical *string                     `json:"instantiatesCanonical"`
	InstantiatesReference *Reference                  `json:"instantiatesReference"`
	BasedOn               []Reference                 `json:"basedOn,omitempty"`
	TriggeredBy           []ObservationTriggeredBy    `json:"triggeredBy,omitempty"`
	PartOf                []Reference                 `json:"partOf,omitempty"`
	Status                string                      `json:"status"`
	Category              []CodeableConcept           `json:"category,omitempty"`
	Code                  CodeableConcept             `json:"code"`
	Subject               *Reference                  `json:"subject,omitempty"`
	Focus                 []Reference                 `json:"focus,omitempty"`
	Encounter             *Reference                  `json:"encounter,omitempty"`
	EffectiveDateTime     *string                     `json:"effectiveDateTime"`
	EffectivePeriod       *Period                     `json:"effectivePeriod"`
	EffectiveTiming       *Timing                     `json:"effectiveTiming"`
	EffectiveInstant      *string                     `json:"effectiveInstant"`
	Issued                *string                     `json:"issued,omitempty"`
	Performer             []Reference                 `json:"performer,omitempty"`
	ValueQuantity         *Quantity                   `json:"valueQuantity"`
	ValueCodeableConcept  *CodeableConcept            `json:"valueCodeableConcept"`
	ValueString           *string                     `json:"valueString"`
	ValueBoolean          *bool                       `json:"valueBoolean"`
	ValueInteger          *int                        `json:"valueInteger"`
	ValueRange            *Range                      `json:"valueRange"`
	ValueRatio            *Ratio                      `json:"valueRatio"`
	ValueSampledData      *SampledData                `json:"valueSampledData"`
	ValueTime             *string                     `json:"valueTime"`
	ValueDateTime         *string                     `json:"valueDateTime"`
	ValuePeriod           *Period                     `json:"valuePeriod"`
	ValueAttachment       *Attachment                 `json:"valueAttachment"`
	ValueReference        *Reference                  `json:"valueReference"`
	DataAbsentReason      *CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation        []CodeableConcept           `json:"interpretation,omitempty"`
	Note                  []Annotation                `json:"note,omitempty"`
	BodySite              *CodeableConcept            `json:"bodySite,omitempty"`
	BodyStructure         *Reference                  `json:"bodyStructure,omitempty"`
	Method                *CodeableConcept            `json:"method,omitempty"`
	Specimen              *Reference                  `json:"specimen,omitempty"`
	Device                *Reference                  `json:"device,omitempty"`
	ReferenceRange        []ObservationReferenceRange `json:"referenceRange,omitempty"`
	HasMember             []Reference                 `json:"hasMember,omitempty"`
	DerivedFrom           []Reference                 `json:"derivedFrom,omitempty"`
	Component             []ObservationComponent      `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Observation
type ObservationTriggeredBy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Observation       Reference   `json:"observation"`
	Type              string      `json:"type"`
	Reason            *string     `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Observation
type ObservationReferenceRange struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Low               *Quantity         `json:"low,omitempty"`
	High              *Quantity         `json:"high,omitempty"`
	NormalValue       *CodeableConcept  `json:"normalValue,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	AppliesTo         []CodeableConcept `json:"appliesTo,omitempty"`
	Age               *Range            `json:"age,omitempty"`
	Text              *string           `json:"text,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Observation
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
	ValueAttachment      *Attachment       `json:"valueAttachment"`
	ValueReference       *Reference        `json:"valueReference"`
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Observation) ObservationStatus() templ.Component {
	optionsValueSet := VSObservation_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Observation) ObservationTriggeredByType(numTriggeredBy int) templ.Component {
	optionsValueSet := VSObservation_triggeredbytype
	currentVal := ""
	if resource != nil && len(resource.TriggeredBy) >= numTriggeredBy {
		currentVal = resource.TriggeredBy[numTriggeredBy].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
