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

// http://hl7.org/fhir/r5/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Subject           []Reference                 `json:"subject,omitempty"`
	Type              *CodeableConcept            `json:"type,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Region            []CodeableConcept           `json:"region,omitempty"`
	Status            *CodeableConcept            `json:"status,omitempty"`
	StatusDate        *time.Time                  `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ValidityPeriod    *Period                     `json:"validityPeriod,omitempty"`
	Indication        []CodeableReference         `json:"indication,omitempty"`
	IntendedUse       *CodeableConcept            `json:"intendedUse,omitempty"`
	Basis             []CodeableConcept           `json:"basis,omitempty"`
	Holder            *Reference                  `json:"holder,omitempty"`
	Regulator         *Reference                  `json:"regulator,omitempty"`
	AttachedDocument  []Reference                 `json:"attachedDocument,omitempty"`
	Case              *RegulatedAuthorizationCase `json:"case,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorizationCase struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	DatePeriod        *Period          `json:"datePeriod,omitempty"`
	DateDateTime      *time.Time       `json:"dateDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

type OtherRegulatedAuthorization RegulatedAuthorization

// on convert struct to json, automatically add resourceType=RegulatedAuthorization
func (r RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRegulatedAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherRegulatedAuthorization: OtherRegulatedAuthorization(r),
		ResourceType:                "RegulatedAuthorization",
	})
}
func (r RegulatedAuthorization) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RegulatedAuthorization/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RegulatedAuthorization"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *RegulatedAuthorization) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_Region(numRegion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRegion >= len(resource.Region) {
		return CodeableConceptSelect("Region["+strconv.Itoa(numRegion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Region["+strconv.Itoa(numRegion)+"]", &resource.Region[numRegion], optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_StatusDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_IntendedUse(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("IntendedUse", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IntendedUse", resource.IntendedUse, optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_Basis(numBasis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBasis >= len(resource.Basis) {
		return CodeableConceptSelect("Basis["+strconv.Itoa(numBasis)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Basis["+strconv.Itoa(numBasis)+"]", &resource.Basis[numBasis], optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_CaseType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CaseType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CaseType", resource.Case.Type, optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_CaseStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CaseStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CaseStatus", resource.Case.Status, optionsValueSet, htmlAttrs)
}
func (resource *RegulatedAuthorization) T_CaseDateDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CaseDateDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("CaseDateDateTime", resource.Case.DateDateTime, htmlAttrs)
}
