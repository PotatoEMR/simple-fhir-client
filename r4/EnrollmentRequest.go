package r4

//generated with command go run ./bultaoreune
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
func (r EnrollmentRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EnrollmentRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EnrollmentRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EnrollmentRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("created", nil, htmlAttrs)
	}
	return DateTimeInput("created", resource.Created, htmlAttrs)
}
