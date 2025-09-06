package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *Encounter) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Encounter.Id", nil)
	}
	return StringInput("Encounter.Id", resource.Id)
}
func (resource *Encounter) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Encounter.ImplicitRules", nil)
	}
	return StringInput("Encounter.ImplicitRules", resource.ImplicitRules)
}
func (resource *Encounter) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Encounter.Language", nil, optionsValueSet)
	}
	return CodeSelect("Encounter.Language", resource.Language, optionsValueSet)
}
func (resource *Encounter) T_Status() templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("Encounter.Status", nil, optionsValueSet)
	}
	return CodeSelect("Encounter.Status", &resource.Status, optionsValueSet)
}
func (resource *Encounter) T_Class(numClass int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Class) >= numClass {
		return CodeableConceptSelect("Encounter.Class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Class["+strconv.Itoa(numClass)+"]", &resource.Class[numClass], optionsValueSet)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Priority", resource.Priority, optionsValueSet)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Encounter.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Encounter) T_SubjectStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.SubjectStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.SubjectStatus", resource.SubjectStatus, optionsValueSet)
}
func (resource *Encounter) T_PlannedStartDate() templ.Component {

	if resource == nil {
		return StringInput("Encounter.PlannedStartDate", nil)
	}
	return StringInput("Encounter.PlannedStartDate", resource.PlannedStartDate)
}
func (resource *Encounter) T_PlannedEndDate() templ.Component {

	if resource == nil {
		return StringInput("Encounter.PlannedEndDate", nil)
	}
	return StringInput("Encounter.PlannedEndDate", resource.PlannedEndDate)
}
func (resource *Encounter) T_DietPreference(numDietPreference int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.DietPreference) >= numDietPreference {
		return CodeableConceptSelect("Encounter.DietPreference["+strconv.Itoa(numDietPreference)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.DietPreference["+strconv.Itoa(numDietPreference)+"]", &resource.DietPreference[numDietPreference], optionsValueSet)
}
func (resource *Encounter) T_SpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialArrangement) >= numSpecialArrangement {
		return CodeableConceptSelect("Encounter.SpecialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.SpecialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", &resource.SpecialArrangement[numSpecialArrangement], optionsValueSet)
}
func (resource *Encounter) T_SpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialCourtesy) >= numSpecialCourtesy {
		return CodeableConceptSelect("Encounter.SpecialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.SpecialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", &resource.SpecialCourtesy[numSpecialCourtesy], optionsValueSet)
}
func (resource *Encounter) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("Encounter.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("Encounter.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant || len(resource.Participant[numParticipant].Type) >= numType {
		return CodeableConceptSelect("Encounter.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet)
}
func (resource *Encounter) T_ReasonId(numReason int) templ.Component {

	if resource == nil || len(resource.Reason) >= numReason {
		return StringInput("Encounter.Reason["+strconv.Itoa(numReason)+"].Id", nil)
	}
	return StringInput("Encounter.Reason["+strconv.Itoa(numReason)+"].Id", resource.Reason[numReason].Id)
}
func (resource *Encounter) T_ReasonUse(numReason int, numUse int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Reason) >= numReason || len(resource.Reason[numReason].Use) >= numUse {
		return CodeableConceptSelect("Encounter.Reason["+strconv.Itoa(numReason)+"].Use["+strconv.Itoa(numUse)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Reason["+strconv.Itoa(numReason)+"].Use["+strconv.Itoa(numUse)+"]", &resource.Reason[numReason].Use[numUse], optionsValueSet)
}
func (resource *Encounter) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, numUse int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis || len(resource.Diagnosis[numDiagnosis].Use) >= numUse {
		return CodeableConceptSelect("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use["+strconv.Itoa(numUse)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use["+strconv.Itoa(numUse)+"]", &resource.Diagnosis[numDiagnosis].Use[numUse], optionsValueSet)
}
func (resource *Encounter) T_AdmissionId() templ.Component {

	if resource == nil {
		return StringInput("Encounter.Admission.Id", nil)
	}
	return StringInput("Encounter.Admission.Id", resource.Admission.Id)
}
func (resource *Encounter) T_AdmissionAdmitSource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.AdmitSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Admission.AdmitSource", resource.Admission.AdmitSource, optionsValueSet)
}
func (resource *Encounter) T_AdmissionReAdmission(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.ReAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Admission.ReAdmission", resource.Admission.ReAdmission, optionsValueSet)
}
func (resource *Encounter) T_AdmissionDischargeDisposition(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.DischargeDisposition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Admission.DischargeDisposition", resource.Admission.DischargeDisposition, optionsValueSet)
}
func (resource *Encounter) T_LocationId(numLocation int) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return StringInput("Encounter.Location["+strconv.Itoa(numLocation)+"].Id", nil)
	}
	return StringInput("Encounter.Location["+strconv.Itoa(numLocation)+"].Id", resource.Location[numLocation].Id)
}
func (resource *Encounter) T_LocationStatus(numLocation int) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil || len(resource.Location) >= numLocation {
		return CodeSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].Status", nil, optionsValueSet)
	}
	return CodeSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].Status", resource.Location[numLocation].Status, optionsValueSet)
}
func (resource *Encounter) T_LocationForm(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return CodeableConceptSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].Form", resource.Location[numLocation].Form, optionsValueSet)
}
