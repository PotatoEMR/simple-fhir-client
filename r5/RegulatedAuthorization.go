package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	DatePeriod        *Period          `json:"datePeriod,omitempty"`
	DateDateTime      *string          `json:"dateDateTime,omitempty"`
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

func (resource *RegulatedAuthorization) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RegulatedAuthorization.Id", nil)
	}
	return StringInput("RegulatedAuthorization.Id", resource.Id)
}
func (resource *RegulatedAuthorization) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RegulatedAuthorization.ImplicitRules", nil)
	}
	return StringInput("RegulatedAuthorization.ImplicitRules", resource.ImplicitRules)
}
func (resource *RegulatedAuthorization) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RegulatedAuthorization.Language", nil, optionsValueSet)
	}
	return CodeSelect("RegulatedAuthorization.Language", resource.Language, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RegulatedAuthorization.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Type", resource.Type, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Description() templ.Component {

	if resource == nil {
		return StringInput("RegulatedAuthorization.Description", nil)
	}
	return StringInput("RegulatedAuthorization.Description", resource.Description)
}
func (resource *RegulatedAuthorization) T_Region(numRegion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Region) >= numRegion {
		return CodeableConceptSelect("RegulatedAuthorization.Region["+strconv.Itoa(numRegion)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Region["+strconv.Itoa(numRegion)+"]", &resource.Region[numRegion], optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RegulatedAuthorization.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Status", resource.Status, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("RegulatedAuthorization.StatusDate", nil)
	}
	return StringInput("RegulatedAuthorization.StatusDate", resource.StatusDate)
}
func (resource *RegulatedAuthorization) T_IntendedUse(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RegulatedAuthorization.IntendedUse", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.IntendedUse", resource.IntendedUse, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_Basis(numBasis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Basis) >= numBasis {
		return CodeableConceptSelect("RegulatedAuthorization.Basis["+strconv.Itoa(numBasis)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Basis["+strconv.Itoa(numBasis)+"]", &resource.Basis[numBasis], optionsValueSet)
}
func (resource *RegulatedAuthorization) T_CaseId() templ.Component {

	if resource == nil {
		return StringInput("RegulatedAuthorization.Case.Id", nil)
	}
	return StringInput("RegulatedAuthorization.Case.Id", resource.Case.Id)
}
func (resource *RegulatedAuthorization) T_CaseType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RegulatedAuthorization.Case.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Case.Type", resource.Case.Type, optionsValueSet)
}
func (resource *RegulatedAuthorization) T_CaseStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RegulatedAuthorization.Case.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RegulatedAuthorization.Case.Status", resource.Case.Status, optionsValueSet)
}
