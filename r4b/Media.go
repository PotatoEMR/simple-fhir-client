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

// http://hl7.org/fhir/r4b/StructureDefinition/Media
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
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Modality(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Modality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Modality", resource.Modality, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_View(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("View", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("View", resource.View, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_CreatedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CreatedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("CreatedDateTime", resource.CreatedDateTime, htmlAttrs)
}
func (resource *Media) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Issued", nil, htmlAttrs)
	}
	return StringInput("Issued", resource.Issued, htmlAttrs)
}
func (resource *Media) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Media) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_DeviceName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceName", nil, htmlAttrs)
	}
	return StringInput("DeviceName", resource.DeviceName, htmlAttrs)
}
func (resource *Media) T_Height(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Height", nil, htmlAttrs)
	}
	return IntInput("Height", resource.Height, htmlAttrs)
}
func (resource *Media) T_Width(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Width", nil, htmlAttrs)
	}
	return IntInput("Width", resource.Width, htmlAttrs)
}
func (resource *Media) T_Frames(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Frames", nil, htmlAttrs)
	}
	return IntInput("Frames", resource.Frames, htmlAttrs)
}
func (resource *Media) T_Duration(htmlAttrs string) templ.Component {
	if resource == nil {
		return Float64Input("Duration", nil, htmlAttrs)
	}
	return Float64Input("Duration", resource.Duration, htmlAttrs)
}
func (resource *Media) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
