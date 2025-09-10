package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []Resource                   `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Version           *string                      `json:"version,omitempty"`
	BasedOn           []Reference                  `json:"basedOn,omitempty"`
	Status            string                       `json:"status"`
	DocStatus         *string                      `json:"docStatus,omitempty"`
	Modality          []CodeableConcept            `json:"modality,omitempty"`
	Type              *CodeableConcept             `json:"type,omitempty"`
	Category          []CodeableConcept            `json:"category,omitempty"`
	Subject           *Reference                   `json:"subject,omitempty"`
	Context           []Reference                  `json:"context,omitempty"`
	Event             []CodeableReference          `json:"event,omitempty"`
	BodySite          []CodeableReference          `json:"bodySite,omitempty"`
	FacilityType      *CodeableConcept             `json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept             `json:"practiceSetting,omitempty"`
	Period            *Period                      `json:"period,omitempty"`
	Date              *string                      `json:"date,omitempty"`
	Author            []Reference                  `json:"author,omitempty"`
	Attester          []DocumentReferenceAttester  `json:"attester,omitempty"`
	Custodian         *Reference                   `json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `json:"content"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DocumentReference
type DocumentReferenceAttester struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Mode              CodeableConcept `json:"mode"`
	Time              *string         `json:"time,omitempty"`
	Party             *Reference      `json:"party,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DocumentReference
type DocumentReferenceRelatesTo struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Target            Reference       `json:"target"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DocumentReference
type DocumentReferenceContent struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Attachment        Attachment                        `json:"attachment"`
	Profile           []DocumentReferenceContentProfile `json:"profile,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DocumentReference
type DocumentReferenceContentProfile struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueUri          string      `json:"valueUri"`
	ValueCanonical    string      `json:"valueCanonical"`
}

type OtherDocumentReference DocumentReference

// on convert struct to json, automatically add resourceType=DocumentReference
func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}
func (r DocumentReference) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DocumentReference/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DocumentReference"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DocumentReference) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *DocumentReference) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_DocStatus(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("docStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("docStatus", resource.DocStatus, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Modality(numModality int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numModality >= len(resource.Modality) {
		return CodeableConceptSelect("modality["+strconv.Itoa(numModality)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("modality["+strconv.Itoa(numModality)+"]", &resource.Modality[numModality], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_FacilityType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("facilityType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("facilityType", resource.FacilityType, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_PracticeSetting(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("practiceSetting", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("practiceSetting", resource.PracticeSetting, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("date", nil, htmlAttrs)
	}
	return StringInput("date", resource.Date, htmlAttrs)
}
func (resource *DocumentReference) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *DocumentReference) T_SecurityLabel(numSecurityLabel int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSecurityLabel >= len(resource.SecurityLabel) {
		return CodeableConceptSelect("securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_AttesterMode(numAttester int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeableConceptSelect("attester["+strconv.Itoa(numAttester)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("attester["+strconv.Itoa(numAttester)+"].mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_AttesterTime(numAttester int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return DateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", nil, htmlAttrs)
	}
	return DateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeableConceptSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueCoding(numContent int, numProfile int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return CodingSelect("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueCoding", &resource.Content[numContent].Profile[numProfile].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueUri(numContent int, numProfile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return StringInput("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueUri", &resource.Content[numContent].Profile[numProfile].ValueUri, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueCanonical(numContent int, numProfile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return StringInput("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueCanonical", nil, htmlAttrs)
	}
	return StringInput("content["+strconv.Itoa(numContent)+"].profile["+strconv.Itoa(numProfile)+"].valueCanonical", &resource.Content[numContent].Profile[numProfile].ValueCanonical, htmlAttrs)
}
