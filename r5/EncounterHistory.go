package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistory struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Encounter         *Reference                 `json:"encounter,omitempty"`
	Identifier        []Identifier               `json:"identifier,omitempty"`
	Status            string                     `json:"status"`
	Class             CodeableConcept            `json:"class"`
	Type              []CodeableConcept          `json:"type,omitempty"`
	ServiceType       []CodeableReference        `json:"serviceType,omitempty"`
	Subject           *Reference                 `json:"subject,omitempty"`
	SubjectStatus     *CodeableConcept           `json:"subjectStatus,omitempty"`
	ActualPeriod      *Period                    `json:"actualPeriod,omitempty"`
	PlannedStartDate  *string                    `json:"plannedStartDate,omitempty"`
	PlannedEndDate    *string                    `json:"plannedEndDate,omitempty"`
	Length            *Duration                  `json:"length,omitempty"`
	Location          []EncounterHistoryLocation `json:"location,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EncounterHistory
type EncounterHistoryLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Location          Reference        `json:"location"`
	Form              *CodeableConcept `json:"form,omitempty"`
}

type OtherEncounterHistory EncounterHistory

// on convert struct to json, automatically add resourceType=EncounterHistory
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounterHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounterHistory: OtherEncounterHistory(r),
		ResourceType:          "EncounterHistory",
	})
}

func (resource *EncounterHistory) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *EncounterHistory) T_Status() templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *EncounterHistory) T_Class(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("class", &resource.Class, optionsValueSet)
}
func (resource *EncounterHistory) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *EncounterHistory) T_SubjectStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subjectStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subjectStatus", resource.SubjectStatus, optionsValueSet)
}
func (resource *EncounterHistory) T_LocationForm(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Location) >= numLocation {
		return CodeableConceptSelect("form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("form", resource.Location[numLocation].Form, optionsValueSet)
}
