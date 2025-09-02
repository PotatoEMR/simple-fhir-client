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

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Encounter) EncounterStatus() templ.Component {
	optionsValueSet := VSEncounter_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Encounter) EncounterClass(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("class", &resource.Class[0], optionsValueSet)
}
func (resource *Encounter) EncounterPriority(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Encounter) EncounterType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *Encounter) EncounterSubjectStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("subjectStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subjectStatus", resource.SubjectStatus, optionsValueSet)
}
func (resource *Encounter) EncounterDietPreference(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("dietPreference", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dietPreference", &resource.DietPreference[0], optionsValueSet)
}
func (resource *Encounter) EncounterSpecialArrangement(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("specialArrangement", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialArrangement", &resource.SpecialArrangement[0], optionsValueSet)
}
func (resource *Encounter) EncounterSpecialCourtesy(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("specialCourtesy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialCourtesy", &resource.SpecialCourtesy[0], optionsValueSet)
}
func (resource *Encounter) EncounterParticipantType(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Participant[numParticipant].Type[0], optionsValueSet)
}
func (resource *Encounter) EncounterReasonUse(numReason int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Reason) >= numReason {
		return CodeableConceptSelect("use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("use", &resource.Reason[numReason].Use[0], optionsValueSet)
}
func (resource *Encounter) EncounterDiagnosisUse(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("use", &resource.Diagnosis[numDiagnosis].Use[0], optionsValueSet)
}
func (resource *Encounter) EncounterAdmissionAdmitSource(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("admitSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("admitSource", resource.Admission.AdmitSource, optionsValueSet)
}
func (resource *Encounter) EncounterAdmissionReAdmission(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reAdmission", resource.Admission.ReAdmission, optionsValueSet)
}
func (resource *Encounter) EncounterAdmissionDischargeDisposition(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("dischargeDisposition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dischargeDisposition", resource.Admission.DischargeDisposition, optionsValueSet)
}
func (resource *Encounter) EncounterLocationStatus(numLocation int) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource != nil && len(resource.Location) >= numLocation {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Location[numLocation].Status, optionsValueSet)
}
func (resource *Encounter) EncounterLocationForm(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Location) >= numLocation {
		return CodeableConceptSelect("form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("form", resource.Location[numLocation].Form, optionsValueSet)
}
