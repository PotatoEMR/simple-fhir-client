package r4b

import "time"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4b/StructureDefinition/MarketingStatus
type MarketingStatus struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           *CodeableConcept `json:"country,omitempty"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Status            CodeableConcept  `json:"status"`
	DateRange         *Period          `json:"dateRange,omitempty"`
	RestoreDate       *time.Time       `json:"restoreDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}
