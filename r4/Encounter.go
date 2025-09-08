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
func (resource *Encounter) T_Class(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("Encounter.Class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Encounter.Class", &resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Encounter.Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Type."+strconv.Itoa(numType)+".", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ServiceType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.ServiceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.ServiceType", resource.ServiceType, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Encounter.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("Encounter.StatusHistory."+strconv.Itoa(numStatusHistory)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Encounter.StatusHistory."+strconv.Itoa(numStatusHistory)+"..Status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ClassHistoryClass(numClassHistory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClassHistory >= len(resource.ClassHistory) {
		return CodingSelect("Encounter.ClassHistory."+strconv.Itoa(numClassHistory)+"..Class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Encounter.ClassHistory."+strconv.Itoa(numClassHistory)+"..Class", &resource.ClassHistory[numClassHistory].Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("Encounter.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisRank(numDiagnosis int, htmlAttrs string) templ.Component {

	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Rank", nil, htmlAttrs)
	}
	return IntInput("Encounter.Diagnosis."+strconv.Itoa(numDiagnosis)+"..Rank", resource.Diagnosis[numDiagnosis].Rank, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationAdmitSource(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.AdmitSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.AdmitSource", resource.Hospitalization.AdmitSource, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationReAdmission(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.ReAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.ReAdmission", resource.Hospitalization.ReAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDietPreference(numDietPreference int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDietPreference >= len(resource.Hospitalization.DietPreference) {
		return CodeableConceptSelect("Encounter.Hospitalization.DietPreference."+strconv.Itoa(numDietPreference)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.DietPreference."+strconv.Itoa(numDietPreference)+".", &resource.Hospitalization.DietPreference[numDietPreference], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialCourtesy >= len(resource.Hospitalization.SpecialCourtesy) {
		return CodeableConceptSelect("Encounter.Hospitalization.SpecialCourtesy."+strconv.Itoa(numSpecialCourtesy)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.SpecialCourtesy."+strconv.Itoa(numSpecialCourtesy)+".", &resource.Hospitalization.SpecialCourtesy[numSpecialCourtesy], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialArrangement >= len(resource.Hospitalization.SpecialArrangement) {
		return CodeableConceptSelect("Encounter.Hospitalization.SpecialArrangement."+strconv.Itoa(numSpecialArrangement)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.SpecialArrangement."+strconv.Itoa(numSpecialArrangement)+".", &resource.Hospitalization.SpecialArrangement[numSpecialArrangement], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDischargeDisposition(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Encounter.Hospitalization.DischargeDisposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Hospitalization.DischargeDisposition", resource.Hospitalization.DischargeDisposition, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationStatus(numLocation int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..Status", resource.Location[numLocation].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationPhysicalType(numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..PhysicalType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Encounter.Location."+strconv.Itoa(numLocation)+"..PhysicalType", resource.Location[numLocation].PhysicalType, optionsValueSet, htmlAttrs)
}
