package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                     `json:"id,omitempty"`
	Meta                     *Meta                       `json:"meta,omitempty"`
	ImplicitRules            *string                     `json:"implicitRules,omitempty"`
	Language                 *string                     `json:"language,omitempty"`
	Text                     *Narrative                  `json:"text,omitempty"`
	Contained                []Resource                  `json:"contained,omitempty"`
	Extension                []Extension                 `json:"extension,omitempty"`
	ModifierExtension        []Extension                 `json:"modifierExtension,omitempty"`
	Identifier               []Identifier                `json:"identifier,omitempty"`
	Status                   string                      `json:"status"`
	StatusReason             *CodeableConcept            `json:"statusReason,omitempty"`
	Description              *string                     `json:"description,omitempty"`
	Subject                  Reference                   `json:"subject"`
	Encounter                *Reference                  `json:"encounter,omitempty"`
	EffectiveDateTime        *string                     `json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                     `json:"effectivePeriod,omitempty"`
	Date                     *string                     `json:"date,omitempty"`
	Performer                *Reference                  `json:"performer,omitempty"`
	Previous                 *Reference                  `json:"previous,omitempty"`
	Problem                  []Reference                 `json:"problem,omitempty"`
	ChangePattern            *CodeableConcept            `json:"changePattern,omitempty"`
	Protocol                 []string                    `json:"protocol,omitempty"`
	Summary                  *string                     `json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding `json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept           `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                 `json:"prognosisReference,omitempty"`
	SupportingInfo           []Reference                 `json:"supportingInfo,omitempty"`
	Note                     []Annotation                `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalImpression
type ClinicalImpressionFinding struct {
	Id                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Item              *CodeableReference `json:"item,omitempty"`
	Basis             *string            `json:"basis,omitempty"`
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
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ClinicalImpression) T_EffectiveDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *ClinicalImpression) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ClinicalImpression) T_ChangePattern(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("changePattern", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("changePattern", resource.ChangePattern, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Protocol(numProtocol int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocol >= len(resource.Protocol) {
		return StringInput("protocol["+strconv.Itoa(numProtocol)+"]", nil, htmlAttrs)
	}
	return StringInput("protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol], htmlAttrs)
}
func (resource *ClinicalImpression) T_Summary(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("summary", nil, htmlAttrs)
	}
	return StringInput("summary", resource.Summary, htmlAttrs)
}
func (resource *ClinicalImpression) T_PrognosisCodeableConcept(numPrognosisCodeableConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrognosisCodeableConcept >= len(resource.PrognosisCodeableConcept) {
		return CodeableConceptSelect("prognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("prognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", &resource.PrognosisCodeableConcept[numPrognosisCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingBasis(numFinding int, htmlAttrs string) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return StringInput("finding["+strconv.Itoa(numFinding)+"].basis", nil, htmlAttrs)
	}
	return StringInput("finding["+strconv.Itoa(numFinding)+"].basis", resource.Finding[numFinding].Basis, htmlAttrs)
}
