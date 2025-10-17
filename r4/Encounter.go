package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r Encounter) ResourceType() string {
	return "Encounter"
}

func (resource *Encounter) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Class(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("class", &resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ServiceType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceType", resource.ServiceType, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
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
func (resource *Encounter) T_Appointment(frs []FhirResource, numAppointment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAppointment >= len(resource.Appointment) {
		return ReferenceInput(frs, "appointment["+strconv.Itoa(numAppointment)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "appointment["+strconv.Itoa(numAppointment)+"]", &resource.Appointment[numAppointment], htmlAttrs)
}
func (resource *Encounter) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *Encounter) T_Length(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("length", nil, htmlAttrs)
	}
	return DurationInput("length", resource.Length, htmlAttrs)
}
func (resource *Encounter) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *Encounter) T_Account(frs []FhirResource, numAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAccount >= len(resource.Account) {
		return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", &resource.Account[numAccount], htmlAttrs)
}
func (resource *Encounter) T_ServiceProvider(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "serviceProvider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "serviceProvider", resource.ServiceProvider, htmlAttrs)
}
func (resource *Encounter) T_PartOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "partOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf", resource.PartOf, htmlAttrs)
}
func (resource *Encounter) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_StatusHistoryPeriod(numStatusHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return PeriodInput("statusHistory["+strconv.Itoa(numStatusHistory)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("statusHistory["+strconv.Itoa(numStatusHistory)+"].period", &resource.StatusHistory[numStatusHistory].Period, htmlAttrs)
}
func (resource *Encounter) T_ClassHistoryClass(numClassHistory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassHistory >= len(resource.ClassHistory) {
		return CodingSelect("classHistory["+strconv.Itoa(numClassHistory)+"].class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("classHistory["+strconv.Itoa(numClassHistory)+"].class", &resource.ClassHistory[numClassHistory].Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ClassHistoryPeriod(numClassHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassHistory >= len(resource.ClassHistory) {
		return PeriodInput("classHistory["+strconv.Itoa(numClassHistory)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("classHistory["+strconv.Itoa(numClassHistory)+"].period", &resource.ClassHistory[numClassHistory].Period, htmlAttrs)
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
func (resource *Encounter) T_ParticipantIndividual(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].individual", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].individual", resource.Participant[numParticipant].Individual, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisCondition(frs []FhirResource, numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", &resource.Diagnosis[numDiagnosis].Condition, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisRank(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", resource.Diagnosis[numDiagnosis].Rank, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationPreAdmissionIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("hospitalization.preAdmissionIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("hospitalization.preAdmissionIdentifier", resource.Hospitalization.PreAdmissionIdentifier, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationOrigin(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "hospitalization.origin", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "hospitalization.origin", resource.Hospitalization.Origin, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationAdmitSource(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.admitSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.admitSource", resource.Hospitalization.AdmitSource, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationReAdmission(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.reAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.reAdmission", resource.Hospitalization.ReAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDietPreference(numDietPreference int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDietPreference >= len(resource.Hospitalization.DietPreference) {
		return CodeableConceptSelect("hospitalization.dietPreference["+strconv.Itoa(numDietPreference)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.dietPreference["+strconv.Itoa(numDietPreference)+"]", &resource.Hospitalization.DietPreference[numDietPreference], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialCourtesy >= len(resource.Hospitalization.SpecialCourtesy) {
		return CodeableConceptSelect("hospitalization.specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", &resource.Hospitalization.SpecialCourtesy[numSpecialCourtesy], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialArrangement >= len(resource.Hospitalization.SpecialArrangement) {
		return CodeableConceptSelect("hospitalization.specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", &resource.Hospitalization.SpecialArrangement[numSpecialArrangement], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDestination(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "hospitalization.destination", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "hospitalization.destination", resource.Hospitalization.Destination, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDischargeDisposition(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.dischargeDisposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.dischargeDisposition", resource.Hospitalization.DischargeDisposition, optionsValueSet, htmlAttrs)
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
func (resource *Encounter) T_LocationPhysicalType(numLocation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].physicalType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].physicalType", resource.Location[numLocation].PhysicalType, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationPeriod(numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return PeriodInput("location["+strconv.Itoa(numLocation)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("location["+strconv.Itoa(numLocation)+"].period", resource.Location[numLocation].Period, htmlAttrs)
}
