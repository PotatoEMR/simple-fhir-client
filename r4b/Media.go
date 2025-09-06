package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

func (resource *Media) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Media.Id", nil)
	}
	return StringInput("Media.Id", resource.Id)
}
func (resource *Media) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Media.ImplicitRules", nil)
	}
	return StringInput("Media.ImplicitRules", resource.ImplicitRules)
}
func (resource *Media) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Media.Language", nil, optionsValueSet)
	}
	return CodeSelect("Media.Language", resource.Language, optionsValueSet)
}
func (resource *Media) T_Status() templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Media.Status", nil, optionsValueSet)
	}
	return CodeSelect("Media.Status", &resource.Status, optionsValueSet)
}
func (resource *Media) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Media.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Media.Type", resource.Type, optionsValueSet)
}
func (resource *Media) T_Modality(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Media.Modality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Media.Modality", resource.Modality, optionsValueSet)
}
func (resource *Media) T_View(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Media.View", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Media.View", resource.View, optionsValueSet)
}
func (resource *Media) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("Media.Issued", nil)
	}
	return StringInput("Media.Issued", resource.Issued)
}
func (resource *Media) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Media.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Media.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Media) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Media.BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Media.BodySite", resource.BodySite, optionsValueSet)
}
func (resource *Media) T_DeviceName() templ.Component {

	if resource == nil {
		return StringInput("Media.DeviceName", nil)
	}
	return StringInput("Media.DeviceName", resource.DeviceName)
}
func (resource *Media) T_Height() templ.Component {

	if resource == nil {
		return IntInput("Media.Height", nil)
	}
	return IntInput("Media.Height", resource.Height)
}
func (resource *Media) T_Width() templ.Component {

	if resource == nil {
		return IntInput("Media.Width", nil)
	}
	return IntInput("Media.Width", resource.Width)
}
func (resource *Media) T_Frames() templ.Component {

	if resource == nil {
		return IntInput("Media.Frames", nil)
	}
	return IntInput("Media.Frames", resource.Frames)
}
func (resource *Media) T_Duration() templ.Component {

	if resource == nil {
		return Float64Input("Media.Duration", nil)
	}
	return Float64Input("Media.Duration", resource.Duration)
}
