package r4

//generated with command go run ./bultaoreune -nodownload
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
func (resource *Encounter) T_Class(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("Encounter.Class", nil, optionsValueSet)
	}
	return CodingSelect("Encounter.Class", &resource.Class, optionsValueSet)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Encounter.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Encounter) T_ServiceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.ServiceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.ServiceType", resource.ServiceType, optionsValueSet)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Priority", resource.Priority, optionsValueSet)
}
func (resource *Encounter) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Encounter.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Encounter) T_StatusHistoryId(numStatusHistory int) templ.Component {

	if resource == nil || len(resource.StatusHistory) >= numStatusHistory {
		return StringInput("Encounter.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Id", nil)
	}
	return StringInput("Encounter.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Id", resource.StatusHistory[numStatusHistory].Id)
}
func (resource *Encounter) T_StatusHistoryStatus(numStatusHistory int) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil || len(resource.StatusHistory) >= numStatusHistory {
		return CodeSelect("Encounter.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", nil, optionsValueSet)
	}
	return CodeSelect("Encounter.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet)
}
func (resource *Encounter) T_ClassHistoryId(numClassHistory int) templ.Component {

	if resource == nil || len(resource.ClassHistory) >= numClassHistory {
		return StringInput("Encounter.ClassHistory["+strconv.Itoa(numClassHistory)+"].Id", nil)
	}
	return StringInput("Encounter.ClassHistory["+strconv.Itoa(numClassHistory)+"].Id", resource.ClassHistory[numClassHistory].Id)
}
func (resource *Encounter) T_ClassHistoryClass(numClassHistory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ClassHistory) >= numClassHistory {
		return CodingSelect("Encounter.ClassHistory["+strconv.Itoa(numClassHistory)+"].Class", nil, optionsValueSet)
	}
	return CodingSelect("Encounter.ClassHistory["+strconv.Itoa(numClassHistory)+"].Class", &resource.ClassHistory[numClassHistory].Class, optionsValueSet)
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
func (resource *Encounter) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet)
}
func (resource *Encounter) T_DiagnosisRank(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return IntInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Rank", nil)
	}
	return IntInput("Encounter.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Rank", resource.Diagnosis[numDiagnosis].Rank)
}
func (resource *Encounter) T_HospitalizationId() templ.Component {

	if resource == nil {
		return StringInput("Encounter.Hospitalization.Id", nil)
	}
	return StringInput("Encounter.Hospitalization.Id", resource.Hospitalization.Id)
}
func (resource *Encounter) T_HospitalizationAdmitSource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.AdmitSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.AdmitSource", resource.Hospitalization.AdmitSource, optionsValueSet)
}
func (resource *Encounter) T_HospitalizationReAdmission(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.ReAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.ReAdmission", resource.Hospitalization.ReAdmission, optionsValueSet)
}
func (resource *Encounter) T_HospitalizationDietPreference(numDietPreference int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Hospitalization.DietPreference) >= numDietPreference {
		return CodeableConceptSelect("Encounter.Hospitalization.DietPreference["+strconv.Itoa(numDietPreference)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.DietPreference["+strconv.Itoa(numDietPreference)+"]", &resource.Hospitalization.DietPreference[numDietPreference], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationSpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Hospitalization.SpecialCourtesy) >= numSpecialCourtesy {
		return CodeableConceptSelect("Encounter.Hospitalization.SpecialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.SpecialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", &resource.Hospitalization.SpecialCourtesy[numSpecialCourtesy], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationSpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Hospitalization.SpecialArrangement) >= numSpecialArrangement {
		return CodeableConceptSelect("Encounter.Hospitalization.SpecialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.SpecialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", &resource.Hospitalization.SpecialArrangement[numSpecialArrangement], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationDischargeDisposition(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.DischargeDisposition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.DischargeDisposition", resource.Hospitalization.DischargeDisposition, optionsValueSet)
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
func (resource *Encounter) T_LocationPhysicalType(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return CodeableConceptSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].PhysicalType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Encounter.Location["+strconv.Itoa(numLocation)+"].PhysicalType", resource.Location[numLocation].PhysicalType, optionsValueSet)
}
