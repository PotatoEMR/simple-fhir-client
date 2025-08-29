package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CiteAsReference   *Reference                  `json:"citeAsReference"`
	CiteAsMarkdown    *string                     `json:"citeAsMarkdown"`
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
func (resource *ArtifactAssessment) ArtifactAssessmentLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ArtifactAssessment) ArtifactAssessmentWorkflowStatus() templ.Component {
	optionsValueSet := VSArtifactassessment_workflow_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.WorkflowStatus
	}
	return CodeSelect("workflowStatus", currentVal, optionsValueSet)
}
func (resource *ArtifactAssessment) ArtifactAssessmentDisposition() templ.Component {
	optionsValueSet := VSArtifactassessment_disposition
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Disposition
	}
	return CodeSelect("disposition", currentVal, optionsValueSet)
}
func (resource *ArtifactAssessment) ArtifactAssessmentContentInformationType(numContent int) templ.Component {
	optionsValueSet := VSArtifactassessment_information_type
	currentVal := ""
	if resource != nil && len(resource.Content) >= numContent {
		currentVal = *resource.Content[numContent].InformationType
	}
	return CodeSelect("informationType", currentVal, optionsValueSet)
}
