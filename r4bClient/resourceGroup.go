package r4bClient

import r4b "github.com/PotatoEMR/simple-fhir-client/r4b"
import "errors"

type ResourceGroup struct {
	Accounts                        []*r4b.Account
	ActivityDefinitions             []*r4b.ActivityDefinition
	AdministrableProductDefinitions []*r4b.AdministrableProductDefinition
	AdverseEvents                   []*r4b.AdverseEvent
	AllergyIntolerances             []*r4b.AllergyIntolerance
	Appointments                    []*r4b.Appointment
	AppointmentResponses            []*r4b.AppointmentResponse
	AuditEvents                     []*r4b.AuditEvent
	Basics                          []*r4b.Basic
	Binarys                         []*r4b.Binary
	BiologicallyDerivedProducts     []*r4b.BiologicallyDerivedProduct
	BodyStructures                  []*r4b.BodyStructure
	CapabilityStatements            []*r4b.CapabilityStatement
	CarePlans                       []*r4b.CarePlan
	CareTeams                       []*r4b.CareTeam
	CatalogEntrys                   []*r4b.CatalogEntry
	ChargeItems                     []*r4b.ChargeItem
	ChargeItemDefinitions           []*r4b.ChargeItemDefinition
	Citations                       []*r4b.Citation
	Claims                          []*r4b.Claim
	ClaimResponses                  []*r4b.ClaimResponse
	ClinicalImpressions             []*r4b.ClinicalImpression
	ClinicalUseDefinitions          []*r4b.ClinicalUseDefinition
	CodeSystems                     []*r4b.CodeSystem
	Communications                  []*r4b.Communication
	CommunicationRequests           []*r4b.CommunicationRequest
	CompartmentDefinitions          []*r4b.CompartmentDefinition
	Compositions                    []*r4b.Composition
	ConceptMaps                     []*r4b.ConceptMap
	Conditions                      []*r4b.Condition
	Consents                        []*r4b.Consent
	Contracts                       []*r4b.Contract
	Coverages                       []*r4b.Coverage
	CoverageEligibilityRequests     []*r4b.CoverageEligibilityRequest
	CoverageEligibilityResponses    []*r4b.CoverageEligibilityResponse
	DetectedIssues                  []*r4b.DetectedIssue
	Devices                         []*r4b.Device
	DeviceDefinitions               []*r4b.DeviceDefinition
	DeviceMetrics                   []*r4b.DeviceMetric
	DeviceRequests                  []*r4b.DeviceRequest
	DeviceUseStatements             []*r4b.DeviceUseStatement
	DiagnosticReports               []*r4b.DiagnosticReport
	DocumentManifests               []*r4b.DocumentManifest
	DocumentReferences              []*r4b.DocumentReference
	Encounters                      []*r4b.Encounter
	Endpoints                       []*r4b.Endpoint
	EnrollmentRequests              []*r4b.EnrollmentRequest
	EnrollmentResponses             []*r4b.EnrollmentResponse
	EpisodeOfCares                  []*r4b.EpisodeOfCare
	EventDefinitions                []*r4b.EventDefinition
	Evidences                       []*r4b.Evidence
	EvidenceReports                 []*r4b.EvidenceReport
	EvidenceVariables               []*r4b.EvidenceVariable
	ExampleScenarios                []*r4b.ExampleScenario
	ExplanationOfBenefits           []*r4b.ExplanationOfBenefit
	FamilyMemberHistorys            []*r4b.FamilyMemberHistory
	Flags                           []*r4b.Flag
	Goals                           []*r4b.Goal
	GraphDefinitions                []*r4b.GraphDefinition
	Groups                          []*r4b.Group
	GuidanceResponses               []*r4b.GuidanceResponse
	HealthcareServices              []*r4b.HealthcareService
	ImagingStudys                   []*r4b.ImagingStudy
	Immunizations                   []*r4b.Immunization
	ImmunizationEvaluations         []*r4b.ImmunizationEvaluation
	ImmunizationRecommendations     []*r4b.ImmunizationRecommendation
	ImplementationGuides            []*r4b.ImplementationGuide
	Ingredients                     []*r4b.Ingredient
	InsurancePlans                  []*r4b.InsurancePlan
	Invoices                        []*r4b.Invoice
	Librarys                        []*r4b.Library
	Linkages                        []*r4b.Linkage
	Lists                           []*r4b.List
	Locations                       []*r4b.Location
	ManufacturedItemDefinitions     []*r4b.ManufacturedItemDefinition
	Measures                        []*r4b.Measure
	MeasureReports                  []*r4b.MeasureReport
	Medias                          []*r4b.Media
	Medications                     []*r4b.Medication
	MedicationAdministrations       []*r4b.MedicationAdministration
	MedicationDispenses             []*r4b.MedicationDispense
	MedicationKnowledges            []*r4b.MedicationKnowledge
	MedicationRequests              []*r4b.MedicationRequest
	MedicationStatements            []*r4b.MedicationStatement
	MedicinalProductDefinitions     []*r4b.MedicinalProductDefinition
	MessageDefinitions              []*r4b.MessageDefinition
	MessageHeaders                  []*r4b.MessageHeader
	MolecularSequences              []*r4b.MolecularSequence
	NamingSystems                   []*r4b.NamingSystem
	NutritionOrders                 []*r4b.NutritionOrder
	NutritionProducts               []*r4b.NutritionProduct
	Observations                    []*r4b.Observation
	ObservationDefinitions          []*r4b.ObservationDefinition
	OperationDefinitions            []*r4b.OperationDefinition
	OperationOutcomes               []*r4b.OperationOutcome
	Organizations                   []*r4b.Organization
	OrganizationAffiliations        []*r4b.OrganizationAffiliation
	PackagedProductDefinitions      []*r4b.PackagedProductDefinition
	Patients                        []*r4b.Patient
	PaymentNotices                  []*r4b.PaymentNotice
	PaymentReconciliations          []*r4b.PaymentReconciliation
	Persons                         []*r4b.Person
	PlanDefinitions                 []*r4b.PlanDefinition
	Practitioners                   []*r4b.Practitioner
	PractitionerRoles               []*r4b.PractitionerRole
	Procedures                      []*r4b.Procedure
	Provenances                     []*r4b.Provenance
	Questionnaires                  []*r4b.Questionnaire
	QuestionnaireResponses          []*r4b.QuestionnaireResponse
	RegulatedAuthorizations         []*r4b.RegulatedAuthorization
	RelatedPersons                  []*r4b.RelatedPerson
	RequestGroups                   []*r4b.RequestGroup
	ResearchDefinitions             []*r4b.ResearchDefinition
	ResearchElementDefinitions      []*r4b.ResearchElementDefinition
	ResearchStudys                  []*r4b.ResearchStudy
	ResearchSubjects                []*r4b.ResearchSubject
	RiskAssessments                 []*r4b.RiskAssessment
	Schedules                       []*r4b.Schedule
	SearchParameters                []*r4b.SearchParameter
	ServiceRequests                 []*r4b.ServiceRequest
	Slots                           []*r4b.Slot
	Specimens                       []*r4b.Specimen
	SpecimenDefinitions             []*r4b.SpecimenDefinition
	StructureDefinitions            []*r4b.StructureDefinition
	StructureMaps                   []*r4b.StructureMap
	Subscriptions                   []*r4b.Subscription
	SubscriptionStatuss             []*r4b.SubscriptionStatus
	SubscriptionTopics              []*r4b.SubscriptionTopic
	Substances                      []*r4b.Substance
	SubstanceDefinitions            []*r4b.SubstanceDefinition
	SupplyDeliverys                 []*r4b.SupplyDelivery
	SupplyRequests                  []*r4b.SupplyRequest
	Tasks                           []*r4b.Task
	TerminologyCapabilitiess        []*r4b.TerminologyCapabilities
	TestReports                     []*r4b.TestReport
	TestScripts                     []*r4b.TestScript
	ValueSets                       []*r4b.ValueSet
	VerificationResults             []*r4b.VerificationResult
	VisionPrescriptions             []*r4b.VisionPrescription
}

