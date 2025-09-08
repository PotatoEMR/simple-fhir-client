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

// http://hl7.org/fhir/r5/StructureDefinition/ArtifactAssessment
type ArtifactAssessment struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Title             *string                     `json:"title,omitempty"`
	CiteAsReference   *Reference                  `json:"citeAsReference,omitempty"`
	CiteAsMarkdown    *string                     `json:"citeAsMarkdown,omitempty"`
	Date              *time.Time                  `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Copyright         *string                     `json:"copyright,omitempty"`
	ApprovalDate      *time.Time                  `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate    *time.Time                  `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
	ArtifactReference Reference                   `json:"artifactReference"`
	ArtifactCanonical string                      `json:"artifactCanonical"`
	ArtifactUri       string                      `json:"artifactUri"`
	Content           []ArtifactAssessmentContent `json:"content,omitempty"`
	WorkflowStatus    *string                     `json:"workflowStatus,omitempty"`
	Disposition       *string                     `json:"disposition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ArtifactAssessment
type ArtifactAssessmentContent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	InformationType   *string           `json:"informationType,omitempty"`
	Summary           *string           `json:"summary,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
	Quantity          *Quantity         `json:"quantity,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
	Path              []string          `json:"path,omitempty"`
	RelatedArtifact   []RelatedArtifact `json:"relatedArtifact,omitempty"`
	FreeToShare       *bool             `json:"freeToShare,omitempty"`
}

type OtherArtifactAssessment ArtifactAssessment

// on convert struct to json, automatically add resourceType=ArtifactAssessment
func (r ArtifactAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherArtifactAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherArtifactAssessment: OtherArtifactAssessment(r),
		ResourceType:            "ArtifactAssessment",
	})
}
func (r ArtifactAssessment) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ArtifactAssessment/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ArtifactAssessment"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ArtifactAssessment) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ArtifactAssessment) T_CiteAsMarkdown(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CiteAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("CiteAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ArtifactAssessment) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ArtifactCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ArtifactCanonical", nil, htmlAttrs)
	}
	return StringInput("ArtifactCanonical", &resource.ArtifactCanonical, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ArtifactUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ArtifactUri", nil, htmlAttrs)
	}
	return StringInput("ArtifactUri", &resource.ArtifactUri, htmlAttrs)
}
func (resource *ArtifactAssessment) T_WorkflowStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSArtifactassessment_workflow_status

	if resource == nil {
		return CodeSelect("WorkflowStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("WorkflowStatus", resource.WorkflowStatus, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Disposition(htmlAttrs string) templ.Component {
	optionsValueSet := VSArtifactassessment_disposition

	if resource == nil {
		return CodeSelect("Disposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Disposition", resource.Disposition, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentInformationType(numContent int, htmlAttrs string) templ.Component {
	optionsValueSet := VSArtifactassessment_information_type

	if resource == nil || numContent >= len(resource.Content) {
		return CodeSelect("Content["+strconv.Itoa(numContent)+"]InformationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Content["+strconv.Itoa(numContent)+"]InformationType", resource.Content[numContent].InformationType, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentSummary(numContent int, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return StringInput("Content["+strconv.Itoa(numContent)+"]Summary", nil, htmlAttrs)
	}
	return StringInput("Content["+strconv.Itoa(numContent)+"]Summary", resource.Content[numContent].Summary, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentType(numContent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return CodeableConceptSelect("Content["+strconv.Itoa(numContent)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Content["+strconv.Itoa(numContent)+"]Type", resource.Content[numContent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentClassifier(numContent int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numClassifier >= len(resource.Content[numContent].Classifier) {
		return CodeableConceptSelect("Content["+strconv.Itoa(numContent)+"]Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Content["+strconv.Itoa(numContent)+"]Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Content[numContent].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentPath(numContent int, numPath int, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numPath >= len(resource.Content[numContent].Path) {
		return StringInput("Content["+strconv.Itoa(numContent)+"]Path["+strconv.Itoa(numPath)+"]", nil, htmlAttrs)
	}
	return StringInput("Content["+strconv.Itoa(numContent)+"]Path["+strconv.Itoa(numPath)+"]", &resource.Content[numContent].Path[numPath], htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentFreeToShare(numContent int, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return BoolInput("Content["+strconv.Itoa(numContent)+"]FreeToShare", nil, htmlAttrs)
	}
	return BoolInput("Content["+strconv.Itoa(numContent)+"]FreeToShare", resource.Content[numContent].FreeToShare, htmlAttrs)
}
