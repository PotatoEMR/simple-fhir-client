package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Media
type Media struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	BasedOn           []Reference       `json:"basedOn,omitempty"`
	PartOf            []Reference       `json:"partOf,omitempty"`
	Status            string            `json:"status"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Modality          *CodeableConcept  `json:"modality,omitempty"`
	View              *CodeableConcept  `json:"view,omitempty"`
	Subject           *Reference        `json:"subject,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	CreatedDateTime   *time.Time        `json:"createdDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	CreatedPeriod     *Period           `json:"createdPeriod,omitempty"`
	Issued            *string           `json:"issued,omitempty"`
	Operator          *Reference        `json:"operator,omitempty"`
	ReasonCode        []CodeableConcept `json:"reasonCode,omitempty"`
	BodySite          *CodeableConcept  `json:"bodySite,omitempty"`
	DeviceName        *string           `json:"deviceName,omitempty"`
	Device            *Reference        `json:"device,omitempty"`
	Height            *int              `json:"height,omitempty"`
	Width             *int              `json:"width,omitempty"`
	Frames            *int              `json:"frames,omitempty"`
	Duration          *float64          `json:"duration,omitempty"`
	Content           Attachment        `json:"content"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherMedia Media

// on convert struct to json, automatically add resourceType=Media
func (r Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedia
		ResourceType string `json:"resourceType"`
	}{
		OtherMedia:   OtherMedia(r),
		ResourceType: "Media",
	})
}
func (r Media) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Media/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Media"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Media) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Media.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Media.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Media.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Media.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Modality(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Media.Modality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Media.Modality", resource.Modality, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_View(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Media.View", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Media.View", resource.View, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_CreatedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Media.CreatedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Media.CreatedDateTime", resource.CreatedDateTime, htmlAttrs)
}
func (resource *Media) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Media.Issued", nil, htmlAttrs)
	}
	return StringInput("Media.Issued", resource.Issued, htmlAttrs)
}
func (resource *Media) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Media.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Media.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Media) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Media.BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Media.BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_DeviceName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Media.DeviceName", nil, htmlAttrs)
	}
	return StringInput("Media.DeviceName", resource.DeviceName, htmlAttrs)
}
func (resource *Media) T_Height(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Media.Height", nil, htmlAttrs)
	}
	return IntInput("Media.Height", resource.Height, htmlAttrs)
}
func (resource *Media) T_Width(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Media.Width", nil, htmlAttrs)
	}
	return IntInput("Media.Width", resource.Width, htmlAttrs)
}
func (resource *Media) T_Frames(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Media.Frames", nil, htmlAttrs)
	}
	return IntInput("Media.Frames", resource.Frames, htmlAttrs)
}
func (resource *Media) T_Duration(htmlAttrs string) templ.Component {
	if resource == nil {
		return Float64Input("Media.Duration", nil, htmlAttrs)
	}
	return Float64Input("Media.Duration", resource.Duration, htmlAttrs)
}
func (resource *Media) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Media.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Media.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
