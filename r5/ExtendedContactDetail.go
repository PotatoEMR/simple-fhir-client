package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r5/StructureDefinition/ExtendedContactDetail
type ExtendedContactDetail struct {
	Id           *string          `json:"id,omitempty"`
	Extension    []Extension      `json:"extension,omitempty"`
	Purpose      *CodeableConcept `json:"purpose,omitempty"`
	Name         []HumanName      `json:"name,omitempty"`
	Telecom      []ContactPoint   `json:"telecom,omitempty"`
	Address      *Address         `json:"address,omitempty"`
	Organization *Reference       `json:"organization,omitempty"`
	Period       *Period          `json:"period,omitempty"`
}
