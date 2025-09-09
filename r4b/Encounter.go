package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
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

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
type EncounterStatusHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
type EncounterClassHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Class             Coding      `json:"class"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
type EncounterParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Individual        *Reference        `json:"individual,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
type EncounterDiagnosis struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Condition         Reference        `json:"condition"`
	Use               *CodeableConcept `json:"use,omitempty"`
	Rank              *int             `json:"rank,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
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

// http://hl7.org/fhir/r4b/StructureDefinition/Encounter
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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Class(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("class", &resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ServiceType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("serviceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceType", resource.ServiceType, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ClassHistoryClass(numClassHistory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassHistory >= len(resource.ClassHistory) {
		return CodingSelect("classHistory["+strconv.Itoa(numClassHistory)+"].class", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("classHistory["+strconv.Itoa(numClassHistory)+"].class", &resource.ClassHistory[numClassHistory].Class, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_DiagnosisRank(numDiagnosis int, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", resource.Diagnosis[numDiagnosis].Rank, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationAdmitSource(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.admitSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.admitSource", resource.Hospitalization.AdmitSource, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationReAdmission(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.reAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.reAdmission", resource.Hospitalization.ReAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDietPreference(numDietPreference int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDietPreference >= len(resource.Hospitalization.DietPreference) {
		return CodeableConceptSelect("hospitalization.dietPreference["+strconv.Itoa(numDietPreference)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.dietPreference["+strconv.Itoa(numDietPreference)+"]", &resource.Hospitalization.DietPreference[numDietPreference], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialCourtesy(numSpecialCourtesy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialCourtesy >= len(resource.Hospitalization.SpecialCourtesy) {
		return CodeableConceptSelect("hospitalization.specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.specialCourtesy["+strconv.Itoa(numSpecialCourtesy)+"]", &resource.Hospitalization.SpecialCourtesy[numSpecialCourtesy], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationSpecialArrangement(numSpecialArrangement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialArrangement >= len(resource.Hospitalization.SpecialArrangement) {
		return CodeableConceptSelect("hospitalization.specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.specialArrangement["+strconv.Itoa(numSpecialArrangement)+"]", &resource.Hospitalization.SpecialArrangement[numSpecialArrangement], optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_HospitalizationDischargeDisposition(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("hospitalization.dischargeDisposition", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("hospitalization.dischargeDisposition", resource.Hospitalization.DischargeDisposition, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationStatus(numLocation int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_location_status

	if resource == nil || numLocation >= len(resource.Location) {
		return CodeSelect("location["+strconv.Itoa(numLocation)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("location["+strconv.Itoa(numLocation)+"].status", resource.Location[numLocation].Status, optionsValueSet, htmlAttrs)
}
func (resource *Encounter) T_LocationPhysicalType(numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].physicalType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"].physicalType", resource.Location[numLocation].PhysicalType, optionsValueSet, htmlAttrs)
}
