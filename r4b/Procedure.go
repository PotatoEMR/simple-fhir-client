package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
type Procedure struct {
	Id                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []Resource             `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	PartOf                []Reference            `json:"partOf,omitempty"`
	Status                string                 `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	Category              *CodeableConcept       `json:"category,omitempty"`
	Code                  *CodeableConcept       `json:"code,omitempty"`
	Subject               Reference              `json:"subject"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	PerformedDateTime     *string                `json:"performedDateTime"`
	PerformedPeriod       *Period                `json:"performedPeriod"`
	PerformedString       *string                `json:"performedString"`
	PerformedAge          *Age                   `json:"performedAge"`
	PerformedRange        *Range                 `json:"performedRange"`
	Recorder              *Reference             `json:"recorder,omitempty"`
	Asserter              *Reference             `json:"asserter,omitempty"`
	Performer             []ProcedurePerformer   `json:"performer,omitempty"`
	Location              *Reference             `json:"location,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	BodySite              []CodeableConcept      `json:"bodySite,omitempty"`
	Outcome               *CodeableConcept       `json:"outcome,omitempty"`
	Report                []Reference            `json:"report,omitempty"`
	Complication          []CodeableConcept      `json:"complication,omitempty"`
	ComplicationDetail    []Reference            `json:"complicationDetail,omitempty"`
	FollowUp              []CodeableConcept      `json:"followUp,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDevice `json:"focalDevice,omitempty"`
	UsedReference         []Reference            `json:"usedReference,omitempty"`
	UsedCode              []CodeableConcept      `json:"usedCode,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
type ProcedurePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
type ProcedureFocalDevice struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `json:"action,omitempty"`
	Manipulated       Reference        `json:"manipulated"`
}

type OtherProcedure Procedure

// on convert struct to json, automatically add resourceType=Procedure
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}

func (resource *Procedure) ProcedureLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Procedure) ProcedureStatus() templ.Component {
	optionsValueSet := VSEvent_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Procedure) ProcedureStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Procedure) ProcedureCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet)
}
func (resource *Procedure) ProcedureCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *Procedure) ProcedureReasonCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *Procedure) ProcedureBodySite(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", &resource.BodySite[0], optionsValueSet)
}
func (resource *Procedure) ProcedureOutcome(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcome", resource.Outcome, optionsValueSet)
}
func (resource *Procedure) ProcedureComplication(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("complication", nil, optionsValueSet)
	}
	return CodeableConceptSelect("complication", &resource.Complication[0], optionsValueSet)
}
func (resource *Procedure) ProcedureFollowUp(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("followUp", nil, optionsValueSet)
	}
	return CodeableConceptSelect("followUp", &resource.FollowUp[0], optionsValueSet)
}
func (resource *Procedure) ProcedureUsedCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("usedCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("usedCode", &resource.UsedCode[0], optionsValueSet)
}
func (resource *Procedure) ProcedurePerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *Procedure) ProcedureFocalDeviceAction(numFocalDevice int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.FocalDevice) >= numFocalDevice {
		return CodeableConceptSelect("action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("action", resource.FocalDevice[numFocalDevice].Action, optionsValueSet)
}
