package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	Id                *string      `json:"id,omitempty"`
	Meta              *Meta        `json:"meta,omitempty"`
	ImplicitRules     *string      `json:"implicitRules,omitempty"`
	Language          *string      `json:"language,omitempty"`
	Text              *Narrative   `json:"text,omitempty"`
	Contained         []Resource   `json:"contained,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier `json:"identifier,omitempty"`
	Status            *string      `json:"status,omitempty"`
	Request           *Reference   `json:"request,omitempty"`
	Outcome           *string      `json:"outcome,omitempty"`
	Disposition       *string      `json:"disposition,omitempty"`
	Created           *string      `json:"created,omitempty"`
	Organization      *Reference   `json:"organization,omitempty"`
	RequestProvider   *Reference   `json:"requestProvider,omitempty"`
}

type OtherEnrollmentResponse EnrollmentResponse

// on convert struct to json, automatically add resourceType=EnrollmentResponse
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
}

func (resource *EnrollmentResponse) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentResponse.Id", nil)
	}
	return StringInput("EnrollmentResponse.Id", resource.Id)
}
func (resource *EnrollmentResponse) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentResponse.ImplicitRules", nil)
	}
	return StringInput("EnrollmentResponse.ImplicitRules", resource.ImplicitRules)
}
func (resource *EnrollmentResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EnrollmentResponse.Language", nil, optionsValueSet)
	}
	return CodeSelect("EnrollmentResponse.Language", resource.Language, optionsValueSet)
}
func (resource *EnrollmentResponse) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("EnrollmentResponse.Status", nil, optionsValueSet)
	}
	return CodeSelect("EnrollmentResponse.Status", resource.Status, optionsValueSet)
}
func (resource *EnrollmentResponse) T_Outcome() templ.Component {
	optionsValueSet := VSEnrollment_outcome

	if resource == nil {
		return CodeSelect("EnrollmentResponse.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("EnrollmentResponse.Outcome", resource.Outcome, optionsValueSet)
}
func (resource *EnrollmentResponse) T_Disposition() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentResponse.Disposition", nil)
	}
	return StringInput("EnrollmentResponse.Disposition", resource.Disposition)
}
func (resource *EnrollmentResponse) T_Created() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentResponse.Created", nil)
	}
	return StringInput("EnrollmentResponse.Created", resource.Created)
}
