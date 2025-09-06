package r5

//generated with command go run ./bultaoreune -nodownload
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
	Date              *string                     `json:"date,omitempty"`
	Copyright         *string                     `json:"copyright,omitempty"`
	ApprovalDate      *string                     `json:"approvalDate,omitempty"`
	LastReviewDate    *string                     `json:"lastReviewDate,omitempty"`
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

func (resource *ArtifactAssessment) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.Id", nil)
	}
	return StringInput("ArtifactAssessment.Id", resource.Id)
}
func (resource *ArtifactAssessment) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.ImplicitRules", nil)
	}
	return StringInput("ArtifactAssessment.ImplicitRules", resource.ImplicitRules)
}
func (resource *ArtifactAssessment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ArtifactAssessment.Language", nil, optionsValueSet)
	}
	return CodeSelect("ArtifactAssessment.Language", resource.Language, optionsValueSet)
}
func (resource *ArtifactAssessment) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.Title", nil)
	}
	return StringInput("ArtifactAssessment.Title", resource.Title)
}
func (resource *ArtifactAssessment) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.Date", nil)
	}
	return StringInput("ArtifactAssessment.Date", resource.Date)
}
func (resource *ArtifactAssessment) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.Copyright", nil)
	}
	return StringInput("ArtifactAssessment.Copyright", resource.Copyright)
}
func (resource *ArtifactAssessment) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.ApprovalDate", nil)
	}
	return StringInput("ArtifactAssessment.ApprovalDate", resource.ApprovalDate)
}
func (resource *ArtifactAssessment) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ArtifactAssessment.LastReviewDate", nil)
	}
	return StringInput("ArtifactAssessment.LastReviewDate", resource.LastReviewDate)
}
func (resource *ArtifactAssessment) T_WorkflowStatus() templ.Component {
	optionsValueSet := VSArtifactassessment_workflow_status

	if resource == nil {
		return CodeSelect("ArtifactAssessment.WorkflowStatus", nil, optionsValueSet)
	}
	return CodeSelect("ArtifactAssessment.WorkflowStatus", resource.WorkflowStatus, optionsValueSet)
}
func (resource *ArtifactAssessment) T_Disposition() templ.Component {
	optionsValueSet := VSArtifactassessment_disposition

	if resource == nil {
		return CodeSelect("ArtifactAssessment.Disposition", nil, optionsValueSet)
	}
	return CodeSelect("ArtifactAssessment.Disposition", resource.Disposition, optionsValueSet)
}
func (resource *ArtifactAssessment) T_ContentId(numContent int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Id", nil)
	}
	return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Id", resource.Content[numContent].Id)
}
func (resource *ArtifactAssessment) T_ContentInformationType(numContent int) templ.Component {
	optionsValueSet := VSArtifactassessment_information_type

	if resource == nil || len(resource.Content) >= numContent {
		return CodeSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].InformationType", nil, optionsValueSet)
	}
	return CodeSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].InformationType", resource.Content[numContent].InformationType, optionsValueSet)
}
func (resource *ArtifactAssessment) T_ContentSummary(numContent int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Summary", nil)
	}
	return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Summary", resource.Content[numContent].Summary)
}
func (resource *ArtifactAssessment) T_ContentType(numContent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return CodeableConceptSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Type", resource.Content[numContent].Type, optionsValueSet)
}
func (resource *ArtifactAssessment) T_ContentClassifier(numContent int, numClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Content) >= numContent || len(resource.Content[numContent].Classifier) >= numClassifier {
		return CodeableConceptSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Content[numContent].Classifier[numClassifier], optionsValueSet)
}
func (resource *ArtifactAssessment) T_ContentPath(numContent int, numPath int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent || len(resource.Content[numContent].Path) >= numPath {
		return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Path["+strconv.Itoa(numPath)+"]", nil)
	}
	return StringInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].Path["+strconv.Itoa(numPath)+"]", &resource.Content[numContent].Path[numPath])
}
func (resource *ArtifactAssessment) T_ContentFreeToShare(numContent int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return BoolInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].FreeToShare", nil)
	}
	return BoolInput("ArtifactAssessment.Content["+strconv.Itoa(numContent)+"].FreeToShare", resource.Content[numContent].FreeToShare)
}
