package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	Id                *string       `json:"id,omitempty"`
	Meta              *Meta         `json:"meta,omitempty"`
	ImplicitRules     *string       `json:"implicitRules,omitempty"`
	Language          *string       `json:"language,omitempty"`
	Text              *Narrative    `json:"text,omitempty"`
	Contained         []Resource    `json:"contained,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Identifier        []Identifier  `json:"identifier,omitempty"`
	Status            *string       `json:"status,omitempty"`
	Request           *Reference    `json:"request,omitempty"`
	Outcome           *string       `json:"outcome,omitempty"`
	Disposition       *string       `json:"disposition,omitempty"`
	Created           *FhirDateTime `json:"created,omitempty"`
	Organization      *Reference    `json:"organization,omitempty"`
	RequestProvider   *Reference    `json:"requestProvider,omitempty"`
}

type OtherEnrollmentResponse EnrollmentResponse

// struct -> json, automatically add resourceType=Patient
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
}

// json -> struct, first reject if resourceType != EnrollmentResponse
func (r *EnrollmentResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "EnrollmentResponse" {
		return errors.New("resourceType not EnrollmentResponse")
	}
	return json.Unmarshal(data, (*OtherEnrollmentResponse)(r))
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
func (resource *EnrollmentResponse) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Request(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request", resource.Request, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", resource.Created, htmlAttrs)
}
func (resource *EnrollmentResponse) T_Organization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "organization", resource.Organization, htmlAttrs)
}
func (resource *EnrollmentResponse) T_RequestProvider(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requestProvider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requestProvider", resource.RequestProvider, htmlAttrs)
}
