package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	StatusDate        *string                     `json:"statusDate,omitempty"`
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
	DatePeriod        *Period          `json:"datePeriod"`
	DateDateTime      *string          `json:"dateDateTime"`
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

func (resource *RegulatedAuthorization) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Region(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("region", nil, optionsValueSet)
	}
	return CodeableConceptSelect("region", &resource.Region[0], optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_IntendedUse(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("intendedUse", nil, optionsValueSet)
	}
	return CodeableConceptSelect("intendedUse", resource.IntendedUse, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Basis(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("basis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("basis", &resource.Basis[0], optionsValueSet)
}
func (resource *RegulatedAuthorization) T_CaseType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Case.Type, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_CaseStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Case.Status, optionsValueSet)
}
