package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Encounter) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Encounter) T_Status() templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Encounter) T_Class(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("class", nil, optionsValueSet)
	}
	return CodingSelect("class", &resource.Class, optionsValueSet)
}
func (resource *Encounter) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *Encounter) T_ServiceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceType", resource.ServiceType, optionsValueSet)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Encounter) T_ReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *Encounter) T_StatusHistoryStatus(numStatusHistory int) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil && len(resource.StatusHistory) >= numStatusHistory {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet)
}
func (resource *Encounter) T_ClassHistoryClass(numClassHistory int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ClassHistory) >= numClassHistory {
		return CodingSelect("class", nil, optionsValueSet)
	}
	return CodingSelect("class", &resource.ClassHistory[numClassHistory].Class, optionsValueSet)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Participant[numParticipant].Type[0], optionsValueSet)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet)
}
func (resource *Encounter) T_HospitalizationAdmitSource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("admitSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("admitSource", resource.Hospitalization.AdmitSource, optionsValueSet)
}
func (resource *Encounter) T_HospitalizationReAdmission(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reAdmission", resource.Hospitalization.ReAdmission, optionsValueSet)
}
func (resource *Encounter) T_HospitalizationDietPreference(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("dietPreference", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dietPreference", &resource.Hospitalization.DietPreference[0], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationSpecialCourtesy(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialCourtesy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialCourtesy", &resource.Hospitalization.SpecialCourtesy[0], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationSpecialArrangement(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialArrangement", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialArrangement", &resource.Hospitalization.SpecialArrangement[0], optionsValueSet)
}
func (resource *Encounter) T_HospitalizationDischargeDisposition(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("dischargeDisposition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dischargeDisposition", resource.Hospitalization.DischargeDisposition, optionsValueSet)
}
func (resource *Encounter) T_LocationStatus(numLocation int) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil && len(resource.Location) >= numLocation {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Location[numLocation].Status, optionsValueSet)
}
func (resource *Encounter) T_LocationPhysicalType(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Location) >= numLocation {
		return CodeableConceptSelect("physicalType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("physicalType", resource.Location[numLocation].PhysicalType, optionsValueSet)
}