func BundleToGroup(bundle *r4b.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r4b.Account:
			grp.Accounts = append(grp.Accounts, res)
		case *r4b.ActivityDefinition:
			grp.ActivityDefinitions = append(grp.ActivityDefinitions, res)
		case *r4b.AdministrableProductDefinition:
			grp.AdministrableProductDefinitions = append(grp.AdministrableProductDefinitions, res)
		case *r4b.AdverseEvent:
			grp.AdverseEvents = append(grp.AdverseEvents, res)
		case *r4b.AllergyIntolerance:
			grp.AllergyIntolerances = append(grp.AllergyIntolerances, res)
		case *r4b.Appointment:
			grp.Appointments = append(grp.Appointments, res)
		case *r4b.AppointmentResponse:
			grp.AppointmentResponses = append(grp.AppointmentResponses, res)
		case *r4b.AuditEvent:
			grp.AuditEvents = append(grp.AuditEvents, res)
		case *r4b.Basic:
			grp.Basics = append(grp.Basics, res)
		case *r4b.Binary:
			grp.Binarys = append(grp.Binarys, res)
		case *r4b.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProducts = append(grp.BiologicallyDerivedProducts, res)
		case *r4b.BodyStructure:
			grp.BodyStructures = append(grp.BodyStructures, res)
		case *r4b.CapabilityStatement:
			grp.CapabilityStatements = append(grp.CapabilityStatements, res)
		case *r4b.CarePlan:
			grp.CarePlans = append(grp.CarePlans, res)
		case *r4b.CareTeam:
			grp.CareTeams = append(grp.CareTeams, res)
		case *r4b.CatalogEntry:
			grp.CatalogEntrys = append(grp.CatalogEntrys, res)
		case *r4b.ChargeItem:
			grp.ChargeItems = append(grp.ChargeItems, res)
		case *r4b.ChargeItemDefinition:
			grp.ChargeItemDefinitions = append(grp.ChargeItemDefinitions, res)
		case *r4b.Citation:
			grp.Citations = append(grp.Citations, res)
		case *r4b.Claim:
			grp.Claims = append(grp.Claims, res)
		case *r4b.ClaimResponse:
			grp.ClaimResponses = append(grp.ClaimResponses, res)
		case *r4b.ClinicalImpression:
			grp.ClinicalImpressions = append(grp.ClinicalImpressions, res)
		case *r4b.ClinicalUseDefinition:
			grp.ClinicalUseDefinitions = append(grp.ClinicalUseDefinitions, res)
		case *r4b.CodeSystem:
			grp.CodeSystems = append(grp.CodeSystems, res)
		case *r4b.Communication:
			grp.Communications = append(grp.Communications, res)
		case *r4b.CommunicationRequest:
			grp.CommunicationRequests = append(grp.CommunicationRequests, res)
		case *r4b.CompartmentDefinition:
			grp.CompartmentDefinitions = append(grp.CompartmentDefinitions, res)
		case *r4b.Composition:
			grp.Compositions = append(grp.Compositions, res)
		case *r4b.ConceptMap:
			grp.ConceptMaps = append(grp.ConceptMaps, res)
		case *r4b.Condition:
			grp.Conditions = append(grp.Conditions, res)
		case *r4b.Consent:
			grp.Consents = append(grp.Consents, res)
		case *r4b.Contract:
			grp.Contracts = append(grp.Contracts, res)
		case *r4b.Coverage:
			grp.Coverages = append(grp.Coverages, res)
		case *r4b.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequests = append(grp.CoverageEligibilityRequests, res)
		case *r4b.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponses = append(grp.CoverageEligibilityResponses, res)
		case *r4b.DetectedIssue:
			grp.DetectedIssues = append(grp.DetectedIssues, res)
		case *r4b.Device:
			grp.Devices = append(grp.Devices, res)
		case *r4b.DeviceDefinition:
			grp.DeviceDefinitions = append(grp.DeviceDefinitions, res)
		case *r4b.DeviceMetric:
			grp.DeviceMetrics = append(grp.DeviceMetrics, res)
		case *r4b.DeviceRequest:
			grp.DeviceRequests = append(grp.DeviceRequests, res)
		case *r4b.DeviceUseStatement:
			grp.DeviceUseStatements = append(grp.DeviceUseStatements, res)
		case *r4b.DiagnosticReport:
			grp.DiagnosticReports = append(grp.DiagnosticReports, res)
		case *r4b.DocumentManifest:
			grp.DocumentManifests = append(grp.DocumentManifests, res)
		case *r4b.DocumentReference:
			grp.DocumentReferences = append(grp.DocumentReferences, res)
		case *r4b.Encounter:
			grp.Encounters = append(grp.Encounters, res)
		case *r4b.Endpoint:
			grp.Endpoints = append(grp.Endpoints, res)
		case *r4b.EnrollmentRequest:
			grp.EnrollmentRequests = append(grp.EnrollmentRequests, res)
		case *r4b.EnrollmentResponse:
			grp.EnrollmentResponses = append(grp.EnrollmentResponses, res)
		case *r4b.EpisodeOfCare:
			grp.EpisodeOfCares = append(grp.EpisodeOfCares, res)
		case *r4b.EventDefinition:
			grp.EventDefinitions = append(grp.EventDefinitions, res)
		case *r4b.Evidence:
			grp.Evidences = append(grp.Evidences, res)
		case *r4b.EvidenceReport:
			grp.EvidenceReports = append(grp.EvidenceReports, res)
		case *r4b.EvidenceVariable:
			grp.EvidenceVariables = append(grp.EvidenceVariables, res)
		case *r4b.ExampleScenario:
			grp.ExampleScenarios = append(grp.ExampleScenarios, res)
		case *r4b.ExplanationOfBenefit:
			grp.ExplanationOfBenefits = append(grp.ExplanationOfBenefits, res)
		case *r4b.FamilyMemberHistory:
			grp.FamilyMemberHistorys = append(grp.FamilyMemberHistorys, res)
		case *r4b.Flag:
			grp.Flags = append(grp.Flags, res)
		case *r4b.Goal:
			grp.Goals = append(grp.Goals, res)
		case *r4b.GraphDefinition:
			grp.GraphDefinitions = append(grp.GraphDefinitions, res)
		case *r4b.Group:
			grp.Groups = append(grp.Groups, res)
		case *r4b.GuidanceResponse:
			grp.GuidanceResponses = append(grp.GuidanceResponses, res)
		case *r4b.HealthcareService:
			grp.HealthcareServices = append(grp.HealthcareServices, res)
		case *r4b.ImagingStudy:
			grp.ImagingStudys = append(grp.ImagingStudys, res)
		case *r4b.Immunization:
			grp.Immunizations = append(grp.Immunizations, res)
		case *r4b.ImmunizationEvaluation:
			grp.ImmunizationEvaluations = append(grp.ImmunizationEvaluations, res)
		case *r4b.ImmunizationRecommendation:
			grp.ImmunizationRecommendations = append(grp.ImmunizationRecommendations, res)
		case *r4b.ImplementationGuide:
			grp.ImplementationGuides = append(grp.ImplementationGuides, res)
		case *r4b.Ingredient:
			grp.Ingredients = append(grp.Ingredients, res)
		case *r4b.InsurancePlan:
			grp.InsurancePlans = append(grp.InsurancePlans, res)
		case *r4b.Invoice:
			grp.Invoices = append(grp.Invoices, res)
		case *r4b.Library:
			grp.Librarys = append(grp.Librarys, res)
		case *r4b.Linkage:
			grp.Linkages = append(grp.Linkages, res)
		case *r4b.List:
			grp.Lists = append(grp.Lists, res)
		case *r4b.Location:
			grp.Locations = append(grp.Locations, res)
		case *r4b.ManufacturedItemDefinition:
			grp.ManufacturedItemDefinitions = append(grp.ManufacturedItemDefinitions, res)
		case *r4b.Measure:
			grp.Measures = append(grp.Measures, res)
		case *r4b.MeasureReport:
			grp.MeasureReports = append(grp.MeasureReports, res)
		case *r4b.Media:
			grp.Medias = append(grp.Medias, res)
		case *r4b.Medication:
			grp.Medications = append(grp.Medications, res)
		case *r4b.MedicationAdministration:
			grp.MedicationAdministrations = append(grp.MedicationAdministrations, res)
		case *r4b.MedicationDispense:
			grp.MedicationDispenses = append(grp.MedicationDispenses, res)
		case *r4b.MedicationKnowledge:
			grp.MedicationKnowledges = append(grp.MedicationKnowledges, res)
		case *r4b.MedicationRequest:
			grp.MedicationRequests = append(grp.MedicationRequests, res)
		case *r4b.MedicationStatement:
			grp.MedicationStatements = append(grp.MedicationStatements, res)
		case *r4b.MedicinalProductDefinition:
			grp.MedicinalProductDefinitions = append(grp.MedicinalProductDefinitions, res)
		case *r4b.MessageDefinition:
			grp.MessageDefinitions = append(grp.MessageDefinitions, res)
		case *r4b.MessageHeader:
			grp.MessageHeaders = append(grp.MessageHeaders, res)
		case *r4b.MolecularSequence:
			grp.MolecularSequences = append(grp.MolecularSequences, res)
		case *r4b.NamingSystem:
			grp.NamingSystems = append(grp.NamingSystems, res)
		case *r4b.NutritionOrder:
			grp.NutritionOrders = append(grp.NutritionOrders, res)
		case *r4b.NutritionProduct:
			grp.NutritionProducts = append(grp.NutritionProducts, res)
		case *r4b.Observation:
			grp.Observations = append(grp.Observations, res)
		case *r4b.ObservationDefinition:
			grp.ObservationDefinitions = append(grp.ObservationDefinitions, res)
		case *r4b.OperationDefinition:
			grp.OperationDefinitions = append(grp.OperationDefinitions, res)
		case *r4b.OperationOutcome:
			grp.OperationOutcomes = append(grp.OperationOutcomes, res)
		case *r4b.Organization:
			grp.Organizations = append(grp.Organizations, res)
		case *r4b.OrganizationAffiliation:
			grp.OrganizationAffiliations = append(grp.OrganizationAffiliations, res)
		case *r4b.PackagedProductDefinition:
			grp.PackagedProductDefinitions = append(grp.PackagedProductDefinitions, res)
		case *r4b.Patient:
			grp.Patients = append(grp.Patients, res)
		case *r4b.PaymentNotice:
			grp.PaymentNotices = append(grp.PaymentNotices, res)
		case *r4b.PaymentReconciliation:
			grp.PaymentReconciliations = append(grp.PaymentReconciliations, res)
		case *r4b.Person:
			grp.Persons = append(grp.Persons, res)
		case *r4b.PlanDefinition:
			grp.PlanDefinitions = append(grp.PlanDefinitions, res)
		case *r4b.Practitioner:
			grp.Practitioners = append(grp.Practitioners, res)
		case *r4b.PractitionerRole:
			grp.PractitionerRoles = append(grp.PractitionerRoles, res)
		case *r4b.Procedure:
			grp.Procedures = append(grp.Procedures, res)
		case *r4b.Provenance:
			grp.Provenances = append(grp.Provenances, res)
		case *r4b.Questionnaire:
			grp.Questionnaires = append(grp.Questionnaires, res)
		case *r4b.QuestionnaireResponse:
			grp.QuestionnaireResponses = append(grp.QuestionnaireResponses, res)
		case *r4b.RegulatedAuthorization:
			grp.RegulatedAuthorizations = append(grp.RegulatedAuthorizations, res)
		case *r4b.RelatedPerson:
			grp.RelatedPersons = append(grp.RelatedPersons, res)
		case *r4b.RequestGroup:
			grp.RequestGroups = append(grp.RequestGroups, res)
		case *r4b.ResearchDefinition:
			grp.ResearchDefinitions = append(grp.ResearchDefinitions, res)
		case *r4b.ResearchElementDefinition:
			grp.ResearchElementDefinitions = append(grp.ResearchElementDefinitions, res)
		case *r4b.ResearchStudy:
			grp.ResearchStudys = append(grp.ResearchStudys, res)
		case *r4b.ResearchSubject:
			grp.ResearchSubjects = append(grp.ResearchSubjects, res)
		case *r4b.RiskAssessment:
			grp.RiskAssessments = append(grp.RiskAssessments, res)
		case *r4b.Schedule:
			grp.Schedules = append(grp.Schedules, res)
		case *r4b.SearchParameter:
			grp.SearchParameters = append(grp.SearchParameters, res)
		case *r4b.ServiceRequest:
			grp.ServiceRequests = append(grp.ServiceRequests, res)
		case *r4b.Slot:
			grp.Slots = append(grp.Slots, res)
		case *r4b.Specimen:
			grp.Specimens = append(grp.Specimens, res)
		case *r4b.SpecimenDefinition:
			grp.SpecimenDefinitions = append(grp.SpecimenDefinitions, res)
		case *r4b.StructureDefinition:
			grp.StructureDefinitions = append(grp.StructureDefinitions, res)
		case *r4b.StructureMap:
			grp.StructureMaps = append(grp.StructureMaps, res)
		case *r4b.Subscription:
			grp.Subscriptions = append(grp.Subscriptions, res)
		case *r4b.SubscriptionStatus:
			grp.SubscriptionStatuss = append(grp.SubscriptionStatuss, res)
		case *r4b.SubscriptionTopic:
			grp.SubscriptionTopics = append(grp.SubscriptionTopics, res)
		case *r4b.Substance:
			grp.Substances = append(grp.Substances, res)
		case *r4b.SubstanceDefinition:
			grp.SubstanceDefinitions = append(grp.SubstanceDefinitions, res)
		case *r4b.SupplyDelivery:
			grp.SupplyDeliverys = append(grp.SupplyDeliverys, res)
		case *r4b.SupplyRequest:
			grp.SupplyRequests = append(grp.SupplyRequests, res)
		case *r4b.Task:
			grp.Tasks = append(grp.Tasks, res)
		case *r4b.TerminologyCapabilities:
			grp.TerminologyCapabilitiess = append(grp.TerminologyCapabilitiess, res)
		case *r4b.TestReport:
			grp.TestReports = append(grp.TestReports, res)
		case *r4b.TestScript:
			grp.TestScripts = append(grp.TestScripts, res)
		case *r4b.ValueSet:
			grp.ValueSets = append(grp.ValueSets, res)
		case *r4b.VerificationResult:
			grp.VerificationResults = append(grp.VerificationResults, res)
		case *r4b.VisionPrescription:
			grp.VisionPrescriptions = append(grp.VisionPrescriptions, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
