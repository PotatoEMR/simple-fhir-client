package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
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
	Created           *FhirDateTime `json:"created,omitempty"`
	Insurer           *Reference    `json:"insurer,omitempty"`
	Provider          *Reference    `json:"provider,omitempty"`
	Candidate         *Reference    `json:"candidate,omitempty"`
	Coverage          *Reference    `json:"coverage,omitempty"`
}

type OtherEnrollmentRequest EnrollmentRequest

// struct -> json, automatically add resourceType=Patient
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
}

// json -> struct, first reject if resourceType != EnrollmentRequest
func (r *EnrollmentRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "EnrollmentRequest" {
		return errors.New("resourceType not EnrollmentRequest")
	}
	return json.Unmarshal(data, (*OtherEnrollmentRequest)(r))
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
func (resource *EnrollmentRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", resource.Created, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Insurer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "insurer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurer", resource.Insurer, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Provider(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "provider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "provider", resource.Provider, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Candidate(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "candidate", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "candidate", resource.Candidate, htmlAttrs)
}
func (resource *EnrollmentRequest) T_Coverage(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "coverage", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "coverage", resource.Coverage, htmlAttrs)
}
