package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/EnrollmentResponse
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

func (resource *EnrollmentResponse) EnrollmentResponseLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *EnrollmentResponse) EnrollmentResponseStatus() templ.Component {
	optionsValueSet := VSFm_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *EnrollmentResponse) EnrollmentResponseOutcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource != nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", resource.Outcome, optionsValueSet)
}
