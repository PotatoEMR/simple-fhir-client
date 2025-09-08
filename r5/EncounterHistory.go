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
	PlannedStartDate  *time.Time                 `json:"plannedStartDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PlannedEndDate    *time.Time                 `json:"plannedEndDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r EncounterHistory) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EncounterHistory/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EncounterHistory"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EncounterHistory) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEncounter_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_Class(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Class", &resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_SubjectStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubjectStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectStatus", resource.SubjectStatus, optionsValueSet, htmlAttrs)
}
func (resource *EncounterHistory) T_PlannedStartDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PlannedStartDate", nil, htmlAttrs)
	}
	return DateTimeInput("PlannedStartDate", resource.PlannedStartDate, htmlAttrs)
}
func (resource *EncounterHistory) T_PlannedEndDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PlannedEndDate", nil, htmlAttrs)
	}
	return DateTimeInput("PlannedEndDate", resource.PlannedEndDate, htmlAttrs)
}
func (resource *EncounterHistory) T_LocationForm(numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("Location["+strconv.Itoa(numLocation)+"]Form", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Location["+strconv.Itoa(numLocation)+"]Form", resource.Location[numLocation].Form, optionsValueSet, htmlAttrs)
}
