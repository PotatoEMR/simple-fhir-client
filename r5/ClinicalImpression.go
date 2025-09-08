package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	EffectiveDateTime        *time.Time                  `json:"effectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	EffectivePeriod          *Period                     `json:"effectivePeriod,omitempty"`
	Date                     *time.Time                  `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ClinicalImpression) T_EffectiveDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("EffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("EffectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *ClinicalImpression) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ClinicalImpression) T_ChangePattern(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ChangePattern", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ChangePattern", resource.ChangePattern, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Protocol(numProtocol int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocol >= len(resource.Protocol) {
		return StringInput("Protocol["+strconv.Itoa(numProtocol)+"]", nil, htmlAttrs)
	}
	return StringInput("Protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol], htmlAttrs)
}
func (resource *ClinicalImpression) T_Summary(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Summary", nil, htmlAttrs)
	}
	return StringInput("Summary", resource.Summary, htmlAttrs)
}
func (resource *ClinicalImpression) T_PrognosisCodeableConcept(numPrognosisCodeableConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrognosisCodeableConcept >= len(resource.PrognosisCodeableConcept) {
		return CodeableConceptSelect("PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PrognosisCodeableConcept["+strconv.Itoa(numPrognosisCodeableConcept)+"]", &resource.PrognosisCodeableConcept[numPrognosisCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalImpression) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ClinicalImpression) T_FindingBasis(numFinding int, htmlAttrs string) templ.Component {
	if resource == nil || numFinding >= len(resource.Finding) {
		return StringInput("Finding["+strconv.Itoa(numFinding)+"]Basis", nil, htmlAttrs)
	}
	return StringInput("Finding["+strconv.Itoa(numFinding)+"]Basis", resource.Finding[numFinding].Basis, htmlAttrs)
}
