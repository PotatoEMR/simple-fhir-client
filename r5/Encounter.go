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
	PlannedStartDate   *FhirDateTime          `json:"plannedStartDate,omitempty"`
	PlannedEndDate     *FhirDateTime          `json:"plannedEndDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounter
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
}

// json -> struct, first reject if resourceType != Encounter
func (r *Encounter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Encounter" {
		return errors.New("resourceType not Encounter")
	}
	return json.Unmarshal(data, (*OtherEncounter)(r))
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
func (resource *Encounter) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Class(numClass int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"]", &resource.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ServiceType(numServiceType int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], htmlAttrs)
}
func (resource *Encounter) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Encounter) T_SubjectStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectStatus", resource.SubjectStatus, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_EpisodeOfCare(frs []FhirResource, numEpisodeOfCare int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEpisodeOfCare >= len(resource.EpisodeOfCare) {
		return ReferenceInput(frs, "episodeOfCare["+strconv.Itoa(numEpisodeOfCare)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "episodeOfCare["+strconv.Itoa(numEpisodeOfCare)+"]", &resource.EpisodeOfCare[numEpisodeOfCare], htmlAttrs)
}
func (resource *Encounter) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Encounter) T_CareTeam(frs []FhirResource, numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"]", &resource.CareTeam[numCareTeam], htmlAttrs)
}
func (resource *Encounter) T_PartOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "partOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf", resource.PartOf, htmlAttrs)
}
func (resource *Encounter) T_ServiceProvider(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "serviceProvider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "serviceProvider", resource.ServiceProvider, htmlAttrs)
}
func (resource *Encounter) T_Appointment(frs []FhirResource, numAppointment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAppointment >= len(resource.Appointment) {
		return ReferenceInput(frs, "appointment["+strconv.Itoa(numAppointment)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "appointment["+strconv.Itoa(numAppointment)+"]", &resource.Appointment[numAppointment], htmlAttrs)
}
func (resource *Encounter) T_VirtualService(numVirtualService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVirtualService >= len(resource.VirtualService) {
		return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", nil, htmlAttrs)
	}
	return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", &resource.VirtualService[numVirtualService], htmlAttrs)
}
func (resource *Encounter) T_ActualPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("actualPeriod", nil, htmlAttrs)
	}
	return PeriodInput("actualPeriod", resource.ActualPeriod, htmlAttrs)
}
func (resource *Encounter) T_PlannedStartDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("plannedStartDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("plannedStartDate", resource.PlannedStartDate, htmlAttrs)
}
func (resource *Encounter) T_PlannedEndDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("plannedEndDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("plannedEndDate", resource.PlannedEndDate, htmlAttrs)
}
func (resource *Encounter) T_Length(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("length", nil, htmlAttrs)
	}
	return DurationInput("length", resource.Length, htmlAttrs)
}
func (resource *Encounter) T_Account(frs []FhirResource, numAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAccount >= len(resource.Account) {
		return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", &resource.Account[numAccount], htmlAttrs)
}
func (resource *Encounter) T_DietPreference(numDietPreference int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDietPreference >= len(resource.DietPreference) {
		return CodeableConceptSelect("dietPreference["+strconv.Itoa(numDietPreference)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dietPreference["+strconv.Itoa(numDietPreference)+"]", &resource.DietPreference[numDietPreference], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_SpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialArrangement >= len(resource.SpecialArrangement) {
		return CodeableConceptSelect("specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", &resource.SpecialArrangement[numSpecialArrangement], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_SpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialCourtesy >= len(resource.SpecialCourtesy) {
		return CodeableConceptSelect("specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", &resource.SpecialCourtesy[numSpecialCourtesy], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ParticipantPeriod(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", resource.Participant[numParticipant].Period, htmlAttrs)
}
func (resource *Encounter) T_ParticipantActor(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", resource.Participant[numParticipant].Actor, htmlAttrs)
}
func (resource *Encounter) T_ReasonUse(numReason int, numUse int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) || numUse >= len(resource.Reason[numReason].Use) {
		return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"].use["+strconv.Itoa(numUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"].use["+strconv.Itoa(numUse)+"]", &resource.Reason[numReason].Use[numUse], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ReasonValue(numReason int, numValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) || numValue >= len(resource.Reason[numReason].Value) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"].value["+strconv.Itoa(numValue)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"].value["+strconv.Itoa(numValue)+"]", &resource.Reason[numReason].Value[numValue], htmlAttrs)
}
func (resource *Encounter) T_DiagnosisCondition(numDiagnosis int, numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numCondition >= len(resource.Diagnosis[numDiagnosis].Condition) {
		return CodeableReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].condition["+strconv.Itoa(numCondition)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].condition["+strconv.Itoa(numCondition)+"]", &resource.Diagnosis[numDiagnosis].Condition[numCondition], htmlAttrs)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, numUse int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numUse >= len(resource.Diagnosis[numDiagnosis].Use) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use["+strconv.Itoa(numUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use["+strconv.Itoa(numUse)+"]", &resource.Diagnosis[numDiagnosis].Use[numUse], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionPreAdmissionIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("admission.preAdmissionIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("admission.preAdmissionIdentifier", resource.Admission.PreAdmissionIdentifier, htmlAttrs)
}
func (resource *Encounter) T_AdmissionOrigin(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "admission.origin", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "admission.origin", resource.Admission.Origin, htmlAttrs)
}
func (resource *Encounter) T_AdmissionAdmitSource(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("admission.admitSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("admission.admitSource", resource.Admission.AdmitSource, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionReAdmission(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("admission.reAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("admission.reAdmission", resource.Admission.ReAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_AdmissionDestination(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "admission.destination", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "admission.destination", resource.Admission.Destination, htmlAttrs)
}
func (resource *Encounter) T_AdmissionDischargeDisposition(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("admission.dischargeDisposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("admission.dischargeDisposition", resource.Admission.DischargeDisposition, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationLocation(frs []FhirResource, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"].location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"].location", &resource.Location[numLocation].Location, htmlAttrs)
}
func (resource *Encounter) T_LocationStatus(numLocation int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeSelect("location["+strconv.Itoa(numLocation)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("location["+strconv.Itoa(numLocation)+"].status", resource.Location[numLocation].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationForm(numLocation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].form", resource.Location[numLocation].Form, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationPeriod(numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return PeriodInput("location["+strconv.Itoa(numLocation)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("location["+strconv.Itoa(numLocation)+"].period", resource.Location[numLocation].Period, htmlAttrs)
}
