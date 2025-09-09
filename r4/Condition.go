package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Condition
type Condition struct {
	Id                 *string             `json:"id,omitempty"`
	Meta               *Meta               `json:"meta,omitempty"`
	ImplicitRules      *string             `json:"implicitRules,omitempty"`
	Language           *string             `json:"language,omitempty"`
	Text               *Narrative          `json:"text,omitempty"`
	Contained          []Resource          `json:"contained,omitempty"`
	Extension          []Extension         `json:"extension,omitempty"`
	ModifierExtension  []Extension         `json:"modifierExtension,omitempty"`
	Identifier         []Identifier        `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept    `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept    `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept   `json:"category,omitempty"`
	Severity           *CodeableConcept    `json:"severity,omitempty"`
	Code               *CodeableConcept    `json:"code,omitempty"`
	BodySite           []CodeableConcept   `json:"bodySite,omitempty"`
	Subject            Reference           `json:"subject"`
	Encounter          *Reference          `json:"encounter,omitempty"`
	OnsetDateTime      *string             `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period             `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range              `json:"onsetRange,omitempty"`
	OnsetString        *string             `json:"onsetString,omitempty"`
	AbatementDateTime  *string             `json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                `json:"abatementAge,omitempty"`
	AbatementPeriod    *Period             `json:"abatementPeriod,omitempty"`
	AbatementRange     *Range              `json:"abatementRange,omitempty"`
	AbatementString    *string             `json:"abatementString,omitempty"`
	RecordedDate       *string             `json:"recordedDate,omitempty"`
	Recorder           *Reference          `json:"recorder,omitempty"`
	Asserter           *Reference          `json:"asserter,omitempty"`
	Stage              []ConditionStage    `json:"stage,omitempty"`
	Evidence           []ConditionEvidence `json:"evidence,omitempty"`
	Note               []Annotation        `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Condition
type ConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Condition
type ConditionEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

type OtherCondition Condition

// on convert struct to json, automatically add resourceType=Condition
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}
func (r Condition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Condition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Condition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Condition) T_ClinicalStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSCondition_clinical

	if resource == nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("clinicalStatus", resource.ClinicalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_VerificationStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSCondition_ver_status

	if resource == nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Severity(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_OnsetDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("onsetDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("onsetDateTime", resource.OnsetDateTime, htmlAttrs)
}
func (resource *Condition) T_OnsetString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("onsetString", nil, htmlAttrs)
	}
	return StringInput("onsetString", resource.OnsetString, htmlAttrs)
}
func (resource *Condition) T_AbatementDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("abatementDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("abatementDateTime", resource.AbatementDateTime, htmlAttrs)
}
func (resource *Condition) T_AbatementString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("abatementString", nil, htmlAttrs)
	}
	return StringInput("abatementString", resource.AbatementString, htmlAttrs)
}
func (resource *Condition) T_RecordedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("recordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("recordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *Condition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Condition) T_StageSummary(numStage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStage >= len(resource.Stage) {
		return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].summary", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].summary", resource.Stage[numStage].Summary, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_StageType(numStage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStage >= len(resource.Stage) {
		return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].type", resource.Stage[numStage].Type, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvidence >= len(resource.Evidence) || numCode >= len(resource.Evidence[numEvidence].Code) {
		return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet, htmlAttrs)
}
