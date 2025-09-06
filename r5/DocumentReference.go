package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *DocumentReference) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.Id", nil)
	}
	return StringInput("DocumentReference.Id", resource.Id)
}
func (resource *DocumentReference) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.ImplicitRules", nil)
	}
	return StringInput("DocumentReference.ImplicitRules", resource.ImplicitRules)
}
func (resource *DocumentReference) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DocumentReference.Language", nil, optionsValueSet)
	}
	return CodeSelect("DocumentReference.Language", resource.Language, optionsValueSet)
}
func (resource *DocumentReference) T_Version() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.Version", nil)
	}
	return StringInput("DocumentReference.Version", resource.Version)
}
func (resource *DocumentReference) T_Status() templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("DocumentReference.Status", nil, optionsValueSet)
	}
	return CodeSelect("DocumentReference.Status", &resource.Status, optionsValueSet)
}
func (resource *DocumentReference) T_DocStatus() templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("DocumentReference.DocStatus", nil, optionsValueSet)
	}
	return CodeSelect("DocumentReference.DocStatus", resource.DocStatus, optionsValueSet)
}
func (resource *DocumentReference) T_Modality(numModality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Modality) >= numModality {
		return CodeableConceptSelect("DocumentReference.Modality["+strconv.Itoa(numModality)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Modality["+strconv.Itoa(numModality)+"]", &resource.Modality[numModality], optionsValueSet)
}
func (resource *DocumentReference) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentReference.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Type", resource.Type, optionsValueSet)
}
func (resource *DocumentReference) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("DocumentReference.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *DocumentReference) T_FacilityType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentReference.FacilityType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.FacilityType", resource.FacilityType, optionsValueSet)
}
func (resource *DocumentReference) T_PracticeSetting(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentReference.PracticeSetting", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.PracticeSetting", resource.PracticeSetting, optionsValueSet)
}
func (resource *DocumentReference) T_Date() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.Date", nil)
	}
	return StringInput("DocumentReference.Date", resource.Date)
}
func (resource *DocumentReference) T_Description() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.Description", nil)
	}
	return StringInput("DocumentReference.Description", resource.Description)
}
func (resource *DocumentReference) T_SecurityLabel(numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SecurityLabel) >= numSecurityLabel {
		return CodeableConceptSelect("DocumentReference.SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.SecurityLabel[numSecurityLabel], optionsValueSet)
}
func (resource *DocumentReference) T_AttesterId(numAttester int) templ.Component {

	if resource == nil || len(resource.Attester) >= numAttester {
		return StringInput("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Id", nil)
	}
	return StringInput("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Id", resource.Attester[numAttester].Id)
}
func (resource *DocumentReference) T_AttesterMode(numAttester int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Attester) >= numAttester {
		return CodeableConceptSelect("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Mode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Mode", &resource.Attester[numAttester].Mode, optionsValueSet)
}
func (resource *DocumentReference) T_AttesterTime(numAttester int) templ.Component {

	if resource == nil || len(resource.Attester) >= numAttester {
		return StringInput("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Time", nil)
	}
	return StringInput("DocumentReference.Attester["+strconv.Itoa(numAttester)+"].Time", resource.Attester[numAttester].Time)
}
func (resource *DocumentReference) T_RelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.RelatesTo[numRelatesTo].Id)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return CodeableConceptSelect("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *DocumentReference) T_ContentId(numContent int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Id", nil)
	}
	return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Id", resource.Content[numContent].Id)
}
func (resource *DocumentReference) T_ContentProfileId(numContent int, numProfile int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent || len(resource.Content[numContent].Profile) >= numProfile {
		return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Profile["+strconv.Itoa(numProfile)+"].Id", nil)
	}
	return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Profile["+strconv.Itoa(numProfile)+"].Id", resource.Content[numContent].Profile[numProfile].Id)
}
