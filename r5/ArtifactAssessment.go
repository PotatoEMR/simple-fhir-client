package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date              *FhirDateTime               `json:"date,omitempty"`
	Copyright         *string                     `json:"copyright,omitempty"`
	ApprovalDate      *FhirDate                   `json:"approvalDate,omitempty"`
	LastReviewDate    *FhirDate                   `json:"lastReviewDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r ArtifactAssessment) ResourceType() string {
	return "ArtifactAssessment"
}

func (resource *ArtifactAssessment) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ArtifactAssessment) T_CiteAsReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "citeAsReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "citeAsReference", resource.CiteAsReference, htmlAttrs)
}
func (resource *ArtifactAssessment) T_CiteAsMarkdown(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("citeAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("citeAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ArtifactAssessment) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ArtifactReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "artifactReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "artifactReference", &resource.ArtifactReference, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ArtifactCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("artifactCanonical", nil, htmlAttrs)
	}
	return StringInput("artifactCanonical", &resource.ArtifactCanonical, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ArtifactUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("artifactUri", nil, htmlAttrs)
	}
	return StringInput("artifactUri", &resource.ArtifactUri, htmlAttrs)
}
func (resource *ArtifactAssessment) T_WorkflowStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSArtifactassessment_workflow_status

	if resource == nil {
		return CodeSelect("workflowStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("workflowStatus", resource.WorkflowStatus, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSArtifactassessment_disposition

	if resource == nil {
		return CodeSelect("disposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("disposition", resource.Disposition, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentInformationType(numContent int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSArtifactassessment_information_type

	if resource == nil || numContent >= len(resource.Content) {
		return CodeSelect("content["+strconv.Itoa(numContent)+"].informationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("content["+strconv.Itoa(numContent)+"].informationType", resource.Content[numContent].InformationType, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentSummary(numContent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return StringInput("content["+strconv.Itoa(numContent)+"].summary", nil, htmlAttrs)
	}
	return StringInput("content["+strconv.Itoa(numContent)+"].summary", resource.Content[numContent].Summary, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentType(numContent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return CodeableConceptSelect("content["+strconv.Itoa(numContent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("content["+strconv.Itoa(numContent)+"].type", resource.Content[numContent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentClassifier(numContent int, numClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numClassifier >= len(resource.Content[numContent].Classifier) {
		return CodeableConceptSelect("content["+strconv.Itoa(numContent)+"].classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("content["+strconv.Itoa(numContent)+"].classifier["+strconv.Itoa(numClassifier)+"]", &resource.Content[numContent].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentQuantity(numContent int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return QuantityInput("content["+strconv.Itoa(numContent)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("content["+strconv.Itoa(numContent)+"].quantity", resource.Content[numContent].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentAuthor(frs []FhirResource, numContent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return ReferenceInput(frs, "content["+strconv.Itoa(numContent)+"].author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "content["+strconv.Itoa(numContent)+"].author", resource.Content[numContent].Author, htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentPath(numContent int, numPath int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numPath >= len(resource.Content[numContent].Path) {
		return StringInput("content["+strconv.Itoa(numContent)+"].path["+strconv.Itoa(numPath)+"]", nil, htmlAttrs)
	}
	return StringInput("content["+strconv.Itoa(numContent)+"].path["+strconv.Itoa(numPath)+"]", &resource.Content[numContent].Path[numPath], htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentRelatedArtifact(numContent int, numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numRelatedArtifact >= len(resource.Content[numContent].RelatedArtifact) {
		return RelatedArtifactInput("content["+strconv.Itoa(numContent)+"].relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("content["+strconv.Itoa(numContent)+"].relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.Content[numContent].RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *ArtifactAssessment) T_ContentFreeToShare(numContent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return BoolInput("content["+strconv.Itoa(numContent)+"].freeToShare", nil, htmlAttrs)
	}
	return BoolInput("content["+strconv.Itoa(numContent)+"].freeToShare", resource.Content[numContent].FreeToShare, htmlAttrs)
}
