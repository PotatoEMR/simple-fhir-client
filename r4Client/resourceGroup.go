package r4Client

import r4 "github.com/PotatoEMR/simple-fhir-client/r4"
import "errors"

type ResourceGroup struct {
	Accounts                           []*r4.Account
	ActivityDefinitions                []*r4.ActivityDefinition
	AdverseEvents                      []*r4.AdverseEvent
	AllergyIntolerances                []*r4.AllergyIntolerance
	Appointments                       []*r4.Appointment
	AppointmentResponses               []*r4.AppointmentResponse
	AuditEvents                        []*r4.AuditEvent
	Basics                             []*r4.Basic
	Binarys                            []*r4.Binary
	BiologicallyDerivedProducts        []*r4.BiologicallyDerivedProduct
	BodyStructures                     []*r4.BodyStructure
	CapabilityStatements               []*r4.CapabilityStatement
	CarePlans                          []*r4.CarePlan
	CareTeams                          []*r4.CareTeam
	CatalogEntrys                      []*r4.CatalogEntry
	ChargeItems                        []*r4.ChargeItem
	ChargeItemDefinitions              []*r4.ChargeItemDefinition
	Claims                             []*r4.Claim
	ClaimResponses                     []*r4.ClaimResponse
	ClinicalImpressions                []*r4.ClinicalImpression
	CodeSystems                        []*r4.CodeSystem
	Communications                     []*r4.Communication
	CommunicationRequests              []*r4.CommunicationRequest
	CompartmentDefinitions             []*r4.CompartmentDefinition
	Compositions                       []*r4.Composition
	ConceptMaps                        []*r4.ConceptMap
	Conditions                         []*r4.Condition
	Consents                           []*r4.Consent
	Contracts                          []*r4.Contract
	Coverages                          []*r4.Coverage
	CoverageEligibilityRequests        []*r4.CoverageEligibilityRequest
	CoverageEligibilityResponses       []*r4.CoverageEligibilityResponse
	DetectedIssues                     []*r4.DetectedIssue
	Devices                            []*r4.Device
	DeviceDefinitions                  []*r4.DeviceDefinition
	DeviceMetrics                      []*r4.DeviceMetric
	DeviceRequests                     []*r4.DeviceRequest
	DeviceUseStatements                []*r4.DeviceUseStatement
	DiagnosticReports                  []*r4.DiagnosticReport
	DocumentManifests                  []*r4.DocumentManifest
	DocumentReferences                 []*r4.DocumentReference
	EffectEvidenceSynthesiss           []*r4.EffectEvidenceSynthesis
	Encounters                         []*r4.Encounter
	Endpoints                          []*r4.Endpoint
	EnrollmentRequests                 []*r4.EnrollmentRequest
	EnrollmentResponses                []*r4.EnrollmentResponse
	EpisodeOfCares                     []*r4.EpisodeOfCare
	EventDefinitions                   []*r4.EventDefinition
	Evidences                          []*r4.Evidence
	EvidenceVariables                  []*r4.EvidenceVariable
	ExampleScenarios                   []*r4.ExampleScenario
	ExplanationOfBenefits              []*r4.ExplanationOfBenefit
	FamilyMemberHistorys               []*r4.FamilyMemberHistory
	Flags                              []*r4.Flag
	Goals                              []*r4.Goal
	GraphDefinitions                   []*r4.GraphDefinition
	Groups                             []*r4.Group
	GuidanceResponses                  []*r4.GuidanceResponse
	HealthcareServices                 []*r4.HealthcareService
	ImagingStudys                      []*r4.ImagingStudy
	Immunizations                      []*r4.Immunization
	ImmunizationEvaluations            []*r4.ImmunizationEvaluation
	ImmunizationRecommendations        []*r4.ImmunizationRecommendation
	ImplementationGuides               []*r4.ImplementationGuide
	InsurancePlans                     []*r4.InsurancePlan
	Invoices                           []*r4.Invoice
	Librarys                           []*r4.Library
	Linkages                           []*r4.Linkage
	Lists                              []*r4.List
	Locations                          []*r4.Location
	Measures                           []*r4.Measure
	MeasureReports                     []*r4.MeasureReport
	Medias                             []*r4.Media
	Medications                        []*r4.Medication
	MedicationAdministrations          []*r4.MedicationAdministration
	MedicationDispenses                []*r4.MedicationDispense
	MedicationKnowledges               []*r4.MedicationKnowledge
	MedicationRequests                 []*r4.MedicationRequest
	MedicationStatements               []*r4.MedicationStatement
	MedicinalProducts                  []*r4.MedicinalProduct
	MedicinalProductAuthorizations     []*r4.MedicinalProductAuthorization
	MedicinalProductContraindications  []*r4.MedicinalProductContraindication
	MedicinalProductIndications        []*r4.MedicinalProductIndication
	MedicinalProductIngredients        []*r4.MedicinalProductIngredient
	MedicinalProductInteractions       []*r4.MedicinalProductInteraction
	MedicinalProductManufactureds      []*r4.MedicinalProductManufactured
	MedicinalProductPackageds          []*r4.MedicinalProductPackaged
	MedicinalProductPharmaceuticals    []*r4.MedicinalProductPharmaceutical
	MedicinalProductUndesirableEffects []*r4.MedicinalProductUndesirableEffect
	MessageDefinitions                 []*r4.MessageDefinition
	MessageHeaders                     []*r4.MessageHeader
	MolecularSequences                 []*r4.MolecularSequence
	NamingSystems                      []*r4.NamingSystem
	NutritionOrders                    []*r4.NutritionOrder
	Observations                       []*r4.Observation
	ObservationDefinitions             []*r4.ObservationDefinition
	OperationDefinitions               []*r4.OperationDefinition
	OperationOutcomes                  []*r4.OperationOutcome
	Organizations                      []*r4.Organization
	OrganizationAffiliations           []*r4.OrganizationAffiliation
	Patients                           []*r4.Patient
	PaymentNotices                     []*r4.PaymentNotice
	PaymentReconciliations             []*r4.PaymentReconciliation
	Persons                            []*r4.Person
	PlanDefinitions                    []*r4.PlanDefinition
	Practitioners                      []*r4.Practitioner
	PractitionerRoles                  []*r4.PractitionerRole
	Procedures                         []*r4.Procedure
	Provenances                        []*r4.Provenance
	Questionnaires                     []*r4.Questionnaire
	QuestionnaireResponses             []*r4.QuestionnaireResponse
	RelatedPersons                     []*r4.RelatedPerson
	RequestGroups                      []*r4.RequestGroup
	ResearchDefinitions                []*r4.ResearchDefinition
	ResearchElementDefinitions         []*r4.ResearchElementDefinition
	ResearchStudys                     []*r4.ResearchStudy
	ResearchSubjects                   []*r4.ResearchSubject
	RiskAssessments                    []*r4.RiskAssessment
	RiskEvidenceSynthesiss             []*r4.RiskEvidenceSynthesis
	Schedules                          []*r4.Schedule
	SearchParameters                   []*r4.SearchParameter
	ServiceRequests                    []*r4.ServiceRequest
	Slots                              []*r4.Slot
	Specimens                          []*r4.Specimen
	SpecimenDefinitions                []*r4.SpecimenDefinition
	StructureDefinitions               []*r4.StructureDefinition
	StructureMaps                      []*r4.StructureMap
	Subscriptions                      []*r4.Subscription
	Substances                         []*r4.Substance
	SubstanceNucleicAcids              []*r4.SubstanceNucleicAcid
	SubstancePolymers                  []*r4.SubstancePolymer
	SubstanceProteins                  []*r4.SubstanceProtein
	SubstanceReferenceInformations     []*r4.SubstanceReferenceInformation
	SubstanceSourceMaterials           []*r4.SubstanceSourceMaterial
	SubstanceSpecifications            []*r4.SubstanceSpecification
	SupplyDeliverys                    []*r4.SupplyDelivery
	SupplyRequests                     []*r4.SupplyRequest
	Tasks                              []*r4.Task
	TerminologyCapabilitiess           []*r4.TerminologyCapabilities
	TestReports                        []*r4.TestReport
	TestScripts                        []*r4.TestScript
	ValueSets                          []*r4.ValueSet
	VerificationResults                []*r4.VerificationResult
	VisionPrescriptions                []*r4.VisionPrescription
}

