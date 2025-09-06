package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
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
	Created           *string      `json:"created,omitempty"`
	Insurer           *Reference   `json:"insurer,omitempty"`
	Provider          *Reference   `json:"provider,omitempty"`
	Candidate         *Reference   `json:"candidate,omitempty"`
	Coverage          *Reference   `json:"coverage,omitempty"`
}

type OtherEnrollmentRequest EnrollmentRequest

// on convert struct to json, automatically add resourceType=EnrollmentRequest
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
}

func (resource *EnrollmentRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentRequest.Id", nil)
	}
	return StringInput("EnrollmentRequest.Id", resource.Id)
}
func (resource *EnrollmentRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentRequest.ImplicitRules", nil)
	}
	return StringInput("EnrollmentRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *EnrollmentRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EnrollmentRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("EnrollmentRequest.Language", resource.Language, optionsValueSet)
}
func (resource *EnrollmentRequest) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("EnrollmentRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("EnrollmentRequest.Status", resource.Status, optionsValueSet)
}
func (resource *EnrollmentRequest) T_Created() templ.Component {

	if resource == nil {
		return StringInput("EnrollmentRequest.Created", nil)
	}
	return StringInput("EnrollmentRequest.Created", resource.Created)
}
