package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"time"

	"github.com/a-h/templ"
)

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
	Created           *time.Time   `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r EnrollmentResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EnrollmentResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EnrollmentResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EnrollmentResponse) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Disposition", nil, htmlAttrs)
	}
	return StringInput("Disposition", resource.Disposition, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Created", nil, htmlAttrs)
	}
	return DateTimeInput("Created", resource.Created, htmlAttrs)
}
