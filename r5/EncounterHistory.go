package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *EncounterHistory) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EncounterHistory.Id", nil)
	}
	return StringInput("EncounterHistory.Id", resource.Id)
}
func (resource *EncounterHistory) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EncounterHistory.ImplicitRules", nil)
	}
	return StringInput("EncounterHistory.ImplicitRules", resource.ImplicitRules)
}
func (resource *EncounterHistory) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EncounterHistory.Language", nil, optionsValueSet)
	}
	return CodeSelect("EncounterHistory.Language", resource.Language, optionsValueSet)
}
func (resource *EncounterHistory) T_Status() templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("EncounterHistory.Status", nil, optionsValueSet)
	}
	return CodeSelect("EncounterHistory.Status", &resource.Status, optionsValueSet)
}
func (resource *EncounterHistory) T_Class(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("EncounterHistory.Class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EncounterHistory.Class", &resource.Class, optionsValueSet)
}
func (resource *EncounterHistory) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("EncounterHistory.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EncounterHistory.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *EncounterHistory) T_SubjectStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("EncounterHistory.SubjectStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EncounterHistory.SubjectStatus", resource.SubjectStatus, optionsValueSet)
}
func (resource *EncounterHistory) T_PlannedStartDate() templ.Component {

	if resource == nil {
		return StringInput("EncounterHistory.PlannedStartDate", nil)
	}
	return StringInput("EncounterHistory.PlannedStartDate", resource.PlannedStartDate)
}
func (resource *EncounterHistory) T_PlannedEndDate() templ.Component {

	if resource == nil {
		return StringInput("EncounterHistory.PlannedEndDate", nil)
	}
	return StringInput("EncounterHistory.PlannedEndDate", resource.PlannedEndDate)
}
func (resource *EncounterHistory) T_LocationId(numLocation int) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return StringInput("EncounterHistory.Location["+strconv.Itoa(numLocation)+"].Id", nil)
	}
	return StringInput("EncounterHistory.Location["+strconv.Itoa(numLocation)+"].Id", resource.Location[numLocation].Id)
}
func (resource *EncounterHistory) T_LocationForm(numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Location) >= numLocation {
		return CodeableConceptSelect("EncounterHistory.Location["+strconv.Itoa(numLocation)+"].Form", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EncounterHistory.Location["+strconv.Itoa(numLocation)+"].Form", resource.Location[numLocation].Form, optionsValueSet)
}
