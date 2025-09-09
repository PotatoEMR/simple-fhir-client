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
	EffectiveDateTime        *time.Time                        `json:"effectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	EffectivePeriod          *Period                           `json:"effectivePeriod,omitempty"`
	Date                     *time.Time                        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// on convert struct to json, automatically add resourceType=ClinicalImpression
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
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
func (resource *ClinicalImpression) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSClinicalimpression_status

	if resource == nil {
		return CodeSelect("ClinicalImpression.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ClinicalImpression.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalImpression.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalImpression.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalImpression.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalImpression.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ClinicalImpression.Description", nil, htmlAttrs)
	}
	return StringInput("ClinicalImpression.Description", resource.Description, htmlAttrs)
}
func (resource *ClinicalImpression) T_EffectiveDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ClinicalImpression.EffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("ClinicalImpression.EffectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *ClinicalImpression) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ClinicalImpression.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ClinicalImpression.Date", resource.Date, htmlAttrs)
}
func (resource *ClinicalImpression) T_Protocol(numProtocol int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocol >= len(resource.Protocol) {
		return StringInput("ClinicalImpression.Protocol["+strconv.Itoa(numProtocol)+"]", nil, htmlAttrs)
	}
	return StringInput("ClinicalImpression.Protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol], htmlAttrs)
}
func (resource *ClinicalImpression) T_Summary(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ClinicalImpression.Summary", nil, htmlAttrs)
	}
	return StringInput("ClinicalImpression.Summary", resource.Summary, htmlAttrs)
}
func (resource *ClinicalImpression) T_PrognosisCodeableConcept(numPrognosisCodeableConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrognosisCodeableConcept >= len(resource.PrognosisCodeableConcept) {
		return CodeableConceptSelect("ClinicalImpression.PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalImpression.PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", &resource.PrognosisCodeableConcept[numPrognosisCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("ClinicalImpression.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("ClinicalImpression.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ClinicalImpression) T_InvestigationCode(numInvestigation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInvestigation >= len(resource.Investigation) {
		return CodeableConceptSelect("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Code", &resource.Investigation[numInvestigation].Code, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingItemCodeableConcept(numFinding int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return CodeableConceptSelect("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].ItemCodeableConcept", resource.Finding[numFinding].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingBasis(numFinding int, htmlAttrs string) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Basis", nil, htmlAttrs)
	}
	return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Basis", resource.Finding[numFinding].Basis, htmlAttrs)
}
