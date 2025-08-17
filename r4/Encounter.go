//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type Encounter struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Status            string                    `json:"status"`
	StatusHistory     []EncounterStatusHistory  `json:"statusHistory,omitempty"`
	Class             Coding                    `json:"class"`
	ClassHistory      []EncounterClassHistory   `json:"classHistory,omitempty"`
	Type              []CodeableConcept         `json:"type,omitempty"`
	ServiceType       *CodeableConcept          `json:"serviceType,omitempty"`
	Priority          *CodeableConcept          `json:"priority,omitempty"`
	Subject           *Reference                `json:"subject,omitempty"`
	EpisodeOfCare     []Reference               `json:"episodeOfCare,omitempty"`
	BasedOn           []Reference               `json:"basedOn,omitempty"`
	Participant       []EncounterParticipant    `json:"participant,omitempty"`
	Appointment       []Reference               `json:"appointment,omitempty"`
	Period            *Period                   `json:"period,omitempty"`
	Length            *Duration                 `json:"length,omitempty"`
	ReasonCode        []CodeableConcept         `json:"reasonCode,omitempty"`
	ReasonReference   []Reference               `json:"reasonReference,omitempty"`
	Diagnosis         []EncounterDiagnosis      `json:"diagnosis,omitempty"`
	Account           []Reference               `json:"account,omitempty"`
	Hospitalization   *EncounterHospitalization `json:"hospitalization,omitempty"`
	Location          []EncounterLocation       `json:"location,omitempty"`
	ServiceProvider   *Reference                `json:"serviceProvider,omitempty"`
	PartOf            *Reference                `json:"partOf,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterStatusHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterClassHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Class             Coding      `json:"class"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Individual        *Reference        `json:"individual,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterDiagnosis struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Condition         Reference        `json:"condition"`
	Use               *CodeableConcept `json:"use,omitempty"`
	Rank              *int             `json:"rank,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterHospitalization struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier       `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *Reference        `json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `json:"dischargeDisposition,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Encounter
type EncounterLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Location          Reference        `json:"location"`
	Status            *string          `json:"status,omitempty"`
	PhysicalType      *CodeableConcept `json:"physicalType,omitempty"`
	Period            *Period          `json:"period,omitempty"`
}

type OtherEncounter Encounter

// on convert struct to json, automatically add resourceType=Encounter
func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounter
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
}
