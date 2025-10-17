package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
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
	OnsetDateTime      *FhirDateTime       `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period             `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range              `json:"onsetRange,omitempty"`
	OnsetString        *string             `json:"onsetString,omitempty"`
	AbatementDateTime  *FhirDateTime       `json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                `json:"abatementAge,omitempty"`
	AbatementPeriod    *Period             `json:"abatementPeriod,omitempty"`
	AbatementRange     *Range              `json:"abatementRange,omitempty"`
	AbatementString    *string             `json:"abatementString,omitempty"`
	RecordedDate       *FhirDateTime       `json:"recordedDate,omitempty"`
	Recorder           *Reference          `json:"recorder,omitempty"`
	Asserter           *Reference          `json:"asserter,omitempty"`
	Stage              []ConditionStage    `json:"stage,omitempty"`
	Evidence           []ConditionEvidence `json:"evidence,omitempty"`
	Note               []Annotation        `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
type ConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
type ConditionEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

type OtherCondition Condition

// struct -> json, automatically add resourceType=Patient
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

// json -> struct, first reject if resourceType != Condition
func (r *Condition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Condition" {
		return errors.New("resourceType not Condition")
	}
	return json.Unmarshal(data, (*OtherCondition)(r))
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
func (resource *Condition) T_ClinicalStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCondition_clinical

	if resource == nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("clinicalStatus", resource.ClinicalStatus, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_VerificationStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCondition_ver_status

	if resource == nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Severity(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *Condition) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Condition) T_OnsetDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("onsetDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("onsetDateTime", resource.OnsetDateTime, htmlAttrs)
}
func (resource *Condition) T_OnsetAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("onsetAge", nil, htmlAttrs)
	}
	return AgeInput("onsetAge", resource.OnsetAge, htmlAttrs)
}
func (resource *Condition) T_OnsetPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("onsetPeriod", nil, htmlAttrs)
	}
	return PeriodInput("onsetPeriod", resource.OnsetPeriod, htmlAttrs)
}
func (resource *Condition) T_OnsetRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("onsetRange", nil, htmlAttrs)
	}
	return RangeInput("onsetRange", resource.OnsetRange, htmlAttrs)
}
func (resource *Condition) T_OnsetString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("onsetString", nil, htmlAttrs)
	}
	return StringInput("onsetString", resource.OnsetString, htmlAttrs)
}
func (resource *Condition) T_AbatementDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("abatementDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("abatementDateTime", resource.AbatementDateTime, htmlAttrs)
}
func (resource *Condition) T_AbatementAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("abatementAge", nil, htmlAttrs)
	}
	return AgeInput("abatementAge", resource.AbatementAge, htmlAttrs)
}
func (resource *Condition) T_AbatementPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("abatementPeriod", nil, htmlAttrs)
	}
	return PeriodInput("abatementPeriod", resource.AbatementPeriod, htmlAttrs)
}
func (resource *Condition) T_AbatementRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("abatementRange", nil, htmlAttrs)
	}
	return RangeInput("abatementRange", resource.AbatementRange, htmlAttrs)
}
func (resource *Condition) T_AbatementString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("abatementString", nil, htmlAttrs)
	}
	return StringInput("abatementString", resource.AbatementString, htmlAttrs)
}
func (resource *Condition) T_RecordedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recordedDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *Condition) T_Recorder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recorder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recorder", resource.Recorder, htmlAttrs)
}
func (resource *Condition) T_Asserter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "asserter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "asserter", resource.Asserter, htmlAttrs)
}
func (resource *Condition) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Condition) T_StageSummary(numStage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStage >= len(resource.Stage) {
		return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].summary", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].summary", resource.Stage[numStage].Summary, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_StageAssessment(frs []FhirResource, numStage int, numAssessment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStage >= len(resource.Stage) || numAssessment >= len(resource.Stage[numStage].Assessment) {
		return ReferenceInput(frs, "stage["+strconv.Itoa(numStage)+"].assessment["+strconv.Itoa(numAssessment)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "stage["+strconv.Itoa(numStage)+"].assessment["+strconv.Itoa(numAssessment)+"]", &resource.Stage[numStage].Assessment[numAssessment], htmlAttrs)
}
func (resource *Condition) T_StageType(numStage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStage >= len(resource.Stage) {
		return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("stage["+strconv.Itoa(numStage)+"].type", resource.Stage[numStage].Type, optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvidence >= len(resource.Evidence) || numCode >= len(resource.Evidence[numEvidence].Code) {
		return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Condition) T_EvidenceDetail(frs []FhirResource, numEvidence int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvidence >= len(resource.Evidence) || numDetail >= len(resource.Evidence[numEvidence].Detail) {
		return ReferenceInput(frs, "evidence["+strconv.Itoa(numEvidence)+"].detail["+strconv.Itoa(numDetail)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "evidence["+strconv.Itoa(numEvidence)+"].detail["+strconv.Itoa(numDetail)+"]", &resource.Evidence[numEvidence].Detail[numDetail], htmlAttrs)
}
