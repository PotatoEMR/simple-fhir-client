package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/Flag
type Flag struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Status            string            `json:"status"`
	Category          []CodeableConcept `json:"category,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Subject           Reference         `json:"subject"`
	Period            *Period           `json:"period,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
}

type OtherFlag Flag

// on convert struct to json, automatically add resourceType=Flag
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

func (resource *Flag) FlagLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Flag) FlagStatus() templ.Component {
	optionsValueSet := VSFlag_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