func BundleToGroup(bundle *r4.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r4.Account:
			grp.Accounts = append(grp.Accounts, res)
		case *r4.ActivityDefinition:
			grp.ActivityDefinitions = append(grp.ActivityDefinitions, res)
		case *r4.AdverseEvent:
			grp.AdverseEvents = append(grp.AdverseEvents, res)
		case *r4.AllergyIntolerance:
			grp.AllergyIntolerances = append(grp.AllergyIntolerances, res)
		case *r4.Appointment:
			grp.Appointments = append(grp.Appointments, res)
		case *r4.AppointmentResponse:
			grp.AppointmentResponses = append(grp.AppointmentResponses, res)
		case *r4.AuditEvent:
			grp.AuditEvents = append(grp.AuditEvents, res)
		case *r4.Basic:
			grp.Basics = append(grp.Basics, res)
		case *r4.Binary:
			grp.Binarys = append(grp.Binarys, res)
		case *r4.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProducts = append(grp.BiologicallyDerivedProducts, res)
		case *r4.BodyStructure:
			grp.BodyStructures = append(grp.BodyStructures, res)
		case *r4.CapabilityStatement:
			grp.CapabilityStatements = append(grp.CapabilityStatements, res)
		case *r4.CarePlan:
			grp.CarePlans = append(grp.CarePlans, res)
		case *r4.CareTeam:
			grp.CareTeams = append(grp.CareTeams, res)
		case *r4.CatalogEntry:
			grp.CatalogEntrys = append(grp.CatalogEntrys, res)
		case *r4.ChargeItem:
			grp.ChargeItems = append(grp.ChargeItems, res)
		case *r4.ChargeItemDefinition:
			grp.ChargeItemDefinitions = append(grp.ChargeItemDefinitions, res)
		case *r4.Claim:
			grp.Claims = append(grp.Claims, res)
		case *r4.ClaimResponse:
			grp.ClaimResponses = append(grp.ClaimResponses, res)
		case *r4.ClinicalImpression:
			grp.ClinicalImpressions = append(grp.ClinicalImpressions, res)
		case *r4.CodeSystem:
			grp.CodeSystems = append(grp.CodeSystems, res)
		case *r4.Communication:
			grp.Communications = append(grp.Communications, res)
		case *r4.CommunicationRequest:
			grp.CommunicationRequests = append(grp.CommunicationRequests, res)
		case *r4.CompartmentDefinition:
			grp.CompartmentDefinitions = append(grp.CompartmentDefinitions, res)
		case *r4.Composition:
			grp.Compositions = append(grp.Compositions, res)
		case *r4.ConceptMap:
			grp.ConceptMaps = append(grp.ConceptMaps, res)
		case *r4.Condition:
			grp.Conditions = append(grp.Conditions, res)
		case *r4.Consent:
			grp.Consents = append(grp.Consents, res)
		case *r4.Contract:
			grp.Contracts = append(grp.Contracts, res)
		case *r4.Coverage:
			grp.Coverages = append(grp.Coverages, res)
		case *r4.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequests = append(grp.CoverageEligibilityRequests, res)
		case *r4.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponses = append(grp.CoverageEligibilityResponses, res)
		case *r4.DetectedIssue:
			grp.DetectedIssues = append(grp.DetectedIssues, res)
		case *r4.Device:
			grp.Devices = append(grp.Devices, res)
		case *r4.DeviceDefinition:
			grp.DeviceDefinitions = append(grp.DeviceDefinitions, res)
		case *r4.DeviceMetric:
			grp.DeviceMetrics = append(grp.DeviceMetrics, res)
		case *r4.DeviceRequest:
			grp.DeviceRequests = append(grp.DeviceRequests, res)
		case *r4.DeviceUseStatement:
			grp.DeviceUseStatements = append(grp.DeviceUseStatements, res)
		case *r4.DiagnosticReport:
			grp.DiagnosticReports = append(grp.DiagnosticReports, res)
		case *r4.DocumentManifest:
			grp.DocumentManifests = append(grp.DocumentManifests, res)
		case *r4.DocumentReference:
			grp.DocumentReferences = append(grp.DocumentReferences, res)
		case *r4.EffectEvidenceSynthesis:
			grp.EffectEvidenceSynthesiss = append(grp.EffectEvidenceSynthesiss, res)
		case *r4.Encounter:
			grp.Encounters = append(grp.Encounters, res)
		case *r4.Endpoint:
			grp.Endpoints = append(grp.Endpoints, res)
		case *r4.EnrollmentRequest:
			grp.EnrollmentRequests = append(grp.EnrollmentRequests, res)
		case *r4.EnrollmentResponse:
			grp.EnrollmentResponses = append(grp.EnrollmentResponses, res)
		case *r4.EpisodeOfCare:
			grp.EpisodeOfCares = append(grp.EpisodeOfCares, res)
		case *r4.EventDefinition:
			grp.EventDefinitions = append(grp.EventDefinitions, res)
		case *r4.Evidence:
			grp.Evidences = append(grp.Evidences, res)
		case *r4.EvidenceVariable:
			grp.EvidenceVariables = append(grp.EvidenceVariables, res)
		case *r4.ExampleScenario:
			grp.ExampleScenarios = append(grp.ExampleScenarios, res)
		case *r4.ExplanationOfBenefit:
			grp.ExplanationOfBenefits = append(grp.ExplanationOfBenefits, res)
		case *r4.FamilyMemberHistory:
			grp.FamilyMemberHistorys = append(grp.FamilyMemberHistorys, res)
		case *r4.Flag:
			grp.Flags = append(grp.Flags, res)
		case *r4.Goal:
			grp.Goals = append(grp.Goals, res)
		case *r4.GraphDefinition:
			grp.GraphDefinitions = append(grp.GraphDefinitions, res)
		case *r4.Group:
			grp.Groups = append(grp.Groups, res)
		case *r4.GuidanceResponse:
			grp.GuidanceResponses = append(grp.GuidanceResponses, res)
		case *r4.HealthcareService:
			grp.HealthcareServices = append(grp.HealthcareServices, res)
		case *r4.ImagingStudy:
			grp.ImagingStudys = append(grp.ImagingStudys, res)
		case *r4.Immunization:
			grp.Immunizations = append(grp.Immunizations, res)
		case *r4.ImmunizationEvaluation:
			grp.ImmunizationEvaluations = append(grp.ImmunizationEvaluations, res)
		case *r4.ImmunizationRecommendation:
			grp.ImmunizationRecommendations = append(grp.ImmunizationRecommendations, res)
		case *r4.ImplementationGuide:
			grp.ImplementationGuides = append(grp.ImplementationGuides, res)
		case *r4.InsurancePlan:
			grp.InsurancePlans = append(grp.InsurancePlans, res)
		case *r4.Invoice:
			grp.Invoices = append(grp.Invoices, res)
		case *r4.Library:
			grp.Librarys = append(grp.Librarys, res)
		case *r4.Linkage:
			grp.Linkages = append(grp.Linkages, res)
		case *r4.List:
			grp.Lists = append(grp.Lists, res)
		case *r4.Location:
			grp.Locations = append(grp.Locations, res)
		case *r4.Measure:
			grp.Measures = append(grp.Measures, res)
		case *r4.MeasureReport:
			grp.MeasureReports = append(grp.MeasureReports, res)
		case *r4.Media:
			grp.Medias = append(grp.Medias, res)
		case *r4.Medication:
			grp.Medications = append(grp.Medications, res)
		case *r4.MedicationAdministration:
			grp.MedicationAdministrations = append(grp.MedicationAdministrations, res)
		case *r4.MedicationDispense:
			grp.MedicationDispenses = append(grp.MedicationDispenses, res)
		case *r4.MedicationKnowledge:
			grp.MedicationKnowledges = append(grp.MedicationKnowledges, res)
		case *r4.MedicationRequest:
			grp.MedicationRequests = append(grp.MedicationRequests, res)
		case *r4.MedicationStatement:
			grp.MedicationStatements = append(grp.MedicationStatements, res)
		case *r4.MedicinalProduct:
			grp.MedicinalProducts = append(grp.MedicinalProducts, res)
		case *r4.MedicinalProductAuthorization:
			grp.MedicinalProductAuthorizations = append(grp.MedicinalProductAuthorizations, res)
		case *r4.MedicinalProductContraindication:
			grp.MedicinalProductContraindications = append(grp.MedicinalProductContraindications, res)
		case *r4.MedicinalProductIndication:
			grp.MedicinalProductIndications = append(grp.MedicinalProductIndications, res)
		case *r4.MedicinalProductIngredient:
			grp.MedicinalProductIngredients = append(grp.MedicinalProductIngredients, res)
		case *r4.MedicinalProductInteraction:
			grp.MedicinalProductInteractions = append(grp.MedicinalProductInteractions, res)
		case *r4.MedicinalProductManufactured:
			grp.MedicinalProductManufactureds = append(grp.MedicinalProductManufactureds, res)
		case *r4.MedicinalProductPackaged:
			grp.MedicinalProductPackageds = append(grp.MedicinalProductPackageds, res)
		case *r4.MedicinalProductPharmaceutical:
			grp.MedicinalProductPharmaceuticals = append(grp.MedicinalProductPharmaceuticals, res)
		case *r4.MedicinalProductUndesirableEffect:
			grp.MedicinalProductUndesirableEffects = append(grp.MedicinalProductUndesirableEffects, res)
		case *r4.MessageDefinition:
			grp.MessageDefinitions = append(grp.MessageDefinitions, res)
		case *r4.MessageHeader:
			grp.MessageHeaders = append(grp.MessageHeaders, res)
		case *r4.MolecularSequence:
			grp.MolecularSequences = append(grp.MolecularSequences, res)
		case *r4.NamingSystem:
			grp.NamingSystems = append(grp.NamingSystems, res)
		case *r4.NutritionOrder:
			grp.NutritionOrders = append(grp.NutritionOrders, res)
		case *r4.Observation:
			grp.Observations = append(grp.Observations, res)
		case *r4.ObservationDefinition:
			grp.ObservationDefinitions = append(grp.ObservationDefinitions, res)
		case *r4.OperationDefinition:
			grp.OperationDefinitions = append(grp.OperationDefinitions, res)
		case *r4.OperationOutcome:
			grp.OperationOutcomes = append(grp.OperationOutcomes, res)
		case *r4.Organization:
			grp.Organizations = append(grp.Organizations, res)
		case *r4.OrganizationAffiliation:
			grp.OrganizationAffiliations = append(grp.OrganizationAffiliations, res)
		case *r4.Patient:
			grp.Patients = append(grp.Patients, res)
		case *r4.PaymentNotice:
			grp.PaymentNotices = append(grp.PaymentNotices, res)
		case *r4.PaymentReconciliation:
			grp.PaymentReconciliations = append(grp.PaymentReconciliations, res)
		case *r4.Person:
			grp.Persons = append(grp.Persons, res)
		case *r4.PlanDefinition:
			grp.PlanDefinitions = append(grp.PlanDefinitions, res)
		case *r4.Practitioner:
			grp.Practitioners = append(grp.Practitioners, res)
		case *r4.PractitionerRole:
			grp.PractitionerRoles = append(grp.PractitionerRoles, res)
		case *r4.Procedure:
			grp.Procedures = append(grp.Procedures, res)
		case *r4.Provenance:
			grp.Provenances = append(grp.Provenances, res)
		case *r4.Questionnaire:
			grp.Questionnaires = append(grp.Questionnaires, res)
		case *r4.QuestionnaireResponse:
			grp.QuestionnaireResponses = append(grp.QuestionnaireResponses, res)
		case *r4.RelatedPerson:
			grp.RelatedPersons = append(grp.RelatedPersons, res)
		case *r4.RequestGroup:
			grp.RequestGroups = append(grp.RequestGroups, res)
		case *r4.ResearchDefinition:
			grp.ResearchDefinitions = append(grp.ResearchDefinitions, res)
		case *r4.ResearchElementDefinition:
			grp.ResearchElementDefinitions = append(grp.ResearchElementDefinitions, res)
		case *r4.ResearchStudy:
			grp.ResearchStudys = append(grp.ResearchStudys, res)
		case *r4.ResearchSubject:
			grp.ResearchSubjects = append(grp.ResearchSubjects, res)
		case *r4.RiskAssessment:
			grp.RiskAssessments = append(grp.RiskAssessments, res)
		case *r4.RiskEvidenceSynthesis:
			grp.RiskEvidenceSynthesiss = append(grp.RiskEvidenceSynthesiss, res)
		case *r4.Schedule:
			grp.Schedules = append(grp.Schedules, res)
		case *r4.SearchParameter:
			grp.SearchParameters = append(grp.SearchParameters, res)
		case *r4.ServiceRequest:
			grp.ServiceRequests = append(grp.ServiceRequests, res)
		case *r4.Slot:
			grp.Slots = append(grp.Slots, res)
		case *r4.Specimen:
			grp.Specimens = append(grp.Specimens, res)
		case *r4.SpecimenDefinition:
			grp.SpecimenDefinitions = append(grp.SpecimenDefinitions, res)
		case *r4.StructureDefinition:
			grp.StructureDefinitions = append(grp.StructureDefinitions, res)
		case *r4.StructureMap:
			grp.StructureMaps = append(grp.StructureMaps, res)
		case *r4.Subscription:
			grp.Subscriptions = append(grp.Subscriptions, res)
		case *r4.Substance:
			grp.Substances = append(grp.Substances, res)
		case *r4.SubstanceNucleicAcid:
			grp.SubstanceNucleicAcids = append(grp.SubstanceNucleicAcids, res)
		case *r4.SubstancePolymer:
			grp.SubstancePolymers = append(grp.SubstancePolymers, res)
		case *r4.SubstanceProtein:
			grp.SubstanceProteins = append(grp.SubstanceProteins, res)
		case *r4.SubstanceReferenceInformation:
			grp.SubstanceReferenceInformations = append(grp.SubstanceReferenceInformations, res)
		case *r4.SubstanceSourceMaterial:
			grp.SubstanceSourceMaterials = append(grp.SubstanceSourceMaterials, res)
		case *r4.SubstanceSpecification:
			grp.SubstanceSpecifications = append(grp.SubstanceSpecifications, res)
		case *r4.SupplyDelivery:
			grp.SupplyDeliverys = append(grp.SupplyDeliverys, res)
		case *r4.SupplyRequest:
			grp.SupplyRequests = append(grp.SupplyRequests, res)
		case *r4.Task:
			grp.Tasks = append(grp.Tasks, res)
		case *r4.TerminologyCapabilities:
			grp.TerminologyCapabilitiess = append(grp.TerminologyCapabilitiess, res)
		case *r4.TestReport:
			grp.TestReports = append(grp.TestReports, res)
		case *r4.TestScript:
			grp.TestScripts = append(grp.TestScripts, res)
		case *r4.ValueSet:
			grp.ValueSets = append(grp.ValueSets, res)
		case *r4.VerificationResult:
			grp.VerificationResults = append(grp.VerificationResults, res)
		case *r4.VisionPrescription:
			grp.VisionPrescriptions = append(grp.VisionPrescriptions, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
