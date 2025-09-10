package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	CreatedDateTime   *string           `json:"createdDateTime,omitempty"`
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
func (resource *Media) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_Modality(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("modality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("modality", resource.Modality, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_View(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("view", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("view", resource.View, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_CreatedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("createdDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("createdDateTime", resource.CreatedDateTime, htmlAttrs)
}
func (resource *Media) T_Issued(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("issued", nil, htmlAttrs)
	}
	return StringInput("issued", resource.Issued, htmlAttrs)
}
func (resource *Media) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Media) T_BodySite(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Media) T_DeviceName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("deviceName", nil, htmlAttrs)
	}
	return StringInput("deviceName", resource.DeviceName, htmlAttrs)
}
func (resource *Media) T_Height(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("height", nil, htmlAttrs)
	}
	return IntInput("height", resource.Height, htmlAttrs)
}
func (resource *Media) T_Width(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("width", nil, htmlAttrs)
	}
	return IntInput("width", resource.Width, htmlAttrs)
}
func (resource *Media) T_Frames(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("frames", nil, htmlAttrs)
	}
	return IntInput("frames", resource.Frames, htmlAttrs)
}
func (resource *Media) T_Duration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("duration", nil, htmlAttrs)
	}
	return Float64Input("duration", resource.Duration, htmlAttrs)
}
func (resource *Media) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
