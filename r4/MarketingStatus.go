package r4

import "time"

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

// http://hl7.org/fhir/r4/StructureDefinition/MarketingStatus
type MarketingStatus struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Status            CodeableConcept  `json:"status"`
	DateRange         Period           `json:"dateRange"`
	RestoreDate       *time.Time       `json:"restoreDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}
