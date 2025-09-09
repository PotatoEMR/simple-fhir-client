package r4

//generated with command go run ./bultaoreune
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
func (resource *DocumentReference) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_reference_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_DocStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("docStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("docStatus", resource.DocStatus, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("date", nil, htmlAttrs)
	}
	return StringInput("date", resource.Date, htmlAttrs)
}
func (resource *DocumentReference) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *DocumentReference) T_SecurityLabel(numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSecurityLabel >= len(resource.SecurityLabel) {
		return CodeableConceptSelect("securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_RelatesToCode(numRelatesTo int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContentFormat(numContent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContent >= len(resource.Content) {
		return CodingSelect("content["+strconv.Itoa(numContent)+"].format", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("content["+strconv.Itoa(numContent)+"].format", resource.Content[numContent].Format, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContextEvent(numEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Context.Event) {
		return CodeableConceptSelect("context.event["+strconv.Itoa(numEvent)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("context.event["+strconv.Itoa(numEvent)+"]", &resource.Context.Event[numEvent], optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContextFacilityType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("context.facilityType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("context.facilityType", resource.Context.FacilityType, optionsValueSet, htmlAttrs)
}
func (resource *DocumentReference) T_ContextPracticeSetting(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("context.practiceSetting", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("context.practiceSetting", resource.Context.PracticeSetting, optionsValueSet, htmlAttrs)
}
