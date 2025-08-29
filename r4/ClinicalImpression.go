package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/ClinicalImpression
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
	EffectiveDateTime        *string                           `json:"effectiveDateTime"`
	EffectivePeriod          *Period                           `json:"effectivePeriod"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ClinicalImpression
type ClinicalImpressionInvestigation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Item              []Reference     `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClinicalImpression
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

func (resource *ClinicalImpression) ClinicalImpressionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ClinicalImpression) ClinicalImpressionStatus() templ.Component {
	optionsValueSet := VSClinicalimpression_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
