package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *DocumentReference) T_RelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.RelatesTo[numRelatesTo].Id)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return CodeSelect("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("DocumentReference.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *DocumentReference) T_ContentId(numContent int) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Id", nil)
	}
	return StringInput("DocumentReference.Content["+strconv.Itoa(numContent)+"].Id", resource.Content[numContent].Id)
}
func (resource *DocumentReference) T_ContentFormat(numContent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Content) >= numContent {
		return CodingSelect("DocumentReference.Content["+strconv.Itoa(numContent)+"].Format", nil, optionsValueSet)
	}
	return CodingSelect("DocumentReference.Content["+strconv.Itoa(numContent)+"].Format", resource.Content[numContent].Format, optionsValueSet)
}
func (resource *DocumentReference) T_ContextId() templ.Component {

	if resource == nil {
		return StringInput("DocumentReference.Context.Id", nil)
	}
	return StringInput("DocumentReference.Context.Id", resource.Context.Id)
}
func (resource *DocumentReference) T_ContextEvent(numEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Context.Event) >= numEvent {
		return CodeableConceptSelect("DocumentReference.Context.Event["+strconv.Itoa(numEvent)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Context.Event["+strconv.Itoa(numEvent)+"]", &resource.Context.Event[numEvent], optionsValueSet)
}
func (resource *DocumentReference) T_ContextFacilityType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentReference.Context.FacilityType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Context.FacilityType", resource.Context.FacilityType, optionsValueSet)
}
func (resource *DocumentReference) T_ContextPracticeSetting(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DocumentReference.Context.PracticeSetting", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DocumentReference.Context.PracticeSetting", resource.Context.PracticeSetting, optionsValueSet)
}
