package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	PlannedStartDate   *time.Time             `json:"plannedStartDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PlannedEndDate     *time.Time             `json:"plannedEndDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Encounter) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Encounter/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Encounter"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Encounter) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("Encounter.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Encounter.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Class(numClass int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("Encounter.Class."+strconv.Itoa(numClass)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Class."+strconv.Itoa(numClass)+".", &resource.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Encounter.Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Type."+strconv.Itoa(numType)+".", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_SubjectStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.SubjectStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.SubjectStatus", resource.SubjectStatus, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_PlannedStartDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Encounter.PlannedStartDate", nil, htmlAttrs)
	}
	return DateTimeInput("Encounter.PlannedStartDate", resource.PlannedStartDate, htmlAttrs)
}
func (resource *Encounter) T_PlannedEndDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Encounter.PlannedEndDate", nil, htmlAttrs)
	}
	return DateTimeInput("Encounter.PlannedEndDate", resource.PlannedEndDate, htmlAttrs)
}
func (resource *Encounter) T_DietPreference(numDietPreference int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDietPreference >= len(resource.DietPreference) {
		return CodeableConceptSelect("Encounter.DietPreference."+strconv.Itoa(numDietPreference)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.DietPreference."+strconv.Itoa(numDietPreference)+".", &resource.DietPreference[numDietPreference], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_SpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialArrangement >= len(resource.SpecialArrangement) {
		return CodeableConceptSelect("Encounter.SpecialArrangement."+strconv.Itoa(numSpecialArrangement)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.SpecialArrangement."+strconv.Itoa(numSpecialArrangement)+".", &resource.SpecialArrangement[numSpecialArrangement], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_SpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialCourtesy >= len(resource.SpecialCourtesy) {
		return CodeableConceptSelect("Encounter.SpecialCourtesy."+strconv.Itoa(numSpecialCourtesy)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.SpecialCourtesy."+strconv.Itoa(numSpecialCourtesy)+".", &resource.SpecialCourtesy[numSpecialCourtesy], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("Encounter.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ReasonUse(numReason int, numUse int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReason >= len(resource.Reason) || numUse >= len(resource.Reason[numReason].Use) {
		return CodeableConceptSelect("Encounter.Reason."+strconv.Itoa(numReason)+"..Use."+strconv.Itoa(numUse)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Reason."+strconv.Itoa(numReason)+"..Use."+strconv.Itoa(numUse)+".", &resource.Reason[numReason].Use[numUse], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, numUse int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numUse >= len(resource.Diagnosis[numDiagnosis].Use) {
		return CodeableConceptSelect("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Use."+strconv.Itoa(numUse)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Use."+strconv.Itoa(numUse)+".", &resource.Diagnosis[numDiagnosis].Use[numUse], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionAdmitSource(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.AdmitSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Admission.AdmitSource", resource.Admission.AdmitSource, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionReAdmission(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.ReAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Admission.ReAdmission", resource.Admission.ReAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionDischargeDisposition(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Admission.DischargeDisposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Admission.DischargeDisposition", resource.Admission.DischargeDisposition, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationStatus(numLocation int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Status", resource.Location[numLocation].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationForm(numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Form", resource.Location[numLocation].Form, optionsValueSet, htmlAttrs)
}
