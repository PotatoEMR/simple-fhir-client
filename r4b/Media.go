package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CreatedDateTime   *string           `json:"createdDateTime"`
	CreatedPeriod     *Period           `json:"createdPeriod"`
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

func (resource *Media) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Media) T_Status() templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Media) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *Media) T_Modality(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("modality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modality", resource.Modality, optionsValueSet)
}
func (resource *Media) T_View(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("view", nil, optionsValueSet)
	}
	return CodeableConceptSelect("view", resource.View, optionsValueSet)
}
func (resource *Media) T_ReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *Media) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet)
}
