// Note: Bundle is special case in gen_bundle.go, to handle entry resources with types
package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type Bundle struct {
	Id            *string       `json:"id,omitempty"`
	Meta          *Meta         `json:"meta,omitempty"`
	ImplicitRules *string       `json:"implicitRules,omitempty"`
	Language      *string       `json:"language,omitempty"`
	Identifier    *Identifier   `json:"identifier,omitempty"`
	Type          string        `json:"type"`
	Timestamp     *string       `json:"timestamp,omitempty"`
	Total         *int          `json:"total,omitempty"`
	Link          []BundleLink  `json:"link,omitempty"`
	Entry         []BundleEntry `json:"entry,omitempty"`
	Signature     *Signature    `json:"signature,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type BundleLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Relation          string      `json:"relation"`
	Url               string      `json:"url"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type BundleEntry struct {
	Id                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	FullUrl           *string              `json:"fullUrl,omitempty"`
	UntypedResource   *json.RawMessage     `json:"resource,omitempty"`
	Resource          any                  `json:"-"`
	Search            *BundleEntrySearch   `json:"search,omitempty"`
	Request           *BundleEntryRequest  `json:"request,omitempty"`
	Response          *BundleEntryResponse `json:"response,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type BundleEntrySearch struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              *string     `json:"mode,omitempty"`
	Score             *float64    `json:"score,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type BundleEntryRequest struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Method            string      `json:"method"`
	Url               string      `json:"url"`
	IfNoneMatch       *string     `json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `json:"ifMatch,omitempty"`
	IfNoneExist       *string     `json:"ifNoneExist,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Bundle
type BundleEntryResponse struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Location          *string     `json:"location,omitempty"`
	Etag              *string     `json:"etag,omitempty"`
	LastModified      *string     `json:"lastModified,omitempty"`
	Outcome           *Resource   `json:"outcome,omitempty"`
}

func (resource *Bundle) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSBundle_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Bundle) T_Timestamp(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Timestamp", nil, htmlAttrs)
	}
	return StringInput("Timestamp", resource.Timestamp, htmlAttrs)
}
func (resource *Bundle) T_Total(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Total", nil, htmlAttrs)
	}
	return IntInput("Total", resource.Total, htmlAttrs)
}
func (resource *Bundle) T_LinkRelation(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("Link["+strconv.Itoa(numLink)+"]Relation", nil, htmlAttrs)
	}
	return StringInput("Link["+strconv.Itoa(numLink)+"]Relation", &resource.Link[numLink].Relation, htmlAttrs)
}
func (resource *Bundle) T_LinkUrl(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("Link["+strconv.Itoa(numLink)+"]Url", nil, htmlAttrs)
	}
	return StringInput("Link["+strconv.Itoa(numLink)+"]Url", &resource.Link[numLink].Url, htmlAttrs)
}
func (resource *Bundle) T_EntryFullUrl(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]FullUrl", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]FullUrl", resource.Entry[numEntry].FullUrl, htmlAttrs)
}
func (resource *Bundle) T_EntrySearchMode(numEntry int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_entry_mode

	if resource == nil || numEntry >= len(resource.Entry) {
		return CodeSelect("Entry["+strconv.Itoa(numEntry)+"]Search.Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Entry["+strconv.Itoa(numEntry)+"]Search.Mode", resource.Entry[numEntry].Search.Mode, optionsValueSet, htmlAttrs)
}
func (resource *Bundle) T_EntrySearchScore(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return Float64Input("Entry["+strconv.Itoa(numEntry)+"]Search.Score", nil, htmlAttrs)
	}
	return Float64Input("Entry["+strconv.Itoa(numEntry)+"]Search.Score", resource.Entry[numEntry].Search.Score, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestMethod(numEntry int, htmlAttrs string) templ.Component {
	optionsValueSet := VSHttp_verb

	if resource == nil || numEntry >= len(resource.Entry) {
		return CodeSelect("Entry["+strconv.Itoa(numEntry)+"]Request.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Entry["+strconv.Itoa(numEntry)+"]Request.Method", &resource.Entry[numEntry].Request.Method, optionsValueSet, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestUrl(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.Url", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.Url", &resource.Entry[numEntry].Request.Url, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestIfNoneMatch(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfNoneMatch", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfNoneMatch", resource.Entry[numEntry].Request.IfNoneMatch, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestIfModifiedSince(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfModifiedSince", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfModifiedSince", resource.Entry[numEntry].Request.IfModifiedSince, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestIfMatch(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfMatch", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfMatch", resource.Entry[numEntry].Request.IfMatch, htmlAttrs)
}
func (resource *Bundle) T_EntryRequestIfNoneExist(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfNoneExist", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Request.IfNoneExist", resource.Entry[numEntry].Request.IfNoneExist, htmlAttrs)
}
func (resource *Bundle) T_EntryResponseStatus(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Status", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Status", &resource.Entry[numEntry].Response.Status, htmlAttrs)
}
func (resource *Bundle) T_EntryResponseLocation(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Location", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Location", resource.Entry[numEntry].Response.Location, htmlAttrs)
}
func (resource *Bundle) T_EntryResponseEtag(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Etag", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.Etag", resource.Entry[numEntry].Response.Etag, htmlAttrs)
}
func (resource *Bundle) T_EntryResponseLastModified(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.LastModified", nil, htmlAttrs)
	}
	return StringInput("Entry["+strconv.Itoa(numEntry)+"]Response.LastModified", resource.Entry[numEntry].Response.LastModified, htmlAttrs)
}

type OtherBundle Bundle

// per Lukas Schmierer, check resourceType string in raw resource json,
// then use string in switch to parse into correct resource,
// and user can use resource with type assertion
func (b *Bundle) UnmarshalJSON(data []byte) error {
	var tmp struct {
		OtherBundle
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = Bundle(tmp.OtherBundle)

	for i := range b.Entry {
		e := &b.Entry[i] //weird looking because need to modify entry at i, not copy of i
		if e.UntypedResource == nil {
			return errors.New("Bundle UnmarshalJSON got nil resource")
		}

		var peek struct {
			ResourceType string `json:"resourceType"`
		}
		if err := json.Unmarshal(*e.UntypedResource, &peek); err != nil {
			return errors.New("Bundle UnmarshalJSON failed to get resource type from bundle entry" + string(*e.UntypedResource))
		}

		switch peek.ResourceType {
		case "Account":
			var res Account
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ActivityDefinition":
			var res ActivityDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "AdverseEvent":
			var res AdverseEvent
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "AllergyIntolerance":
			var res AllergyIntolerance
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Appointment":
			var res Appointment
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "AppointmentResponse":
			var res AppointmentResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "AuditEvent":
			var res AuditEvent
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Basic":
			var res Basic
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Binary":
			var res Binary
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "BiologicallyDerivedProduct":
			var res BiologicallyDerivedProduct
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "BodyStructure":
			var res BodyStructure
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CapabilityStatement":
			var res CapabilityStatement
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CarePlan":
			var res CarePlan
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CareTeam":
			var res CareTeam
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CatalogEntry":
			var res CatalogEntry
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ChargeItem":
			var res ChargeItem
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ChargeItemDefinition":
			var res ChargeItemDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Claim":
			var res Claim
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ClaimResponse":
			var res ClaimResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ClinicalImpression":
			var res ClinicalImpression
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CodeSystem":
			var res CodeSystem
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Communication":
			var res Communication
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CommunicationRequest":
			var res CommunicationRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CompartmentDefinition":
			var res CompartmentDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Composition":
			var res Composition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ConceptMap":
			var res ConceptMap
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Condition":
			var res Condition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Consent":
			var res Consent
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Contract":
			var res Contract
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Coverage":
			var res Coverage
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CoverageEligibilityRequest":
			var res CoverageEligibilityRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "CoverageEligibilityResponse":
			var res CoverageEligibilityResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DetectedIssue":
			var res DetectedIssue
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Device":
			var res Device
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DeviceDefinition":
			var res DeviceDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DeviceMetric":
			var res DeviceMetric
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DeviceRequest":
			var res DeviceRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DeviceUseStatement":
			var res DeviceUseStatement
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DiagnosticReport":
			var res DiagnosticReport
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DocumentManifest":
			var res DocumentManifest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "DocumentReference":
			var res DocumentReference
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EffectEvidenceSynthesis":
			var res EffectEvidenceSynthesis
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Encounter":
			var res Encounter
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Endpoint":
			var res Endpoint
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EnrollmentRequest":
			var res EnrollmentRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EnrollmentResponse":
			var res EnrollmentResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EpisodeOfCare":
			var res EpisodeOfCare
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EventDefinition":
			var res EventDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Evidence":
			var res Evidence
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "EvidenceVariable":
			var res EvidenceVariable
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ExampleScenario":
			var res ExampleScenario
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ExplanationOfBenefit":
			var res ExplanationOfBenefit
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "FamilyMemberHistory":
			var res FamilyMemberHistory
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Flag":
			var res Flag
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Goal":
			var res Goal
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "GraphDefinition":
			var res GraphDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Group":
			var res Group
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "GuidanceResponse":
			var res GuidanceResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "HealthcareService":
			var res HealthcareService
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ImagingStudy":
			var res ImagingStudy
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Immunization":
			var res Immunization
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ImmunizationEvaluation":
			var res ImmunizationEvaluation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ImmunizationRecommendation":
			var res ImmunizationRecommendation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ImplementationGuide":
			var res ImplementationGuide
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "InsurancePlan":
			var res InsurancePlan
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Invoice":
			var res Invoice
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Library":
			var res Library
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Linkage":
			var res Linkage
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "List":
			var res List
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Location":
			var res Location
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Measure":
			var res Measure
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MeasureReport":
			var res MeasureReport
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Media":
			var res Media
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Medication":
			var res Medication
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicationAdministration":
			var res MedicationAdministration
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicationDispense":
			var res MedicationDispense
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicationKnowledge":
			var res MedicationKnowledge
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicationRequest":
			var res MedicationRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicationStatement":
			var res MedicationStatement
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProduct":
			var res MedicinalProduct
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductAuthorization":
			var res MedicinalProductAuthorization
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductContraindication":
			var res MedicinalProductContraindication
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductIndication":
			var res MedicinalProductIndication
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductIngredient":
			var res MedicinalProductIngredient
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductInteraction":
			var res MedicinalProductInteraction
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductManufactured":
			var res MedicinalProductManufactured
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductPackaged":
			var res MedicinalProductPackaged
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductPharmaceutical":
			var res MedicinalProductPharmaceutical
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MedicinalProductUndesirableEffect":
			var res MedicinalProductUndesirableEffect
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MessageDefinition":
			var res MessageDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MessageHeader":
			var res MessageHeader
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "MolecularSequence":
			var res MolecularSequence
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "NamingSystem":
			var res NamingSystem
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "NutritionOrder":
			var res NutritionOrder
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Observation":
			var res Observation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ObservationDefinition":
			var res ObservationDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "OperationDefinition":
			var res OperationDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "OperationOutcome":
			var res OperationOutcome
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Organization":
			var res Organization
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "OrganizationAffiliation":
			var res OrganizationAffiliation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Patient":
			var res Patient
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "PaymentNotice":
			var res PaymentNotice
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "PaymentReconciliation":
			var res PaymentReconciliation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Person":
			var res Person
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "PlanDefinition":
			var res PlanDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Practitioner":
			var res Practitioner
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "PractitionerRole":
			var res PractitionerRole
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Procedure":
			var res Procedure
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Provenance":
			var res Provenance
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Questionnaire":
			var res Questionnaire
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "QuestionnaireResponse":
			var res QuestionnaireResponse
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "RelatedPerson":
			var res RelatedPerson
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "RequestGroup":
			var res RequestGroup
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ResearchDefinition":
			var res ResearchDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ResearchElementDefinition":
			var res ResearchElementDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ResearchStudy":
			var res ResearchStudy
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ResearchSubject":
			var res ResearchSubject
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "RiskAssessment":
			var res RiskAssessment
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "RiskEvidenceSynthesis":
			var res RiskEvidenceSynthesis
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Schedule":
			var res Schedule
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SearchParameter":
			var res SearchParameter
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ServiceRequest":
			var res ServiceRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Slot":
			var res Slot
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Specimen":
			var res Specimen
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SpecimenDefinition":
			var res SpecimenDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "StructureDefinition":
			var res StructureDefinition
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "StructureMap":
			var res StructureMap
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Subscription":
			var res Subscription
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Substance":
			var res Substance
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstanceNucleicAcid":
			var res SubstanceNucleicAcid
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstancePolymer":
			var res SubstancePolymer
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstanceProtein":
			var res SubstanceProtein
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstanceReferenceInformation":
			var res SubstanceReferenceInformation
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstanceSourceMaterial":
			var res SubstanceSourceMaterial
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SubstanceSpecification":
			var res SubstanceSpecification
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SupplyDelivery":
			var res SupplyDelivery
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "SupplyRequest":
			var res SupplyRequest
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "Task":
			var res Task
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "TerminologyCapabilities":
			var res TerminologyCapabilities
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "TestReport":
			var res TestReport
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "TestScript":
			var res TestScript
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "ValueSet":
			var res ValueSet
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "VerificationResult":
			var res VerificationResult
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		case "VisionPrescription":
			var res VisionPrescription
			err := json.Unmarshal(*e.UntypedResource, &res)
			if err != nil {
				return err
			}
			e.Resource = &res
		default:
			return errors.New("Bundle UnmarshalJSON switch default, no resource type" + peek.ResourceType)
		}
	}
	return nil
}

// MarshalJSON copies Resource fields to UntypedResource before marshaling
func (b *Bundle) MarshalJSON() ([]byte, error) {
	tmp := struct {
		*OtherBundle
	}{
		OtherBundle: (*OtherBundle)(b),
	}

	// Convert Resource to UntypedResource for each entry
	for i := range tmp.Entry {
		e := &tmp.Entry[i]
		if e.Resource != nil {
			resourceBytes, err := json.Marshal(e.Resource)
			if err != nil {
				return nil, err
			}
			rawMsg := json.RawMessage(resourceBytes)
			e.UntypedResource = &rawMsg
		}
	}
	return json.Marshal(tmp.OtherBundle)
}
