package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Active            *bool                   `json:"active,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	Practitioner      *Reference              `json:"practitioner,omitempty"`
	Organization      *Reference              `json:"organization,omitempty"`
	Code              []CodeableConcept       `json:"code,omitempty"`
	Specialty         []CodeableConcept       `json:"specialty,omitempty"`
	Location          []Reference             `json:"location,omitempty"`
	HealthcareService []Reference             `json:"healthcareService,omitempty"`
	Contact           []ExtendedContactDetail `json:"contact,omitempty"`
	Characteristic    []CodeableConcept       `json:"characteristic,omitempty"`
	Communication     []CodeableConcept       `json:"communication,omitempty"`
	Availability      []Availability          `json:"availability,omitempty"`
	Endpoint          []Reference             `json:"endpoint,omitempty"`
}

type OtherPractitionerRole PractitionerRole

// struct -> json, automatically add resourceType=Patient
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

// json -> struct, first reject if resourceType != PractitionerRole
func (r *PractitionerRole) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "PractitionerRole" {
		return errors.New("resourceType not PractitionerRole")
	}
	return json.Unmarshal(data, (*OtherPractitionerRole)(r))
}

func (r PractitionerRole) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PractitionerRole/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PractitionerRole"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PractitionerRole) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *PractitionerRole) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *PractitionerRole) T_Practitioner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "practitioner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "practitioner", resource.Practitioner, htmlAttrs)
}
func (resource *PractitionerRole) T_Organization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "organization", resource.Organization, htmlAttrs)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Location(frs []FhirResource, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], htmlAttrs)
}
func (resource *PractitionerRole) T_HealthcareService(frs []FhirResource, numHealthcareService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHealthcareService >= len(resource.HealthcareService) {
		return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", &resource.HealthcareService[numHealthcareService], htmlAttrs)
}
func (resource *PractitionerRole) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *PractitionerRole) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Communication(numCommunication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Availability(numAvailability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAvailability >= len(resource.Availability) {
		return AvailabilityInput("availability["+strconv.Itoa(numAvailability)+"]", nil, htmlAttrs)
	}
	return AvailabilityInput("availability["+strconv.Itoa(numAvailability)+"]", &resource.Availability[numAvailability], htmlAttrs)
}
func (resource *PractitionerRole) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
