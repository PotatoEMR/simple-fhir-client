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
	Time              *time.Time      `json:"time,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *DocumentReference) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *DocumentReference) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_DocStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("DocStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DocStatus", resource.DocStatus, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Modality(numModality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numModality >= len(resource.Modality) {
		return CodeableConceptSelect("Modality["+strconv.Itoa(numModality)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Modality["+strconv.Itoa(numModality)+"]", &resource.Modality[numModality], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_FacilityType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FacilityType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FacilityType", resource.FacilityType, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_PracticeSetting(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PracticeSetting", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PracticeSetting", resource.PracticeSetting, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Date", nil, htmlAttrs)
	}
	return StringInput("Date", resource.Date, htmlAttrs)
}
func (resource *DocumentReference) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *DocumentReference) T_SecurityLabel(numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSecurityLabel >= len(resource.SecurityLabel) {
		return CodeableConceptSelect("SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_AttesterMode(numAttester int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeableConceptSelect("Attester["+strconv.Itoa(numAttester)+"]Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Attester["+strconv.Itoa(numAttester)+"]Mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_AttesterTime(numAttester int, htmlAttrs string) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return DateTimeInput("Attester["+strconv.Itoa(numAttester)+"]Time", nil, htmlAttrs)
	}
	return DateTimeInput("Attester["+strconv.Itoa(numAttester)+"]Time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeableConceptSelect("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueCoding(numContent int, numProfile int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return CodingSelect("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueCoding", &resource.Content[numContent].Profile[numProfile].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueUri(numContent int, numProfile int, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return StringInput("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueUri", nil, htmlAttrs)
	}
	return StringInput("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueUri", &resource.Content[numContent].Profile[numProfile].ValueUri, htmlAttrs)
}
func (resource *DocumentReference) T_ContentProfileValueCanonical(numContent int, numProfile int, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) || numProfile >= len(resource.Content[numContent].Profile) {
		return StringInput("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueCanonical", nil, htmlAttrs)
	}
	return StringInput("Content["+strconv.Itoa(numContent)+"]Profile["+strconv.Itoa(numProfile)+"].ValueCanonical", &resource.Content[numContent].Profile[numProfile].ValueCanonical, htmlAttrs)
}
