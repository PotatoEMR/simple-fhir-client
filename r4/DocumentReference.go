package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []Resource                   `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier                  `json:"masterIdentifier,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Status            string                       `json:"status"`
	DocStatus         *string                      `json:"docStatus,omitempty"`
	Type              *CodeableConcept             `json:"type,omitempty"`
	Category          []CodeableConcept            `json:"category,omitempty"`
	Subject           *Reference                   `json:"subject,omitempty"`
	Date              *string                      `json:"date,omitempty"`
	Author            []Reference                  `json:"author,omitempty"`
	Authenticator     *Reference                   `json:"authenticator,omitempty"`
	Custodian         *Reference                   `json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `json:"content"`
	Context           *DocumentReferenceContext    `json:"context,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DocumentReference
type DocumentReferenceRelatesTo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Target            Reference   `json:"target"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DocumentReference
type DocumentReferenceContent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Attachment        Attachment  `json:"attachment"`
	Format            *Coding     `json:"format,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DocumentReference
type DocumentReferenceContext struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Encounter         []Reference       `json:"encounter,omitempty"`
	Event             []CodeableConcept `json:"event,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	FacilityType      *CodeableConcept  `json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept  `json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference        `json:"sourcePatientInfo,omitempty"`
	Related           []Reference       `json:"related,omitempty"`
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

func (resource *DocumentReference) DocumentReferenceLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceStatus() templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceDocStatus() templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("docStatus", nil, optionsValueSet)
	}
	return CodeSelect("docStatus", resource.DocStatus, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceSecurityLabel(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("securityLabel", nil, optionsValueSet)
	}
	return CodeableConceptSelect("securityLabel", &resource.SecurityLabel[0], optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceRelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil && len(resource.RelatesTo) >= numRelatesTo {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceContentFormat(numContent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Content) >= numContent {
		return CodingSelect("format", nil, optionsValueSet)
	}
	return CodingSelect("format", resource.Content[numContent].Format, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceContextEvent(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("event", nil, optionsValueSet)
	}
	return CodeableConceptSelect("event", &resource.Context.Event[0], optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceContextFacilityType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("facilityType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("facilityType", resource.Context.FacilityType, optionsValueSet)
}
func (resource *DocumentReference) DocumentReferenceContextPracticeSetting(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("practiceSetting", nil, optionsValueSet)
	}
	return CodeableConceptSelect("practiceSetting", resource.Context.PracticeSetting, optionsValueSet)
}
