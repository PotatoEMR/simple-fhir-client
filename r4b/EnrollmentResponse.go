//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/EnrollmentResponse
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
