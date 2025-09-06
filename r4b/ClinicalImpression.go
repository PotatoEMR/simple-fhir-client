package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
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
	EffectiveDateTime        *string                           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                           `json:"effectivePeriod,omitempty"`
	Date                     *string                           `json:"date,omitempty"`
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

func (resource *ClinicalImpression) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ClinicalImpression.Id", nil)
	}
	return StringInput("ClinicalImpression.Id", resource.Id)
}
func (resource *ClinicalImpression) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ClinicalImpression.ImplicitRules", nil)
	}
	return StringInput("ClinicalImpression.ImplicitRules", resource.ImplicitRules)
}
func (resource *ClinicalImpression) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ClinicalImpression.Language", nil, optionsValueSet)
	}
	return CodeSelect("ClinicalImpression.Language", resource.Language, optionsValueSet)
}
func (resource *ClinicalImpression) T_Status() templ.Component {
	optionsValueSet := VSClinicalimpression_status

	if resource == nil {
		return CodeSelect("ClinicalImpression.Status", nil, optionsValueSet)
	}
	return CodeSelect("ClinicalImpression.Status", &resource.Status, optionsValueSet)
}
func (resource *ClinicalImpression) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalImpression.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalImpression.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *ClinicalImpression) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalImpression.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalImpression.Code", resource.Code, optionsValueSet)
}
func (resource *ClinicalImpression) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ClinicalImpression.Description", nil)
	}
	return StringInput("ClinicalImpression.Description", resource.Description)
}
func (resource *ClinicalImpression) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ClinicalImpression.Date", nil)
	}
	return StringInput("ClinicalImpression.Date", resource.Date)
}
func (resource *ClinicalImpression) T_Protocol(numProtocol int) templ.Component {

	if resource == nil || len(resource.Protocol) >= numProtocol {
		return StringInput("ClinicalImpression.Protocol["+strconv.Itoa(numProtocol)+"]", nil)
	}
	return StringInput("ClinicalImpression.Protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol])
}
func (resource *ClinicalImpression) T_Summary() templ.Component {

	if resource == nil {
		return StringInput("ClinicalImpression.Summary", nil)
	}
	return StringInput("ClinicalImpression.Summary", resource.Summary)
}
func (resource *ClinicalImpression) T_PrognosisCodeableConcept(numPrognosisCodeableConcept int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrognosisCodeableConcept) >= numPrognosisCodeableConcept {
		return CodeableConceptSelect("ClinicalImpression.PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalImpression.PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", &resource.PrognosisCodeableConcept[numPrognosisCodeableConcept], optionsValueSet)
}
func (resource *ClinicalImpression) T_InvestigationId(numInvestigation int) templ.Component {

	if resource == nil || len(resource.Investigation) >= numInvestigation {
		return StringInput("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Id", nil)
	}
	return StringInput("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Id", resource.Investigation[numInvestigation].Id)
}
func (resource *ClinicalImpression) T_InvestigationCode(numInvestigation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Investigation) >= numInvestigation {
		return CodeableConceptSelect("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalImpression.Investigation["+strconv.Itoa(numInvestigation)+"].Code", &resource.Investigation[numInvestigation].Code, optionsValueSet)
}
func (resource *ClinicalImpression) T_FindingId(numFinding int) templ.Component {

	if resource == nil || len(resource.Finding) >= numFinding {
		return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Id", nil)
	}
	return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Id", resource.Finding[numFinding].Id)
}
func (resource *ClinicalImpression) T_FindingItemCodeableConcept(numFinding int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Finding) >= numFinding {
		return CodeableConceptSelect("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].ItemCodeableConcept", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].ItemCodeableConcept", resource.Finding[numFinding].ItemCodeableConcept, optionsValueSet)
}
func (resource *ClinicalImpression) T_FindingBasis(numFinding int) templ.Component {

	if resource == nil || len(resource.Finding) >= numFinding {
		return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Basis", nil)
	}
	return StringInput("ClinicalImpression.Finding["+strconv.Itoa(numFinding)+"].Basis", resource.Finding[numFinding].Basis)
}
