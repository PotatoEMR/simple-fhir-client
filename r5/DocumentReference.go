package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *DocumentReference) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DocumentReference) T_Status() templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *DocumentReference) T_DocStatus() templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("docStatus", nil, optionsValueSet)
	}
	return CodeSelect("docStatus", resource.DocStatus, optionsValueSet)
}
func (resource *DocumentReference) T_Modality(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("modality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modality", &resource.Modality[0], optionsValueSet)
}
func (resource *DocumentReference) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *DocumentReference) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *DocumentReference) T_FacilityType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("facilityType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("facilityType", resource.FacilityType, optionsValueSet)
}
func (resource *DocumentReference) T_PracticeSetting(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("practiceSetting", nil, optionsValueSet)
	}
	return CodeableConceptSelect("practiceSetting", resource.PracticeSetting, optionsValueSet)
}
func (resource *DocumentReference) T_SecurityLabel(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("securityLabel", nil, optionsValueSet)
	}
	return CodeableConceptSelect("securityLabel", &resource.SecurityLabel[0], optionsValueSet)
}
func (resource *DocumentReference) T_AttesterMode(numAttester int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Attester) >= numAttester {
		return CodeableConceptSelect("mode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("mode", &resource.Attester[numAttester].Mode, optionsValueSet)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.RelatesTo) >= numRelatesTo {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
