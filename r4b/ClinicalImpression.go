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

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                           `json:"id,omitempty"`
	Meta                     *Meta                             `json:"meta,omitempty"`
	ImplicitRules            *string                           `json:"implicitRules,omitempty"`
	Language                 *string                           `json:"language,omitempty"`
	Text                     *Narrative                        `json:"text,omitempty"`
	Contained                []Resource                        `json:"contained,omitempty"`
	Extension                []Extension                       `json:"extension,omitempty"`
	ModifierExtension        []Extension                       `json:"modifierExtension,omitempty"`
	Identifier               []Identifier                      `json:"identifier,omitempty"`
	Status                   string                            `json:"status"`
	StatusReason             *CodeableConcept                  `json:"statusReason,omitempty"`
	Code                     *CodeableConcept                  `json:"code,omitempty"`
	Description              *string                           `json:"description,omitempty"`
	Subject                  Reference                         `json:"subject"`
	Encounter                *Reference                        `json:"encounter,omitempty"`
	EffectiveDateTime        *FhirDateTime                     `json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                           `json:"effectivePeriod,omitempty"`
	Date                     *FhirDateTime                     `json:"date,omitempty"`
	Assessor                 *Reference                        `json:"assessor,omitempty"`
	Previous                 *Reference                        `json:"previous,omitempty"`
	Problem                  []Reference                       `json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	Protocol                 []string                          `json:"protocol,omitempty"`
	Summary                  *string                           `json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `json:"prognosisReference,omitempty"`
	SupportingInfo           []Reference                       `json:"supportingInfo,omitempty"`
	Note                     []Annotation                      `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalImpression
type ClinicalImpressionInvestigation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Item              []Reference     `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalImpression
type ClinicalImpressionFinding struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	Basis               *string          `json:"basis,omitempty"`
}

type OtherClinicalImpression ClinicalImpression

// struct -> json, automatically add resourceType=Patient
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
}

// json -> struct, first reject if resourceType != ClinicalImpression
func (r *ClinicalImpression) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "ClinicalImpression" {
		return errors.New("resourceType not ClinicalImpression")
	}
	return json.Unmarshal(data, (*OtherClinicalImpression)(r))
}

func (r ClinicalImpression) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ClinicalImpression/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ClinicalImpression"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ClinicalImpression) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClinicalimpression_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ClinicalImpression) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *ClinicalImpression) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *ClinicalImpression) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *ClinicalImpression) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *ClinicalImpression) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ClinicalImpression) T_Assessor(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "assessor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "assessor", resource.Assessor, htmlAttrs)
}
func (resource *ClinicalImpression) T_Previous(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "previous", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "previous", resource.Previous, htmlAttrs)
}
func (resource *ClinicalImpression) T_Problem(frs []FhirResource, numProblem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProblem >= len(resource.Problem) {
		return ReferenceInput(frs, "problem["+strconv.Itoa(numProblem)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "problem["+strconv.Itoa(numProblem)+"]", &resource.Problem[numProblem], htmlAttrs)
}
func (resource *ClinicalImpression) T_Protocol(numProtocol int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocol >= len(resource.Protocol) {
		return StringInput("protocol["+strconv.Itoa(numProtocol)+"]", nil, htmlAttrs)
	}
	return StringInput("protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol], htmlAttrs)
}
func (resource *ClinicalImpression) T_Summary(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("summary", nil, htmlAttrs)
	}
	return StringInput("summary", resource.Summary, htmlAttrs)
}
func (resource *ClinicalImpression) T_PrognosisCodeableConcept(numPrognosisCodeableConcept int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrognosisCodeableConcept >= len(resource.PrognosisCodeableConcept) {
		return CodeableConceptSelect("prognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("prognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", &resource.PrognosisCodeableConcept[numPrognosisCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_PrognosisReference(frs []FhirResource, numPrognosisReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrognosisReference >= len(resource.PrognosisReference) {
		return ReferenceInput(frs, "prognosisReference["+strconv.Itoa(numPrognosisReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "prognosisReference["+strconv.Itoa(numPrognosisReference)+"]", &resource.PrognosisReference[numPrognosisReference], htmlAttrs)
}
func (resource *ClinicalImpression) T_SupportingInfo(frs []FhirResource, numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *ClinicalImpression) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ClinicalImpression) T_InvestigationCode(numInvestigation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInvestigation >= len(resource.Investigation) {
		return CodeableConceptSelect("investigation["+strconv.Itoa(numInvestigation)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("investigation["+strconv.Itoa(numInvestigation)+"].code", &resource.Investigation[numInvestigation].Code, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_InvestigationItem(frs []FhirResource, numInvestigation int, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInvestigation >= len(resource.Investigation) || numItem >= len(resource.Investigation[numInvestigation].Item) {
		return ReferenceInput(frs, "investigation["+strconv.Itoa(numInvestigation)+"].item["+strconv.Itoa(numItem)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "investigation["+strconv.Itoa(numInvestigation)+"].item["+strconv.Itoa(numItem)+"]", &resource.Investigation[numInvestigation].Item[numItem], htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingItemCodeableConcept(numFinding int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return CodeableConceptSelect("finding["+strconv.Itoa(numFinding)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("finding["+strconv.Itoa(numFinding)+"].itemCodeableConcept", resource.Finding[numFinding].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingItemReference(frs []FhirResource, numFinding int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return ReferenceInput(frs, "finding["+strconv.Itoa(numFinding)+"].itemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "finding["+strconv.Itoa(numFinding)+"].itemReference", resource.Finding[numFinding].ItemReference, htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingBasis(numFinding int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return StringInput("finding["+strconv.Itoa(numFinding)+"].basis", nil, htmlAttrs)
	}
	return StringInput("finding["+strconv.Itoa(numFinding)+"].basis", resource.Finding[numFinding].Basis, htmlAttrs)
}
