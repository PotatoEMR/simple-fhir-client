package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type Encounter struct {
	Id                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Contained          []Resource             `json:"contained,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	Status             string                 `json:"status"`
	Class              []CodeableConcept      `json:"class,omitempty"`
	Priority           *CodeableConcept       `json:"priority,omitempty"`
	Type               []CodeableConcept      `json:"type,omitempty"`
	ServiceType        []CodeableReference    `json:"serviceType,omitempty"`
	Subject            *Reference             `json:"subject,omitempty"`
	SubjectStatus      *CodeableConcept       `json:"subjectStatus,omitempty"`
	EpisodeOfCare      []Reference            `json:"episodeOfCare,omitempty"`
	BasedOn            []Reference            `json:"basedOn,omitempty"`
	CareTeam           []Reference            `json:"careTeam,omitempty"`
	PartOf             *Reference             `json:"partOf,omitempty"`
	ServiceProvider    *Reference             `json:"serviceProvider,omitempty"`
	Participant        []EncounterParticipant `json:"participant,omitempty"`
	Appointment        []Reference            `json:"appointment,omitempty"`
	VirtualService     []VirtualServiceDetail `json:"virtualService,omitempty"`
	ActualPeriod       *Period                `json:"actualPeriod,omitempty"`
	PlannedStartDate   *string                `json:"plannedStartDate,omitempty"`
	PlannedEndDate     *string                `json:"plannedEndDate,omitempty"`
	Length             *Duration              `json:"length,omitempty"`
	Reason             []EncounterReason      `json:"reason,omitempty"`
	Diagnosis          []EncounterDiagnosis   `json:"diagnosis,omitempty"`
	Account            []Reference            `json:"account,omitempty"`
	DietPreference     []CodeableConcept      `json:"dietPreference,omitempty"`
	SpecialArrangement []CodeableConcept      `json:"specialArrangement,omitempty"`
	SpecialCourtesy    []CodeableConcept      `json:"specialCourtesy,omitempty"`
	Admission          *EncounterAdmission    `json:"admission,omitempty"`
	Location           []EncounterLocation    `json:"location,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type EncounterParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Actor             *Reference        `json:"actor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type EncounterReason struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Use               []CodeableConcept   `json:"use,omitempty"`
	Value             []CodeableReference `json:"value,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type EncounterDiagnosis struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Condition         []CodeableReference `json:"condition,omitempty"`
	Use               []CodeableConcept   `json:"use,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type EncounterAdmission struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier      `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference       `json:"origin,omitempty"`
	AdmitSource            *CodeableConcept `json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept `json:"reAdmission,omitempty"`
	Destination            *Reference       `json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept `json:"dischargeDisposition,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Encounter
type EncounterLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Location          Reference        `json:"location"`
	Status            *string          `json:"status,omitempty"`
	Form              *CodeableConcept `json:"form,omitempty"`
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

func (resource *Encounter) EncounterLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Encounter) EncounterStatus() templ.Component {
	optionsValueSet := VSEncounter_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Encounter) EncounterLocationStatus(numLocation int) templ.Component {
	optionsValueSet := VSEncounter_location_status
	currentVal := ""
	if resource != nil && len(resource.Location) >= numLocation {
		currentVal = *resource.Location[numLocation].Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
